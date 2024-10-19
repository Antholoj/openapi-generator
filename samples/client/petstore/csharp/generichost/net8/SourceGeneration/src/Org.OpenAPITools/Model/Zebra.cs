// <auto-generated>
/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */
#nullable enable

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
using System.Text.Json.Serialization.Metadata;
using Org.OpenAPITools.Client;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// Zebra
    /// </summary>
    public partial class Zebra : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="Zebra" /> class.
        /// </summary>
        /// <param name="className">className</param>
        /// <param name="type">type</param>
        [JsonConstructor]
        public Zebra(string className, Option<TypeEnum?> type = default)
        {
            ClassName = className;
            TypeOption = type;
            OnCreated();
        }

        partial void OnCreated();

        /// <summary>
        /// Defines Type
        /// </summary>
        public enum TypeEnum
        {
            /// <summary>
            /// Enum Plains for value: plains
            /// </summary>
            Plains = 1,

            /// <summary>
            /// Enum Mountain for value: mountain
            /// </summary>
            Mountain = 2,

            /// <summary>
            /// Enum Grevys for value: grevys
            /// </summary>
            Grevys = 3
        }

        /// <summary>
        /// Returns a <see cref="TypeEnum"/>
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        /// <exception cref="NotImplementedException"></exception>
        public static TypeEnum TypeEnumFromString(string value)
        {
            if (value.Equals("plains"))
                return TypeEnum.Plains;

            if (value.Equals("mountain"))
                return TypeEnum.Mountain;

            if (value.Equals("grevys"))
                return TypeEnum.Grevys;

            throw new NotImplementedException($"Could not convert value to type TypeEnum: '{value}'");
        }

        /// <summary>
        /// Returns a <see cref="TypeEnum"/>
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        public static TypeEnum? TypeEnumFromStringOrDefault(string value)
        {
            if (value.Equals("plains"))
                return TypeEnum.Plains;

            if (value.Equals("mountain"))
                return TypeEnum.Mountain;

            if (value.Equals("grevys"))
                return TypeEnum.Grevys;

            return null;
        }

        /// <summary>
        /// Converts the <see cref="TypeEnum"/> to the json value
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        /// <exception cref="NotImplementedException"></exception>
        public static string TypeEnumToJsonValue(TypeEnum? value)
        {
            if (value == TypeEnum.Plains)
                return "plains";

            if (value == TypeEnum.Mountain)
                return "mountain";

            if (value == TypeEnum.Grevys)
                return "grevys";

            throw new NotImplementedException($"Value could not be handled: '{value}'");
        }
        /// <summary>
        /// Used to track the state of Type
        /// </summary>
        [JsonIgnore]
        [global::System.ComponentModel.EditorBrowsable(global::System.ComponentModel.EditorBrowsableState.Never)]
        public Option<TypeEnum?> TypeOption { get; private set; }

        /// <summary>
        /// Gets or Sets Type
        /// </summary>
        [JsonPropertyName("type")]
        public TypeEnum? Type { get { return this.TypeOption; } set { this.TypeOption = new(value); } }

        /// <summary>
        /// Gets or Sets ClassName
        /// </summary>
        [JsonPropertyName("className")]
        public string ClassName { get; set; }

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
            sb.Append("class Zebra {\n");
            sb.Append("  ClassName: ").Append(ClassName).Append("\n");
            sb.Append("  Type: ").Append(Type).Append("\n");
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
    /// A Json converter for type <see cref="Zebra" />
    /// </summary>
    public class ZebraJsonConverter : JsonConverter<Zebra>
    {
        /// <summary>
        /// Deserializes json to <see cref="Zebra" />
        /// </summary>
        /// <param name="utf8JsonReader"></param>
        /// <param name="typeToConvert"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <returns></returns>
        /// <exception cref="JsonException"></exception>
        public override Zebra Read(ref Utf8JsonReader utf8JsonReader, Type typeToConvert, JsonSerializerOptions jsonSerializerOptions)
        {
            int currentDepth = utf8JsonReader.CurrentDepth;

            if (utf8JsonReader.TokenType != JsonTokenType.StartObject && utf8JsonReader.TokenType != JsonTokenType.StartArray)
                throw new JsonException();

            JsonTokenType startingTokenType = utf8JsonReader.TokenType;

            Option<string?> className = default;
            Option<Zebra.TypeEnum?> type = default;

            while (utf8JsonReader.Read())
            {
                if (startingTokenType == JsonTokenType.StartObject && utf8JsonReader.TokenType == JsonTokenType.EndObject && currentDepth == utf8JsonReader.CurrentDepth)
                    break;

                if (startingTokenType == JsonTokenType.StartArray && utf8JsonReader.TokenType == JsonTokenType.EndArray && currentDepth == utf8JsonReader.CurrentDepth)
                    break;

                if (utf8JsonReader.TokenType == JsonTokenType.PropertyName && currentDepth == utf8JsonReader.CurrentDepth - 1)
                {
                    string? localVarJsonPropertyName = utf8JsonReader.GetString();
                    utf8JsonReader.Read();

                    switch (localVarJsonPropertyName)
                    {
                        case "className":
                            className = new Option<string?>(utf8JsonReader.GetString()!);
                            break;
                        case "type":
                            string? typeRawValue = utf8JsonReader.GetString();
                            if (typeRawValue != null)
                                type = new Option<Zebra.TypeEnum?>(Zebra.TypeEnumFromStringOrDefault(typeRawValue));
                            break;
                        default:
                            break;
                    }
                }
            }

            if (!className.IsSet)
                throw new ArgumentException("Property is required for class Zebra.", nameof(className));

            if (className.IsSet && className.Value == null)
                throw new ArgumentNullException(nameof(className), "Property is not nullable for class Zebra.");

            if (type.IsSet && type.Value == null)
                throw new ArgumentNullException(nameof(type), "Property is not nullable for class Zebra.");

            return new Zebra(className.Value!, type);
        }

        /// <summary>
        /// Serializes a <see cref="Zebra" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="zebra"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public override void Write(Utf8JsonWriter writer, Zebra zebra, JsonSerializerOptions jsonSerializerOptions)
        {
            writer.WriteStartObject();

            WriteProperties(writer, zebra, jsonSerializerOptions);
            writer.WriteEndObject();
        }

        /// <summary>
        /// Serializes the properties of <see cref="Zebra" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="zebra"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public void WriteProperties(Utf8JsonWriter writer, Zebra zebra, JsonSerializerOptions jsonSerializerOptions)
        {
            if (zebra.ClassName == null)
                throw new ArgumentNullException(nameof(zebra.ClassName), "Property is required for class Zebra.");

                                    writer.WriteString("className", zebra.ClassName);
                                    var typeRawValue = Zebra.TypeEnumToJsonValue(zebra.TypeOption.Value!.Value);
            writer.WriteString("type", typeRawValue);
        }
    }
    /// <summary>
    /// The ZebraSerializationContext
    /// </summary>
    [JsonSourceGenerationOptions(WriteIndented = true, GenerationMode = JsonSourceGenerationMode.Metadata | JsonSourceGenerationMode.Serialization)]
    [JsonSerializable(typeof(Zebra))]
    public partial class ZebraSerializationContext : JsonSerializerContext { }
}
