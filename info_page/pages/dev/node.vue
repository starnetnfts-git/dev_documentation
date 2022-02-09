<template>
  <v-layout
    align-center
    justify-center
    class="ma-5 pa-5 metamask-info-template"
  >
    <v-flex xs12 sm8 md6 class="metamask-info-template__row">
      <v-card elevation="0" class="mb-5 metamask-info-template__card">
        <v-card-title
          >Running full node for Star Net NFT Blockchain</v-card-title
        >
        <v-card-text>Prerequisites: Java JDK (11-16)</v-card-text>
        <v-card-title>Linux / Unix</v-card-title>
        <v-card-text>
          <p>
            Download the latest besu release from
            <a
              target="_blank"
              href="https://github.com/hyperledger/besu/releases"
              >HERE</a
            >
          </p>

          <p>
            Unpack the downloaded files and change into the besu-release
            directory. Display Besu command line help to confirm installation:
            <br />
            <br />
            <pre>bin/besu --help</pre>
          </p>
        </v-card-text>
      </v-card>

      <v-card elevation="0" class="metamask-info-template__card">
        <v-card-title>Configuration</v-card-title>
        <v-card-text>
          <br />
          <br />
          In order to run a node you'll need our genesis.json file
          <br /> <br />
          <pre>
{
  "config" : {
    "chainId" : 4444,
    "muirglacierblock" : 0,
    "ibft2" : {
      "blockperiodseconds" : 3,
      "epochlength" : 10000,
      "requesttimeoutseconds" : 6
    }
  },
  "nonce" : "0x0",
  "timestamp" : "0x61EFEC19",
  "gasLimit" : "0x1312D00",
  "difficulty" : "0x1",
  "mixHash" : "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365",
  "coinbase" : "0x0000000000000000000000000000000000000000",
  "alloc" : {
    "63E5Cf7d08595e7A92846c6d4B75892EC012325a" : {
      "balance" : "100000000000000000000000000"
    }
  },
  "extraData" : "0xf87ea00000000000000000000000000000000000000000000000000000000000000000f8549469f32789bf976f6211aac6efb71d4745bfd7c5cd943eecc7080522ea66391ee71e32dd8fa23b556550945b61eba48ab0d4812a0a76b67d2a229b3671608894c8b4db4094528dadf363d6fd40af726838fff4a4808400000000c0"
}</pre
          >
          <br />
          <br />
          <p>and an enode URI that our team will give you. Please contact us for this information</p>
        </v-card-text>

        <v-card-title>Besu Example Configuration</v-card-title>
        <v-card-text>
            <pre>
/usr/bin/besu --data-path=/PATH_TO_NODE_FILES \
 --genesis-file=/PATH_TO_NODE_FILES/genesis.json --rpc-http-enabled \
 --rpc-http-api=ETH,NET,WEB3 --host-allowlist="*" \
 --rpc-http-host=127.0.0.1 \
 --rpc-http-port=8545 \
 --rpc-http-cors-origins="all" \
 --rpc-ws-api=ETH,NET,WEB3 \
 --rpc-ws-enabled \
 --rpc-ws-host=127.0.0.1 \
 --rpc-ws-max-active-connections=200 \
 --rpc-http-max-active-connections=200 \
 --rpc-ws-port=8546 \
 --static-nodes-file=/PATH_TO_NODE_FILES/static-nodes.json \
 --network-id=4444 \
 --api-gas-price-max=100000000000 \
 --max-peers=20 \
 --identity=YOUR_NODE \
 --min-gas-price=1000000000 \
 --tx-pool-max-size=4096 \
 --tx-pool-price-bump=10 \
 --p2p-port=30303
            </pre>
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>
<script>
export default {
  auth: false,
  data: () => ({
    variable: null,
  }),
  mounted() {},
}
</script>

<style>
.v-card__title.quick-font-fix-for-mobile {
  font-size: 2em;
}
</style>
