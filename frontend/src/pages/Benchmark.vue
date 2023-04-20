<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { onMounted, reactive } from 'vue'

interface Status {
  benchmarking: boolean,
  current: number
  size: number
}

const status: Status = reactive({
  benchmarking: false,
  current: 0,
  size: 20
})

onMounted(() => {
  // setInterval(() => {
  //   status.current++
  //   if (status.current >= status.size) status.current = 0
  // }, 1000)
})
</script>
<template>
  <div class="flex flex-col my-auto">
    <div v-if="status.benchmarking">
      <div class="flex flex-col gap-8 my-auto items-center justify-center">
        <div class="p-3 bg-gray-700 rounded shadow-md shadow-black opacity-70 scale-90">
          <font-awesome-icon :icon="['fas', 'check']" />
          Benchmarking 1/20 tags
        </div>
        <hr class="h-[2px] w-11/12 mx-8 text-white bg-gray-500 border-0" />
        <div class="p-3 bg-gray-700 rounded shadow-md shadow-black">
          <font-awesome-icon class="animate-spin" :icon="['fas', 'spinner']" />
          Benchmarking {{ `${status.current}/${status.size}` }} tags
        </div>
        <hr class="h-[2px] w-11/12 mx-8 text-white bg-gray-500 border-0" />
        <div class="p-3 bg-gray-700 rounded shadow-md shadow-black opacity-70 scale-90">
          <font-awesome-icon :icon="['fas', 'check']" />
          Benchmarking 0/20 tags
        </div>
      </div>
    </div>
    <div v-else>
      <button class="p-5 bg-blue-500 rounded shadow-md shadow-black hover:scale-105 transition" @click="status.benchmarking = true">ベンチマーク開始</button>
    </div>
  </div>
</template>
