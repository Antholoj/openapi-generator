/* tslint:disable */
/* eslint-disable */
/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { mapValues } from '../runtime';
import type { DeprecatedObject } from './DeprecatedObject';
import {
    DeprecatedObjectFromJSON,
    DeprecatedObjectFromJSONTyped,
    DeprecatedObjectToJSON,
    DeprecatedObjectToJSONTyped,
} from './DeprecatedObject';

/**
 * 
 * @export
 * @interface ObjectWithDeprecatedFields
 */
export interface ObjectWithDeprecatedFields {
    /**
     * 
     * @type {string}
     * @memberof ObjectWithDeprecatedFields
     */
    uuid?: string;
    /**
     * 
     * @type {number}
     * @memberof ObjectWithDeprecatedFields
     * @deprecated
     */
    id?: number;
    /**
     * 
     * @type {DeprecatedObject}
     * @memberof ObjectWithDeprecatedFields
     * @deprecated
     */
    deprecatedRef?: DeprecatedObject;
    /**
     * 
     * @type {Array<string>}
     * @memberof ObjectWithDeprecatedFields
     * @deprecated
     */
    bars?: Array<string>;
}
/**
 * Check if a given object implements the ObjectWithDeprecatedFields interface.
 */
export function instanceOfObjectWithDeprecatedFields(value: object): value is ObjectWithDeprecatedFields {
    return true;
}

export function ObjectWithDeprecatedFieldsFromJSON(json: any): ObjectWithDeprecatedFields {
    return ObjectWithDeprecatedFieldsFromJSONTyped(json, false);
}

export function ObjectWithDeprecatedFieldsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ObjectWithDeprecatedFields {
    if (json == null) {
        return json;
    }
    return {
        
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
        'id': json['id'] == null ? undefined : json['id'],
        'deprecatedRef': json['deprecatedRef'] == null ? undefined : DeprecatedObjectFromJSON(json['deprecatedRef']),
        'bars': json['bars'] == null ? undefined : json['bars'],
    };
}

  export function ObjectWithDeprecatedFieldsToJSON(json: any): ObjectWithDeprecatedFields {
      return ObjectWithDeprecatedFieldsToJSONTyped(json, false);
  }

  export function ObjectWithDeprecatedFieldsToJSONTyped(value?: ObjectWithDeprecatedFields | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'uuid': value['uuid'],
        'id': value['id'],
        'deprecatedRef': DeprecatedObjectToJSON(value['deprecatedRef']),
        'bars': value['bars'],
    };
}
