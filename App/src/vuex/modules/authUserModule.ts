import { ActionContext } from "vuex"
import { User, emptyUser } from "../../models/user"

export default {
  namespaced: true,
  state() {
    return {
      checkingAuth: false,
      authenticated: false,
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
      state.user.avatar = setUserPicturePrefix(state.user.avatar)
    },
    setUserPicture(state: UserState, payload: string) {
      state.user.avatar = setUserPicturePrefix(payload)
    },
    logout(state: UserState) {
      state.authenticated = false
      state.user = emptyUser
    },
  },
}

function setUserPicturePrefix(filename: string): string {
  return (
    import.meta.env.VITE_BACKEND_URL +
    "/static/images/" +
    (filename || "default_avatar.jpg")
  )
}

interface UserState {
  checkingAuth: boolean,
  authenticated: boolean,
  user: User
}
