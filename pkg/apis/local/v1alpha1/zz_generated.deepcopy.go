// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	operatorv1alpha1 "github.com/openshift/api/operator/v1alpha1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskMakerImageVersion) DeepCopyInto(out *DiskMakerImageVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskMakerImageVersion.
func (in *DiskMakerImageVersion) DeepCopy() *DiskMakerImageVersion {
	if in == nil {
		return nil
	}
	out := new(DiskMakerImageVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalProvisionerImageVersion) DeepCopyInto(out *LocalProvisionerImageVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalProvisionerImageVersion.
func (in *LocalProvisionerImageVersion) DeepCopy() *LocalProvisionerImageVersion {
	if in == nil {
		return nil
	}
	out := new(LocalProvisionerImageVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageProvider) DeepCopyInto(out *LocalStorageProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageProvider.
func (in *LocalStorageProvider) DeepCopy() *LocalStorageProvider {
	if in == nil {
		return nil
	}
	out := new(LocalStorageProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalStorageProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageProviderList) DeepCopyInto(out *LocalStorageProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalStorageProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageProviderList.
func (in *LocalStorageProviderList) DeepCopy() *LocalStorageProviderList {
	if in == nil {
		return nil
	}
	out := new(LocalStorageProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalStorageProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageProviderSpec) DeepCopyInto(out *LocalStorageProviderSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.NodeSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClassDevices != nil {
		in, out := &in.StorageClassDevices, &out.StorageClassDevices
		*out = make([]StorageClassDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.LocalProvisionerImageVersion = in.LocalProvisionerImageVersion
	out.DiskMakerImageVersion = in.DiskMakerImageVersion
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageProviderSpec.
func (in *LocalStorageProviderSpec) DeepCopy() *LocalStorageProviderSpec {
	if in == nil {
		return nil
	}
	out := new(LocalStorageProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageProviderStatus) DeepCopyInto(out *LocalStorageProviderStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]operatorv1alpha1.GenerationHistory, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]operatorv1alpha1.OperatorCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageProviderStatus.
func (in *LocalStorageProviderStatus) DeepCopy() *LocalStorageProviderStatus {
	if in == nil {
		return nil
	}
	out := new(LocalStorageProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassDevice) DeepCopyInto(out *StorageClassDevice) {
	*out = *in
	if in.DeviceNames != nil {
		in, out := &in.DeviceNames, &out.DeviceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeviceUUIDs != nil {
		in, out := &in.DeviceUUIDs, &out.DeviceUUIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassDevice.
func (in *StorageClassDevice) DeepCopy() *StorageClassDevice {
	if in == nil {
		return nil
	}
	out := new(StorageClassDevice)
	in.DeepCopyInto(out)
	return out
}
