<script setup lang="ts">
import RankItem from '../components/RankItem.vue'
import TopRank from '../components/TopRank.vue'
import * as proto from '../../../proto-gen/ts/backend/messages'

const groups: { name: string; score: number }[] = [
  {
    name: 'a1',
    score: 1000
  },
  {
    name: 'b2',
    score: 10000
  },
  {
    name: 'a2',
    score: 2000
  },
  {
    name: 'a3',
    score: 1234
  },
  {
    name: 'a4',
    score: 5431
  },
  {
    name: 'b2',
    score: 50
  },
  {
    name: 'b3',
    score: 51
  }
]
</script>
<template>
    <!-- container -->
    <div class="flex flex-col items-center gap-5 w-full px-4">
      <!-- separator -->
      <TopRank
        v-for="(g, idx) in groups.sort((a, b) => a.score < b.score ? 1 : 0).filter((_, i:number) => i < 3)"
        :key="g.name"
        :rank="idx + 1"
        :name="g.name"
        :score="g.score"
      />
      <!-- top rank and normal rank separator -->
      <hr class="h-[2px] w-11/12 mx-8 text-white bg-gray-500 border-0" />
      <RankItem
        v-for="(g, idx) in groups.sort((a, b) => a.score < b.score ? 1 : 0).filter((_, i:number) => i >= 3)"
        :key="g.name"
        :rank="idx + 4"
        :name="g.name"
        :score="g.score"
      />
    </div>
</template>
