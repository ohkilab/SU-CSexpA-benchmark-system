import { defineStore } from "pinia";

export interface IAdminState {
}

export const useAdminStateStore = defineStore<"adminState", IAdminState>("adminState", {
  state: (): IAdminState => ({
  }),
});
