import { createStore } from "vuex";

// Create a new store instance.
export default createStore({
  state() {
    return {
      authenticated: false,
    };
  },
  mutations: {
    authenticated(state: { authenticated: boolean }, payload: boolean) {
      state.authenticated = payload;
    },
  },
});
