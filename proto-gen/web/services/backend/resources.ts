// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "services/backend/resources.proto" (package "backend", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Timestamp } from "../../google/protobuf/timestamp";
/**
 * @generated from protobuf message backend.Contest
 */
export interface Contest {
    /**
     * @generated from protobuf field: int32 id = 1;
     */
    id: number;
    /**
     * @generated from protobuf field: string title = 2;
     */
    title: string;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp start_at = 4;
     */
    startAt?: Timestamp;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp end_at = 5;
     */
    endAt?: Timestamp;
    /**
     * @generated from protobuf field: int32 submit_limit = 6;
     */
    submitLimit: number;
    /**
     * @generated from protobuf field: string slug = 8;
     */
    slug: string;
    /**
     * @generated from protobuf field: backend.TagSelectionLogicType tag_selection_logic = 9;
     */
    tagSelectionLogic: TagSelectionLogicType;
    /**
     * @generated from protobuf field: backend.Validator validator = 10;
     */
    validator: Validator;
}
/**
 * @generated from protobuf message backend.TagSelectionLogicManual
 */
export interface TagSelectionLogicManual {
    /**
     * @generated from protobuf field: backend.TagSelectionLogicType type = 1;
     */
    type: TagSelectionLogicType;
    /**
     * @generated from protobuf field: repeated backend.Tags tags_list = 2;
     */
    tagsList: Tags[]; // tags_list[i] .. used if the attempt count is i+1
}
/**
 * @generated from protobuf message backend.TagSelectionLogicAuto
 */
export interface TagSelectionLogicAuto {
    /**
     * @generated from protobuf field: backend.TagSelectionLogicType type = 1;
     */
    type: TagSelectionLogicType;
    /**
     * @generated from protobuf field: backend.Tags tags = 2;
     */
    tags?: Tags;
}
/**
 * @generated from protobuf message backend.Tags
 */
export interface Tags {
    /**
     * @generated from protobuf field: repeated string tags = 1;
     */
    tags: string[];
}
/**
 * @generated from protobuf message backend.Group
 */
export interface Group {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: backend.Role role = 4;
     */
    role: Role;
}
/**
 * @generated from protobuf message backend.Submit
 */
export interface Submit {
    /**
     * @generated from protobuf field: int32 id = 1;
     */
    id: number;
    /**
     * @generated from protobuf field: string group_name = 2;
     */
    groupName: string;
    /**
     * @generated from protobuf field: int32 score = 4;
     */
    score: number;
    /**
     * @generated from protobuf field: backend.Language language = 5;
     */
    language: Language;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp submited_at = 6;
     */
    submitedAt?: Timestamp;
    /**
     * @generated from protobuf field: optional google.protobuf.Timestamp completed_at = 7;
     */
    completedAt?: Timestamp; // it this field is not null, this submit is completed
    /**
     * @generated from protobuf field: repeated backend.TaskResult task_results = 8;
     */
    taskResults: TaskResult[];
    /**
     * @generated from protobuf field: backend.Status status = 9;
     */
    status: Status;
    /**
     * @generated from protobuf field: optional string error_message = 10;
     */
    errorMessage?: string; // if the connection error occurs, then this field is filled
    /**
     * @generated from protobuf field: int32 tag_count = 11;
     */
    tagCount: number;
}
/**
 * @generated from protobuf message backend.TaskResult
 */
export interface TaskResult {
    /**
     * @generated from protobuf field: int32 id = 1;
     */
    id: number;
    /**
     * @generated from protobuf field: int32 request_per_sec = 2;
     */
    requestPerSec: number;
    /**
     * @generated from protobuf field: string url = 3;
     */
    url: string;
    /**
     * @generated from protobuf field: string method = 4;
     */
    method: string;
    /**
     * @generated from protobuf field: string request_content_type = 5;
     */
    requestContentType: string;
    /**
     * @generated from protobuf field: optional string request_body = 6;
     */
    requestBody?: string;
    /**
     * @generated from protobuf field: string response_code = 7;
     */
    responseCode: string;
    /**
     * @generated from protobuf field: string response_content_type = 8;
     */
    responseContentType: string;
    /**
     * @generated from protobuf field: string response_body = 9;
     */
    responseBody: string;
    /**
     * @generated from protobuf field: int32 thread_num = 10;
     */
    threadNum: number;
    /**
     * @generated from protobuf field: int32 attempt_count = 11;
     */
    attemptCount: number;
    /**
     * @generated from protobuf field: int32 attempt_time = 12;
     */
    attemptTime: number;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp created_at = 13;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional google.protobuf.Timestamp deleted_at = 14;
     */
    deletedAt?: Timestamp;
    /**
     * @generated from protobuf field: optional string error_message = 15;
     */
    errorMessage?: string;
    /**
     * @generated from protobuf field: backend.Status status = 16;
     */
    status: Status;
}
/**
 * @generated from protobuf enum backend.TagSelectionLogicType
 */
