<script setup lang="ts">
import { onMounted, Ref, ref } from 'vue';
import { IState, useStateStore } from '../stores/state';
import { Status, Submit } from 'proto-gen-web/services/backend/resources';
import { GetSubmitRequest, ListSubmitsRequest } from 'proto-gen-web/services/backend/messages';
import Result from '../components/Result.vue'
import { useBackendStore } from '../stores/backend'

const state:IState = useStateStore()
const { backend } = useBackendStore()

const noSubmissions:Ref<boolean> = ref(false)

const modalItem:Ref<Partial<Submit>> = ref({})

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
    contestSlug: "test-contest", // TODO: fix
    // groupName: 'a01',
    // status: Status.VALIDATION_ERROR
  }

  backend.listSubmits(listSubmitsRequest, opt)
    .then(res => {
      if(import.meta.env.DEV) console.log('Submits', res.response.submits)
      state.submits = res.response.submits

      if(res.response.submits.length == 0) {
        noSubmissions.value = true
      }
    })
})
</script>
<template>
  <!-- TODO: show "no submissions" when server returns no submissions -->

  <!-- modal -->
  <transition enter-active-class="duration-100 ease-out" enter-from-class="transform opacity-0"
    enter-to-class="opacity-100" leave-active-class="duration-100 ease-in" leave-from-class="opacity-100"
    leave-to-class="transform opacity-0">


    <div v-if="Object.keys(modalItem).length > 0" @click.self="modalItem = {}"
      class="fixed flex justify-center items-center inset-0 bg-black bg-opacity-50 overflow-y-auto h-full w-full">
      <div class="w-5/6 h-5/6">
        <result :submit="modalItem" :title="'結果詳細'" :show-close-button="true" @close-modal="modalItem = {}" />
      </div>
    </div>

  </transition>
  <table v-if="state.submits.length > 0" class="table-auto">
    <thead class="bg-gray-700">
      <tr>
        <th class="px-2 py-3">提出ID</th>
        <th class="">提出日時</th>
        <th class="px-2">グループ名</th>
        <th class="">得点</th>
        <th class="w-48">結果</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(s, idx) in state.submits"
        class="bg-gray-900 border-b-2 border-gray-800 hover:bg-gray-700 cursor-pointer transition"
        @click.prevent="handleModal(s)" key="idx">
        <td class="w-20 text-center">{{ s.id }}</td>
        <td class="w-60 text-center">{{ formatDate(Number(s.submitedAt?.seconds)) }}</td>
        <td class="w-30 text-center">{{ s.groupName }}</td>
        <td class="w-20 text-center px-5">
          <div class="w-20 bg-gray-500 rounded text-center justify-center">
            {{ s.score }}
          </div>
        </td>
        <td class="py-2 transition-colors text-center">
          <div v-if="s.status == Status.WAITING" class="p-1 w-40 bg-teal-500 rounded mx-auto">Waiting</div>
          <div v-else-if="s.status == Status.IN_PROGRESS" class="p-1 w-40 bg-teal-500 rounded mx-auto">In Progress</div>
          <div v-else-if="s.status == Status.SUCCESS" class="p-1 w-40 bg-blue-600 rounded mx-auto">Success</div>
          <div v-else-if="s.status == Status.CONNECTION_FAILED" class="p-1 w-40 bg-red-600 rounded mx-auto">Connection
            Failed</div>
          <div v-else-if="s.status == Status.VALIDATION_ERROR" class="p-1 w-40 bg-orange-500 rounded mx-auto">Validation
            Error</div>
          <div v-else-if="s.status == Status.TIMEOUT" class="p-1 w-40 bg-orange-500 rounded mx-auto">Timeout</div>
          <div v-else-if="s.status == Status.INTERNAL_ERROR" class="p-1 w-40 bg-orange-500 rounded mx-auto">Internal Error
          </div>
          <div v-else class="p-1 w-40 bg-orange-500 rounded">Unknown Error</div>
        </td>
      </tr>
    </tbody>
  </table>

  <div class="mt-auto" v-else>
    <font-awesome-icon v-if="!noSubmissions" class="animate-spin text-3xl" :icon="['fas', 'spinner']"></font-awesome-icon>
    <div v-else>まだベンチマーク結果がありません。</div>
  </div>
  <div v-if="state.debug" class="flex flex-col gap-2 w-full">
    <div class="p-1 w-40 bg-teal-500 rounded">Waiting</div>
    <div class="p-1 w-40 bg-teal-500 rounded">In Progress</div>
    <div class="p-1 w-40 bg-blue-600 rounded">Success</div>
    <div class="p-1 w-40 bg-red-600 rounded">Connection Failed</div>
    <div class="p-1 w-40 bg-orange-500 rounded">Validation Error</div>
    <div class="p-1 w-40 bg-orange-500 rounded">Internal Error</div>
    {{ Status }}
  </div>
</template>
