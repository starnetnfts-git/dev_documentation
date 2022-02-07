<template>
  <v-app>
    <v-app-bar
      class="app-nav"
      color="transparent"
      clipped-left
      app
      elevation="0"
    >
      <v-toolbar-title>
        <nuxt-link style="text-decoration: none" to="/">
          <img
            src="https://starnetscan.com/images/k_logo.png"
            style="max-height: 70px"
            alt="starnet nft"
          />
        </nuxt-link>
      </v-toolbar-title>
    </v-app-bar>

    <v-main>
      <nuxt />
    </v-main>

    <v-footer class="transparent text-xs-center justify-center">
      <p class="text-xs-center justify-center" style="color: #333">
        <a href="/other/disclaimer">disclaimer</a>
      </p>
    </v-footer>
  </v-app>
</template>
<script>
export default {
  auth: false,
  data() {
    return {}
  },
  mounted() {},
  methods: {
    goToUrl(url) {
      this.$router.push(url)
    },
    goToExternalUrl(url) {
      window.open(url, '_blank')
    },
  },
}
</script>
<style scoped>
#styled-input {
  height: 48px;
  width: 48px;
}
.styled-input label[for] {
  height: 48px;
}
.v-navigation-drawer > .list:not(.list--dense) .list__tile {
  font-size: 17px;
}
.avatar {
  max-width: 75px;
}
/* .list__tile--active.list__tile.list__tile--link {
} */
a.nuxt-link-exact-active.list__tile--active.list__tile.list__tile--link {
  font-weight: 900 !important;
}
.v-list-item {
  border-left: 10px solid transparent;
}
.v-list-item--active {
  color: #333;
  border-left: 10px solid #ff5122;
}
.glow {
  -webkit-text-stroke: 1px #97adc5;
  text-shadow: 0 0 15px rgb(137 246 143 / 77%), 0 0 10px transparent;
  -webkit-text-fill-color: transparent;
}
h1 {
  letter-spacing: 0.5px;
}

.v-main {
  padding-top: 0 !important;
}

.social-btns {
  display: flex;
}

::v-deep .social-btns a {
  width: 72px;
  height: 72px;
}

::v-deep .social-btns i {
  font-size: 42px !important;
}

.logo-img {
  width: 100%;
  max-height: 820px;
  position: absolute;
  bottom: 0;
  z-index: 0;
}
</style>
