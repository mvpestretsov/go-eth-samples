pragma solidity ^0.4.16;


contract WhiteList {

    /*
    Questions:

    1. Q:rules to voting - 50%+1  or smth else ? Maybe admin votes?
    A: Admin must be voted by all amins
    2. Q:if voting is expired and there is no decision - discard voting?
    A:Yes. New admin must be voted all admins.
    3. Q:trusted=>admin ?
    A:No. Trusted is only for receiving info,admins for voting
    */

    // biznet originator
    address beginer;
    //proved participants
    mapping (address => bool) trusted;
    //super participants
    mapping (address => bool) admins;
    //PGP - FIXME define length of key to change type
    mapping (address => bytes32) keys;

    //constructor - set initial admins + sender to trusted
    function WhiteList(address[] _address){
        //save initiator
        beginer = msg.sender;
        //add initiator to admins
        admins[msg.sender] = true;
        //add initiator to trusted
        trusted[msg.sender] = true;
        for (uint i = 0; i < _admins.length; i++) {
            //add admin to admins
            admins[_admins[i]] = true;
            //add admin to trusted
            trusted[_admins[i]] = true;
        }
    }


    //ADMINS Section

    //check participant in admins list
    function isAdmin(address _address) constant returns (bool) {
        return admins[_address];
    }

    function addAdmin(){
        //TODO by voting
    }

    function removeAdmin(){
        //TODO by voting
    }

    //WHITELIST Section

    function isAddressTrusted(address _address) constant returns (bool) {
        return trusted[_address];
    }

    function addAddressToTrusted(address _address){
        //trusted participant must be added only by admin
        assert(isAdmin(msg.sender));
        //TODO here must be voting
        trusted[_address] = true;
        // -> dope
    }

    //VOTING Section
    //TODO

    //PGP Section
    //WIP
    function addKey(bytes32 key){
        //check sender is trusty
        assert(isAddressTrusted(msg.sender));
        //
    }
}