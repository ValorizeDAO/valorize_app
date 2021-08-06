<template>
  <main class="min-h-screen">
    <nav
      class="
        h-20
        bg-white
        border-b-2 border-black
        absolute
        w-screen
        flex
        items-center
        px-16
        justify-between
      "
    >
      <router-link to="/">
        <img src="./assets/logo_large.png" alt="Valorize" class="h-8" />
      </router-link>
      <div v-if="!authenticated" class="pr-6 flex">
        <router-link class="pr-4" to="/login">Login</router-link>
        <router-link to="/register">Register</router-link>
      </div>
      <div v-else class="flex">
        <router-link to="/edit-profile" class="mr-4">Edit Profile</router-link>
        <button @click="logout">Logout</button>
      </div>
    </nav>
    <div class="pt-20 min-h-screen w-screen overflow-hidden bg-purple-50">
      <div
        v-if="checkingAuth"
        class="
          flex
          items-center
          justify-center
          mt-20
          bg-purple-900
          absolute
          top-0
          left-0
          w-screen
          h-screen
          z-50
        "
      >
        <SvgLoader fill="#fff" class="h-12 mx-auto"></SvgLoader>
      </div>
      <router-view />
    </div>
  </main>
</template>

<script lang="ts">
import { mapGetters, mapMutations } from "vuex";
import { defineComponent } from "vue";
import SvgLoader from "./components/SvgLoader.vue";
import auth from "./services/authentication";

export default defineComponent({
  name: "App",
  components: { SvgLoader },
  methods: {
    ...mapMutations({ logoutState : "authUser/logout" }),
    async logout() {
      await auth.logout();
      this.logoutState()
      this.$router.push("/login");
    },
  },
  computed: {
    ...mapGetters({
      checkingAuth: "authUser/checkingAuth",
      authenticated: "authUser/authenticated",
      user: "authUser/user",
    }),
  },
});
</script>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
