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
 * The Client model module.
 * @module model/Client
 * @version 1.0.0
 */
class Client {
    /**
     * Constructs a new <code>Client</code>.
     * @alias module:model/Client
     */
    constructor() { 
        
        Client.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Client</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Client} obj Optional instance to populate.
     * @return {module:model/Client} The populated <code>Client</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Client();

            if (data.hasOwnProperty('client')) {
                obj['client'] = ApiClient.convertToType(data['client'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>Client</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>Client</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['client'] && !(typeof data['client'] === 'string' || data['client'] instanceof String)) {
            throw new Error("Expected the field `client` to be a primitive type in the JSON string but got " + data['client']);
        }

        return true;
    }


}



/**
 * @member {String} client
 */
Client.prototype['client'] = undefined;






export default Client;
