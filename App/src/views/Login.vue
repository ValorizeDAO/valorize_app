<template>
<div class="text-center">
  <h1>Login</h1>
  <form @submit.prevent="sendLogin" class="flex flex-col max-w-md items-center mx-auto">
    <input type="text" name="name" id="name" class="border border-gray-700 p-2 mb-2" v-model="name">
    <input type="password" name="password" class="border border-gray-700 p-2" id="password" v-model="password">
    <button class="px-4 py-2 mt-4 cursor-pointer bg-gray-100">Submit</button>
  </form>
  <!-- <div class="mt-8"> -->
  <!-- Or: <br> -->
  <!-- <metamask-login></metamask-login> -->
  <!-- </div> -->
</div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
// import MetamaskLogin from "../components/MetamaskLogin.vue"
export default defineComponent({
  name: "Login",
  // components: {MetamaskLogin},
  props: {

  },
  setup: () => {
    const name = ref("");
    const password = ref("");

    async function sendLogin() {
      const formdata = new FormData();
      console.log(name);
      formdata.append("username", name.value);
      formdata.append("password", password.value);

      const requestOptions = {
        method: "POST",
        body: formdata,
        credentials: "include",
        redirect: "follow",
      } as RequestInit;

      fetch(import.meta.env.VITE_BACKEND_URL+"/login", requestOptions)
        .then(response => response.text())
        .then(result => console.log(result))
        .catch(error => console.log("error", error));
    }
    return { name, password, sendLogin };
  },
  mounted: () => {
    console.log("mounted");
  }
});
</script>

<style scoped>
a {
  color: #42b983;
}

label {
  margin: 0 0.5em;
  font-weight: bold;
}

code {
  background-color: #eee;
  padding: 2px 4px;
  border-radius: 4px;
  color: #304455;
}
</style>
