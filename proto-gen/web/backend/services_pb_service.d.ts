// package: 
// file: backend/services.proto

import * as backend_services_pb from "../backend/services_pb";
import * as backend_messages_pb from "../backend/messages_pb";
import {grpc} from "@improbable-eng/grpc-web";

type BackendServiceGetRanking = {
  readonly methodName: string;
  readonly service: typeof BackendService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof backend_messages_pb.GetRankingRequest;
  readonly responseType: typeof backend_messages_pb.GetRankingResponse;
};

type BackendServicePostSubmit = {
  readonly methodName: string;
  readonly service: typeof BackendService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof backend_messages_pb.PostSubmitRequest;
  readonly responseType: typeof backend_messages_pb.PostSubmitResponse;
};

type BackendServiceGetSubmit = {
  readonly methodName: string;
  readonly service: typeof BackendService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof backend_messages_pb.GetSubmitRequest;
  readonly responseType: typeof backend_messages_pb.GetSubmitResponse;
};

type BackendServicePostLogin = {
  readonly methodName: string;
  readonly service: typeof BackendService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof backend_messages_pb.PostLoginRequest;
  readonly responseType: typeof backend_messages_pb.PostLoginResponse;
};

export class BackendService {
  static readonly serviceName: string;
  static readonly GetRanking: BackendServiceGetRanking;
  static readonly PostSubmit: BackendServicePostSubmit;
  static readonly GetSubmit: BackendServiceGetSubmit;
  static readonly PostLogin: BackendServicePostLogin;
}

type HealthcheckServicePingUnary = {
  readonly methodName: string;
  readonly service: typeof HealthcheckService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof backend_messages_pb.PingUnaryRequest;
  readonly responseType: typeof backend_messages_pb.PingUnaryResponse;
};

type HealthcheckServicePingServerSideStreaming = {
  readonly methodName: string;
  readonly service: typeof HealthcheckService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof backend_messages_pb.PingServerSideStreamingRequest;
  readonly responseType: typeof backend_messages_pb.PingServerSideStreamingResponse;
};

export class HealthcheckService {
  static readonly serviceName: string;
  static readonly PingUnary: HealthcheckServicePingUnary;
  static readonly PingServerSideStreaming: HealthcheckServicePingServerSideStreaming;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class BackendServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getRanking(
    requestMessage: backend_messages_pb.GetRankingRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: backend_messages_pb.GetRankingResponse|null) => void
  ): UnaryResponse;
  getRanking(
    requestMessage: backend_messages_pb.GetRankingRequest,
    callback: (error: ServiceError|null, responseMessage: backend_messages_pb.GetRankingResponse|null) => void
  ): UnaryResponse;
  postSubmit(
    requestMessage: backend_messages_pb.PostSubmitRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: backend_messages_pb.PostSubmitResponse|null) => void
  ): UnaryResponse;
  postSubmit(
    requestMessage: backend_messages_pb.PostSubmitRequest,
    callback: (error: ServiceError|null, responseMessage: backend_messages_pb.PostSubmitResponse|null) => void
  ): UnaryResponse;
  getSubmit(requestMessage: backend_messages_pb.GetSubmitRequest, metadata?: grpc.Metadata): ResponseStream<backend_messages_pb.GetSubmitResponse>;
  postLogin(
    requestMessage: backend_messages_pb.PostLoginRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: backend_messages_pb.PostLoginResponse|null) => void
  ): UnaryResponse;
  postLogin(
    requestMessage: backend_messages_pb.PostLoginRequest,
    callback: (error: ServiceError|null, responseMessage: backend_messages_pb.PostLoginResponse|null) => void
  ): UnaryResponse;
}

export class HealthcheckServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  pingUnary(
    requestMessage: backend_messages_pb.PingUnaryRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: backend_messages_pb.PingUnaryResponse|null) => void
  ): UnaryResponse;
  pingUnary(
    requestMessage: backend_messages_pb.PingUnaryRequest,
    callback: (error: ServiceError|null, responseMessage: backend_messages_pb.PingUnaryResponse|null) => void
  ): UnaryResponse;
  pingServerSideStreaming(requestMessage: backend_messages_pb.PingServerSideStreamingRequest, metadata?: grpc.Metadata): ResponseStream<backend_messages_pb.PingServerSideStreamingResponse>;
}

