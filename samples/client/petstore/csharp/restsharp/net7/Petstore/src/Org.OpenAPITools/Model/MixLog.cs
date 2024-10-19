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
    /// MixLog
    /// </summary>
    [DataContract(Name = "MixLog")]
    public partial class MixLog : IEquatable<MixLog>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="MixLog" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected MixLog() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="MixLog" /> class.
        /// </summary>
        /// <param name="id">id (required).</param>
        /// <param name="description">description (required).</param>
        /// <param name="mixDate">mixDate (required).</param>
        /// <param name="shopId">shopId.</param>
        /// <param name="totalPrice">totalPrice.</param>
        /// <param name="totalRecalculations">totalRecalculations (required).</param>
        /// <param name="totalOverPoors">totalOverPoors (required).</param>
        /// <param name="totalSkips">totalSkips (required).</param>
        /// <param name="totalUnderPours">totalUnderPours (required).</param>
        /// <param name="formulaVersionDate">formulaVersionDate (required).</param>
        /// <param name="someCode">SomeCode is only required for color mixes.</param>
        /// <param name="batchNumber">batchNumber.</param>
        /// <param name="brandCode">BrandCode is only required for non-color mixes.</param>
        /// <param name="brandId">BrandId is only required for color mixes.</param>
        /// <param name="brandName">BrandName is only required for color mixes.</param>
        /// <param name="categoryCode">CategoryCode is not used anymore.</param>
        /// <param name="color">Color is only required for color mixes.</param>
        /// <param name="colorDescription">colorDescription.</param>
        /// <param name="comment">comment.</param>
        /// <param name="commercialProductCode">commercialProductCode.</param>
        /// <param name="productLineCode">ProductLineCode is only required for color mixes.</param>
        /// <param name="country">country.</param>
        /// <param name="createdBy">createdBy.</param>
        /// <param name="createdByFirstName">createdByFirstName.</param>
        /// <param name="createdByLastName">createdByLastName.</param>
        /// <param name="deltaECalculationRepaired">deltaECalculationRepaired.</param>
        /// <param name="deltaECalculationSprayout">deltaECalculationSprayout.</param>
        /// <param name="ownColorVariantNumber">ownColorVariantNumber.</param>
        /// <param name="primerProductId">primerProductId.</param>
        /// <param name="productId">ProductId is only required for color mixes.</param>
        /// <param name="productName">ProductName is only required for color mixes.</param>
        /// <param name="selectedVersionIndex">selectedVersionIndex.</param>
        public MixLog(Guid id = default(Guid), string description = default(string), DateTime mixDate = default(DateTime), Guid shopId = default(Guid), float? totalPrice = default(float?), int totalRecalculations = default(int), int totalOverPoors = default(int), int totalSkips = default(int), int totalUnderPours = default(int), DateTime formulaVersionDate = default(DateTime), string someCode = default(string), string batchNumber = default(string), string brandCode = default(string), string brandId = default(string), string brandName = default(string), string categoryCode = default(string), string color = default(string), string colorDescription = default(string), string comment = default(string), string commercialProductCode = default(string), string productLineCode = default(string), string country = default(string), string createdBy = default(string), string createdByFirstName = default(string), string createdByLastName = default(string), string deltaECalculationRepaired = default(string), string deltaECalculationSprayout = default(string), int? ownColorVariantNumber = default(int?), string primerProductId = default(string), string productId = default(string), string productName = default(string), int selectedVersionIndex = default(int))
        {
            this.Id = id;
            // to ensure "description" is required (not null)
            if (description == null)
            {
                throw new ArgumentNullException("description is a required property for MixLog and cannot be null");
            }
            this.Description = description;
            this.MixDate = mixDate;
            this.TotalRecalculations = totalRecalculations;
            this.TotalOverPoors = totalOverPoors;
            this.TotalSkips = totalSkips;
            this.TotalUnderPours = totalUnderPours;
            this.FormulaVersionDate = formulaVersionDate;
            this.ShopId = shopId;
            this.TotalPrice = totalPrice;
            this.SomeCode = someCode;
            this.BatchNumber = batchNumber;
            this.BrandCode = brandCode;
            this.BrandId = brandId;
            this.BrandName = brandName;
            this.CategoryCode = categoryCode;
            this.Color = color;
            this.ColorDescription = colorDescription;
            this.Comment = comment;
            this.CommercialProductCode = commercialProductCode;
            this.ProductLineCode = productLineCode;
            this.Country = country;
            this.CreatedBy = createdBy;
            this.CreatedByFirstName = createdByFirstName;
            this.CreatedByLastName = createdByLastName;
            this.DeltaECalculationRepaired = deltaECalculationRepaired;
            this.DeltaECalculationSprayout = deltaECalculationSprayout;
            this.OwnColorVariantNumber = ownColorVariantNumber;
            this.PrimerProductId = primerProductId;
            this.ProductId = productId;
            this.ProductName = productName;
            this.SelectedVersionIndex = selectedVersionIndex;
        }

        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = true)]
        public Guid Id { get; set; }

        /// <summary>
        /// Gets or Sets Description
        /// </summary>
        [DataMember(Name = "description", IsRequired = true, EmitDefaultValue = true)]
        public string Description { get; set; }

        /// <summary>
        /// Gets or Sets MixDate
        /// </summary>
        [DataMember(Name = "mixDate", IsRequired = true, EmitDefaultValue = true)]
        public DateTime MixDate { get; set; }

        /// <summary>
        /// Gets or Sets ShopId
        /// </summary>
        [DataMember(Name = "shopId", EmitDefaultValue = false)]
        public Guid ShopId { get; set; }

        /// <summary>
        /// Gets or Sets TotalPrice
        /// </summary>
        [DataMember(Name = "totalPrice", EmitDefaultValue = true)]
        public float? TotalPrice { get; set; }

        /// <summary>
        /// Gets or Sets TotalRecalculations
        /// </summary>
        [DataMember(Name = "totalRecalculations", IsRequired = true, EmitDefaultValue = true)]
        public int TotalRecalculations { get; set; }

        /// <summary>
        /// Gets or Sets TotalOverPoors
        /// </summary>
        [DataMember(Name = "totalOverPoors", IsRequired = true, EmitDefaultValue = true)]
        public int TotalOverPoors { get; set; }

        /// <summary>
        /// Gets or Sets TotalSkips
        /// </summary>
        [DataMember(Name = "totalSkips", IsRequired = true, EmitDefaultValue = true)]
        public int TotalSkips { get; set; }

        /// <summary>
        /// Gets or Sets TotalUnderPours
        /// </summary>
        [DataMember(Name = "totalUnderPours", IsRequired = true, EmitDefaultValue = true)]
        public int TotalUnderPours { get; set; }

        /// <summary>
        /// Gets or Sets FormulaVersionDate
        /// </summary>
        [DataMember(Name = "formulaVersionDate", IsRequired = true, EmitDefaultValue = true)]
        public DateTime FormulaVersionDate { get; set; }

        /// <summary>
        /// SomeCode is only required for color mixes
        /// </summary>
        /// <value>SomeCode is only required for color mixes</value>
        [DataMember(Name = "someCode", EmitDefaultValue = true)]
        public string SomeCode { get; set; }

        /// <summary>
        /// Gets or Sets BatchNumber
        /// </summary>
        [DataMember(Name = "batchNumber", EmitDefaultValue = false)]
        public string BatchNumber { get; set; }

        /// <summary>
        /// BrandCode is only required for non-color mixes
        /// </summary>
        /// <value>BrandCode is only required for non-color mixes</value>
        [DataMember(Name = "brandCode", EmitDefaultValue = false)]
        public string BrandCode { get; set; }

        /// <summary>
        /// BrandId is only required for color mixes
        /// </summary>
        /// <value>BrandId is only required for color mixes</value>
        [DataMember(Name = "brandId", EmitDefaultValue = false)]
        public string BrandId { get; set; }

        /// <summary>
        /// BrandName is only required for color mixes
        /// </summary>
        /// <value>BrandName is only required for color mixes</value>
        [DataMember(Name = "brandName", EmitDefaultValue = false)]
        public string BrandName { get; set; }

        /// <summary>
        /// CategoryCode is not used anymore
        /// </summary>
        /// <value>CategoryCode is not used anymore</value>
        [DataMember(Name = "categoryCode", EmitDefaultValue = false)]
        public string CategoryCode { get; set; }

        /// <summary>
        /// Color is only required for color mixes
        /// </summary>
        /// <value>Color is only required for color mixes</value>
        [DataMember(Name = "color", EmitDefaultValue = false)]
        public string Color { get; set; }

        /// <summary>
        /// Gets or Sets ColorDescription
        /// </summary>
        [DataMember(Name = "colorDescription", EmitDefaultValue = false)]
        public string ColorDescription { get; set; }

        /// <summary>
        /// Gets or Sets Comment
        /// </summary>
        [DataMember(Name = "comment", EmitDefaultValue = false)]
        public string Comment { get; set; }

        /// <summary>
        /// Gets or Sets CommercialProductCode
        /// </summary>
        [DataMember(Name = "commercialProductCode", EmitDefaultValue = false)]
        public string CommercialProductCode { get; set; }

        /// <summary>
        /// ProductLineCode is only required for color mixes
        /// </summary>
        /// <value>ProductLineCode is only required for color mixes</value>
        [DataMember(Name = "productLineCode", EmitDefaultValue = false)]
        public string ProductLineCode { get; set; }

        /// <summary>
        /// Gets or Sets Country
        /// </summary>
        [DataMember(Name = "country", EmitDefaultValue = false)]
        public string Country { get; set; }

        /// <summary>
        /// Gets or Sets CreatedBy
        /// </summary>
        [DataMember(Name = "createdBy", EmitDefaultValue = false)]
        public string CreatedBy { get; set; }

        /// <summary>
        /// Gets or Sets CreatedByFirstName
        /// </summary>
        [DataMember(Name = "createdByFirstName", EmitDefaultValue = false)]
        public string CreatedByFirstName { get; set; }

        /// <summary>
        /// Gets or Sets CreatedByLastName
        /// </summary>
        [DataMember(Name = "createdByLastName", EmitDefaultValue = false)]
        public string CreatedByLastName { get; set; }

        /// <summary>
        /// Gets or Sets DeltaECalculationRepaired
        /// </summary>
        [DataMember(Name = "deltaECalculationRepaired", EmitDefaultValue = false)]
        public string DeltaECalculationRepaired { get; set; }

        /// <summary>
        /// Gets or Sets DeltaECalculationSprayout
        /// </summary>
        [DataMember(Name = "deltaECalculationSprayout", EmitDefaultValue = false)]
        public string DeltaECalculationSprayout { get; set; }

        /// <summary>
        /// Gets or Sets OwnColorVariantNumber
        /// </summary>
        [DataMember(Name = "ownColorVariantNumber", EmitDefaultValue = true)]
        public int? OwnColorVariantNumber { get; set; }

        /// <summary>
        /// Gets or Sets PrimerProductId
        /// </summary>
        [DataMember(Name = "primerProductId", EmitDefaultValue = false)]
        public string PrimerProductId { get; set; }

        /// <summary>
        /// ProductId is only required for color mixes
        /// </summary>
        /// <value>ProductId is only required for color mixes</value>
        [DataMember(Name = "productId", EmitDefaultValue = false)]
        public string ProductId { get; set; }

        /// <summary>
        /// ProductName is only required for color mixes
        /// </summary>
        /// <value>ProductName is only required for color mixes</value>
        [DataMember(Name = "productName", EmitDefaultValue = false)]
        public string ProductName { get; set; }

        /// <summary>
        /// Gets or Sets SelectedVersionIndex
        /// </summary>
        [DataMember(Name = "selectedVersionIndex", EmitDefaultValue = false)]
        public int SelectedVersionIndex { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class MixLog {\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Description: ").Append(Description).Append("\n");
            sb.Append("  MixDate: ").Append(MixDate).Append("\n");
            sb.Append("  ShopId: ").Append(ShopId).Append("\n");
            sb.Append("  TotalPrice: ").Append(TotalPrice).Append("\n");
            sb.Append("  TotalRecalculations: ").Append(TotalRecalculations).Append("\n");
            sb.Append("  TotalOverPoors: ").Append(TotalOverPoors).Append("\n");
            sb.Append("  TotalSkips: ").Append(TotalSkips).Append("\n");
            sb.Append("  TotalUnderPours: ").Append(TotalUnderPours).Append("\n");
            sb.Append("  FormulaVersionDate: ").Append(FormulaVersionDate).Append("\n");
            sb.Append("  SomeCode: ").Append(SomeCode).Append("\n");
            sb.Append("  BatchNumber: ").Append(BatchNumber).Append("\n");
            sb.Append("  BrandCode: ").Append(BrandCode).Append("\n");
            sb.Append("  BrandId: ").Append(BrandId).Append("\n");
            sb.Append("  BrandName: ").Append(BrandName).Append("\n");
            sb.Append("  CategoryCode: ").Append(CategoryCode).Append("\n");
            sb.Append("  Color: ").Append(Color).Append("\n");
            sb.Append("  ColorDescription: ").Append(ColorDescription).Append("\n");
            sb.Append("  Comment: ").Append(Comment).Append("\n");
            sb.Append("  CommercialProductCode: ").Append(CommercialProductCode).Append("\n");
            sb.Append("  ProductLineCode: ").Append(ProductLineCode).Append("\n");
            sb.Append("  Country: ").Append(Country).Append("\n");
            sb.Append("  CreatedBy: ").Append(CreatedBy).Append("\n");
            sb.Append("  CreatedByFirstName: ").Append(CreatedByFirstName).Append("\n");
            sb.Append("  CreatedByLastName: ").Append(CreatedByLastName).Append("\n");
            sb.Append("  DeltaECalculationRepaired: ").Append(DeltaECalculationRepaired).Append("\n");
            sb.Append("  DeltaECalculationSprayout: ").Append(DeltaECalculationSprayout).Append("\n");
            sb.Append("  OwnColorVariantNumber: ").Append(OwnColorVariantNumber).Append("\n");
            sb.Append("  PrimerProductId: ").Append(PrimerProductId).Append("\n");
            sb.Append("  ProductId: ").Append(ProductId).Append("\n");
            sb.Append("  ProductName: ").Append(ProductName).Append("\n");
            sb.Append("  SelectedVersionIndex: ").Append(SelectedVersionIndex).Append("\n");
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
            return OpenAPIClientUtils.compareLogic.Compare(this, input as MixLog).AreEqual;
        }

        /// <summary>
        /// Returns true if MixLog instances are equal
        /// </summary>
        /// <param name="input">Instance of MixLog to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(MixLog input)
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
                if (this.Id != null)
                {
                    hashCode = (hashCode * 59) + this.Id.GetHashCode();
                }
                if (this.Description != null)
                {
                    hashCode = (hashCode * 59) + this.Description.GetHashCode();
                }
                if (this.MixDate != null)
                {
                    hashCode = (hashCode * 59) + this.MixDate.GetHashCode();
                }
                if (this.ShopId != null)
                {
                    hashCode = (hashCode * 59) + this.ShopId.GetHashCode();
                }
                if (this.TotalPrice != null)
                {
                    hashCode = (hashCode * 59) + this.TotalPrice.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.TotalRecalculations.GetHashCode();
                hashCode = (hashCode * 59) + this.TotalOverPoors.GetHashCode();
                hashCode = (hashCode * 59) + this.TotalSkips.GetHashCode();
                hashCode = (hashCode * 59) + this.TotalUnderPours.GetHashCode();
                if (this.FormulaVersionDate != null)
                {
                    hashCode = (hashCode * 59) + this.FormulaVersionDate.GetHashCode();
                }
                if (this.SomeCode != null)
                {
                    hashCode = (hashCode * 59) + this.SomeCode.GetHashCode();
                }
                if (this.BatchNumber != null)
                {
                    hashCode = (hashCode * 59) + this.BatchNumber.GetHashCode();
                }
                if (this.BrandCode != null)
                {
                    hashCode = (hashCode * 59) + this.BrandCode.GetHashCode();
                }
                if (this.BrandId != null)
                {
                    hashCode = (hashCode * 59) + this.BrandId.GetHashCode();
                }
                if (this.BrandName != null)
                {
                    hashCode = (hashCode * 59) + this.BrandName.GetHashCode();
                }
                if (this.CategoryCode != null)
                {
                    hashCode = (hashCode * 59) + this.CategoryCode.GetHashCode();
                }
                if (this.Color != null)
                {
                    hashCode = (hashCode * 59) + this.Color.GetHashCode();
                }
                if (this.ColorDescription != null)
                {
                    hashCode = (hashCode * 59) + this.ColorDescription.GetHashCode();
                }
                if (this.Comment != null)
                {
                    hashCode = (hashCode * 59) + this.Comment.GetHashCode();
                }
                if (this.CommercialProductCode != null)
                {
                    hashCode = (hashCode * 59) + this.CommercialProductCode.GetHashCode();
                }
                if (this.ProductLineCode != null)
                {
                    hashCode = (hashCode * 59) + this.ProductLineCode.GetHashCode();
                }
                if (this.Country != null)
                {
                    hashCode = (hashCode * 59) + this.Country.GetHashCode();
                }
                if (this.CreatedBy != null)
                {
                    hashCode = (hashCode * 59) + this.CreatedBy.GetHashCode();
                }
                if (this.CreatedByFirstName != null)
                {
                    hashCode = (hashCode * 59) + this.CreatedByFirstName.GetHashCode();
                }
                if (this.CreatedByLastName != null)
                {
                    hashCode = (hashCode * 59) + this.CreatedByLastName.GetHashCode();
                }
                if (this.DeltaECalculationRepaired != null)
                {
                    hashCode = (hashCode * 59) + this.DeltaECalculationRepaired.GetHashCode();
                }
                if (this.DeltaECalculationSprayout != null)
                {
                    hashCode = (hashCode * 59) + this.DeltaECalculationSprayout.GetHashCode();
                }
                if (this.OwnColorVariantNumber != null)
                {
                    hashCode = (hashCode * 59) + this.OwnColorVariantNumber.GetHashCode();
                }
                if (this.PrimerProductId != null)
                {
                    hashCode = (hashCode * 59) + this.PrimerProductId.GetHashCode();
                }
                if (this.ProductId != null)
                {
                    hashCode = (hashCode * 59) + this.ProductId.GetHashCode();
                }
                if (this.ProductName != null)
                {
                    hashCode = (hashCode * 59) + this.ProductName.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.SelectedVersionIndex.GetHashCode();
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
