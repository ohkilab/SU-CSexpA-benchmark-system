<script setup lang="ts">
import { useAdminStateStore } from '../stores/adminState';
import { onMounted, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useBackendStore } from '../stores/backend';

const router = useRouter()
const route = useRoute()
const adminState = useAdminStateStore()
const {backend} = useBackendStore()

onMounted(() => {
  if (adminState.currentPath === '') {
    adminState.currentPath = '/admin/contests'
    router.push(adminState.currentPath)
  } else {
    router.push(adminState.currentPath)
  }
})


watch(
  () => route.path,
  (to) => { //toParams, prevParams
    adminState.currentPath = to
  })
</script>
<template>
  <div class="w-full h-full flex flex-col items-center gap-6 grow px-20 pb-16">
    <div class="text-xl">管理者パネル</div>
    <div class="w-full flex flex-grow rounded-md bg-gray-700">
      <div class="flex flex-col w-52">
        <router-link class="h-10 flex items-center justify-center rounded-tl hover:bg-gray-600 transition" active-class="bg-gray-600" :to="'/admin/contests'">コンテスト</router-link>
        <router-link class="h-10 flex items-center justify-center hover:bg-gray-600 transition"  active-class="bg-gray-600" :to="'/admin/users'">ユーザ</router-link>
      </div>
      <div class="flex flex-col flex-grow">
        <div class="h-full m-5">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </div>
</template> 
