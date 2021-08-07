import { ActionContext } from "vuex"
import { User, emptyUser } from "../../models/user"
import auth from "../../services/authentication"
import backendImageFilePathService from "../../services/backendImageService"

export default {
  namespaced: true,
  state() {
    return {
      checkingAuth: false,
      authenticated: false,
      hasToken: false,
      user: {
        id: 0,
        email: "",
        name: "",
        username: "",
        avatar: "",
        about: "",
      },
    }
  },
  getters: {
    authenticated: (state: UserState) => state.authenticated,
    user: (state: UserState) => state.user,
    checkingAuth: (state: UserState) => state.checkingAuth,
    profileImage: (state: UserState) => state.user.avatar,
    hasToken: (state: UserState) => state.hasToken,
  },
  mutations: {
    authenticated(state: UserState, payload: boolean) {
      state.checkingAuth = false
      state.authenticated = payload
    },
    setUser(state: UserState, payload: User) {
      state.authenticated = true // assumes setUser is only called by logging in
      state.checkingAuth = false
      state.user = payload
      state.user.avatar = backendImageFilePathService(state.user.avatar)
    },
    setUserPicture(state: UserState, payload: string) {
      state.user.avatar = backendImageFilePathService(payload)
    },
    logout(state: UserState) {
      state.authenticated = false
      state.user = emptyUser
    },
  },
  actions: {
    async checkAuth({ commit, state }: ActionContext<UserState, any>) {
      if (state.checkingAuth) {
        return
      }
      state.checkingAuth = true
      const { isLoggedIn, user } = await auth.isLoggedIn()
      if (isLoggedIn) {
        commit("authenticated", true)
        commit("setUser", user)
      } else {
        commit("authenticated", false)
      }
      state.checkingAuth = false
    }
  }
}

interface UserState {
  checkingAuth: boolean,
  authenticated: boolean,
  hasToken: boolean,
  user: User
}
