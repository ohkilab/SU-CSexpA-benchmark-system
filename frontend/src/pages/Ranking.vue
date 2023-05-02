<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client';
import { onMounted, Ref, ref } from 'vue';
import RankItem from '../components/RankItem.vue'
import TopRank from '../components/TopRank.vue'
import { GetRankingResponse_Record } from 'proto-gen-web/src/backend/messages';

import { useStateStore } from '../stores/state';

const state = useStateStore()

const backend = new BackendServiceClient(
  new GrpcWebFetchTransport({
    baseUrl: "http://localhost:8080"
  })
)

onMounted(() => {
  let opt = {meta: {'authorization' : 'Bearer ' + state.token}}

  backend.getRanking({
    year: 2023,
    containGuest: false
  },opt).then(res => {
    state.records = res.response.records
  })
})
</script>
<template>
    <!-- container -->
    <div v-if="state.records.length > 0" class="flex flex-col items-center gap-5 w-full px-4">
      <!-- separator -->
      <TopRank
        v-for="(g, idx) in state.records.sort((a, b) => a.group.score < b.group.score ? 1 : 0).filter((_, i:number) => i < 3)"
        :key="g.group.id"
        :rank="idx + 1"
        :class="state.id == g.group.id ? 'bg-blue-700' : 'bg-gray-700'"
        :name="g.group.id"
        :score="g.group.score"
      />
      <!-- top rank and normal rank separator -->
      <hr class="h-[2px] w-11/12 mx-8 text-white bg-gray-500 border-0" />
      <RankItem
        v-for="(g, idx) in state.records.sort((a, b) => a.group.score < b.group.score ? 1 : 0).filter((_, i:number) => i >= 3)"
        :key="g.group.id"
        :rank="idx + 4"
        :class="state.id == g.group.id ? 'bg-blue-700' : 'bg-gray-700'"
        :name="g.group.id"
        :score="g.group.score"
      />
    </div>
    <div class="mt-auto" v-else>
      <font-awesome-icon  class="animate-spin text-3xl" :icon="['fas', 'spinner']"></font-awesome-icon>
    </div>
</template>
