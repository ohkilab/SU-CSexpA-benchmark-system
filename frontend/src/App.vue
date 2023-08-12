<script setup lang="ts">
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { BackendServiceClient } from 'proto-gen-web/services/backend/services.client';
import { onMounted, ref, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import { useStateStore } from './stores/state'
import { useBackendStore } from './stores/backend'
import { useAdminStateStore } from './stores/adminState';
import { Role } from 'proto-gen-web/services/backend/resources';

const loggedIn = ref(false)
const state = useStateStore()
const adminState = useAdminStateStore()
const backendStore = useBackendStore()
const router = useRouter()
const route = useRoute()

const handleLogout = () => {
  // clear all credentials
  loggedIn.value = false
  state.lastResult = 0
  router.push('/login')
  adminState.$reset()
  state.$reset()
}

const PROD = import.meta.env.PROD

const errMsg = ref('')

const handleLogin = (id: string, password: string) => {

  backendStore.backend = new BackendServiceClient(
    new GrpcWebFetchTransport({
      baseUrl: import.meta.env.PROD ? `http://${window.location.hostname}:8080` : state.devBaseUrl
    })
  )

  backendStore.backend.postLogin({ id, password }).then(value => {
    if (import.meta.env.DEV) console.log('Login', value)

    state.token = value.response.token
    state.group = id
    state.role = value.response.group?.role ?? Role.GUEST
    loggedIn.value = true
    errMsg.value = ''
    router.push('/contests')
  }).catch(err => {
    console.log(err)
    errMsg.value = err.message
  })
}

watch(loggedIn, loggedIn => {
  if (loggedIn && route.name === 'index') {
    router.push('/benchmark')
  }
})

onMounted(() => {
  // if(import.meta.env.DEV) BigInt.prototype.toJSON = function() {return this.toString()}

  // try login with token
  if (state.token) {
    console.log('try login with token')
    backendStore.backend = new BackendServiceClient(
      new GrpcWebFetchTransport({
        baseUrl: import.meta.env.PROD ? `http://${window.location.hostname}:8080` : state.devBaseUrl
      })
    )

    backendStore.backend.verifyToken({ token: state.token })
      .then(_ => {
        console.log(_)
        // successfully logged in with token, load token and group name into app
        loggedIn.value = true
      }).catch(_ => {
        console.log(_)
        // login with token failed
        router.push('/login')
        errMsg.value = 'Session expired, please login again'
      })
  } else {
    router.push('/login')
  }
})
</script>
<template>
  <div class="text-white items-center bg-gray-800 flex flex-col min-h-screen">
    <!-- app bar -->
    <div class="w-full h-16 items-center bg-gray-700 flex shadow-md shadow-gray-950 px-5">
      <div v-if="loggedIn" class="w-40 text-center p-2 border border-gray-500">グループ：{{ state.group }}</div>
      <div class="w-32 flex justify-center">
        <button @click="state.debug = !state.debug" v-if="!PROD"
          class="p-2 rounded border border-red-500 transition hover:bg-red-700">Debug: {{ state.debug ? 'on' :
            'off' }}</button>
      </div>
      <div class="mx-auto">
        情報科学実験A：ベンチマークサーバ
      </div>
      <div class="w-32 flex justify-center">
        <button v-if="state.role === Role.ADMIN && !route.name?.toString().startsWith('admin')" @click="router.push('/admin')"
          class="p-2 rounded transition border border-red-500 hover:bg-red-700">管理者モード</button>
        <button v-else
          @click="router.push('/contests')"
          class="p-2 rounded border border-blue-500 transition hover:bg-blue-700">ユーザモード</button>
      </div>
      <button @click="handleLogout" v-if="loggedIn"
        class="p-2 w-40 rounded bg-red-500 transition hover:bg-red-700">ログアウト</button>
    </div>
    <!-- debug mode -->
    <fieldset v-if="state.debug" class="mx-8 border border-red-500 p-2 flex flex-col gap-2">
      <legend>Debug Panel</legend>
      <!-- <pre class="break-all whitespace-pre-wrap">state: {{JSON.stringify(state, null, 4)}}</pre> -->
      <pre class="break-all whitespace-pre-wrap">token: {{ state.token }}</pre>
      <pre class="break-all whitespace-pre-wrap">group: {{ state.benchmarking }}</pre>
      <pre class="break-all whitespace-pre-wrap">benchmarking: {{ state.benchmarkInterval }}</pre>
      <pre class="break-all whitespace-pre-wrap">lastResult: {{ state.lastResult }}</pre>
      <pre class="break-all whitespace-pre-wrap">showResult: {{ state.showResult }}</pre>
      <pre class="break-all whitespace-pre-wrap">current: {{ state.current }}</pre>
      <pre class="break-all whitespace-pre-wrap">result: {{ state.result }}</pre>
      <pre class="break-all whitespace-pre-wrap">debug: {{ state.debug }}</pre>
      <pre class="break-all whitespace-pre-wrap">devBaseUrl: {{ state.devBaseUrl }}</pre>
      <pre class="break-all whitespace-pre-wrap">contestSlug: {{ state.contestSlug }}</pre>


      <button class="bg-green-500 p-2" @click="state.benchmarking = !state.benchmarking">Toggle benchmarking</button>
      <input type="text" v-model="state.devBaseUrl" placeholder="baseUrl"
        class="bg-gray-700 p-2 rounded transition hover:bg-gray-600 focus:outline-none" />
    </fieldset>
    <!-- main view -->
    <div v-if="loggedIn" class="w-full h-full flex flex-col gap-6 flex-grow align-middle items-center">
      <div v-if="$route.name !== 'contests'" class="text-xl mt-5">{{ state.selectedContestName }}</div>
      <div v-if="loggedIn && !state.benchmarking && $route.name !== 'contests'" class="flex w-full px-12 text-lg">
        <router-link to="/contests"
          class="w-48 rounded transition hover:scale-105 shadow-md shadow-black p-2 text-center border border-gray-500 bg-red-500 mr-auto">
          &#x2190 <div class="inline">コンテスト一覧</div> </router-link>
        <div class="mx-auto flex gap-5">
          <router-link class="p-2 rounded shadow-md shadow-black hover:scale-105 transition border border-gray-500"
            active-class="bg-blue-500" to="/benchmark">ベンチマーク</router-link>
          <router-link class="p-2 rounded shadow-md shadow-black hover:scale-105 transition border border-gray-500"
            active-class="bg-blue-500" to="/submissions">結果一覧</router-link>
          <router-link class="p-2 rounded shadow-md shadow-black hover:scale-105 transition border border-gray-500"
            active-class="bg-blue-500" to="/ranking">ランキング</router-link>
        </div>
        <div to="/contests" class="w-48 p-2 ml-auto"></div>
      </div>
      <div v-if="errMsg" class="text-red-500">{{ errMsg }}</div>
      <router-view></router-view>
    </div>
    <div v-else class="flex flex-col grow justify-center items-center">
      <div class="text-red-500 h-10">{{ errMsg }}</div>
      <router-view @login="(id: string, password: string) => { handleLogin(id, password) }"></router-view>
    </div>
    <!-- footer -->
    <div class="flex items-center justify-center bg-gray-700 w-full h-8">© 2023 <a href="https://sec.inf.shizuoka.ac.jp/"
        class="text-blue-500 mx-1" target="_blank">Ohkilab.</a> All rights
      reserved.</div>
  </div>
</template>
