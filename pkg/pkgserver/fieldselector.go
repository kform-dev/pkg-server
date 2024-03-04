/*
Copyright 2024 Nokia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkgserver

/*

// packageFilter filters
type packageFilter struct {
	// Name filters by the name of the objects
	Name string

	// Namespace filters by the namespace of the objects
	Namespace string

	// Repository filters to repositories with the given name.
	Repository string

	// Package filters to packages with the given name.
	Package string

	// Revision filters to revisions with the given name.
	Revision string

	// Workspace filters the workspaces with the given name.
	Workspace string
}

// parsePackageFieldSelector parses client-provided fields.Selector into a packageFilter
func parsePackageFieldSelector(fieldSelector fields.Selector) (*packageFilter, error) {
	var filter *packageFilter

	if fieldSelector == nil {
		return filter, nil
	}

	requirements := fieldSelector.Requirements()
	for _, requirement := range requirements {
		filter = &packageFilter{}
		switch requirement.Operator {
		case selection.Equals, selection.DoesNotExist:
			if requirement.Value == "" {
				return filter, apierrors.NewBadRequest(fmt.Sprintf("unsupported fieldSelector value %q for field %q with operator %q", requirement.Value, requirement.Field, requirement.Operator))
			}
		default:
			return filter, apierrors.NewBadRequest(fmt.Sprintf("unsupported fieldSelector operator %q for field %q", requirement.Operator, requirement.Field))
		}
		fmt.Println("requirement.Field", requirement.Field)
		fmt.Println("requirement.Value", requirement.Value)

		switch requirement.Field {
		case "metadata.name":
			filter.Name = requirement.Value
		case "metadata.namespace":
			filter.Namespace = requirement.Value
		case "spec.repository":
			filter.Repository = requirement.Value
		case "spec.package":
			filter.Package = requirement.Value
		case "spec.revision":
			filter.Revision = requirement.Value
		case "spec.workspace":
			filter.Workspace = requirement.Value
		default:
			return filter, apierrors.NewBadRequest(fmt.Sprintf("unknown fieldSelector field %q", requirement.Field))
		}
	}

	return filter, nil
}
*/
