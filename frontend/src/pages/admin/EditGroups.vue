<script setup lang="ts">
</script>
<template>
  <div class="flex flex-col gap-2 h-full">
    <h1 class="text-xl">グループ</h1>
    <h2>グループ名</h2>
    <input class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
           placeholder="タイトル" type="text">
    <h2>パスワード</h2>
    <input type="number"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none" value="0">
    <h2>年度</h2>
    <input type="number"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none" value="2000">
    <h2>ロール</h2>
    <input type=""
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none">

    <button class="mt-auto bg-blue-500 w-fit p-2 rounded">グループ作成</button>
  </div>
</template>
