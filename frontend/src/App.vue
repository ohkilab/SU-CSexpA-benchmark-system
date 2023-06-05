<script setup lang="ts">
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client';
import { onMounted, Ref, ref, watch } from 'vue';
import Login from './pages/Login.vue'
import {useStateStore} from './stores/state'

const loggedIn = ref(false)
const token:Ref<string> = ref('')
const group:Ref<string> = ref('')
const state = useStateStore()

const handleLogout = () => {
  // clear all credentials
  loggedIn.value = false
  token.value = ''
  group.value = ''
}

const backend = new BackendServiceClient(
  new GrpcWebFetchTransport({
    baseUrl: import.meta.env.PROD ? `http://${window.location.hostname}:8080` : state.devBaseUrl
  })
)

const PROD = import.meta.env.PROD

const errMsg = ref('')

const handleLogin = (id:string, password:string) => {
  backend.postLogin({ id, password }).then(value => {
    if(import.meta.env.DEV) console.log('Login', value)
    token.value = value.response.token
    group.value = id
    loggedIn.value = true
  }).catch(err => {
    console.log(err)
    errMsg.value = err.message
  })
}

// update localStorage and state based on refs
watch(
  group,
  group => {
    state.group = group
    localStorage.setItem('group', group)
  },
  {deep: true}
)

watch(
  token,
  token => {
    state.token = token
    localStorage.setItem('token', token)
  },
  {deep: true}
)

onMounted(() => {
  // try login with token
  if(localStorage.getItem('token')) {
    let opt = {meta: {'authorization' : 'Bearer ' + localStorage.getItem('token')}}

    backend.getRanking({
      year: 2023,
      containGuest: false
    },opt).then(_ => {
      // successfully logged in with token, load token and group name into app
      token.value = localStorage.getItem('token') ?? ''
      group.value = localStorage.getItem('group') ?? ''
      loggedIn.value = true
    }).catch(_ => {
      // login with token failed
      localStorage.removeItem('token')
      localStorage.removeItem('group')

      errMsg.value = 'Session expired, please login again'
    })
  }
})
</script>
<template>
  <div
    class="text-white items-center bg-gray-800 flex flex-col gap-6 min-h-screen"
  >
    <!-- app bar -->
    <div
      class="w-full h-16 items-center bg-gray-700 flex shadow-md shadow-gray-950 px-5"
    >
    <div v-if="loggedIn" class="w-32 p-2 border border-gray-500">グループ：{{state.group}}</div>
    <button @click="state.debug = !state.debug" v-if="!PROD" class="p-2 ml-2 w-32 rounded border border-red-500 transition hover:bg-red-700">Debug: {{state.debug ? 'on' : 'off'}}</button>
      <div class="mx-auto text-lg sm:text-xl">
        情報科学実験A：ベンチマークサーバ
      </div>
      <button @click="handleLogout" v-if="loggedIn" class="p-2 w-32 rounded border border-red-500 transition hover:bg-red-700">ログアウト</button>
    </div>
    <!-- debug mode -->
    <fieldset v-if="state.debug" class="mx-8 border border-red-500 p-2 flex flex-col gap-2">
      <legend>Debug Panel</legend>
      <pre class="break-all whitespace-pre-wrap">state: {{JSON.stringify(state, null, 4)}}</pre>
      <button class="bg-green-500 p-2" @click="state.benchmarking = !state.benchmarking">Toggle benchmarking</button>
      <input type="text" v-model="state.devBaseUrl" placeholder="baseUrl" class="bg-gray-700 p-2 rounded transition hover:bg-gray-600 focus:outline-none" />
    </fieldset>
    <div v-if="loggedIn && !state.benchmarking" class="flex gap-5 text-lg">
        <!-- TODO: fix active class -->
        <router-link
          class="p-2 rounded shadow-md shadow-black hover:scale-105 transition border border-gray-500"
          active-class="bg-blue-500"
          to="/benchmark"
          >ベンチマーク</router-link
        >
        <router-link
          class="p-2 rounded shadow-md shadow-black hover:scale-105 transition border border-gray-500"
          active-class="bg-blue-500"
          to="/submissions"
          >結果一覧</router-link
        >
        <router-link
          class="p-2 rounded shadow-md shadow-black hover:scale-105 transition border border-gray-500"
          active-class="bg-blue-500"
          to="/ranking"
          >ランキング</router-link
        >
      </div>
      <router-view v-if="loggedIn"></router-view>
      <Login class="mt-auto" :err-msg="errMsg" @login=" (id, password) => {handleLogin(id, password)}" v-else></Login>
    <!-- footer -->
    <div class="flex items-center justify-center bg-gray-700 w-full mt-auto h-8">© 2023 <a href="https://sec.inf.shizuoka.ac.jp/" class="text-blue-500 mx-1" target="_blank">Ohkilab.</a> All rights reserved.</div>
  </div>
</template>
