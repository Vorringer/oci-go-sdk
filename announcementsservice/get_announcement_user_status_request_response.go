// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package announcementsservice

import (
	"github.com/oracle/oci-go-sdk/v46/common"
	"net/http"
)

// GetAnnouncementUserStatusRequest wrapper for the GetAnnouncementUserStatus operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/GetAnnouncementUserStatus.go.html to see an example of how to use GetAnnouncementUserStatusRequest.
type GetAnnouncementUserStatusRequest struct {

	// The OCID of the announcement.
	AnnouncementId *string `mandatory:"true" contributesTo:"path" name:"announcementId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAnnouncementUserStatusRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAnnouncementUserStatusRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAnnouncementUserStatusRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAnnouncementUserStatusRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAnnouncementUserStatusResponse wrapper for the GetAnnouncementUserStatus operation
type GetAnnouncementUserStatusResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AnnouncementUserStatusDetails instance
	AnnouncementUserStatusDetails `presentIn:"body"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the complete request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAnnouncementUserStatusResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAnnouncementUserStatusResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
