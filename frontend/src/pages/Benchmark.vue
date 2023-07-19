<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { onMounted, ref, watch } from 'vue'
import { useStateStore, IState } from '../stores/state'

import type { Ref } from 'vue'

import { Status, Submit, TaskResult } from 'proto-gen-web/services/backend/resources'

import Result from '../components/Result.vue'
import { useBackendStore } from '../stores/backend'

const state:IState = useStateStore()
const { backend } = useBackendStore()

const latestSubmit:Ref<Partial<Submit>> = ref({})

const currentStatus:Ref<Status> = ref(0)

const errorMsg = ref('')

const taskResults: Ref<TaskResult[]> = ref([])

const url: Ref<string> = ref('')

const urlList: Ref<string[]> = ref([])

const noSubmissions: Ref<boolean> = ref(false)

const fetchLatestSubmit = () => {

  let opt = {meta: {'authorization' : 'Bearer ' + state.token}}

  backend.getLatestSubmit({}, opt).then(res => {
    console.log(res)
    latestSubmit.value = res.response.submit ?? {}
  }).catch(err => {
    console.log('getLatestSubmit err:' + err)
  })
}

const benchmark = () => {
  let opt = {meta: {'authorization' : 'Bearer ' + state.token}}
  taskResults.value = []

  backend.postSubmit({
    url: url.value, //'http://host.docker.internal:3001',
    contestSlug: state.contestSlug,
  },opt).then(async res => {
    if(import.meta.env.DEV) console.log("postSubmit res: ", res)
    state.benchmarking = true

    let call = backend.getSubmit({submitId: res.response.id}, opt)
    for await (let message of call.responses) {
      if(!state.benchmarking) break

      if(import.meta.env.DEV) console.log('Submit', message)

      currentStatus.value = message.submit?.status ?? Status.WAITING

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
    if(import.meta.env.DEV) console.log(err)
    if(err.code === "FAILED_PRECONDITION") {
      errorMsg.value = "Contest is over"
    } else {
      errorMsg.value = JSON.stringify(err) ?? ''
    }
  })
}

const handleStopBenchmark = () => {
  state.current = 0
  state.benchmarking = false
  state.showResult = true
  fetchLatestSubmit()
}

onMounted(() => {
  url.value = localStorage.getItem('currentUrl') ?? ''
  urlList.value = JSON.parse(localStorage.getItem('urlList') ?? '[]')

  fetchLatestSubmit()
  // fix BigInt problem
  // if(import.meta.env.DEV) BigInt.prototype.toJSON = function() {return this.toString()}
})

const statusMessage = (status: Status) => {
  switch (status) {
    case Status.WAITING: //1
      return 'Waiting'
    case Status.IN_PROGRESS: //2
      return 'In progress'
    case Status.SUCCESS: //3
      return 'Success'
    case Status.CONNECTION_FAILED: //4
      return 'Connection Failed'
    case Status.VALIDATION_ERROR: //5
      return 'Validation Error'
  }
}

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
      <pre>{{ taskResults.length }}</pre>
      <pre>{{ Status }}</pre>
      <pre>{{ statusMessage(currentStatus) }}</pre>
      <input class="text-red-500" v-model="currentStatus" type="number">
      <pre>{{ Status[currentStatus] }}</pre>
      <button class="border border-red-500 p-2" @click="latestSubmit = {}">clear submit</button>
    </fieldset>
    <div class="text-red-500 text-center" v-if="errorMsg">Error: {{ errorMsg }}</div>
    <div v-if="state.benchmarking">
      <div class="flex flex-col gap-8 my-auto text-xl items-center justify-center">
        <div class="p-4 border-2 rounded border-gray-600 flex flex-col items-center gap-3">
          <div>
            ステータス: {{ statusMessage(currentStatus) }}
          </div>
          <div>
            {{ currentStatus == Status.IN_PROGRESS ? `タグ ${state.current + 1}/${state.size} ベンチマーク中` :
              currentStatus == Status.SUCCESS ? `ベンチマーク成功` : ''
            }}
          </div>
          <font-awesome-icon v-if="currentStatus == Status.WAITING || currentStatus == Status.IN_PROGRESS"
            class="animate-spin" :icon="['fas', 'spinner']" />
        </div>

        <div v-if="currentStatus != Status.WAITING" class="flex flex-wrap gap-5 max-w-[1000px] justify-center">
          <div v-for="(t, i) in taskResults" :key="i"
            class="transition-all ease-out duration-200 w-20 p-3 text-center rounded shadow-md shadow-black" :class="t.status == Status.WAITING ? 'opacity-70' :
                t.status == Status.IN_PROGRESS ? 'bg-teal-500' :
                  t.status == Status.SUCCESS ? 'bg-blue-600' :
                    t.status == Status.CONNECTION_FAILED ? 'bg-red-500' :
                      t.status == Status.VALIDATION_ERROR ? 'bg-orange-500' :
                        t.status == Status.INTERNAL_ERROR ? 'bg-orange-500' : 'bg-gray-700 opacity-70'
              ">

            <font-awesome-icon v-if="state.current == i" class="animate-spin"
              :icon="['fas', 'spinner']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.WAITING" :icon="['fas', 'minus']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.IN_PROGRESS" :icon="['fas', 'spinner']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.SUCCESS" :icon="['fas', 'check']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.CONNECTION_FAILED" :icon="['fas', 'x']"></font-awesome-icon>
            <font-awesome-icon v-else-if="t.status == Status.VALIDATION_ERROR"
              :icon="['fas', 'exclamation']"></font-awesome-icon>
            <font-awesome-icon v-else :icon="['fas', 'minus']"></font-awesome-icon>
            {{ i + 1 }}
          </div>
        </div>
        <button @click="handleStopBenchmark"
          class="p-5 bg-red-500 rounded shadow-black shadow-md hover:scale-105 transition">ベンチマーク停止</button>
      </div>
    </div>
    <div class="flex flex-col gap-5 w-full items-center" v-else>
      <!-- deprecated: url list -->
      <div class="flex gap-4 w-5/6" v-for="(u, idx) in urlList" :key="u">
        <button @click="url = u"
          class="w-full rounded bg-gray-600 p-2 text-left hover:bg-gray-600 transition">{{ u }}</button>
        <button @click="urlList.splice(idx, 1)"
          class="px-4 bg-red-500 rounded shadow-black shadow-md transition hover:scale-105">-</button>
      </div>
      <div class="w-5/6 flex gap-5">
        <input class="w-full rounded bg-gray-700 p-2 hover:bg-gray-600 transition focus:outline-none focus:bg-gray-600"
          placeholder="Raspberry Pi の IPアドレス。例：http://192.168.1.10:3000" type="text" v-model="url">
        <!-- <button @click="urlList.push(url)" class="px-4 bg-blue-500 rounded shadow-black shadow-md transition hover:scale-105">+</button> -->
      </div>
      <button class="p-5 mb-5 bg-blue-500 w-64 rounded text-xl shadow-md shadow-black hover:scale-105 transition"
        @click="benchmark">ベンチマーク開始</button>
      <div v-if="Object.keys(latestSubmit).length > 0" class="w-5/6 h-[500px]">
        <result :submit="latestSubmit" :title="`最新結果`" />
      </div>
      <div v-else class="w-5/6 bg-gray-700 rounded-md flex justify-center items-center h-[500px]">
        {{ noSubmissions ? 'まだベンチマーク結果がありません。' : '読み込み中...' }}
      </div>
    </div>
  </div>
</template>
