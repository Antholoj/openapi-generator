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
 */

package org.openapitools.client.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import org.openapitools.jackson.nullable.JsonNullable;
import com.fasterxml.jackson.annotation.JsonIgnore;
import java.util.NoSuchElementException;



/**
  * Just a string to inform instance is up and running. Make it nullable in hope to get it as pointer in generated model.
 **/
public class HealthCheckResult  {
  
  private String nullableMessage;

 /**
   * Get nullableMessage
   * @return nullableMessage
  **/
  public String getNullableMessage() {
    return nullableMessage;
  }

  /**
    * Set nullableMessage
  **/
  public void setNullableMessage(String nullableMessage) {
    this.nullableMessage = nullableMessage;
  }

  public HealthCheckResult nullableMessage(String nullableMessage) {
    this.nullableMessage = nullableMessage;
    return this;
  }


  /**
    * Create a string representation of this pojo.
  **/
  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HealthCheckResult {\n");
    
    sb.append("    nullableMessage: ").append(toIndentedString(nullableMessage)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private static String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}
