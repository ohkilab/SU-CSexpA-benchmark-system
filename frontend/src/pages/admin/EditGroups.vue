<script setup lang="ts">
import { ref } from "vue";
import type { CreateGroupsRequest_CreateGroupsGroup } from "proto-gen-web/services/backend/messages";
import { Role } from "proto-gen-web/services/backend/resources";
import { useAdminStateStore } from "../../stores/adminState";
import { useStateStore } from "../../stores/state";

type NoticeType = "success" | "error";

interface GroupRow {
  id: number;
  name: string;
  password: string;
  year: number;
  role: Role;
}

let nextRowId = 1;
const currentYear = new Date().getFullYear();

const newGroupRow = (): GroupRow => ({
  id: nextRowId++,
  name: "",
  password: "",
  year: currentYear,
  role: Role.CONTESTANT,
});

const adminState = useAdminStateStore();
const state = useStateStore();

const rows = ref<GroupRow[]>([newGroupRow()]);
const saving = ref(false);
const notice = ref<{ type: NoticeType; message: string } | null>(null);

const roleOptions = [
  { label: "CONTESTANT", value: Role.CONTESTANT },
  { label: "GUEST", value: Role.GUEST },
  { label: "ADMIN", value: Role.ADMIN },
];

const authOptions = () => ({
  meta: { authorization: "Bearer " + state.token },
});

const setNotice = (type: NoticeType, message: string) => {
  notice.value = { type, message };
};

const formatError = (error: unknown): string => {
  if (error instanceof Error) return error.message;
  return String(error);
};

const addRow = () => {
  rows.value.push(newGroupRow());
};

const removeRow = (index: number) => {
  if (rows.value.length <= 1) return;
  rows.value.splice(index, 1);
};

const buildGroups = (): CreateGroupsRequest_CreateGroupsGroup[] | null => {
  const groups = rows.value.map((row) => ({
    name: row.name.trim(),
    password: row.password,
    year: row.year,
    role: Number(row.role) as Role,
  }));

  const invalidIndex = groups.findIndex(
    (group) =>
      group.name.length === 0 ||
      group.password.trim().length === 0 ||
      group.year <= 0 ||
      !roleOptions.some((option) => option.value === group.role),
  );
  if (invalidIndex >= 0) {
    setNotice("error", `${invalidIndex + 1}行目の入力を確認してください。`);
    return null;
  }

  const seen = new Set<string>();
  const duplicateIndex = groups.findIndex((group) => {
    const key = `${group.name}:${group.year}`;
    if (seen.has(key)) return true;
    seen.add(key);
    return false;
  });
  if (duplicateIndex >= 0) {
    setNotice(
      "error",
      `${duplicateIndex + 1}行目のグループ名と年度が重複しています。`,
    );
    return null;
  }

  return groups;
};

const createGroups = async () => {
  const groups = buildGroups();
  if (!groups) return;

  saving.value = true;
  try {
    const res = await adminState.admin.createGroups({ groups }, authOptions());
    setNotice(
      "success",
      `${res.response.groups.length}件のグループを作成しました。`,
    );
    rows.value = [newGroupRow()];
  } catch (error) {
    setNotice("error", "グループ作成に失敗しました: " + formatError(error));
  } finally {
    saving.value = false;
  }
};
</script>

<template>
  <div class="flex h-full flex-col gap-5">
    <div class="flex flex-wrap items-center gap-3">
      <h1 class="text-xl">グループ管理</h1>
      <button
        class="rounded bg-gray-600 px-3 py-2 transition hover:bg-gray-500"
        @click="addRow"
      >
        行を追加
      </button>
      <div
        v-if="notice"
        class="rounded px-3 py-2 text-sm"
        :class="notice.type === 'success' ? 'bg-green-700' : 'bg-red-700'"
      >
        {{ notice.message }}
      </div>
    </div>

    <div class="overflow-x-auto">
      <table class="w-full min-w-[760px] table-fixed text-sm">
        <thead class="bg-gray-800">
          <tr>
            <th class="w-52 px-3 py-2 text-left">グループ名</th>
            <th class="w-56 px-3 py-2 text-left">パスワード</th>
            <th class="w-28 px-3 py-2 text-left">年度</th>
            <th class="w-40 px-3 py-2 text-left">ロール</th>
            <th class="w-24 px-3 py-2"></th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(row, index) in rows"
            :key="row.id"
            class="border-b border-gray-800 bg-gray-900"
          >
            <td class="px-3 py-2">
              <input
                v-model="row.name"
                class="w-full rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
                placeholder="group-name"
                type="text"
              />
            </td>
            <td class="px-3 py-2">
              <input
                v-model="row.password"
                class="w-full rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
                placeholder="password"
                type="password"
              />
            </td>
            <td class="px-3 py-2">
              <input
                v-model.number="row.year"
                class="w-full rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
                min="1"
                type="number"
              />
            </td>
            <td class="px-3 py-2">
              <select
                v-model.number="row.role"
                class="w-full rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
              >
                <option
                  v-for="option in roleOptions"
                  :key="option.value"
                  :value="option.value"
                >
                  {{ option.label }}
                </option>
              </select>
            </td>
            <td class="px-3 py-2 text-right">
              <button
                class="rounded bg-red-600 px-3 py-2 transition hover:bg-red-500 disabled:opacity-50"
                :disabled="rows.length <= 1"
                @click="removeRow(index)"
              >
                削除
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <button
      class="w-fit rounded bg-blue-600 px-4 py-2 transition hover:bg-blue-500 disabled:opacity-50"
      :disabled="saving"
      @click="createGroups"
    >
      {{ saving ? "作成中..." : "グループを作成" }}
    </button>
  </div>
</template>
