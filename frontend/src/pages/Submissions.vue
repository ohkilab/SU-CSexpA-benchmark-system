<script setup lang="ts">
import { onMounted, Ref, ref } from 'vue';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client';
import { IState, useStateStore } from '../stores/state';
import { Status, Submit, TaskResult } from 'proto-gen-web/src/backend/resources';
import { GetSubmitRequest, ListSubmitsRequest } from 'proto-gen-web/src/backend/messages';

const backend = new BackendServiceClient(
  new GrpcWebFetchTransport({
    baseUrl: import.meta.env.PROD ? `http://${window.location.hostname}:8080` : 'http://localhost:8080'
  })
)

const state:IState = useStateStore()

const submits:Ref<Submit[]> = ref([])

const modalItem:Ref<Submit> = ref({
  id: 0,
  groupId: 0,
  year: 0,
  score: 0,
  language: 0,
  taskResults: [],
  tagCount: 0,
  status: 0,
  errorMessage: ''
})

const taskResults:Ref<TaskResult[]> = ref([])

const modalSubmitId:Ref<number> = ref(0)

const formatDate = (timestamp: number):string => {
  const dateObject: Date = new Date(timestamp * 1000)
  const date: string = dateObject.toLocaleDateString()
  const time: string = dateObject.toLocaleTimeString()
  
  return `${date} ${time}`
}

const handleModal = async (submit: Submit) => {
    // get taskResults
    const getSubmitRequest:GetSubmitRequest = {
      submitId: submit.id
    }

    const opt = {meta: {'authorization' : 'Bearer ' + state.token}}
    const call = backend.getSubmit(getSubmitRequest, opt)
    for await (let message of call.responses) {
      if(import.meta.env.DEV) console.log('Submit', message)
      submit.taskResults = message.submit?.taskResults ?? []
    }

    // set modal item
    modalItem.value = submit
}

