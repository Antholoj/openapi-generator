// <auto-generated>
/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */
using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Text;
using System.Text.RegularExpressions;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.ComponentModel.DataAnnotations;
using OpenAPIClientUtils = Org.OpenAPITools.Client.ClientUtils;
using Org.OpenAPITools.Client;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// ZeroBasedEnumClass
    /// </summary>
    public partial class ZeroBasedEnumClass : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ZeroBasedEnumClass" /> class.
        /// </summary>
        /// <param name="zeroBasedEnum">zeroBasedEnum</param>
        [JsonConstructor]
        public ZeroBasedEnumClass(Option<ZeroBasedEnumClassZeroBasedEnum?> zeroBasedEnum = default)
        {
            ZeroBasedEnumOption = zeroBasedEnum;
            OnCreated();
        }

        partial void OnCreated();

        /// <summary>
        /// Used to track the state of ZeroBasedEnum
        /// </summary>
        [JsonIgnore]
        [global::System.ComponentModel.EditorBrowsable(global::System.ComponentModel.EditorBrowsableState.Never)]
        public Option<ZeroBasedEnumClassZeroBasedEnum?> ZeroBasedEnumOption { get; private set; }

        /// <summary>
        /// Gets or Sets ZeroBasedEnum
        /// </summary>
        [JsonPropertyName("ZeroBasedEnum")]
        public ZeroBasedEnumClassZeroBasedEnum? ZeroBasedEnum { get { return this.ZeroBasedEnumOption; } set { this.ZeroBasedEnumOption = new Option<ZeroBasedEnumClassZeroBasedEnum?>(value); } }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public Dictionary<string, JsonElement> AdditionalProperties { get; } = new Dictionary<string, JsonElement>();

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class ZeroBasedEnumClass {\n");
            sb.Append("  ZeroBasedEnum: ").Append(ZeroBasedEnum).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }    }
    /// <summary>
    /// A Json converter for type <see cref="ZeroBasedEnumClass" />
    /// </summary>
    public class ZeroBasedEnumClassJsonConverter : JsonConverter<ZeroBasedEnumClass>
    {
        /// <summary>
        /// Deserializes json to <see cref="ZeroBasedEnumClass" />
        /// </summary>
        /// <param name="utf8JsonReader"></param>
        /// <param name="typeToConvert"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <returns></returns>
        /// <exception cref="JsonException"></exception>
        public override ZeroBasedEnumClass Read(ref Utf8JsonReader utf8JsonReader, Type typeToConvert, JsonSerializerOptions jsonSerializerOptions)
        {
            int currentDepth = utf8JsonReader.CurrentDepth;

            if (utf8JsonReader.TokenType != JsonTokenType.StartObject && utf8JsonReader.TokenType != JsonTokenType.StartArray)
                throw new JsonException();

            JsonTokenType startingTokenType = utf8JsonReader.TokenType;

            Option<ZeroBasedEnumClassZeroBasedEnum?> zeroBasedEnum = default;

            while (utf8JsonReader.Read())
            {
                if (startingTokenType == JsonTokenType.StartObject && utf8JsonReader.TokenType == JsonTokenType.EndObject && currentDepth == utf8JsonReader.CurrentDepth)
                    break;

                if (startingTokenType == JsonTokenType.StartArray && utf8JsonReader.TokenType == JsonTokenType.EndArray && currentDepth == utf8JsonReader.CurrentDepth)
                    break;

                if (utf8JsonReader.TokenType == JsonTokenType.PropertyName && currentDepth == utf8JsonReader.CurrentDepth - 1)
                {
                    string localVarJsonPropertyName = utf8JsonReader.GetString();
                    utf8JsonReader.Read();

                    switch (localVarJsonPropertyName)
                    {
                        case "ZeroBasedEnum":
                            string zeroBasedEnumRawValue = utf8JsonReader.GetString();
                            if (zeroBasedEnumRawValue != null)
                                zeroBasedEnum = new Option<ZeroBasedEnumClassZeroBasedEnum?>(ZeroBasedEnumClassZeroBasedEnumValueConverter.FromStringOrDefault(zeroBasedEnumRawValue));
                            break;
                        default:
                            break;
                    }
                }
            }

            if (zeroBasedEnum.IsSet && zeroBasedEnum.Value == null)
                throw new ArgumentNullException(nameof(zeroBasedEnum), "Property is not nullable for class ZeroBasedEnumClass.");

            return new ZeroBasedEnumClass(zeroBasedEnum);
        }

        /// <summary>
        /// Serializes a <see cref="ZeroBasedEnumClass" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="zeroBasedEnumClass"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public override void Write(Utf8JsonWriter writer, ZeroBasedEnumClass zeroBasedEnumClass, JsonSerializerOptions jsonSerializerOptions)
        {
            writer.WriteStartObject();

            WriteProperties(writer, zeroBasedEnumClass, jsonSerializerOptions);
            writer.WriteEndObject();
        }

        /// <summary>
        /// Serializes the properties of <see cref="ZeroBasedEnumClass" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="zeroBasedEnumClass"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public void WriteProperties(Utf8JsonWriter writer, ZeroBasedEnumClass zeroBasedEnumClass, JsonSerializerOptions jsonSerializerOptions)
        {
            if (zeroBasedEnumClass.ZeroBasedEnumOption.IsSet)
            {
                var zeroBasedEnumRawValue = ZeroBasedEnumClassZeroBasedEnumValueConverter.ToJsonValue(zeroBasedEnumClass.ZeroBasedEnum.Value);
                writer.WriteString("ZeroBasedEnum", zeroBasedEnumRawValue);
            }
        }
    }
}
