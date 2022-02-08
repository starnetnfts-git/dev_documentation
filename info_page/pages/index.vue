<template>
  <v-container>
    <v-layout
      align-center
      justify-center
      class="ma-5 pa-5 metamask-info-template"
    >
      <v-flex xs12 sm8 md6 class="metamask-info-template__row">
        <div>
          <p class="display-1">
            <span class="underline">StarnetNFT</span> blockchain
          </p>

          <p class="mt-5 title">Connect to the blockchain</p>
          <p class="body">
            Interacting with the StarnetNFT blockchain requires a compatible
            crypto wallet such as
            <a target="_blank" href="https://metamask.io/">MetaMask</a>.
            However, MetaMask doesn’t have StarnetNFT added as a network by
            default. Setting up your browser wallet to connect to StarnetNFT is
            simple and can be done in just a few minutes.
          </p>
          <p class="mt-5 body">
            <a target="_blank" href="https://metamask.io/">MetaMask</a> is
            available to download and install on Chrome, iOS, or Android through
            the MetaMask website. Always check if you are using the official
            website to make sure you’re downloading the real MetaMask extension.
          </p>

          <p class="mt-5 body">
            Try adding automatically by clicking the following buttons and
            accepting the prompt in metamask
            <br />
            <v-btn class="mt-3 mb-3 primary" @click="connectWallet()">{{
              connectWalletText
            }}</v-btn>
            <v-btn
              class="mt-3 mb-3 primary"
              @click="addStarentNFTChainToMetamask()"
              >2. AND ADD STARNETNFT TO METAMASK</v-btn
            >
          </p>

          <p class="mt-5 body">
            You can also connect manually. Open MetaMask and click the network
            dropdown menu. Click Add Network button and add the following
            details
          </p>

          <v-card class="ma-5 pa-5">
            <v-card-title>
              <p>
                <span style="font-weight: bold">Network Name:</span> Star Net
              </p>
              <p>
                <span style="font-weight: bold">New RPC URL:</span>
                https://rpc1.starnetnft.com
              </p>
              <p><span style="font-weight: bold">Chain ID:</span> 4444</p>
              <p>
                <span style="font-weight: bold">Block Explorer URL:</span>
                https://starnetscan.com
              </p>
            </v-card-title>
          </v-card>

          <p>
            <img
              src="/images/starnetnft_explanations/1.png"
              alt="starnetnft explanation 1"
            />
            <img
              src="/images/starnetnft_explanations/2.png"
              alt="starnetnft explanation 2"
            />
            <img
              src="/images/starnetnft_explanations/3.png"
              alt="starnetnft explanation 3"
            />
          </p>

          <p class="ma-5 body-1">
            With StarnetNFT set up in MetaMask, you’re free to start making
            transactions, collecting NFTs interacting with DeFi DApps, and
            managing your crypto.
          </p>
        </div>

        <v-dialog v-model="dialogSuccess" class="ma-5 pa-5" max-width="600px">
          <v-card class="success">
            <v-card-title>
              <span>{{ successText }}</span>
            </v-card-title>
            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn
                color="white"
                text
                @click="
                  dialogSuccess = false
                  successText = ''
                "
              >
                EXIT
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>

        <v-dialog v-model="dialogError" class="ma-5 pa-5" max-width="600px">
          <v-card class="warning">
            <v-card-title>
              <span>{{ errorText }}</span>
            </v-card-title>
            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn
                color="white"
                text
                @click="
                  dialogError = false
                  errorText = ''
                "
              >
                EXIT
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { ethers } from 'ethers'
import { RPC_PROVIDER, NETWORK_ID } from '../constants'
const EthersUtils = require('ethers').utils

export default {
  auth: false,
  data() {
    return {
      isLoading: false,
      dialogConfirmation: false,
      tokenIDs: null,
      contract: null,
      loadingText: 'loading...',
      account: null,
      dialogSuccess: false,
      txHash: null,
      isOwned: false,
      ethers: null,
      signer: null,
      provider: null,
      errorText: '',
      dialogError: false,
      walletAddress: null,
      connectWalletText: '1. CONNECT WALLET',
      successText: '',
    }
  },
  async mounted() {},
  methods: {
    async connectWallet() {
      this.ethers = new ethers.providers.Web3Provider(window.ethereum, 'any')
      await this.ethers.send('eth_requestAccounts', [])

      this.signer = this.ethers.getSigner()
      this.account = await this.signer.getAddress()
      this.connectWalletText = 'CONNECTED'
    },
    async addStarentNFTChainToMetamask() {
      console.log('adding starnetnft to metamask...')
      if (window.ethereum) {
        await window.ethereum.request({
          method: 'wallet_addEthereumChain',
          params: [
            {
              chainId: '0x115C',
              chainName: 'Star Net',
              nativeCurrency: {
                name: 'STAR',
                symbol: 'STAR',
                decimals: 18,
              },
              rpcUrls: ['https://rpc1.starnetnft.com'],
              blockExplorerUrls: ['https://starnetscan.com'],
              iconUrls: [],
            },
          ],
        })
      } else {
        this.$router.push('/other/install_metamask')
      }
    },
  },
}
</script>

