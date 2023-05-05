<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { computed, onMounted, reactive, ref } from 'vue'
import { useStateStore, IState } from '../stores/state'

import type { Ref } from 'vue'

import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client'

import type { GetRankingRequest, GetSubmitRequest } from 'proto-gen-web/src/backend/messages'
import type { Group } from 'proto-gen-web/src/backend/resources'
import { Role } from 'proto-gen-web/src/backend/resources'

const state:IState = useStateStore()

const backend = new BackendServiceClient(
  new GrpcWebFetchTransport({
    baseUrl: "http://localhost:8080"
  })
)

interface Status {
  benchmarking: boolean,
  showResult: boolean,
  result: number,
  current: number,
  size: number,
}

const status: Status = reactive({
  benchmarking: false,
  showResult: false,
  result: 0,
  current: 0,
  size: 20
})

const tags: Ref<Array<{tag: string, idx: number}>> = ref([])

const benchmark = () => {
  state.benchmarking = true

  state.benchmarkInterval = setInterval(() => {
    if(status.current < status.size) {
      status.current++
    } else {
      status.current = 0
      state.benchmarking = false
      status.showResult = true
      state.lastResult = Math.floor(Math.random() * (20000 - 200) + 200)

      clearInterval(state.benchmarkInterval)
      state.benchmarkInterval = 0
    }
  }, 100)
}

const handleStopBenchmark = () => {
  state.benchmarking = false
  status.current = 0
  state.benchmarking = false
  status.showResult = true
  clearInterval(state.benchmarkInterval)
}

onMounted(() => {
  tags.value = Array.from(new Array(status.size)).map((_, idx) => {

  return {
    tag: 'tag_string ' + idx,
    idx
  }
})

  let opt = {meta: {'authorization' : 'Bearer ' + state.token}}

  backend.getSubmit({
    submitId: 'test'
  },opt).then(res => {
    console.log(res)
  })
})

const filteredTags = computed(() => tags.value.slice(status.current - 2, status.current + 2))
</script>
<template>
  <div  class="flex flex-col mt-auto">
    <div v-if="!state.benchmarking && state.lastResult != 0" class="border flex flex-col border-gray-500 text-center p-5 rounded mb-5">
      <div class="mb-2">最新結果</div>
      <div class="flex self-center">
        <div class="rounded bg-gray-500 px-2">{{state.lastResult}}</div>
        &nbsp;req/s
      </div>
    </div>
    <div v-if="state.benchmarking">
      <div class="flex flex-col gap-8 my-auto text-xl items-center justify-center">
        <div class="p-4 border-2 rounded border-gray-600 flex flex-col items-center gap-3">
          タグ {{ `${status.current+1}/${status.size}` }} ベンチマーク中
          <br>
          <font-awesome-icon class="animate-spin" :icon="['fas', 'spinner']" />
        </div>

        <div class="flex flex-wrap gap-5 self-center max-w-[600px]">
          <div
            v-for="(t, i) in tags"
            :key="i"
              class="transition-all ease-out duration-200 w-32 p-3 text-center bg-gray-700 rounded shadow-md shadow-black"
              :class="status.current > i ? 'bg-green-700' :
                      status.current == i ? 'bg-gray-700' :
                      'opacity-70'
                     "
          >
            <font-awesome-icon v-if="status.current > i" :icon="['fas', 'check']"></font-awesome-icon>
            <font-awesome-icon v-else-if="status.current == i" class="animate-spin" :icon="['fas', 'spinner']"></font-awesome-icon>
            <font-awesome-icon v-else :icon="['fas', 'minus']"></font-awesome-icon>
            {{ i+1 }}
          </div>
        </div>
        <button @click="handleStopBenchmark" class="p-5 bg-red-500 rounded shadow-black shadow-md hover:scale-105 transition">ベンチマーク停止</button>
      </div>
    </div>
    <div v-else>
      <button class="p-5 bg-blue-500 rounded text-xl shadow-md shadow-black hover:scale-105 transition" @click="benchmark">ベンチマーク開始</button>
    </div>
  </div>
</template>
