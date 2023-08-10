<script setup lang="ts">
defineProps<{ rank: number; score: number; name: string }>();
</script>
<template>
  <div
    class="flex w-full cursor-pointer justify-around gap-4 rounded px-2 py-4 shadow-md shadow-gray-950 transition hover:scale-105 sm:w-[400px]"
  >
    <div class="w-8 justify-center text-center">
      {{ rank }}
    </div>
    <div>{{ name }}</div>
    <div class="flex items-center gap-2">
      <div class="w-20 justify-center rounded bg-gray-500 text-center">
        {{ score }}
      </div>
      <div class="text-sm">req/s</div>
    </div>
  </div>
</template>
