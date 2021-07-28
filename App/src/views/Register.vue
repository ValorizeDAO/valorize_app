<template>
  <div class="bg-purple-50 h-screen">
    <div class="max-w-sm mx-auto pt-24 px-6 md:px-0">
      <h1 class="font-black text-3xl">Register</h1>
      <div class="my-8">
        {{ status }}
        <form @submit.prevent="sendLogin" class="my-12">
          <label for="name" class="">
            <span class="text-xl font-black">Email</span>
          </label>
          <input
            type="email"
            name="email"
            id="email"
            class="
              p-2
              mb-12
              w-full
              border-0 border-b-2 border-black
              bg-purple-50
            "
            v-model="email"
          />
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
            class="
              p-2
              w-full
              border-0 border-b-2 border-black
              bg-purple-50
              mb-12
            "
            id="password"
            v-model="password"
          />
          <br />
          <label for="password" class="mt-24">
            <span class="text-xl font-black">Re-Enter Password</span>
          </label>
          <input
            type="password"
            name="password"
            class="p-2 w-full border-0 border-b-2 border-black bg-purple-50"
            id="password"
            v-model="passwordVerify"
          />
          <br />
          <div class="mt-20">
            <transition name="fade" mode="out-in">
              <div v-if="status === 'pending'">
                <SvgLoader fill="#cecece"  class="h-12 mx-auto"></SvgLoader>
              </div>
              <input
                v-else-if="['init', 'error', 'conflict'].includes(status)"
                type="submit"
                class="px-4 py-2 border-2 rounded-md font-black w-full"
                :class="
                  ready
                    ? 'text-black border-black bg-purple-50 cursor-pointer'
                    : 'cursor-not-allowed text-gray-400'
                "
                :disabled="!ready"
              />
            </transition>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent, computed } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
// import MetamaskLogin from "../components/MetamaskLogin.vue"
import SvgLoader from "../components/SvgLoader.vue"
export default defineComponent({
  name: "Register",
  components: {SvgLoader},
  setup: () => {
    const requestStatuses = ["init", "pending", "conflict", "error", "success"];
    const store = useStore();
    const router = useRouter();
    const name = ref("");
    const email = ref("");
    const password = ref("");
    const passwordVerify = ref("");
    const status = ref(requestStatuses[0]);
    const ready = computed(() => {
      const passwordIsNotEmpty =
        password.value != "" && passwordVerify.value != "";
      return passwordIsNotEmpty && password.value === passwordVerify.value;
    });

    async function sendLogin() {
      status.value = requestStatuses[1];

      const formdata = new FormData();
      formdata.append("username", name.value);
      formdata.append("password", password.value);
      formdata.append("email", email.value);

      const requestOptions = {
        method: "POST",
        body: formdata,
        redirect: "follow",
      } as RequestInit;

      fetch(import.meta.env.VITE_BACKEND_URL + "/register", requestOptions)
        .then((response) => {
          console.log(response);
          if (response.status === 409) {
            status.value = requestStatuses[2];
          } else if (response.status === 201) {
            status.value = requestStatuses[4];
          } else {
            status.value = requestStatuses[3];
          }
          return response.json();
        })
        .then((result) => {
          if (!result.error) {
            store.state.authenticated = true;
            router.push("/");
          }
        })
        .catch((error) => console.log("error", error));
    }

    return { name, email, password, passwordVerify, status, ready, sendLogin };
  },
});
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.spinner {
  animation: rotate 2s linear infinite;
  z-index: 2;
  position: absolute;
  top: 50%;
  left: 50%;
  margin: -25px 0 0 -25px;
  width: 50px;
  height: 50px;
}
.spinner .path {
  stroke: hsl(210, 70, 75);
  stroke-linecap: round;
  animation: dash 1.5s ease-in-out infinite;
}

@keyframes rotate {
  100% {
    transform: rotate(360deg);
  }
}

@keyframes dash {
  0% {
    stroke-dasharray: 1, 150;
    stroke-dashoffset: 0;
  }
  50% {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -35;
  }
  100% {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -124;
  }
}
</style>
