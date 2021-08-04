<template>
  <div id="profile-page" class="grid grid-cols-12">
    <div id="user-info" class="col-span-5 pl-16 bg-purple-200 min-h-screen pt-8 border-r-2 border-black">
      <h1 class="text-3xl font-black">{{ userInfo.username }}</h1>
      <ImageContainer class="my-8">
        <img class="h-52 w-52 object-cover" :src="userInfo.avatar" alt="">
      </ImageContainer>
      <p class="mt-8">{{ userInfo.about }}</p>
    </div>
    <div id="token-info" class="col-span-7 pl-16 min-h-screen pt-8 border-r-2 border-black">
      <div v-if="tokenInfo">
        <h2 class="font-black text-2xl">
          {{tokenInfo.name}} ({{tokenInfo.symbol}})
        </h2>
        Address: {{tokenInfo.address}}
        <transition name="fade" mode="out-in">
          <img v-if="showImage" class="rounded border-2 border-black p-6 bg-purple-100" :src="'https://chart.googleapis.com/chart?chs=150x150&cht=qr&chl='+tokenInfo.address+'&choe=UTF-8'" alt="">
          <button v-else @click.prevent="showImage=true" class="btn bg-burple-100 my-12 w-48">Show QR Code</button>
        </transition>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent, onMounted } from "vue";
import { useRoute } from 'vue-router';
import {useStore} from "vuex";
import backendImageFilePathService from "../services/backendImageService"
import ImageContainer from "../components/ImageContainer.vue";
export default defineComponent({
  name: 'ProfilePage',
  components: { ImageContainer },
  setup() {
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
})
</script>
