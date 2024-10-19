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
package org.openapitools.database.models

import org.ktorm.dsl.*
import org.ktorm.schema.*
import org.ktorm.database.Database
import org.openapitools.client.models.*


/**
 * A User who is purchasing from the pet store
 * @param id 
 * @param username 
 * @param firstName 
 * @param lastName 
 * @param email 
 * @param password 
 * @param phone 
 * @param userStatus User Status
 */
object Users : BaseTable<User>("User") {
    val id = long("id") /* null */
    val username = text("username") /* null */
    val firstName = text("firstName") /* null */
    val lastName = text("lastName") /* null */
    val email = text("email") /* null */
    val password = text("password") /* null */
    val phone = text("phone") /* null */
    val userStatus = int("userStatus") /* null */ /* User Status */

    /**
     * Create an entity of type User from the model
     */
    override fun doCreateEntity(row: QueryRowSet, withReferences: Boolean) = User(
        id = row[id]  /* kotlin.Long? */,
        username = row[username]  /* kotlin.String? */,
        firstName = row[firstName]  /* kotlin.String? */,
        lastName = row[lastName]  /* kotlin.String? */,
        email = row[email]  /* kotlin.String? */,
        password = row[password]  /* kotlin.String? */,
        phone = row[phone]  /* kotlin.String? */,
        userStatus = row[userStatus]  /* kotlin.Int? */ /* User Status */
    )

    /**
    * Assign all the columns from the entity of type User to the DML expression.
    *
    * Usage:
    *
    * ```kotlin
    * let entity = User()
    * database.update(Users, {
    *     assignFrom(entity)
    * })
    * ```
    * @return the builder with the columns for the update or insert.
    */
    fun AssignmentsBuilder.assignFrom(entity: User) {
        this.apply {
            set(Users.id, entity.id)            set(Users.username, entity.username)            set(Users.firstName, entity.firstName)            set(Users.lastName, entity.lastName)            set(Users.email, entity.email)            set(Users.password, entity.password)            set(Users.phone, entity.phone)            set(Users.userStatus, entity.userStatus)        }
    }

}


