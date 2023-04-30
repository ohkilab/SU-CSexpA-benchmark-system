<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { computed, onMounted, reactive, ref } from 'vue'
import type { Ref } from 'vue'

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

const tags: Ref<Array<{tag: string, idx: number}>> = ref([])

onMounted(() => {
  tags.value = Array.from(new Array(status.size)).map((_, idx) => {
  return {
    tag: 'tag_string ' + idx,
    idx
  }
})

setInterval(() => {
  if(status.current < status.size) {
    status.current++
  } else {
    status.current = 0
  }
}, 500)
  // setInterval(() => {
  //   status.current++
  //   if (status.current >= status.size) status.current = 0
  // }, 1000)
})

const filteredTags = computed(() => tags.value.slice(status.current - 2, status.current + 2))
</script>
<template>
  <div class="flex flex-col mt-auto">
    <div v-if="status.benchmarking">
      <div class="flex flex-col gap-8 my-auto text-xl items-center justify-center">
        <div class="p-4 border-2 rounded border-gray-600 flex flex-col items-center gap-3">
          タグ {{ `${status.current+1}/${status.size}` }} ベンチマーク中
          <br>
          <font-awesome-icon class="animate-spin" :icon="['fas', 'spinner']" />
        </div>

        <div
          v-for="(t, i) in tags"
          :key="i"
        >
          <div
            class="transition-all ease-out duration-200 p-3 bg-gray-700 rounded shadow-md shadow-black"
            :class="status.current > i ? 'bg-green-700' :
                    status.current == i ? 'bg-gray-700 p-5' :
                    'opacity-70'
                   "
          >
          <font-awesome-icon v-if="status.current > i" :icon="['fas', 'check']"></font-awesome-icon>
            <font-awesome-icon v-else-if="status.current == i" class="animate-spin" :icon="['fas', 'spinner']"></font-awesome-icon>
            <font-awesome-icon v-else :icon="['fas', 'minus']"></font-awesome-icon>
            {{ i+1 }}: {{ t.tag }}
          </div>
        </div>
      </div>
    </div>
    <div v-else>
      <button class="p-5 bg-blue-500 rounded text-xl shadow-md shadow-black hover:scale-105 transition" @click="status.benchmarking = true">ベンチマーク開始</button>
    </div>
  </div>
</template>
