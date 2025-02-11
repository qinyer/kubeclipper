/*
 *
 *  * Copyright 2021 KubeClipper Authors.
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

// Package validation include all tenant resource validation func.
package validation

import (
	tenantv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/tenant/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateUserCreate validate project when create.
func ValidateUserCreate(p *tenantv1.Project) field.ErrorList {
	return validateProjectSpecSpec(&p.Spec, field.NewPath("spec"))
}

// ValidateProjectUpdate validate project when update.
func ValidateProjectUpdate(p *tenantv1.Project) field.ErrorList {
	return validateProjectSpecSpec(&p.Spec, field.NewPath("spec"))
}

func validateProjectSpecSpec(spec *tenantv1.ProjectSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if spec.Manager == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("manager"), "must be set"))
	}
	for _, nodeID := range spec.Nodes {
		if nodeID == "" {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("nodes"), nodeID, "nodeID can't be empty string"))
		}
	}
	// TODO: validate other field
	return allErrs
}
