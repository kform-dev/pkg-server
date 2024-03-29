//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	condition "github.com/kform-dev/pkg-server/apis/condition"
	pkgid "github.com/kform-dev/pkg-server/apis/pkgid"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevision) DeepCopyInto(out *PackageRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevision.
func (in *PackageRevision) DeepCopy() *PackageRevision {
	if in == nil {
		return nil
	}
	out := new(PackageRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionApproval) DeepCopyInto(out *PackageRevisionApproval) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionApproval.
func (in *PackageRevisionApproval) DeepCopy() *PackageRevisionApproval {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionApproval)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionList) DeepCopyInto(out *PackageRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PackageRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionList.
func (in *PackageRevisionList) DeepCopy() *PackageRevisionList {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionResources) DeepCopyInto(out *PackageRevisionResources) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionResources.
func (in *PackageRevisionResources) DeepCopy() *PackageRevisionResources {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageRevisionResources) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionResourcesList) DeepCopyInto(out *PackageRevisionResourcesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PackageRevisionResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionResourcesList.
func (in *PackageRevisionResourcesList) DeepCopy() *PackageRevisionResourcesList {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionResourcesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageRevisionResourcesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionResourcesSpec) DeepCopyInto(out *PackageRevisionResourcesSpec) {
	*out = *in
	out.PackageID = in.PackageID
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionResourcesSpec.
func (in *PackageRevisionResourcesSpec) DeepCopy() *PackageRevisionResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionResourcesStatus) DeepCopyInto(out *PackageRevisionResourcesStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionResourcesStatus.
func (in *PackageRevisionResourcesStatus) DeepCopy() *PackageRevisionResourcesStatus {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionResourcesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionSpec) DeepCopyInto(out *PackageRevisionSpec) {
	*out = *in
	out.PackageID = in.PackageID
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]Task, len(*in))
		copy(*out, *in)
	}
	if in.ReadinessGates != nil {
		in, out := &in.ReadinessGates, &out.ReadinessGates
		*out = make([]condition.ReadinessGate, len(*in))
		copy(*out, *in)
	}
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = new(pkgid.Upstream)
		**out = **in
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionSpec.
func (in *PackageRevisionSpec) DeepCopy() *PackageRevisionSpec {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionStatus) DeepCopyInto(out *PackageRevisionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	in.PublishedAt.DeepCopyInto(&out.PublishedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionStatus.
func (in *PackageRevisionStatus) DeepCopy() *PackageRevisionStatus {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}
