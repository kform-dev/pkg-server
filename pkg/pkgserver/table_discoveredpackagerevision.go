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
import (
	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewDiscoveredPackageRevisionTableConvertor(gr schema.GroupResource) tableConvertor {
	return tableConvertor{
		resource: gr,
		cells: func(obj runtime.Object) []interface{} {
			discoveredPkgRev, ok := obj.(*pkgv1alpha1.PackageRevision)
			if !ok {
				return nil
			}
			return []interface{}{
				discoveredPkgRev.Name,
				discoveredPkgRev.Spec.PackageID.Repository,
				discoveredPkgRev.Spec.PackageID.Package,
				discoveredPkgRev.Spec.PackageID.Revision,
				discoveredPkgRev.Spec.PackageID.Workspace,
			}
		},
		columns: []metav1.TableColumnDefinition{
			{Name: "Name", Type: "string"},
			{Name: "Repository", Type: "string"},
			{Name: "Package", Type: "string"},
			{Name: "Revision", Type: "string"},
			{Name: "Workspace", Type: "string"},
		},
	}
}
*/