import { defineStore } from "pinia";
import { GetRankingResponse_Record } from 'proto-gen-web/src/backend/messages';

// TODO: add logout action

export interface IState {
  token: string
  group: string
  benchmarking: boolean
  benchmarkInterval: number
  lastResult: number
  records: Array<GetRankingResponse_Record> | null
}

export const useStateStore = defineStore('state', {
  state: (): IState => ({
    token: '',
    group: '',
    benchmarking: false,
    benchmarkInterval: 0,
    lastResult: 0,
    records: []
  })
})
