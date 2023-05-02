import { defineStore } from "pinia";

export const useStateStore = defineStore('state', {
  state: () => ({
    token: '',
    group: '',
    records: []
  })
})
