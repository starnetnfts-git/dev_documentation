<template>
  <v-layout
    align-center
    justify-center
    class="ma-5 pa-5 metamask-info-template"
  >
    <v-flex xs12 sm8 md6 class="metamask-info-template__row">
      <v-card elevation="0" class="metamask-info-template__card">
        <v-card-title class="quick-font-fix-for-mobile"
          >Disclaimer</v-card-title
        >
        <v-card-text>
          <br />
          <br />
          Our mission is to make NFTs as easy to use for as many people as
          possible.
        </v-card-text>

        <v-card-text>
          <br />WE WILL NOT BE LIABLE FOR ANY LIABILITY OR DAMAGES WHATSOEVER
          ASSOCIATED WITH YOUR USE, INABILITY TO USE, OR YOUR INTERACTION WITH
          OTHER USERS OF, THIS DAPP, OTHER DAPPS PLATFORM, INCLUDING ANY DIRECT,
          INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY OR CONSEQUENTIAL DAMAGES,
          LOSS OF PROFITS, OR DATA. <br />WE DO NOT OFFER ANY REFUNDS AND ALL
          TRANSACTIONS ARE CONSIDERED FINAL
        </v-card-text>

        <v-card-text class="blue--text">
          <v-btn class="success white--text" href="/">BACK TO HOME PAGE</v-btn>
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
