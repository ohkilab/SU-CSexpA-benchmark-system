<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { onMounted, Ref, ref } from "vue";
import RankItem from "../components/RankItem.vue";
import TopRank from "../components/TopRank.vue";
import { ListSubmitsRequest } from "proto-gen-web/services/backend/messages";
import Graph from "../components/Graph.vue";

import { useStateStore } from "../stores/state";
import { useBackendStore } from "../stores/backend";

const state = useStateStore();
const { backend } = useBackendStore();

const hasRanking: Ref<boolean> = ref(false);
const hasSubmits: Ref<boolean> = ref(false);

onMounted(() => {
  let opt = { meta: { authorization: "Bearer " + state.token } };
  const listSubmitsRequest: ListSubmitsRequest = {
    contestSlug: state.contestSlug,
    page: 1,
    // groupName: 'a01',
    // status: Status.VALIDATION_ERROR
  };

  backend
    .getRanking(
      {
        contestSlug: state.contestSlug,
        containGuest: false,
      },
      opt,
    )
    .then((res) => {
      if (import.meta.env.DEV) console.log(res.response.records);
      state.records = res.response.records ?? [];

      hasRanking.value = true;
    });

  backend.listSubmits(listSubmitsRequest, opt).then((res) => {
    if (import.meta.env.DEV) console.log("Submits", res.response.submits);
    state.submits = res.response.submits;

    hasSubmits.value = true;
  });
});
</script>
<template>
  <div
    class="flex h-80 w-11/12 items-center justify-center rounded-md bg-gray-200 p-4 text-black sm:w-5/6"
  >
    <graph v-if="hasRanking && hasSubmits"></graph>
    <div v-else class="">読み込み中...</div>
  </div>
  <!-- container -->
  <div
    v-if="state.records.length > 0"
    class="flex w-full flex-col items-center gap-5 px-4"
  >
    <!-- separator -->
    <TopRank
      v-for="g in state.records.filter((_, i: number) => i < 3)"
      :key="g.group?.name"
      :rank="g.rank"
      :class="state.group == g.group?.name ? 'bg-blue-700' : 'bg-gray-700'"
      :name="g.group?.name ?? ''"
      :score="g.score ?? 0"
    />
    <!-- top rank and normal rank separator -->
    <hr class="mx-8 h-[2px] w-11/12 border-0 bg-gray-500 text-white" />
    <RankItem
      v-for="g in state.records.filter((_, i: number) => i >= 3)"
      :key="g.group?.name"
      :rank="g.rank"
      :class="state.group == g.group?.name ? 'bg-blue-700' : 'bg-gray-700'"
      :name="g.group?.name ?? ''"
      :score="g.score ?? 0"
    />
  </div>
  <div class="mt-auto" v-else>
    <font-awesome-icon
      class="animate-spin text-3xl"
      :icon="['fas', 'spinner']"
    ></font-awesome-icon>
  </div>
</template>
