import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { defineStore } from "pinia";
import { AdminServiceClient } from "proto-gen-web/services/backend/services.client";

export interface IAdminState {
  currentPath: string;
  admin: AdminServiceClient;
}

export const useAdminStateStore = defineStore("adminState", {
  state: (): IAdminState => ({
    currentPath: "",
    admin: new AdminServiceClient(
      new GrpcWebFetchTransport({
        baseUrl: `http://${window.location.hostname}:8080`,
      }),
    ),
  }),
  actions: {
    setBaseUrl(baseUrl: string) {
      this.admin = new AdminServiceClient(
        new GrpcWebFetchTransport({
          baseUrl,
        }),
      );
    },
  },
  persist: {
    paths: ["currentPath"],
  },
});
