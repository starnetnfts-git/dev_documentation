<template>
  <v-layout align-center justify-center class="metamask-info-template">
    <v-flex xs12 sm8 md6 class="metamask-info-template__row">
      <v-card class="metamask-info-template__card">
        <v-card-title class="headline quick-font-fix-for-mobile"
          >Metamask is Needed</v-card-title
        >
        <v-card-text>
          MetaMask is a web extension, which allows you to manage your Ethereum
          private keys via your web browser.
        </v-card-text>

        <v-card-text>
          Install MetaMask from
          <a target="_blank" href="https://metamask.io/"
            >https://metamask.io/</a
          >
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
  font-size: 13px;
}
</style>
