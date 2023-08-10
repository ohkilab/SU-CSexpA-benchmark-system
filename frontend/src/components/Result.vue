<script setup lang="ts">
import { Submit, Status } from "proto-gen-web/services/backend/resources";

const formatDate = (timestamp: number): string => {
  const dateObject: Date = new Date(timestamp * 1000);
  const date: string = dateObject.toLocaleDateString();
  const time: string = dateObject.toLocaleTimeString();

  return `${date} ${time}`;
};

const props = defineProps<{
  submit: Partial<Submit>;
  title: string;
  showCloseButton?: boolean;
}>();

const emit = defineEmits(["closeModal"]);
</script>
<template>
  <div
    class="mx-auto flex h-full w-full flex-col gap-2 overflow-y-auto rounded bg-gray-700 p-10"
  >
    <!-- title row -->
    <div class="flex items-center">
      <!-- <div class="text-2xl">グループ {{ props.submit.groupName }} 提出ID: {{ props.submit.id }} 結果詳細</div> -->
      <div class="text-2xl">{{ props.title }}</div>
      <div class="ml-auto flex items-center justify-end gap-2">
        <div class="text-center">
          <div
            v-if="props.submit.status == Status.WAITING"
            class="mx-2 w-40 rounded bg-teal-500 p-1"
          >
            Waiting
          </div>
          <div
            v-else-if="props.submit.status == Status.IN_PROGRESS"
            class="mx-2 w-40 rounded bg-teal-500 p-1"
          >
            In Progress
          </div>
          <div
            v-else-if="props.submit.status == Status.SUCCESS"
            class="mx-2 w-40 rounded bg-blue-600 p-1"
          >
            Success
          </div>
          <div
            v-else-if="props.submit.status == Status.CONNECTION_FAILED"
            class="mx-2 w-40 rounded bg-red-600 p-1"
          >
            Connection Failed
          </div>
          <div
            v-else-if="props.submit.status == Status.VALIDATION_ERROR"
            class="mx-2 w-40 rounded bg-orange-500 p-1"
          >
            Validation Error
          </div>
          <div
            v-else-if="props.submit.status == Status.TIMEOUT"
            class="mx-2 w-40 rounded bg-orange-500 p-1"
          >
            Timeout
          </div>
          <div
            v-else-if="props.submit.status == Status.INTERNAL_ERROR"
            class="mx-2 w-40 rounded bg-orange-500 p-1"
          >
            Internal Error
          </div>
          <div v-else class="w-40 rounded bg-orange-500 p-1">Unknown Error</div>
        </div>
        <button
          v-if="showCloseButton"
          @click="() => emit('closeModal')"
          class="rounded bg-red-500 px-4 py-2 shadow-md shadow-black transition hover:bg-red-600"
        >
          <font-awesome-icon :icon="['fas', 'x']"></font-awesome-icon>
        </button>
      </div>
    </div>
    <div class="flex gap-2">
      <div class="flex gap-2">
        グループ {{ props.submit.groupName }} 提出 ID: {{ props.submit.id }}
      </div>
      <div class="flex gap-2">
        得点:
        <div class="w-20 justify-center rounded bg-gray-500 text-center">
          {{ props.submit.score }}
        </div>
      </div>
    </div>
    <div class="text-md text-gray-300">
      提出日時: {{ formatDate(Number(props.submit.submitedAt?.seconds)) }}
    </div>
    <div
      class="flex h-full w-full flex-col gap-2 overflow-y-auto rounded bg-gray-900 p-8"
    >
      <div
        v-for="(t, i) in props.submit.taskResults"
        :key="i"
        class="flex items-center justify-between gap-2 rounded px-5 py-3 shadow-md shadow-black"
        :class="
          t.status == Status.WAITING
            ? 'opacity-70'
            : t.status == Status.IN_PROGRESS
            ? 'bg-teal-500'
            : t.status == Status.SUCCESS
            ? 'bg-blue-600'
            : t.status == Status.CONNECTION_FAILED
            ? 'bg-red-500'
            : t.status == Status.VALIDATION_ERROR
            ? 'bg-orange-500'
            : t.status == Status.INTERNAL_ERROR
            ? 'bg-orange-500'
            : 'bg-gray-700 opacity-70'
        "
      >
        <div class="flex items-center justify-center gap-2">
          タグ {{ i + 1 }}：
          <div class="rounded bg-gray-500 px-2">{{ t.requestPerSec }}</div>
          req/s
        </div>
        <div class="flex items-center gap-5">
          {{ t.errorMessage != "" ? `エラー: ${t.errorMessage}` : "" }}
          <font-awesome-icon
            v-if="t.status == Status.IN_PROGRESS"
            :icon="['fas', 'spinner']"
          ></font-awesome-icon>
          <font-awesome-icon
            v-else-if="t.status == Status.WAITING"
            :icon="['fas', 'minus']"
          ></font-awesome-icon>
          <font-awesome-icon
            v-else-if="t.status == Status.SUCCESS"
            :icon="['fas', 'check']"
          ></font-awesome-icon>
          <font-awesome-icon
            v-else-if="t.status == Status.CONNECTION_FAILED"
            :icon="['fas', 'x']"
          ></font-awesome-icon>
          <font-awesome-icon
            v-else-if="t.status == Status.VALIDATION_ERROR"
            :icon="['fas', 'exclamation']"
          ></font-awesome-icon>
          <font-awesome-icon
            v-else
            :icon="['fas', 'minus']"
          ></font-awesome-icon>
        </div>
      </div>
    </div>
  </div>
</template>
