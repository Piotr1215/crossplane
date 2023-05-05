//
//Copyright 2022 The Crossplane Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// TODO(negz): Add a make target such that `make lint` runs `buf lint`.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: apiextensions/fn/proto/v1alpha1/run_function.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ImagePullPolicy specifies when a Composition Function container should be
// pulled from a remote OCI registry.
type ImagePullPolicy int32

const (
	ImagePullPolicy_IMAGE_PULL_POLICY_UNSPECIFIED    ImagePullPolicy = 0
	ImagePullPolicy_IMAGE_PULL_POLICY_IF_NOT_PRESENT ImagePullPolicy = 1
	ImagePullPolicy_IMAGE_PULL_POLICY_ALWAYS         ImagePullPolicy = 2
	ImagePullPolicy_IMAGE_PULL_POLICY_NEVER          ImagePullPolicy = 3
)

// Enum value maps for ImagePullPolicy.
var (
	ImagePullPolicy_name = map[int32]string{
		0: "IMAGE_PULL_POLICY_UNSPECIFIED",
		1: "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
		2: "IMAGE_PULL_POLICY_ALWAYS",
		3: "IMAGE_PULL_POLICY_NEVER",
	}
	ImagePullPolicy_value = map[string]int32{
		"IMAGE_PULL_POLICY_UNSPECIFIED":    0,
		"IMAGE_PULL_POLICY_IF_NOT_PRESENT": 1,
		"IMAGE_PULL_POLICY_ALWAYS":         2,
		"IMAGE_PULL_POLICY_NEVER":          3,
	}
)

func (x ImagePullPolicy) Enum() *ImagePullPolicy {
	p := new(ImagePullPolicy)
	*p = x
	return p
}

func (x ImagePullPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImagePullPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_enumTypes[0].Descriptor()
}

func (ImagePullPolicy) Type() protoreflect.EnumType {
	return &file_apiextensions_fn_proto_v1alpha1_run_function_proto_enumTypes[0]
}

func (x ImagePullPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImagePullPolicy.Descriptor instead.
func (ImagePullPolicy) EnumDescriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{0}
}

// NetworkPolicy configures whether a container is isolated from the network.
type NetworkPolicy int32

const (
	NetworkPolicy_NETWORK_POLICY_UNSPECIFIED NetworkPolicy = 0
	// Run the container without network access. The default.
	NetworkPolicy_NETWORK_POLICY_ISOLATED NetworkPolicy = 1
	// Allow the container to access the same network as the function runner.
	NetworkPolicy_NETWORK_POLICY_RUNNER NetworkPolicy = 2
)

// Enum value maps for NetworkPolicy.
var (
	NetworkPolicy_name = map[int32]string{
		0: "NETWORK_POLICY_UNSPECIFIED",
		1: "NETWORK_POLICY_ISOLATED",
		2: "NETWORK_POLICY_RUNNER",
	}
	NetworkPolicy_value = map[string]int32{
		"NETWORK_POLICY_UNSPECIFIED": 0,
		"NETWORK_POLICY_ISOLATED":    1,
		"NETWORK_POLICY_RUNNER":      2,
	}
)

func (x NetworkPolicy) Enum() *NetworkPolicy {
	p := new(NetworkPolicy)
	*p = x
	return p
}

func (x NetworkPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_enumTypes[1].Descriptor()
}

func (NetworkPolicy) Type() protoreflect.EnumType {
	return &file_apiextensions_fn_proto_v1alpha1_run_function_proto_enumTypes[1]
}

func (x NetworkPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkPolicy.Descriptor instead.
func (NetworkPolicy) EnumDescriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{1}
}

