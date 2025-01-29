/*
 * Seaweedfs Master Server API
 *
 * The Seaweedfs Master Server API allows you to store blobs
 *
 * The version of the OpenAPI document: 3.43.0
 * 
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using RestSharp;

namespace Org.OpenAPITools.Client
{
    /// <summary>
    /// A delegate to ExceptionFactory method
    /// </summary>
    /// <param name="methodName">Method name</param>
    /// <param name="response">Response</param>
    /// <returns>Exceptions</returns>
    public delegate Exception ExceptionFactory(string methodName, IRestResponse response);
}
