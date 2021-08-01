import { createStore } from "vuex"
import authUser from "./modules/authUserModule"

export default createStore({
  modules: {
    authUser,
  },
})
