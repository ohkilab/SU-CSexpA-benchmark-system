<script setup lang="ts">
import { CreateContestRequest } from 'proto-gen-web/services/backend/messages';
import { useAdminStateStore } from '../../stores/adminState';
import { ref, Ref } from 'vue';

const { admin } = useAdminStateStore();

const contest: Ref<CreateContestRequest> = ref({})

const createContest = () => {
  console.log(contest.value)
}

</script>
<template>
  <div class="flex flex-col gap-2 h-full">
    <h1 class="text-xl">コンテスト</h1>
    <h2>タイトル</h2>
    <input class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
           placeholder="タイトル" type="text">
    <h2>提出回数</h2>
    <input type="number"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none" value="0">
    <h2>開始日</h2>
    <input type="date"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none" value="0">
    <h2>終了日</h2>
    <input type="date"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none" value="0">
    <h2>slug</h2>
    <input class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
      placeholder="slug" type="text">

    <button class="mt-auto bg-blue-500 w-fit p-2 rounded">コンテストを作成</button>
  </div>
</template>
