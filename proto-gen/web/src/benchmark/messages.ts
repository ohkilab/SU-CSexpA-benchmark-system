/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Status, statusFromJSON, statusToJSON } from "../backend/resources";

export const protobufPackage = "benchmark";

export enum HttpMethod {
  GET = 0,
  POST = 1,
  PUT = 2,
  DELETE = 3,
  UNRECOGNIZED = -1,
}

export function httpMethodFromJSON(object: any): HttpMethod {
  switch (object) {
    case 0:
    case "GET":
      return HttpMethod.GET;
    case 1:
    case "POST":
      return HttpMethod.POST;
    case 2:
    case "PUT":
      return HttpMethod.PUT;
    case 3:
    case "DELETE":
      return HttpMethod.DELETE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return HttpMethod.UNRECOGNIZED;
  }
}

export function httpMethodToJSON(object: HttpMethod): string {
  switch (object) {
    case HttpMethod.GET:
      return "GET";
    case HttpMethod.POST:
      return "POST";
    case HttpMethod.PUT:
      return "PUT";
    case HttpMethod.DELETE:
      return "DELETE";
    case HttpMethod.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface ExecuteRequest {
  tasks: Task[];
  /** for logging */
  groupId: string;
  contestSlug: string;
}

export interface Task {
  request:
    | HttpRequest
    | undefined;
  /** the number of threads for a task */
  threadNum: number;
  /** the count of attempting for a task */
  attemptCount: number;
}

export interface ExecuteResponse {
  ok: boolean;
  /** if ok is false, this field is set */
  errorMessage?:
    | string
    | undefined;
  /** in milliseconds */
  timeElapsed: number;
  totalRequests: number;
  requestsPerSecond: number;
  task: Task | undefined;
  status: Status;
}

export interface HttpRequest {
  /** e.g.) http://10.255.255.255/endpoint */
  url: string;
  method: HttpMethod;
  contentType: string;
  body: string;
}

function createBaseExecuteRequest(): ExecuteRequest {
  return { tasks: [], groupId: "", contestSlug: "" };
}

export const ExecuteRequest = {
  encode(message: ExecuteRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tasks) {
      Task.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.groupId !== "") {
      writer.uint32(18).string(message.groupId);
    }
    if (message.contestSlug !== "") {
      writer.uint32(26).string(message.contestSlug);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ExecuteRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExecuteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.tasks.push(Task.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.groupId = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.contestSlug = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ExecuteRequest {
    return {
      tasks: Array.isArray(object?.tasks) ? object.tasks.map((e: any) => Task.fromJSON(e)) : [],
      groupId: isSet(object.groupId) ? String(object.groupId) : "",
      contestSlug: isSet(object.contestSlug) ? String(object.contestSlug) : "",
    };
  },

  toJSON(message: ExecuteRequest): unknown {
    const obj: any = {};
    if (message.tasks) {
      obj.tasks = message.tasks.map((e) => e ? Task.toJSON(e) : undefined);
    } else {
      obj.tasks = [];
    }
    message.groupId !== undefined && (obj.groupId = message.groupId);
    message.contestSlug !== undefined && (obj.contestSlug = message.contestSlug);
    return obj;
  },

  create<I extends Exact<DeepPartial<ExecuteRequest>, I>>(base?: I): ExecuteRequest {
    return ExecuteRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ExecuteRequest>, I>>(object: I): ExecuteRequest {
    const message = createBaseExecuteRequest();
    message.tasks = object.tasks?.map((e) => Task.fromPartial(e)) || [];
    message.groupId = object.groupId ?? "";
    message.contestSlug = object.contestSlug ?? "";
    return message;
  },
};

function createBaseTask(): Task {
  return { request: undefined, threadNum: 0, attemptCount: 0 };
}

export const Task = {
  encode(message: Task, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.request !== undefined) {
      HttpRequest.encode(message.request, writer.uint32(10).fork()).ldelim();
    }
    if (message.threadNum !== 0) {
      writer.uint32(48).int32(message.threadNum);
    }
    if (message.attemptCount !== 0) {
      writer.uint32(56).int32(message.attemptCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Task {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTask();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.request = HttpRequest.decode(reader, reader.uint32());
          continue;
        case 6:
          if (tag != 48) {
            break;
          }

          message.threadNum = reader.int32();
          continue;
        case 7:
          if (tag != 56) {
            break;
          }

          message.attemptCount = reader.int32();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Task {
    return {
      request: isSet(object.request) ? HttpRequest.fromJSON(object.request) : undefined,
      threadNum: isSet(object.threadNum) ? Number(object.threadNum) : 0,
      attemptCount: isSet(object.attemptCount) ? Number(object.attemptCount) : 0,
    };
  },

  toJSON(message: Task): unknown {
    const obj: any = {};
    message.request !== undefined && (obj.request = message.request ? HttpRequest.toJSON(message.request) : undefined);
    message.threadNum !== undefined && (obj.threadNum = Math.round(message.threadNum));
    message.attemptCount !== undefined && (obj.attemptCount = Math.round(message.attemptCount));
    return obj;
  },

  create<I extends Exact<DeepPartial<Task>, I>>(base?: I): Task {
    return Task.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Task>, I>>(object: I): Task {
    const message = createBaseTask();
    message.request = (object.request !== undefined && object.request !== null)
      ? HttpRequest.fromPartial(object.request)
      : undefined;
    message.threadNum = object.threadNum ?? 0;
    message.attemptCount = object.attemptCount ?? 0;
    return message;
  },
};

function createBaseExecuteResponse(): ExecuteResponse {
  return {
    ok: false,
    errorMessage: undefined,
    timeElapsed: 0,
    totalRequests: 0,
    requestsPerSecond: 0,
    task: undefined,
    status: 0,
  };
}

export const ExecuteResponse = {
  encode(message: ExecuteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ok === true) {
      writer.uint32(8).bool(message.ok);
    }
    if (message.errorMessage !== undefined) {
      writer.uint32(18).string(message.errorMessage);
    }
    if (message.timeElapsed !== 0) {
      writer.uint32(24).int64(message.timeElapsed);
    }
    if (message.totalRequests !== 0) {
      writer.uint32(32).int32(message.totalRequests);
    }
    if (message.requestsPerSecond !== 0) {
      writer.uint32(40).int32(message.requestsPerSecond);
    }
    if (message.task !== undefined) {
      Task.encode(message.task, writer.uint32(50).fork()).ldelim();
    }
    if (message.status !== 0) {
      writer.uint32(56).int32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ExecuteResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExecuteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.ok = reader.bool();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.errorMessage = reader.string();
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.timeElapsed = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag != 32) {
            break;
          }

          message.totalRequests = reader.int32();
          continue;
        case 5:
          if (tag != 40) {
            break;
          }

          message.requestsPerSecond = reader.int32();
          continue;
        case 6:
          if (tag != 50) {
            break;
          }

          message.task = Task.decode(reader, reader.uint32());
          continue;
        case 7:
          if (tag != 56) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ExecuteResponse {
    return {
      ok: isSet(object.ok) ? Boolean(object.ok) : false,
      errorMessage: isSet(object.errorMessage) ? String(object.errorMessage) : undefined,
      timeElapsed: isSet(object.timeElapsed) ? Number(object.timeElapsed) : 0,
      totalRequests: isSet(object.totalRequests) ? Number(object.totalRequests) : 0,
      requestsPerSecond: isSet(object.requestsPerSecond) ? Number(object.requestsPerSecond) : 0,
      task: isSet(object.task) ? Task.fromJSON(object.task) : undefined,
      status: isSet(object.status) ? statusFromJSON(object.status) : 0,
    };
  },

  toJSON(message: ExecuteResponse): unknown {
    const obj: any = {};
    message.ok !== undefined && (obj.ok = message.ok);
    message.errorMessage !== undefined && (obj.errorMessage = message.errorMessage);
    message.timeElapsed !== undefined && (obj.timeElapsed = Math.round(message.timeElapsed));
    message.totalRequests !== undefined && (obj.totalRequests = Math.round(message.totalRequests));
    message.requestsPerSecond !== undefined && (obj.requestsPerSecond = Math.round(message.requestsPerSecond));
    message.task !== undefined && (obj.task = message.task ? Task.toJSON(message.task) : undefined);
    message.status !== undefined && (obj.status = statusToJSON(message.status));
    return obj;
  },

  create<I extends Exact<DeepPartial<ExecuteResponse>, I>>(base?: I): ExecuteResponse {
    return ExecuteResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ExecuteResponse>, I>>(object: I): ExecuteResponse {
    const message = createBaseExecuteResponse();
    message.ok = object.ok ?? false;
    message.errorMessage = object.errorMessage ?? undefined;
    message.timeElapsed = object.timeElapsed ?? 0;
    message.totalRequests = object.totalRequests ?? 0;
    message.requestsPerSecond = object.requestsPerSecond ?? 0;
    message.task = (object.task !== undefined && object.task !== null) ? Task.fromPartial(object.task) : undefined;
    message.status = object.status ?? 0;
    return message;
  },
};

function createBaseHttpRequest(): HttpRequest {
  return { url: "", method: 0, contentType: "", body: "" };
}

export const HttpRequest = {
  encode(message: HttpRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.method !== 0) {
      writer.uint32(16).int32(message.method);
    }
    if (message.contentType !== "") {
      writer.uint32(34).string(message.contentType);
    }
    if (message.body !== "") {
      writer.uint32(42).string(message.body);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HttpRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHttpRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.url = reader.string();
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.method = reader.int32() as any;
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.contentType = reader.string();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.body = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): HttpRequest {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      method: isSet(object.method) ? httpMethodFromJSON(object.method) : 0,
      contentType: isSet(object.contentType) ? String(object.contentType) : "",
      body: isSet(object.body) ? String(object.body) : "",
    };
  },

  toJSON(message: HttpRequest): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    message.method !== undefined && (obj.method = httpMethodToJSON(message.method));
    message.contentType !== undefined && (obj.contentType = message.contentType);
    message.body !== undefined && (obj.body = message.body);
    return obj;
  },

  create<I extends Exact<DeepPartial<HttpRequest>, I>>(base?: I): HttpRequest {
    return HttpRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<HttpRequest>, I>>(object: I): HttpRequest {
    const message = createBaseHttpRequest();
    message.url = object.url ?? "";
    message.method = object.method ?? 0;
    message.contentType = object.contentType ?? "";
    message.body = object.body ?? "";
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
