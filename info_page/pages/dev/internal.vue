<template>
  <v-layout
    v-if="isVisible"
    align-center
    justify-center
    class="ma-5 pa-5 metamask-info-template"
  >
    <v-flex xs12 sm8 md6 class="metamask-info-template__row">
      <v-card elevation="0" class="mb-5 metamask-info-template__card">
        <v-card-title class="display-1">Useful Accounts</v-card-title>
        <v-card-text>
          ETH Holder: 0x63E5Cf7d08595e7A92846c6d4B75892EC012325a
        </v-card-text>
        <v-card-text>
          HOT Wallet: 0x76695e6af103C09B7bD4B518c123b2266f3931Bd
        </v-card-text>
        <v-card-text
          >Bala's Metamask:
          0xeCeC39831aa330C9c65655d260AEEd392a20692e</v-card-text
        >
        <v-card-text
          >Faucet Testnet STAR:
          0x22C7883CB16B48520CE275FE7049DFb831bB4953</v-card-text
        >
      </v-card>

      <v-card elevation="0" class="mb-5 metamask-info-template__card">
        <v-card-title class="display-1">STAR NET TESTNET</v-card-title>
        <v-card-text
          >Multisend: 0x8523E9194788C7F4a4e00b7945FA0c29FA9D7c3D</v-card-text
        >
        <v-card-text
          >Receiver: 0xf7144d94BAc0F1593b8166Fb71C448469eF6e261</v-card-text
        >
      </v-card>

      <v-card elevation="0" class="mt-5 metamask-info-template__card">
        <v-card-title class="display-1">ARBITRUM TESTNET</v-card-title>
        <v-card-text
          >STAR:
          https://testnet.arbiscan.io/address/0x007d3561191d984bB6F8cAA95B9Eb73e6bc64792#code</v-card-text
        >
        <v-card-text
          >Bridge Proxy: 0x9AfD35596eC89682235116A9Ec45ca2f9AA6bdfd</v-card-text
        >
        <v-card-text
          >Multisend: 0xCbA082F3265a18245AE0E6952C5DA646Add96fd9</v-card-text
        >
      </v-card>

      <v-card elevation="0" class="mt-5 metamask-info-template__card">
        <v-card-text>
          <pre>
DEBUG=true
SERVER_HOST="localhost"
SERVER_PORT=5000
# STARNET CONFIGURATIONS
RPC_STARNET="wss://rpc-testnet.starnetnft.com/ws"
RPC_STARNET_HTTP="https://rpc-testnet.starnetnft.com"
CHAIN_ID_STARNET=4445
CHAIN_ID_HEX_STARNET = "0x115d"
BRIDGE_STARNET=0x9Eb759db64A3C5E8F2bE42f17dB1E566dFfcc30a
MULTISEND_STARNET=0x8523E9194788C7F4a4e00b7945FA0c29FA9D7c3D
RECEIVER_STARNET=0xf7144d94BAc0F1593b8166Fb71C448469eF6e261
EXPLORER_URI_STARNET="https://testnet.starnetscan.com"
# PARTNER CONFIGURATIONS
PARTNER_NAME=arbitrum_testnet
RPC_PARTNER="wss://arb-rinkeby.g.alchemy.com/v2/2acnyf2B0tm8uuaPumwj4cX85qT_pgJW"
RPC_PARTNER_HTTP="https://rinkeby.arbitrum.io/rpc"
CHAIN_ID_PARTNER=421611
CHAIN_ID_HEX_PARTNER="0x66eeb"
STAR_CONTRACT_PARTNER=0x007d3561191d984bB6F8cAA95B9Eb73e6bc64792
BRIDGE_PARTNER=0x9AfD35596eC89682235116A9Ec45ca2f9AA6bdfd
MULTISEND_PARTNER=0xCbA082F3265a18245AE0E6952C5DA646Add96fd9
EXPLORER_URI_PARTNER="https://testnet.arbiscan.io"
# OTHER CONFIGS
API_KEY_METRICS=XXXXXXXXXXXXXXXXXXXXXXXXX
HOT_WALLET_PRIVATE_KEY=XXXXXXXXXXXXXXXXXXXXXXXXX
</pre
          >
        </v-card-text>
      </v-card>
      <v-card elevation="0" class="mt-5 metamask-info-template__card">
        <v-card-text>
          <pre>
# STARNET CONFIGURATIONS
export const RPC_STARNET_HTTP="https://rpc-testnet.starnetnft.com"
export const CHAIN_ID_STARNET="4445"
export const CHAIN_ID_HEX_STARNET = "0x115d"
export const BRIDGE_STARNET="0x9Eb759db64A3C5E8F2bE42f17dB1E566dFfcc30a"
export const MULTISEND_STARNET="0x8523E9194788C7F4a4e00b7945FA0c29FA9D7c3D"
export const RECEIVER_STARNET="0xf7144d94BAc0F1593b8166Fb71C448469eF6e261"
export const EXPLORER_URI_STARNET="https://testnet.starnetscan.com"
# PARTNER CONFIGURATIONS
export const PARTNER_NAME="arbitrum_testnet"
export const RPC_PARTNER_HTTP="https://rinkeby.arbitrum.io/rpc"
export const CHAIN_ID_PARTNER="421611"
export const CHAIN_ID_HEX_PARTNER="0x66eeb"
export const STAR_CONTRACT_PARTNER="0x007d3561191d984bB6F8cAA95B9Eb73e6bc64792"
export const BRIDGE_PARTNER="0x9AfD35596eC89682235116A9Ec45ca2f9AA6bdfd"
export const MULTISEND_PARTNER="0xCbA082F3265a18245AE0E6952C5DA646Add96fd9"
export const EXPLORER_URI_PARTNER="https://testnet.arbiscan.io"
</pre
          >
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>
<script>
export default {
  auth: false,
  data: () => ({
    isVisible: false,
  }),
  async mounted() {
    const {
      data: { ip },
    } = await this.$axios.get('https://www.cloudflare.com/cdn-cgi/trace', {
      responseType: 'text',
      transformResponse: (data) =>
        Object.fromEntries(
          data
            .trim()
            .split('\n')
            .map((line) => line.split('='))
        ),
    })
    if (ip !== '212.36.9.186') {
      this.$router.push('/error/404')
      return
    }
    this.isVisible = true
  },
}
</script>

<style>
.v-card__title.quick-font-fix-for-mobile {
  font-size: 2em;
}
</style>
