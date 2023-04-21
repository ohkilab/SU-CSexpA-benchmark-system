<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue'
import RankItem from './components/RankItem.vue'
import TopRank from './components/TopRank.vue'

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
  <div
    class="text-white items-center bg-gray-800 min-h-screen flex flex-col gap-6"
  >
    <!-- app bar -->
    <div
      class="w-screen h-16 items-center bg-gray-700 flex shadow-md shadow-gray-950"
    >
      <div class="mx-auto text-lg sm:text-xl">
        情報科学実験A：ベンチマークサーバランキング
      </div>
    </div>
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
    <!-- footer -->
    <div class="mt-20"></div>
  </div>
</template>
