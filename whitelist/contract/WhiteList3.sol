pragma solidity ^ 0.4.16;

contract Constants {

	enum VoteConsensus {
		ALL,
		HALF
	}

	enum VoteResult {
		UNKNOWN,
		SUCCESS,
		FAIL,
		CANCEL,
		TIMEOUT
	}

	enum VoteState {
		VOTED,
		FINISHED,
		CANCELED,
		TIMEOUTED
	}

	enum AdminAction {
		UNKNOWN,
		ADDADMIN,
		REMOVEADMIN,
		ADDTRUSTEE,
		REMOVETRUSTEE
	}

}

contract MutexProtected {
	// Mutex contract on all state mutating external/public functions for protecting against
	// Re-entry and `Solar Storm` attacks
	modifier isMutexed() {
		require(!mutex);
		// Exclusion already set so bail.
		mutex = true;
		// Set exclusion and continue with function code
		_;
		delete mutex;
		// release contract from exclusion
		return;
		// functions require single exit and return parameters
	}
	bool mutex; // the exclusion state variable
}

contract Administration is Constants {

	/***
	 *
	TODO LIST:
	2. Add Events for Adding\Removing admins and trustees
	4. Fix public modifiers for contract variables
	5. PKI - realize replacing of another's key via voting
	6. PKI  - save key's add/update date (???)
	7. Maybe refactor remove and add functions - separate self and others cases in distinct functions?
	10. Manipulate timeout's values through votings ?
	11. Realize algorithm of migration to other version of Admininstration
	13. Do recovery of key not only by voting but with revealing hash of key
	14. Clear archive of votings by time
	15. Need remove title for removed trustee ???


	Won't Fixed :

	1. Return voting contract address in CRUD functions
	16. How to participant add self to trustee ? through voting ?  - by admin

	 * **/

	address[]public trustees;

	address[]public admins;

	mapping(address => string)public titles;

	mapping(address => bool)public trueAdmin;

	mapping(address => bool)public trueTrustee;

	function Administration(address[]_admins)public {
		//initialize admins and add their to trustees
		for (uint i = 0; i < _admins.length; i++) {
			if (!isAdmin(_admins[i]) && _admins[i] != 0) {
				trueAdmin[_admins[i]] = true;
				admins.push(_admins[i]);
				trueTrustee[_admins[i]] = true;
				trustees.push(_admins[i]);
			}
		}
	}

	/***********
	ADMIN section
	 *************/

	function isAdmin(address _addr)constant public returns(bool) {
		return trueAdmin[_addr];
	}

	function adminSize()constant public returns(uint) {
		return admins.length;
	}

	modifier callOnlyAdmin() {
		require(isAdmin(msg.sender));
		_;
	}

	modifier doesNotAdmin(address _admin) {
		require(!isAdmin(_admin));
		_;
	}

	modifier adminExists(address _admin) {
		require(isAdmin(_admin));
		_;
	}

	modifier notNull(address _address) {
		require(_address != 0);
		_;
	}

	function addAdmin(address _admin)callOnlyAdmin trusteeExists(_admin)notNull(_admin)doesNotAdmin(_admin)public {
		createVoting(msg.sender, AdminAction.ADDADMIN, _admin, admins, VoteConsensus.ALL, ADDADMINDURATION);
	}

	function removeAdmin(address _admin)callOnlyAdmin adminExists(_admin)public {
		//there must be at least one admin
		require(admins.length > 1);
		if (_admin == msg.sender) {
			//need to prohibit removing if there is actual voting
			if (actual.voting == 0x0) {
				//admin can remove self without voting
				trueAdmin[_admin] = false;
				admins = removeFromArray(admins, _admin);
			}
		} else {
			//remove contender from voters
			address[]memory voters = exCopyArray(admins, _admin);
			createVoting(msg.sender, AdminAction.REMOVEADMIN, _admin, voters, VoteConsensus.ALL, REMOVEADMINDURATION);
		}
	}

	/********
	TRUSTEES Section
	 ********/
	function isTrusted(address _addr)constant public returns(bool) {
		return trueTrustee[_addr];
	}

	modifier callOnlyTrustee() {
		require(isTrusted(msg.sender));
		_;
	}

	modifier trusteeExists(address _addr) {
		require(isTrusted(_addr));
		_;
	}

	modifier doesNotTrustee(address _addr) {
		require(!isTrusted(_addr));
		_;
	}

	function addTrustee(address _addr)callOnlyAdmin doesNotTrustee(_addr)public {
		createVoting(msg.sender, AdminAction.ADDTRUSTEE, _addr, admins, VoteConsensus.HALF, ADDTRUSTEEDURATION);
	}

	function removeTrusted(address _addr)doesNotAdmin(_addr)public {
		if (msg.sender == _addr) { //trustee may remove himself without voting
			trueTrustee[_addr] = false;
			trustees = removeFromArray(trustees, _addr);
			//remove PKI
			removeKey(_addr);
		} else {
			//check if caller is admin
			require(isAdmin(msg.sender));
			//by voting
			createVoting(msg.sender, AdminAction.REMOVETRUSTEE, _addr, admins, VoteConsensus.HALF, REMOVETRUSTEEDURATION);
		}
	}

	function getTitle(address _trustee)public constant returns(string) {
		return titles[_trustee];
	}

	function setTitle(string _title)public callOnlyTrustee {
		assert(utfStringLength(_title) > 0);
		titles[msg.sender] = _title;
	}

	function removeTitile()public callOnlyTrustee {
		delete titles[msg.sender];
	}

	/***
	 * PKI Section
	 *
	 *
	 *
	 **/

	struct Key {
		uint date;
		bytes32 value;
		bool isActive;
	}

	//PKI - FIXME define length of key
	mapping(address => Key)public activeKeys;
	mapping(address => Key[])public archiveKeys;

	//PKI Section
	function addOrUpdateKey(bytes32 _key)public callOnlyTrustee {
		//TODO must separate self and others cases
		//check does active key exist
		if (activeKeys[msg.sender].isActive) {
			//save current key to archive
			Key memory oldKey = Key({
					date: activeKeys[msg.sender].date,
					value: activeKeys[msg.sender].value,
					isActive: false
				});
			archiveKeys[msg.sender].push(oldKey);

		}
		//make new Key
		activeKeys[msg.sender] = Key({
				date: now,
				value: _key,
				isActive: true
			});
	}

	function getKey(address trustee)public constant returns(bool, bytes32) {
		return (activeKeys[trustee].isActive, activeKeys[trustee].value);
	}

	function removeOwnKey()public callOnlyTrustee {
		if (activeKeys[msg.sender].isActive) { //if keys exists
			delete activeKeys[msg.sender];
		}
	}

	//remove key of participant
	function removeKey(address trustee)private {
		if (activeKeys[trustee].isActive) { //if keys exists
			delete activeKeys[trustee];
		}
	}

	/****
	 * VOTING Section
	 *
	 ***/
	//vote constants

	uint constant ADDADMINDURATION = 120; //2 minutes
	uint constant REMOVEADMINDURATION = 120; //2 minutes
	uint constant ADDTRUSTEEDURATION = 60; // 1 minutes
	uint constant REMOVETRUSTEEDURATION = 60; // 1 minutes


	struct Proposal {
		address voting;
		AdminAction action;
		address contender;
	}

	Proposal public actual;

	struct ArchiveProposal {
		Proposal proposal;
		uint finishDate;
		VoteResult tally;
	}

	ArchiveProposal[]archiveProposals;

	function createVoting(address owner, AdminAction action, address contender, address[]voters, VoteConsensus consensus, uint durationInSeconds)private {
		//no sumultaneous votings
		require(actual.voting == 0x0);
		//create votin contract
		Voting voting = new Voting(owner, voters, consensus, durationInSeconds);
		actual = Proposal({
				voting: voting,
				action: action,
				contender: contender
			});
	}

	function callback(VoteResult tally)public {
		//check if voting is created by this administration
		require(msg.sender == actual.voting);
		ArchiveProposal memory arch = ArchiveProposal({
				proposal: actual,
				finishDate: now,
				tally: tally
			});
		if (tally == VoteResult.SUCCESS) { //SUCCESS
			if (actual.action == AdminAction.ADDADMIN) {
				//check is contender trustee yet and not admin
				assert(isTrusted(actual.contender) && !isAdmin(actual.contender));
				admins.push(actual.contender);
				trueAdmin[actual.contender] = true;
			}
			if (actual.action == AdminAction.REMOVEADMIN) {
				//check if exists one admin yet
				assert(admins.length > 1);
				admins = removeFromArray(admins, actual.contender);
				delete trueAdmin[actual.contender];
			}
			if (actual.action == AdminAction.ADDTRUSTEE) {
				//check is contender trustee
				assert(!isTrusted(actual.contender));
				//save contender in trusted
				trustees.push(actual.contender);
				trueTrustee[actual.contender] = true;
			}
			if (actual.action == AdminAction.REMOVETRUSTEE) {
				trustees = removeFromArray(trustees, actual.contender);
				delete trueTrustee[actual.contender];
				//remove PKI
				removeKey(actual.contender);
			}
			//disable actual voting
			actual.voting = 0x0;
			actual.action = AdminAction.UNKNOWN;
			actual.contender = 0x0;
			//archiveVoting
			archiveProposals.push(arch);
		} else if (tally == VoteResult.FAIL || tally == VoteResult.CANCEL || tally == VoteResult.TIMEOUT) { //FAIL or CANCEL or TIMEOUT
			//remove voting from actual;
			actual.voting = 0x0;
			actual.action = AdminAction.UNKNOWN;
			actual.contender = 0x0;
			//archiveVoting
			archiveProposals.push(arch);
		}
	}

	/****
	UTILS
	 ******/

	function removeFromArray(address[]storage _array, address _addr)private returns(address[]) {
		for (uint i = 0; i < _array.length - 1; i++) {
			if (_array[i] == _addr) {
				_array[i] = _array[_array.length - 1];
				break;
			}
		}
		_array.length -= 1;
		return _array;
	}

	function exCopyArray(address[]_array, address _addr)private pure returns(address[]) {
		address[]memory dArray = new address[](_array.length - 1);
		uint count = 0;
		for (uint i = 0; i < _array.length; i++) {
			if (_array[i] != _addr) {
				dArray[count++] = _array[i];
			}
		}
		return dArray;
	}

	function utfStringLength(string str)internal pure
	returns(uint length) {
		uint i = 0;
		bytes memory string_rep = bytes(str);
		while (i < string_rep.length) {
			if (string_rep[i] >> 7 == 0)
				i += 1;
			else if (string_rep[i] >> 5 == 0x6)
				i += 2;
			else if (string_rep[i] >> 4 == 0xE)
				i += 3;
			else if (string_rep[i] >> 3 == 0x1E)
				i += 4;
			else
				//For safety
				i += 1;

			length++;
		}
	}
}

