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
    CANCEL
    }

    enum VoteState {
    VOTE,
    FINISHED,
    CANCELED
    }

    enum AdminAction {
    UNKNOWN,
    ADDADMIN,
    REMOVEADMIN,
    ADDTRUSTEE,
    REMOVETRUSTEE
    }
}

contract Administration is Constants {

    /***
     *
    TODO LIST:
    1. refactor "removeXXX" voting costructor - candidate must be removed from voters
    2. Add Events for Adding\Removing admins and trustees
    3. Save List of voting with status (?)
    4. Fix public modifiers
    5. PGP - realize replacing of another's key via voting
    6. PGP  - save key's add/update date (???)
    7. Make list of proposals

     * **/

    address[]public trustees;

    address[]public admins;

    mapping (address => bool)public trueAdmin;

    mapping (address => bool)public trueTrustee;

    function Administration(address[] _admins) {
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

    function isAdmin(address _addr) constant returns (bool) {
        return trueAdmin[_addr];
    }

    function adminSize() constant returns (uint) {
        return admins.length;
    }

    modifier callOnlyAdmin() {
        require(isAdmin(msg.sender));
        _;
    }

    modifier onlyAdministration() {
        require(msg.sender == address(this));
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

    function addAdmin(address _admin) callOnlyAdmin trusteeExists(_admin) notNull(_admin) doesNotAdmin(_admin) {
        createVoting(msg.sender, AdminAction.ADDADMIN, _admin, VoteConsensus.ALL);
    }

    function removeAdmin(address _admin) callOnlyAdmin adminExists(_admin) {
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
            createVoting(msg.sender, AdminAction.REMOVEADMIN, _admin, VoteConsensus.ALL);
        }
    }

    /********
    TRUSTEES Section
     ********/
    function isTrusted(address _addr) constant returns (bool) {
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

    function addTrustee(address _addr) callOnlyAdmin doesNotTrustee(_addr) {
        createVoting(msg.sender, AdminAction.ADDTRUSTEE, _addr, VoteConsensus.HALF);
    }

    function removeTrusted(address _addr) doesNotAdmin(_addr) {
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
            createVoting(msg.sender, AdminAction.REMOVETRUSTEE, _addr, VoteConsensus.HALF);
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
    function addKey(bytes32 key) callOnlyTrustee {
        //add key to keys
        keys[msg.sender] = key;
    }

    function getKey(address trustee) constant returns (bytes32) {
        return keys[trustee];
    }

    //FIXME check key's existance
    function removeKey(address _addr) onlyAdministration {
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

    //FIXME onlyAdministration - why doesn't work ??
    function createVoting(address owner, AdminAction action, address contender, VoteConsensus consensus) {
        //no sumultaneous votings
        require(actual.voting == 0x0);
        Voting voting = new Voting(owner, admins, consensus);
        actual = Proposal({
        voting : voting,
        action : action,
        contender : contender
        });
    }

    function callback(VoteResult tally) {
        //TODO check if voting is created by this administration
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
        else if (tally == VoteResult.FAIL || tally == VoteResult.CANCEL) {//FAIL or CANCEL
            //remove voting from actual; TODO save result in archive
            actual.voting = 0x0;
            actual.action = AdminAction.UNKNOWN;
            actual.contender = 0x0;
        }
    }

    /****
    UTILS
     ******/

    function removeFromArray(address[]storage _array, address _addr) internal returns (address[]) {
        for (uint i = 0; i < _array.length - 1; i++) {
            if (_array[i] == _addr) {
                _array[i] = _array[_array.length - 1];
                break;
            }
        }
        _array.length -= 1;
        return _array;
    }
}

contract Voting is Constants {

    /****
    TODO List

    1. Add Events
    2. Add Timeouts
    3. Disable multithreading
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

    //constructor
    function Voting(address _owner, address[] _voters, VoteConsensus consensus) {
        //check length
        assert(_voters.length > 0);
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
        state = VoteState.VOTE;
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
    }

    function accept() inState(VoteState.VOTE) returns (bool) {
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

    function reject() inState(VoteState.VOTE) returns (bool) {
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

    function countResult() inState(VoteState.VOTE) returns (VoteResult) {
        //all must vote
        require(totaleligible == totalvoted);
        uint pro = 0;
        uint contra = 0;
        for (uint i = 0; i < totaleligible; i++) {
            //  Check all votes have been cast
            require(votecast[voters[i].addr]);
            //sum vote
            if (voters[i].vote == VoteResult.SUCCESS) {
                pro += 1;
            }
            if (voters[i].vote == VoteResult.FAIL) {
                contra += 1;
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
        return tally;
    }

    function cancelVoting() inState(VoteState.VOTE) {
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

}
