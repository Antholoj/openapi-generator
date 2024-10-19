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
    /// Defines Enum_Test_enum_string
    /// </summary>
    public enum EnumTestEnumString
    {
        /// <summary>
        /// Enum UPPER for value: UPPER
        /// </summary>
        UPPER = 1,

        /// <summary>
        /// Enum Lower for value: lower
        /// </summary>
        Lower = 2,

        /// <summary>
        /// Enum Empty for value: 
        /// </summary>
        Empty = 3,

        /// <summary>
        /// Enum ValuewithTab for value: Value\twith tab
        /// </summary>
        ValuewithTab = 4,

        /// <summary>
        /// Enum ValueWithQuote for value: Value with \&quot; quote
        /// </summary>
        ValueWithQuote = 5,

        /// <summary>
        /// Enum ValueWithEscapedQuote for value: Value with escaped \&quot; quote
        /// </summary>
        ValueWithEscapedQuote = 6,

        /// <summary>
        /// Enum Duplicatevalue for value: Duplicate\nvalue
        /// </summary>
        Duplicatevalue = 7,

        /// <summary>
        /// Enum Duplicatevalue2 for value: Duplicate\r\nvalue
        /// </summary>
        Duplicatevalue2 = 8
    }

    /// <summary>
    /// Converts <see cref="EnumTestEnumString"/> to and from the JSON value
    /// </summary>
    public static class EnumTestEnumStringValueConverter
    {
        /// <summary>
        /// Parses a given value to <see cref="EnumTestEnumString"/>
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        public static EnumTestEnumString FromString(string value)
        {
            if (value.Equals("UPPER"))
                return EnumTestEnumString.UPPER;

            if (value.Equals("lower"))
                return EnumTestEnumString.Lower;

            if (value.Equals(""))
                return EnumTestEnumString.Empty;

            if (value.Equals("Value\twith tab"))
                return EnumTestEnumString.ValuewithTab;

            if (value.Equals("Value with \" quote"))
                return EnumTestEnumString.ValueWithQuote;

            if (value.Equals("Value with escaped \" quote"))
                return EnumTestEnumString.ValueWithEscapedQuote;

            if (value.Equals("Duplicate\nvalue"))
                return EnumTestEnumString.Duplicatevalue;

            if (value.Equals("Duplicate\r\nvalue"))
                return EnumTestEnumString.Duplicatevalue2;

            throw new NotImplementedException($"Could not convert value to type EnumTestEnumString: '{value}'");
        }

        /// <summary>
        /// Parses a given value to <see cref="EnumTestEnumString"/>
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        public static EnumTestEnumString? FromStringOrDefault(string value)
        {
            if (value.Equals("UPPER"))
                return EnumTestEnumString.UPPER;

            if (value.Equals("lower"))
                return EnumTestEnumString.Lower;

            if (value.Equals(""))
                return EnumTestEnumString.Empty;

            if (value.Equals("Value\twith tab"))
                return EnumTestEnumString.ValuewithTab;

            if (value.Equals("Value with \" quote"))
                return EnumTestEnumString.ValueWithQuote;

            if (value.Equals("Value with escaped \" quote"))
                return EnumTestEnumString.ValueWithEscapedQuote;

            if (value.Equals("Duplicate\nvalue"))
                return EnumTestEnumString.Duplicatevalue;

            if (value.Equals("Duplicate\r\nvalue"))
                return EnumTestEnumString.Duplicatevalue2;

            return null;
        }

        /// <summary>
        /// Converts the <see cref="EnumTestEnumString"/> to the json value
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        /// <exception cref="NotImplementedException"></exception>
        public static string ToJsonValue(EnumTestEnumString value)
        {
            if (value == EnumTestEnumString.UPPER)
                return "UPPER";

            if (value == EnumTestEnumString.Lower)
                return "lower";

            if (value == EnumTestEnumString.Empty)
                return "";

            if (value == EnumTestEnumString.ValuewithTab)
                return "Value\twith tab";

            if (value == EnumTestEnumString.ValueWithQuote)
                return "Value with \" quote";

            if (value == EnumTestEnumString.ValueWithEscapedQuote)
                return "Value with escaped \" quote";

            if (value == EnumTestEnumString.Duplicatevalue)
                return "Duplicate\nvalue";

            if (value == EnumTestEnumString.Duplicatevalue2)
                return "Duplicate\r\nvalue";

            throw new NotImplementedException($"Value could not be handled: '{value}'");
        }
    }

    /// <summary>
    /// A Json converter for type <see cref="EnumTestEnumString"/>
    /// </summary>
    /// <exception cref="NotImplementedException"></exception>
    public class EnumTestEnumStringJsonConverter : JsonConverter<EnumTestEnumString>
    {
        /// <summary>
        /// Returns a  from the Json object
        /// </summary>
        /// <param name="reader"></param>
        /// <param name="typeToConvert"></param>
        /// <param name="options"></param>
        /// <returns></returns>
        public override EnumTestEnumString Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        {
            string rawValue = reader.GetString();

            EnumTestEnumString? result = rawValue == null
                ? null
                : EnumTestEnumStringValueConverter.FromStringOrDefault(rawValue);

            if (result != null)
                return result.Value;

            throw new JsonException();
        }

        /// <summary>
        /// Writes the EnumTestEnumString to the json writer
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="enumTestEnumString"></param>
        /// <param name="options"></param>
        public override void Write(Utf8JsonWriter writer, EnumTestEnumString enumTestEnumString, JsonSerializerOptions options)
        {
            writer.WriteStringValue(enumTestEnumString.ToString());
        }
    }

    /// <summary>
    /// A Json converter for type <see cref="EnumTestEnumString"/>
    /// </summary>
    public class EnumTestEnumStringNullableJsonConverter : JsonConverter<EnumTestEnumString?>
    {
        /// <summary>
        /// Returns a EnumTestEnumString from the Json object
        /// </summary>
        /// <param name="reader"></param>
        /// <param name="typeToConvert"></param>
        /// <param name="options"></param>
        /// <returns></returns>
        public override EnumTestEnumString? Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        {
            string rawValue = reader.GetString();

            EnumTestEnumString? result = rawValue == null
                ? null
                : EnumTestEnumStringValueConverter.FromStringOrDefault(rawValue);

            if (result != null)
                return result.Value;

            throw new JsonException();
        }

        /// <summary>
        /// Writes the DateTime to the json writer
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="enumTestEnumString"></param>
        /// <param name="options"></param>
        public override void Write(Utf8JsonWriter writer, EnumTestEnumString? enumTestEnumString, JsonSerializerOptions options)
        {
            writer.WriteStringValue(enumTestEnumString?.ToString() ?? "null");
        }
    }
}
