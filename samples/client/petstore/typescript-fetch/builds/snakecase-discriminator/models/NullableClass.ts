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
 * @interface NullableClass
 */
export interface NullableClass {
    [key: string]: object | any;
    /**
     * 
     * @type {number}
     * @memberof NullableClass
     */
    integerProp?: number | null;
    /**
     * 
     * @type {number}
     * @memberof NullableClass
     */
    numberProp?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof NullableClass
     */
    booleanProp?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof NullableClass
     */
    stringProp?: string | null;
    /**
     * 
     * @type {Date}
     * @memberof NullableClass
     */
    dateProp?: Date | null;
    /**
     * 
     * @type {Date}
     * @memberof NullableClass
     */
    datetimeProp?: Date | null;
    /**
     * 
     * @type {Array<object>}
     * @memberof NullableClass
     */
    arrayNullableProp?: Array<object> | null;
    /**
     * 
     * @type {Array<object | null>}
     * @memberof NullableClass
     */
    arrayAndItemsNullableProp?: Array<object | null> | null;
    /**
     * 
     * @type {Array<object | null>}
     * @memberof NullableClass
     */
    arrayItemsNullable?: Array<object | null>;
    /**
     * 
     * @type {{ [key: string]: object; }}
     * @memberof NullableClass
     */
    objectNullableProp?: { [key: string]: object; } | null;
    /**
     * 
     * @type {{ [key: string]: object | null; }}
     * @memberof NullableClass
     */
    objectAndItemsNullableProp?: { [key: string]: object | null; } | null;
    /**
     * 
     * @type {{ [key: string]: object | null; }}
     * @memberof NullableClass
     */
    objectItemsNullable?: { [key: string]: object | null; };
}
/**
 * Check if a given object implements the NullableClass interface.
 */
export function instanceOfNullableClass(value: object): value is NullableClass {
    return true;
}

export function NullableClassFromJSON(json: any): NullableClass {
    return NullableClassFromJSONTyped(json, false);
}

export function NullableClassFromJSONTyped(json: any, ignoreDiscriminator: boolean): NullableClass {
    if (json == null) {
        return json;
    }
    return {
        
            ...json,
        'integerProp': json['integer_prop'] == null ? undefined : json['integer_prop'],
        'numberProp': json['number_prop'] == null ? undefined : json['number_prop'],
        'booleanProp': json['boolean_prop'] == null ? undefined : json['boolean_prop'],
        'stringProp': json['string_prop'] == null ? undefined : json['string_prop'],
        'dateProp': json['date_prop'] == null ? undefined : (new Date(json['date_prop'])),
        'datetimeProp': json['datetime_prop'] == null ? undefined : (new Date(json['datetime_prop'])),
        'arrayNullableProp': json['array_nullable_prop'] == null ? undefined : json['array_nullable_prop'],
        'arrayAndItemsNullableProp': json['array_and_items_nullable_prop'] == null ? undefined : json['array_and_items_nullable_prop'],
        'arrayItemsNullable': json['array_items_nullable'] == null ? undefined : json['array_items_nullable'],
        'objectNullableProp': json['object_nullable_prop'] == null ? undefined : json['object_nullable_prop'],
        'objectAndItemsNullableProp': json['object_and_items_nullable_prop'] == null ? undefined : json['object_and_items_nullable_prop'],
        'objectItemsNullable': json['object_items_nullable'] == null ? undefined : json['object_items_nullable'],
    };
}

  export function NullableClassToJSON(json: any): NullableClass {
      return NullableClassToJSONTyped(json, false);
  }

  export function NullableClassToJSONTyped(value?: NullableClass | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
            ...value,
        'integer_prop': value['integerProp'],
        'number_prop': value['numberProp'],
        'boolean_prop': value['booleanProp'],
        'string_prop': value['stringProp'],
        'date_prop': value['dateProp'] == null ? undefined : ((value['dateProp'] as any).toISOString().substring(0,10)),
        'datetime_prop': value['datetimeProp'] == null ? undefined : ((value['datetimeProp'] as any).toISOString()),
        'array_nullable_prop': value['arrayNullableProp'],
        'array_and_items_nullable_prop': value['arrayAndItemsNullableProp'],
        'array_items_nullable': value['arrayItemsNullable'],
        'object_nullable_prop': value['objectNullableProp'],
        'object_and_items_nullable_prop': value['objectAndItemsNullableProp'],
        'object_items_nullable': value['objectItemsNullable'],
    };
}
