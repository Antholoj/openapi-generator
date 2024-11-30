/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 * Generated by: https://openapi-generator.tech
 */

using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Http;
using Swashbuckle.AspNetCore.Annotations;
using Swashbuckle.AspNetCore.SwaggerGen;
using Newtonsoft.Json;
using Org.OpenAPITools.Attributes;
using Org.OpenAPITools.Models;

namespace Org.OpenAPITools.Controllers
{ 
    /// <summary>
    /// 
    /// </summary>
    [ApiController]
    public class FakeApiController : ControllerBase
    { 
        /// <summary>
        /// Fake endpoint to test nullable example (object)
        /// </summary>
        /// <response code="200">Successful operation</response>
        [HttpGet]
        [Route("/v2/fake/nullable_example_test")]
        [ValidateModelState]
        [SwaggerOperation("FakeNullableExampleTest")]
        [SwaggerResponse(statusCode: 200, type: typeof(TestNullable), description: "Successful operation")]
        public virtual IActionResult FakeNullableExampleTest()
        {

            //TODO: Uncomment the next line to return response 200 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(200, default(TestNullable));
            string exampleJson = null;
            exampleJson = "{\r\n  \"nullableName\" : \"nullableName\",\r\n  \"name\" : \"name\"\r\n}";
            
            var example = exampleJson != null
            ? JsonConvert.DeserializeObject<TestNullable>(exampleJson)
            : default(TestNullable);
            //TODO: Change the data returned
            return new ObjectResult(example);
        }

        /// <summary>
        /// fake endpoint to test parameter example (object)
        /// </summary>
        /// <param name="data"></param>
        /// <response code="0">successful operation</response>
        [HttpGet]
        [Route("/v2/fake/parameter_example_test")]
        [ValidateModelState]
        [SwaggerOperation("FakeParameterExampleTest")]
        public virtual IActionResult FakeParameterExampleTest([FromQuery (Name = "data")][Required()]Pet data)
        {

            //TODO: Uncomment the next line to return response 0 or use other options such as return this.NotFound(), return this.BadRequest(..), ...
            // return StatusCode(0);

            throw new NotImplementedException();
        }
    }
}
