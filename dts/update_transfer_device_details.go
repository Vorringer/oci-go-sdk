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

// UpdateTransferDeviceDetails The representation of UpdateTransferDeviceDetails
type UpdateTransferDeviceDetails struct {
	LifecycleState UpdateTransferDeviceDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m UpdateTransferDeviceDetails) String() string {
	return common.PointerString(m)
}

// UpdateTransferDeviceDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferDeviceDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferDeviceDetailsLifecycleStateEnum
const (
	UpdateTransferDeviceDetailsLifecycleStatePreparing UpdateTransferDeviceDetailsLifecycleStateEnum = "PREPARING"
	UpdateTransferDeviceDetailsLifecycleStateReady     UpdateTransferDeviceDetailsLifecycleStateEnum = "READY"
	UpdateTransferDeviceDetailsLifecycleStateCancelled UpdateTransferDeviceDetailsLifecycleStateEnum = "CANCELLED"
)

var mappingUpdateTransferDeviceDetailsLifecycleState = map[string]UpdateTransferDeviceDetailsLifecycleStateEnum{
	"PREPARING": UpdateTransferDeviceDetailsLifecycleStatePreparing,
	"READY":     UpdateTransferDeviceDetailsLifecycleStateReady,
	"CANCELLED": UpdateTransferDeviceDetailsLifecycleStateCancelled,
}

// GetUpdateTransferDeviceDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferDeviceDetailsLifecycleStateEnum
func GetUpdateTransferDeviceDetailsLifecycleStateEnumValues() []UpdateTransferDeviceDetailsLifecycleStateEnum {
	values := make([]UpdateTransferDeviceDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferDeviceDetailsLifecycleState {
		values = append(values, v)
	}
	return values
}
