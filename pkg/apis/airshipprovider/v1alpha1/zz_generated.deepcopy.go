// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressingItems) DeepCopyInto(out *AddressingItems) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressingItems.
func (in *AddressingItems) DeepCopy() *AddressingItems {
	if in == nil {
		return nil
	}
	out := new(AddressingItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirshipClusterProviderSpec) DeepCopyInto(out *AirshipClusterProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CAPrivateKey != nil {
		in, out := &in.CAPrivateKey, &out.CAPrivateKey
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirshipClusterProviderSpec.
func (in *AirshipClusterProviderSpec) DeepCopy() *AirshipClusterProviderSpec {
	if in == nil {
		return nil
	}
	out := new(AirshipClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AirshipClusterProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirshipClusterProviderStatus) DeepCopyInto(out *AirshipClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Network.DeepCopyInto(&out.Network)
	in.Bastion.DeepCopyInto(&out.Bastion)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirshipClusterProviderStatus.
func (in *AirshipClusterProviderStatus) DeepCopy() *AirshipClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AirshipClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AirshipClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirshipMachineProviderCondition) DeepCopyInto(out *AirshipMachineProviderCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirshipMachineProviderCondition.
func (in *AirshipMachineProviderCondition) DeepCopy() *AirshipMachineProviderCondition {
	if in == nil {
		return nil
	}
	out := new(AirshipMachineProviderCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirshipMachineProviderSpec) DeepCopyInto(out *AirshipMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]MachineRole, len(*in))
		copy(*out, *in)
	}
	in.BaremetalSpec.DeepCopyInto(&out.BaremetalSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirshipMachineProviderSpec.
func (in *AirshipMachineProviderSpec) DeepCopy() *AirshipMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(AirshipMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AirshipMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirshipMachineProviderStatus) DeepCopyInto(out *AirshipMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.VMID != nil {
		in, out := &in.VMID, &out.VMID
		*out = new(string)
		**out = **in
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AirshipMachineProviderCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirshipMachineProviderStatus.
func (in *AirshipMachineProviderStatus) DeepCopy() *AirshipMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AirshipMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AirshipMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirshipResourceReference) DeepCopyInto(out *AirshipResourceReference) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirshipResourceReference.
func (in *AirshipResourceReference) DeepCopy() *AirshipResourceReference {
	if in == nil {
		return nil
	}
	out := new(AirshipResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaremetalSpec) DeepCopyInto(out *BaremetalSpec) {
	*out = *in
	if in.Addressing != nil {
		in, out := &in.Addressing, &out.Addressing
		*out = make([]*AddressingItems, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AddressingItems)
				**out = **in
			}
		}
	}
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make(map[string]*InterfacesItem, len(*in))
		for key, val := range *in {
			var outVal *InterfacesItem
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(InterfacesItem)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(Metadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Oob != nil {
		in, out := &in.Oob, &out.Oob
		*out = new(Oob)
		**out = **in
	}
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(Platform)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaremetalSpec.
func (in *BaremetalSpec) DeepCopy() *BaremetalSpec {
	if in == nil {
		return nil
	}
	out := new(BaremetalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filesystem) DeepCopyInto(out *Filesystem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filesystem.
func (in *Filesystem) DeepCopy() *Filesystem {
	if in == nil {
		return nil
	}
	out := new(Filesystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	*out = *in
	if in.CidrBlocks != nil {
		in, out := &in.CidrBlocks, &out.CidrBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceSecurityGroupIDs != nil {
		in, out := &in.SourceSecurityGroupIDs, &out.SourceSecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRule.
func (in *IngressRule) DeepCopy() *IngressRule {
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in IngressRules) DeepCopyInto(out *IngressRules) {
	{
		in := &in
		*out = make(IngressRules, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IngressRule)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRules.
func (in IngressRules) DeepCopy() IngressRules {
	if in == nil {
		return nil
	}
	out := new(IngressRules)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(string)
		**out = **in
	}
	if in.PublicIP != nil {
		in, out := &in.PublicIP, &out.PublicIP
		*out = new(string)
		**out = **in
	}
	if in.ENASupport != nil {
		in, out := &in.ENASupport, &out.ENASupport
		*out = new(bool)
		**out = **in
	}
	if in.EBSOptimized != nil {
		in, out := &in.EBSOptimized, &out.EBSOptimized
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfacesItem) DeepCopyInto(out *InterfacesItem) {
	*out = *in
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Slaves != nil {
		in, out := &in.Slaves, &out.Slaves
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfacesItem.
func (in *InterfacesItem) DeepCopy() *InterfacesItem {
	if in == nil {
		return nil
	}
	out := new(InterfacesItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelParams) DeepCopyInto(out *KernelParams) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelParams.
func (in *KernelParams) DeepCopy() *KernelParams {
	if in == nil {
		return nil
	}
	out := new(KernelParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make([]*LoadBalancerListener, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(LoadBalancerListener)
				**out = **in
			}
		}
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(LoadBalancerHealthCheck)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerHealthCheck) DeepCopyInto(out *LoadBalancerHealthCheck) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerHealthCheck.
func (in *LoadBalancerHealthCheck) DeepCopy() *LoadBalancerHealthCheck {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerListener) DeepCopyInto(out *LoadBalancerListener) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerListener.
func (in *LoadBalancerListener) DeepCopy() *LoadBalancerListener {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalVolumesItems) DeepCopyInto(out *LogicalVolumesItems) {
	*out = *in
	if in.Filesystem != nil {
		in, out := &in.Filesystem, &out.Filesystem
		*out = new(Filesystem)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalVolumesItems.
func (in *LogicalVolumesItems) DeepCopy() *LogicalVolumesItems {
	if in == nil {
		return nil
	}
	out := new(LogicalVolumesItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.OwnerData != nil {
		in, out := &in.OwnerData, &out.OwnerData
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	in.Vnet.DeepCopyInto(&out.Vnet)
	if in.InternetGatewayID != nil {
		in, out := &in.InternetGatewayID, &out.InternetGatewayID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make(map[SecurityGroupRole]*SecurityGroup, len(*in))
		for key, val := range *in {
			var outVal *SecurityGroup
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(SecurityGroup)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(Subnets, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Subnet)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.APIServerLB.DeepCopyInto(&out.APIServerLB)
	out.APIServerIP = in.APIServerIP
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Oob) DeepCopyInto(out *Oob) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Oob.
func (in *Oob) DeepCopy() *Oob {
	if in == nil {
		return nil
	}
	out := new(Oob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionsItems) DeepCopyInto(out *PartitionsItems) {
	*out = *in
	if in.Filesystem != nil {
		in, out := &in.Filesystem, &out.Filesystem
		*out = new(Filesystem)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionsItems.
func (in *PartitionsItems) DeepCopy() *PartitionsItems {
	if in == nil {
		return nil
	}
	out := new(PartitionsItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalDevicesItem) DeepCopyInto(out *PhysicalDevicesItem) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = make([]*PartitionsItems, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PartitionsItems)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalDevicesItem.
func (in *PhysicalDevicesItem) DeepCopy() *PhysicalDevicesItem {
	if in == nil {
		return nil
	}
	out := new(PhysicalDevicesItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platform) DeepCopyInto(out *Platform) {
	*out = *in
	if in.KernelParams != nil {
		in, out := &in.KernelParams, &out.KernelParams
		*out = new(KernelParams)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platform.
func (in *Platform) DeepCopy() *Platform {
	if in == nil {
		return nil
	}
	out := new(Platform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicIPAddress) DeepCopyInto(out *PublicIPAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicIPAddress.
func (in *PublicIPAddress) DeepCopy() *PublicIPAddress {
	if in == nil {
		return nil
	}
	out := new(PublicIPAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTable) DeepCopyInto(out *RouteTable) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTable.
func (in *RouteTable) DeepCopy() *RouteTable {
	if in == nil {
		return nil
	}
	out := new(RouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	if in.IngressRules != nil {
		in, out := &in.IngressRules, &out.IngressRules
		*out = make(IngressRules, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IngressRule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.PhysicalDevices != nil {
		in, out := &in.PhysicalDevices, &out.PhysicalDevices
		*out = make(map[string]*PhysicalDevicesItem, len(*in))
		for key, val := range *in {
			var outVal *PhysicalDevicesItem
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(PhysicalDevicesItem)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.VolumeGroups != nil {
		in, out := &in.VolumeGroups, &out.VolumeGroups
		*out = make(map[string]*VolumeGroupsItem, len(*in))
		for key, val := range *in {
			var outVal *VolumeGroupsItem
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(VolumeGroupsItem)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	if in.RouteTableID != nil {
		in, out := &in.RouteTableID, &out.RouteTableID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Subnets) DeepCopyInto(out *Subnets) {
	{
		in := &in
		*out = make(Subnets, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Subnet)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnets.
func (in Subnets) DeepCopy() Subnets {
	if in == nil {
		return nil
	}
	out := new(Subnets)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vnet) DeepCopyInto(out *Vnet) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vnet.
func (in *Vnet) DeepCopy() *Vnet {
	if in == nil {
		return nil
	}
	out := new(Vnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupsItem) DeepCopyInto(out *VolumeGroupsItem) {
	*out = *in
	if in.LogicalVolumes != nil {
		in, out := &in.LogicalVolumes, &out.LogicalVolumes
		*out = make([]*LogicalVolumesItems, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(LogicalVolumesItems)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupsItem.
func (in *VolumeGroupsItem) DeepCopy() *VolumeGroupsItem {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupsItem)
	in.DeepCopyInto(out)
	return out
}
