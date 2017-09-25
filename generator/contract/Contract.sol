pragma solidity ^0.4.16;


contract Parent {
    uint public t;

    function generateChild(uint i) payable returns (address){
        Child child = new Child(i);
        t = i * 2;
        return child;
    }
}


contract Child {
    uint private str;

    function Child(uint i){
        str = i;
    }

    function getStrorage() constant returns (uint)  {
        return str;
    }
}