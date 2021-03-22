// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v37/common"
	"net/http"
)

// ListComputeCapacityReservationInstancesRequest wrapper for the ListComputeCapacityReservationInstances operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListComputeCapacityReservationInstances.go.html to see an example of how to use ListComputeCapacityReservationInstancesRequest.
type ListComputeCapacityReservationInstancesRequest struct {

	// The OCID of the compute capacity reservation.
	CapacityReservationId *string `mandatory:"true" contributesTo:"path" name:"capacityReservationId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListComputeCapacityReservationInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListComputeCapacityReservationInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListComputeCapacityReservationInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListComputeCapacityReservationInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListComputeCapacityReservationInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListComputeCapacityReservationInstancesResponse wrapper for the ListComputeCapacityReservationInstances operation
type ListComputeCapacityReservationInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CapacityReservationInstanceSummary instances
	Items []CapacityReservationInstanceSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListComputeCapacityReservationInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListComputeCapacityReservationInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListComputeCapacityReservationInstancesSortByEnum Enum with underlying type: string
type ListComputeCapacityReservationInstancesSortByEnum string

// Set of constants representing the allowable values for ListComputeCapacityReservationInstancesSortByEnum
const (
	ListComputeCapacityReservationInstancesSortByTimecreated ListComputeCapacityReservationInstancesSortByEnum = "TIMECREATED"
	ListComputeCapacityReservationInstancesSortByDisplayname ListComputeCapacityReservationInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListComputeCapacityReservationInstancesSortBy = map[string]ListComputeCapacityReservationInstancesSortByEnum{
	"TIMECREATED": ListComputeCapacityReservationInstancesSortByTimecreated,
	"DISPLAYNAME": ListComputeCapacityReservationInstancesSortByDisplayname,
}

// GetListComputeCapacityReservationInstancesSortByEnumValues Enumerates the set of values for ListComputeCapacityReservationInstancesSortByEnum
func GetListComputeCapacityReservationInstancesSortByEnumValues() []ListComputeCapacityReservationInstancesSortByEnum {
	values := make([]ListComputeCapacityReservationInstancesSortByEnum, 0)
	for _, v := range mappingListComputeCapacityReservationInstancesSortBy {
		values = append(values, v)
	}
	return values
}

// ListComputeCapacityReservationInstancesSortOrderEnum Enum with underlying type: string
type ListComputeCapacityReservationInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListComputeCapacityReservationInstancesSortOrderEnum
const (
	ListComputeCapacityReservationInstancesSortOrderAsc  ListComputeCapacityReservationInstancesSortOrderEnum = "ASC"
	ListComputeCapacityReservationInstancesSortOrderDesc ListComputeCapacityReservationInstancesSortOrderEnum = "DESC"
)

var mappingListComputeCapacityReservationInstancesSortOrder = map[string]ListComputeCapacityReservationInstancesSortOrderEnum{
	"ASC":  ListComputeCapacityReservationInstancesSortOrderAsc,
	"DESC": ListComputeCapacityReservationInstancesSortOrderDesc,
}

// GetListComputeCapacityReservationInstancesSortOrderEnumValues Enumerates the set of values for ListComputeCapacityReservationInstancesSortOrderEnum
func GetListComputeCapacityReservationInstancesSortOrderEnumValues() []ListComputeCapacityReservationInstancesSortOrderEnum {
	values := make([]ListComputeCapacityReservationInstancesSortOrderEnum, 0)
	for _, v := range mappingListComputeCapacityReservationInstancesSortOrder {
		values = append(values, v)
	}
	return values
}
