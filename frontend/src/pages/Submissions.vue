<script setup lang="ts">
import { onMounted, Ref, ref, watch } from "vue";
import { IState, useStateStore } from "../stores/state";
import { Status, Submit } from "proto-gen-web/services/backend/resources";
import {
  GetSubmitRequest,
  ListSubmitsRequest,
} from "proto-gen-web/services/backend/messages";
import Result from "../components/Result.vue";
import { useBackendStore } from "../stores/backend";
import { useRouter, useRoute } from "vue-router";

const state: IState = useStateStore();
const { backend } = useBackendStore();

const noSubmissions: Ref<boolean> = ref(false);

const modalItem: Ref<Partial<Submit>> = ref({});
const totalPages: Ref<number> = ref(0);
const currentPage: Ref<number> = ref(1);

const router = useRouter();
const route = useRoute();

const formatDate = (timestamp: number): string => {
  const dateObject: Date = new Date(timestamp * 1000);
  const date: string = dateObject.toLocaleDateString();
  const time: string = dateObject.toLocaleTimeString();

  return `${date} ${time}`;
};

const handleModal = (id: number) => {
  router.push(`/submissions/${id}`);
};

const getSubmitById = async (id: number): Promise<Submit | undefined> => {
  const getSubmitRequest: GetSubmitRequest = {
    submitId: id,
  };

  const opt = { meta: { authorization: "Bearer " + state.token } };
  const call = backend.getSubmit(getSubmitRequest, opt);
  for await (let message of call.responses) {
    if (import.meta.env.DEV) console.log("Submit", message);
    return message.submit;
  }
};

const handleCloseModal = () => {
  router.push("/submissions");
};

const listSubmitsByPage = (page: number) => {
  const opt = { meta: { authorization: "Bearer " + state.token } };
  // TODO: get own submissions, filter functionality
  const listSubmitsRequest: ListSubmitsRequest = {
    contestSlug: state.contestSlug,
    page,
    // groupName: 'a01',
    // status: Status.VALIDATION_ERROR
  };

  backend.listSubmits(listSubmitsRequest, opt).then((res) => {
    if (import.meta.env.DEV) console.log("Submits", res.response.submits);
    state.submits = res.response.submits;
    totalPages.value = res.response.totalPages;

    if (res.response.submits.length == 0) {
      noSubmissions.value = true;
    }
  });
};

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
    listSubmitsByPage(currentPage.value);
  }
};

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--;
    listSubmitsByPage(currentPage.value);
  }
};

const gotoPage = (page: number) => {
  if (page > 0 && page <= totalPages.value) {
    currentPage.value = page;
    listSubmitsByPage(currentPage.value);
  }
};

watch(
  () => route.params.id,
  async (to) => {
    //toParams, prevParams
    if (import.meta.env.DEV) console.log(to);
    if (to) {
      const submit = await getSubmitById(Number(to));
      modalItem.value = submit ?? {};
    }
  },
);

watch(
  () => currentPage,
  (to) => {
    listSubmitsByPage(to.value);
  },
);

