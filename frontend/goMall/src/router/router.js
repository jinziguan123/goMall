/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-13 15:47:52
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 16:32:22
 * @FilePath: /goMall/frontend/goMall/src/router/router.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import {createRouter, createWebHistory} from 'vue-router'

/**
 * 路由表
 */
const router = createRouter({
    // 去掉url的#
    history: createWebHistory(),
    routes:[
        {path:'/login', component: ()=> import('@/views/login.vue')}
    ]
})

export default router