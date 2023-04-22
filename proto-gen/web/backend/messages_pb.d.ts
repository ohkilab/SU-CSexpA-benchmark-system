import * as jspb from 'google-protobuf'

import * as backend_resources_pb from '../backend/resources_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class PostLoginRequest extends jspb.Message {
  getId(): string;
  setId(value: string): PostLoginRequest;

  getPassword(): string;
  setPassword(value: string): PostLoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PostLoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PostLoginRequest): PostLoginRequest.AsObject;
  static serializeBinaryToWriter(message: PostLoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PostLoginRequest;
  static deserializeBinaryFromReader(message: PostLoginRequest, reader: jspb.BinaryReader): PostLoginRequest;
}

export namespace PostLoginRequest {
  export type AsObject = {
    id: string,
    password: string,
  }
}

export class PostLoginResponse extends jspb.Message {
  getGroup(): backend_resources_pb.Group | undefined;
  setGroup(value?: backend_resources_pb.Group): PostLoginResponse;
  hasGroup(): boolean;
  clearGroup(): PostLoginResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PostLoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PostLoginResponse): PostLoginResponse.AsObject;
  static serializeBinaryToWriter(message: PostLoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PostLoginResponse;
  static deserializeBinaryFromReader(message: PostLoginResponse, reader: jspb.BinaryReader): PostLoginResponse;
}

export namespace PostLoginResponse {
  export type AsObject = {
    group?: backend_resources_pb.Group.AsObject,
  }
}

export class PostSubmitRequest extends jspb.Message {
  getIpAddr(): string;
  setIpAddr(value: string): PostSubmitRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PostSubmitRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PostSubmitRequest): PostSubmitRequest.AsObject;
  static serializeBinaryToWriter(message: PostSubmitRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PostSubmitRequest;
  static deserializeBinaryFromReader(message: PostSubmitRequest, reader: jspb.BinaryReader): PostSubmitRequest;
}

export namespace PostSubmitRequest {
  export type AsObject = {
    ipAddr: string,
  }
}

export class PostSubmitResponse extends jspb.Message {
  getId(): string;
  setId(value: string): PostSubmitResponse;

  getIpAddr(): string;
  setIpAddr(value: string): PostSubmitResponse;

  getSubmitedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setSubmitedAt(value?: google_protobuf_timestamp_pb.Timestamp): PostSubmitResponse;
  hasSubmitedAt(): boolean;
  clearSubmitedAt(): PostSubmitResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PostSubmitResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PostSubmitResponse): PostSubmitResponse.AsObject;
  static serializeBinaryToWriter(message: PostSubmitResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PostSubmitResponse;
  static deserializeBinaryFromReader(message: PostSubmitResponse, reader: jspb.BinaryReader): PostSubmitResponse;
}

