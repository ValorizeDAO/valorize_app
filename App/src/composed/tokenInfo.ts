import { computed, onMounted, ref } from "vue"

export default function composeTokenInfo(username: string) {
  const userStatuses = ["INIT", "LOADING", "SUCCESS", "FAIL"]
  const userStatus = ref<string>(userStatuses[0])
  const tokenCap = ref<number>(0)
  const tokenEthBalance = ref<number>(0)
  const ethPrice = ref<number>(0.0)
  const tokenPrice = computed(() => {
    return (ethPrice.value * tokenEthBalance.value) / (tokenCap.value === 0 ? 1 : tokenCap.value)
  })
  onMounted(async () => {
    fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/users/" + username + "/token")
      .then((response) => {
        if (response.status !== 200) {
          userStatus.value = userStatuses[3]
          return
        }
        return response.json()
      })
      .then((result) => {
        tokenEthBalance.value = parseInt(result.ether_staked)
        tokenCap.value = parseInt(result.total_minted)
      })
      .catch((error) => console.log(error))
    fetch(import.meta.env.VITE_BACKEND_URL + "/api/v0/utils/price")
      .then((response) => {
        if (response.status !== 200) {
          userStatus.value = userStatuses[3]
          return
        }
        return response.json()
      })
      .then((result) => {
        ethPrice.value = parseFloat(result.USD)
      })
      .catch((error) => console.log(error))
  })
  return {
    tokenPrice,
    tokenCap,
    tokenEthBalance,
    ethPrice,
  }
}