export enum TagSelectionLogicType {
    /**
     * @generated from protobuf enum value: AUTO = 0;
     */
    AUTO = 0,
    /**
     * @generated from protobuf enum value: MANUAL = 1;
     */
    MANUAL = 1
}
/**
 * @generated from protobuf enum backend.Status
 */
export enum Status {
    /**
     * waiting for benchmark
     *
     * @generated from protobuf enum value: WAITING = 0;
     */
    WAITING = 0,
    /**
     * in progress
     *
     * @generated from protobuf enum value: IN_PROGRESS = 1;
     */
    IN_PROGRESS = 1,
    /**
     * benchmark succeeded
     *
     * @generated from protobuf enum value: SUCCESS = 2;
     */
    SUCCESS = 2,
    /**
     * failed to connect
     *
     * @generated from protobuf enum value: CONNECTION_FAILED = 3;
     */
    CONNECTION_FAILED = 3,
    /**
     * validation error
     *
     * @generated from protobuf enum value: VALIDATION_ERROR = 4;
     */
    VALIDATION_ERROR = 4,
    /**
     * backend error
     *
     * @generated from protobuf enum value: INTERNAL_ERROR = 5;
     */
    INTERNAL_ERROR = 5,
    /**
     * timeout
     *
     * @generated from protobuf enum value: TIMEOUT = 6;
     */
    TIMEOUT = 6
}
/**
 * @generated from protobuf enum backend.Language
 */
export enum Language {
    /**
     * @generated from protobuf enum value: PHP = 0;
     */
    PHP = 0,
    /**
     * @generated from protobuf enum value: GO = 1;
     */
    GO = 1,
    /**
     * @generated from protobuf enum value: RUST = 2;
     */
    RUST = 2,
    /**
     * @generated from protobuf enum value: JAVASCRIPT = 3;
     */
    JAVASCRIPT = 3,
    /**
     * @generated from protobuf enum value: CSHARP = 4;
     */
    CSHARP = 4,
    /**
     * @generated from protobuf enum value: CPP = 5;
     */
    CPP = 5,
    /**
     * @generated from protobuf enum value: RUBY = 6;
     */
    RUBY = 6,
    /**
     * @generated from protobuf enum value: PYTHON = 7;
     */
    PYTHON = 7
}
/**
 * @generated from protobuf enum backend.Role
 */
export enum Role {
    /**
     * @generated from protobuf enum value: CONTESTANT = 0;
     */
    CONTESTANT = 0,
    /**
     * @generated from protobuf enum value: GUEST = 1;
     */
    GUEST = 1
}
/**
 * 運用的に難があるけど仕方ない・・
 * DB だけでここら辺をやるとしたら、AtCoder のスペシャルジャッジみたいに
 * シングルの Go や C++ で書かれた validator を download & compile して
 * request と response を渡してチェックしてもらうとかの形にしないといけない気がする
 *
 * @generated from protobuf enum backend.Validator
 */
