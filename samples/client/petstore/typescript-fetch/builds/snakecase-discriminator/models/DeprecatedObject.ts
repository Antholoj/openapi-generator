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
/**
 * 
 * @export
 * @interface DeprecatedObject
 */
export interface DeprecatedObject {
    /**
     * 
     * @type {string}
     * @memberof DeprecatedObject
     */
    name?: string;
}
/**
 * Check if a given object implements the DeprecatedObject interface.
 */
export function instanceOfDeprecatedObject(value: object): value is DeprecatedObject {
    return true;
}

export function DeprecatedObjectFromJSON(json: any): DeprecatedObject {
    return DeprecatedObjectFromJSONTyped(json, false);
}

export function DeprecatedObjectFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeprecatedObject {
    if (json == null) {
        return json;
    }
    return {
        
        'name': json['name'] == null ? undefined : json['name'],
    };
}

  export function DeprecatedObjectToJSON(json: any): DeprecatedObject {
      return DeprecatedObjectToJSONTyped(json, false);
  }

  export function DeprecatedObjectToJSONTyped(value?: DeprecatedObject | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'name': value['name'],
    };
}
