<script setup lang="ts">
import { reactive, Ref, ref } from 'vue';
import Login from './pages/Login.vue'
import {useStateStore} from './stores/state'

const loggedIn = ref(false)
const token:Ref<string> = ref('')
const msg:Ref<string> = ref('')
const state = useStateStore()

const handleLoggedIn = (t:string) => {
  token.value = t
  loggedIn.value = true
  state.token = t
}
</script>
<template>
  <div
    class="text-white items-center bg-gray-800 flex flex-col gap-6 min-h-screen"
  >
    <!-- app bar -->
    <div
      class="w-full h-16 items-center bg-gray-700 flex shadow-md shadow-gray-950"
    >
      <div class="mx-auto text-lg sm:text-xl">
        情報科学実験A：ベンチマークサーバランキング
      </div>
    </div>
    <div v-if="loggedIn" class="flex gap-5 text-lg">
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
          to="/ranking"
          >ランキング</router-link
        >
      </div>
      <router-view v-if="loggedIn"></router-view>
      <Login class="mt-auto" @logged-in="t => {handleLoggedIn(t)}" v-else></Login>
    <!-- footer -->
    <div class="flex items-center justify-center bg-gray-700 w-full mt-auto">© 2023 Ohkilab. All rights reserved.</div>
  </div>
</template>
