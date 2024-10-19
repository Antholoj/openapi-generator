/*
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

package org.openapitools.model;

import java.util.Objects;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.*;

import javax.validation.constraints.*;
import javax.validation.Valid;
import io.micronaut.core.annotation.*;
import javax.annotation.Generated;

/**
 * ModelClient
 */
@JsonPropertyOrder({
  ModelClient.JSON_PROPERTY_CLIENT
})
@JsonTypeName("Client")
@Generated(value="org.openapitools.codegen.languages.JavaMicronautClientCodegen", comments = "Generator version: 7.10.0-SNAPSHOT")
@Introspected
public class ModelClient {
    public static final String JSON_PROPERTY_CLIENT = "client";
    private String _client;

    public ModelClient() {
    }
    public ModelClient _client(String _client) {
        this._client = _client;
        return this;
    }

    /**
     * Get _client
     * @return _client
     */
    @Nullable
    @JsonProperty(JSON_PROPERTY_CLIENT)
    @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
    public String getClient() {
        return _client;
    }

    @JsonProperty(JSON_PROPERTY_CLIENT)
    @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
    public void setClient(String _client) {
        this._client = _client;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        ModelClient _client = (ModelClient) o;
        return Objects.equals(this._client, _client._client);
    }

    @Override
    public int hashCode() {
        return Objects.hash(_client);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("class ModelClient {\n");
        sb.append("    _client: ").append(toIndentedString(_client)).append("\n");
        sb.append("}");
        return sb.toString();
    }

    /**
     * Convert the given object to string with each line indented by 4 spaces
     * (except the first line).
     */
    private String toIndentedString(Object o) {
        if (o == null) {
            return "null";
        }
        return o.toString().replace("\n", "\n    ");
    }

}
