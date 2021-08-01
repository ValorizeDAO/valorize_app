<template>
  <div class="bg-purple-50 mt-20">
    <div class="max-w-sm mx-auto pt-24 px-6 md:px-0">
      <h1 class="font-black text-3xl">Login</h1>
      <div class="my-8">
        <form @submit.prevent="sendLogin" class="my-12">
          <label for="name" class="">
            <span class="text-xl font-black">Username</span>
          </label>
          <input
            type="text"
            name="name"
            id="name"
            class="
              p-2
              mb-12
              w-full
              border-0 border-b-2 border-black
              bg-purple-50
            "
            v-model="name"
          />
          <label for="password" class="mt-24">
            <span class="text-xl font-black">Password</span>
          </label>
          <input
            type="password"
            name="password"
            class="p-2 w-full border-0 border-b-2 border-black bg-purple-50"
            id="password"
            v-model="password"
          />

          <div class="text-sm h-12">
            <transition name="fade">
              <div class="text-red-700" v-if="authError">
                Username or password incorrect
              </div>
            </transition>
          </div>
          <br />
          <transition name="fade" mode="out-in">
            <div v-if="authenticating">
              <SvgLoader fill="#cecece" class="h-12 mx-auto"></SvgLoader>
            </div>
            <button
              v-else
              class="
                px-4
                py-2
                cursor-pointer
                border-2 border-black
                rounded-md
                font-black
                w-full
              "
            >
              Submit
            </button>
          </transition>
          <div class="py-8">
            Or
            <router-link to="/register" class="font-black underline"
              >Register New Account</router-link
            >
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import SvgLoader from "../components/SvgLoader.vue";

export default defineComponent({
  name: "Login",
  components: { SvgLoader },
  props: {},
  setup: () => {
    const store = useStore();
    const router = useRouter();
    const name = ref("");
    const password = ref("");
    const authError = ref(false);
    const authenticating = ref(false);
    async function sendLogin() {
      authError.value = false;
      authenticating.value = true;
      const formdata = new FormData();
      formdata.append("username", name.value);
      formdata.append("password", password.value);

      const requestOptions = {
        method: "POST",
        body: formdata,
        credentials: "include",
      } as RequestInit;

      fetch(import.meta.env.VITE_BACKEND_URL + "/login", requestOptions)
        .then((response) => {
          if (response.status !== 200) {
            authError.value = true;
          }
          return response.json();
        })
        .then((result) => {
          if (!authError.value) {
            store.state.authenticated = true;
            store.commit("authUser/setUser", result);
            router.push("/");
          }
          authenticating.value = false;
        })
        .catch((error) => console.log("error", error));
    }

    return { name, password, sendLogin, authError, authenticating };
  },
});
</script>

<style scoped>
</style>