contract Voting is Constants, MutexProtected {

	/****
	TODO List

	1. Add Events
	4. Fix public modifiers
	5. Check if calling "callback" function returns true
	 *****/

	VoteState public state;

	//addresses of voters and their indexes
	address[]public addresses;

	mapping(address => uint)public addressid;

	mapping(uint => Voter)public voters;

	mapping(address => bool)public eligible; // White list of addresses allowed to vote
	mapping(address => bool)public votecast; // Address voted?

	struct Voter {
		address addr;
		VoteResult vote; //enum ?
	}
	//counters
	uint public totaleligible;

	uint public totalvoted;

	VoteResult public tally;

	address public owner;

	address public administration;

	modifier inState(VoteState s) {
		require(state == s);
		_;
	}

	uint public neededAmount;

	uint public timeoutTime;

	//constructor
	function Voting(address _owner, address[]_voters, VoteConsensus consensus, uint durationInSeconds)public {
		//check length
		assert(_voters.length > 0);
		//check valid duration
		assert(durationInSeconds > 0);
		for (uint i = 0; i < _voters.length; i++) {
			// init voter state
			if (!eligible[_voters[i]]) {
				eligible[_voters[i]] = true;
				addresses.push(_voters[i]);
				addressid[_voters[i]] = totaleligible;
				voters[totaleligible] = Voter({
						addr: _voters[i],
						vote: VoteResult.UNKNOWN
					});
				votecast[_voters[i]] = false;
				totaleligible += 1;
			}
		}
		administration = msg.sender;
		owner = _owner;
		state = VoteState.VOTED;
		tally = VoteResult.UNKNOWN;
		if (consensus == VoteConsensus.ALL) {
			neededAmount = totaleligible;
		} else if (consensus == VoteConsensus.HALF) {
			neededAmount = (totaleligible / 2) + 1;
		} else {
			//ALL by default;
			neededAmount == totaleligible;
		}
		//calc timeout time
		timeoutTime = now + durationInSeconds;
	}

	function getRemainsTime()public constant returns(uint) {
		if (timeoutTime >= now) {
			return timeoutTime - now;
		}
		return 0;
	}

	function accept()inState(VoteState.VOTED)public returns(bool) {
		//check participant can vote and hasn't already voted
		//TODO maybe assert to withdraw gas
		if (eligible[msg.sender] && !votecast[msg.sender]) {
			votecast[msg.sender] = true;
			voters[addressid[msg.sender]].vote = VoteResult.SUCCESS;
			totalvoted += 1;
			return true;
		}
		return false;
	}

	function reject()inState(VoteState.VOTED)public returns(bool) {
		//check participant can vote and hasn't already voted
		//TODO maybe assert to withdraw gas
		if (eligible[msg.sender] && !votecast[msg.sender]) {
			votecast[msg.sender] = true;
			voters[addressid[msg.sender]].vote = VoteResult.FAIL;
			totalvoted += 1;
			return true;
		}
		return false;
	}

	function countResult()inState(VoteState.VOTED)isMutexed public {

		bool timeIsGone = (getRemainsTime() == 0);
		bool allVote = (totalvoted == totaleligible);
		bool enoughVoted = (totalvoted >= neededAmount);

		if (timeIsGone) {
			if (enoughVoted) { // if the voting time has expired, but the votes are enough to get the result
				//calc result
				calcResult();
			} else { // cancel voting by timeout
				//set state and result
				state = VoteState.TIMEOUTED;
				tally = VoteResult.TIMEOUT;
				//send callback to whitelist to remove this vote
				Administration(administration).callback(tally);
			}
		} else { //voting is still going
			if (allVote) { // get all votes - nothing will change
				//calc result
				calcResult();
			}
		}
	}

	function cancelVoting()inState(VoteState.VOTED)isMutexed public {
		//only owner can stop this voting
		require(msg.sender == owner);
		//save this vote in vote array
		if (!votecast[msg.sender]) {
			//mark as voted
			votecast[msg.sender] = true;
			totalvoted += 1;
		}
		voters[addressid[msg.sender]].vote = VoteResult.CANCEL;
		//set state and result
		state = VoteState.CANCELED;
		tally = VoteResult.CANCEL;
		//send callback to whitelist to remove this vote
		Administration(administration).callback(tally);
	}

	function calcResult()private {
		uint pro = 0;
		uint contra = 0;
		for (uint i = 0; i < totaleligible; i++) {
			//  Sum only vote have been cast
			if (votecast[voters[i].addr]) {
				//sum vote
				if (voters[i].vote == VoteResult.SUCCESS) {
					pro += 1;
				}
				if (voters[i].vote == VoteResult.FAIL) {
					contra += 1;
				}
			}
		}
		state = VoteState.FINISHED;
		//check result
		if (pro >= neededAmount) {
			tally = VoteResult.SUCCESS;
		} else {
			tally = VoteResult.FAIL;
		}
		//make normal callback
		Administration(administration).callback(tally);
	}

}
