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
 * @interface MapTest
 */
export interface MapTest {
    /**
     * 
     * @type {{ [key: string]: { [key: string]: string; }; }}
     * @memberof MapTest
     */
    mapMapOfString?: { [key: string]: { [key: string]: string; }; };
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof MapTest
     */
    mapOfEnumString?: { [key: string]: string; };
    /**
     * 
     * @type {{ [key: string]: boolean; }}
     * @memberof MapTest
     */
    directMap?: { [key: string]: boolean; };
    /**
     * 
     * @type {{ [key: string]: boolean; }}
     * @memberof MapTest
     */
    indirectMap?: { [key: string]: boolean; };
}


/**
 * @export
 */
export const MapTestMapOfEnumStringEnum = {
    Upper: 'UPPER',
    Lower: 'lower'
} as const;
export type MapTestMapOfEnumStringEnum = typeof MapTestMapOfEnumStringEnum[keyof typeof MapTestMapOfEnumStringEnum];

/**
 * Check if a given object implements the MapTest interface.
 */
export function instanceOfMapTest(value: object): value is MapTest {
    return true;
}

export function MapTestFromJSON(json: any): MapTest {
    return MapTestFromJSONTyped(json, false);
}

export function MapTestFromJSONTyped(json: any, ignoreDiscriminator: boolean): MapTest {
    if (json == null) {
        return json;
    }
    return {
        
        'mapMapOfString': json['map_map_of_string'] == null ? undefined : json['map_map_of_string'],
        'mapOfEnumString': json['map_of_enum_string'] == null ? undefined : json['map_of_enum_string'],
        'directMap': json['direct_map'] == null ? undefined : json['direct_map'],
        'indirectMap': json['indirect_map'] == null ? undefined : json['indirect_map'],
    };
}

  export function MapTestToJSON(json: any): MapTest {
      return MapTestToJSONTyped(json, false);
  }

  export function MapTestToJSONTyped(value?: MapTest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'map_map_of_string': value['mapMapOfString'],
        'map_of_enum_string': value['mapOfEnumString'],
        'direct_map': value['directMap'],
        'indirect_map': value['indirectMap'],
    };
}
