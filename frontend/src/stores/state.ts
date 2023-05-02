import { defineStore } from "pinia";

// TODO: add logout action

export const useStateStore = defineStore('state', {
  state: () => ({
    token: '',
    group: '',
    benchmarking: false,
    benchmarkInterval: 0,
    lastResult: 0,
    records: []
  })
})
