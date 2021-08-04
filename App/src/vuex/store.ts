import { createStore } from "vuex"
import authUser from "./modules/authUserModule"
import userToken from "./modules/userTokenModule"

export default createStore({
  modules: {
    authUser,
    userToken
  },
})
