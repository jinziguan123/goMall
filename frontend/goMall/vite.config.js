/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-13 15:36:40
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 16:24:25
 * @FilePath: /goMall/frontend/goMall/vite.config.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  resolve:{
    alias:{
        '@':path.resolve(__dirname, './src')
    }
  },
  plugins: [vue()],
})
