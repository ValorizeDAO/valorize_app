import { createApp } from "vue";
import router from "./router/router";
import "./index.css";
import App from "./App.vue";

createApp(App).use(router).mount("#app");
