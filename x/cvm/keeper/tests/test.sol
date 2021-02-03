contract Token {
    uint256 totalSupply;
    mapping(address => uint256) balances;
    uint256 realint;

    event Transfer(address indexed _from, address indexed _to, uint256 _value);

    constructor() public {
        totalSupply = 100000000;
        balances[msg.sender] = totalSupply;
        realint = 123;
    }

    function balanceOf(address account) view public returns (uint256) {
        return balances[account];
    }

    function transfer(address to, uint256 amount) public returns (bool) {
        balances[msg.sender] = balances[msg.sender] - amount;
        balances[to] = balances[to] + amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    function wow() public returns (uint256) {
        return realint;
    }
}