onMounted(() => {
  listSubmitsByPage(currentPage.value);

  // in case of direct access
  if (route.params.id) {
    (async () => {
      const submit = await getSubmitById(Number(route.params.id));
      modalItem.value = submit ?? {};
    })();
  }
});
</script>
<template>
  <!-- TODO: show "no submissions" when server returns no submissions -->

  <!-- modal -->
  <transition
    enter-active-class="duration-100 ease-out"
    enter-from-class="transform opacity-0"
    enter-to-class="opacity-100"
    leave-active-class="duration-100 ease-in"
    leave-from-class="opacity-100"
    leave-to-class="transform opacity-0"
  >
    <div
      v-if="$route.params.id"
      @click.self="modalItem = {}"
      class="fixed inset-0 flex h-full w-full items-center justify-center overflow-y-auto bg-black bg-opacity-50"
    >
      <div class="h-5/6 w-5/6">
        <result
          :submit="modalItem"
          :title="'結果詳細'"
          :show-close-button="true"
          @close-modal="handleCloseModal"
        />
      </div>
    </div>
  </transition>
  <div class="flex w-full justify-center gap-5">
    <button
      class="rounded bg-gray-700 px-3 py-2 shadow shadow-black transition hover:scale-105"
      @click="prevPage"
    >
      <font-awesome-icon :icon="['fas', 'arrow-left']"></font-awesome-icon>
    </button>
    <button
      v-for="page in totalPages"
      :key="page"
      @click="gotoPage(page)"
      class="rounded px-3 py-2 shadow shadow-black transition hover:scale-105"
      :class="page == currentPage ? 'bg-blue-500' : 'bg-gray-700'"
    >
      {{ page }}
    </button>
    <button
      class="rounded bg-gray-700 px-3 py-2 shadow shadow-black transition hover:scale-105"
      @click="nextPage"
    >
      <font-awesome-icon :icon="['fas', 'arrow-right']"></font-awesome-icon>
    </button>
  </div>
  <table v-if="state.submits.length > 0" class="table-auto mb-4">
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
      <tr
        v-for="s in state.submits"
        class="cursor-pointer border-b-2 border-gray-800 bg-gray-900 transition hover:bg-gray-700"
        @click.prevent="handleModal(s.id)"
        key="s.id"
      >
        <td class="w-20 text-center">{{ s.id }}</td>
        <td class="w-60 text-center">
          {{ formatDate(Number(s.submitedAt?.seconds)) }}
        </td>
        <td class="w-30 text-center">{{ s.groupName }}</td>
        <td class="w-20 px-5 text-center">
          <div class="w-20 justify-center rounded bg-gray-500 text-center">
            {{ s.score }}
          </div>
        </td>
        <td class="py-2 text-center transition-colors">
          <div
            v-if="s.status == Status.WAITING"
            class="mx-auto w-40 rounded bg-teal-500 p-1"
          >
            Waiting
          </div>
          <div
            v-else-if="s.status == Status.IN_PROGRESS"
            class="mx-auto w-40 rounded bg-teal-500 p-1"
          >
            In Progress
          </div>
          <div
            v-else-if="s.status == Status.SUCCESS"
            class="mx-auto w-40 rounded bg-blue-600 p-1"
          >
            Success
          </div>
          <div
            v-else-if="s.status == Status.CONNECTION_FAILED"
            class="mx-auto w-40 rounded bg-red-600 p-1"
          >
            Connection Failed
          </div>
          <div
            v-else-if="s.status == Status.VALIDATION_ERROR"
            class="mx-auto w-40 rounded bg-orange-500 p-1"
          >
            Validation Error
          </div>
          <div
            v-else-if="s.status == Status.TIMEOUT"
            class="mx-auto w-40 rounded bg-orange-500 p-1"
          >
            Timeout
          </div>
          <div
            v-else-if="s.status == Status.INTERNAL_ERROR"
            class="mx-auto w-40 rounded bg-orange-500 p-1"
          >
            Internal Error
          </div>
          <div v-else class="w-40 rounded bg-orange-500 p-1">Unknown Error</div>
        </td>
      </tr>
    </tbody>
  </table>

  <div class="" v-else>
    <font-awesome-icon
      v-if="!noSubmissions"
      class="animate-spin text-3xl"
      :icon="['fas', 'spinner']"
    ></font-awesome-icon>
    <div v-else>まだベンチマーク結果がありません。</div>
  </div>
  <div v-if="state.debug" class="flex w-full flex-col gap-2">
    <div class="w-40 rounded bg-teal-500 p-1">Waiting</div>
    <div class="w-40 rounded bg-teal-500 p-1">In Progress</div>
    <div class="w-40 rounded bg-blue-600 p-1">Success</div>
    <div class="w-40 rounded bg-red-600 p-1">Connection Failed</div>
    <div class="w-40 rounded bg-orange-500 p-1">Validation Error</div>
    <div class="w-40 rounded bg-orange-500 p-1">Internal Error</div>
    {{ Status }}
  </div>
</template>
