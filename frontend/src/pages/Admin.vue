<script setup lang="ts"></script>
<template>
  <div class="flex h-full w-full grow flex-col items-center gap-6 px-20 pb-16">
    <div class="text-xl">管理者パネル</div>
    <div class="flex w-full flex-grow rounded-md bg-gray-700">
      <div class="flex w-52 flex-col">
        <router-link
          class="flex h-10 items-center justify-center rounded-tl transition hover:bg-gray-600"
          active-class="bg-gray-600"
          :to="'/admin/contests'"
          >コンテスト</router-link
        >
        <router-link
          class="flex h-10 items-center justify-center transition hover:bg-gray-600"
          active-class="bg-gray-600"
          :to="'/admin/groups'"
          >グループ</router-link
        >
      </div>
      <div class="flex flex-grow flex-col">
        <div class="m-5 h-full">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </div>
</template>
