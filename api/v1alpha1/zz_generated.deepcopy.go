//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
func (in *APIServer) DeepCopyInto(out *APIServer) {
	*out = *in
	out.ArtifactScriptConfigMap = in.ArtifactScriptConfigMap
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServer.
func (in *APIServer) DeepCopy() *APIServer {
	if in == nil {
		return nil
	}
	out := new(APIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactScriptConfigMap) DeepCopyInto(out *ArtifactScriptConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactScriptConfigMap.
func (in *ArtifactScriptConfigMap) DeepCopy() *ArtifactScriptConfigMap {
	if in == nil {
		return nil
	}
	out := new(ArtifactScriptConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSPipeline) DeepCopyInto(out *DSPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSPipeline.
func (in *DSPipeline) DeepCopy() *DSPipeline {
	if in == nil {
		return nil
	}
	out := new(DSPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DSPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSPipelineList) DeepCopyInto(out *DSPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DSPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSPipelineList.
func (in *DSPipelineList) DeepCopy() *DSPipelineList {
	if in == nil {
		return nil
	}
	out := new(DSPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DSPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSPipelineSpec) DeepCopyInto(out *DSPipelineSpec) {
	*out = *in
	in.APIServer.DeepCopyInto(&out.APIServer)
	in.PersistenceAgent.DeepCopyInto(&out.PersistenceAgent)
	in.ScheduledWorkflow.DeepCopyInto(&out.ScheduledWorkflow)
	in.ViewerCRD.DeepCopyInto(&out.ViewerCRD)
	in.Database.DeepCopyInto(&out.Database)
	in.ObjectStorage.DeepCopyInto(&out.ObjectStorage)
	in.MlPipelineUI.DeepCopyInto(&out.MlPipelineUI)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSPipelineSpec.
func (in *DSPipelineSpec) DeepCopy() *DSPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(DSPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSPipelineStatus) DeepCopyInto(out *DSPipelineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSPipelineStatus.
func (in *DSPipelineStatus) DeepCopy() *DSPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(DSPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	in.MariaDB.DeepCopyInto(&out.MariaDB)
	out.ExternalDB = in.ExternalDB
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDB) DeepCopyInto(out *ExternalDB) {
	*out = *in
	out.PasswordSecret = in.PasswordSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDB.
func (in *ExternalDB) DeepCopy() *ExternalDB {
	if in == nil {
		return nil
	}
	out := new(ExternalDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalStorage) DeepCopyInto(out *ExternalStorage) {
	*out = *in
	out.S3CredentialSecret = in.S3CredentialSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalStorage.
func (in *ExternalStorage) DeepCopy() *ExternalStorage {
	if in == nil {
		return nil
	}
	out := new(ExternalStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDB) DeepCopyInto(out *MariaDB) {
	*out = *in
	out.PasswordSecret = in.PasswordSecret
	out.PVCSize = in.PVCSize.DeepCopy()
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDB.
func (in *MariaDB) DeepCopy() *MariaDB {
	if in == nil {
		return nil
	}
	out := new(MariaDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Minio) DeepCopyInto(out *Minio) {
	*out = *in
	out.S3CredentialSecret = in.S3CredentialSecret
	out.PVCSize = in.PVCSize.DeepCopy()
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Minio.
func (in *Minio) DeepCopy() *Minio {
	if in == nil {
		return nil
	}
	out := new(Minio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MlPipelineUI) DeepCopyInto(out *MlPipelineUI) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MlPipelineUI.
func (in *MlPipelineUI) DeepCopy() *MlPipelineUI {
	if in == nil {
		return nil
	}
	out := new(MlPipelineUI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorage) DeepCopyInto(out *ObjectStorage) {
	*out = *in
	in.Minio.DeepCopyInto(&out.Minio)
	out.ExternalStorage = in.ExternalStorage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorage.
func (in *ObjectStorage) DeepCopy() *ObjectStorage {
	if in == nil {
		return nil
	}
	out := new(ObjectStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceAgent) DeepCopyInto(out *PersistenceAgent) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceAgent.
func (in *PersistenceAgent) DeepCopy() *PersistenceAgent {
	if in == nil {
		return nil
	}
	out := new(PersistenceAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	in.Limits.DeepCopyInto(&out.Limits)
	in.Requests.DeepCopyInto(&out.Requests)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3CredentialSecret) DeepCopyInto(out *S3CredentialSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3CredentialSecret.
func (in *S3CredentialSecret) DeepCopy() *S3CredentialSecret {
	if in == nil {
		return nil
	}
	out := new(S3CredentialSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledWorkflow) DeepCopyInto(out *ScheduledWorkflow) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledWorkflow.
func (in *ScheduledWorkflow) DeepCopy() *ScheduledWorkflow {
	if in == nil {
		return nil
	}
	out := new(ScheduledWorkflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyValue) DeepCopyInto(out *SecretKeyValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyValue.
func (in *SecretKeyValue) DeepCopy() *SecretKeyValue {
	if in == nil {
		return nil
	}
	out := new(SecretKeyValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewerCRD) DeepCopyInto(out *ViewerCRD) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewerCRD.
func (in *ViewerCRD) DeepCopy() *ViewerCRD {
	if in == nil {
		return nil
	}
	out := new(ViewerCRD)
	in.DeepCopyInto(out)
	return out
}
