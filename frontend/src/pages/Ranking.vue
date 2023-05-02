<script setup lang="ts">
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { Group } from 'proto-gen-web/src/backend/resources';
import { BackendServiceClient } from 'proto-gen-web/src/backend/services.client';
import { onMounted, Ref, ref } from 'vue';
import RankItem from '../components/RankItem.vue'
import TopRank from '../components/TopRank.vue'

import { useCredStore } from '../stores/cred';

const cred = useCredStore()

const records: Ref<{ rank: number; score: number, group: Group }>[] = ref([
  {
    rank: 1,
    score: 1000,
    group: {
      id: 'a01',
      year: 2023,
      role: 0
    }
  },
  {
    rank: 1,
    score: 1000,
    group: {
      id: 'a03',
      year: 2023,
      role: 0
    }
  },
  {
    rank: 1,
    score: 1050,
    group: {
      id: 'a02',
      year: 2023,
      role: 0
    }
  },
  {
    rank: 1,
    score: 2230,
    group: {
      id: 'b01',
      year: 2023,
      role: 0
    }
  },
  {
    rank: 1,
    score: 10000,
    group: {
      id: 'b02',
      year: 2023,
      role: 0
    }
  },
])

const backend = new BackendServiceClient(
  new GrpcWebFetchTransport({
    baseUrl: "http://localhost:8080"
  })
)

onMounted(() => {
  let opt = {meta: {'authorization' : 'Bearer ' + cred.token}}

  backend.getRanking({
    year: 2023,
    containGuest: false
  },opt).then(res => {records.value = res.response.records})
})
</script>
<template>
    <!-- container -->
    <div class="flex flex-col items-center gap-5 w-full px-4">
      <!-- separator -->
      <TopRank
        v-for="(g, idx) in records.sort((a, b) => a.score < b.score ? 1 : 0).filter((_, i:number) => i < 3)"
        :key="g.group.id"
        :rank="idx + 1"
        :name="g.group.id"
        :score="g.score"
      />
      <!-- top rank and normal rank separator -->
      <hr class="h-[2px] w-11/12 mx-8 text-white bg-gray-500 border-0" />
      <RankItem
        v-for="(g, idx) in records.sort((a, b) => a.score < b.score ? 1 : 0).filter((_, i:number) => i >= 3)"
        :key="g.group.id"
        :rank="idx + 4"
        :name="g.group.id"
        :score="g.score"
      />
    </div>
</template>
