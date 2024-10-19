/* tslint:disable */
/* eslint-disable */
/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { mapValues } from '../runtime';
import type { ResponseMeta } from './ResponseMeta';
import {
    ResponseMetaFromJSON,
    ResponseMetaFromJSONTyped,
    ResponseMetaToJSON,
    ResponseMetaToJSONTyped,
} from './ResponseMeta';
import type { PetPartType } from './PetPartType';
import {
    PetPartTypeFromJSON,
    PetPartTypeFromJSONTyped,
    PetPartTypeToJSON,
    PetPartTypeToJSONTyped,
} from './PetPartType';

/**
 * 
 * @export
 * @interface GetPetPartTypeResponse
 */
export interface GetPetPartTypeResponse {
    /**
     * 
     * @type {ResponseMeta}
     * @memberof GetPetPartTypeResponse
     */
    meta: ResponseMeta;
    /**
     * 
     * @type {PetPartType}
     * @memberof GetPetPartTypeResponse
     */
    data?: PetPartType;
}


/**
 * Check if a given object implements the GetPetPartTypeResponse interface.
 */
export function instanceOfGetPetPartTypeResponse(value: object): value is GetPetPartTypeResponse {
    if (!('meta' in value) || value['meta'] === undefined) return false;
    return true;
}

export function GetPetPartTypeResponseFromJSON(json: any): GetPetPartTypeResponse {
    return GetPetPartTypeResponseFromJSONTyped(json, false);
}

export function GetPetPartTypeResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetPetPartTypeResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'meta': ResponseMetaFromJSON(json['meta']),
        'data': json['data'] == null ? undefined : PetPartTypeFromJSON(json['data']),
    };
}

  export function GetPetPartTypeResponseToJSON(json: any): GetPetPartTypeResponse {
      return GetPetPartTypeResponseToJSONTyped(json, false);
  }

  export function GetPetPartTypeResponseToJSONTyped(value?: GetPetPartTypeResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'meta': ResponseMetaToJSON(value['meta']),
        'data': PetPartTypeToJSON(value['data']),
    };
}
