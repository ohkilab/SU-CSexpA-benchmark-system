<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { useStateStore, IState } from '../stores/state'

import type { Ref } from 'vue'

import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client'
import { HealthcheckServiceClient } from 'proto-gen-web/src/backend/services.client'

import type { GetRankingRequest, GetSubmitRequest, GetSubmitResponse } from 'proto-gen-web/src/backend/messages'
import type { Group } from 'proto-gen-web/src/backend/resources'
import { Role } from 'proto-gen-web/src/backend/resources'

const state:IState = useStateStore()

const webfetchTransport = new GrpcWebFetchTransport({baseUrl: 'http://localhost:8080'})

const backend = new BackendServiceClient(webfetchTransport)

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

const submits: Ref<GetSubmitResponse> = ref({})

const url: Ref<string> = ref('')

const urlList: Ref<string[]> = ref([])

const benchmark = () => {

  let opt = {meta: {'authorization' : 'Bearer ' + state.token}}

  backend.postSubmit({
    url: url.value, //'http://host.docker.internal:3001',
    contestId: 1
  },opt).then(async res => {
    console.log(res)
    state.benchmarking = true

    let call = backend.getSubmit({submitId: res.response.id}, opt)
    for await (let message of call.responses) {
      if(!state.benchmarking) break

      console.log("got a message", message)
      // status.current++
      status.current = message.submit?.taskResults.length ?? -1
    }

    status.current = 0
    state.benchmarking = false
    status.showResult = true
    console.log(call.status)
    console.log(call.trailers)

  }).catch(err => {
    console.log(err)
  })

  // state.benchmarkInterval = setInterval(() => {
  //   if(status.current < status.size) {
  //     status.current++
  //   } else {
  //     status.current = 0
  //     state.benchmarking = false
  //     status.showResult = true
  //     state.lastResult = Math.floor(Math.random() * (20000 - 200) + 200)
  //
  //     clearInterval(state.benchmarkInterval)
  //     state.benchmarkInterval = 0
  //   }
  // }, 100)
}

const handleStopBenchmark = () => {
  state.benchmarking = false
  status.current = 0
  state.benchmarking = false
  status.showResult = true
  clearInterval(state.benchmarkInterval)
}

onMounted(() => {
  let opt = {meta: {'authorization' : 'Bearer ' + state.token}}

  url.value = localStorage.getItem('currentUrl') ?? ''
  urlList.value = JSON.parse(localStorage.getItem('urlList') ?? '[]')
  
  tags.value = Array.from(new Array(status.size)).map((_, idx) => {
    return {
      tag: 'tag_string ' + idx,
      idx
    }
  })

  // backend.getSubmit({
  //   submitId: 1
  // }, opt).then(res => {
  //   // console.log(res)
  // })
})

watch(url, url => {
  localStorage.setItem('currentUrl', url)
})

watch(urlList, urlList => {
  localStorage.setItem('urlList', JSON.stringify(urlList))
}, {deep: true})

const filteredTags = computed(() => tags.value.slice(status.current - 2, status.current + 2))
</script>
<template>
  <div class="flex flex-col mt-auto w-full items-center">
    <div v-if="!state.benchmarking && state.lastResult != 0" class="border flex flex-col border-gray-500 text-center p-5 rounded mb-5 w-max">
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
    <div class="flex flex-col gap-5 w-full items-center" v-else>
      <div class="flex gap-4 w-5/6" v-for="(u, idx) in urlList" :key="u">
        <button @click="url = u" class="w-full rounded bg-gray-600 p-2 text-left hover:bg-gray-600 transition">{{u}}</button>
        <button @click="urlList.splice(idx, 1)" class="px-4 bg-red-500 rounded shadow-black shadow-md transition hover:scale-105">-</button>
      </div>
      <div class="w-5/6 flex gap-5">
        <input class="w-full rounded bg-gray-700 p-2 hover:bg-gray-600 transition focus:outline-none focus:bg-gray-600" placeholder="Raspberry Pi の IPアドレス" type="text" v-model="url">
        <button @click="urlList.push(url)" class="px-4 bg-blue-500 rounded shadow-black shadow-md transition hover:scale-105">+</button>
      </div>
      <button class="p-5 bg-blue-500 w-64 rounded text-xl shadow-md shadow-black hover:scale-105 transition" @click="benchmark">ベンチマーク開始</button>
    </div>
  </div>
</template>
