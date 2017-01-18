/* 
 * Trakerr API
 *
 * Get your application events and errors to Trakerr via the *Trakerr API*.
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package trakerr_client

import (
	"net/url"
	"fmt"
)

type EventsApi struct {
	Configuration Configuration
}

func NewEventsApi() *EventsApi {
	configuration := NewConfiguration()
	return &EventsApi{
		Configuration: *configuration,
	}
}

func NewEventsApiWithBasePath(basePath string) *EventsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &EventsApi{
		Configuration: *configuration,
	}
}

/**
 * Submit an application event or error to Trakerr
 *  The events endpoint submits an application event or an application error / exception with an optional stacktrace field to Trakerr.  ##### Sample POST request body: &#x60;&#x60;&#x60; {  \&quot;apiKey\&quot;: \&quot;a9a2807a2e8fd4602adae9e8f819790a267213234083\&quot;,  \&quot;classification\&quot;: \&quot;Error\&quot;,  \&quot;eventType\&quot;: \&quot;System.Exception\&quot;,  \&quot;eventMessage\&quot;: \&quot;This is a test exception.\&quot;,  \&quot;eventTime\&quot;: 1479477482291,  \&quot;eventStacktrace\&quot;: [    {      \&quot;type\&quot;: \&quot;System.Exception\&quot;,      \&quot;message\&quot;: \&quot;This is a test exception.\&quot;,      \&quot;traceLines\&quot;: [        {          \&quot;function\&quot;: \&quot;Main\&quot;,          \&quot;line\&quot;: 19,          \&quot;file\&quot;: \&quot;TrakerrSampleApp\\\\Program.cs\&quot;        }      ]    }  ],  \&quot;contextAppVersion\&quot;: \&quot;1.0\&quot;,  \&quot;contextEnvName\&quot;: \&quot;development\&quot;,  \&quot;contextEnvHostname\&quot;: \&quot;trakerr.io\&quot;,  \&quot;contextAppOS\&quot;: \&quot;Win32NT Service Pack 1\&quot;,  \&quot;contextAppOSVersion\&quot;: \&quot;6.1.7601.65536\&quot; } &#x60;&#x60;&#x60; ##### Sample POST response body (200 OK): &#x60;&#x60;&#x60; { } &#x60;&#x60;&#x60;
 *
 * @param data Event to submit
 * @return void
 */
func (a EventsApi) EventsPost(data AppEvent) (*APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/events"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &data


	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		fmt.Println(err)
		return NewAPIResponse(httpResponse.RawResponse), err
	}
	fmt.Println(httpResponse.RawResponse)

	return NewAPIResponse(httpResponse.RawResponse), err
}
