import { createApp } from "vue"
import router from "./router/router"
import "./index.css"
import App from "./App.vue"
import store from "./vuex/vuex"

createApp(App).use(router).use(store).mount("#app")