<style lang="scss" scoped>
.main-block {
  max-width: 960px;
  margin: auto;
  position: relative;
}
.main-block .m-right {
  position: absolute;
  top: 180px;
  right: -220px;
}
.main-block p {
  line-height: 1.6;
  font-size: 20px;
  margin: auto;
}
.main-block img {
  margin: 20px auto;
  max-width: 520px;
}
.container {
  max-width: 1500px;
}
.black-text {
  color: black i !important;
}
.theme--dark.v-input input,
.theme--dark.v-input textarea {
  color: #ea700c;
}

.centered-input input {
  text-align: center;
}
.container div {
  z-index: 1;
}
.banner {
  display: flex;
}
.banner h2 {
  font-size: 48px;
  font-family: Poppins-ExtraBold;
  padding: 18px 45px;
  background: #346dac;
  color: #000000;
  text-align: center;
  margin-left: auto;
  width: 100%;
}
.timer {
  display: flex;
  padding: 35px 45px;
  margin-left: 37%;
  line-height: 1.15;
}
.timer div {
  font-family: Poppins-ExtraBold !important;
  font-size: 50px !important;
  text-align: right;
}
.timer span {
  font-size: 51px !important;
  color: #346dac !important;
  font-family: Poppins-ExtraBold !important;
  padding: 0 6px;
  line-height: 1;
}
.timer p {
  font-size: 10px !important;
  font-family: Poppins-ExtraBold !important;
  text-align: center;
  margin: auto;
}
.main-right-block {
  padding-left: 30px;
}
.main-right-block > div {
  display: flex;
  margin-bottom: 40px;
  line-height: 1.2;
}
.main-right-block h3 {
  font-size: 34px;
  font-family: Poppins-Bold;
  width: 185px;
  min-width: 185px;
  text-align: right;
  word-break: break-word;
}
.main-right-block ul {
  margin-left: 25px;
  padding-left: 25px;
  border-left: 4px solid #df0;
}
.main-right-block ul li {
  margin-bottom: 25px;
}
.main-right-block form {
  margin-left: 25px;
  width: 100%;
}
.sel-btn {
  display: flex;
  justify-content: space-between;
  border: 4px dotted #cc0000;
  padding: 10px 16px;
  align-items: center;
  max-width: 242px;
  width: 100%;
  margin-right: 20px;
}
.sel-btn p {
  font-size: 17px;
  margin-bottom: 0;
}
::v-deep .v-icon {
  color: #346dac !important;
}
::v-deep .quantity-input {
  max-width: 64px;
  border-left: 1px solid;
  border-radius: 0;
}
::v-deep .v-input__control {
  min-height: 24px !important;
  height: 24px;
}
::v-deep .quantity-input .v-input__slot {
  padding-right: 0px !important;
  margin: auto;
  background: transparent !important;
  box-shadow: none !important;
}
::v-deep .quantity-input .v-select__slot .v-select__selection,
::v-deep .quantity-input .qty-amount,
::v-deep .quantity-input select {
  color: #fff !important;
  text-align: center;
  font-size: 30px !important;
  right: 0 !important;
}
::v-deep .quantity-input .v-text-field__details {
  display: none !important;
}
::v-deep .quantity-input .v-select__slot input {
  display: none;
}

.kingMessage {
  font-size: 1.2em;
  padding: 15px;
  border: 1px dashed lightblue;
}

.currentKing {
  padding: 10px;
  border: 1px dashed lightgreen;
}

.treasury {
  font-size: 1.2em;
  padding: 15px;
  border: 1px dashed lightcoral;
}

::v-deep .mint-btn-grey {
  font-size: 20px;
  font-weight: bold;
  height: 64px !important;
  background: transparent !important;
  border: 6px solid #888;
  text-transform: capitalize !important;
  border-radius: 0 !important;
  max-width: 342px;
  width: 100%;
  margin-top: 30px;
}
::v-deep .mint-btn-grey .v-btn__content {
  color: #fff !important;
}
::v-deep .mint-btn-grey {
  will-change: transform;
  transition: transform 250ms;
}
::v-deep .mint-btn-grey:hover {
  transform: translateY(-3px);
}
::v-deep .mint-btn-grey:active {
  transform: translateY(0px);
}

