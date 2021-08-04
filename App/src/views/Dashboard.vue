
<template>
  <div class="px-16 py-4">
    <h1 class="text-5xl font-sans font-black">Welcome, {{ user.username }}!</h1>
    <h2 class="text-3xl font-sans font-black my-4">Get started by setting a profile and deploying your own token</h2>
    <router-link to="edit-profile" class="btn mx-auto my-3 w-40 bg-paper-light">Setup Profile</router-link>
  </div>

</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import {useStore} from "vuex";
export default defineComponent({
  name: "Dashboard",
  props: {

  },
  setup: () => {
    const store = useStore();
    const user = store.getters["authUser/user"]
    return { store, user };
  },
});
</script>

<style scoped>
</style>
