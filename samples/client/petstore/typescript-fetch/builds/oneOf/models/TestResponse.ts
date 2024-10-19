/* tslint:disable */
/* eslint-disable */
/**
 * testing oneOf without discriminator
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { TestA } from './TestA';
import {
    instanceOfTestA,
    TestAFromJSON,
    TestAFromJSONTyped,
    TestAToJSON,
} from './TestA';
import type { TestB } from './TestB';
import {
    instanceOfTestB,
    TestBFromJSON,
    TestBFromJSONTyped,
    TestBToJSON,
} from './TestB';

/**
 * @type TestResponse
 * 
 * @export
 */
export type TestResponse = TestA | TestB;
export function TestResponseFromJSON(json: any): TestResponse {
    return TestResponseFromJSONTyped(json, false);
}

export function TestResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): TestResponse {
    if (json == null) {
        return json;
    }
    if (instanceOfTestA(json)) {
        return TestAFromJSONTyped(json, true);
    }
    if (instanceOfTestB(json)) {
        return TestBFromJSONTyped(json, true);
    }

    return {} as any;
}

export function TestResponseToJSON(value?: TestResponse | null): any {
    if (value == null) {
        return value;
    }

    if (instanceOfTestA(value)) {
        return TestAToJSON(value as TestA);
    }
    if (instanceOfTestB(value)) {
        return TestBToJSON(value as TestB);
    }

    return {};
}
