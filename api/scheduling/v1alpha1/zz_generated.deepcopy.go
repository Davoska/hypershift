//go:build !ignore_autogenerated

/*


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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSizingConfiguration) DeepCopyInto(out *ClusterSizingConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSizingConfiguration.
func (in *ClusterSizingConfiguration) DeepCopy() *ClusterSizingConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClusterSizingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSizingConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSizingConfigurationList) DeepCopyInto(out *ClusterSizingConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSizingConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSizingConfigurationList.
func (in *ClusterSizingConfigurationList) DeepCopy() *ClusterSizingConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ClusterSizingConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSizingConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSizingConfigurationSpec) DeepCopyInto(out *ClusterSizingConfigurationSpec) {
	*out = *in
	if in.Sizes != nil {
		in, out := &in.Sizes, &out.Sizes
		*out = make([]SizeConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Concurrency = in.Concurrency
	out.TransitionDelay = in.TransitionDelay
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSizingConfigurationSpec.
func (in *ClusterSizingConfigurationSpec) DeepCopy() *ClusterSizingConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSizingConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSizingConfigurationStatus) DeepCopyInto(out *ClusterSizingConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSizingConfigurationStatus.
func (in *ClusterSizingConfigurationStatus) DeepCopy() *ClusterSizingConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSizingConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConcurrencyConfiguration) DeepCopyInto(out *ConcurrencyConfiguration) {
	*out = *in
	out.SlidingWindow = in.SlidingWindow
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConcurrencyConfiguration.
func (in *ConcurrencyConfiguration) DeepCopy() *ConcurrencyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ConcurrencyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Effects) DeepCopyInto(out *Effects) {
	*out = *in
	if in.KASGoMemLimit != nil {
		in, out := &in.KASGoMemLimit, &out.KASGoMemLimit
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.ControlPlanePriorityClassName != nil {
		in, out := &in.ControlPlanePriorityClassName, &out.ControlPlanePriorityClassName
		*out = new(string)
		**out = **in
	}
	if in.EtcdPriorityClassName != nil {
		in, out := &in.EtcdPriorityClassName, &out.EtcdPriorityClassName
		*out = new(string)
		**out = **in
	}
	if in.APICriticalPriorityClassName != nil {
		in, out := &in.APICriticalPriorityClassName, &out.APICriticalPriorityClassName
		*out = new(string)
		**out = **in
	}
	if in.ResourceRequests != nil {
		in, out := &in.ResourceRequests, &out.ResourceRequests
		*out = make([]ResourceRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Effects.
func (in *Effects) DeepCopy() *Effects {
	if in == nil {
		return nil
	}
	out := new(Effects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Management) DeepCopyInto(out *Management) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Management.
func (in *Management) DeepCopy() *Management {
	if in == nil {
		return nil
	}
	out := new(Management)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeCountCriteria) DeepCopyInto(out *NodeCountCriteria) {
	*out = *in
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeCountCriteria.
func (in *NodeCountCriteria) DeepCopy() *NodeCountCriteria {
	if in == nil {
		return nil
	}
	out := new(NodeCountCriteria)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequest) DeepCopyInto(out *ResourceRequest) {
	*out = *in
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequest.
func (in *ResourceRequest) DeepCopy() *ResourceRequest {
	if in == nil {
		return nil
	}
	out := new(ResourceRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SizeConfiguration) DeepCopyInto(out *SizeConfiguration) {
	*out = *in
	in.Criteria.DeepCopyInto(&out.Criteria)
	if in.Effects != nil {
		in, out := &in.Effects, &out.Effects
		*out = new(Effects)
		(*in).DeepCopyInto(*out)
	}
	if in.Management != nil {
		in, out := &in.Management, &out.Management
		*out = new(Management)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SizeConfiguration.
func (in *SizeConfiguration) DeepCopy() *SizeConfiguration {
	if in == nil {
		return nil
	}
	out := new(SizeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitionDelayConfiguration) DeepCopyInto(out *TransitionDelayConfiguration) {
	*out = *in
	out.Increase = in.Increase
	out.Decrease = in.Decrease
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitionDelayConfiguration.
func (in *TransitionDelayConfiguration) DeepCopy() *TransitionDelayConfiguration {
	if in == nil {
		return nil
	}
	out := new(TransitionDelayConfiguration)
	in.DeepCopyInto(out)
	return out
}
