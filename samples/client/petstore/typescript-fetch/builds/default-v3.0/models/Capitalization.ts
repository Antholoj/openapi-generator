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
 * @interface Capitalization
 */
export interface Capitalization {
    /**
     * 
     * @type {string}
     * @memberof Capitalization
     */
    smallCamel?: string;
    /**
     * 
     * @type {string}
     * @memberof Capitalization
     */
    capitalCamel?: string;
    /**
     * 
     * @type {string}
     * @memberof Capitalization
     */
    smallSnake?: string;
    /**
     * 
     * @type {string}
     * @memberof Capitalization
     */
    capitalSnake?: string;
    /**
     * 
     * @type {string}
     * @memberof Capitalization
     */
    sCAETHFlowPoints?: string;
    /**
     * Name of the pet
     * 
     * @type {string}
     * @memberof Capitalization
     */
    aTTNAME?: string;
}
/**
 * Check if a given object implements the Capitalization interface.
 */
export function instanceOfCapitalization(value: object): value is Capitalization {
    return true;
}

export function CapitalizationFromJSON(json: any): Capitalization {
    return CapitalizationFromJSONTyped(json, false);
}

export function CapitalizationFromJSONTyped(json: any, ignoreDiscriminator: boolean): Capitalization {
    if (json == null) {
        return json;
    }
    return {
        
        'smallCamel': json['smallCamel'] == null ? undefined : json['smallCamel'],
        'capitalCamel': json['CapitalCamel'] == null ? undefined : json['CapitalCamel'],
        'smallSnake': json['small_Snake'] == null ? undefined : json['small_Snake'],
        'capitalSnake': json['Capital_Snake'] == null ? undefined : json['Capital_Snake'],
        'sCAETHFlowPoints': json['SCA_ETH_Flow_Points'] == null ? undefined : json['SCA_ETH_Flow_Points'],
        'aTTNAME': json['ATT_NAME'] == null ? undefined : json['ATT_NAME'],
    };
}

  export function CapitalizationToJSON(json: any): Capitalization {
      return CapitalizationToJSONTyped(json, false);
  }

  export function CapitalizationToJSONTyped(value?: Capitalization | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'smallCamel': value['smallCamel'],
        'CapitalCamel': value['capitalCamel'],
        'small_Snake': value['smallSnake'],
        'Capital_Snake': value['capitalSnake'],
        'SCA_ETH_Flow_Points': value['sCAETHFlowPoints'],
        'ATT_NAME': value['aTTNAME'],
    };
}
