<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
defineProps<{ rank: number; score: number; name: string }>();
</script>
<template>
  <div
    class="flex w-full cursor-pointer items-center justify-around gap-2 rounded p-5 text-xl shadow-md shadow-gray-950 transition hover:scale-105 sm:w-[500px]"
  >
    <div class="w-10 justify-center text-center">
      <div class="text-2xl">
        <font-awesome-icon v-if="rank == 1" :icon="['fas', 'crown']" />
        {{ rank != 1 ? rank : "" }}
      </div>
    </div>
    <div>{{ name }}</div>
    <div class="flex items-center gap-2">
      <div class="w-24 justify-center rounded bg-gray-500 text-center">
        {{ score }}
      </div>
      <div class="text-sm">req/s</div>
    </div>
  </div>
</template>
