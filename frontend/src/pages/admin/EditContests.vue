<script setup lang="ts">
import { CreateContestRequest } from "proto-gen-web/services/backend/messages";
import { useAdminStateStore } from "../../stores/adminState";
import { reactive, ref } from "vue";
import {
  TagSelectionLogicType,
  Validator,
} from "proto-gen-web/services/backend/resources";
import { Timestamp } from "proto-gen-web/google/protobuf/timestamp";
import { AdminServiceClient } from "proto-gen-web/services/backend/services.client";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { useStateStore } from "../../stores/state";

const state = useStateStore();

const startAt = ref(new Date());

// export interface CreateContestRequest {
//     title: string;
//     startAt?: Timestamp;
//     endAt?: Timestamp;
//     submitLimit: number;
//     slug: string;
//     tagSelection: {
//         oneofKind: "auto";
//         auto: TagSelectionLogicAuto;
//     } | {
//         oneofKind: "manual";
//         manual: TagSelectionLogicManual;
//     } | {
//         oneofKind: undefined;
//     };
//     validator: Validator;
//     timeLimitPerTask: number; // sec
// }

const contest: CreateContestRequest = reactive({
  title: "",
  startAt: {
    seconds: BigInt(0),
    nanos: 0,
  },
  endAt: {
    seconds: BigInt(0),
    nanos: 0,
  },
  submitLimit: 0,
  slug: "",
  tagSelection: {
    oneofKind: undefined,
  },
  validator: Validator.V2023,
  timeLimitPerTask: 0,
});

const updateStartAt = (e: Event) => {
  contest.startAt = Timestamp.fromDate(
    new Date((e.target as HTMLInputElement).value),
  );
};

const updateEndAt = (e: Event) => {
  contest.endAt = Timestamp.fromDate(
    new Date((e.target as HTMLInputElement).value),
  );
};

// fix BigInt problem
if (import.meta.env.DEV)
  BigInt.prototype.toJSON = function () {
    return this.toString();
  };

const createContest = () => {
  console.log(contest);
  const admin = new AdminServiceClient(
    new GrpcWebFetchTransport({
      baseUrl: import.meta.env.PROD
        ? `http://${window.location.hostname}:8080`
        : state.devBaseUrl,
    }),
  );

  const opt = { meta: { authorization: "Bearer " + state.token } };

  admin.createContest(contest, opt).then((res) => {
    console.log(res);
  });
};
</script>
<template>
  <!-- {{ contest }} -->
  <div class="flex h-full flex-col gap-2">
    <h1 class="text-xl">コンテスト</h1>
    <h2>タイトル</h2>
    <input
      v-model="contest.title"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
      placeholder="タイトル"
      type="text"
    />
    <h2>提出回数</h2>
    <input
      type="number"
      v-model="contest.submitLimit"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
    />
    <h2>開始日</h2>
    <input
      @input="updateStartAt"
      type="date"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
    />
    <h2>終了日</h2>
    <input
      @input="updateEndAt"
      type="date"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
    />
    <h2>slug</h2>
    <input
      v-model="contest.slug"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
      placeholder="slug"
      type="text"
    />

    <h2>timeLimitPerTask(seconds)</h2>
    <input
      type="number"
      v-model="contest.timeLimitPerTask"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
    />
    <button
      @click="createContest"
      class="mt-auto w-fit rounded bg-blue-500 p-2"
    >
      コンテストを作成
    </button>
  </div>
</template>
