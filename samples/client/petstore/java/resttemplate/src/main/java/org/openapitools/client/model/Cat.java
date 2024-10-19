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


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import org.openapitools.client.model.Animal;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

/**
 * Cat
 */
@JsonPropertyOrder({
  Cat.JSON_PROPERTY_DECLAWED
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", comments = "Generator version: 7.10.0-SNAPSHOT")

@JsonIgnoreProperties(
  value = "className", // ignore manually set className, it will be automatically generated by Jackson during serialization
  allowSetters = true // allows the className to be set during deserialization
)
@JsonTypeInfo(use = JsonTypeInfo.Id.NAME, include = JsonTypeInfo.As.PROPERTY, property = "className", visible = true)

public class Cat extends Animal {
  public static final String JSON_PROPERTY_DECLAWED = "declawed";
  @javax.annotation.Nullable  private Boolean declawed;

  public Cat() {

  }

  /**
   * Constructor with all args parameters
   */
  public Cat(@JsonProperty(JSON_PROPERTY_DECLAWED) Boolean declawed, @JsonProperty(JSON_PROPERTY_CLASS_NAME) String className, @JsonProperty(JSON_PROPERTY_COLOR) String color) {
    super(className, color);
    this.declawed = declawed;
  }

  public Cat declawed(@javax.annotation.Nullable Boolean declawed) {
    
    this.declawed = declawed;
    return this;
  }

  /**
   * Get declawed
   * @return declawed
   */
  @javax.annotation.Nullable  @JsonProperty(JSON_PROPERTY_DECLAWED)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Boolean getDeclawed() {
    return declawed;
  }


  @JsonProperty(JSON_PROPERTY_DECLAWED)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDeclawed(@javax.annotation.Nullable Boolean declawed) {
    this.declawed = declawed;
  }

  @Override
  public Cat className(@javax.annotation.Nonnull String className) {
    this.setClassName(className);
    return this;
  }

  @Override
  public Cat color(@javax.annotation.Nullable String color) {
    this.setColor(color);
    return this;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Cat cat = (Cat) o;
    return Objects.equals(this.declawed, cat.declawed) &&
        super.equals(o);
  }

  @Override
  public int hashCode() {
    return Objects.hash(declawed, super.hashCode());
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Cat {\n");
    sb.append("    ").append(toIndentedString(super.toString())).append("\n");
    sb.append("    declawed: ").append(toIndentedString(declawed)).append("\n");
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

  public static class Builder extends Animal.Builder {
  
      private Cat instance;
  
      public Builder() {
        this(new Cat());
      }
  
      protected Builder(Cat instance) {
        super(instance);
        this.instance = instance;
      }
  
      public Cat.Builder declawed(Boolean declawed) {
          this.instance.declawed = declawed;
        return this;
      }
    
      public Cat.Builder className(String className) { // inherited: true
        super.className(className);
        return this;
     }
  
      public Cat.Builder color(String color) { // inherited: true
        super.color(color);
        return this;
     }
  
  
      /**
      * returns a built Cat instance.
      *
      * The builder is not reusable.
      */
      public Cat build() {
        try {
          return this.instance;
        } finally {
          // ensure that this.instance is not reused
          super.build();
          this.instance = null;
        }
      }
  
      @Override
      public String toString() {
        return getClass() + "=(" + instance + ")";
      }
    }
  
    /**
    * Create a builder with no initialized field.
    */
    public static Cat.Builder builder() {
      return new Cat.Builder();
    }
  
    /**
    * Create a builder with a shallow copy of this instance.
    */
    public Cat.Builder toBuilder() {
      return new Cat.Builder()
        .className(getClassName())
        .color(getColor())
        .declawed(getDeclawed());
    }

}

