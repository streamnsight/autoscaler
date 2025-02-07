// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v55/common"
	"strings"
)

// WorkRequestResource The properties that define a work request resource.
type WorkRequestResource struct {

	// The way in which this resource was affected by the work tracked by the work request.
	ActionType WorkRequestResourceActionTypeEnum `mandatory:"false" json:"actionType,omitempty"`

	// The resource type the work request affects.
	EntityType *string `mandatory:"false" json:"entityType"`

	// The OCID of the resource the work request affects.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The URI path on which the user can issue a GET request to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingWorkRequestResourceActionTypeEnum[string(m.ActionType)]; !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetWorkRequestResourceActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestResourceActionTypeEnum Enum with underlying type: string
type WorkRequestResourceActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestResourceActionTypeEnum
const (
	WorkRequestResourceActionTypeCreated        WorkRequestResourceActionTypeEnum = "CREATED"
	WorkRequestResourceActionTypeUpdated        WorkRequestResourceActionTypeEnum = "UPDATED"
	WorkRequestResourceActionTypeDeleted        WorkRequestResourceActionTypeEnum = "DELETED"
	WorkRequestResourceActionTypeRelated        WorkRequestResourceActionTypeEnum = "RELATED"
	WorkRequestResourceActionTypeInProgress     WorkRequestResourceActionTypeEnum = "IN_PROGRESS"
	WorkRequestResourceActionTypeFailed         WorkRequestResourceActionTypeEnum = "FAILED"
	WorkRequestResourceActionTypeCanceledCreate WorkRequestResourceActionTypeEnum = "CANCELED_CREATE"
	WorkRequestResourceActionTypeCanceledUpdate WorkRequestResourceActionTypeEnum = "CANCELED_UPDATE"
	WorkRequestResourceActionTypeCanceledDelete WorkRequestResourceActionTypeEnum = "CANCELED_DELETE"
)

var mappingWorkRequestResourceActionTypeEnum = map[string]WorkRequestResourceActionTypeEnum{
	"CREATED":         WorkRequestResourceActionTypeCreated,
	"UPDATED":         WorkRequestResourceActionTypeUpdated,
	"DELETED":         WorkRequestResourceActionTypeDeleted,
	"RELATED":         WorkRequestResourceActionTypeRelated,
	"IN_PROGRESS":     WorkRequestResourceActionTypeInProgress,
	"FAILED":          WorkRequestResourceActionTypeFailed,
	"CANCELED_CREATE": WorkRequestResourceActionTypeCanceledCreate,
	"CANCELED_UPDATE": WorkRequestResourceActionTypeCanceledUpdate,
	"CANCELED_DELETE": WorkRequestResourceActionTypeCanceledDelete,
}

// GetWorkRequestResourceActionTypeEnumValues Enumerates the set of values for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumValues() []WorkRequestResourceActionTypeEnum {
	values := make([]WorkRequestResourceActionTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceActionTypeEnumStringValues Enumerates the set of values in String for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
		"RELATED",
		"IN_PROGRESS",
		"FAILED",
		"CANCELED_CREATE",
		"CANCELED_UPDATE",
		"CANCELED_DELETE",
	}
}
