
<template>
  <div v-if="tokenInfo">
    <h2 class="font-black text-2xl">
      {{tokenInfo.name}} ({{tokenInfo.symbol}})
    </h2>
    <p class="my-4">
      Address: {{tokenInfo.address}}
    </p>
    <div id="token-info-actions" class="mx-auto flex justify-center h-56 items-center">

      <transition name="fade" mode="out-in">
        <ImageContainer v-if="showImage">
          <img class="rounded border-2 border-black p-6 bg-purple-100 mb-6 text-center md:text-left" :src="'https://chart.googleapis.com/chart?chs=150x150&cht=qr&chl='+tokenInfo.address+'&choe=UTF-8'" alt="">
        </ImageContainer>
        <button v-else @click.prevent="showImage=true" class="btn bg-burple-100 mb-4 w-48">Show QR Code</button>
      </transition>
    </div>
    <div id="coin-data" class="flex justify-between flex-wrap">
    <div><h3 class="text-xl font-black">Price Per Coin: {{ tokenPrice }} </h3></div>
    <div><h3 class="text-xl font-black">Total Supply: {{ tokenCap }} </h3></div>
    <div><h3 class="text-xl font-black">USD locked in <strong>{{tokenInfo.symbol}}</strong>: $ {{ tokenEthBalance * ethPrice }}</h3></div>
  </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import composeTokenInfo from "../composed/tokenInfo"
import composeUserInfo from "../composed/userInfo"
import ImageContainer from "./ImageContainer.vue";
export default defineComponent({
  name: "TokenInfoComponent",
  props: ["username"],
  components: { ImageContainer },
  setup: (props) => {
    return {
      ...composeTokenInfo(props.username),
      ...composeUserInfo(props.username)
    }
  },
});
</script>

<style scoped>
</style>
