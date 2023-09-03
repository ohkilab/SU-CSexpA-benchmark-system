import { defineStore } from 'pinia'
import { BackendServiceClient } from 'proto-gen-web/services/backend/services.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'

export interface IBackendStore {
  backend: BackendServiceClient
}

export const useBackendStore = defineStore<'state', IBackendStore>('state', {
  state: (): IBackendStore => ({
    // backend is initialized here, so do this later
    // - only reinitialize for dev environment
    backend: new BackendServiceClient(
      new GrpcWebFetchTransport({
        baseUrl: `http://${window.location.hostname}:8080`
      })
    )
  }),
  actions: {
    getSubmissions() {}
  },
  persist: false
})
