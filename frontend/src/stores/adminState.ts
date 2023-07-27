import { defineStore } from "pinia";

export interface IAdminState {
  currentPath: string;
}

export const useAdminStateStore = defineStore<"adminState", IAdminState>("adminState", {
  state: (): IAdminState => ({
    currentPath: ''
  }),
  persist: true,
});
