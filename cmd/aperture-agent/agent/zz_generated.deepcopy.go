//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package agent

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentOTELConfig) DeepCopyInto(out *AgentOTELConfig) {
	*out = *in
	in.CommonOTELConfig.DeepCopyInto(&out.CommonOTELConfig)
	in.BatchPrerollup.DeepCopyInto(&out.BatchPrerollup)
	in.BatchPostrollup.DeepCopyInto(&out.BatchPostrollup)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentOTELConfig.
func (in *AgentOTELConfig) DeepCopy() *AgentOTELConfig {
	if in == nil {
		return nil
	}
	out := new(AgentOTELConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchPostrollupConfig) DeepCopyInto(out *BatchPostrollupConfig) {
	*out = *in
	in.Timeout.DeepCopyInto(&out.Timeout)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchPostrollupConfig.
func (in *BatchPostrollupConfig) DeepCopy() *BatchPostrollupConfig {
	if in == nil {
		return nil
	}
	out := new(BatchPostrollupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchPrerollupConfig) DeepCopyInto(out *BatchPrerollupConfig) {
	*out = *in
	in.Timeout.DeepCopyInto(&out.Timeout)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchPrerollupConfig.
func (in *BatchPrerollupConfig) DeepCopy() *BatchPrerollupConfig {
	if in == nil {
		return nil
	}
	out := new(BatchPrerollupConfig)
	in.DeepCopyInto(out)
	return out
}