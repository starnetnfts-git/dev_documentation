<template>
  <v-layout
    align-center
    justify-center
    class="ma-5 pa-5 metamask-info-template"
  >
    <v-flex xs12 sm8 md6 class="metamask-info-template__row">
      <v-card elevation="0" class="mb-5 metamask-info-template__card">
        <v-card-title>Arbitrum Developer Quickstart</v-card-title>
        <v-card-text
          >You can find a hardhat sample project
          <a
            target="_blank"
            href="https://github.com/starnetnfts-official/dev_documentation/tree/main/mock_contracts"
            >here</a
          >
        </v-card-text>
        <v-card-title>Hardhat Configuration</v-card-title>
        <v-card-text>
          <pre>
starnet: {
	url: "https://rpc1.starnetnft.com",
	gas: 4500000,
	gasPrice: 1000000000, //1 gwei
	timeout: 99000,
	accounts: [process.env.PRIVATE_KEY],
},
          </pre>
          <p>
            if needed, a websocket connection is available at
            <strong>wss://rpc1.starnetnft.com/ws</strong>
          </p>
        </v-card-text>
      </v-card>

      <v-card elevation="0" class="metamask-info-template__card">
        <v-card-title>Verifying Smart Contracts</v-card-title>
        <v-card-text>
          <br />
          <br />
          The easiest way is to flatten your smart contract before publishing.
          <br />
          <pre>
npx hardhat flatten contracts/YOUR_CONTRACT.sol > contracts/YOUR_CONTRACT_Flatten.sol</pre
          >
          <br />
          Then just verify it normaly in starnetscan
        </v-card-text>

        <v-card-text>
          <img
            src="/images/devtip1.png"
            style="max-width: 700px"
            alt="devtip1"
          />
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
