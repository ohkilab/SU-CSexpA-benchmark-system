<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
defineProps<{ rank: number; score: number; name: string }>()
</script>
<template>
  <div
    class="cursor-pointer text-xl w-full sm:w-[500px] rounded justify-around items-center shadow-md shadow-gray-950 flex p-5 gap-2 hover:scale-105 transition"
  >
    <div class="w-10 text-center justify-center">
      <div class="text-2xl">
        <font-awesome-icon v-if="rank == 1" :icon="['fas', 'crown']" />
        {{ rank != 1 ? rank : '' }}
      </div>
    </div>
    <div>{{ name }}</div>
    <div class="flex gap-2 items-center">
      <div class="w-24 bg-gray-500 rounded text-center justify-center">
        {{ score }}
      </div>
      <div class="text-sm">req/s</div>
    </div>
  </div>
</template>
