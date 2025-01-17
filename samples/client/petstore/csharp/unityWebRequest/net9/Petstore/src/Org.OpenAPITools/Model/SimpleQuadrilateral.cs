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
using OpenAPIDateConverter = Org.OpenAPITools.Client.OpenAPIDateConverter;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// SimpleQuadrilateral
    /// </summary>
    [DataContract(Name = "SimpleQuadrilateral")]
    public partial class SimpleQuadrilateral : IEquatable<SimpleQuadrilateral>
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="SimpleQuadrilateral" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected SimpleQuadrilateral() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="SimpleQuadrilateral" /> class.
        /// </summary>
        /// <param name="shapeType">shapeType (required).</param>
        /// <param name="quadrilateralType">quadrilateralType (required).</param>
        public SimpleQuadrilateral(string shapeType = default(string), string quadrilateralType = default(string))
        {
            // to ensure "shapeType" is required (not null)
            if (shapeType == null)
            {
                throw new ArgumentNullException("shapeType is a required property for SimpleQuadrilateral and cannot be null");
            }
            this.ShapeType = shapeType;
            // to ensure "quadrilateralType" is required (not null)
            if (quadrilateralType == null)
            {
                throw new ArgumentNullException("quadrilateralType is a required property for SimpleQuadrilateral and cannot be null");
            }
            this.QuadrilateralType = quadrilateralType;
        }

        /// <summary>
        /// Gets or Sets ShapeType
        /// </summary>
        [DataMember(Name = "shapeType", IsRequired = true, EmitDefaultValue = true)]
        public string ShapeType { get; set; }

        /// <summary>
        /// Gets or Sets QuadrilateralType
        /// </summary>
        [DataMember(Name = "quadrilateralType", IsRequired = true, EmitDefaultValue = true)]
        public string QuadrilateralType { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class SimpleQuadrilateral {\n");
            sb.Append("  ShapeType: ").Append(ShapeType).Append("\n");
            sb.Append("  QuadrilateralType: ").Append(QuadrilateralType).Append("\n");
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
            return this.Equals(input as SimpleQuadrilateral);
        }

        /// <summary>
        /// Returns true if SimpleQuadrilateral instances are equal
        /// </summary>
        /// <param name="input">Instance of SimpleQuadrilateral to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(SimpleQuadrilateral input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.ShapeType == input.ShapeType ||
                    (this.ShapeType != null &&
                    this.ShapeType.Equals(input.ShapeType))
                ) && 
                (
                    this.QuadrilateralType == input.QuadrilateralType ||
                    (this.QuadrilateralType != null &&
                    this.QuadrilateralType.Equals(input.QuadrilateralType))
                );
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
                if (this.ShapeType != null)
                {
                    hashCode = (hashCode * 59) + this.ShapeType.GetHashCode();
                }
                if (this.QuadrilateralType != null)
                {
                    hashCode = (hashCode * 59) + this.QuadrilateralType.GetHashCode();
                }
                return hashCode;
            }
        }

    }

}
