// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v46/common"
)

// OccSourceDetails Details about the Oracle Cloud@Customer account, the source environment from which you want to migrate the application.
type OccSourceDetails struct {

	// If you are using an Oracle Cloud@Customer account with Identity Cloud Service (IDCS), enter the service instance ID.
	// For example, if Compute-567890123 is the account name of your Oracle Cloud@Customer Compute service entitlement,
	// then enter 567890123.
	ComputeAccount *string `mandatory:"true" json:"computeAccount"`
}

func (m OccSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OccSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOccSourceDetails OccSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOccSourceDetails
	}{
		"OCC",
		(MarshalTypeOccSourceDetails)(m),
	}

	return json.Marshal(&s)
}
