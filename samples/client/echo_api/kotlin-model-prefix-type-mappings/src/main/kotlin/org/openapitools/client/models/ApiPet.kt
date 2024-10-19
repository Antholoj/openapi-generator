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

import org.openapitools.client.models.ApiCategory
import org.openapitools.client.models.ApiTag

import com.google.gson.annotations.SerializedName

/**
 * 
 *
 * @param name 
 * @param photoUrls 
 * @param id 
 * @param category 
 * @param tags 
 * @param status pet status in the store
 */

data class ApiPet (

    @SerializedName("name")
    val name: kotlin.String,

    @SerializedName("photoUrls")
    val photoUrls: kotlin.collections.List<kotlin.String>,

    @SerializedName("id")
    val id: kotlin.Long? = null,

    @SerializedName("category")
    val category: ApiCategory? = null,

    @SerializedName("tags")
    val tags: kotlin.collections.List<ApiTag>? = null,

    /* pet status in the store */
    @SerializedName("status")
    val status: ApiPet.Status? = null

) {

    /**
     * pet status in the store
     *
     * Values: AVAILABLE,PENDING,SOLD
     */
    enum class Status(val value: kotlin.String) {
        @SerializedName(value = "available") AVAILABLE("available"),
        @SerializedName(value = "pending") PENDING("pending"),
        @SerializedName(value = "sold") SOLD("sold");
    }

}

