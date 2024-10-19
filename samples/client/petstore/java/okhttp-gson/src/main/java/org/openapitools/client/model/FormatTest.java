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
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.File;
import java.io.IOException;
import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.UUID;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import com.google.gson.TypeAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * FormatTest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", comments = "Generator version: 7.10.0-SNAPSHOT")

public class FormatTest {
  public static final String SERIALIZED_NAME_INTEGER = "integer";
  @SerializedName(SERIALIZED_NAME_INTEGER)
  @javax.annotation.Nullable  private Integer integer;

  public static final String SERIALIZED_NAME_INT32 = "int32";
  @SerializedName(SERIALIZED_NAME_INT32)
  @javax.annotation.Nullable  private Integer int32;

  public static final String SERIALIZED_NAME_INT64 = "int64";
  @SerializedName(SERIALIZED_NAME_INT64)
  @javax.annotation.Nullable  private Long int64;

  public static final String SERIALIZED_NAME_NUMBER = "number";
  @SerializedName(SERIALIZED_NAME_NUMBER)
  @javax.annotation.Nonnull  private BigDecimal number;

  public static final String SERIALIZED_NAME_FLOAT = "float";
  @SerializedName(SERIALIZED_NAME_FLOAT)
  @javax.annotation.Nullable  private Float _float;

  public static final String SERIALIZED_NAME_DOUBLE = "double";
  @SerializedName(SERIALIZED_NAME_DOUBLE)
  @javax.annotation.Nullable  private Double _double;

  public static final String SERIALIZED_NAME_DECIMAL = "decimal";
  @SerializedName(SERIALIZED_NAME_DECIMAL)
  @javax.annotation.Nullable  private BigDecimal decimal;

  public static final String SERIALIZED_NAME_STRING = "string";
  @SerializedName(SERIALIZED_NAME_STRING)
  @javax.annotation.Nullable  private String string;

  public static final String SERIALIZED_NAME_BYTE = "byte";
  @SerializedName(SERIALIZED_NAME_BYTE)
  @javax.annotation.Nonnull  private byte[] _byte;

  public static final String SERIALIZED_NAME_BINARY = "binary";
  @SerializedName(SERIALIZED_NAME_BINARY)
  @javax.annotation.Nullable  private File binary;

  public static final String SERIALIZED_NAME_DATE = "date";
  @SerializedName(SERIALIZED_NAME_DATE)
  @javax.annotation.Nonnull  private LocalDate date;

  public static final String SERIALIZED_NAME_DATE_TIME = "dateTime";
  @SerializedName(SERIALIZED_NAME_DATE_TIME)
  @javax.annotation.Nullable  private OffsetDateTime dateTime;

  public static final String SERIALIZED_NAME_UUID = "uuid";
  @SerializedName(SERIALIZED_NAME_UUID)
  @javax.annotation.Nullable  private UUID uuid;

  public static final String SERIALIZED_NAME_UUID_WITH_DEFAULT = "uuid_with_default";
  @SerializedName(SERIALIZED_NAME_UUID_WITH_DEFAULT)
  @javax.annotation.Nullable  private UUID uuidWithDefault = UUID.fromString("11111111-206d-4f12-9f12-3d1e525a8e84");

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  @javax.annotation.Nonnull  private String password;

  public static final String SERIALIZED_NAME_PATTERN_WITH_DIGITS = "pattern_with_digits";
  @SerializedName(SERIALIZED_NAME_PATTERN_WITH_DIGITS)
  @javax.annotation.Nullable  private String patternWithDigits;

  public static final String SERIALIZED_NAME_PATTERN_WITH_DIGITS_AND_DELIMITER = "pattern_with_digits_and_delimiter";
  @SerializedName(SERIALIZED_NAME_PATTERN_WITH_DIGITS_AND_DELIMITER)
  @javax.annotation.Nullable  private String patternWithDigitsAndDelimiter;

  public FormatTest() {
  }

  public FormatTest integer(@javax.annotation.Nullable Integer integer) {
    this.integer = integer;
    return this;
  }

  /**
   * Get integer
   * minimum: 10
   * maximum: 100
   * @return integer
   */
  @javax.annotation.Nullable  public Integer getInteger() {
    return integer;
  }

