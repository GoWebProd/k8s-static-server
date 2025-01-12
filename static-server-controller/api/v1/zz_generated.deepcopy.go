//go:build !ignore_autogenerated

/*
Copyright 2025.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRepo) DeepCopyInto(out *StaticRepo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRepo.
func (in *StaticRepo) DeepCopy() *StaticRepo {
	if in == nil {
		return nil
	}
	out := new(StaticRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StaticRepo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRepoList) DeepCopyInto(out *StaticRepoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StaticRepo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRepoList.
func (in *StaticRepoList) DeepCopy() *StaticRepoList {
	if in == nil {
		return nil
	}
	out := new(StaticRepoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StaticRepoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRepoSpec) DeepCopyInto(out *StaticRepoSpec) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(StaticRepoSpecGit)
		**out = **in
	}
	if in.Cmd != nil {
		in, out := &in.Cmd, &out.Cmd
		*out = new(StaticRepoSpecCmd)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageDir != nil {
		in, out := &in.ImageDir, &out.ImageDir
		*out = new(StaticRepoSpecImageDir)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRepoSpec.
func (in *StaticRepoSpec) DeepCopy() *StaticRepoSpec {
	if in == nil {
		return nil
	}
	out := new(StaticRepoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRepoSpecCmd) DeepCopyInto(out *StaticRepoSpecCmd) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRepoSpecCmd.
func (in *StaticRepoSpecCmd) DeepCopy() *StaticRepoSpecCmd {
	if in == nil {
		return nil
	}
	out := new(StaticRepoSpecCmd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRepoSpecGit) DeepCopyInto(out *StaticRepoSpecGit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRepoSpecGit.
func (in *StaticRepoSpecGit) DeepCopy() *StaticRepoSpecGit {
	if in == nil {
		return nil
	}
	out := new(StaticRepoSpecGit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRepoSpecImageDir) DeepCopyInto(out *StaticRepoSpecImageDir) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRepoSpecImageDir.
func (in *StaticRepoSpecImageDir) DeepCopy() *StaticRepoSpecImageDir {
	if in == nil {
		return nil
	}
	out := new(StaticRepoSpecImageDir)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticRepoStatus) DeepCopyInto(out *StaticRepoStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticRepoStatus.
func (in *StaticRepoStatus) DeepCopy() *StaticRepoStatus {
	if in == nil {
		return nil
	}
	out := new(StaticRepoStatus)
	in.DeepCopyInto(out)
	return out
}
