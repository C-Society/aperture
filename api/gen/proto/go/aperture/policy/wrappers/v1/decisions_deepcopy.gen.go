// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package wrappersv1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using LoadShedDecisionWrapper within kubernetes types, where deepcopy-gen is used.
func (in *LoadShedDecisionWrapper) DeepCopyInto(out *LoadShedDecisionWrapper) {
	p := proto.Clone(in).(*LoadShedDecisionWrapper)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadShedDecisionWrapper. Required by controller-gen.
func (in *LoadShedDecisionWrapper) DeepCopy() *LoadShedDecisionWrapper {
	if in == nil {
		return nil
	}
	out := new(LoadShedDecisionWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadShedDecisionWrapper. Required by controller-gen.
func (in *LoadShedDecisionWrapper) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TokensDecisionWrapper within kubernetes types, where deepcopy-gen is used.
func (in *TokensDecisionWrapper) DeepCopyInto(out *TokensDecisionWrapper) {
	p := proto.Clone(in).(*TokensDecisionWrapper)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokensDecisionWrapper. Required by controller-gen.
func (in *TokensDecisionWrapper) DeepCopy() *TokensDecisionWrapper {
	if in == nil {
		return nil
	}
	out := new(TokensDecisionWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TokensDecisionWrapper. Required by controller-gen.
func (in *TokensDecisionWrapper) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimiterDecisionWrapper within kubernetes types, where deepcopy-gen is used.
func (in *RateLimiterDecisionWrapper) DeepCopyInto(out *RateLimiterDecisionWrapper) {
	p := proto.Clone(in).(*RateLimiterDecisionWrapper)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiterDecisionWrapper. Required by controller-gen.
func (in *RateLimiterDecisionWrapper) DeepCopy() *RateLimiterDecisionWrapper {
	if in == nil {
		return nil
	}
	out := new(RateLimiterDecisionWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiterDecisionWrapper. Required by controller-gen.
func (in *RateLimiterDecisionWrapper) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}