  public void setInteger(@javax.annotation.Nullable Integer integer) {
    this.integer = integer;
  }


  public FormatTest int32(@javax.annotation.Nullable Integer int32) {
    this.int32 = int32;
    return this;
  }

  /**
   * Get int32
   * minimum: 20
   * maximum: 200
   * @return int32
   */
  @javax.annotation.Nullable  public Integer getInt32() {
    return int32;
  }

  public void setInt32(@javax.annotation.Nullable Integer int32) {
    this.int32 = int32;
  }


  public FormatTest int64(@javax.annotation.Nullable Long int64) {
    this.int64 = int64;
    return this;
  }

  /**
   * Get int64
   * @return int64
   */
  @javax.annotation.Nullable  public Long getInt64() {
    return int64;
  }

  public void setInt64(@javax.annotation.Nullable Long int64) {
    this.int64 = int64;
  }


  public FormatTest number(@javax.annotation.Nonnull BigDecimal number) {
    this.number = number;
    return this;
  }

  /**
   * Get number
   * minimum: 32.1
   * maximum: 543.2
   * @return number
   */
  @javax.annotation.Nonnull  public BigDecimal getNumber() {
    return number;
  }

  public void setNumber(@javax.annotation.Nonnull BigDecimal number) {
    this.number = number;
  }


  public FormatTest _float(@javax.annotation.Nullable Float _float) {
    this._float = _float;
    return this;
  }

  /**
   * Get _float
   * minimum: 54.3
   * maximum: 987.6
   * @return _float
   */
  @javax.annotation.Nullable  public Float getFloat() {
    return _float;
  }

  public void setFloat(@javax.annotation.Nullable Float _float) {
    this._float = _float;
  }


  public FormatTest _double(@javax.annotation.Nullable Double _double) {
    this._double = _double;
    return this;
  }

  /**
   * Get _double
   * minimum: 67.8
   * maximum: 123.4
   * @return _double
   */
  @javax.annotation.Nullable  public Double getDouble() {
    return _double;
  }

  public void setDouble(@javax.annotation.Nullable Double _double) {
    this._double = _double;
  }


  public FormatTest decimal(@javax.annotation.Nullable BigDecimal decimal) {
    this.decimal = decimal;
    return this;
  }

  /**
   * Get decimal
   * @return decimal
   */
  @javax.annotation.Nullable  public BigDecimal getDecimal() {
    return decimal;
  }

  public void setDecimal(@javax.annotation.Nullable BigDecimal decimal) {
    this.decimal = decimal;
  }


  public FormatTest string(@javax.annotation.Nullable String string) {
    this.string = string;
    return this;
  }

  /**
   * Get string
   * @return string
   */
  @javax.annotation.Nullable  public String getString() {
    return string;
  }

  public void setString(@javax.annotation.Nullable String string) {
    this.string = string;
  }


  public FormatTest _byte(@javax.annotation.Nonnull byte[] _byte) {
    this._byte = _byte;
    return this;
  }

  /**
   * Get _byte
   * @return _byte
   */
  @javax.annotation.Nonnull  public byte[] getByte() {
    return _byte;
  }

  public void setByte(@javax.annotation.Nonnull byte[] _byte) {
    this._byte = _byte;
  }


  public FormatTest binary(@javax.annotation.Nullable File binary) {
    this.binary = binary;
    return this;
  }

  /**
   * Get binary
   * @return binary
   */
  @javax.annotation.Nullable  public File getBinary() {
    return binary;
  }

  public void setBinary(@javax.annotation.Nullable File binary) {
    this.binary = binary;
  }


  public FormatTest date(@javax.annotation.Nonnull LocalDate date) {
    this.date = date;
    return this;
  }

  /**
   * Get date
   * @return date
   */
  @javax.annotation.Nonnull  public LocalDate getDate() {
    return date;
  }

  public void setDate(@javax.annotation.Nonnull LocalDate date) {
    this.date = date;
  }


  public FormatTest dateTime(@javax.annotation.Nullable OffsetDateTime dateTime) {
    this.dateTime = dateTime;
    return this;
  }

