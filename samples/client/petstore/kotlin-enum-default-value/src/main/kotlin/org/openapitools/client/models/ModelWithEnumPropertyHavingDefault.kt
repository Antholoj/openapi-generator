/**
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models


import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass
import java.io.Serializable

/**
 * 
 *
 * @param propertyName 
 */

data class ModelWithEnumPropertyHavingDefault (

    @Json(name = "propertyName")
    val propertyName: ModelWithEnumPropertyHavingDefault.PropertyName = PropertyName.VALUE

) : Serializable {
    companion object {
        private const val serialVersionUID: Long = 123
    }

    /**
     * 
     *
     * Values: VALUE,unknown_default_open_api
     */
    @JsonClass(generateAdapter = false)
    enum class PropertyName(val value: kotlin.String) {
        @Json(name = "VALUE") VALUE("VALUE"),
        @Json(name = "unknown_default_open_api") unknown_default_open_api("unknown_default_open_api");
    }

}

