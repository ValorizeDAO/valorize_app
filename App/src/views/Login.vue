<template>
  <div class="bg-purple-50 h-screen">
    <div class="max-w-sm mx-auto pt-24 px-6 md:px-0">
      <h1 class="font-black text-3xl">Login</h1>
      <div class="my-8">
        <form @submit.prevent="sendLogin" class="my-12">
          <label for="name" class="">
            <span class="text-xl font-black">Username</span>
          </label>
          <input type="text" name="name" id="name" class="p-2 mb-12 w-full border-0 border-b-2 border-black bg-purple-50"
                 v-model="name">
          <label for="password" class="mt-24">
            <span class="text-xl font-black">Password</span>
          </label>
          <input type="password" name="password" class="p-2 w-full border-0 border-b-2 border-black bg-purple-50"
                 id="password" v-model="password"> <br>
          <button class="px-4 py-2 mt-4 cursor-pointer border-2 border-black rounded-md font-black w-full mt-20">Submit</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {ref, defineComponent} from "vue";
import {useStore} from 'vuex';
import {useRouter} from 'vue-router'
// import MetamaskLogin from "../components/MetamaskLogin.vue"
export default defineComponent({
  name: "Login",
  // components: {MetamaskLogin},
  props: {},
  setup: () => {
    const store = useStore();
    const router = useRouter();
    const name = ref("");
    const password = ref("");

    async function sendLogin() {
      const formdata = new FormData();
      formdata.append("username", name.value);
      formdata.append("password", password.value);

      const requestOptions = {
        method: "POST",
        body: formdata,
        redirect: "follow",
      } as RequestInit;

      fetch(import.meta.env.VITE_BACKEND_URL + "/login", requestOptions)
          .then(response => response.text())
          .then(result => {
            store.state.authenticated = true;
            router.push('/')
          })
          .catch(error => console.log("error", error));
    }

    return {name, password, sendLogin};
  },
  mounted: () => {
    console.log("mounted");
  }
});
</script>

<style scoped>

</style>
