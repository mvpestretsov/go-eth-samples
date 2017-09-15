pragma solidity ^ 0.4.16;


contract Administration {

    /***
     *
    TODO LIST:
    1. Add Events for Adding\Removing admins and trustees

     * **/

    Array public trustees;

    Array public admins;

    function Administration(address[] _admins) {
        //initialize admins and add their to trustees
        admins = new Array(_admins);
        trustees = new Array(_admins);
    }

    /***********
    ADMIN section
     *************/

    function isAdmin(address _addr) constant returns (bool) {
        return admins.isExistElement(_addr);
    }

    function adminSize() constant returns (uint) {
        return admins.ArrayLength();
    }

    modifier onlyAdmin() {
        require(isAdmin(msg.sender));
        _;
    }

    //FIXME trustees
    function addAdmin(address _admin) onlyAdmin {
        //don't be admin
        require(!isAdmin(_admin));
        //must be already trusted
        require(isTrusted(_admin));
        //TODO by voting
        admins.addElement(_admin);
    }

    //FIXME check admins>2
    function removeAdmin(address _admin) onlyAdmin {
        //must be already admin
        require(isAdmin(_admin));
        if (_admin == msg.sender) {
            //admin may remove self without voting
            admins.removeElement(_admin);
        }
        else {
            //FIXME by voting
            admins.removeElement(_admin);
        }
    }

    /********
    TRUSTEES Section
     ********/
    function isTrusted(address _addr) constant returns (bool) {
        return trustees.isExistElement(_addr);
    }

    modifier onlyTrustee() {
        require(isTrusted(msg.sender));
        _;
    }

    function addTrustee(address _addr) onlyAdmin {
        //can't be trusted already
        require(!isTrusted(_addr));
        //FIXME by voting
        trustees.addElement(_addr);
    }

    function removeTrusted(address _addr) onlyAdmin {
        // can't be admin, should be removed from admins before
        require(!isAdmin(_addr));
        if (msg.sender == _addr) {
            //trustee may remove self without voting
            trustees.removeElement(_addr);
            //TODO remove PGP
        }
        else {
            //FIXME by voting
            trustees.removeElement(_addr);
            //TODO remove PGP
        }
    }

    //PGP - FIXME define length of key
    mapping (address => bytes32)keys;

    //PGP Section
    function addKey(bytes32 key) onlyTrustee {
        //add key to keys
        keys[msg.sender] = key;
    }

    function getKey(address trustee) constant returns (bytes32) {
        return keys[trustee];
    }

}


contract Array {

    address[]public array;

    mapping (address => uint)public arrayIndex;

    address public owner;

    modifier onlyOwner() {
        require(owner == msg.sender);
        _;
    }

    function Array(address[] _array) {
        assert(_array.length > 0);
        for (uint i = 0; i < _array.length; i++) {
            if (!isExistElement(_array[i])) {
                add(_array[i]);
            }
        }
        owner = msg.sender;
    }

    function addElement(address _addr) onlyOwner {
        if (!isExistElement(_addr)) {
            array.push(_addr);
            arrayIndex[_addr] = array.length;

        }
    }

    function removeElement(address _addr) onlyOwner {
        if (isExistElement(_addr)) {
            uint index = arrayIndex[_addr];
            array = remove(array, index);
            delete arrayIndex[_addr];

        }
    }

    function add(address _addr) internal {
        array.push(_addr);
        arrayIndex[_addr] = array.length;
    }

    function remove(address[] _array, uint index) internal returns (address[]) {
        if (index >= _array.length)
        return;
        for (uint i = index; i < _array.length - 1; i++) {
            _array[i] = _array[i + 1];
        }
        delete _array[_array.length - 1];
        array.length--;
        return _array;
    }

    function isExistElement(address _element) constant returns (bool) {
        return arrayIndex[_element] > 0;
    }

    function ArrayLength() constant returns (uint) {
        return array.length;
    }
}
