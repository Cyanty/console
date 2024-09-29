// @generated by protoc-gen-es v1.6.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/console/v1alpha1/license.proto (package redpanda.api.console.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message redpanda.api.console.v1alpha1.License
 */
export class License extends Message<License> {
  /**
   * Source is where the license is used (e.g. Redpanda Cluster, Console).
   *
   * @generated from field: redpanda.api.console.v1alpha1.License.Source source = 1;
   */
  source = License_Source.UNSPECIFIED;

  /**
   * Type is the type of license (community, trial, enterprise).
   *
   * @generated from field: redpanda.api.console.v1alpha1.License.Type type = 2;
   */
  type = License_Type.UNSPECIFIED;

  /**
   * UnixEpochSeconds is the timestamp when the license is going to expire.
   *
   * @generated from field: int64 expires_at = 3;
   */
  expiresAt = protoInt64.zero;

  constructor(data?: PartialMessage<License>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.License";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "source", kind: "enum", T: proto3.getEnumType(License_Source) },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(License_Type) },
    { no: 3, name: "expires_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): License {
    return new License().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): License {
    return new License().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): License {
    return new License().fromJsonString(jsonString, options);
  }

  static equals(a: License | PlainMessage<License> | undefined, b: License | PlainMessage<License> | undefined): boolean {
    return proto3.util.equals(License, a, b);
  }
}

/**
 * @generated from enum redpanda.api.console.v1alpha1.License.Source
 */
export enum License_Source {
  /**
   * @generated from enum value: SOURCE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SOURCE_REDPANDA_CONSOLE = 1;
   */
  REDPANDA_CONSOLE = 1,

  /**
   * @generated from enum value: SOURCE_REDPANDA_CORE = 2;
   */
  REDPANDA_CORE = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(License_Source)
proto3.util.setEnumType(License_Source, "redpanda.api.console.v1alpha1.License.Source", [
  { no: 0, name: "SOURCE_UNSPECIFIED" },
  { no: 1, name: "SOURCE_REDPANDA_CONSOLE" },
  { no: 2, name: "SOURCE_REDPANDA_CORE" },
]);

/**
 * @generated from enum redpanda.api.console.v1alpha1.License.Type
 */
export enum License_Type {
  /**
   * @generated from enum value: TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: TYPE_COMMUNITY = 1;
   */
  COMMUNITY = 1,

  /**
   * @generated from enum value: TYPE_TRIAL = 2;
   */
  TRIAL = 2,

