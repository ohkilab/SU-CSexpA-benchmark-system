/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { map } from "rxjs/operators";
import { ExecuteRequest, ExecuteResponse } from "./messages";

export const protobufPackage = "benchmark";

export interface BenchmarkService {
  Execute(request: ExecuteRequest): Observable<ExecuteResponse>;
}

export class BenchmarkServiceClientImpl implements BenchmarkService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "benchmark.BenchmarkService";
    this.rpc = rpc;
    this.Execute = this.Execute.bind(this);
  }
  Execute(request: ExecuteRequest): Observable<ExecuteResponse> {
    const data = ExecuteRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(this.service, "Execute", data);
    return result.pipe(map((data) => ExecuteResponse.decode(_m0.Reader.create(data))));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
  clientStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Promise<Uint8Array>;
  serverStreamingRequest(service: string, method: string, data: Uint8Array): Observable<Uint8Array>;
  bidirectionalStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Observable<Uint8Array>;
}
