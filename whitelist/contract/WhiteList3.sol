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
    1. Add Events for Adding\Removing admins and trustees
    2. Save List of voting with status (?)
    3. Fix public modifiers

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
        //admins.push(_admin);
        //trueAdmin[_admin] = true;
        createVoting(msg.sender, AdminAction.ADDADMIN, _admin, VoteConsensus.ALL);
    }

    function removeAdmin(address _admin) callOnlyAdmin adminExists(_admin) {
        //there must be at least one admin
        require(admins.length > 2);
        //FIXME may be exists voting for delete admin so need to prohibit removing
        if (_admin == msg.sender) {
            //admin can remove self without voting
            trueAdmin[_admin] = false;
            removeFromArray(admins, _admin);
        }
        else {
            //FIXME by voting
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
        //TODO by voting
        trustees.push(_addr);
        trueTrustee[_addr] = true;
    }

    function removeTrusted(address _addr) callOnlyAdmin doesNotAdmin(_addr) {
        if (msg.sender == _addr) {//trustee may remove self without voting
            trueTrustee[_addr] = false;
            removeFromArray(trustees, _addr);
            //remove PGP
            removeKey(_addr);
        }
        else {
            //FIXME by voting
            //FIXME trustees.removeElement(_addr);
            //TODO remove PGP
        }
    }

    /***
     * PGP Section
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
    address participant;
    }

    //TODO must be list of proposal
    Proposal public actual;

    //FIXME onlyAdministration - why doesn't work ??
    function createVoting(address owner, AdminAction action, address participant, VoteConsensus consensus) {
        //no sumultaneous votings
        require(actual.voting == 0x0);
        Voting voting = new Voting(owner, admins, consensus);
        actual = Proposal({
        voting : voting,
        action : action,
        participant : participant
        });
    }

    function callback(VoteResult tally) {
        require(msg.sender == actual.voting);
        if (tally == VoteResult.SUCCESS) {//SUCCESS
            //TODO case action
            if (actual.action == AdminAction.ADDADMIN) {
                //FIXME add check is participant trustee yet and not admin
                admins.push(actual.participant);
                trueAdmin[actual.participant] = true;
                //disable voting
                actual.voting = 0x0;
                actual.action = AdminAction.UNKNOWN;
                actual.participant = 0x0;
            }
        }
        else if (tally == VoteResult.FAIL || tally == VoteResult.CANCEL) {//FAIL or CANCEL
            //remove voting from actual
            actual.voting = 0x0;
            actual.action = AdminAction.UNKNOWN;
            actual.participant = 0x0;
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
    }
}


contract Voting is Constants {

    /****
    TODO List

    1. Add Events
    2. Add Timeouts
    3. Disable multithreading
    4. Fix public modifiers
    5. Constant and enums to external contract
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
    uint vote; // FIXME enum ?
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
        //FIXME check length
        for (uint i = 0; i < _voters.length; i++) {
            // init voter state
            if (!eligible[_voters[i]]) {
                eligible[_voters[i]] = true;
                addresses.push(_voters[i]);
                addressid[_voters[i]] = totaleligible;
                voters[totaleligible] = Voter({
                addr : _voters[i],
                vote : 0
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
            voters[addressid[msg.sender]].vote = uint(VoteResult.SUCCESS);
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
            voters[addressid[msg.sender]].vote = uint(VoteResult.FAIL);
            totalvoted += 1;
            return true;
        }
        return false;
    }

    function countResult() inState(VoteState.VOTE) returns (VoteResult) {
        //all must vote
        assert(totaleligible == totalvoted);
        uint pro = 0;
        uint contra = 0;
        //TODO add timeout rule
        for (uint i = 0; i < totaleligible; i++) {
            //  Check all votes have been cast
            assert(votecast[voters[i].addr]);
            //sum vote
            if (voters[i].vote == uint(VoteResult.SUCCESS)) {
                pro += 1;
            }
            if (voters[i].vote == uint(VoteResult.FAIL)) {
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
        //FIXME make normal callback
        Administration(administration).callback(tally);
        return tally;
    }

    function cancelVoting() inState(VoteState.VOTE) {
        require(msg.sender == owner);
        state = VoteState.CANCELED;
        tally = VoteResult.CANCEL;
        //FIXME send callback to whitelist to remove this vote
        Administration(administration).callback(tally);
    }

}
