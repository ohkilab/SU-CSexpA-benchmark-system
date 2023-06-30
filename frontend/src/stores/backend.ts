import { defineStore } from "pinia";
import { BackendServiceClient } from 'proto-gen-web/services/backend/services.client';
import {
  GetRankingResponse_Record,
  PostLoginRequest,
} from "proto-gen-web/services/backend/messages";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";

export interface IBackendStore {
  backend: BackendServiceClient
}

export const useBackendStore = defineStore<"state", IBackendStore>("state", {
  state: (): IBackendStore => ({
    backend: null
  }),
  actions: {},
  persist: false,
});