onMounted(() => {
  const opt = {meta: {'authorization' : 'Bearer ' + state.token}}
  // TODO: get own submissions, filter functionality
  const listSubmitsRequest:ListSubmitsRequest = {
    // groupId: '2',
    // status: Status.VALIDATION_ERROR
  }

  backend.listSubmits(listSubmitsRequest, opt)
    .then(res => {
      if(import.meta.env.DEV) console.log('Submits', res.response.submits)
      submits.value = res.response.submits
    })

})
</script>
<template>
  <!-- TODO: show "no submissions" when server returns no submissions -->

  <!-- modal -->
  <div v-if="modalItem.taskResults.length > 0" @click="modalItem.taskResults = []" class="fixed flex justify-center items-center inset-0 bg-black bg-opacity-50 overflow-y-auto h-full w-full">
        <div class="bg-gray-700 w-5/6 h-5/6 rounded overflow-y-auto mx-auto p-10 gap-4 flex flex-col">
          <div class="text-2xl">提出ID: {{modalItem.id}} 結果詳細</div>
          <div class="text-md text-gray-300">提出日時: {{formatDate(Number(modalItem.submitedAt?.seconds))}}</div>
          <div class="text-md text-gray-300">グループID: {{modalItem.groupId}}</div>
          <div class="w-full h-full bg-gray-900 rounded p-8 overflow-y-auto flex flex-col gap-2">
          <div
            v-for="(t, i) in modalItem.taskResults"
            :key="i"
            class="flex gap-2 py-3 px-6 rounded shadow-md shadow-black items-center justify-between"
            :class="
              t.status == Status.WAITING ? 'opacity-70' :
              t.status == Status.IN_PROGRESS ? 'bg-teal-500' :
              t.status == Status.SUCCESS ? 'bg-blue-600' :
              t.status == Status.CONNECTION_FAILED ? 'bg-red-500' :
              t.status == Status.VALIDATION_ERROR ? 'bg-orange-500' :
              t.status == Status.INTERNAL_ERROR ? 'bg-orange-500' : 'bg-gray-700 opacity-70'
            "
          >
            <div class="flex justify-center items-center gap-2">
              タグ {{ i+1 }}：
              <div class="rounded bg-gray-500 px-2">{{t.requestPerSec}}</div>
              req/s
            </div>
            <div class='flex gap-5 items-center'>{{t.errorMessage != '' ? `エラー: ${t.errorMessage}` : ''}}
              <font-awesome-icon v-if="t.status == Status.IN_PROGRESS" :icon="['fas', 'spinner']"></font-awesome-icon>
              <font-awesome-icon v-else-if="t.status == Status.WAITING" :icon="['fas', 'minus']"></font-awesome-icon>
              <font-awesome-icon v-else-if="t.status == Status.SUCCESS" :icon="['fas', 'check']"></font-awesome-icon>
              <font-awesome-icon v-else-if="t.status == Status.CONNECTION_FAILED" :icon="['fas', 'x']"></font-awesome-icon>
              <font-awesome-icon v-else-if="t.status == Status.VALIDATION_ERROR" :icon="['fas', 'exclamation']"></font-awesome-icon>
              <font-awesome-icon v-else :icon="['fas', 'minus']"></font-awesome-icon>
            </div>
          </div>
          </div>
        </div>
      </div>
  <table v-if="submits.length > 0" class="table-auto">
    <thead class="bg-gray-700">
      <tr>
        <th class="px-2 py-3">提出ID</th>
        <th class="">提出日時</th>
        <th class="px-2">グループID</th>
        <th class="">得点</th>
        <th class="w-48">結果</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="(s, idx) in submits" class="bg-gray-900 border-b-2 border-gray-800 hover:bg-gray-700 cursor-pointer transition" 
        @click.prevent="handleModal(s)"
        key="idx"
      >
       <td class="w-20 text-center">{{s.id}}</td>
       <td class="w-60 text-center">{{formatDate(Number(s.submitedAt?.seconds))}}</td>
       <td class="w-30 text-center">{{s.groupId}}</td>
       <td class="w-20 text-center px-5">
         <div class="w-20 bg-gray-500 rounded text-center justify-center">
          {{s.score}}
          </div>
       </td>
       <td class="py-2 transition-colors text-center">
          <div v-if="s.status == Status.WAITING" class="p-1 w-40 bg-teal-500 rounded mx-auto">Waiting</div>
          <div v-else-if="s.status == Status.IN_PROGRESS" class="p-1 w-40 bg-teal-500 rounded mx-auto">In Progress</div>
          <div v-else-if="s.status == Status.SUCCESS" class="p-1 w-40 bg-blue-600 rounded mx-auto">Success</div>
          <div v-else-if="s.status == Status.CONNECTION_FAILED" class="p-1 w-40 bg-red-600 rounded mx-auto">Connection Failed</div>
          <div v-else-if="s.status == Status.VALIDATION_ERROR" class="p-1 w-40 bg-orange-500 rounded mx-auto">Validation Error</div>
          <div v-else-if="s.status == Status.INTERNAL_ERROR" class="p-1 w-40 bg-orange-500 rounded mx-auto">Internal Error</div>
          <div v-else class="p-1 w-40 bg-orange-500 rounded">Unknown Error</div>
       </td>
      </tr>
    </tbody>
  </table>

  <div class="mt-auto" v-else>
    <font-awesome-icon  class="animate-spin text-3xl" :icon="['fas', 'spinner']"></font-awesome-icon>
  </div>
  <div v-if="state.debug" class="flex flex-col gap-2 w-full">
    <div class="p-1 w-40 bg-teal-500 rounded">Waiting</div>
    <div class="p-1 w-40 bg-teal-500 rounded">In Progress</div>
    <div class="p-1 w-40 bg-blue-600 rounded">Success</div>
    <div class="p-1 w-40 bg-red-600 rounded">Connection Failed</div>
    <div class="p-1 w-40 bg-orange-500 rounded">Validation Error</div>
    <div class="p-1 w-40 bg-orange-500 rounded">Internal Error</div>
    {{Status}}
  </div>
</template>
