# Star Net NFT Blockchain Benchmarks

### TL/DR;

With a firewall, max per IP is 60 TPS
Without firewall, max per IP is 

### Testing Procedures

Basically there are two main parameters you can tweak on:

big discussion for someone that has time: https://github.com/ethereum/go-ethereum/issues/18402

- block time (how fast to produce blocks?)
- block gas limit (how much to include in a block?)

#### Propagation time can't be translated directly into TPS, because

TPS depends mainly on gas limit and blocktime. But it can affect synchronization in the network, degrading the TPS that
can be calculated theoretically from the configuration of the network.

## Important info

#### Sealers Number

In PoA network consensus is achieved by a majority agreement among the sealers nodes. A sealer node is a special client
which is allowed to include blocks on the blockchain. Sealers are set in a whitelist in the blockchain genesis block.
Once the blockchain is running, new sealers could be added by majority voting. To consider a block as valid, it must be
validated by at least 51% of sealers. By increasing the number of sealers the network latency could be also increased.
This can generate synchronization problems during the generation of blocks. It is necessary to study how the amount of
sealers affects the performance of the network.

#### Block Time

Clique consensus algorithm divides time into epochs. At the beginning of each epoch, a sealer is selected using a Round
Robin algorithm as the leader to propose a new block. During the epoch, the leader validates transactions and includes
them in the new block, and once the epoch is finished, it broadcasts the block to the other sealers. If the majority of
the sealers accepted it, the block can (finally) be considered as valid. In case that the leader delays in submitting
the block, some back-up sealers can take its place and propose another block instead. The time between two consecutive
blocks is called block time. Despite the fact that in PoA networks theoretical block time is fixed by configuration, it
can fluctuate due to synchronization and network delays. That is why it is interesting to measure real block times
given other varying blockchain parameters, e.g. number of sealers and Gas Limit (which determines the block size). The
block time configuration parameter can be used to set the maximum network throughput, as evaluated in Gas Limit section.

#### Gas Limit

Ethereum platform prevents transaction spamming and rewards block miners by charging a gas fee on transactions. Each
block contains a maximum amount of gas that can be collected from transactions, defining a maximum block size. That gas
limit could be set as a configuration parameter. In the long term, the block gas limit approaches a target gas limit set
also as a configuration parameter (it can also be changed at runtime if needed). The theoretical maximum transactions
per second (TPS) can be calculated using the following equation:

TPS = (GAS Limit ) / (TxGas \* BlockTime)

where GasLimit is the Block gas limit, TxGas is the gas needed to compute the simplest transaction and Block time is the
blockchain block time.

### So what is the optimal gas limit ?

According to this
research: https://www.researchgate.net/publication/343924396_Performance_Evaluation_of_Private_Ethereum_Networks
Additionally, we can conclude that the maximum block gas limits hould be between 0x8000000 and 0x10000000.

## 1 second time. 3 sealers

---

- wallets stop functioning correctly (metamask). errors like transactions stuck in pending and insufficient funds for
  transaction (even thought they are there)

Conclusion: while sealers appear to produce blocks correctly, rpc works, etc. the wallets malfunction for some reason.
Setting the time of 1 second for block production is not possible.

## 2 seconds time. 3 sealers

- metamask functions correctly 50% of the time. sometimes it shows wrong amount. 1 out of 10 transactions show up as
  failed, even though they went through

Conclusion: no. geth is stable but the wallets no.

## 3 seconds, 3 sealers

"gasLimit": "0x10000000",
1500 transactions per block = 500 TPS

## 3 seconds, 3 sealers

"gasLimit": infinite
1300 transactions per block = 433 TPS

### Optimal Hardware Specs:

RPC Nodes: 2 CPU, 2GB RAM, SSD HDD
Sealer Nodes: 2 CPU, 2GB RAM, SSD HDD, Needs Better Network Connection than a RPC

### Optimal Node Configurations:

- 2 sealers or 3, 5, 7, 9, 11... (prime numbers)
- 1 or more bootnodes
- as many RPC nodes as possible

defaults:

block gas limit of 4,700,000 running on 5-second block intervals and we can support 44 transactions per second at 21,000 gas per transaction

The Rinkeby network has a current gas limit of around 7,000,000 per block and roughly 15-second block intervals, which equates to 22 transactions per second at the same 21,000 gas price per transaction

### Final recommendation

10,000,000 block limit for 5 seconds is a sensible choice... giving 95 TPS

what happens if there are more than 95\*5 transactions at some point:

- RPC nodes will buffer them

what happens if the RPC nodes are full and can't buffer more ?

- wallet should retry them

what happens if both RPC nodes are full continuously

- attacker should run out of ETH
- the firewall before the official RPCs will block the IPs that flood the network
- the attacker should run out of ETH fast. address will be blacklisted.
