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

package pkgid

import (
	"fmt"
	"path"
	"strings"
)

const (
	PkgTarget_Catalog = "catalog"
)

// +k8s:openapi-gen=true
type PackageID struct {
	// Target defines the target for the package; not relevant for catalog packages
	// e.g. a cluster
	Target string `json:"target,omitempty" protobuf:"bytes,1,opt,name=target"`
	// Repository defines the name of the Repository object containing this package.
	Repository string `json:"repository,omitempty" protobuf:"bytes,2,opt,name=repository"`
	// Realm defines the scope in which the package is relevant
	Realm string `json:"realm,omitempty" protobuf:"bytes,3,opt,name=realm"`
	// Package defines the name of the package in the repository.
	Package string `json:"package,omitempty" protobuf:"bytes,4,opt,name=package"`
	// Revision defines the revision of the package once published
	Revision string `json:"revision,omitempty" protobuf:"bytes,5,opt,name=revision"`
	// Workspace defines the workspace of the package
	Workspace string `json:"workspace,omitempty" protobuf:"bytes,6,opt,name=workspace"`
}

func ParsePkgID(pkgstr string) (*PackageID, error) {
	parts := strings.Split(pkgstr, ".")
	if len(parts) != 5 {
		return &PackageID{}, fmt.Errorf("pkgString should contain 3 parameters <TARGET>.<REPO>.<REALM>.<PKG>.<WORKSPACE>, got: %s", pkgstr)
	}
	return &PackageID{
		Target:     parts[0],
		Repository: parts[1],
		Realm:      parts[2],
		Package:    parts[3],
		Workspace:  parts[4],
	}, nil
}

func (r *PackageID) PkgRevString() string {
	return fmt.Sprintf("%s.%s.%s.%s.%s", r.Target, r.Repository, RealmToName(r.Realm), r.Package, r.Workspace)
}

func (r *PackageID) Path() string {
	if r.Target == "catalog" {
		return path.Join(r.Realm, r.Package)
	}
	return path.Join(r.Target, r.Realm, r.Package)
}

func (r *PackageID) PkgString() string {
	return path.Join(r.Realm, r.Package)
}

func GetRepoNameFromPkgRevName(name string) string {
	return strings.Split(name, ".")[1]
}

func RealmToName(pkg string) string {
	return strings.ReplaceAll(pkg, "/", ":")
}

func PackageToDir(pkg string) string {
	return strings.ReplaceAll(pkg, ":", "/")
}

func (r PackageID) Branch() string {
	return path.Join(r.Target, r.Realm, r.Package, r.Workspace)
}

func (r PackageID) GitRevision() string {
	return path.Join(r.Target, r.Realm, r.Package, r.Revision)
}

func (r PackageID) OutDir() string {
	return path.Join(r.Target, r.Realm, r.Package, "out")
}

func (r PackageID) DNSName() string {
	// had to trim this to please  config-management
	//return fmt.Sprintf("%s.%s.%s", r.Target, strings.ReplaceAll(r.Realm, "/", "."), r.Package)
	return fmt.Sprintf("%s.%s", r.Target, r.Package)
}
