/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { AggregatedKeyShare } from "./aggregated_key_share";
import { AuthorizedAddress } from "./authorized_address";
import { GeneralKeyShare } from "./general_key_share";
import { KeyShare } from "./key_share";
import { Params } from "./params";
import { ActivePubKey, QueuedPubKey } from "./pub_key";
import { ValidatorSet } from "./validator_set";

export const protobufPackage = "fairyring.keyshare";

/** GenesisState defines the keyshare module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  portId: string;
  validatorSetList: ValidatorSet[];
  keyShareList: KeyShare[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  aggregatedKeyShareList: AggregatedKeyShare[];
  activePubKey: ActivePubKey | undefined;
  queuedPubKey: QueuedPubKey | undefined;
  authorizedAddressList: AuthorizedAddress[];
  requestCount: number;
  generalKeyShareList: GeneralKeyShare[];
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    portId: "",
    validatorSetList: [],
    keyShareList: [],
    aggregatedKeyShareList: [],
    activePubKey: undefined,
    queuedPubKey: undefined,
    authorizedAddressList: [],
    requestCount: 0,
    generalKeyShareList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.portId !== "") {
      writer.uint32(18).string(message.portId);
    }
    for (const v of message.validatorSetList) {
      ValidatorSet.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.keyShareList) {
      KeyShare.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.aggregatedKeyShareList) {
      AggregatedKeyShare.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.activePubKey !== undefined) {
      ActivePubKey.encode(message.activePubKey, writer.uint32(50).fork()).ldelim();
    }
    if (message.queuedPubKey !== undefined) {
      QueuedPubKey.encode(message.queuedPubKey, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.authorizedAddressList) {
      AuthorizedAddress.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.requestCount !== 0) {
      writer.uint32(72).uint64(message.requestCount);
    }
    for (const v of message.generalKeyShareList) {
      GeneralKeyShare.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    if (message.activePubKey !== undefined) {
      ActivePubKey.encode(message.activePubKey, writer.uint32(42).fork()).ldelim();
    }
    if (message.queuedPubKey !== undefined) {
      QueuedPubKey.encode(message.queuedPubKey, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.authorizedAddressList) {
      AuthorizedAddress.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.portId = reader.string();
          break;
        case 3:
          message.validatorSetList.push(ValidatorSet.decode(reader, reader.uint32()));
          break;
        case 4:
          message.keyShareList.push(KeyShare.decode(reader, reader.uint32()));
          break;
        case 5:
          message.aggregatedKeyShareList.push(AggregatedKeyShare.decode(reader, reader.uint32()));
          break;
        case 6:
          message.activePubKey = ActivePubKey.decode(reader, reader.uint32());
          break;
        case 7:
          message.queuedPubKey = QueuedPubKey.decode(reader, reader.uint32());
          break;
        case 8:
          message.authorizedAddressList.push(AuthorizedAddress.decode(reader, reader.uint32()));
          break;
        case 9:
          message.requestCount = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.generalKeyShareList.push(GeneralKeyShare.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      portId: isSet(object.portId) ? String(object.portId) : "",
      validatorSetList: Array.isArray(object?.validatorSetList)
        ? object.validatorSetList.map((e: any) => ValidatorSet.fromJSON(e))
        : [],
      keyShareList: Array.isArray(object?.keyShareList)
        ? object.keyShareList.map((e: any) => KeyShare.fromJSON(e))
        : [],
      aggregatedKeyShareList: Array.isArray(object?.aggregatedKeyShareList)
        ? object.aggregatedKeyShareList.map((e: any) => AggregatedKeyShare.fromJSON(e))
        : [],
      activePubKey: isSet(object.activePubKey) ? ActivePubKey.fromJSON(object.activePubKey) : undefined,
      queuedPubKey: isSet(object.queuedPubKey) ? QueuedPubKey.fromJSON(object.queuedPubKey) : undefined,
      authorizedAddressList: Array.isArray(object?.authorizedAddressList)
        ? object.authorizedAddressList.map((e: any) => AuthorizedAddress.fromJSON(e))
        : [],
      requestCount: isSet(object.requestCount) ? Number(object.requestCount) : 0,
      generalKeyShareList: Array.isArray(object?.generalKeyShareList)
        ? object.generalKeyShareList.map((e: any) => GeneralKeyShare.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.portId !== undefined && (obj.portId = message.portId);
    if (message.validatorSetList) {
      obj.validatorSetList = message.validatorSetList.map((e) => e ? ValidatorSet.toJSON(e) : undefined);
    } else {
      obj.validatorSetList = [];
    }
    if (message.keyShareList) {
      obj.keyShareList = message.keyShareList.map((e) => e ? KeyShare.toJSON(e) : undefined);
    } else {
      obj.keyShareList = [];
    }
    if (message.aggregatedKeyShareList) {
      obj.aggregatedKeyShareList = message.aggregatedKeyShareList.map((e) =>
        e ? AggregatedKeyShare.toJSON(e) : undefined
      );
    } else {
      obj.aggregatedKeyShareList = [];
    }
    message.activePubKey !== undefined
      && (obj.activePubKey = message.activePubKey ? ActivePubKey.toJSON(message.activePubKey) : undefined);
    message.queuedPubKey !== undefined
      && (obj.queuedPubKey = message.queuedPubKey ? QueuedPubKey.toJSON(message.queuedPubKey) : undefined);
    if (message.authorizedAddressList) {
      obj.authorizedAddressList = message.authorizedAddressList.map((e) => e ? AuthorizedAddress.toJSON(e) : undefined);
    } else {
      obj.authorizedAddressList = [];
    }
    message.requestCount !== undefined && (obj.requestCount = Math.round(message.requestCount));
    if (message.generalKeyShareList) {
      obj.generalKeyShareList = message.generalKeyShareList.map((e) => e ? GeneralKeyShare.toJSON(e) : undefined);
    } else {
      obj.generalKeyShareList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.portId = object.portId ?? "";
    message.validatorSetList = object.validatorSetList?.map((e) => ValidatorSet.fromPartial(e)) || [];
    message.keyShareList = object.keyShareList?.map((e) => KeyShare.fromPartial(e)) || [];
    message.aggregatedKeyShareList = object.aggregatedKeyShareList?.map((e) => AggregatedKeyShare.fromPartial(e)) || [];
    message.activePubKey = (object.activePubKey !== undefined && object.activePubKey !== null)
      ? ActivePubKey.fromPartial(object.activePubKey)
      : undefined;
    message.queuedPubKey = (object.queuedPubKey !== undefined && object.queuedPubKey !== null)
      ? QueuedPubKey.fromPartial(object.queuedPubKey)
      : undefined;
    message.authorizedAddressList = object.authorizedAddressList?.map((e) => AuthorizedAddress.fromPartial(e)) || [];
    message.requestCount = object.requestCount ?? 0;
    message.generalKeyShareList = object.generalKeyShareList?.map((e) => GeneralKeyShare.fromPartial(e)) || [];
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
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
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
