// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// ShippingVendors The representation of ShippingVendors
type ShippingVendors struct {

	// List of available shipping vendors for package delivery
	Vendors []string `mandatory:"false" json:"vendors"`
}

func (m ShippingVendors) String() string {
	return common.PointerString(m)
}
