/**
 * Petstore API
 * API for managing pets in a pet store
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package org.openapitools.client.model;

import java.lang.reflect.Type;
import javax.json.bind.annotation.JsonbTypeDeserializer;
import javax.json.bind.annotation.JsonbTypeSerializer;
import javax.json.bind.serializer.DeserializationContext;
import javax.json.bind.serializer.JsonbDeserializer;
import javax.json.bind.serializer.JsonbSerializer;
import javax.json.bind.serializer.SerializationContext;
import javax.json.stream.JsonGenerator;
import javax.json.stream.JsonParser;
import javax.json.bind.annotation.JsonbProperty;


/**
 * Gets or Sets Status
 */
@JsonbTypeSerializer(Status.Serializer.class)
@JsonbTypeDeserializer(Status.Deserializer.class)
public enum Status {
  
  AVAILABLE("available"),
  
  PENDING("pending"),
  
  SOLD("sold");

  private String value;

  Status(String value) {
    this.value = value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static final class Deserializer implements JsonbDeserializer<Status> {
    @Override
    public Status deserialize(JsonParser parser, DeserializationContext ctx, Type rtType) {
      return fromValue(parser.getString());
    }
  }

  public static final class Serializer implements JsonbSerializer<Status> {
    @Override
    public void serialize(Status obj, JsonGenerator generator, SerializationContext ctx) {
      generator.write(obj.value);
    }
  }

  public static Status fromValue(String text) {
    for (Status b : Status.values()) {
      if (String.valueOf(b.value).equals(text)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + text + "'");
  }
}
