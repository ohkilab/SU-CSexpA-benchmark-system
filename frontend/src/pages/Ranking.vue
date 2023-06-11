<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client';
import { computed, onMounted, reactive, Ref, ref } from 'vue';
import RankItem from '../components/RankItem.vue'
import TopRank from '../components/TopRank.vue'
import { GetRankingResponse_Record } from 'proto-gen-web/src/backend/messages';

import { useStateStore, IState } from '../stores/state';

const state = useStateStore()
const records: Ref<GetRankingResponse_Record[]> = ref(state.records ?? [])

const backend = new BackendServiceClient(
  new GrpcWebFetchTransport({
    baseUrl: import.meta.env.PROD ? `http://${window.location.hostname}:8080` : state.devBaseUrl
  })
)

onMounted(() => {
  let opt = { meta: { 'authorization': 'Bearer ' + state.token } }

  backend.getRanking({
    contestId: 1,
    containGuest: false
  }, opt).then(res => {
    records.value = res.response.records
    if (import.meta.env.DEV) console.log(res.response.records)
    // state.records = records.value
  })
})

const sortedRecords = computed(() =>
  records.value.sort((a: GetRankingResponse_Record, b: GetRankingResponse_Record) => a.rank.toString().localeCompare(b.rank.toString()))
)

</script>
<template>
  <!-- container -->
  <div v-if="records.length > 0" class="flex flex-col items-center gap-5 w-full px-4">
    <!-- separator -->
    <TopRank
      v-for="(g, idx) in records.sort((a: GetRankingResponse_Record, b: GetRankingResponse_Record) => a.rank.toString().localeCompare(b.rank.toString())).filter((_, i: number) => i < 3)"
      :key="g.group?.id" :rank="idx + 1" :class="state.group == g.group?.id ? 'bg-blue-700' : 'bg-gray-700'"
      :name="g.group?.id ?? ''" :score="g.score ?? 0" />
    <!-- top rank and normal rank separator -->
    <hr class="h-[2px] w-11/12 mx-8 text-white bg-gray-500 border-0" />
    <RankItem v-for="(g, idx) in sortedRecords.filter((_, i: number) => i >= 3)" :key="g.group?.id" :rank="idx + 4"
      :class="state.group == g.group?.id ? 'bg-blue-700' : 'bg-gray-700'" :name="g.group?.id ?? ''"
      :score="g.score ?? 0" />
  </div>
  <div class="mt-auto" v-else>
    <font-awesome-icon class="animate-spin text-3xl" :icon="['fas', 'spinner']"></font-awesome-icon>
  </div>
</template>
