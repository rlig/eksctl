// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
	ipnet "github.com/weaveworks/eksctl/pkg/utils/ipnet"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonIAM) DeepCopyInto(out *AddonIAM) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonIAM.
func (in *AddonIAM) DeepCopy() *AddonIAM {
	if in == nil {
		return nil
	}
	out := new(AddonIAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAddons) DeepCopyInto(out *ClusterAddons) {
	*out = *in
	out.WithIAM = in.WithIAM
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAddons.
func (in *ClusterAddons) DeepCopy() *ClusterAddons {
	if in == nil {
		return nil
	}
	out := new(ClusterAddons)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(ClusterMeta)
		(*in).DeepCopyInto(*out)
	}
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		*out = new(ClusterVPC)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeGroups != nil {
		in, out := &in.NodeGroups, &out.NodeGroups
		*out = make([]*NodeGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Addons = in.Addons
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigList) DeepCopyInto(out *ClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigList.
func (in *ClusterConfigList) DeepCopy() *ClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMeta) DeepCopyInto(out *ClusterMeta) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMeta.
func (in *ClusterMeta) DeepCopy() *ClusterMeta {
	if in == nil {
		return nil
	}
	out := new(ClusterMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVPC) DeepCopyInto(out *ClusterVPC) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(map[SubnetTopology]map[string]Network, len(*in))
		for key, val := range *in {
			var outVal map[string]Network
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]Network, len(*in))
				for key, val := range *in {
					(*out)[key] = *val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.ExtraCIDRs != nil {
		in, out := &in.ExtraCIDRs, &out.ExtraCIDRs
		*out = make([]*ipnet.IPNet, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVPC.
func (in *ClusterVPC) DeepCopy() *ClusterVPC {
	if in == nil {
		return nil
	}
	out := new(ClusterVPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.CIDR != nil {
		in, out := &in.CIDR, &out.CIDR
		*out = (*in).DeepCopy()
	}
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
func (in *NodeGroup) DeepCopyInto(out *NodeGroup) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PolicyARNs != nil {
		in, out := &in.PolicyARNs, &out.PolicyARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SSHPublicKey != nil {
		in, out := &in.SSHPublicKey, &out.SSHPublicKey
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroup.
func (in *NodeGroup) DeepCopy() *NodeGroup {
	if in == nil {
		return nil
	}
	out := new(NodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfig) DeepCopyInto(out *ProviderConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfig.
func (in *ProviderConfig) DeepCopy() *ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderConfig)
	in.DeepCopyInto(out)
	return out
}
