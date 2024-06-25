//go:build !ignore_autogenerated

/*
Copyright 2024.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsOIDCSpec) DeepCopyInto(out *AwsOIDCSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsOIDCSpec.
func (in *AwsOIDCSpec) DeepCopy() *AwsOIDCSpec {
	if in == nil {
		return nil
	}
	out := new(AwsOIDCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsProviderSpec) DeepCopyInto(out *AwsProviderSpec) {
	*out = *in
	if in.Oidc != nil {
		in, out := &in.Oidc, &out.Oidc
		*out = new(AwsOIDCSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsProviderSpec.
func (in *AwsProviderSpec) DeepCopy() *AwsProviderSpec {
	if in == nil {
		return nil
	}
	out := new(AwsProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomProviderSpec) DeepCopyInto(out *CustomProviderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomProviderSpec.
func (in *CustomProviderSpec) DeepCopy() *CustomProviderSpec {
	if in == nil {
		return nil
	}
	out := new(CustomProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecret) DeepCopyInto(out *ExternalSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecret.
func (in *ExternalSecret) DeepCopy() *ExternalSecret {
	if in == nil {
		return nil
	}
	out := new(ExternalSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretAccess) DeepCopyInto(out *ExternalSecretAccess) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretAccess.
func (in *ExternalSecretAccess) DeepCopy() *ExternalSecretAccess {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecretAccess) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretAccessList) DeepCopyInto(out *ExternalSecretAccessList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalSecretAccess, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretAccessList.
func (in *ExternalSecretAccessList) DeepCopy() *ExternalSecretAccessList {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretAccessList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecretAccessList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretAccessSpec) DeepCopyInto(out *ExternalSecretAccessSpec) {
	*out = *in
	if in.AccessSubjects != nil {
		in, out := &in.AccessSubjects, &out.AccessSubjects
		*out = make([]SecretAccessSubject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretAccessSpec.
func (in *ExternalSecretAccessSpec) DeepCopy() *ExternalSecretAccessSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretAccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretAccessStatus) DeepCopyInto(out *ExternalSecretAccessStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]SecretAccessSubject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccountAnnotation != nil {
		in, out := &in.ServiceAccountAnnotation, &out.ServiceAccountAnnotation
		*out = new(string)
		**out = **in
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretAccessStatus.
func (in *ExternalSecretAccessStatus) DeepCopy() *ExternalSecretAccessStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretAccessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretList) DeepCopyInto(out *ExternalSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretList.
func (in *ExternalSecretList) DeepCopy() *ExternalSecretList {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretProvider) DeepCopyInto(out *ExternalSecretProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretProvider.
func (in *ExternalSecretProvider) DeepCopy() *ExternalSecretProvider {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecretProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretProviderList) DeepCopyInto(out *ExternalSecretProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalSecretProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretProviderList.
func (in *ExternalSecretProviderList) DeepCopy() *ExternalSecretProviderList {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecretProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretProviderSpec) DeepCopyInto(out *ExternalSecretProviderSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretProviderSpec.
func (in *ExternalSecretProviderSpec) DeepCopy() *ExternalSecretProviderSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretProviderStatus) DeepCopyInto(out *ExternalSecretProviderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretProviderStatus.
func (in *ExternalSecretProviderStatus) DeepCopy() *ExternalSecretProviderStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretSpec) DeepCopyInto(out *ExternalSecretSpec) {
	*out = *in
	if in.SecretString != nil {
		in, out := &in.SecretString, &out.SecretString
		*out = new(string)
		**out = **in
	}
	if in.SecretName != nil {
		in, out := &in.SecretName, &out.SecretName
		*out = new(string)
		**out = **in
	}
	if in.Random != nil {
		in, out := &in.Random, &out.Random
		*out = new(RandomSecretSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderSpec != nil {
		in, out := &in.ProviderSpec, &out.ProviderSpec
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretSpec.
func (in *ExternalSecretSpec) DeepCopy() *ExternalSecretSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretStatus) DeepCopyInto(out *ExternalSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeletionDate != nil {
		in, out := &in.DeletionDate, &out.DeletionDate
		*out = (*in).DeepCopy()
	}
	if in.NextRotateDate != nil {
		in, out := &in.NextRotateDate, &out.NextRotateDate
		*out = (*in).DeepCopy()
	}
	if in.RandomGenRegex != nil {
		in, out := &in.RandomGenRegex, &out.RandomGenRegex
		*out = new(string)
		**out = **in
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretStatus.
func (in *ExternalSecretStatus) DeepCopy() *ExternalSecretStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RandomSecretSpec) DeepCopyInto(out *RandomSecretSpec) {
	*out = *in
	if in.Rotate != nil {
		in, out := &in.Rotate, &out.Rotate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RandomSecretSpec.
func (in *RandomSecretSpec) DeepCopy() *RandomSecretSpec {
	if in == nil {
		return nil
	}
	out := new(RandomSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretAccessSubject) DeepCopyInto(out *SecretAccessSubject) {
	*out = *in
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(SecretAccessSubjectServiceAccount)
		**out = **in
	}
	if in.ProviderIdentifier != nil {
		in, out := &in.ProviderIdentifier, &out.ProviderIdentifier
		*out = new(SecretAccessSubjectProviderIdentifier)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretAccessSubject.
func (in *SecretAccessSubject) DeepCopy() *SecretAccessSubject {
	if in == nil {
		return nil
	}
	out := new(SecretAccessSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretAccessSubjectProviderIdentifier) DeepCopyInto(out *SecretAccessSubjectProviderIdentifier) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretAccessSubjectProviderIdentifier.
func (in *SecretAccessSubjectProviderIdentifier) DeepCopy() *SecretAccessSubjectProviderIdentifier {
	if in == nil {
		return nil
	}
	out := new(SecretAccessSubjectProviderIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretAccessSubjectServiceAccount) DeepCopyInto(out *SecretAccessSubjectServiceAccount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretAccessSubjectServiceAccount.
func (in *SecretAccessSubjectServiceAccount) DeepCopy() *SecretAccessSubjectServiceAccount {
	if in == nil {
		return nil
	}
	out := new(SecretAccessSubjectServiceAccount)
	in.DeepCopyInto(out)
	return out
}
