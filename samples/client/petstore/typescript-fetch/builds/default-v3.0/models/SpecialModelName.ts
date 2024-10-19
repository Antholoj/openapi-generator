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
 * @interface SpecialModelName
 */
export interface SpecialModelName {
    /**
     * 
     * @type {number}
     * @memberof SpecialModelName
     */
    $specialPropertyName?: number;
}
/**
 * Check if a given object implements the SpecialModelName interface.
 */
export function instanceOfSpecialModelName(value: object): value is SpecialModelName {
    return true;
}

export function SpecialModelNameFromJSON(json: any): SpecialModelName {
    return SpecialModelNameFromJSONTyped(json, false);
}

export function SpecialModelNameFromJSONTyped(json: any, ignoreDiscriminator: boolean): SpecialModelName {
    if (json == null) {
        return json;
    }
    return {
        
        '$specialPropertyName': json['$special[property.name]'] == null ? undefined : json['$special[property.name]'],
    };
}

  export function SpecialModelNameToJSON(json: any): SpecialModelName {
      return SpecialModelNameToJSONTyped(json, false);
  }

  export function SpecialModelNameToJSONTyped(value?: SpecialModelName | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        '$special[property.name]': value['$specialPropertyName'],
    };
}
