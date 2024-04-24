/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-13 17:51:41
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 17:55:01
 * @FilePath: /goMall/frontend/goMall/src/utils/request.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/**
 * axios封装
 */

import axios from "axios";
import { de } from "element-plus/es/locale";

// 创建axios
const service = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_API,
    timeout: 5000
})

// 核心函数
function request(options) {
    options.method = options.method || 'get'
    if (options.method.toLowerCase() == 'get'){
        options.params = options.data
    }
    service.defaults.baseURL = import.meta.env.VITE_APP_BASE_API
    return service(options)
}

export default request