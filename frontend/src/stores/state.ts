import { defineStore } from "pinia";
import { GetRankingResponse_Record } from "proto-gen-web/services/backend/messages";
import { Submit } from 'proto-gen-web/services/backend/resources';

// TODO: add logout action

export interface IState {
  token: string;
  group: string;
  benchmarking: boolean;
  benchmarkInterval: number;
  lastResult: number;
  records: Array<GetRankingResponse_Record>;
  showResult: boolean;
  current: number;
  size: number;
  result: number;
  debug: boolean;
  devBaseUrl: string;
  submits: Submit[];
  contestSlug: string;
  selectedContestName: string;
}

export const useStateStore = defineStore<"state", IState>("state", {
  state: (): IState => ({
    token: "",
    group: "",
    benchmarking: false,
    benchmarkInterval: 0,
    lastResult: 0,
    records: [],
    showResult: false,
    result: 0,
    current: 0,
    size: 0,
    debug: false,
    devBaseUrl: "http://localhost:8080",
    submits: [],
    contestSlug: '2023-qual',
    selectedContestName: ''
  }),
  persist: true,
});
