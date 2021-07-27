
<template>
  <h1 class="text-9xl font-sans font-black">Home</h1>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
export default defineComponent({
  name: "Dashboard",
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
