<template>
  <main class="min-h-screen">
    <nav class="h-20 bg-white border-b-2 border-black absolute w-screen flex items-center pl-6">
      <router-link to="/">
        <img src="./assets/logo_large.png" alt="Valorize" class="h-8" />
      </router-link>
    </nav>
    <div class="pt-20">
      <router-view />
    </div>
  </main>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
});
</script>

<style lang="postcss">
</style>
