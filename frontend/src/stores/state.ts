import { defineStore } from "pinia";

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
