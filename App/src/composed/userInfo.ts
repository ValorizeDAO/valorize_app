import { onMounted, ref } from "vue"
import { useStore } from "vuex"
import backendImageFilePathService from "../services/backendImageService"
import {emptyUser} from "../models/User";

export default function composeUserInfo(username: string) {
  const store = useStore()
  const userStatuses = ["INIT", "LOADING", "SUCCESS", "FAIL"]
  const userStatus = ref(userStatuses[0])
  const userInfo = ref(emptyUser)
  const tokenInfo = ref([])
  const showImage = ref(false)

  onMounted(() => {
    store.dispatch("authUser/checkAuth")
    fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/users/" + username)
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
      .catch((error) => console.log(error))
  })
  return {
    username,
    userInfo,
    tokenInfo,
    showImage,
  }
}
