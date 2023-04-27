<script setup lang="ts">
import Login from './pages/Login.vue'
const loggedIn:boolean = false

const groups: { name: string; score: number }[] = [
  {
    name: 'a1',
    score: 1000
  },
  {
    name: 'b2',
    score: 10000
  },
  {
    name: 'a2',
    score: 2000
  },
  {
    name: 'a3',
    score: 1234
  },
  {
    name: 'a4',
    score: 5431
  },
  {
    name: 'b2',
    score: 50
  },
  {
    name: 'b3',
    score: 51
  }
]
</script>
<template>
  <div
    class="text-white items-center bg-gray-800 flex flex-col gap-6 min-h-screen"
  >
    <!-- app bar -->
    <div
      class="w-screen h-16 items-center bg-gray-700 flex shadow-md shadow-gray-950"
    >
      <div class="mx-auto text-lg sm:text-xl">
        情報科学実験A：ベンチマークサーバランキング
      </div>
    </div>
    <div v-if="loggedIn" class="flex gap-5 text-lg">
        <!-- TODO: fix active class -->
        <router-link
          class="p-2 rounded bg-gray-500 shadow-md shadow-black hover:scale-105 transition"
          active-class="bg-blue-500 underline"
          to="/benchmark"
          >ベンチマーク</router-link
        >
        <router-link
          class="p-2 rounded bg-gray-500 shadow-md shadow-black hover:scale-105 transition"
          active-class="bg-blue-500 underline"
          to="/ranking"
          >ランキング</router-link
        >
      </div>
      <router-view v-if="loggedIn"></router-view>
      <Login class="mt-auto" v-else></Login>
    <!-- footer -->
    <div class="flex items-center justify-center bg-gray-700 w-full mt-auto">footer text</div>
  </div>
</template>
