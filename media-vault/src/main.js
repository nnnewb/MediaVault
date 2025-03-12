import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import "dayjs/locale/zh-cn";
import "@/assets/main.css";
import axios from "axios";
import components from "@/components/components.js";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";

const app = createApp(App);
const axiosInstance = axios.create({
  baseURL: "/",
  timeout: 1000,
});
app.provide("axios", axiosInstance);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

app.use(router);
app.use(components);
app.use(ElementPlus);

app.mount("#app");