  /**
   * Get dateTime
   * @return dateTime
   */
  @javax.annotation.Nullable  public OffsetDateTime getDateTime() {
    return dateTime;
  }

  public void setDateTime(@javax.annotation.Nullable OffsetDateTime dateTime) {
    this.dateTime = dateTime;
  }


  public FormatTest uuid(@javax.annotation.Nullable UUID uuid) {
    this.uuid = uuid;
    return this;
  }

  /**
   * Get uuid
   * @return uuid
   */
  @javax.annotation.Nullable  public UUID getUuid() {
    return uuid;
  }

  public void setUuid(@javax.annotation.Nullable UUID uuid) {
    this.uuid = uuid;
  }


  public FormatTest uuidWithDefault(@javax.annotation.Nullable UUID uuidWithDefault) {
    this.uuidWithDefault = uuidWithDefault;
    return this;
  }

  /**
   * Get uuidWithDefault
   * @return uuidWithDefault
   */
  @javax.annotation.Nullable  public UUID getUuidWithDefault() {
    return uuidWithDefault;
  }

  public void setUuidWithDefault(@javax.annotation.Nullable UUID uuidWithDefault) {
    this.uuidWithDefault = uuidWithDefault;
  }


  public FormatTest password(@javax.annotation.Nonnull String password) {
    this.password = password;
    return this;
  }

  /**
   * Get password
   * @return password
   */
  @javax.annotation.Nonnull  public String getPassword() {
    return password;
  }

  public void setPassword(@javax.annotation.Nonnull String password) {
    this.password = password;
  }


  public FormatTest patternWithDigits(@javax.annotation.Nullable String patternWithDigits) {
    this.patternWithDigits = patternWithDigits;
    return this;
  }

  /**
   * A string that is a 10 digit number. Can have leading zeros.
   * @return patternWithDigits
   */
  @javax.annotation.Nullable  public String getPatternWithDigits() {
    return patternWithDigits;
  }

  public void setPatternWithDigits(@javax.annotation.Nullable String patternWithDigits) {
    this.patternWithDigits = patternWithDigits;
  }


  public FormatTest patternWithDigitsAndDelimiter(@javax.annotation.Nullable String patternWithDigitsAndDelimiter) {
    this.patternWithDigitsAndDelimiter = patternWithDigitsAndDelimiter;
    return this;
  }

  /**
   * A string starting with &#39;image_&#39; (case insensitive) and one to three digits following i.e. Image_01.
   * @return patternWithDigitsAndDelimiter
   */
  @javax.annotation.Nullable  public String getPatternWithDigitsAndDelimiter() {
    return patternWithDigitsAndDelimiter;
  }

  public void setPatternWithDigitsAndDelimiter(@javax.annotation.Nullable String patternWithDigitsAndDelimiter) {
    this.patternWithDigitsAndDelimiter = patternWithDigitsAndDelimiter;
  }

  /**
   * A container for additional, undeclared properties.
   * This is a holder for any undeclared properties as specified with
   * the 'additionalProperties' keyword in the OAS document.
   */
  private Map<String, Object> additionalProperties;

  /**
   * Set the additional (undeclared) property with the specified name and value.
   * If the property does not already exist, create it otherwise replace it.
   *
   * @param key name of the property
   * @param value value of the property
   * @return the FormatTest instance itself
   */
  public FormatTest putAdditionalProperty(String key, Object value) {
    if (this.additionalProperties == null) {
        this.additionalProperties = new HashMap<String, Object>();
    }
    this.additionalProperties.put(key, value);
    return this;
  }

  /**
   * Return the additional (undeclared) property.
   *
   * @return a map of objects
   */
  public Map<String, Object> getAdditionalProperties() {
    return additionalProperties;
  }

