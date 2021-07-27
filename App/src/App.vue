<template>
  <main>
    <nav class="p-4 bg-white border-b-2 border-black">
      <router-link to="/">
        <img src="./assets/logo_large.png" alt="Valorize" class="h-8" />
      </router-link>
    </nav>
    <router-view />
  </main>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
});
</script>

<style>
</style>
