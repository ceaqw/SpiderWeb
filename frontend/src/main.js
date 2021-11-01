import { createApp } from 'vue'

import App from './App'
import router from './route'
import store from '@/store'

import Axios from 'axios'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './style/index.css'

const app = createApp(App)
app.config.globalProperties.Axios = Axios
app.use(router)
app.use(ElementPlus)
app.use(store)

app.mount('#app')
