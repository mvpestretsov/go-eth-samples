pragma solidity ^ 0.4.16;

contract WhiteList {

    /*
    Questions:

    1. Q:rules to voting - 50%+1  or smth else ? Maybe admin votes?
    A: Admin must be voted by all admins
    2. Q:if voting is expired and there is no decision - discard voting?
    A:Yes. New admin must be voted all admins.
    3. Q:trusted=>admin ?
    A:No. Trusted is only for receiving info, admins for voting
    4. Fix number of admins for every voting ?
    5. How system will allow changing admin address if he lose his ethereum key + 100% participating ?
    6. ADD KEY in PGP with add to trusted ?
     */

    // biznet originator
    address beginner;
    //proved participants
    mapping (address => bool)trusted;
    //super participants
    mapping (address => bool)admins;

    address[]adminsArray;
    //PGP - FIXME define length of key
    mapping (address => bytes32)keys;

    //constructor - set initial admins + sender to trusted
    function WhiteList(address[] _admins) {
        //save initiator
        beginner = msg.sender;
        //add initiator to admins
        admins[msg.sender] = true;
        //add initiator to trusted
        trusted[msg.sender] = true;
        for (uint i = 0; i < _admins.length; i++) {
            //add admins to admins
            admins[_admins[i]] = true;
            adminsArray.push(_admins[i]);
            //add admins to trusted
            trusted[_admins[i]] = true;
        }
    }

    //ADMINS Section

    //check participant in admins list
    function isAdmin(address _address) constant returns (bool) {
        return admins[_address];
    }

    function addAdmin() {
        //TODO by voting
    }

    function removeAdmin() {
        //TODO by voting
    }

    //WHITELIST Section

    function isAddressWhitelisted(address _address) constant returns (bool) {
        return trusted[_address];
    }

    function addAddressToWhitelist(address _address) {
        //trusted participant must be added only by admin
        assert(isAdmin(msg.sender));
        //FIXME here must be voting
        trusted[_address] = true;
    }

    //VOTING Section
    //TODO
    address public voting;

    uint public result;

    function makeVoting() {
        voting = new Voting(adminsArray);
    }

    function callback(uint _result) {
        result = _result;
    }
    //PGP Section

    function addKey(bytes32 key) {
        //check sender is trusty
        assert(isAddressWhitelisted(msg.sender));
        //add key to keys
        keys[msg.sender] = key;
    }

    function getKey(address participant) constant returns (bytes32) {
        return keys[participant];
    }

}


contract Voting {

    /*
    TODO List

    1. Add Events
    2. Add Timeouts
    3. Disable multithreading
    4. Fix public modifiers
    5. Constant and enums to external contract
     */

    enum State {
    VOTE,
    FINISHED,
    CANCELED
    }
    State public state;

    //addresses of voters and their indexes
    address[]public addresses; //fixme remove public
    mapping (address => uint)public addressid; //fixme remove public
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

    enum Result {
    UNKNOWN,
    SUCCESS,
    FAIL
    }

    Result public tally;

    address public owner;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier inState(State s) {
        require(state == s);
        _;
    }

    //constructor
    function Voting(address[] _voters) {
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
        owner = msg.sender;
        state = State.VOTE;
        tally = Result.UNKNOWN;
    }

    function accept() inState(State.VOTE) returns (bool) {
        //check participant can vote and hasn't already voted
        //TODO maybe assert to withdraw gas
        if (eligible[msg.sender] && !votecast[msg.sender]) {
            votecast[msg.sender] = true;
            voters[addressid[msg.sender]].vote = uint(Result.SUCCESS);
            totalvoted += 1;
            return true;
        }
        return false;
    }

    function reject() inState(State.VOTE) returns (bool) {
        //check participant can vote and hasn't already voted
        //TODO maybe assert to withdraw gas
        if (eligible[msg.sender] && !votecast[msg.sender]) {
            votecast[msg.sender] = true;
            voters[addressid[msg.sender]].vote = uint(Result.FAIL);
            totalvoted += 1;
            return true;
        }
        return false;
    }

    function countResult() inState(State.VOTE) returns (uint) {
        //all must vote
        assert(totaleligible == totalvoted);
        uint pro = 0;
        uint contra = 0;
        //TODO add timeout rule
        for (uint i = 0; i < totaleligible; i++) {
            //  Check all votes have been cast
            assert(votecast[voters[i].addr]);
            //sum vote
            if (voters[i].vote == uint(Result.SUCCESS)) {
                pro += 1;
            }
            if (voters[i].vote == uint(Result.FAIL)) {
                contra += 1;
            }
        }
        state = State.FINISHED;
        //check result
        if (pro > contra) {
            tally = Result.SUCCESS;
        }
        else {
            tally = Result.FAIL;
        }
        //FIXME make normal callback
        WhiteList(owner).callback(uint(tally));
        return uint(tally);
    }

    function cancelVoting() onlyOwner {
        state = State.CANCELED;
    }

}