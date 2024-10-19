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
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.openapitools.client.model.DeprecatedObject;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;


import org.openapitools.client.ApiClient;
/**
 * ObjectWithDeprecatedFields
 */
@JsonPropertyOrder({
  ObjectWithDeprecatedFields.JSON_PROPERTY_UUID,
  ObjectWithDeprecatedFields.JSON_PROPERTY_ID,
  ObjectWithDeprecatedFields.JSON_PROPERTY_DEPRECATED_REF,
  ObjectWithDeprecatedFields.JSON_PROPERTY_BARS
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", comments = "Generator version: 7.10.0-SNAPSHOT")
public class ObjectWithDeprecatedFields {
  public static final String JSON_PROPERTY_UUID = "uuid";
  @javax.annotation.Nullable  private String uuid;

  public static final String JSON_PROPERTY_ID = "id";
  @javax.annotation.Nullable  private BigDecimal id;

  public static final String JSON_PROPERTY_DEPRECATED_REF = "deprecatedRef";
  @javax.annotation.Nullable  private DeprecatedObject deprecatedRef;

  public static final String JSON_PROPERTY_BARS = "bars";
  @javax.annotation.Nullable  private List<String> bars = new ArrayList<>();

  public ObjectWithDeprecatedFields() { 
  }

  public ObjectWithDeprecatedFields uuid(@javax.annotation.Nullable String uuid) {
    this.uuid = uuid;
    return this;
  }

  /**
   * Get uuid
   * @return uuid
   */
  @javax.annotation.Nullable  @JsonProperty(JSON_PROPERTY_UUID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public String getUuid() {
    return uuid;
  }


  @JsonProperty(JSON_PROPERTY_UUID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setUuid(@javax.annotation.Nullable String uuid) {
    this.uuid = uuid;
  }


  public ObjectWithDeprecatedFields id(@javax.annotation.Nullable BigDecimal id) {
    this.id = id;
    return this;
  }

  /**
   * Get id
   * @return id
   * @deprecated
   */
  @Deprecated
  @javax.annotation.Nullable  @JsonProperty(JSON_PROPERTY_ID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public BigDecimal getId() {
    return id;
  }


  @JsonProperty(JSON_PROPERTY_ID)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setId(@javax.annotation.Nullable BigDecimal id) {
    this.id = id;
  }


  public ObjectWithDeprecatedFields deprecatedRef(@javax.annotation.Nullable DeprecatedObject deprecatedRef) {
    this.deprecatedRef = deprecatedRef;
    return this;
  }

  /**
   * Get deprecatedRef
   * @return deprecatedRef
   * @deprecated
   */
  @Deprecated
  @javax.annotation.Nullable  @JsonProperty(JSON_PROPERTY_DEPRECATED_REF)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public DeprecatedObject getDeprecatedRef() {
    return deprecatedRef;
  }


  @JsonProperty(JSON_PROPERTY_DEPRECATED_REF)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDeprecatedRef(@javax.annotation.Nullable DeprecatedObject deprecatedRef) {
    this.deprecatedRef = deprecatedRef;
  }


  public ObjectWithDeprecatedFields bars(@javax.annotation.Nullable List<String> bars) {
    this.bars = bars;
    return this;
  }

  public ObjectWithDeprecatedFields addBarsItem(String barsItem) {
    if (this.bars == null) {
      this.bars = new ArrayList<>();
    }
    this.bars.add(barsItem);
    return this;
  }

  /**
   * Get bars
   * @return bars
   * @deprecated
   */
  @Deprecated
  @javax.annotation.Nullable  @JsonProperty(JSON_PROPERTY_BARS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public List<String> getBars() {
    return bars;
  }


  @JsonProperty(JSON_PROPERTY_BARS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setBars(@javax.annotation.Nullable List<String> bars) {
    this.bars = bars;
  }

  /**
   * Return true if this ObjectWithDeprecatedFields object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ObjectWithDeprecatedFields objectWithDeprecatedFields = (ObjectWithDeprecatedFields) o;
    return Objects.equals(this.uuid, objectWithDeprecatedFields.uuid) &&
        Objects.equals(this.id, objectWithDeprecatedFields.id) &&
        Objects.equals(this.deprecatedRef, objectWithDeprecatedFields.deprecatedRef) &&
        Objects.equals(this.bars, objectWithDeprecatedFields.bars);
  }

  @Override
  public int hashCode() {
    return Objects.hash(uuid, id, deprecatedRef, bars);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ObjectWithDeprecatedFields {\n");
    sb.append("    uuid: ").append(toIndentedString(uuid)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    deprecatedRef: ").append(toIndentedString(deprecatedRef)).append("\n");
    sb.append("    bars: ").append(toIndentedString(bars)).append("\n");
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

    // add `uuid` to the URL query string
    if (getUuid() != null) {
      joiner.add(String.format("%suuid%s=%s", prefix, suffix, URLEncoder.encode(ApiClient.valueToString(getUuid()), StandardCharsets.UTF_8).replaceAll("\\+", "%20")));
    }

    // add `id` to the URL query string
    if (getId() != null) {
      joiner.add(String.format("%sid%s=%s", prefix, suffix, URLEncoder.encode(ApiClient.valueToString(getId()), StandardCharsets.UTF_8).replaceAll("\\+", "%20")));
    }

    // add `deprecatedRef` to the URL query string
    if (getDeprecatedRef() != null) {
      joiner.add(getDeprecatedRef().toUrlQueryString(prefix + "deprecatedRef" + suffix));
    }

    // add `bars` to the URL query string
    if (getBars() != null) {
      for (int i = 0; i < getBars().size(); i++) {
        joiner.add(String.format("%sbars%s%s=%s", prefix, suffix,
            "".equals(suffix) ? "" : String.format("%s%d%s", containerPrefix, i, containerSuffix),
            URLEncoder.encode(ApiClient.valueToString(getBars().get(i)), StandardCharsets.UTF_8).replaceAll("\\+", "%20")));
      }
    }

    return joiner.toString();
  }

    public static class Builder {

    private ObjectWithDeprecatedFields instance;

    public Builder() {
      this(new ObjectWithDeprecatedFields());
    }

    protected Builder(ObjectWithDeprecatedFields instance) {
      this.instance = instance;
    }

    public ObjectWithDeprecatedFields.Builder uuid(String uuid) {
      this.instance.uuid = uuid;
      return this;
    }
    public ObjectWithDeprecatedFields.Builder id(BigDecimal id) {
      this.instance.id = id;
      return this;
    }
    public ObjectWithDeprecatedFields.Builder deprecatedRef(DeprecatedObject deprecatedRef) {
      this.instance.deprecatedRef = deprecatedRef;
      return this;
    }
    public ObjectWithDeprecatedFields.Builder bars(List<String> bars) {
      this.instance.bars = bars;
      return this;
    }


    /**
    * returns a built ObjectWithDeprecatedFields instance.
    *
    * The builder is not reusable.
    */
    public ObjectWithDeprecatedFields build() {
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
  public static ObjectWithDeprecatedFields.Builder builder() {
    return new ObjectWithDeprecatedFields.Builder();
  }

  /**
  * Create a builder with a shallow copy of this instance.
  */
  public ObjectWithDeprecatedFields.Builder toBuilder() {
    return new ObjectWithDeprecatedFields.Builder()
      .uuid(getUuid())
      .id(getId())
      .deprecatedRef(getDeprecatedRef())
      .bars(getBars());
  }
}