  /**
   * Return the additional (undeclared) property with the specified name.
   *
   * @param key name of the property
   * @return an object
   */
  public Object getAdditionalProperty(String key) {
    if (this.additionalProperties == null) {
        return null;
    }
    return this.additionalProperties.get(key);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    FormatTest formatTest = (FormatTest) o;
    return Objects.equals(this.integer, formatTest.integer) &&
        Objects.equals(this.int32, formatTest.int32) &&
        Objects.equals(this.int64, formatTest.int64) &&
        Objects.equals(this.number, formatTest.number) &&
        Objects.equals(this._float, formatTest._float) &&
        Objects.equals(this._double, formatTest._double) &&
        Objects.equals(this.decimal, formatTest.decimal) &&
        Objects.equals(this.string, formatTest.string) &&
        Arrays.equals(this._byte, formatTest._byte) &&
        Objects.equals(this.binary, formatTest.binary) &&
        Objects.equals(this.date, formatTest.date) &&
        Objects.equals(this.dateTime, formatTest.dateTime) &&
        Objects.equals(this.uuid, formatTest.uuid) &&
        Objects.equals(this.uuidWithDefault, formatTest.uuidWithDefault) &&
        Objects.equals(this.password, formatTest.password) &&
        Objects.equals(this.patternWithDigits, formatTest.patternWithDigits) &&
        Objects.equals(this.patternWithDigitsAndDelimiter, formatTest.patternWithDigitsAndDelimiter)&&
        Objects.equals(this.additionalProperties, formatTest.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(integer, int32, int64, number, _float, _double, decimal, string, Arrays.hashCode(_byte), binary, date, dateTime, uuid, uuidWithDefault, password, patternWithDigits, patternWithDigitsAndDelimiter, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class FormatTest {\n");
    sb.append("    integer: ").append(toIndentedString(integer)).append("\n");
    sb.append("    int32: ").append(toIndentedString(int32)).append("\n");
    sb.append("    int64: ").append(toIndentedString(int64)).append("\n");
    sb.append("    number: ").append(toIndentedString(number)).append("\n");
    sb.append("    _float: ").append(toIndentedString(_float)).append("\n");
    sb.append("    _double: ").append(toIndentedString(_double)).append("\n");
    sb.append("    decimal: ").append(toIndentedString(decimal)).append("\n");
    sb.append("    string: ").append(toIndentedString(string)).append("\n");
    sb.append("    _byte: ").append(toIndentedString(_byte)).append("\n");
    sb.append("    binary: ").append(toIndentedString(binary)).append("\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    dateTime: ").append(toIndentedString(dateTime)).append("\n");
    sb.append("    uuid: ").append(toIndentedString(uuid)).append("\n");
    sb.append("    uuidWithDefault: ").append(toIndentedString(uuidWithDefault)).append("\n");
    sb.append("    password: ").append("*").append("\n");
    sb.append("    patternWithDigits: ").append(toIndentedString(patternWithDigits)).append("\n");
    sb.append("    patternWithDigitsAndDelimiter: ").append(toIndentedString(patternWithDigitsAndDelimiter)).append("\n");
    sb.append("    additionalProperties: ").append(toIndentedString(additionalProperties)).append("\n");
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


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("integer");
    openapiFields.add("int32");
    openapiFields.add("int64");
    openapiFields.add("number");
    openapiFields.add("float");
    openapiFields.add("double");
    openapiFields.add("decimal");
    openapiFields.add("string");
    openapiFields.add("byte");
    openapiFields.add("binary");
    openapiFields.add("date");
    openapiFields.add("dateTime");
    openapiFields.add("uuid");
    openapiFields.add("uuid_with_default");
    openapiFields.add("password");
    openapiFields.add("pattern_with_digits");
    openapiFields.add("pattern_with_digits_and_delimiter");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("number");
    openapiRequiredFields.add("byte");
    openapiRequiredFields.add("date");
    openapiRequiredFields.add("password");
  }

  /**
   * Validates the JSON Element and throws an exception if issues found
   *
   * @param jsonElement JSON Element
   * @throws IOException if the JSON Element is invalid with respect to FormatTest
   */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!FormatTest.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in FormatTest is not found in the empty JSON string", FormatTest.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : FormatTest.openapiRequiredFields) {
        if (jsonElement.getAsJsonObject().get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("string") != null && !jsonObj.get("string").isJsonNull()) && !jsonObj.get("string").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `string` to be a primitive type in the JSON string but got `%s`", jsonObj.get("string").toString()));
      }
      if ((jsonObj.get("uuid") != null && !jsonObj.get("uuid").isJsonNull()) && !jsonObj.get("uuid").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `uuid` to be a primitive type in the JSON string but got `%s`", jsonObj.get("uuid").toString()));
      }
      if ((jsonObj.get("uuid_with_default") != null && !jsonObj.get("uuid_with_default").isJsonNull()) && !jsonObj.get("uuid_with_default").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `uuid_with_default` to be a primitive type in the JSON string but got `%s`", jsonObj.get("uuid_with_default").toString()));
      }
      if (!jsonObj.get("password").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `password` to be a primitive type in the JSON string but got `%s`", jsonObj.get("password").toString()));
      }
      if ((jsonObj.get("pattern_with_digits") != null && !jsonObj.get("pattern_with_digits").isJsonNull()) && !jsonObj.get("pattern_with_digits").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `pattern_with_digits` to be a primitive type in the JSON string but got `%s`", jsonObj.get("pattern_with_digits").toString()));
      }
      if ((jsonObj.get("pattern_with_digits_and_delimiter") != null && !jsonObj.get("pattern_with_digits_and_delimiter").isJsonNull()) && !jsonObj.get("pattern_with_digits_and_delimiter").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `pattern_with_digits_and_delimiter` to be a primitive type in the JSON string but got `%s`", jsonObj.get("pattern_with_digits_and_delimiter").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!FormatTest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'FormatTest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<FormatTest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(FormatTest.class));

       return (TypeAdapter<T>) new TypeAdapter<FormatTest>() {
           @Override
           public void write(JsonWriter out, FormatTest value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             obj.remove("additionalProperties");
             // serialize additional properties
             if (value.getAdditionalProperties() != null) {
               for (Map.Entry<String, Object> entry : value.getAdditionalProperties().entrySet()) {
                 if (entry.getValue() instanceof String)
                   obj.addProperty(entry.getKey(), (String) entry.getValue());
                 else if (entry.getValue() instanceof Number)
                   obj.addProperty(entry.getKey(), (Number) entry.getValue());
                 else if (entry.getValue() instanceof Boolean)
                   obj.addProperty(entry.getKey(), (Boolean) entry.getValue());
                 else if (entry.getValue() instanceof Character)
                   obj.addProperty(entry.getKey(), (Character) entry.getValue());
                 else {
                   JsonElement jsonElement = gson.toJsonTree(entry.getValue());
                   if (jsonElement.isJsonArray()) {
                     obj.add(entry.getKey(), jsonElement.getAsJsonArray());
                   } else {
                     obj.add(entry.getKey(), jsonElement.getAsJsonObject());
                   }
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public FormatTest read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             JsonObject jsonObj = jsonElement.getAsJsonObject();
             // store additional fields in the deserialized instance
             FormatTest instance = thisAdapter.fromJsonTree(jsonObj);
             for (Map.Entry<String, JsonElement> entry : jsonObj.entrySet()) {
               if (!openapiFields.contains(entry.getKey())) {
                 if (entry.getValue().isJsonPrimitive()) { // primitive type
                   if (entry.getValue().getAsJsonPrimitive().isString())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsString());
                   else if (entry.getValue().getAsJsonPrimitive().isNumber())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsNumber());
                   else if (entry.getValue().getAsJsonPrimitive().isBoolean())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsBoolean());
                   else
                     throw new IllegalArgumentException(String.format("The field `%s` has unknown primitive type. Value: %s", entry.getKey(), entry.getValue().toString()));
                 } else if (entry.getValue().isJsonArray()) {
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), List.class));
                 } else { // JSON object
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), HashMap.class));
                 }
               }
             }
             return instance;
           }

       }.nullSafe();
    }
  }

  /**
   * Create an instance of FormatTest given an JSON string
   *
   * @param jsonString JSON string
   * @return An instance of FormatTest
   * @throws IOException if the JSON string is invalid with respect to FormatTest
   */
  public static FormatTest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, FormatTest.class);
  }

  /**
   * Convert an instance of FormatTest to an JSON string
   *
   * @return JSON string
   */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

