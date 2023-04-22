// package: 
// file: backend/resources.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Contest extends jspb.Message {
  getYear(): number;
  setYear(value: number): void;

  hasQualifierStartAt(): boolean;
  clearQualifierStartAt(): void;
  getQualifierStartAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setQualifierStartAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasQualifierEndAt(): boolean;
  clearQualifierEndAt(): void;
  getQualifierEndAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setQualifierEndAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getQualifierSubmitLimit(): number;
  setQualifierSubmitLimit(value: number): void;

  hasFinalStartAt(): boolean;
  clearFinalStartAt(): void;
  getFinalStartAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFinalStartAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasFinalEndAt(): boolean;
  clearFinalEndAt(): void;
  getFinalEndAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFinalEndAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getFinalSubmitLimit(): number;
  setFinalSubmitLimit(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Contest.AsObject;
  static toObject(includeInstance: boolean, msg: Contest): Contest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Contest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Contest;
  static deserializeBinaryFromReader(message: Contest, reader: jspb.BinaryReader): Contest;
}

export namespace Contest {
  export type AsObject = {
    year: number,
    qualifierStartAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    qualifierEndAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    qualifierSubmitLimit: number,
    finalStartAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    finalEndAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    finalSubmitLimit: number,
  }
}

export class Group extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getYear(): number;
  setYear(value: number): void;

  getRole(): RoleMap[keyof RoleMap];
  setRole(value: RoleMap[keyof RoleMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Group.AsObject;
  static toObject(includeInstance: boolean, msg: Group): Group.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Group, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Group;
  static deserializeBinaryFromReader(message: Group, reader: jspb.BinaryReader): Group;
}

export namespace Group {
  export type AsObject = {
    id: string,
    year: number,
    role: RoleMap[keyof RoleMap],
  }
}

export class Submit extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getGroupId(): string;
  setGroupId(value: string): void;

  getYear(): number;
  setYear(value: number): void;

  getScore(): number;
  setScore(value: number): void;

  getLanguage(): LanguageMap[keyof LanguageMap];
  setLanguage(value: LanguageMap[keyof LanguageMap]): void;

  hasSubmitedAt(): boolean;
  clearSubmitedAt(): void;
  getSubmitedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setSubmitedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  clearTagResultsList(): void;
  getTagResultsList(): Array<TagResult>;
  setTagResultsList(value: Array<TagResult>): void;
  addTagResults(value?: TagResult, index?: number): TagResult;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Submit.AsObject;
  static toObject(includeInstance: boolean, msg: Submit): Submit.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Submit, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Submit;
  static deserializeBinaryFromReader(message: Submit, reader: jspb.BinaryReader): Submit;
}

export namespace Submit {
  export type AsObject = {
    id: string,
    groupId: string,
    year: number,
    score: number,
    language: LanguageMap[keyof LanguageMap],
    submitedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    tagResultsList: Array<TagResult.AsObject>,
  }
}

export class TagResult extends jspb.Message {
  getSubmitId(): string;
  setSubmitId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getScore(): number;
  setScore(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TagResult.AsObject;
  static toObject(includeInstance: boolean, msg: TagResult): TagResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TagResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TagResult;
  static deserializeBinaryFromReader(message: TagResult, reader: jspb.BinaryReader): TagResult;
}

export namespace TagResult {
  export type AsObject = {
    submitId: string,
    name: string,
    score: number,
  }
}

export interface LanguageMap {
  PHP: 0;
  GO: 1;
  RUST: 2;
  JAVASCRIPT: 3;
  CSHARP: 4;
  CPP: 5;
  RUBY: 6;
  PYTHON: 7;
}

export const Language: LanguageMap;

export interface RoleMap {
  CONTESTANT: 0;
  GUEST: 1;
}

export const Role: RoleMap;

