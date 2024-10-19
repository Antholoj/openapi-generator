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
    /// ComplexQuadrilateral
    /// </summary>
    public partial class ComplexQuadrilateral : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ComplexQuadrilateral" /> class.
        /// </summary>
        /// <param name="quadrilateralType">quadrilateralType</param>
        /// <param name="shapeType">shapeType</param>
        [JsonConstructor]
        public ComplexQuadrilateral(string quadrilateralType, string shapeType)
        {
            QuadrilateralType = quadrilateralType;
            ShapeType = shapeType;
            OnCreated();
        }

        partial void OnCreated();

        /// <summary>
        /// Gets or Sets QuadrilateralType
        /// </summary>
        [JsonPropertyName("quadrilateralType")]
        public string QuadrilateralType { get; set; }

        /// <summary>
        /// Gets or Sets ShapeType
        /// </summary>
        [JsonPropertyName("shapeType")]
        public string ShapeType { get; set; }

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
            sb.Append("class ComplexQuadrilateral {\n");
            sb.Append("  QuadrilateralType: ").Append(QuadrilateralType).Append("\n");
            sb.Append("  ShapeType: ").Append(ShapeType).Append("\n");
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
    /// A Json converter for type <see cref="ComplexQuadrilateral" />
    /// </summary>
    public class ComplexQuadrilateralJsonConverter : JsonConverter<ComplexQuadrilateral>
    {
        /// <summary>
        /// Deserializes json to <see cref="ComplexQuadrilateral" />
        /// </summary>
        /// <param name="utf8JsonReader"></param>
        /// <param name="typeToConvert"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <returns></returns>
        /// <exception cref="JsonException"></exception>
        public override ComplexQuadrilateral Read(ref Utf8JsonReader utf8JsonReader, Type typeToConvert, JsonSerializerOptions jsonSerializerOptions)
        {
            int currentDepth = utf8JsonReader.CurrentDepth;

            if (utf8JsonReader.TokenType != JsonTokenType.StartObject && utf8JsonReader.TokenType != JsonTokenType.StartArray)
                throw new JsonException();

            JsonTokenType startingTokenType = utf8JsonReader.TokenType;

            Option<string> quadrilateralType = default;
            Option<string> shapeType = default;

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
                        case "quadrilateralType":
                            quadrilateralType = new Option<string>(utf8JsonReader.GetString());
                            break;
                        case "shapeType":
                            shapeType = new Option<string>(utf8JsonReader.GetString());
                            break;
                        default:
                            break;
                    }
                }
            }

            if (!quadrilateralType.IsSet)
                throw new ArgumentException("Property is required for class ComplexQuadrilateral.", nameof(quadrilateralType));

            if (!shapeType.IsSet)
                throw new ArgumentException("Property is required for class ComplexQuadrilateral.", nameof(shapeType));

            if (quadrilateralType.IsSet && quadrilateralType.Value == null)
                throw new ArgumentNullException(nameof(quadrilateralType), "Property is not nullable for class ComplexQuadrilateral.");

            if (shapeType.IsSet && shapeType.Value == null)
                throw new ArgumentNullException(nameof(shapeType), "Property is not nullable for class ComplexQuadrilateral.");

            return new ComplexQuadrilateral(quadrilateralType.Value, shapeType.Value);
        }

        /// <summary>
        /// Serializes a <see cref="ComplexQuadrilateral" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="complexQuadrilateral"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public override void Write(Utf8JsonWriter writer, ComplexQuadrilateral complexQuadrilateral, JsonSerializerOptions jsonSerializerOptions)
        {
            writer.WriteStartObject();

            WriteProperties(writer, complexQuadrilateral, jsonSerializerOptions);
            writer.WriteEndObject();
        }

        /// <summary>
        /// Serializes the properties of <see cref="ComplexQuadrilateral" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="complexQuadrilateral"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public void WriteProperties(Utf8JsonWriter writer, ComplexQuadrilateral complexQuadrilateral, JsonSerializerOptions jsonSerializerOptions)
        {
            if (complexQuadrilateral.QuadrilateralType == null)
                throw new ArgumentNullException(nameof(complexQuadrilateral.QuadrilateralType), "Property is required for class ComplexQuadrilateral.");

            if (complexQuadrilateral.ShapeType == null)
                throw new ArgumentNullException(nameof(complexQuadrilateral.ShapeType), "Property is required for class ComplexQuadrilateral.");

                                    writer.WriteString("quadrilateralType", complexQuadrilateral.QuadrilateralType);
                                                            writer.WriteString("shapeType", complexQuadrilateral.ShapeType);
        }
    }
}
