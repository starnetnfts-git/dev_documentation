const hre = require("hardhat")
require("@nomiclabs/hardhat-web3")
const fs = require("fs-extra")

function sleep(ms) {
	return new Promise((resolve) => {
		setTimeout(resolve, ms)
	})
}

async function main() {
	fs.removeSync("cache")
	fs.removeSync("artifacts")
	await hre.run("compile")

	let network = process.env.NETWORK ? process.env.NETWORK : "starnetscan"
	console.log(">-> Network is set to " + network)

	// ethers is avaialble in the global scope
	const [deployer] = await ethers.getSigners()
	const deployerAddress = await deployer.getAddress()
	const account = await web3.utils.toChecksumAddress(deployerAddress)
	const balance = await web3.eth.getBalance(account)

	console.log(
		"Deployer Account " + deployerAddress + " has balance: " + web3.utils.fromWei(balance, "ether"),
		"ETH"
	)

	const TContract = await hre.ethers.getContractFactory("MetaboxNFTToken")
	console.log("Deploying MetaboxNFTToken Contract...")

	// const TContract = await hre.ethers.getContractFactory("ERC20Mock")
	// console.log("Deploying ERC20Mock Contract...")

	// const TContract = await hre.ethers.getContractFactory("ERC721Mock")
	// console.log("Deploying ERC721Mock Contract...")

	const deployed = await TContract.deploy()
	let dep = await deployed.deployed()
	console.log("Contract deployed to:", dep.address)
}

main()
	.then(() => process.exit(0))
	.catch((error) => {
		console.error(error)
		process.exit(1)
	})
