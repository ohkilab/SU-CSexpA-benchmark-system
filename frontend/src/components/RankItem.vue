<script setup lang="ts">
defineProps<{ rank: number; score: number; name: string }>()
</script>
<template>
  <div class="cursor-pointer bg-gray-700 rounded shadow-gray-950 shadow-md flex gap-4 px-2 py-4 w-full sm:w-[400px] justify-around hover:scale-105 transition">
  <div class="w-8 text-center justify-center">
    {{ rank }}
  </div>
  <div>{{ name }}</div>
  <div class="flex gap-2 items-center">
    <div class="w-20 bg-gray-500 rounded text-center justify-center">
      {{ score }}
    </div>
    <div class="text-sm">req/s</div>
  </div>
  </div>
</template>
