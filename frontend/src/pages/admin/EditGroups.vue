<script setup lang="ts">
import { ref, Ref } from "vue";
import {
  CreateGroupsRequest,
  CreateGroupsRequest_CreateGroupsGroup,
} from "proto-gen-web/services/backend/messages";
import { AdminServiceClient } from "proto-gen-web/services/backend/services.client";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { useStateStore } from "../../stores/state";
import { Group, Role } from "proto-gen-web/services/backend/resources";

const state = useStateStore();

const groups: Ref<CreateGroupsRequest> = ref({ groups: [] });
const group: Ref<CreateGroupsRequest_CreateGroupsGroup> = ref({
  name: "",
  password: "",
  year: 0,
  role: Role.GUEST,
});

const createGroups = () => {
  let admin = new AdminServiceClient(
    new GrpcWebFetchTransport({
      baseUrl: import.meta.env.PROD
        ? `http://${window.location.hostname}:8080`
        : state.devBaseUrl,
    }),
  );

  const opt = { meta: { authorization: "Bearer " + state.token } };

  admin.createGroups({ groups: [group.value] }, opt).then((res) => {
    if (import.meta.env.DEV) console.log(res);
  });
};
</script>
<template>
  <div class="flex h-full flex-col gap-2">
    <h1 class="text-xl">グループ</h1>
    <h2>グループ名</h2>
    <input
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
      placeholder="タイトル"
      type="text"
    />
    <h2>パスワード</h2>
    <input
      type="number"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
      value="0"
    />
    <h2>年度</h2>
    <input
      type="number"
      class="w-full rounded bg-gray-500 p-2 transition hover:bg-gray-600 focus:bg-gray-600 focus:outline-none"
      value="2000"
    />
    <h2>ロール</h2>

    <div
      v-for="(r, idx) in Object.values(Role).filter((el) => typeof el === 'string')"
    >
    <input type="radio" name="role" :value="idx" />
      {{
        r.toString().toLowerCase().charAt(0).toUpperCase() +
        r.toString().toLowerCase().slice(1)
      }}
    </div>

    <button @click="createGroups" class="mt-auto w-fit rounded bg-blue-500 p-2">
      グループ作成
    </button>
  </div>
</template>
