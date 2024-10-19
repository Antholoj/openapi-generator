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


/**
 * A User who is purchasing from the pet store
 * @export
 * @interface User
 */
export interface User {
    /**
     * 
     * @type {number}
     * @memberof User
     */
    'id'?: number;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    'username'?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    'firstName'?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    'lastName'?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    'email'?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    'password'?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    'phone'?: string;
    /**
     * User Status
     * @type {number}
     * @memberof User
     */
    'userStatus'?: number;
}

