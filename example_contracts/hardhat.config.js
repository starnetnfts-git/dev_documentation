require("@nomiclabs/hardhat-waffle")
require("dotenv").config()
require("hardhat-gas-reporter")
require("@nomiclabs/hardhat-web3")
require("@nomiclabs/hardhat-etherscan")

module.exports = {
	solidity: {
		version: "0.8.12",
		settings: {
			optimizer: {
				enabled: true,
				runs: 200,
			},
		},
	},
	networks: {
		hardhat: {
			chainId: 127001,
			accounts: {
				mnemonic: "test test test test test test test test test test test junk",
			},
			blockGasLimit: 199022552,
			gas: 1500000,
			gasPrice: 100,
			allowUnlimitedContractSize: false,
			throwOnTransactionFailures: false,
			throwOnCallFailures: true,
		},
		ganache: {
			url: "http://127.0.0.1:7545",
			blockGasLimit: 10000000,
		},
		starnet: {
			url: "https://rpc1.starnetnft.com",
			gas: 4500000,
			gasPrice: 1000000000, //1 gwei
			timeout: 99000,
			accounts: [process.env.PRIVATE_KEY],
		},
		starnet_testnet: {
			url: "https://rpc-testnet.starnetnft.com",
			gas: 4500000,
			gasPrice: 1000000000, //1 gwei
			timeout: 99000,
			accounts: [process.env.PRIVATE_KEY],
		},
	},
	gasReporter: {
		enabled: !!process.env.REPORT_GAS === true,
		currency: "USD",
		gasPrice: 100,
		showTimeSpent: true,
	},
	mocha: {
		timeout: 25000,
	},
	etherscan: {
		apiKey: "abc",
	},
}
