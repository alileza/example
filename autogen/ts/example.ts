/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';


export interface Empty {
}

export interface StatusResponse {
  ok: boolean;
  message: string;
}

export interface HelloRequest {
  name: string;
}

export interface HelloResponse {
  world: string;
}

const baseEmpty: object = {
};

const baseStatusResponse: object = {
  ok: false,
  message: "",
};

const baseHelloRequest: object = {
  name: "",
};

const baseHelloResponse: object = {
  world: "",
};

export interface ExampleServiceV1 {

  status(request: Empty): Promise<StatusResponse>;

  hello(request: HelloRequest): Promise<HelloResponse>;

}

export const protobufPackage = 'example'

export const Empty = {
  encode(_: Empty, writer: Writer = Writer.create()): Writer {
    return writer;
  },
  decode(input: Uint8Array | Reader, length?: number): Empty {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEmpty } as Empty;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(_: any): Empty {
    const message = { ...baseEmpty } as Empty;
    return message;
  },
  fromPartial(_: DeepPartial<Empty>): Empty {
    const message = { ...baseEmpty } as Empty;
    return message;
  },
  toJSON(_: Empty): unknown {
    const obj: any = {};
    return obj;
  },
};

export const StatusResponse = {
  encode(message: StatusResponse, writer: Writer = Writer.create()): Writer {
    writer.uint32(8).bool(message.ok);
    writer.uint32(18).string(message.message);
    return writer;
  },
  decode(input: Uint8Array | Reader, length?: number): StatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStatusResponse } as StatusResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ok = reader.bool();
          break;
        case 2:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): StatusResponse {
    const message = { ...baseStatusResponse } as StatusResponse;
    if (object.ok !== undefined && object.ok !== null) {
      message.ok = Boolean(object.ok);
    } else {
      message.ok = false;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<StatusResponse>): StatusResponse {
    const message = { ...baseStatusResponse } as StatusResponse;
    if (object.ok !== undefined && object.ok !== null) {
      message.ok = object.ok;
    } else {
      message.ok = false;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    return message;
  },
  toJSON(message: StatusResponse): unknown {
    const obj: any = {};
    message.ok !== undefined && (obj.ok = message.ok);
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },
};

export const HelloRequest = {
  encode(message: HelloRequest, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.name);
    return writer;
  },
  decode(input: Uint8Array | Reader, length?: number): HelloRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseHelloRequest } as HelloRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): HelloRequest {
    const message = { ...baseHelloRequest } as HelloRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<HelloRequest>): HelloRequest {
    const message = { ...baseHelloRequest } as HelloRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
  toJSON(message: HelloRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },
};

export const HelloResponse = {
  encode(message: HelloResponse, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).string(message.world);
    return writer;
  },
  decode(input: Uint8Array | Reader, length?: number): HelloResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseHelloResponse } as HelloResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.world = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): HelloResponse {
    const message = { ...baseHelloResponse } as HelloResponse;
    if (object.world !== undefined && object.world !== null) {
      message.world = String(object.world);
    } else {
      message.world = "";
    }
    return message;
  },
  fromPartial(object: DeepPartial<HelloResponse>): HelloResponse {
    const message = { ...baseHelloResponse } as HelloResponse;
    if (object.world !== undefined && object.world !== null) {
      message.world = object.world;
    } else {
      message.world = "";
    }
    return message;
  },
  toJSON(message: HelloResponse): unknown {
    const obj: any = {};
    message.world !== undefined && (obj.world = message.world);
    return obj;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;