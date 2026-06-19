<script setup lang="ts">
import { useAdminStateStore } from "../stores/adminState";
import { onMounted, watch } from "vue";
import { useRouter, useRoute } from "vue-router";

const router = useRouter();
const route = useRoute();
const adminState = useAdminStateStore();

onMounted(() => {
  if (adminState.currentPath === "" || adminState.currentPath === "/admin") {
    adminState.currentPath = "/admin/contests";
    router.replace(adminState.currentPath);
  } else {
    router.replace(adminState.currentPath);
  }
});

watch(
  () => route.path,
  (to) => {
    //toParams, prevParams
    adminState.currentPath = to;
  },
);
</script>
<template>
  <div
    class="flex h-full w-full grow flex-col items-center gap-6 px-4 pb-16 sm:px-8 lg:px-20"
  >
    <div class="text-xl">管理者パネル</div>
    <div
      class="flex w-full min-w-0 flex-grow overflow-x-auto rounded-md bg-gray-700"
    >
      <div class="flex w-40 shrink-0 flex-col sm:w-52">
        <router-link
          class="flex h-10 items-center justify-center whitespace-nowrap rounded-tl px-3 transition hover:bg-gray-600"
          active-class="bg-gray-600"
          :to="'/admin/contests'"
          >コンテスト</router-link
        >
        <router-link
          class="flex h-10 items-center justify-center whitespace-nowrap px-3 transition hover:bg-gray-600"
          active-class="bg-gray-600"
          :to="'/admin/groups'"
          >グループ</router-link
        >
      </div>
      <div class="flex min-w-[720px] flex-grow flex-col">
        <div class="m-5 h-full">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </div>
</template>
