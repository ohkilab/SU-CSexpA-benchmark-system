<script setup lang="ts">
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client';
import { ref } from 'vue';

import { useStateStore, IState } from '../stores/state';

const id = ref('')
const password = ref('')
const errMsg = ref('')

const state:IState = useStateStore()

const emit = defineEmits(['loggedIn'])

const backend = new BackendServiceClient(
  new GrpcWebFetchTransport({
    baseUrl: "http://localhost:8080"
  })
)

const handleLogin = () => {
  backend.postLogin({ id: id.value , password: password.value }).then(value => {
    console.log(value.response)
    state.id = id
    emit('loggedIn', value.response.token)
  }).catch(err => {
    console.log(err.message)
    errMsg.value = err.message
  })
}
</script>
<template>
  <form class="flex flex-col items-center justify-center h-full w-full px-5 md:w-96 gap-5 text-xl" @submit.prevent="handleLogin">
    <div class="text-red-500">
      {{errMsg}}
    </div>
    <input class="w-full rounded bg-gray-700 p-2 hover:bg-gray-600 transition focus:outline-none focus:bg-gray-600" placeholder="グループ名" type="text" v-model="id">
    <input class="w-full rounded bg-gray-700 p-2 hover:bg-gray-600 transition focus:outline-none focus:bg-gray-600" placeholder="パスワード" type="password" v-model="password">
    <button class="w-full md:w-2/3 border rounded py-2 hover:bg-gray-700 transition">ログイン</button>
  </form>
</template>
