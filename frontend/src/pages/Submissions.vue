<script setup lang="ts">
import { onMounted, Ref, ref } from 'vue';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client';
import { IState, useStateStore } from '../stores/state';
import { Status, Submit } from 'proto-gen-web/src/backend/resources';
import { ListSubmitsRequest } from 'proto-gen-web/src/backend/messages';

const backend = new BackendServiceClient(
  new GrpcWebFetchTransport({
    baseUrl: import.meta.env.PROD ? `http://${window.location.hostname}:8080` : 'http://localhost:8080'
  })
)

const state:IState = useStateStore()

const submits:Ref<Submit[]> = ref([])

const formatDate = (timestamp: number):string => {
  const dateObject: Date = new Date(timestamp * 1000)
  const date: string = dateObject.toLocaleDateString()
  const time: string = dateObject.toLocaleTimeString()
  
  return `${date} ${time}`
}

onMounted(() => {
  const opt = {meta: {'authorization' : 'Bearer ' + state.token}}
  const listSubmitsRequest:ListSubmitsRequest = {
    // groupId: '2',
    // status: Status.SUCCESS
  }

  backend.listSubmits(listSubmitsRequest, opt)
      .then(res => {
        console.log(res.response.submits)
        submits.value = res.response.submits
      })
})
</script>
<template>
  <!-- TODO: show "no submissions when server returns no submissions" -->
  <table v-if="submits.length > 0" class="table-auto">
    <thead class="bg-gray-700">
      <tr>
        <th class="px-2 py-3">提出ID</th>
        <th class="">提出日時</th>
        <th class="px-2">グループID</th>
        <th class="">得点</th>
        <th class="">結果</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(s, idx) in submits" class="bg-gray-900 border-b-2 border-gray-800 hover:bg-gray-700 cursor-pointer transition" key="idx">
       <td class="w-20 text-center">{{s.id}}</td>
       <td class="w-60 text-center">{{formatDate(Number(s.submitedAt?.seconds))}}</td>
       <td class="w-30 text-center">{{s.groupId}}</td>
       <td class="w-20 text-center px-4">
         <div class="w-20 bg-gray-500 rounded text-center justify-center">
          {{s.score}}
          </div>
       </td>
       <td class="py-2 transition-colors text-center">
          <div v-if="s.status == Status.WAITING" class="p-1 w-40 bg-teal-500 rounded">Waiting</div>
          <div v-else-if="s.status == Status.IN_PROGRESS" class="p-1 w-40 bg-teal-500 rounded">In Progress</div>
          <div v-else-if="s.status == Status.SUCCESS" class="p-1 w-40 bg-blue-600 rounded">Success</div>
          <div v-else-if="s.status == Status.CONNECTION_FAILED" class="p-1 w-40 bg-red-600 rounded">Connection Failed</div>
          <div v-else-if="s.status == Status.VALIDATION_ERROR" class="p-1 w-40 bg-orange-500 rounded">Validation Error</div>
          <div v-else-if="s.status == 5" class="p-1 w-40 bg-orange-500 rounded">Internal Error</div>
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
