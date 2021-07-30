import { createStore, ActionContext } from "vuex"
import { User, emptyUser } from "../models/user"
import auth from "../services/authentication"

export default createStore({
  state() {
    return {
      checkingAuth: false,
      authenticated: false,
      user: {
        id: "",
        email: "",
        name: "",
        username: "",
      },
    }
  },
  getters: {
    authenticated: (state: State) => state.authenticated,
    user: (state: State) => state.user,
    checkingAuth: (state: State) => state.checkingAuth,
  },
  mutations: {
    authenticated(state: State, payload: boolean) {
      state.checkingAuth = false
      state.authenticated = payload
    },
    setUser(state: State, payload: User) {
      state.authenticated = true // assumes setUser is only called by logging in
      state.checkingAuth = false
      state.user = payload
    },
    logout(state: State) {
      state.authenticated = false
      state.user = emptyUser
    },
  },
  actions: {
    async logout(context: ActionContext<State, any>) {
      await auth.logout()
      context.commit("logout")
    },
  },
})

interface State {
  checkingAuth: boolean,
  authenticated: boolean,
  user: User
}
