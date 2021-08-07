
<template>
  <div>
    <button @click="metamaskAuthenticate" class="border border-1 border-grey-600 p-2 rounded-md">
      Sign in with metamask
    </button>
    <div v-if="error" v-html="errorText" class="text-red-700 mt-2"></div>
  </div>

</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import Ethereum from "../types/Ethereum"
export default defineComponent({
  name: "MetamaskLogin",
  setup: () => {
    const count = ref(0)
    const error = ref(false)
    const errorText = ref("")
    async function metamaskAuthenticate(): Promise<void> {
      const ethereum: Ethereum = (window as any).ethereum
      if (ethereum) {
        setTimeout(() => {
          error.value = true;
          errorText.value = "Use the metamask extension to sign in"
        }, 10000)
        //const accounts = await ethereum.request({ method: "eth_requestAccounts" })
      } else {
        error.value = true;
        errorText.value = "Download <a href='https://metamask.io/' target='_blank' class='underline'>metamask</a> to authenticate"
      }
    }
    return { count, error, errorText, metamaskAuthenticate }
  },
});
</script>

<style scoped>
</style>