// ImagePullAuth configures authentication to a remote OCI registry.
// It corresponds to go-containerregistry's AuthConfig type.
// https://pkg.go.dev/github.com/google/go-containerregistry@v0.11.0/pkg/authn#AuthConfig
type ImagePullAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Auth          string `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	IdentityToken string `protobuf:"bytes,4,opt,name=identity_token,json=identityToken,proto3" json:"identity_token,omitempty"`
	RegistryToken string `protobuf:"bytes,5,opt,name=registry_token,json=registryToken,proto3" json:"registry_token,omitempty"`
}

func (x *ImagePullAuth) Reset() {
	*x = ImagePullAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagePullAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagePullAuth) ProtoMessage() {}

func (x *ImagePullAuth) ProtoReflect() protoreflect.Message {
	mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagePullAuth.ProtoReflect.Descriptor instead.
func (*ImagePullAuth) Descriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{0}
}

func (x *ImagePullAuth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ImagePullAuth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ImagePullAuth) GetAuth() string {
	if x != nil {
		return x.Auth
	}
	return ""
}

func (x *ImagePullAuth) GetIdentityToken() string {
	if x != nil {
		return x.IdentityToken
	}
	return ""
}

func (x *ImagePullAuth) GetRegistryToken() string {
	if x != nil {
		return x.RegistryToken
	}
	return ""
}

// ImagePullConfig configures how a Composition Function container should be
// pulled from a remote OCI registry.
type ImagePullConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PullPolicy ImagePullPolicy `protobuf:"varint,1,opt,name=pull_policy,json=pullPolicy,proto3,enum=apiextensions.fn.proto.v1alpha1.ImagePullPolicy" json:"pull_policy,omitempty"`
	Auth       *ImagePullAuth  `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *ImagePullConfig) Reset() {
	*x = ImagePullConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagePullConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagePullConfig) ProtoMessage() {}

func (x *ImagePullConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagePullConfig.ProtoReflect.Descriptor instead.
func (*ImagePullConfig) Descriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{1}
}

func (x *ImagePullConfig) GetPullPolicy() ImagePullPolicy {
	if x != nil {
		return x.PullPolicy
	}
	return ImagePullPolicy_IMAGE_PULL_POLICY_UNSPECIFIED
}

func (x *ImagePullConfig) GetAuth() *ImagePullAuth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// NetworkConfig configures whether and how a Composition Function container may
// access the network.
type NetworkConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not the container can access the network.
	Policy NetworkPolicy `protobuf:"varint,1,opt,name=policy,proto3,enum=apiextensions.fn.proto.v1alpha1.NetworkPolicy" json:"policy,omitempty"`
}

func (x *NetworkConfig) Reset() {
	*x = NetworkConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkConfig) ProtoMessage() {}

func (x *NetworkConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkConfig.ProtoReflect.Descriptor instead.
func (*NetworkConfig) Descriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkConfig) GetPolicy() NetworkPolicy {
	if x != nil {
		return x.Policy
	}
	return NetworkPolicy_NETWORK_POLICY_UNSPECIFIED
}

// Resources configures what compute resources should be available to a
// Composition Function container.
type ResourceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limits *ResourceLimits `protobuf:"bytes,1,opt,name=limits,proto3" json:"limits,omitempty"`
}

func (x *ResourceConfig) Reset() {
	*x = ResourceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceConfig) ProtoMessage() {}

func (x *ResourceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceConfig.ProtoReflect.Descriptor instead.
func (*ResourceConfig) Descriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceConfig) GetLimits() *ResourceLimits {
	if x != nil {
		return x.Limits
	}
	return nil
}

// ResourceLimits configures the maximum compute resources that will be
// available to a Composition Function container.
type ResourceLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CPU, in cores. (500m = .5 cores)
	// Specified in Kubernetes-style resource.Quantity form.
	Memory string `protobuf:"bytes,1,opt,name=memory,proto3" json:"memory,omitempty"`
	// Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	// Specified in Kubernetes-style resource.Quantity form.
	Cpu string `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
}

func (x *ResourceLimits) Reset() {
	*x = ResourceLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLimits) ProtoMessage() {}

func (x *ResourceLimits) ProtoReflect() protoreflect.Message {
	mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLimits.ProtoReflect.Descriptor instead.
func (*ResourceLimits) Descriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{4}
}

func (x *ResourceLimits) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *ResourceLimits) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

// RunFunctionConfig configures how a Composition Function container is run.
type RunFunctionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resources available to the container.
	Resources *ResourceConfig `protobuf:"bytes,1,opt,name=resources,proto3" json:"resources,omitempty"`
	// Network configuration for the container.
	Network *NetworkConfig `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	// Timeout after which the container will be killed.
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *RunFunctionConfig) Reset() {
	*x = RunFunctionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunFunctionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunFunctionConfig) ProtoMessage() {}

func (x *RunFunctionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunFunctionConfig.ProtoReflect.Descriptor instead.
func (*RunFunctionConfig) Descriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{5}
}

