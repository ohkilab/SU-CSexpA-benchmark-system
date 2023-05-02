import { defineStore } from 'pinia'

export const useCredStore = defineStore('cred', {
  state: () => ({
    token: ''
  })
})
