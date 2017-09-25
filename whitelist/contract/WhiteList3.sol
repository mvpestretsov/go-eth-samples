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
    3. Save List of voting with status (?)
    4. Fix public modifiers for contract variables
    5. PGP - realize replacing of another's key via voting
    6. PGP  - save key's add/update date (???)
    7. Maybe refactor remove and add functions - separate self and others cases in distinct functions?
    8. Return voting contract address in CRUD functions
    9. Maybe return in calback boolean value and then check it in Voting  ?
    10. Extract timeouts of voting in constants or variables
     * **/

    address[]public trustees;

    address[]public admins;

    mapping (address => bool)public trueAdmin;

    mapping (address => bool)public trueTrustee;

    function Administration(address[] _admins) public {
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

    function isAdmin(address _addr) constant public returns (bool) {
        return trueAdmin[_addr];
    }

    function adminSize() constant public returns (uint) {
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

    function addAdmin(address _admin) callOnlyAdmin trusteeExists(_admin) notNull(_admin) doesNotAdmin(_admin) public {
        createVoting(msg.sender, AdminAction.ADDADMIN, _admin, admins, VoteConsensus.ALL);
    }

    function removeAdmin(address _admin) callOnlyAdmin adminExists(_admin) public {
        //there must be at least one admin
        require(admins.length > 1);
        if (_admin == msg.sender) {
            //need to prohibit removing if there is actual voting
            if (actual.voting == 0x0) {
                //admin can remove self without voting
                trueAdmin[_admin] = false;
                admins = removeFromArray(admins, _admin);
            }
        }
        else {
            //remove contender from voters
            address[]memory voters = exCopyArray(admins, _admin);
            createVoting(msg.sender, AdminAction.REMOVEADMIN, _admin, voters, VoteConsensus.ALL);
        }
    }

    /********
    TRUSTEES Section
     ********/
    function isTrusted(address _addr) constant public returns (bool) {
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

    function addTrustee(address _addr) callOnlyAdmin doesNotTrustee(_addr) public {
        createVoting(msg.sender, AdminAction.ADDTRUSTEE, _addr, admins, VoteConsensus.HALF);
    }

    function removeTrusted(address _addr) doesNotAdmin(_addr) public {
        if (msg.sender == _addr) {//trustee may remove himself without voting
            trueTrustee[_addr] = false;
            trustees = removeFromArray(trustees, _addr);
            //remove PGP
            removeKey(_addr);
        }
        else {
            //check if caller is admin
            require(isAdmin(msg.sender));
            //by voting
            createVoting(msg.sender, AdminAction.REMOVETRUSTEE, _addr, admins, VoteConsensus.HALF);
        }
    }

    /***
     * PGP Section
     *
     *
     *
     **/

    //PGP - FIXME define length of key
    mapping (address => bytes32)keys;

    //PGP Section
    function addKey(bytes32 key) public callOnlyTrustee {
        //add key to keys
        keys[msg.sender] = key;
    }

    function getKey(address trustee) public constant returns (bytes32) {
        return keys[trustee];
    }

    //FIXME check key's existance
    function removeKey(address _addr) private {
        delete keys[_addr];
    }

    /****
     * VOTING Section
     *
     ***/

    struct Proposal {
    address voting;
    AdminAction action;
    address contender;
    }

    Proposal public actual;

    function createVoting(address owner, AdminAction action, address contender, address[] voters, VoteConsensus consensus) private {
        //no sumultaneous votings
        require(actual.voting == 0x0);
        //create votin contract
        //FIXME set timeout constants
        Voting voting = new Voting(owner, voters, consensus, 60);
        actual = Proposal({
        voting : voting,
        action : action,
        contender : contender
        });
    }

    function callback(VoteResult tally) public {
        //check if voting is created by this administration
        require(msg.sender == actual.voting);
        if (tally == VoteResult.SUCCESS) {//SUCCESS
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
                //remove PGP
                removeKey(actual.contender);
            }
            //disable voting
            actual.voting = 0x0;
            actual.action = AdminAction.UNKNOWN;
            actual.contender = 0x0;
        }
        else if (tally == VoteResult.FAIL || tally == VoteResult.CANCEL || tally == VoteResult.TIMEOUT) {//FAIL or CANCEL
            //remove voting from actual; TODO save result in archive
            actual.voting = 0x0;
            actual.action = AdminAction.UNKNOWN;
            actual.contender = 0x0;
        }
    }

    /****
    UTILS
     ******/

    function removeFromArray(address[]storage _array, address _addr) private returns (address[]) {
        for (uint i = 0; i < _array.length - 1; i++) {
            if (_array[i] == _addr) {
                _array[i] = _array[_array.length - 1];
                break;
            }
        }
        _array.length -= 1;
        return _array;
    }

    function exCopyArray(address[] _array, address _addr) private pure returns (address[]) {
        address[]memory dArray = new address[](_array.length - 1);
        uint count = 0;
        for (uint i = 0; i < _array.length; i++) {
            if (_array[i] != _addr) {
                dArray[count++] = _array[i];
            }
        }
        return dArray;
    }
}


contract Voting is Constants, MutexProtected {

    /****
    TODO List

    1. Add Events

    4. Fix public modifiers

     *****/

    VoteState public state;

    //addresses of voters and their indexes
    address[]public addresses;

    mapping (address => uint)public addressid;

    mapping (uint => Voter)public voters;

    mapping (address => bool)public eligible; // White list of addresses allowed to vote
    mapping (address => bool)public votecast; // Address voted?

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

    uint public timeout;

    //constructor
    function Voting(address _owner, address[] _voters, VoteConsensus consensus, uint durationInSeconds) public {
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
                addr : _voters[i],
                vote : VoteResult.UNKNOWN
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
        }
        else if (consensus == VoteConsensus.HALF) {
            neededAmount = (totaleligible / 2) + 1;
        }
        else {
            //ALL by default;
            neededAmount == totaleligible;
        }
        //calc timeout time
        timeout = now + durationInSeconds;
    }

    function getRemainsTime() public constant returns (uint) {
        if (timeout >= now) {
            return timeout - now;
        }
        return 0;
    }

    function accept() inState(VoteState.VOTED) public returns (bool) {
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

    function reject() inState(VoteState.VOTED) public returns (bool) {
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

    function countResult() inState(VoteState.VOTED) isMutexed public {

        bool timeIsGone = (getRemainsTime() == 0);
        bool allVote = (totalvoted == totaleligible);
        bool enoughVoted = (totalvoted >= neededAmount);

        if (timeIsGone) {
            if (enoughVoted) {
                //calc result
                calcResult();
            }
            else {// cancel voting by timeout
                //set state and result
                state = VoteState.TIMEOUTED;
                tally = VoteResult.TIMEOUT;
                //send callback to whitelist to remove this vote
                Administration(administration).callback(tally);
            }
        }
        else {
            if (allVote) {
                //calc result
                calcResult();
            }
        }
    }

    function cancelVoting() inState(VoteState.VOTED) isMutexed public {
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

    function calcResult() private {
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
        }
        else {
            tally = VoteResult.FAIL;
        }
        //make normal callback
        Administration(administration).callback(tally);
    }

}
