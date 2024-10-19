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
 *
 */

import ApiClient from '../ApiClient';
/**
* Enum class OuterEnumInteger.
* @enum {}
* @readonly
*/
export default class OuterEnumInteger {
    
        /**
         * value: 0
         * @const
         */
        "0" = 0;

    
        /**
         * value: 1
         * @const
         */
        "1" = 1;

    
        /**
         * value: 2
         * @const
         */
        "2" = 2;

    

    /**
    * Returns a <code>OuterEnumInteger</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/OuterEnumInteger} The enum <code>OuterEnumInteger</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}
