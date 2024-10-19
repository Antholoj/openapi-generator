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
package org.openapitools.server.apis

import com.google.gson.Gson
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.auth.*
import io.ktor.server.response.*
import org.openapitools.server.Paths
import io.ktor.server.resources.options
import io.ktor.server.resources.get
import io.ktor.server.resources.post
import io.ktor.server.resources.put
import io.ktor.server.resources.delete
import io.ktor.server.resources.head
import io.ktor.server.resources.patch
import io.ktor.server.routing.*
import org.openapitools.server.infrastructure.ApiPrincipal
import org.openapitools.server.models.ModelApiResponse
import org.openapitools.server.models.Pet

fun Route.PetApi() {
    val gson = Gson()
    val empty = mutableMapOf<String, Any?>()

    authenticate("petstore_auth") {
    post<Paths.addPet> {
        
        val principal = call.authentication.principal<OAuthAccessTokenResponse>()!!
        
        call.respond(HttpStatusCode.NotImplemented)
        
    }
    }

    authenticate("petstore_auth") {
    delete<Paths.deletePet> {
        
        val principal = call.authentication.principal<OAuthAccessTokenResponse>()!!
        
        call.respond(HttpStatusCode.NotImplemented)
        
    }
    }

    authenticate("petstore_auth") {
    get<Paths.findPetsByStatus> {
        
        val principal = call.authentication.principal<OAuthAccessTokenResponse>()!!
        
        val exampleContentType = "application/json"
            val exampleContentString = """[ {
              "photoUrls" : [ "photoUrls", "photoUrls" ],
              "name" : "doggie",
              "id" : 0,
              "category" : {
                "name" : "name",
                "id" : 6
              },
              "tags" : [ {
                "name" : "name",
                "id" : 1
              }, {
                "name" : "name",
                "id" : 1
              } ],
              "status" : "available"
            }, {
              "photoUrls" : [ "photoUrls", "photoUrls" ],
              "name" : "doggie",
              "id" : 0,
              "category" : {
                "name" : "name",
                "id" : 6
              },
              "tags" : [ {
                "name" : "name",
                "id" : 1
              }, {
                "name" : "name",
                "id" : 1
              } ],
              "status" : "available"
            } ]"""
            
            when (exampleContentType) {
                "application/json" -> call.respond(gson.fromJson(exampleContentString, empty::class.java))
                "application/xml" -> call.respondText(exampleContentString, ContentType.Text.Xml)
                else -> call.respondText(exampleContentString)
            }
        
    }
    }

    authenticate("petstore_auth") {
    get<Paths.findPetsByTags> {
        
        val principal = call.authentication.principal<OAuthAccessTokenResponse>()!!
        
        val exampleContentType = "application/json"
            val exampleContentString = """[ {
              "photoUrls" : [ "photoUrls", "photoUrls" ],
              "name" : "doggie",
              "id" : 0,
              "category" : {
                "name" : "name",
                "id" : 6
              },
              "tags" : [ {
                "name" : "name",
                "id" : 1
              }, {
                "name" : "name",
                "id" : 1
              } ],
              "status" : "available"
            }, {
              "photoUrls" : [ "photoUrls", "photoUrls" ],
              "name" : "doggie",
              "id" : 0,
              "category" : {
                "name" : "name",
                "id" : 6
              },
              "tags" : [ {
                "name" : "name",
                "id" : 1
              }, {
                "name" : "name",
                "id" : 1
              } ],
              "status" : "available"
            } ]"""
            
            when (exampleContentType) {
                "application/json" -> call.respond(gson.fromJson(exampleContentString, empty::class.java))
                "application/xml" -> call.respondText(exampleContentString, ContentType.Text.Xml)
                else -> call.respondText(exampleContentString)
            }
        
    }
    }

    authenticate("api_key") {
    get<Paths.getPetById> {
        
        val principal = call.authentication.principal<ApiPrincipal>()!!
        
        val exampleContentType = "application/json"
            val exampleContentString = """{
              "photoUrls" : [ "photoUrls", "photoUrls" ],
              "name" : "doggie",
              "id" : 0,
              "category" : {
                "name" : "name",
                "id" : 6
              },
              "tags" : [ {
                "name" : "name",
                "id" : 1
              }, {
                "name" : "name",
                "id" : 1
              } ],
              "status" : "available"
            }"""
            
            when (exampleContentType) {
                "application/json" -> call.respond(gson.fromJson(exampleContentString, empty::class.java))
                "application/xml" -> call.respondText(exampleContentString, ContentType.Text.Xml)
                else -> call.respondText(exampleContentString)
            }
        
    }
    }

    authenticate("petstore_auth") {
    put<Paths.updatePet> {
        
        val principal = call.authentication.principal<OAuthAccessTokenResponse>()!!
        
        call.respond(HttpStatusCode.NotImplemented)
        
    }
    }

    authenticate("petstore_auth") {
    post<Paths.updatePetWithForm> {
        
        val principal = call.authentication.principal<OAuthAccessTokenResponse>()!!
        
        call.respond(HttpStatusCode.NotImplemented)
        
    }
    }

    authenticate("petstore_auth") {
    post<Paths.uploadFile> {
        
        val principal = call.authentication.principal<OAuthAccessTokenResponse>()!!
        
        val exampleContentType = "application/json"
            val exampleContentString = """{
              "code" : 0,
              "type" : "type",
              "message" : "message"
            }"""
            
            when (exampleContentType) {
                "application/json" -> call.respond(gson.fromJson(exampleContentString, empty::class.java))
                "application/xml" -> call.respondText(exampleContentString, ContentType.Text.Xml)
                else -> call.respondText(exampleContentString)
            }
        
    }
    }

}
