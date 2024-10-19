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
    /// Pet
    /// </summary>
    [DataContract(Name = "Pet")]
    public partial class Pet : IEquatable<Pet>
    {
        /// <summary>
        /// pet status in the store
        /// </summary>
        /// <value>pet status in the store</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum StatusEnum
        {
            /// <summary>
            /// Enum Available for value: available
            /// </summary>
            [EnumMember(Value = "available")]
            Available = 1,

            /// <summary>
            /// Enum Pending for value: pending
            /// </summary>
            [EnumMember(Value = "pending")]
            Pending = 2,

            /// <summary>
            /// Enum Sold for value: sold
            /// </summary>
            [EnumMember(Value = "sold")]
            Sold = 3
        }

        /// <summary>
        /// pet status in the store
        /// </summary>
        /// <value>pet status in the store</value>
        [DataMember(Name = "status", EmitDefaultValue = false)]
        public StatusEnum? Status { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="Pet" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected Pet() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="Pet" /> class.
        /// </summary>
        /// <param name="id">id.</param>
        /// <param name="category">category.</param>
        /// <param name="name">name (required).</param>
        /// <param name="photoUrls">photoUrls (required).</param>
        /// <param name="tags">tags.</param>
        /// <param name="status">pet status in the store.</param>
        public Pet(long id = default(long), Category category = default(Category), string name = default(string), List<string> photoUrls = default(List<string>), List<Tag> tags = default(List<Tag>), StatusEnum? status = default(StatusEnum?))
        {
            // to ensure "name" is required (not null)
            if (name == null)
            {
                throw new ArgumentNullException("name is a required property for Pet and cannot be null");
            }
            this.Name = name;
            // to ensure "photoUrls" is required (not null)
            if (photoUrls == null)
            {
                throw new ArgumentNullException("photoUrls is a required property for Pet and cannot be null");
            }
            this.PhotoUrls = photoUrls;
            this.Id = id;
            this.Category = category;
            this.Tags = tags;
            this.Status = status;
        }

        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name = "id", EmitDefaultValue = false)]
        public long Id { get; set; }

        /// <summary>
        /// Gets or Sets Category
        /// </summary>
        [DataMember(Name = "category", EmitDefaultValue = false)]
        public Category Category { get; set; }

        /// <summary>
        /// Gets or Sets Name
        /// </summary>
        /*
        <example>doggie</example>
        */
        [DataMember(Name = "name", IsRequired = true, EmitDefaultValue = true)]
        public string Name { get; set; }

        /// <summary>
        /// Gets or Sets PhotoUrls
        /// </summary>
        [DataMember(Name = "photoUrls", IsRequired = true, EmitDefaultValue = true)]
        public List<string> PhotoUrls { get; set; }

        /// <summary>
        /// Gets or Sets Tags
        /// </summary>
        [DataMember(Name = "tags", EmitDefaultValue = false)]
        public List<Tag> Tags { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class Pet {\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Category: ").Append(Category).Append("\n");
            sb.Append("  Name: ").Append(Name).Append("\n");
            sb.Append("  PhotoUrls: ").Append(PhotoUrls).Append("\n");
            sb.Append("  Tags: ").Append(Tags).Append("\n");
            sb.Append("  Status: ").Append(Status).Append("\n");
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
            return this.Equals(input as Pet);
        }

        /// <summary>
        /// Returns true if Pet instances are equal
        /// </summary>
        /// <param name="input">Instance of Pet to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(Pet input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Id == input.Id ||
                    this.Id.Equals(input.Id)
                ) && 
                (
                    this.Category == input.Category ||
                    (this.Category != null &&
                    this.Category.Equals(input.Category))
                ) && 
                (
                    this.Name == input.Name ||
                    (this.Name != null &&
                    this.Name.Equals(input.Name))
                ) && 
                (
                    this.PhotoUrls == input.PhotoUrls ||
                    this.PhotoUrls != null &&
                    input.PhotoUrls != null &&
                    this.PhotoUrls.SequenceEqual(input.PhotoUrls)
                ) && 
                (
                    this.Tags == input.Tags ||
                    this.Tags != null &&
                    input.Tags != null &&
                    this.Tags.SequenceEqual(input.Tags)
                ) && 
                (
                    this.Status == input.Status ||
                    this.Status.Equals(input.Status)
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
                hashCode = (hashCode * 59) + this.Id.GetHashCode();
                if (this.Category != null)
                {
                    hashCode = (hashCode * 59) + this.Category.GetHashCode();
                }
                if (this.Name != null)
                {
                    hashCode = (hashCode * 59) + this.Name.GetHashCode();
                }
                if (this.PhotoUrls != null)
                {
                    hashCode = (hashCode * 59) + this.PhotoUrls.GetHashCode();
                }
                if (this.Tags != null)
                {
                    hashCode = (hashCode * 59) + this.Tags.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Status.GetHashCode();
                return hashCode;
            }
        }

    }

}
