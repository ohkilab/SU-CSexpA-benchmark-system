<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import type {
  CreateContestRequest,
  UpdateContestRequest,
} from "proto-gen-web/services/backend/messages";
import {
  Contest,
  TagSelectionLogicType,
  Validator,
} from "proto-gen-web/services/backend/resources";
import type { Timestamp } from "proto-gen-web/google/protobuf/timestamp";
import { useAdminStateStore } from "../../stores/adminState";
import { useBackendStore } from "../../stores/backend";
import { useStateStore } from "../../stores/state";

type TagMode = "auto" | "manual";
type NoticeType = "success" | "error";

interface ManualAttempt {
  id: number;
  tags: string;
}

interface CreateContestForm {
  title: string;
  slug: string;
  startAt: string;
  endAt: string;
  submitLimit: number;
  validator: Validator;
  timeLimitPerTask: number;
  tagMode: TagMode;
  autoTags: string;
  manualAttempts: ManualAttempt[];
}

interface EditContestForm {
  contestSlug: string;
  title: string;
  startAt: string;
  endAt: string;
  submitLimit: number;
  validator: Validator;
}

let nextAttemptId = 1;

const newManualAttempt = (): ManualAttempt => ({
  id: nextAttemptId++,
  tags: "",
});

const newCreateForm = (): CreateContestForm => ({
  title: "",
  slug: "",
  startAt: "",
  endAt: "",
  submitLimit: 1,
  validator: Validator.V2023,
  timeLimitPerTask: 30,
  tagMode: "auto",
  autoTags: "",
  manualAttempts: [newManualAttempt()],
});

const adminState = useAdminStateStore();
const backendStore = useBackendStore();
const state = useStateStore();

const contests = ref<Contest[]>([]);
const selectedContestSlug = ref("");
const createForm = ref<CreateContestForm>(newCreateForm());
const editForm = ref<EditContestForm>({
  contestSlug: "",
  title: "",
  startAt: "",
  endAt: "",
  submitLimit: 1,
  validator: Validator.V2023,
});
const loading = ref(false);
const savingCreate = ref(false);
const savingUpdate = ref(false);
const notice = ref<{ type: NoticeType; message: string } | null>(null);

const validatorOptions = [
  { label: "V2022", value: Validator.V2022 },
  { label: "V2023", value: Validator.V2023 },
];

const selectedContest = computed(() =>
  contests.value.find((contest) => contest.slug === selectedContestSlug.value),
);

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

const toTimestamp = (value: string): Timestamp | undefined => {
  if (!value) return undefined;
  const milliseconds = new Date(value).getTime();
  if (Number.isNaN(milliseconds)) return undefined;
  return {
    seconds: BigInt(Math.floor(milliseconds / 1000)),
    nanos: (milliseconds % 1000) * 1000000,
  };
};