::v-deep .mint-btn-grey2 {
  font-size: 14px;
  height: 64px !important;
  background: transparent !important;
  border: 4px dotted #444;
  text-transform: capitalize !important;
  border-radius: 0 !important;
  max-width: 342px;
  width: 100%;
  margin-top: 30px;
}
::v-deep .mint-btn-grey2 .v-btn__content {
  color: #fff !important;
}
::v-deep .mint-btn-grey2 {
  will-change: transform;
  transition: transform 250ms;
}
::v-deep .mint-btn-grey2:hover {
  transform: translateY(-3px);
}
::v-deep .mint-btn-grey2:active {
  transform: translateY(0px);
}

::v-deep .mint-btn {
  font-size: 20px;
  font-weight: bold;
  height: 64px !important;
  border: 6px solid #df0;
  text-transform: capitalize !important;
  border-radius: 0 !important;
  max-width: 342px;
  width: 100%;
  margin-top: 30px;
}
::v-deep .mint-btn .v-btn__content {
  color: #333 !important;
}
::v-deep .mint-btn {
  will-change: transform;
  transition: transform 250ms;
}
::v-deep .mint-btn:hover {
  transform: translateY(-3px);
}
::v-deep .mint-btn:active {
  transform: translateY(0px);
}
@media (max-width: 1200px) {
  .timer {
    margin-left: 15%;
  }
}
@media (max-width: 767px) {
  .banner h2 {
    font-size: 35px;
    margin: auto;
  }
  .timer {
    text-align: center;
    margin-left: 10%;
  }
  .col {
    padding: 0 !important;
  }
  .main-right-block {
    display: flex;
    flex-direction: column-reverse;
  }
  .main-right-block > div {
    flex-direction: column;
    padding-left: 40px;
    padding-right: 40px;
  }
  .main-right-block h3 {
    width: 100%;
    text-align: left;
  }
  .main-right-block ul {
    margin-left: 0;
    padding-left: 0px;
    padding-top: 25px;
    margin-top: 25px;
    border-left: none;
    border-top: 4px solid #346dac;
  }
  .main-right-block form {
    margin-left: 0;
  }
  .presale-minting {
    margin-bottom: 300px !important;
  }
}
.container {
  max-width: 1500px;
}
.black-text {
  color: black i !important;
}
.theme--dark.v-input input,
.theme--dark.v-input textarea {
  color: #ea700c;
}
.underline {
  text-decoration-line: underline;
  text-decoration-style: 'solid';
  text-decoration-color: rgb(0, 197, 232);
  text-decoration-thickness: 5px;
}

.centered-input input {
  text-align: center;
}
.main-block {
  max-width: 560px;
  margin: auto;
  position: relative;
}
.main-block .m-right {
  position: absolute;
  top: 180px;
  right: -220px;
}
.main-block p {
  line-height: 1.6;
  text-align: center;
  font-size: 20px;
  margin: auto;
}
.main-block img {
  margin: 20px auto;
  max-width: 520px;
}
::v-deep .quantity-input {
  width: 180px;
}
::v-deep .quantity-input .v-input__slot {
  height: 48px;
  margin: auto;
}
::v-deep .quantity-input .v-select__slot .v-select__selection,
::v-deep .quantity-input .qty-amount,
::v-deep .quantity-input select {
  color: #fff !important;
  text-align: center;
  font-size: 16px !important;
  right: 0 !important;
}
::v-deep .quantity-input .v-text-field__details {
  display: none !important;
}
::v-deep .quantity-input .v-select__slot input {
  display: none;
}
::v-deep .mint-btn {
  height: 48px !important;
}
::v-deep .mint-btn-grey {
  height: 48px !important;
}
::v-deep .mint-btn .v-btn__content {
  color: #333 !important;
}
::v-deep .mint-btn-grey .v-btn__content {
  color: #333 !important;
}
::v-deep .mint-btn {
  will-change: transform;
  transition: transform 250ms;
}
::v-deep .mint-btn:hover {
  transform: translateY(-3px);
}
::v-deep .mint-btn:active {
  transform: translateY(0px);
}
.search-form__row p {
  line-height: 2;
  text-align: center;
  font-size: 20px;
  margin: auto;
}
@media (max-width: 767px) {
  .main-block img {
    width: 100%;
  }
  ::v-deep .quantity-input {
    width: 100%;
  }
  .main-block .m-right {
    position: relative !important;
    margin-top: 20px;
    top: 0;
    right: 0;
  }
}
</style>