  /**
   * @generated from enum value: TYPE_ENTERPRISE = 3;
   */
  ENTERPRISE = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(License_Type)
proto3.util.setEnumType(License_Type, "redpanda.api.console.v1alpha1.License.Type", [
  { no: 0, name: "TYPE_UNSPECIFIED" },
  { no: 1, name: "TYPE_COMMUNITY" },
  { no: 2, name: "TYPE_TRIAL" },
  { no: 3, name: "TYPE_ENTERPRISE" },
]);

/**
 * @generated from message redpanda.api.console.v1alpha1.ListLicensesRequest
 */
export class ListLicensesRequest extends Message<ListLicensesRequest> {
  constructor(data?: PartialMessage<ListLicensesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.ListLicensesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListLicensesRequest {
    return new ListLicensesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListLicensesRequest {
    return new ListLicensesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListLicensesRequest {
    return new ListLicensesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListLicensesRequest | PlainMessage<ListLicensesRequest> | undefined, b: ListLicensesRequest | PlainMessage<ListLicensesRequest> | undefined): boolean {
    return proto3.util.equals(ListLicensesRequest, a, b);
  }
}

/**
 * @generated from message redpanda.api.console.v1alpha1.ListLicensesResponse
 */
export class ListLicensesResponse extends Message<ListLicensesResponse> {
  /**
   * @generated from field: repeated redpanda.api.console.v1alpha1.License licenses = 1;
   */
  licenses: License[] = [];

  /**
   * @generated from field: bool violation = 2;
   */
  violation = false;

  constructor(data?: PartialMessage<ListLicensesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.ListLicensesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "licenses", kind: "message", T: License, repeated: true },
    { no: 2, name: "violation", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListLicensesResponse {
    return new ListLicensesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListLicensesResponse {
    return new ListLicensesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListLicensesResponse {
    return new ListLicensesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListLicensesResponse | PlainMessage<ListLicensesResponse> | undefined, b: ListLicensesResponse | PlainMessage<ListLicensesResponse> | undefined): boolean {
    return proto3.util.equals(ListLicensesResponse, a, b);
  }
}

/**
 * @generated from message redpanda.api.console.v1alpha1.SetLicenseRequest
 */
export class SetLicenseRequest extends Message<SetLicenseRequest> {
  /**
   * @generated from field: string license = 1;
   */
  license = "";

  constructor(data?: PartialMessage<SetLicenseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.SetLicenseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "license", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetLicenseRequest {
    return new SetLicenseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetLicenseRequest {
    return new SetLicenseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetLicenseRequest {
    return new SetLicenseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SetLicenseRequest | PlainMessage<SetLicenseRequest> | undefined, b: SetLicenseRequest | PlainMessage<SetLicenseRequest> | undefined): boolean {
    return proto3.util.equals(SetLicenseRequest, a, b);
  }
}

/**
 * @generated from message redpanda.api.console.v1alpha1.SetLicenseResponse
 */
export class SetLicenseResponse extends Message<SetLicenseResponse> {
  /**
   * @generated from field: redpanda.api.console.v1alpha1.License license = 1;
   */
  license?: License;

  constructor(data?: PartialMessage<SetLicenseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.SetLicenseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "license", kind: "message", T: License },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetLicenseResponse {
    return new SetLicenseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetLicenseResponse {
    return new SetLicenseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetLicenseResponse {
    return new SetLicenseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SetLicenseResponse | PlainMessage<SetLicenseResponse> | undefined, b: SetLicenseResponse | PlainMessage<SetLicenseResponse> | undefined): boolean {
    return proto3.util.equals(SetLicenseResponse, a, b);
  }
}

/**
 * @generated from message redpanda.api.console.v1alpha1.ListEnterpriseFeaturesRequest
 */
export class ListEnterpriseFeaturesRequest extends Message<ListEnterpriseFeaturesRequest> {
  constructor(data?: PartialMessage<ListEnterpriseFeaturesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.ListEnterpriseFeaturesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEnterpriseFeaturesRequest {
    return new ListEnterpriseFeaturesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEnterpriseFeaturesRequest {
    return new ListEnterpriseFeaturesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEnterpriseFeaturesRequest {
    return new ListEnterpriseFeaturesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListEnterpriseFeaturesRequest | PlainMessage<ListEnterpriseFeaturesRequest> | undefined, b: ListEnterpriseFeaturesRequest | PlainMessage<ListEnterpriseFeaturesRequest> | undefined): boolean {
    return proto3.util.equals(ListEnterpriseFeaturesRequest, a, b);
  }
}

/**
 * @generated from message redpanda.api.console.v1alpha1.ListEnterpriseFeaturesResponse
 */
export class ListEnterpriseFeaturesResponse extends Message<ListEnterpriseFeaturesResponse> {
  /**
   * LicenseStatus reports the status of the installed license in the Redpanda cluster.
   *
   * @generated from field: redpanda.api.console.v1alpha1.ListEnterpriseFeaturesResponse.LicenseStatus license_status = 1;
   */
  licenseStatus = ListEnterpriseFeaturesResponse_LicenseStatus.UNSPECIFIED;

  /**
   * Violation is true if license_status is not 'valid' AND one or more enterprise features are enabled
   *
   * @generated from field: bool violation = 2;
   */
  violation = false;

  /**
   * Features is a ist of enterprise features (name and whether in use)
   *
   * @generated from field: repeated redpanda.api.console.v1alpha1.ListEnterpriseFeaturesResponse.Feature features = 3;
   */
  features: ListEnterpriseFeaturesResponse_Feature[] = [];

  constructor(data?: PartialMessage<ListEnterpriseFeaturesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.ListEnterpriseFeaturesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "license_status", kind: "enum", T: proto3.getEnumType(ListEnterpriseFeaturesResponse_LicenseStatus) },
    { no: 2, name: "violation", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "features", kind: "message", T: ListEnterpriseFeaturesResponse_Feature, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEnterpriseFeaturesResponse {
    return new ListEnterpriseFeaturesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEnterpriseFeaturesResponse {
    return new ListEnterpriseFeaturesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEnterpriseFeaturesResponse {
    return new ListEnterpriseFeaturesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListEnterpriseFeaturesResponse | PlainMessage<ListEnterpriseFeaturesResponse> | undefined, b: ListEnterpriseFeaturesResponse | PlainMessage<ListEnterpriseFeaturesResponse> | undefined): boolean {
    return proto3.util.equals(ListEnterpriseFeaturesResponse, a, b);
  }
}

/**
 * @generated from enum redpanda.api.console.v1alpha1.ListEnterpriseFeaturesResponse.LicenseStatus
 */
export enum ListEnterpriseFeaturesResponse_LicenseStatus {
  /**
   * @generated from enum value: LICENSE_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: LICENSE_STATUS_VALID = 1;
   */
  VALID = 1,

  /**
   * @generated from enum value: LICENSE_STATUS_EXPIRED = 2;
   */
  EXPIRED = 2,

  /**
   * @generated from enum value: LICENSE_STATUS_NOT_PRESENT = 3;
   */
  NOT_PRESENT = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(ListEnterpriseFeaturesResponse_LicenseStatus)
proto3.util.setEnumType(ListEnterpriseFeaturesResponse_LicenseStatus, "redpanda.api.console.v1alpha1.ListEnterpriseFeaturesResponse.LicenseStatus", [
  { no: 0, name: "LICENSE_STATUS_UNSPECIFIED" },
  { no: 1, name: "LICENSE_STATUS_VALID" },
  { no: 2, name: "LICENSE_STATUS_EXPIRED" },
  { no: 3, name: "LICENSE_STATUS_NOT_PRESENT" },
]);

/**
 * Feature is an enterprise feature and a bool indicating whether it's in use.
 *
 * @generated from message redpanda.api.console.v1alpha1.ListEnterpriseFeaturesResponse.Feature
 */
export class ListEnterpriseFeaturesResponse_Feature extends Message<ListEnterpriseFeaturesResponse_Feature> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: bool enabled = 2;
   */
  enabled = false;

  constructor(data?: PartialMessage<ListEnterpriseFeaturesResponse_Feature>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.ListEnterpriseFeaturesResponse.Feature";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEnterpriseFeaturesResponse_Feature {
    return new ListEnterpriseFeaturesResponse_Feature().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEnterpriseFeaturesResponse_Feature {
    return new ListEnterpriseFeaturesResponse_Feature().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEnterpriseFeaturesResponse_Feature {
    return new ListEnterpriseFeaturesResponse_Feature().fromJsonString(jsonString, options);
  }

  static equals(a: ListEnterpriseFeaturesResponse_Feature | PlainMessage<ListEnterpriseFeaturesResponse_Feature> | undefined, b: ListEnterpriseFeaturesResponse_Feature | PlainMessage<ListEnterpriseFeaturesResponse_Feature> | undefined): boolean {
    return proto3.util.equals(ListEnterpriseFeaturesResponse_Feature, a, b);
  }
}