const toDateTimeLocal = (timestamp?: Timestamp): string => {
  if (!timestamp) return "";
  const date = new Date(
    Number(timestamp.seconds) * 1000 + timestamp.nanos / 1000000,
  );
  const pad = (value: number) => value.toString().padStart(2, "0");
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(
    date.getDate(),
  )}T${pad(date.getHours())}:${pad(date.getMinutes())}`;
};

const formatDateTime = (timestamp?: Timestamp): string => {
  if (!timestamp) return "";
  return new Date(
    Number(timestamp.seconds) * 1000 + timestamp.nanos / 1000000,
  ).toLocaleString();
};

const splitTags = (value: string): string[] =>
  value
    .split(/\r?\n/)
    .map((tag) => tag.trim())
    .filter((tag) => tag.length > 0);

const importAutoTagsFromFile = async (event: Event) => {
  const input = event.target as HTMLInputElement;
  const file = input.files?.[0];
  input.value = "";
  if (!file) return;

  try {
    const text = await file.text();
    const tags = splitTags(text);
    if (tags.length === 0) {
      setNotice("error", "Autoタグファイルにタグが含まれていません。");
      return;
    }
    createForm.value.autoTags = text;
    setNotice("success", `${tags.length}件のAutoタグを読み込みました。`);
  } catch (error) {
    setNotice(
      "error",
      "Autoタグファイルの読み込みに失敗しました: " + formatError(error),
    );
  }
};

const getDirectChildFileName = (file: File): string | null => {
  const relativePath = file.webkitRelativePath || file.name;
  const parts = relativePath.split("/").filter((part) => part.length > 0);
  if (parts.length === 0 || parts.length > 2) return null;
  return parts[parts.length - 1];
};

const importManualTagsFromDirectory = async (event: Event) => {
  const input = event.target as HTMLInputElement;
  const files = Array.from(input.files ?? []);
  input.value = "";
  if (files.length === 0) return;

  try {
    const filesByAttempt = new Map<number, File>();

    for (const file of files) {
      const fileName = getDirectChildFileName(file);
      const match = fileName?.match(/^([1-9]\d*)\.txt$/);
      if (!match) continue;

      const attemptNumber = Number(match[1]);
      if (filesByAttempt.has(attemptNumber)) {
        throw new Error(`${attemptNumber}.txt が重複しています。`);
      }
      filesByAttempt.set(attemptNumber, file);
    }

    if (!filesByAttempt.has(1)) {
      throw new Error("1.txt が見つかりません。");
    }

    const attemptNumbers = [...filesByAttempt.keys()].sort((a, b) => a - b);
    for (let expected = 1; expected <= attemptNumbers.length; expected++) {
      if (attemptNumbers[expected - 1] !== expected) {
        throw new Error(`${expected}.txt が見つかりません。`);
      }
    }

    const tagTexts: string[] = [];
    for (const attemptNumber of attemptNumbers) {
      const text = await filesByAttempt.get(attemptNumber)!.text();
      if (splitTags(text).length === 0) {
        throw new Error(`${attemptNumber}.txt にタグが含まれていません。`);
      }
      tagTexts.push(text);
    }

    createForm.value.manualAttempts = tagTexts.map((tags) => ({
      id: nextAttemptId++,
      tags,
    }));
    setNotice(
      "success",
      `${tagTexts.length}試行分のManualタグを読み込みました。`,
    );
  } catch (error) {
    setNotice(
      "error",
      "Manualタグディレクトリの読み込みに失敗しました: " +
        formatError(error),
    );
  }
};

const validateDateRange = (startAt: string, endAt: string): boolean => {
  const start = new Date(startAt).getTime();
  const end = new Date(endAt).getTime();
  return !Number.isNaN(start) && !Number.isNaN(end) && start < end;
};

const fetchContests = async () => {
  loading.value = true;
  try {
    const res = await backendStore.backend.listContests({}, authOptions());
    contests.value = res.response.contests ?? [];
    if (selectedContestSlug.value) {
      const current = contests.value.find(
        (contest) => contest.slug === selectedContestSlug.value,
      );
      if (current) selectContest(current);
    }
  } catch (error) {
    setNotice(
      "error",
      "コンテスト一覧の取得に失敗しました: " + formatError(error),
    );
  } finally {
    loading.value = false;
  }
};

const selectContest = (contest: Contest) => {
  selectedContestSlug.value = contest.slug;
  editForm.value = {
    contestSlug: contest.slug,
    title: contest.title,
    startAt: toDateTimeLocal(contest.startAt),
    endAt: toDateTimeLocal(contest.endAt),
    submitLimit: contest.submitLimit,
    validator: contest.validator,
  };
};

const addManualAttempt = () => {
  createForm.value.manualAttempts.push(newManualAttempt());
};

const removeManualAttempt = (index: number) => {
  if (createForm.value.manualAttempts.length <= 1) return;
  createForm.value.manualAttempts.splice(index, 1);
};

const buildCreateRequest = (): CreateContestRequest | null => {
  const form = createForm.value;
  const startAt = toTimestamp(form.startAt);
  const endAt = toTimestamp(form.endAt);
  if (!form.title.trim() || !form.slug.trim()) {
    setNotice("error", "タイトルとslugを入力してください。");
    return null;
  }
  if (!startAt || !endAt || !validateDateRange(form.startAt, form.endAt)) {
    setNotice("error", "開始日時と終了日時を正しい順序で入力してください。");
    return null;
  }
  if (form.submitLimit <= 0 || form.timeLimitPerTask <= 0) {
    setNotice("error", "提出回数とタスク制限時間は1以上にしてください。");
    return null;
  }

  const baseRequest = {
    title: form.title.trim(),
    slug: form.slug.trim(),
    startAt,
    endAt,
    submitLimit: form.submitLimit,
    validator: Number(form.validator) as Validator,
    timeLimitPerTask: form.timeLimitPerTask,
  };

  if (form.tagMode === "auto") {
    const tags = splitTags(form.autoTags);
    if (tags.length === 0) {
      setNotice("error", "Autoタグを1件以上入力してください。");
      return null;
    }
    return {
      ...baseRequest,
      tagSelection: {
        oneofKind: "auto",
        auto: {
          type: TagSelectionLogicType.AUTO,
          tags: { tags },
        },
      },
    };
  }

  const tagsList = form.manualAttempts.map((attempt) => ({
    tags: splitTags(attempt.tags),
  }));
  if (tagsList.some((tags) => tags.tags.length === 0)) {
    setNotice("error", "Manualタグは各試行に1件以上入力してください。");
    return null;
  }

  return {
    ...baseRequest,
    tagSelection: {
      oneofKind: "manual",
      manual: {
        type: TagSelectionLogicType.MANUAL,
        tagsList,
      },
    },
  };
};

const createContest = async () => {
  const request = buildCreateRequest();
  if (!request) return;

  savingCreate.value = true;
  try {
    const res = await adminState.admin.createContest(request, authOptions());
    setNotice(
      "success",
      `コンテスト「${
        res.response.contest?.title ?? request.title
      }」を作成しました。`,
    );
    createForm.value = newCreateForm();
    await fetchContests();
  } catch (error) {
    setNotice("error", "コンテスト作成に失敗しました: " + formatError(error));
  } finally {
    savingCreate.value = false;
  }
};

const updateContest = async () => {
  const form = editForm.value;
  const startAt = toTimestamp(form.startAt);
  const endAt = toTimestamp(form.endAt);
  if (!form.contestSlug) {
    setNotice("error", "編集するコンテストを選択してください。");
    return;
  }
  if (!form.title.trim()) {
    setNotice("error", "タイトルを入力してください。");
    return;
  }
  if (!startAt || !endAt || !validateDateRange(form.startAt, form.endAt)) {
    setNotice("error", "開始日時と終了日時を正しい順序で入力してください。");
    return;
  }
  if (form.submitLimit <= 0) {
    setNotice("error", "提出回数は1以上にしてください。");
    return;
  }

  const request: UpdateContestRequest = {
    contestSlug: form.contestSlug,
    title: form.title.trim(),
    startAt,
    endAt,
    submitLimit: form.submitLimit,
    validator: Number(form.validator) as Validator,
  };

  savingUpdate.value = true;
  try {
    const res = await adminState.admin.updateContest(request, authOptions());
    setNotice(
      "success",
      `コンテスト「${
        res.response.contest?.title ?? request.title
      }」を更新しました。`,
    );
    await fetchContests();
  } catch (error) {
    setNotice("error", "コンテスト更新に失敗しました: " + formatError(error));
  } finally {
    savingUpdate.value = false;
  }
};

onMounted(() => {
  fetchContests();
});
</script>

<template>
  <div class="flex h-full flex-col gap-5">
    <div class="flex flex-wrap items-center gap-3">
      <h1 class="text-xl">コンテスト管理</h1>
      <button
        class="rounded bg-gray-600 px-3 py-2 transition hover:bg-gray-500 disabled:opacity-50"
        :disabled="loading"
        @click="fetchContests"
      >
        再読み込み
      </button>
      <div
        v-if="notice"
        class="rounded px-3 py-2 text-sm"
        :class="notice.type === 'success' ? 'bg-green-700' : 'bg-red-700'"
      >
        {{ notice.message }}
      </div>
    </div>

    <div
      class="grid gap-5 xl:grid-cols-[minmax(320px,1fr)_minmax(420px,1.3fr)]"
    >
      <section class="flex min-w-0 flex-col gap-3">
        <h2 class="text-lg">既存コンテスト</h2>
        <div class="overflow-x-auto">
          <table class="w-full table-fixed text-sm">
            <thead class="bg-gray-800">
              <tr>
                <th class="w-32 px-3 py-2 text-left">slug</th>
                <th class="w-40 px-3 py-2 text-left">タイトル</th>
                <th class="w-40 px-3 py-2 text-left">開始</th>
                <th class="w-40 px-3 py-2 text-left">終了</th>
                <th class="w-24 px-3 py-2 text-right">提出回数</th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="loading">
                <td class="px-3 py-4 text-center text-gray-300" colspan="5">
                  読み込み中...
                </td>
              </tr>
              <tr v-else-if="contests.length === 0">
                <td class="px-3 py-4 text-center text-gray-300" colspan="5">
                  コンテストがありません
                </td>
              </tr>
              <template v-else>
                <tr
                  v-for="contest in contests"
                  :key="contest.slug"
                  class="cursor-pointer border-b border-gray-800 bg-gray-900 transition hover:bg-gray-700"
                  :class="
                    selectedContestSlug === contest.slug ? 'bg-blue-900' : ''
                  "
                  @click="selectContest(contest)"
                >
                  <td class="truncate px-3 py-2">{{ contest.slug }}</td>
                  <td class="truncate px-3 py-2">{{ contest.title }}</td>
                  <td class="px-3 py-2">
                    {{ formatDateTime(contest.startAt) }}
                  </td>
                  <td class="px-3 py-2">{{ formatDateTime(contest.endAt) }}</td>
                  <td class="px-3 py-2 text-right">
                    {{ contest.submitLimit }}
                  </td>
                </tr>
              </template>
            </tbody>
          </table>
        </div>
      </section>

      <section class="flex min-w-0 flex-col gap-4">
        <h2 class="text-lg">新規作成</h2>
        <div class="grid gap-3 md:grid-cols-2">
          <label class="flex flex-col gap-1">
            <span>タイトル</span>
            <input
              v-model="createForm.title"
              class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
              type="text"
            />
          </label>
          <label class="flex flex-col gap-1">
            <span>slug</span>
            <input
              v-model="createForm.slug"
              class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
              type="text"
            />
          </label>
          <label class="flex flex-col gap-1">
            <span>開始日時</span>
            <input
              v-model="createForm.startAt"
              class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
              type="datetime-local"
            />
          </label>
          <label class="flex flex-col gap-1">
            <span>終了日時</span>
            <input
              v-model="createForm.endAt"
              class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
              type="datetime-local"
            />
          </label>
          <label class="flex flex-col gap-1">
            <span>提出回数</span>
            <input
              v-model.number="createForm.submitLimit"
              class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
              min="1"
              type="number"
            />
          </label>
          <label class="flex flex-col gap-1">
            <span>タスク制限時間(sec)</span>
            <input
              v-model.number="createForm.timeLimitPerTask"
              class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
              min="1"
              type="number"
            />
          </label>
          <label class="flex flex-col gap-1">
            <span>validator</span>
            <select
              v-model.number="createForm.validator"
              class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
            >
              <option
                v-for="option in validatorOptions"
                :key="option.value"
                :value="option.value"
              >
                {{ option.label }}
              </option>
            </select>
          </label>
          <label class="flex flex-col gap-1">
            <span>タグ選択</span>
            <select
              v-model="createForm.tagMode"
              class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
            >
              <option value="auto">Auto</option>
              <option value="manual">Manual</option>
            </select>
          </label>
        </div>

        <div v-if="createForm.tagMode === 'auto'" class="flex flex-col gap-2">
          <div class="flex flex-wrap items-center justify-between gap-3">
            <span>Autoタグ（1行1タグ）</span>
            <label
              class="cursor-pointer rounded bg-gray-600 px-3 py-2 transition hover:bg-gray-500"
            >
              ファイルを読み込み
              <input
                class="sr-only"
                accept=".txt,text/plain"
                type="file"
                @change="importAutoTagsFromFile"
              />
            </label>
          </div>
          <textarea
            v-model="createForm.autoTags"
            class="min-h-28 rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
          ></textarea>
        </div>

        <div v-else class="flex flex-col gap-2">
          <div class="flex flex-wrap items-center justify-between gap-3">
            <span>Manualタグ（試行ごと、1行1タグ）</span>
            <div class="flex flex-wrap gap-2">
              <label
                class="cursor-pointer rounded bg-gray-600 px-3 py-2 transition hover:bg-gray-500"
              >
                ディレクトリを読み込み
                <input
                  class="sr-only"
                  type="file"
                  multiple
                  webkitdirectory
                  @change="importManualTagsFromDirectory"
                />
              </label>
              <button
                class="rounded bg-gray-600 px-3 py-2 transition hover:bg-gray-500"
                @click="addManualAttempt"
              >
                試行を追加
              </button>
            </div>
          </div>
          <div
            v-for="(attempt, index) in createForm.manualAttempts"
            :key="attempt.id"
            class="flex gap-2"
          >
            <label class="flex min-w-0 flex-1 flex-col gap-1">
              <span>{{ index + 1 }}回目</span>
              <textarea
                v-model="attempt.tags"
                class="min-h-24 rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
              ></textarea>
            </label>
            <button
              class="mt-7 h-10 rounded bg-red-600 px-3 transition hover:bg-red-500 disabled:opacity-50"
              :disabled="createForm.manualAttempts.length <= 1"
              @click="removeManualAttempt(index)"
            >
              削除
            </button>
          </div>
        </div>

        <button
          class="w-fit rounded bg-blue-600 px-4 py-2 transition hover:bg-blue-500 disabled:opacity-50"
          :disabled="savingCreate"
          @click="createContest"
        >
          {{ savingCreate ? "作成中..." : "コンテストを作成" }}
        </button>
      </section>
    </div>

    <section class="flex flex-col gap-4 border-t border-gray-600 pt-5">
      <div class="flex flex-wrap items-center gap-3">
        <h2 class="text-lg">選択中コンテストの編集</h2>
        <span
          v-if="selectedContest"
          class="rounded bg-gray-800 px-3 py-1 text-sm"
          >{{ selectedContest.slug }}</span
        >
      </div>
      <div v-if="!selectedContest" class="text-gray-300">
        一覧から編集するコンテストを選択してください。
      </div>
      <div v-else class="grid gap-3 md:grid-cols-2 xl:grid-cols-4">
        <label class="flex flex-col gap-1">
          <span>タイトル</span>
          <input
            v-model="editForm.title"
            class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
            type="text"
          />
        </label>
        <label class="flex flex-col gap-1">
          <span>開始日時</span>
          <input
            v-model="editForm.startAt"
            class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
            type="datetime-local"
          />
        </label>
        <label class="flex flex-col gap-1">
          <span>終了日時</span>
          <input
            v-model="editForm.endAt"
            class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
            type="datetime-local"
          />
        </label>
        <label class="flex flex-col gap-1">
          <span>提出回数</span>
          <input
            v-model.number="editForm.submitLimit"
            class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
            min="1"
            type="number"
          />
        </label>
        <label class="flex flex-col gap-1">
          <span>validator</span>
          <select
            v-model.number="editForm.validator"
            class="rounded bg-gray-500 p-2 focus:bg-gray-600 focus:outline-none"
          >
            <option
              v-for="option in validatorOptions"
              :key="option.value"
              :value="option.value"
            >
              {{ option.label }}
            </option>
          </select>
        </label>
        <label class="flex flex-col gap-1">
          <span>タグ選択</span>
          <input
            class="rounded bg-gray-700 p-2 text-gray-300"
            :value="
              selectedContest.tagSelectionLogic === TagSelectionLogicType.AUTO
                ? 'Auto'
                : 'Manual'
            "
            disabled
          />
        </label>
        <label class="flex flex-col gap-1">
          <span>slug</span>
          <input
            class="rounded bg-gray-700 p-2 text-gray-300"
            :value="selectedContest.slug"
            disabled
          />
        </label>
      </div>
      <button
        v-if="selectedContest"
        class="w-fit rounded bg-blue-600 px-4 py-2 transition hover:bg-blue-500 disabled:opacity-50"
        :disabled="savingUpdate"
        @click="updateContest"
      >
        {{ savingUpdate ? "更新中..." : "コンテストを更新" }}
      </button>
    </section>
  </div>
</template>
