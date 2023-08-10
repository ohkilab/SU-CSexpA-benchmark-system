<script setup lang="ts">
import { ref } from "vue";

defineProps({
  errMsg: String,
});

const id = ref("");
const password = ref("");

const emit = defineEmits(["login"]);
</script>
<template>
  <form
    class="flex h-full w-full flex-col items-center justify-center gap-5 px-5 text-xl md:w-96"
    @submit.prevent="() => emit('login', id, password)"
  >
    <!-- <router-link class="border border-gray-300 p-2 rounded hover:bg-gray-600 transition" :to="'/ranking'">ログインなしでランキングを見る</router-link> -->

    <div class="text-red-500">
      {{ errMsg }}
    </div>
    <input
      class="w-full rounded bg-gray-700 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
      placeholder="グループ名"
      type="text"
      v-model="id"
    />
    <input
      class="w-full rounded bg-gray-700 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
      placeholder="パスワード"
      type="password"
      v-model="password"
    />
    <button
      class="w-full rounded border py-2 transition hover:bg-gray-700 md:w-2/3"
    >
      ログイン
    </button>
  </form>
</template>
