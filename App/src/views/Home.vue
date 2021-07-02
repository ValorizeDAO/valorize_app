
<template>
  <h1 class="text-9xl font-sans font-black">Home</h1>
  <h1 class="text-2xl">Home</h1>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
export default defineComponent({
  name: "Login",
  props: {

  },
  setup: () => {
    const count = ref(0);
    return { count };
  },
});
</script>

<style scoped>
</style>
