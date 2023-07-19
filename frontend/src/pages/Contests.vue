<script setup lang="ts">
import { onMounted, Ref, ref } from 'vue'
import { useBackendStore } from '../stores/backend';
import { useStateStore } from '../stores/state';
import { Contest } from 'proto-gen-web/services/backend/resources';
import { useRouter } from 'vue-router';

const { backend } = useBackendStore()
const state = useStateStore()

const router = useRouter()

const contests: Ref<Contest[]> = ref([])

const formatDate = (timestamp: number): string => {
  const dateObject: Date = new Date(timestamp * 1000)
  const date: string = dateObject.toLocaleDateString()

  const time: string = dateObject.toLocaleTimeString()

  return `${date} ${time}`
}

const handleSelectContest = (contest: Contest) => {
  state.contestSlug = contest.slug
  state.selectedContestName = contest.title

  router.push('/benchmark')
}

onMounted(() => {
  let opt = { meta: { 'authorization': 'Bearer ' + state.token } }

  backend.listContests({}, opt).then((res) => {
    if(import.meta.env.DEV) console.log(res)
    contests.value = res.response.contests ?? []
  })
})

</script>
<template>
  <div class="">
    <h1 class="text-center py-3 text-xl">コンテスト一覧</h1>
    <table v-if="contests.length > 0" class="table-auto">
      <thead class="bg-gray-700">
        <tr>
          <th class="w-36 px-2 py-3">コンテストID</th>
          <th class="w-40">コンテスト名</th>
          <th class="w-52">開始時間</th>
          <th class="w-30">終了時間</th>
          <th class="w-32">提出回数</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="c in contests"
            @click.prevent="handleSelectContest(c)"
          class="bg-gray-900 border-b-2 border-gray-800 hover:bg-gray-700 cursor-pointer transition"
           key="c.id"
           >
          <td class="text-center py-3">{{c.id}}</td>
          <td class="text-center">{{c.title}}</td>
          <td class="text-center">{{formatDate(Math.floor(Number(c.startAt?.seconds)))}}</td>
          <td class="text-center">{{formatDate(Number(c.endAt?.seconds))}}</td>
          <td class="text-center px-5">{{c.submitLimit}}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
