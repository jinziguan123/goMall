<!--
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-13 15:49:14
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 17:47:27
 * @FilePath: /goMall/frontend/goMall/src/views/login.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div class="login-container">
        <el-form class="login-form" :rules="rules" :model="loginForm">
            <div class="title-container">goMall System 登陆界面</div>
            <!--用户名-->
            <el-form-item style="margin-bottom: 30px;" prop="username">
                <el-input placeholder="请输入用户名" name="username" type="text" prefix-icon="User"
                    v-model="loginForm.username"></el-input>
            </el-form-item>
            <!--密码-->
            <el-form-item style="margin-bottom: 30px;" prop="password">
                <el-input placeholder="请输入密码" name="password" :type="flagType" v-model="loginForm.password">
                    <template #prefix>
                        <el-icon>
                            <lock/>
                        </el-icon>
                    </template>
                    <template #suffix>
                        <span @click="changeView">
                            <el-icon v-if="flag == true" style="cursor: pointer;">
                                <Hide/>
                            </el-icon>
                            <el-icon v-else-if="flag == false" style="cursor: pointer;">
                                <View/>
                            </el-icon>
                        </span>
                    </template>
                </el-input>
            </el-form-item>
            <!--登陆和注册按钮-->
            <el-button type="primary" style="width: 100%; margin-bottom: 30px">登陆</el-button>
        </el-form>
    </div>
</template>

<script setup>
import { ref } from 'vue'

// 表单校验
const rules = {
    username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
}

const loginForm = ref({})

// 切换小眼睛实现
const flagType = ref("password")
const flag = ref(true)
const changeView = ()=>{
    flag.value = !flag.value
    flagType.value = flag.value ? "password" : "text"
}
</script>

<style lang="scss">
.login-container {
    background-color: #2d3a4b;
    height: 100%;

    .login-form {
        width: 400px;
        border-radius: 1px;
        position: absolute;
        left: 50%;
        top: 30%;
        transform: translate(-50%);

        .title-container {
            font-size: 30px;
            line-height: 1.5;
            text-align: center;
            margin-bottom: 30px;
            font-family: Impact, Haettenschweiler, 'Arial Narrow Bold', sans-serif;
        }
    }
}
</style>