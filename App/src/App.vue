<template>
  <main class="min-h-screen">
    <nav class="h-20 bg-white border-b-2 border-black absolute w-screen flex items-center pl-6 justify-between">
      <router-link to="/">
        <img src="./assets/logo_large.png" alt="Valorize" class="h-8" />
      </router-link>
      <div v-if="!authenticated" class="pr-6 flex">
        <router-link class="pr-4" to="/login">Login</router-link>
        <router-link to="/register">Register</router-link>
      </div>
      <div v-else class="pr-6 flex">
        <button @click="logout">Logout</button>
      </div>
    </nav>
    <div class="pt-20 min-h-screen w-screen overflow-hidden bg-purple-50">
      <router-view v-if="!checkingAuth" />
      <div v-if="checkingAuth" class="flex items-center justify-center">
        <SvgLoader fill="#cecece" class="h-12 mx-auto"></SvgLoader>
      </div>
    </div>
  </main>
</template>

<script lang="ts">
import { mapGetters } from "vuex"
import { defineComponent } from "vue";
import SvgLoader from "./components/SvgLoader.vue";

export default defineComponent({
  name: "App",
  components: { SvgLoader },
  methods: {
    async logout() {
      this.$store.dispatch('logout')
      this.$router.push('login')
    }
  },
  computed: {
    ...mapGetters(['checkingAuth', 'authenticated', 'user'])
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
