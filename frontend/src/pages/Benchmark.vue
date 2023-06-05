<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { useStateStore, IState } from '../stores/state'

import type { Ref } from 'vue'

import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client'

import { Group, Status, TaskResult } from 'proto-gen-web/src/backend/resources'

const state:IState = useStateStore()

const webfetchTransport = new GrpcWebFetchTransport({baseUrl: 'http://localhost:8080'})

const backend = new BackendServiceClient(webfetchTransport)

const errorMsg = ref('')

const taskResults: Ref<TaskResult[]> = ref([])

const url: Ref<string> = ref('')

const urlList: Ref<string[]> = ref([])

const benchmark = () => {

  let opt = {meta: {'authorization' : 'Bearer ' + state.token}}
  taskResults.value = []

  backend.postSubmit({
    url: url.value, //'http://host.docker.internal:3001',
    contestId: 1
  },opt).then(async res => {
    if(import.meta.env.DEV) console.log(res)
    state.benchmarking = true

    let call = backend.getSubmit({submitId: res.response.id}, opt)
    for await (let message of call.responses) {
      if(!state.benchmarking) break

      if(import.meta.env.DEV) console.log('Submit', message)

      taskResults.value = Array.from(Array(message.submit?.tagCount)).map((_, i) => message.submit?.taskResults[i] ?? {} as TaskResult)
      errorMsg.value = message.submit?.errorMessage ?? ''
      state.lastResult = message.submit?.score ?? 0
      state.current = message.submit?.taskResults.length ?? -1
      state.size = message.submit?.tagCount ?? 0
    }

    handleStopBenchmark()

    if(import.meta.env.DEV) {
      console.log(call.status)
      console.log(call.trailers)
    }
  }).catch(err => {
    console.log(err)
  })

}

const handleStopBenchmark = () => {
  state.current = 0
  state.benchmarking = false
  state.showResult = true
}

onMounted(() => {
  url.value = localStorage.getItem('currentUrl') ?? ''
  urlList.value = JSON.parse(localStorage.getItem('urlList') ?? '[]')

  // fix BigInt problem
  // if(import.meta.env.DEV) BigInt.prototype.toJSON = function() {return this.toString()}
})

watch(url, url => {
  localStorage.setItem('currentUrl', url)
})

watch(urlList, urlList => {
  localStorage.setItem('urlList', JSON.stringify(urlList))
}, {deep: true})

</script>
<template>
  <div class="flex flex-col mt-auto w-full items-center">
    <fieldset v-if="state.debug" class="border border-red-500 p-2">
      <legend>Debug</legend>
      <pre>{{taskResults.length}}</pre>

      <div v-if="state.debug" class="flex flex-col gap-2 w-full">
        <div class="p-1 w-40 bg-teal-500 rounded">Waiting</div>
        <div class="p-1 w-40 bg-teal-500 rounded">In Progress</div>
        <div class="p-1 w-40 bg-blue-600 rounded">Success</div>
        <div class="p-1 w-40 bg-red-600 rounded">Connection Failed</div>
        <div class="p-1 w-40 bg-orange-500 rounded">Validation Error</div>
        <div class="p-1 w-40 bg-orange-500 rounded">Internal Error</div>
        {{Status}}
        <div class="p-1 w-40 bg-teal-500 rounded">teal-500</div>
        <div class="p-1 w-40 bg-teal-500 rounded">teal-500</div>
        <div class="p-1 w-40 bg-blue-600 rounded">blue-600</div>
        <div class="p-1 w-40 bg-red-600 rounded">red-600</div>
        <div class="p-1 w-40 bg-orange-500 rounded">orange-500</div>
        <div class="p-1 w-40 bg-orange-500 rounded">orange-500</div>
      </div>
    </fieldset>
    <div v-if="!state.benchmarking" class="border flex flex-col border-gray-500 p-5 text-center rounded mb-5">
      <div class="mb-2">最新結果</div>
      <div class="flex self-center mb-5">
        <div class="rounded bg-gray-500 px-2">{{state.lastResult}}</div>
        &nbsp;req/s
      </div>
      <div class="flex flex-wrap gap-5 max-w-[1000px] items-center justify-center">
        <div
          v-for="(t, i) in taskResults"
          :key="i"
          class="flex gap-1 w-40 p-3 bg-gray-700 justify-center items-center rounded shadow-md shadow-black"
          :class="
            t.status == Status.SUCCESS ? 'bg-blue-600' :
            t.status == Status.WAITING ? 'opacity-70' :
            t.status == Status.IN_PROGRESS ? 'bg-teal-500' :
            t.status == Status.CONNECTION_FAILED ? 'bg-red-500' :
            t.status == Status.VALIDATION_ERROR ? 'bg-orange-500' :
            t.status == Status.INTERNAL_ERROR ? 'bg-orange-500' : ''
          "
        >
            <font-awesome-icon v-if="t.status == Status.IN_PROGRESS" :icon="['fas', 'spinner']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.WAITING" :icon="['fas', 'minus']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.SUCCESS" :icon="['fas', 'check']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.CONNECTION_FAILED" :icon="['fas', 'x']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.VALIDATION_ERROR" :icon="['fas', 'exclamation']"></font-awesome-icon>
            <font-awesome-icon v-else :icon="['fas', 'minus']"></font-awesome-icon>
            {{ i+1 }}:
            <div class="rounded bg-gray-500 px-2">{{t.requestPerSec}}</div>
            req/s
        </div>
      </div>
    </div>
    <div class="text-red-500" v-if="errorMsg">Error: {{errorMsg}}</div>
    <div v-if="state.benchmarking">
      <div class="flex flex-col gap-8 my-auto text-xl items-center justify-center">
        <div class="p-4 border-2 rounded border-gray-600 flex flex-col items-center gap-3">
          タグ {{ `${state.current+1}/${state.size}` }} ベンチマーク中
          <br>
          <font-awesome-icon class="animate-spin" :icon="['fas', 'spinner']" />
        </div>

        <div class="flex flex-wrap gap-5 max-w-[1000px] justify-center">
          <div
            v-for="(t, i) in taskResults"
            :key="i"
              class="transition-all ease-out duration-200 w-20 p-3 text-center bg-gray-700 rounded shadow-md shadow-black"
              :class="
                t.status == Status.WAITING ? 'opacity-70' :
                t.status == Status.IN_PROGRESS ? 'bg-teal-500' :
                t.status == Status.SUCCESS ? 'bg-blue-600' :
                t.status == Status.CONNECTION_FAILED ? 'bg-red-500' :
                t.status == Status.VALIDATION_ERROR ? 'bg-orange-500' :
                t.status == Status.INTERNAL_ERROR ? 'bg-orange-500' : 'opacity-70'
              "
          >

            <font-awesome-icon v-if="state.current == i" class="animate-spin" :icon="['fas', 'spinner']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.WAITING" :icon="['fas', 'minus']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.IN_PROGRESS" :icon="['fas', 'spinner']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.SUCCESS" :icon="['fas', 'check']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.CONNECTION_FAILED" :icon="['fas', 'x']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.VALIDATION_ERROR" :icon="['fas', 'exclamation']"></font-awesome-icon>
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
        <!-- <button @click="urlList.push(url)" class="px-4 bg-blue-500 rounded shadow-black shadow-md transition hover:scale-105">+</button> -->
      </div>
      <button class="p-5 bg-blue-500 w-64 rounded text-xl shadow-md shadow-black hover:scale-105 transition" @click="benchmark">ベンチマーク開始</button>
    </div>
  </div>
</template>
