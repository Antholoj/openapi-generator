package org.openapitools.model;

import java.net.URI;
import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import org.openapitools.jackson.nullable.JsonNullable;
import java.time.OffsetDateTime;
import javax.validation.Valid;
import javax.validation.constraints.*;
import io.swagger.v3.oas.annotations.media.Schema;


import java.util.*;
import javax.annotation.Generated;

/**
 * Entity
 */
@JsonIgnoreProperties(
  value = "@type", // ignore manually set @type, it will be automatically generated by Jackson during serialization
  allowSetters = true // allows the @type to be set during deserialization
)
@JsonTypeInfo(use = JsonTypeInfo.Id.NAME, include = JsonTypeInfo.As.PROPERTY, property = "@type", visible = true)
@JsonSubTypes({
  @JsonSubTypes.Type(value = Bar.class, name = "Bar"),
  @JsonSubTypes.Type(value = BarCreate.class, name = "Bar_Create"),
  @JsonSubTypes.Type(value = Foo.class, name = "Foo"),
  @JsonSubTypes.Type(value = Pasta.class, name = "Pasta"),
  @JsonSubTypes.Type(value = Pizza.class, name = "Pizza"),
  @JsonSubTypes.Type(value = PizzaSpeziale.class, name = "PizzaSpeziale")
})
@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", comments = "Generator version: 7.10.0-SNAPSHOT")
public class Entity {

  private String href;

  private String id;

  private String atSchemaLocation;

  private String atBaseType;

  private String atType;

  public Entity() {
    super();
  }

  /**
   * Constructor with only required parameters
   */
  public Entity(String atType) {
    this.atType = atType;
  }

  public Entity href(String href) {
    this.href = href;
    return this;
  }

  /**
   * Hyperlink reference
   * @return href
   */
    @Schema(name = "href", description = "Hyperlink reference", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("href")
  public String getHref() {
    return href;
  }

  public void setHref(String href) {
    this.href = href;
  }

  public Entity id(String id) {
    this.id = id;
    return this;
  }

  /**
   * unique identifier
   * @return id
   */
    @Schema(name = "id", description = "unique identifier", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("id")
  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }

  public Entity atSchemaLocation(String atSchemaLocation) {
    this.atSchemaLocation = atSchemaLocation;
    return this;
  }

  /**
   * A URI to a JSON-Schema file that defines additional attributes and relationships
   * @return atSchemaLocation
   */
    @Schema(name = "@schemaLocation", description = "A URI to a JSON-Schema file that defines additional attributes and relationships", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("@schemaLocation")
  public String getAtSchemaLocation() {
    return atSchemaLocation;
  }

  public void setAtSchemaLocation(String atSchemaLocation) {
    this.atSchemaLocation = atSchemaLocation;
  }

  public Entity atBaseType(String atBaseType) {
    this.atBaseType = atBaseType;
    return this;
  }

  /**
   * When sub-classing, this defines the super-class
   * @return atBaseType
   */
    @Schema(name = "@baseType", description = "When sub-classing, this defines the super-class", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("@baseType")
  public String getAtBaseType() {
    return atBaseType;
  }

  public void setAtBaseType(String atBaseType) {
    this.atBaseType = atBaseType;
  }

  public Entity atType(String atType) {
    this.atType = atType;
    return this;
  }

  /**
   * When sub-classing, this defines the sub-class Extensible name
   * @return atType
   */
  @NotNull   @Schema(name = "@type", description = "When sub-classing, this defines the sub-class Extensible name", requiredMode = Schema.RequiredMode.REQUIRED)
  @JsonProperty("@type")
  public String getAtType() {
    return atType;
  }

  public void setAtType(String atType) {
    this.atType = atType;
  }
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Entity entity = (Entity) o;
    return Objects.equals(this.href, entity.href) &&
        Objects.equals(this.id, entity.id) &&
        Objects.equals(this.atSchemaLocation, entity.atSchemaLocation) &&
        Objects.equals(this.atBaseType, entity.atBaseType) &&
        Objects.equals(this.atType, entity.atType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(href, id, atSchemaLocation, atBaseType, atType);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Entity {\n");
    sb.append("    href: ").append(toIndentedString(href)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    atSchemaLocation: ").append(toIndentedString(atSchemaLocation)).append("\n");
    sb.append("    atBaseType: ").append(toIndentedString(atBaseType)).append("\n");
    sb.append("    atType: ").append(toIndentedString(atType)).append("\n");
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
  
    public static class Builder {
  
      private Entity instance;
  
      public Builder() {
        this(new Entity());
      }
  
      protected Builder(Entity instance) {
        this.instance = instance;
      }
  
      protected Builder copyOf(Entity value) { 
        this.instance.setHref(value.href);
        this.instance.setId(value.id);
        this.instance.setAtSchemaLocation(value.atSchemaLocation);
        this.instance.setAtBaseType(value.atBaseType);
        this.instance.setAtType(value.atType);
        return this;
      }
  
        public Entity.Builder href(String href) {
        this.instance.href(href);
        return this;
      }
      
        public Entity.Builder id(String id) {
        this.instance.id(id);
        return this;
      }
      
        public Entity.Builder atSchemaLocation(String atSchemaLocation) {
        this.instance.atSchemaLocation(atSchemaLocation);
        return this;
      }
      
        public Entity.Builder atBaseType(String atBaseType) {
        this.instance.atBaseType(atBaseType);
        return this;
      }
      
        public Entity.Builder atType(String atType) {
        this.instance.atType(atType);
        return this;
      }
      
          /**
      * returns a built Entity instance.
      *
      * The builder is not reusable (NullPointerException)
      */
      public Entity build() {
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
    * Create a builder with no initialized field (except for the default values).
    */
    public static Entity.Builder builder() {
      return new Entity.Builder();
    }
  
    /**
    * Create a builder with a shallow copy of this instance.
    */
    public Entity.Builder toBuilder() {
      Entity.Builder builder = new Entity.Builder();
      return builder.copyOf(this);
    }
}

