import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import "element-plus/theme-chalk/index.css";
import * as ElIconModules from "@element-plus/icons-vue";
import service from './utils/request';

const app = createApp(App)
app.use(store)
app.use(router)

Object.keys(ElIconModules).forEach(function (key) {
  app.component(ElIconModules[key].name, ElIconModules[key]);
});
app.config.globalProperties.$service = service

app.mount('#app')
