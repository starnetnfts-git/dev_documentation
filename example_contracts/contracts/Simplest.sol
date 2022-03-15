// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

contract Simplest {
    uint256 public version = 1;
	constructor(){}
    function setVersion(uint256 newVersion) external {
        version = newVersion;
    }
    function yy1(uint256 something) internal {}
}
