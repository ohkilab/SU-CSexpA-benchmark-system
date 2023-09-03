import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { defineStore } from 'pinia'
import { AdminServiceClient } from 'proto-gen-web/services/backend/services.client'

export interface IAdminState {
  currentPath: string
  admin: AdminServiceClient
}

// backend: new BackendServiceClient(
//   new GrpcWebFetchTransport({
//     baseUrl: `http://${window.location.hostname}:8080`,
//   }),

export const useAdminStateStore = defineStore<'adminState', IAdminState>(
  'adminState',
  {
    state: (): IAdminState => ({
      currentPath: '',
      admin: new AdminServiceClient(
        new GrpcWebFetchTransport({
          baseUrl: `http://${window.location.hostname}:8080`
        })
      )
    }),
    persist: true
  }
)
