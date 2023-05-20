<script setup lang="ts">
import { ref } from 'vue';

defineProps({
  errMsg: String
})

const id = ref('')
const password = ref('')

const emit = defineEmits(['login'])
</script>
<template>
  <form class="flex flex-col items-center justify-center h-full w-full px-5 md:w-96 gap-5 text-xl" @submit.prevent="() => emit('login', id, password)">
    <div class="text-red-500">
      {{errMsg}}
    </div>
    <input class="w-full rounded bg-gray-700 p-2 hover:bg-gray-600 transition focus:outline-none focus:bg-gray-600" placeholder="グループ名" type="text" v-model="id">
    <input class="w-full rounded bg-gray-700 p-2 hover:bg-gray-600 transition focus:outline-none focus:bg-gray-600" placeholder="パスワード" type="password" v-model="password">
    <button class="w-full md:w-2/3 border rounded py-2 hover:bg-gray-700 transition">ログイン</button>
  </form>
</template>
