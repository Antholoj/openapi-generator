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
import { Cat, CatFromJSONTyped, CatToJSON, CatToJSONTyped } from './Cat';
import { Dog, DogFromJSONTyped, DogToJSON, DogToJSONTyped } from './Dog';
/**
 * 
 * @export
 * @interface Animal
 */
export interface Animal {
    /**
     * 
     * @type {string}
     * @memberof Animal
     */
    className: string;
    /**
     * 
     * @type {string}
     * @memberof Animal
     */
    color?: string;
}
/**
 * Check if a given object implements the Animal interface.
 */
export function instanceOfAnimal(value: object): value is Animal {
    if (!('className' in value) || value['className'] === undefined) return false;
    return true;
}

export function AnimalFromJSON(json: any): Animal {
    return AnimalFromJSONTyped(json, false);
}

export function AnimalFromJSONTyped(json: any, ignoreDiscriminator: boolean): Animal {
    if (json == null) {
        return json;
    }
    if (!ignoreDiscriminator) {
        if (json['class_name'] === 'CAT') {
            return CatFromJSONTyped(json, ignoreDiscriminator);
        }
        if (json['class_name'] === 'DOG') {
            return DogFromJSONTyped(json, ignoreDiscriminator);
        }
    }
    return {
        
        'className': json['class_name'],
        'color': json['color'] == null ? undefined : json['color'],
    };
}

  export function AnimalToJSON(json: any): Animal {
      return AnimalToJSONTyped(json, false);
  }

  export function AnimalToJSONTyped(value?: Animal | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    if (!ignoreDiscriminator) {
        switch (value['className']) {
            case 'CAT':
                return CatToJSONTyped(value as Cat, ignoreDiscriminator);
            case 'DOG':
                return DogToJSONTyped(value as Dog, ignoreDiscriminator);
            default:
                throw new Error(`No variant of Animal exists with 'className=${value['className']}'`);
        }
    }

    return {
        
        'class_name': value['className'],
        'color': value['color'],
    };
}