export enum Validator {
    /**
     * @generated from protobuf enum value: V2022 = 0;
     */
    V2022 = 0,
    /**
     * @generated from protobuf enum value: V2023 = 1;
     */
    V2023 = 1
}
// @generated message type with reflection information, may provide speed optimized methods
class Contest$Type extends MessageType<Contest> {
    constructor() {
        super("backend.Contest", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "start_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "end_at", kind: "message", T: () => Timestamp },
            { no: 6, name: "submit_limit", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "slug", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 9, name: "tag_selection_logic", kind: "enum", T: () => ["backend.TagSelectionLogicType", TagSelectionLogicType] },
            { no: 10, name: "validator", kind: "enum", T: () => ["backend.Validator", Validator] }
        ]);
    }
    create(value?: PartialMessage<Contest>): Contest {
        const message = { id: 0, title: "", submitLimit: 0, slug: "", tagSelectionLogic: 0, validator: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Contest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Contest): Contest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 id */ 1:
                    message.id = reader.int32();
                    break;
                case /* string title */ 2:
                    message.title = reader.string();
                    break;
                case /* google.protobuf.Timestamp start_at */ 4:
                    message.startAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.startAt);
                    break;
                case /* google.protobuf.Timestamp end_at */ 5:
                    message.endAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.endAt);
                    break;
                case /* int32 submit_limit */ 6:
                    message.submitLimit = reader.int32();
                    break;
                case /* string slug */ 8:
                    message.slug = reader.string();
                    break;
                case /* backend.TagSelectionLogicType tag_selection_logic */ 9:
                    message.tagSelectionLogic = reader.int32();
                    break;
                case /* backend.Validator validator */ 10:
                    message.validator = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Contest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).int32(message.id);
        /* string title = 2; */
        if (message.title !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.title);
        /* google.protobuf.Timestamp start_at = 4; */
        if (message.startAt)
            Timestamp.internalBinaryWrite(message.startAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* google.protobuf.Timestamp end_at = 5; */
        if (message.endAt)
            Timestamp.internalBinaryWrite(message.endAt, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* int32 submit_limit = 6; */
        if (message.submitLimit !== 0)
            writer.tag(6, WireType.Varint).int32(message.submitLimit);
        /* string slug = 8; */
        if (message.slug !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.slug);
        /* backend.TagSelectionLogicType tag_selection_logic = 9; */
        if (message.tagSelectionLogic !== 0)
            writer.tag(9, WireType.Varint).int32(message.tagSelectionLogic);
        /* backend.Validator validator = 10; */
        if (message.validator !== 0)
            writer.tag(10, WireType.Varint).int32(message.validator);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message backend.Contest
 */
export const Contest = new Contest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TagSelectionLogicManual$Type extends MessageType<TagSelectionLogicManual> {
    constructor() {
        super("backend.TagSelectionLogicManual", [
            { no: 1, name: "type", kind: "enum", T: () => ["backend.TagSelectionLogicType", TagSelectionLogicType] },
            { no: 2, name: "tags_list", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Tags }
        ]);
    }
    create(value?: PartialMessage<TagSelectionLogicManual>): TagSelectionLogicManual {
        const message = { type: 0, tagsList: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TagSelectionLogicManual>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TagSelectionLogicManual): TagSelectionLogicManual {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* backend.TagSelectionLogicType type */ 1:
                    message.type = reader.int32();
                    break;
                case /* repeated backend.Tags tags_list */ 2:
                    message.tagsList.push(Tags.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: TagSelectionLogicManual, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* backend.TagSelectionLogicType type = 1; */
        if (message.type !== 0)
            writer.tag(1, WireType.Varint).int32(message.type);
        /* repeated backend.Tags tags_list = 2; */
        for (let i = 0; i < message.tagsList.length; i++)
            Tags.internalBinaryWrite(message.tagsList[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message backend.TagSelectionLogicManual
 */
export const TagSelectionLogicManual = new TagSelectionLogicManual$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TagSelectionLogicAuto$Type extends MessageType<TagSelectionLogicAuto> {
    constructor() {
        super("backend.TagSelectionLogicAuto", [
            { no: 1, name: "type", kind: "enum", T: () => ["backend.TagSelectionLogicType", TagSelectionLogicType] },
            { no: 2, name: "tags", kind: "message", T: () => Tags }
        ]);
    }
    create(value?: PartialMessage<TagSelectionLogicAuto>): TagSelectionLogicAuto {
        const message = { type: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TagSelectionLogicAuto>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TagSelectionLogicAuto): TagSelectionLogicAuto {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* backend.TagSelectionLogicType type */ 1:
                    message.type = reader.int32();
                    break;
                case /* backend.Tags tags */ 2:
                    message.tags = Tags.internalBinaryRead(reader, reader.uint32(), options, message.tags);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: TagSelectionLogicAuto, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* backend.TagSelectionLogicType type = 1; */
        if (message.type !== 0)
            writer.tag(1, WireType.Varint).int32(message.type);
        /* backend.Tags tags = 2; */
        if (message.tags)
            Tags.internalBinaryWrite(message.tags, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message backend.TagSelectionLogicAuto
 */
export const TagSelectionLogicAuto = new TagSelectionLogicAuto$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Tags$Type extends MessageType<Tags> {
    constructor() {
        super("backend.Tags", [
            { no: 1, name: "tags", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Tags>): Tags {
        const message = { tags: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Tags>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Tags): Tags {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated string tags */ 1:
                    message.tags.push(reader.string());
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Tags, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated string tags = 1; */
        for (let i = 0; i < message.tags.length; i++)
            writer.tag(1, WireType.LengthDelimited).string(message.tags[i]);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message backend.Tags
 */
export const Tags = new Tags$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Group$Type extends MessageType<Group> {
    constructor() {
        super("backend.Group", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "role", kind: "enum", T: () => ["backend.Role", Role] }
        ]);
    }
    create(value?: PartialMessage<Group>): Group {
        const message = { id: "", role: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Group>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Group): Group {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* backend.Role role */ 4:
                    message.role = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Group, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* backend.Role role = 4; */
        if (message.role !== 0)
            writer.tag(4, WireType.Varint).int32(message.role);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message backend.Group
 */
export const Group = new Group$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Submit$Type extends MessageType<Submit> {
    constructor() {
        super("backend.Submit", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "group_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "score", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "language", kind: "enum", T: () => ["backend.Language", Language] },
            { no: 6, name: "submited_at", kind: "message", T: () => Timestamp },
            { no: 7, name: "completed_at", kind: "message", T: () => Timestamp },
            { no: 8, name: "task_results", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => TaskResult },
            { no: 9, name: "status", kind: "enum", T: () => ["backend.Status", Status] },
            { no: 10, name: "error_message", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 11, name: "tag_count", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<Submit>): Submit {
        const message = { id: 0, groupName: "", score: 0, language: 0, taskResults: [], status: 0, tagCount: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Submit>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Submit): Submit {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 id */ 1:
                    message.id = reader.int32();
                    break;
                case /* string group_name */ 2:
                    message.groupName = reader.string();
                    break;
                case /* int32 score */ 4:
                    message.score = reader.int32();
                    break;
                case /* backend.Language language */ 5:
                    message.language = reader.int32();
                    break;
                case /* google.protobuf.Timestamp submited_at */ 6:
                    message.submitedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.submitedAt);
                    break;
                case /* optional google.protobuf.Timestamp completed_at */ 7:
                    message.completedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.completedAt);
                    break;
                case /* repeated backend.TaskResult task_results */ 8:
                    message.taskResults.push(TaskResult.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* backend.Status status */ 9:
                    message.status = reader.int32();
                    break;
                case /* optional string error_message */ 10:
                    message.errorMessage = reader.string();
                    break;
                case /* int32 tag_count */ 11:
                    message.tagCount = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Submit, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).int32(message.id);
        /* string group_name = 2; */
        if (message.groupName !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.groupName);
        /* int32 score = 4; */
        if (message.score !== 0)
            writer.tag(4, WireType.Varint).int32(message.score);
        /* backend.Language language = 5; */
        if (message.language !== 0)
            writer.tag(5, WireType.Varint).int32(message.language);
        /* google.protobuf.Timestamp submited_at = 6; */
        if (message.submitedAt)
            Timestamp.internalBinaryWrite(message.submitedAt, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* optional google.protobuf.Timestamp completed_at = 7; */
        if (message.completedAt)
            Timestamp.internalBinaryWrite(message.completedAt, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* repeated backend.TaskResult task_results = 8; */
        for (let i = 0; i < message.taskResults.length; i++)
            TaskResult.internalBinaryWrite(message.taskResults[i], writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* backend.Status status = 9; */
        if (message.status !== 0)
            writer.tag(9, WireType.Varint).int32(message.status);
        /* optional string error_message = 10; */
        if (message.errorMessage !== undefined)
            writer.tag(10, WireType.LengthDelimited).string(message.errorMessage);
        /* int32 tag_count = 11; */
        if (message.tagCount !== 0)
            writer.tag(11, WireType.Varint).int32(message.tagCount);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message backend.Submit
 */
export const Submit = new Submit$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TaskResult$Type extends MessageType<TaskResult> {
    constructor() {
        super("backend.TaskResult", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "request_per_sec", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "url", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "method", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "request_content_type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "request_body", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "response_code", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "response_content_type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 9, name: "response_body", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 10, name: "thread_num", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "attempt_count", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "attempt_time", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 14, name: "deleted_at", kind: "message", T: () => Timestamp },
            { no: 15, name: "error_message", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 16, name: "status", kind: "enum", T: () => ["backend.Status", Status] }
        ]);
    }
    create(value?: PartialMessage<TaskResult>): TaskResult {
        const message = { id: 0, requestPerSec: 0, url: "", method: "", requestContentType: "", responseCode: "", responseContentType: "", responseBody: "", threadNum: 0, attemptCount: 0, attemptTime: 0, status: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TaskResult>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TaskResult): TaskResult {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 id */ 1:
                    message.id = reader.int32();
                    break;
                case /* int32 request_per_sec */ 2:
                    message.requestPerSec = reader.int32();
                    break;
                case /* string url */ 3:
                    message.url = reader.string();
                    break;
                case /* string method */ 4:
                    message.method = reader.string();
                    break;
                case /* string request_content_type */ 5:
                    message.requestContentType = reader.string();
                    break;
                case /* optional string request_body */ 6:
                    message.requestBody = reader.string();
                    break;
                case /* string response_code */ 7:
                    message.responseCode = reader.string();
                    break;
                case /* string response_content_type */ 8:
                    message.responseContentType = reader.string();
                    break;
                case /* string response_body */ 9:
                    message.responseBody = reader.string();
                    break;
                case /* int32 thread_num */ 10:
                    message.threadNum = reader.int32();
                    break;
                case /* int32 attempt_count */ 11:
                    message.attemptCount = reader.int32();
                    break;
                case /* int32 attempt_time */ 12:
                    message.attemptTime = reader.int32();
                    break;
                case /* google.protobuf.Timestamp created_at */ 13:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* optional google.protobuf.Timestamp deleted_at */ 14:
                    message.deletedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.deletedAt);
                    break;
                case /* optional string error_message */ 15:
                    message.errorMessage = reader.string();
                    break;
                case /* backend.Status status */ 16:
                    message.status = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: TaskResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).int32(message.id);
        /* int32 request_per_sec = 2; */
        if (message.requestPerSec !== 0)
            writer.tag(2, WireType.Varint).int32(message.requestPerSec);
        /* string url = 3; */
        if (message.url !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.url);
        /* string method = 4; */
        if (message.method !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.method);
        /* string request_content_type = 5; */
        if (message.requestContentType !== "")
            writer.tag(5, WireType.LengthDelimited).string(message.requestContentType);
        /* optional string request_body = 6; */
        if (message.requestBody !== undefined)
            writer.tag(6, WireType.LengthDelimited).string(message.requestBody);
        /* string response_code = 7; */
        if (message.responseCode !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.responseCode);
        /* string response_content_type = 8; */
        if (message.responseContentType !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.responseContentType);
        /* string response_body = 9; */
        if (message.responseBody !== "")
            writer.tag(9, WireType.LengthDelimited).string(message.responseBody);
        /* int32 thread_num = 10; */
        if (message.threadNum !== 0)
            writer.tag(10, WireType.Varint).int32(message.threadNum);
        /* int32 attempt_count = 11; */
        if (message.attemptCount !== 0)
            writer.tag(11, WireType.Varint).int32(message.attemptCount);
        /* int32 attempt_time = 12; */
        if (message.attemptTime !== 0)
            writer.tag(12, WireType.Varint).int32(message.attemptTime);
        /* google.protobuf.Timestamp created_at = 13; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(13, WireType.LengthDelimited).fork(), options).join();
        /* optional google.protobuf.Timestamp deleted_at = 14; */
        if (message.deletedAt)
            Timestamp.internalBinaryWrite(message.deletedAt, writer.tag(14, WireType.LengthDelimited).fork(), options).join();
        /* optional string error_message = 15; */
        if (message.errorMessage !== undefined)
            writer.tag(15, WireType.LengthDelimited).string(message.errorMessage);
        /* backend.Status status = 16; */
        if (message.status !== 0)
            writer.tag(16, WireType.Varint).int32(message.status);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message backend.TaskResult
 */
export const TaskResult = new TaskResult$Type();
