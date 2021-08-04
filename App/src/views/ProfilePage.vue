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
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent, watchEffect, onMounted, Ref } from "vue";
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

    onMounted( () => {
      const isLoggedInUser = route.params.username === store.getters["authUser/authenticated"]

      if(!isLoggedInUser) {
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
            })
            .catch((error) => console.log(error));
        } else {
          userInfo.value = store.state.getters["authUser/user"]
        }
    })
    return {
      username,
      userInfo,
    }
  }
})
</script>
