pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract StakingContract is Ownable, ReentrancyGuard {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    IERC20 public stakingToken; 

    struct Stake {
        uint256 amount; 
        uint256 since; 
    }

    mapping(address => Stake) public stakes;

    event Staked(address indexed user, uint256 amount, uint256 timestamp);
    event Withdrawn(address indexed user, uint256 amount);

    constructor(IERC20 _stakingToken) {
        stakingToken = _stakingToken;
    }

    function stakeTokens(uint256 _amount) external nonReentrant {
        require(_amount > 0, "Cannot stake 0 tokens");

        stakes[msg.sender].amount = stakes[msg.sender].amount.add(_amount);
        stakes[msg.sender].since = block.timestamp;

        stakingToken.safeTransferFrom(msg.sender, address(this), _amount);

        emit Staked(msg.sender, _amount, block.timestamp);
    }

    function viewStaked(address _staker) external view returns (uint256) {
        return stakes[_staker].amount;
    }

    function withdrawTokens(uint256 _amount) external nonReentrant {
        require(_amount > 0, "Cannot withdraw 0 tokens");
        require(stakes[msg.sender].amount >= _amount, "Withdrawal request exceeds staked amount");

        stakes[msg.sender].amount = stakes[msg.sender].amount.sub(_amount);

        stakingToken.safeTransfer(msg.sender, _amount);

        emit Withdrawn(msg.sender, _amount);
    }
}