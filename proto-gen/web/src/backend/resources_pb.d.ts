import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Contest extends jspb.Message {
  getYear(): number;
  setYear(value: number): Contest;

  getQualifierStartAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setQualifierStartAt(value?: google_protobuf_timestamp_pb.Timestamp): Contest;
  hasQualifierStartAt(): boolean;
  clearQualifierStartAt(): Contest;

  getQualifierEndAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setQualifierEndAt(value?: google_protobuf_timestamp_pb.Timestamp): Contest;
  hasQualifierEndAt(): boolean;
  clearQualifierEndAt(): Contest;

  getQualifierSubmitLimit(): number;
  setQualifierSubmitLimit(value: number): Contest;

  getFinalStartAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFinalStartAt(value?: google_protobuf_timestamp_pb.Timestamp): Contest;
  hasFinalStartAt(): boolean;
  clearFinalStartAt(): Contest;

  getFinalEndAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFinalEndAt(value?: google_protobuf_timestamp_pb.Timestamp): Contest;
  hasFinalEndAt(): boolean;
  clearFinalEndAt(): Contest;

  getFinalSubmitLimit(): number;
  setFinalSubmitLimit(value: number): Contest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Contest.AsObject;
  static toObject(includeInstance: boolean, msg: Contest): Contest.AsObject;
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
  setId(value: string): Group;

  getYear(): number;
  setYear(value: number): Group;

  getRole(): Role;
  setRole(value: Role): Group;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Group.AsObject;
  static toObject(includeInstance: boolean, msg: Group): Group.AsObject;
  static serializeBinaryToWriter(message: Group, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Group;
  static deserializeBinaryFromReader(message: Group, reader: jspb.BinaryReader): Group;
}

export namespace Group {
  export type AsObject = {
    id: string,
    year: number,
    role: Role,
  }
}

export class Submit extends jspb.Message {
  getId(): string;
  setId(value: string): Submit;

  getGroupId(): string;
  setGroupId(value: string): Submit;

  getYear(): number;
  setYear(value: number): Submit;

  getScore(): number;
  setScore(value: number): Submit;

  getLanguage(): Language;
  setLanguage(value: Language): Submit;

  getSubmitedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setSubmitedAt(value?: google_protobuf_timestamp_pb.Timestamp): Submit;
  hasSubmitedAt(): boolean;
  clearSubmitedAt(): Submit;

  getTagResultsList(): Array<TagResult>;
  setTagResultsList(value: Array<TagResult>): Submit;
  clearTagResultsList(): Submit;
  addTagResults(value?: TagResult, index?: number): TagResult;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Submit.AsObject;
  static toObject(includeInstance: boolean, msg: Submit): Submit.AsObject;
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
    language: Language,
    submitedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    tagResultsList: Array<TagResult.AsObject>,
  }
}

export class TagResult extends jspb.Message {
  getSubmitId(): string;
  setSubmitId(value: string): TagResult;

  getName(): string;
  setName(value: string): TagResult;

  getScore(): number;
  setScore(value: number): TagResult;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TagResult.AsObject;
  static toObject(includeInstance: boolean, msg: TagResult): TagResult.AsObject;
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

export enum Language { 
  PHP = 0,
  GO = 1,
  RUST = 2,
  JAVASCRIPT = 3,
  CSHARP = 4,
  CPP = 5,
  RUBY = 6,
  PYTHON = 7,
}
export enum Role { 
  CONTESTANT = 0,
  GUEST = 1,
}
