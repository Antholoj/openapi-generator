/*
 * test
 * test
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import org.openapitools.client.JSON;


/**
 * Parent
 */
@JsonPropertyOrder({
  Parent.JSON_PROPERTY_OBJECT_TYPE
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", comments = "Generator version: 7.10.0-SNAPSHOT")

@JsonIgnoreProperties(
  value = "objectType", // ignore manually set objectType, it will be automatically generated by Jackson during serialization
  allowSetters = true // allows the objectType to be set during deserialization
)
@JsonTypeInfo(use = JsonTypeInfo.Id.NAME, include = JsonTypeInfo.As.PROPERTY, property = "objectType", visible = true)
@JsonSubTypes({
  @JsonSubTypes.Type(value = ChildSchema.class, name = "ChildSchema"),
  @JsonSubTypes.Type(value = MySchemaNameCharacters.class, name = "MySchemaName._-Characters"),
})
public class Parent {
  public static final String JSON_PROPERTY_OBJECT_TYPE = "objectType";
  @javax.annotation.Nullable
  private String objectType;

  public Parent() { 
  }

  public Parent objectType(@javax.annotation.Nullable String objectType) {
    this.objectType = objectType;
    return this;
  }

  /**
   * Get objectType
   * @return objectType
   */
  @javax.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_OBJECT_TYPE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getObjectType() {
    return objectType;
  }


  @JsonProperty(JSON_PROPERTY_OBJECT_TYPE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setObjectType(@javax.annotation.Nullable String objectType) {
    this.objectType = objectType;
  }

  /**
   * Return true if this Parent object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Parent parent = (Parent) o;
    return Objects.equals(this.objectType, parent.objectType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(objectType);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Parent {\n");
    sb.append("    objectType: ").append(toIndentedString(objectType)).append("\n");
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

  static {
    // Initialize and register the discriminator mappings.
    Map<String, Class<?>> mappings = new HashMap<>();
    mappings.put("ChildSchema", ChildSchema.class);
    mappings.put("MySchemaName._-Characters", MySchemaNameCharacters.class);
    mappings.put("Parent", Parent.class);
    JSON.registerDiscriminator(Parent.class, "objectType", mappings);
  }
}