export namespace PostSubmitResponse {
  export type AsObject = {
    id: string,
    ipAddr: string,
    submitedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class GetSubmitRequest extends jspb.Message {
  getSubmitId(): string;
  setSubmitId(value: string): GetSubmitRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSubmitRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetSubmitRequest): GetSubmitRequest.AsObject;
  static serializeBinaryToWriter(message: GetSubmitRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSubmitRequest;
  static deserializeBinaryFromReader(message: GetSubmitRequest, reader: jspb.BinaryReader): GetSubmitRequest;
}

export namespace GetSubmitRequest {
  export type AsObject = {
    submitId: string,
  }
}

export class GetSubmitResponse extends jspb.Message {
  getTagProgress(): GetSubmitResponse.TagProgress | undefined;
  setTagProgress(value?: GetSubmitResponse.TagProgress): GetSubmitResponse;
  hasTagProgress(): boolean;
  clearTagProgress(): GetSubmitResponse;

  getSubmit(): backend_resources_pb.Submit | undefined;
  setSubmit(value?: backend_resources_pb.Submit): GetSubmitResponse;
  hasSubmit(): boolean;
  clearSubmit(): GetSubmitResponse;

  getResultCase(): GetSubmitResponse.ResultCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSubmitResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetSubmitResponse): GetSubmitResponse.AsObject;
  static serializeBinaryToWriter(message: GetSubmitResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSubmitResponse;
  static deserializeBinaryFromReader(message: GetSubmitResponse, reader: jspb.BinaryReader): GetSubmitResponse;
}

export namespace GetSubmitResponse {
  export type AsObject = {
    tagProgress?: GetSubmitResponse.TagProgress.AsObject,
    submit?: backend_resources_pb.Submit.AsObject,
  }

  export class TagProgress extends jspb.Message {
    getSubmitId(): string;
    setSubmitId(value: string): TagProgress;

    getScore(): number;
    setScore(value: number): TagProgress;

    getName(): string;
    setName(value: string): TagProgress;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TagProgress.AsObject;
    static toObject(includeInstance: boolean, msg: TagProgress): TagProgress.AsObject;
    static serializeBinaryToWriter(message: TagProgress, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TagProgress;
    static deserializeBinaryFromReader(message: TagProgress, reader: jspb.BinaryReader): TagProgress;
  }

  export namespace TagProgress {
    export type AsObject = {
      submitId: string,
      score: number,
      name: string,
    }
  }


  export enum ResultCase { 
    RESULT_NOT_SET = 0,
    TAG_PROGRESS = 1,
    SUBMIT = 2,
  }
}

export class GetRankingRequest extends jspb.Message {
  getYear(): number;
  setYear(value: number): GetRankingRequest;

  getContainGuest(): boolean;
  setContainGuest(value: boolean): GetRankingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRankingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRankingRequest): GetRankingRequest.AsObject;
  static serializeBinaryToWriter(message: GetRankingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRankingRequest;
  static deserializeBinaryFromReader(message: GetRankingRequest, reader: jspb.BinaryReader): GetRankingRequest;
}

export namespace GetRankingRequest {
  export type AsObject = {
    year: number,
    containGuest: boolean,
  }
}

export class GetRankingResponse extends jspb.Message {
  getRecordsList(): Array<GetRankingResponse.Record>;
  setRecordsList(value: Array<GetRankingResponse.Record>): GetRankingResponse;
  clearRecordsList(): GetRankingResponse;
  addRecords(value?: GetRankingResponse.Record, index?: number): GetRankingResponse.Record;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRankingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetRankingResponse): GetRankingResponse.AsObject;
  static serializeBinaryToWriter(message: GetRankingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRankingResponse;
  static deserializeBinaryFromReader(message: GetRankingResponse, reader: jspb.BinaryReader): GetRankingResponse;
}

export namespace GetRankingResponse {
  export type AsObject = {
    recordsList: Array<GetRankingResponse.Record.AsObject>,
  }

  export class Record extends jspb.Message {
    getRank(): number;
    setRank(value: number): Record;

    getGroup(): backend_resources_pb.Group | undefined;
    setGroup(value?: backend_resources_pb.Group): Record;
    hasGroup(): boolean;
    clearGroup(): Record;

    getScore(): number;
    setScore(value: number): Record;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Record.AsObject;
    static toObject(includeInstance: boolean, msg: Record): Record.AsObject;
    static serializeBinaryToWriter(message: Record, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Record;
    static deserializeBinaryFromReader(message: Record, reader: jspb.BinaryReader): Record;
  }

  export namespace Record {
    export type AsObject = {
      rank: number,
      group?: backend_resources_pb.Group.AsObject,
      score: number,
    }
  }

}

export class GetGroupRequest extends jspb.Message {
  getGroupId(): string;
  setGroupId(value: string): GetGroupRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetGroupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetGroupRequest): GetGroupRequest.AsObject;
  static serializeBinaryToWriter(message: GetGroupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetGroupRequest;
  static deserializeBinaryFromReader(message: GetGroupRequest, reader: jspb.BinaryReader): GetGroupRequest;
}

export namespace GetGroupRequest {
  export type AsObject = {
    groupId: string,
  }
}

export class GetGroupResponse extends jspb.Message {
  getGroupsList(): Array<GetGroupResponse.GroupInfo>;
  setGroupsList(value: Array<GetGroupResponse.GroupInfo>): GetGroupResponse;
  clearGroupsList(): GetGroupResponse;
  addGroups(value?: GetGroupResponse.GroupInfo, index?: number): GetGroupResponse.GroupInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetGroupResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetGroupResponse): GetGroupResponse.AsObject;
  static serializeBinaryToWriter(message: GetGroupResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetGroupResponse;
  static deserializeBinaryFromReader(message: GetGroupResponse, reader: jspb.BinaryReader): GetGroupResponse;
}

export namespace GetGroupResponse {
  export type AsObject = {
    groupsList: Array<GetGroupResponse.GroupInfo.AsObject>,
  }

  export class GroupInfo extends jspb.Message {
    getGroup(): backend_resources_pb.Group | undefined;
    setGroup(value?: backend_resources_pb.Group): GroupInfo;
    hasGroup(): boolean;
    clearGroup(): GroupInfo;

    getSubmitsList(): Array<backend_resources_pb.Submit>;
    setSubmitsList(value: Array<backend_resources_pb.Submit>): GroupInfo;
    clearSubmitsList(): GroupInfo;
    addSubmits(value?: backend_resources_pb.Submit, index?: number): backend_resources_pb.Submit;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GroupInfo.AsObject;
    static toObject(includeInstance: boolean, msg: GroupInfo): GroupInfo.AsObject;
    static serializeBinaryToWriter(message: GroupInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GroupInfo;
    static deserializeBinaryFromReader(message: GroupInfo, reader: jspb.BinaryReader): GroupInfo;
  }

  export namespace GroupInfo {
    export type AsObject = {
      group?: backend_resources_pb.Group.AsObject,
      submitsList: Array<backend_resources_pb.Submit.AsObject>,
    }
  }

}

export class PingUnaryRequest extends jspb.Message {
  getPing(): string;
  setPing(value: string): PingUnaryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingUnaryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PingUnaryRequest): PingUnaryRequest.AsObject;
  static serializeBinaryToWriter(message: PingUnaryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingUnaryRequest;
  static deserializeBinaryFromReader(message: PingUnaryRequest, reader: jspb.BinaryReader): PingUnaryRequest;
}

export namespace PingUnaryRequest {
  export type AsObject = {
    ping: string,
  }
}

export class PingUnaryResponse extends jspb.Message {
  getPong(): string;
  setPong(value: string): PingUnaryResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingUnaryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PingUnaryResponse): PingUnaryResponse.AsObject;
  static serializeBinaryToWriter(message: PingUnaryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingUnaryResponse;
  static deserializeBinaryFromReader(message: PingUnaryResponse, reader: jspb.BinaryReader): PingUnaryResponse;
}

export namespace PingUnaryResponse {
  export type AsObject = {
    pong: string,
  }
}

export class PingServerSideStreamingRequest extends jspb.Message {
  getPing(): string;
  setPing(value: string): PingServerSideStreamingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingServerSideStreamingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PingServerSideStreamingRequest): PingServerSideStreamingRequest.AsObject;
  static serializeBinaryToWriter(message: PingServerSideStreamingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingServerSideStreamingRequest;
  static deserializeBinaryFromReader(message: PingServerSideStreamingRequest, reader: jspb.BinaryReader): PingServerSideStreamingRequest;
}

export namespace PingServerSideStreamingRequest {
  export type AsObject = {
    ping: string,
  }
}

export class PingServerSideStreamingResponse extends jspb.Message {
  getPong(): string;
  setPong(value: string): PingServerSideStreamingResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingServerSideStreamingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PingServerSideStreamingResponse): PingServerSideStreamingResponse.AsObject;
  static serializeBinaryToWriter(message: PingServerSideStreamingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingServerSideStreamingResponse;
  static deserializeBinaryFromReader(message: PingServerSideStreamingResponse, reader: jspb.BinaryReader): PingServerSideStreamingResponse;
}

export namespace PingServerSideStreamingResponse {
  export type AsObject = {
    pong: string,
  }
}

