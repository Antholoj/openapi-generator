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
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Org.OpenAPITools.Client.OpenAPIDateConverter;
using OpenAPIClientUtils = Org.OpenAPITools.Client.ClientUtils;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// ZeroBasedEnumClass
    /// </summary>
    [DataContract(Name = "ZeroBasedEnumClass")]
    public partial class ZeroBasedEnumClass : IEquatable<ZeroBasedEnumClass>, IValidatableObject
    {
        /// <summary>
        /// Defines ZeroBasedEnum
        /// </summary>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum ZeroBasedEnumEnum
        {
            /// <summary>
            /// Enum Unknown for value: unknown
            /// </summary>
            [EnumMember(Value = "unknown")]
            Unknown,

            /// <summary>
            /// Enum NotUnknown for value: notUnknown
            /// </summary>
            [EnumMember(Value = "notUnknown")]
            NotUnknown
        }

        /// <summary>
        /// Gets or Sets ZeroBasedEnum
        /// </summary>
        [DataMember(Name = "ZeroBasedEnum", EmitDefaultValue = false)]
        public ZeroBasedEnumEnum? ZeroBasedEnum { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="ZeroBasedEnumClass" /> class.
        /// </summary>
        /// <param name="zeroBasedEnum">zeroBasedEnum.</param>
        public ZeroBasedEnumClass(ZeroBasedEnumEnum? zeroBasedEnum = default(ZeroBasedEnumEnum?))
        {
            this.ZeroBasedEnum = zeroBasedEnum;
        }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class ZeroBasedEnumClass {\n");
            sb.Append("  ZeroBasedEnum: ").Append(ZeroBasedEnum).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input as ZeroBasedEnumClass).AreEqual;
        }

        /// <summary>
        /// Returns true if ZeroBasedEnumClass instances are equal
        /// </summary>
        /// <param name="input">Instance of ZeroBasedEnumClass to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ZeroBasedEnumClass input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input).AreEqual;
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                hashCode = (hashCode * 59) + this.ZeroBasedEnum.GetHashCode();
                return hashCode;
            }
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

}
