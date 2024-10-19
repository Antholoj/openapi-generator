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
 * @interface ModelApiResponse
 */
export interface ModelApiResponse {
    /**
     * 
     * @type {number}
     * @memberof ModelApiResponse
     */
    code?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelApiResponse
     */
    type?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelApiResponse
     */
    message?: string;
}
/**
 * Check if a given object implements the ModelApiResponse interface.
 */
export function instanceOfModelApiResponse(value: object): value is ModelApiResponse {
    return true;
}

export function ModelApiResponseFromJSON(json: any): ModelApiResponse {
    return ModelApiResponseFromJSONTyped(json, false);
}

export function ModelApiResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelApiResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'code': json['code'] == null ? undefined : json['code'],
        'type': json['type'] == null ? undefined : json['type'],
        'message': json['message'] == null ? undefined : json['message'],
    };
}

  export function ModelApiResponseToJSON(json: any): ModelApiResponse {
      return ModelApiResponseToJSONTyped(json, false);
  }

  export function ModelApiResponseToJSONTyped(value?: ModelApiResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'code': value['code'],
        'type': value['type'],
        'message': value['message'],
    };
}
