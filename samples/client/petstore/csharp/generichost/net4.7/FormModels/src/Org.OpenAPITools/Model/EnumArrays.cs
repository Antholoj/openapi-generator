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
    /// EnumArrays
    /// </summary>
    public partial class EnumArrays : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="EnumArrays" /> class.
        /// </summary>
        /// <param name="arrayEnum">arrayEnum</param>
        /// <param name="justSymbol">justSymbol</param>
        [JsonConstructor]
        public EnumArrays(Option<List<EnumArraysArrayEnumInner>> arrayEnum = default, Option<EnumArraysJustSymbol?> justSymbol = default)
        {
            ArrayEnumOption = arrayEnum;
            JustSymbolOption = justSymbol;
            OnCreated();
        }

        partial void OnCreated();

        /// <summary>
        /// Used to track the state of JustSymbol
        /// </summary>
        [JsonIgnore]
        [global::System.ComponentModel.EditorBrowsable(global::System.ComponentModel.EditorBrowsableState.Never)]
        public Option<EnumArraysJustSymbol?> JustSymbolOption { get; private set; }

        /// <summary>
        /// Gets or Sets JustSymbol
        /// </summary>
        [JsonPropertyName("just_symbol")]
        public EnumArraysJustSymbol? JustSymbol { get { return this.JustSymbolOption; } set { this.JustSymbolOption = new Option<EnumArraysJustSymbol?>(value); } }

        /// <summary>
        /// Used to track the state of ArrayEnum
        /// </summary>
        [JsonIgnore]
        [global::System.ComponentModel.EditorBrowsable(global::System.ComponentModel.EditorBrowsableState.Never)]
        public Option<List<EnumArraysArrayEnumInner>> ArrayEnumOption { get; private set; }

        /// <summary>
        /// Gets or Sets ArrayEnum
        /// </summary>
        [JsonPropertyName("array_enum")]
        public List<EnumArraysArrayEnumInner> ArrayEnum { get { return this.ArrayEnumOption; } set { this.ArrayEnumOption = new Option<List<EnumArraysArrayEnumInner>>(value); } }

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
            sb.Append("class EnumArrays {\n");
            sb.Append("  ArrayEnum: ").Append(ArrayEnum).Append("\n");
            sb.Append("  JustSymbol: ").Append(JustSymbol).Append("\n");
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
    /// A Json converter for type <see cref="EnumArrays" />
    /// </summary>
    public class EnumArraysJsonConverter : JsonConverter<EnumArrays>
    {
        /// <summary>
        /// Deserializes json to <see cref="EnumArrays" />
        /// </summary>
        /// <param name="utf8JsonReader"></param>
        /// <param name="typeToConvert"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <returns></returns>
        /// <exception cref="JsonException"></exception>
        public override EnumArrays Read(ref Utf8JsonReader utf8JsonReader, Type typeToConvert, JsonSerializerOptions jsonSerializerOptions)
        {
            int currentDepth = utf8JsonReader.CurrentDepth;

            if (utf8JsonReader.TokenType != JsonTokenType.StartObject && utf8JsonReader.TokenType != JsonTokenType.StartArray)
                throw new JsonException();

            JsonTokenType startingTokenType = utf8JsonReader.TokenType;

            Option<List<EnumArraysArrayEnumInner>> arrayEnum = default;
            Option<EnumArraysJustSymbol?> justSymbol = default;

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
                        case "array_enum":
                            if (utf8JsonReader.TokenType != JsonTokenType.Null)
                                arrayEnum = new Option<List<EnumArraysArrayEnumInner>>(JsonSerializer.Deserialize<List<EnumArraysArrayEnumInner>>(ref utf8JsonReader, jsonSerializerOptions));
                            break;
                        case "just_symbol":
                            string justSymbolRawValue = utf8JsonReader.GetString();
                            if (justSymbolRawValue != null)
                                justSymbol = new Option<EnumArraysJustSymbol?>(EnumArraysJustSymbolValueConverter.FromStringOrDefault(justSymbolRawValue));
                            break;
                        default:
                            break;
                    }
                }
            }

            if (arrayEnum.IsSet && arrayEnum.Value == null)
                throw new ArgumentNullException(nameof(arrayEnum), "Property is not nullable for class EnumArrays.");

            if (justSymbol.IsSet && justSymbol.Value == null)
                throw new ArgumentNullException(nameof(justSymbol), "Property is not nullable for class EnumArrays.");

            return new EnumArrays(arrayEnum, justSymbol);
        }

        /// <summary>
        /// Serializes a <see cref="EnumArrays" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="enumArrays"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public override void Write(Utf8JsonWriter writer, EnumArrays enumArrays, JsonSerializerOptions jsonSerializerOptions)
        {
            writer.WriteStartObject();

            WriteProperties(writer, enumArrays, jsonSerializerOptions);
            writer.WriteEndObject();
        }

        /// <summary>
        /// Serializes the properties of <see cref="EnumArrays" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="enumArrays"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public void WriteProperties(Utf8JsonWriter writer, EnumArrays enumArrays, JsonSerializerOptions jsonSerializerOptions)
        {
            if (enumArrays.ArrayEnumOption.IsSet && enumArrays.ArrayEnum == null)
                throw new ArgumentNullException(nameof(enumArrays.ArrayEnum), "Property is required for class EnumArrays.");

            if (enumArrays.ArrayEnumOption.IsSet)
            {
                writer.WritePropertyName("array_enum");
                JsonSerializer.Serialize(writer, enumArrays.ArrayEnum, jsonSerializerOptions);
            }
            if (enumArrays.JustSymbolOption.IsSet)
            {
                var justSymbolRawValue = EnumArraysJustSymbolValueConverter.ToJsonValue(enumArrays.JustSymbol.Value);
                writer.WriteString("just_symbol", justSymbolRawValue);
            }
        }
    }
}
