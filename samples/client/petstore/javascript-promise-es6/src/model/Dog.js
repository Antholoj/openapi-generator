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
import Animal from './Animal';

/**
 * The Dog model module.
 * @module model/Dog
 * @version 1.0.0
 */
class Dog {
    /**
     * Constructs a new <code>Dog</code>.
     * @alias module:model/Dog
     * @extends module:model/Animal
     * @implements module:model/Animal
     * @param className {String} 
     */
    constructor(className) { 
        Animal.initialize(this, className);
        Dog.initialize(this, className);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, className) { 
    }

    /**
     * Constructs a <code>Dog</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Dog} obj Optional instance to populate.
     * @return {module:model/Dog} The populated <code>Dog</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Dog();
            Animal.constructFromObject(data, obj);
            Animal.constructFromObject(data, obj);

            if (data.hasOwnProperty('breed')) {
                obj['breed'] = ApiClient.convertToType(data['breed'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>Dog</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>Dog</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of Dog.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['breed'] && !(typeof data['breed'] === 'string' || data['breed'] instanceof String)) {
            throw new Error("Expected the field `breed` to be a primitive type in the JSON string but got " + data['breed']);
        }

        return true;
    }


}

Dog.RequiredProperties = ["className"];

/**
 * @member {String} breed
 */
Dog.prototype['breed'] = undefined;


// Implement Animal interface:
/**
 * @member {String} className
 */
Animal.prototype['className'] = undefined;
/**
 * @member {String} color
 * @default 'red'
 */
Animal.prototype['color'] = 'red';




export default Dog;
