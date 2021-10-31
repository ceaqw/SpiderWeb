import { createApp } from 'vue'

import App from './App'
import router from './route'

import Axios from 'axios'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './style/index.css'


const app = createApp(App)
app.config.globalProperties.Axios = Axios
app.use(router)
app.use(ElementPlus)

app.mount('#app')
