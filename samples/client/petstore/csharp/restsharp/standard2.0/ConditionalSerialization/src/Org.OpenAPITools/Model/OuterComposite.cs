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
    /// OuterComposite
    /// </summary>
    [DataContract(Name = "OuterComposite")]
    public partial class OuterComposite : IEquatable<OuterComposite>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="OuterComposite" /> class.
        /// </summary>
        /// <param name="myNumber">myNumber.</param>
        /// <param name="myString">myString.</param>
        /// <param name="myBoolean">myBoolean.</param>
        public OuterComposite(decimal myNumber = default(decimal), string myString = default(string), bool myBoolean = default(bool))
        {
            this._MyNumber = myNumber;
            if (this.MyNumber != null)
            {
                this._flagMyNumber = true;
            }
            this._MyString = myString;
            if (this.MyString != null)
            {
                this._flagMyString = true;
            }
            this._MyBoolean = myBoolean;
            if (this.MyBoolean != null)
            {
                this._flagMyBoolean = true;
            }
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets MyNumber
        /// </summary>
        [DataMember(Name = "my_number", EmitDefaultValue = false)]
        public decimal MyNumber
        {
            get{ return _MyNumber;}
            set
            {
                _MyNumber = value;
                _flagMyNumber = true;
            }
        }
        private decimal _MyNumber;
        private bool _flagMyNumber;

        /// <summary>
        /// Returns false as MyNumber should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeMyNumber()
        {
            return _flagMyNumber;
        }
        /// <summary>
        /// Gets or Sets MyString
        /// </summary>
        [DataMember(Name = "my_string", EmitDefaultValue = false)]
        public string MyString
        {
            get{ return _MyString;}
            set
            {
                _MyString = value;
                _flagMyString = true;
            }
        }
        private string _MyString;
        private bool _flagMyString;

        /// <summary>
        /// Returns false as MyString should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeMyString()
        {
            return _flagMyString;
        }
        /// <summary>
        /// Gets or Sets MyBoolean
        /// </summary>
        [DataMember(Name = "my_boolean", EmitDefaultValue = true)]
        public bool MyBoolean
        {
            get{ return _MyBoolean;}
            set
            {
                _MyBoolean = value;
                _flagMyBoolean = true;
            }
        }
        private bool _MyBoolean;
        private bool _flagMyBoolean;

        /// <summary>
        /// Returns false as MyBoolean should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeMyBoolean()
        {
            return _flagMyBoolean;
        }
        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class OuterComposite {\n");
            sb.Append("  MyNumber: ").Append(MyNumber).Append("\n");
            sb.Append("  MyString: ").Append(MyString).Append("\n");
            sb.Append("  MyBoolean: ").Append(MyBoolean).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
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
            return OpenAPIClientUtils.compareLogic.Compare(this, input as OuterComposite).AreEqual;
        }

        /// <summary>
        /// Returns true if OuterComposite instances are equal
        /// </summary>
        /// <param name="input">Instance of OuterComposite to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(OuterComposite input)
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
                hashCode = (hashCode * 59) + this.MyNumber.GetHashCode();
                if (this.MyString != null)
                {
                    hashCode = (hashCode * 59) + this.MyString.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.MyBoolean.GetHashCode();
                if (this.AdditionalProperties != null)
                {
                    hashCode = (hashCode * 59) + this.AdditionalProperties.GetHashCode();
                }
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