func (x *RunFunctionConfig) GetResources() *ResourceConfig {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *RunFunctionConfig) GetNetwork() *NetworkConfig {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *RunFunctionConfig) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// A RunFunctionRequest requests that a Composition Function be run.
type RunFunctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OCI image of the Composition Function.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// A FunctionIO serialized as YAML.
	Input []byte `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	// Configures how the function image is pulled.
	ImagePullConfig *ImagePullConfig `protobuf:"bytes,3,opt,name=image_pull_config,json=imagePullConfig,proto3" json:"image_pull_config,omitempty"`
	// Configures how the function container is run.
	RunFunctionConfig *RunFunctionConfig `protobuf:"bytes,4,opt,name=run_function_config,json=runFunctionConfig,proto3" json:"run_function_config,omitempty"`
}

func (x *RunFunctionRequest) Reset() {
	*x = RunFunctionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunFunctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunFunctionRequest) ProtoMessage() {}

func (x *RunFunctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunFunctionRequest.ProtoReflect.Descriptor instead.
func (*RunFunctionRequest) Descriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{6}
}

func (x *RunFunctionRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *RunFunctionRequest) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *RunFunctionRequest) GetImagePullConfig() *ImagePullConfig {
	if x != nil {
		return x.ImagePullConfig
	}
	return nil
}

func (x *RunFunctionRequest) GetRunFunctionConfig() *RunFunctionConfig {
	if x != nil {
		return x.RunFunctionConfig
	}
	return nil
}

// A RunFunctionResponse contains the response from a Composition Function run.
// The output FunctionIO is returned as opaque bytes. Errors encountered while
// running a function (as opposed to errors returned _by_ a function) will be
// encapsulated as gRPC errors.
type RunFunctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *RunFunctionResponse) Reset() {
	*x = RunFunctionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunFunctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunFunctionResponse) ProtoMessage() {}

func (x *RunFunctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunFunctionResponse.ProtoReflect.Descriptor instead.
func (*RunFunctionResponse) Descriptor() ([]byte, []int) {
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP(), []int{7}
}

func (x *RunFunctionResponse) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_apiextensions_fn_proto_v1alpha1_run_function_proto protoreflect.FileDescriptor

var file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDesc = []byte{
	0x0a, 0x32, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x66, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50,
	0x75, 0x6c, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0xa8, 0x01, 0x0a, 0x0f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a, 0x0b, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x61, 0x70, 0x69,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0a, 0x70, 0x75,
	0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x42, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75,
	0x6c, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x57, 0x0a, 0x0d,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x0a,
	0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x59, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x22, 0x3a, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70,
	0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x22, 0xe1, 0x01, 0x0a,
	0x11, 0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x4d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x48, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x33, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0x82, 0x02, 0x0a, 0x12, 0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x5c, 0x0a, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x75, 0x6c,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x62, 0x0a, 0x13, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x11, 0x72, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x2d, 0x0a, 0x13, 0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2a, 0x95, 0x01, 0x0a, 0x0f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75,
	0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4d, 0x41, 0x47,
	0x45, 0x5f, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59,
	0x5f, 0x49, 0x46, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x55, 0x4c, 0x4c, 0x5f,
	0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x41, 0x4c, 0x57, 0x41, 0x59, 0x53, 0x10, 0x02, 0x12,
	0x1b, 0x0a, 0x17, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x50, 0x4f,
	0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4e, 0x45, 0x56, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x67, 0x0a, 0x0d,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1e, 0x0a,
	0x1a, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a,
	0x17, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f,
	0x49, 0x53, 0x4f, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x45,
	0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x52, 0x55, 0x4e,
	0x4e, 0x45, 0x52, 0x10, 0x02, 0x32, 0xa0, 0x01, 0x0a, 0x22, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x0b,
	0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x75,
	0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x66, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescOnce sync.Once
	file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescData = file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDesc
)

func file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescGZIP() []byte {
	file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescOnce.Do(func() {
		file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescData)
	})
	return file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDescData
}

var file_apiextensions_fn_proto_v1alpha1_run_function_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apiextensions_fn_proto_v1alpha1_run_function_proto_goTypes = []interface{}{
	(ImagePullPolicy)(0),        // 0: apiextensions.fn.proto.v1alpha1.ImagePullPolicy
	(NetworkPolicy)(0),          // 1: apiextensions.fn.proto.v1alpha1.NetworkPolicy
	(*ImagePullAuth)(nil),       // 2: apiextensions.fn.proto.v1alpha1.ImagePullAuth
	(*ImagePullConfig)(nil),     // 3: apiextensions.fn.proto.v1alpha1.ImagePullConfig
	(*NetworkConfig)(nil),       // 4: apiextensions.fn.proto.v1alpha1.NetworkConfig
	(*ResourceConfig)(nil),      // 5: apiextensions.fn.proto.v1alpha1.ResourceConfig
	(*ResourceLimits)(nil),      // 6: apiextensions.fn.proto.v1alpha1.ResourceLimits
	(*RunFunctionConfig)(nil),   // 7: apiextensions.fn.proto.v1alpha1.RunFunctionConfig
	(*RunFunctionRequest)(nil),  // 8: apiextensions.fn.proto.v1alpha1.RunFunctionRequest
	(*RunFunctionResponse)(nil), // 9: apiextensions.fn.proto.v1alpha1.RunFunctionResponse
	(*durationpb.Duration)(nil), // 10: google.protobuf.Duration
}
var file_apiextensions_fn_proto_v1alpha1_run_function_proto_depIdxs = []int32{
	0,  // 0: apiextensions.fn.proto.v1alpha1.ImagePullConfig.pull_policy:type_name -> apiextensions.fn.proto.v1alpha1.ImagePullPolicy
	2,  // 1: apiextensions.fn.proto.v1alpha1.ImagePullConfig.auth:type_name -> apiextensions.fn.proto.v1alpha1.ImagePullAuth
	1,  // 2: apiextensions.fn.proto.v1alpha1.NetworkConfig.policy:type_name -> apiextensions.fn.proto.v1alpha1.NetworkPolicy
	6,  // 3: apiextensions.fn.proto.v1alpha1.ResourceConfig.limits:type_name -> apiextensions.fn.proto.v1alpha1.ResourceLimits
	5,  // 4: apiextensions.fn.proto.v1alpha1.RunFunctionConfig.resources:type_name -> apiextensions.fn.proto.v1alpha1.ResourceConfig
	4,  // 5: apiextensions.fn.proto.v1alpha1.RunFunctionConfig.network:type_name -> apiextensions.fn.proto.v1alpha1.NetworkConfig
	10, // 6: apiextensions.fn.proto.v1alpha1.RunFunctionConfig.timeout:type_name -> google.protobuf.Duration
	3,  // 7: apiextensions.fn.proto.v1alpha1.RunFunctionRequest.image_pull_config:type_name -> apiextensions.fn.proto.v1alpha1.ImagePullConfig
	7,  // 8: apiextensions.fn.proto.v1alpha1.RunFunctionRequest.run_function_config:type_name -> apiextensions.fn.proto.v1alpha1.RunFunctionConfig
	8,  // 9: apiextensions.fn.proto.v1alpha1.ContainerizedFunctionRunnerService.RunFunction:input_type -> apiextensions.fn.proto.v1alpha1.RunFunctionRequest
	9,  // 10: apiextensions.fn.proto.v1alpha1.ContainerizedFunctionRunnerService.RunFunction:output_type -> apiextensions.fn.proto.v1alpha1.RunFunctionResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_apiextensions_fn_proto_v1alpha1_run_function_proto_init() }
func file_apiextensions_fn_proto_v1alpha1_run_function_proto_init() {
	if File_apiextensions_fn_proto_v1alpha1_run_function_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagePullAuth); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagePullConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLimits); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunFunctionConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunFunctionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunFunctionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apiextensions_fn_proto_v1alpha1_run_function_proto_goTypes,
		DependencyIndexes: file_apiextensions_fn_proto_v1alpha1_run_function_proto_depIdxs,
		EnumInfos:         file_apiextensions_fn_proto_v1alpha1_run_function_proto_enumTypes,
		MessageInfos:      file_apiextensions_fn_proto_v1alpha1_run_function_proto_msgTypes,
	}.Build()
	File_apiextensions_fn_proto_v1alpha1_run_function_proto = out.File
	file_apiextensions_fn_proto_v1alpha1_run_function_proto_rawDesc = nil
	file_apiextensions_fn_proto_v1alpha1_run_function_proto_goTypes = nil
	file_apiextensions_fn_proto_v1alpha1_run_function_proto_depIdxs = nil
}
