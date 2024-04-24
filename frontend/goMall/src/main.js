/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-13 15:36:40
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 17:59:23
 * @FilePath: /goMall/frontend/goMall/src/main.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import { createApp } from 'vue'
import App from './App.vue'
// 引入router
import router from './router/router'
// 引入element-plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// 全局样式
import '@/styles/index.scss'
// 输出环境
console.log("环境地址为: ", import.meta.env.VITE_APP_ENV)
console.log("baseUrl地址为: ", import.meta.env.VITE_APP_BASE_API)


// app实例
const app = createApp(App)
app.use(router)
app.use(ElementPlus)
for(const [key, component] of Object.entries(ElementPlusIconsVue)){
    app.component(key, component)
}
app.mount('#app')
