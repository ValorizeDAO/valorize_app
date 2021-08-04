<template>
  <div id="profile-page" class="grid grid-cols-12">
    <div id="user-info" class="col-span-5 pl-16 bg-purple-200 min-h-screen pt-8 border-r-2 border-black">
      <h1 class="text-3xl font-black">{{ userInfo.username }}</h1>
      <ImageContainer class="my-8">
        <img class="h-52 w-52 object-cover" :src="userInfo.avatar" alt="">
      </ImageContainer>
      <p class="mt-8">{{ userInfo.about }}</p>
    </div>
    <div id="token-info" class="col-span-7 px-16 min-h-screen pt-8 border-black">
      <div v-if="tokenInfo">
        <h2 class="font-black text-2xl">
          {{tokenInfo.name}} ({{tokenInfo.symbol}})
        </h2>
        <p class="my-4">
          Address: {{tokenInfo.address}}
        </p>
        <transition name="fade" mode="out-in">
          <img v-if="showImage" class="rounded border-2 border-black p-6 bg-purple-100" :src="'https://chart.googleapis.com/chart?chs=150x150&cht=qr&chl='+tokenInfo.address+'&choe=UTF-8'" alt="">
          <button v-else @click.prevent="showImage=true" class="btn bg-burple-100 mb-4 w-48">Show QR Code</button>
        </transition>
        <div id="coin-data" class="flex justify-between flex-wrap">
          <div>Price Per Coin: {{ tokenPrice }}</div>
          <div>Total Supply: {{ tokenCap }}</div>
          <div>USD locked in <strong>{{tokenInfo.symbol}}</strong> {{ tokenEthBalance * ethPrice }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent, onMounted, computed } from "vue";
import { useRoute } from 'vue-router';
import {useStore} from "vuex";
import backendImageFilePathService from "../services/backendImageService"
import ImageContainer from "../components/ImageContainer.vue";
export default defineComponent({
  name: 'ProfilePage',
  components: { ImageContainer },
  setup() {
    return {
      ...composeUserInfo(),
      ...composeTokenInfo()
    }
  }
})
function composeUserInfo() {
  const route = useRoute()
  const store = useStore();
  const username = ref ("")
  const userStatuses = ["INIT", "LOADING", "SUCCESS", "FAIL"]
  const userStatus = ref(userStatuses[0])
  const userInfo = ref({})
  const tokenInfo = ref([])
  const showImage = ref(false)

  onMounted( () => {
    store.dispatch("authUser/checkAuth")
    fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/users/" + route.params.username)
        .then((response) => {
          if (response.status !== 200) {
            userStatus.value = userStatuses[3]
            return
          }
          return response.json()
        })
        .then((result) => {
          userInfo.value = result
          userInfo.value.avatar = backendImageFilePathService(result.avatar)
          tokenInfo.value = result.token
        })
        .catch((error) => console.log(error));
  })
  return {
    username,
    userInfo,
    tokenInfo,
    showImage
  }
}
function composeTokenInfo() {
  const route = useRoute()
  const store = useStore();
  const userStatuses = ["INIT", "LOADING", "SUCCESS", "FAIL"]
  const userStatus = ref(userStatuses[0])
  const tokenCap = ref(0)
  const tokenEthBalance = ref(0)
  const ethPrice = ref(0.0)
  const tokenPrice = computed(() => {
    return (ethPrice.value * tokenEthBalance.value )/ tokenCap.value
  })
  onMounted(async () => {
    fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/users/" + route.params.username + "/token")
      .then((response) => {
        if (response.status !== 200) {
          userStatus.value = userStatuses[3]
          return
        }
        return response.json()
      })
      .then((result) => {
        tokenEthBalance.value = parseInt(result["ether_staked"])
        tokenCap.value = parseInt(result["total_minted"])
      })
      .catch((error) => console.log(error));
    fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/utils/price")
      .then((response) => {
        if (response.status !== 200) {
          userStatus.value = userStatuses[3]
          return
        }
        return response.json()
      })
      .then((result) => {
        ethPrice.value = parseFloat(result["USD"])
      })
      .catch((error) => console.log(error));
  })
    //fetch utils/prce to map to tokenPrice

  return {
    tokenPrice,
    tokenCap,
    tokenEthBalance,
    ethPrice,
  }
}
</script>
