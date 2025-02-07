// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Search Service API
//
// Search for resources in your cloud network.
//

package resourcesearch

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// SearchContext Contains search context, such as highlighting, for found resources.
type SearchContext struct {

	// Describes what in each field matched the search criteria by showing highlighted values, but only for free text searches or for structured
	// queries that use a MATCHING clause. The list of strings represents fragments of values that matched the query conditions. Highlighted
	// values are wrapped with &lt;h1&gt;..&lt;/h1&gt; tags. All values are HTML-encoded (except &lt;h1&gt; tags).
	Highlights map[string][]string `mandatory:"false" json:"highlights"`
}

func (m SearchContext) String() string {
	return common.PointerString(m)
}
