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

import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.StringJoiner;
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
import org.openapitools.client.ApiClient;
/**
 * Animal
 */
@JsonPropertyOrder({
  Animal.JSON_PROPERTY_CLASS_NAME,
  Animal.JSON_PROPERTY_COLOR
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", comments = "Generator version: 7.10.0-SNAPSHOT")
@JsonIgnoreProperties(
  value = "className", // ignore manually set className, it will be automatically generated by Jackson during serialization
  allowSetters = true // allows the className to be set during deserialization
)
@JsonTypeInfo(use = JsonTypeInfo.Id.NAME, include = JsonTypeInfo.As.PROPERTY, property = "className", visible = true)
@JsonSubTypes({
  @JsonSubTypes.Type(value = Cat.class, name = "Cat"),
  @JsonSubTypes.Type(value = Dog.class, name = "Dog"),
})
public class Animal {
  public static final String JSON_PROPERTY_CLASS_NAME = "className";
  @javax.annotation.Nonnull  private String className;

  public static final String JSON_PROPERTY_COLOR = "color";
  @javax.annotation.Nullable  private String color = "red";

  public Animal() { 
  }

  public Animal className(@javax.annotation.Nonnull String className) {
    this.className = className;
    return this;
  }

  /**
   * Get className
   * @return className
   */
  @javax.annotation.Nonnull  @JsonProperty(JSON_PROPERTY_CLASS_NAME)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public String getClassName() {
    return className;
  }


  @JsonProperty(JSON_PROPERTY_CLASS_NAME)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setClassName(@javax.annotation.Nonnull String className) {
    this.className = className;
  }


  public Animal color(@javax.annotation.Nullable String color) {
    this.color = color;
    return this;
  }

  /**
   * Get color
   * @return color
   */
  @javax.annotation.Nullable  @JsonProperty(JSON_PROPERTY_COLOR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public String getColor() {
    return color;
  }


  @JsonProperty(JSON_PROPERTY_COLOR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setColor(@javax.annotation.Nullable String color) {
    this.color = color;
  }

  /**
   * Return true if this Animal object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Animal animal = (Animal) o;
    return Objects.equals(this.className, animal.className) &&
        Objects.equals(this.color, animal.color);
  }

  @Override
  public int hashCode() {
    return Objects.hash(className, color);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Animal {\n");
    sb.append("    className: ").append(toIndentedString(className)).append("\n");
    sb.append("    color: ").append(toIndentedString(color)).append("\n");
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

  /**
   * Convert the instance into URL query string.
   *
   * @return URL query string
   */
  public String toUrlQueryString() {
    return toUrlQueryString(null);
  }

  /**
   * Convert the instance into URL query string.
   *
   * @param prefix prefix of the query string
   * @return URL query string
   */
  public String toUrlQueryString(String prefix) {
    String suffix = "";
    String containerSuffix = "";
    String containerPrefix = "";
    if (prefix == null) {
      // style=form, explode=true, e.g. /pet?name=cat&type=manx
      prefix = "";
    } else {
      // deepObject style e.g. /pet?id[name]=cat&id[type]=manx
      prefix = prefix + "[";
      suffix = "]";
      containerSuffix = "]";
      containerPrefix = "[";
    }

    StringJoiner joiner = new StringJoiner("&");

    // add `className` to the URL query string
    if (getClassName() != null) {
      joiner.add(String.format("%sclassName%s=%s", prefix, suffix, URLEncoder.encode(ApiClient.valueToString(getClassName()), StandardCharsets.UTF_8).replaceAll("\\+", "%20")));
    }

    // add `color` to the URL query string
    if (getColor() != null) {
      joiner.add(String.format("%scolor%s=%s", prefix, suffix, URLEncoder.encode(ApiClient.valueToString(getColor()), StandardCharsets.UTF_8).replaceAll("\\+", "%20")));
    }

    return joiner.toString();
  }
static {
  // Initialize and register the discriminator mappings.
  Map<String, Class<?>> mappings = new HashMap<String, Class<?>>();
  mappings.put("Cat", Cat.class);
  mappings.put("Dog", Dog.class);
  mappings.put("Animal", Animal.class);
  JSON.registerDiscriminator(Animal.class, "className", mappings);
}

    public static class Builder {

    private Animal instance;

    public Builder() {
      this(new Animal());
    }

    protected Builder(Animal instance) {
      this.instance = instance;
    }

    public Animal.Builder className(String className) {
      this.instance.className = className;
      return this;
    }
    public Animal.Builder color(String color) {
      this.instance.color = color;
      return this;
    }


    /**
    * returns a built Animal instance.
    *
    * The builder is not reusable.
    */
    public Animal build() {
      try {
        return this.instance;
      } finally {
        // ensure that this.instance is not reused
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
  public static Animal.Builder builder() {
    return new Animal.Builder();
  }

  /**
  * Create a builder with a shallow copy of this instance.
  */
  public Animal.Builder toBuilder() {
    return new Animal.Builder()
      .className(getClassName())
      .color(getColor());
  }
}

