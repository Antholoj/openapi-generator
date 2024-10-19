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

/**
 * 
 * @export
 */
export const OuterEnumIntegerDefaultValue = {
    NUMBER_0: 0,
    NUMBER_1: 1,
    NUMBER_2: 2
} as const;
export type OuterEnumIntegerDefaultValue = typeof OuterEnumIntegerDefaultValue[keyof typeof OuterEnumIntegerDefaultValue];


export function instanceOfOuterEnumIntegerDefaultValue(value: any): boolean {
    for (const key in OuterEnumIntegerDefaultValue) {
        if (Object.prototype.hasOwnProperty.call(OuterEnumIntegerDefaultValue, key)) {
            if (OuterEnumIntegerDefaultValue[key as keyof typeof OuterEnumIntegerDefaultValue] === value) {
                return true;
            }
        }
    }
    return false;
}

export function OuterEnumIntegerDefaultValueFromJSON(json: any): OuterEnumIntegerDefaultValue {
    return OuterEnumIntegerDefaultValueFromJSONTyped(json, false);
}

export function OuterEnumIntegerDefaultValueFromJSONTyped(json: any, ignoreDiscriminator: boolean): OuterEnumIntegerDefaultValue {
    return json as OuterEnumIntegerDefaultValue;
}

export function OuterEnumIntegerDefaultValueToJSON(value?: OuterEnumIntegerDefaultValue | null): any {
    return value as any;
}

export function OuterEnumIntegerDefaultValueToJSONTyped(value: any, ignoreDiscriminator: boolean): OuterEnumIntegerDefaultValue {
    return value as OuterEnumIntegerDefaultValue;
}
