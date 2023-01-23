// Copyright 2023 Jeremy Edwards
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/hardware.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CpuDeviceMetrics holds details about CPU utilization and temperatures.
type CpuDeviceMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Load as a percentage [0-100].
	Load []int32 `protobuf:"varint,1,rep,packed,name=load,proto3" json:"load,omitempty"`
	// Temperature is the temperature of each core in celcius.
	Temperature []float64 `protobuf:"fixed64,2,rep,packed,name=temperature,proto3" json:"temperature,omitempty"`
	// NumCores is the total number of cores across all CPUs on the machine.
	NumCores int32 `protobuf:"varint,3,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`
	// Frequency is the clock frequency (Hz) of the CPU.
	FrequencyMhz float64 `protobuf:"fixed64,4,opt,name=frequency_mhz,json=frequencyMhz,proto3" json:"frequency_mhz,omitempty"`
	// FSBFrequency is the clock frequency of the front side bus.
	FsbFrequencyMhz float64 `protobuf:"fixed64,5,opt,name=fsb_frequency_mhz,json=fsbFrequencyMhz,proto3" json:"fsb_frequency_mhz,omitempty"`
}

func (x *CpuDeviceMetrics) Reset() {
	*x = CpuDeviceMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hardware_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuDeviceMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuDeviceMetrics) ProtoMessage() {}

func (x *CpuDeviceMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hardware_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuDeviceMetrics.ProtoReflect.Descriptor instead.
func (*CpuDeviceMetrics) Descriptor() ([]byte, []int) {
	return file_proto_hardware_proto_rawDescGZIP(), []int{0}
}

func (x *CpuDeviceMetrics) GetLoad() []int32 {
	if x != nil {
		return x.Load
	}
	return nil
}

func (x *CpuDeviceMetrics) GetTemperature() []float64 {
	if x != nil {
		return x.Temperature
	}
	return nil
}

func (x *CpuDeviceMetrics) GetNumCores() int32 {
	if x != nil {
		return x.NumCores
	}
	return 0
}

func (x *CpuDeviceMetrics) GetFrequencyMhz() float64 {
	if x != nil {
		return x.FrequencyMhz
	}
	return 0
}

func (x *CpuDeviceMetrics) GetFsbFrequencyMhz() float64 {
	if x != nil {
		return x.FsbFrequencyMhz
	}
	return 0
}

// DeviceMetrics holds the health metrics (temperature and other measurements) of the device.
type DeviceMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Kind of device that is described.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Temperature of the device in celcius if available.
	Temperature float64 `protobuf:"fixed64,3,opt,name=temperature,proto3" json:"temperature,omitempty"`
	// CPU is populated if the device is a CPU.
	Cpu *CpuDeviceMetrics `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
}

func (x *DeviceMetrics) Reset() {
	*x = DeviceMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hardware_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceMetrics) ProtoMessage() {}

func (x *DeviceMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hardware_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceMetrics.ProtoReflect.Descriptor instead.
func (*DeviceMetrics) Descriptor() ([]byte, []int) {
	return file_proto_hardware_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceMetrics) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceMetrics) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *DeviceMetrics) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *DeviceMetrics) GetCpu() *CpuDeviceMetrics {
	if x != nil {
		return x.Cpu
	}
	return nil
}

// MachineMetrics holds a list of devices that can be instrumented for health.
type MachineMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the hostname of the machine.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Device is a list of all the instrumented devices on the machine.
	Device []*DeviceMetrics `protobuf:"bytes,2,rep,name=device,proto3" json:"device,omitempty"`
	// Timestamp of the sample of metrics.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MachineMetrics) Reset() {
	*x = MachineMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hardware_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineMetrics) ProtoMessage() {}

func (x *MachineMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hardware_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineMetrics.ProtoReflect.Descriptor instead.
func (*MachineMetrics) Descriptor() ([]byte, []int) {
	return file_proto_hardware_proto_rawDescGZIP(), []int{2}
}

func (x *MachineMetrics) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MachineMetrics) GetDevice() []*DeviceMetrics {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *MachineMetrics) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_proto_hardware_proto protoreflect.FileDescriptor

var file_proto_hardware_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6a, 0x65, 0x72, 0x65, 0x6d, 0x79, 0x6a, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x10, 0x43, 0x70,
	0x75, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x72, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d,
	0x68, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x4d, 0x68, 0x7a, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x73, 0x62, 0x5f, 0x66, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x68, 0x7a, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0f, 0x66, 0x73, 0x62, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x4d,
	0x68, 0x7a, 0x22, 0x9f, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x44,
	0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6a, 0x65,
	0x72, 0x65, 0x6d, 0x79, 0x6a, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x5f,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x70, 0x75, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x03, 0x63, 0x70, 0x75, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6a, 0x65,
	0x72, 0x65, 0x6d, 0x79, 0x6a, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x5f,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x72,
	0x65, 0x6d, 0x79, 0x6a, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x2d, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hardware_proto_rawDescOnce sync.Once
	file_proto_hardware_proto_rawDescData = file_proto_hardware_proto_rawDesc
)

func file_proto_hardware_proto_rawDescGZIP() []byte {
	file_proto_hardware_proto_rawDescOnce.Do(func() {
		file_proto_hardware_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hardware_proto_rawDescData)
	})
	return file_proto_hardware_proto_rawDescData
}

var file_proto_hardware_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_hardware_proto_goTypes = []interface{}{
	(*CpuDeviceMetrics)(nil),      // 0: jeremyje.coretemp_exporter.proto.CpuDeviceMetrics
	(*DeviceMetrics)(nil),         // 1: jeremyje.coretemp_exporter.proto.DeviceMetrics
	(*MachineMetrics)(nil),        // 2: jeremyje.coretemp_exporter.proto.MachineMetrics
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proto_hardware_proto_depIdxs = []int32{
	0, // 0: jeremyje.coretemp_exporter.proto.DeviceMetrics.cpu:type_name -> jeremyje.coretemp_exporter.proto.CpuDeviceMetrics
	1, // 1: jeremyje.coretemp_exporter.proto.MachineMetrics.device:type_name -> jeremyje.coretemp_exporter.proto.DeviceMetrics
	3, // 2: jeremyje.coretemp_exporter.proto.MachineMetrics.timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_hardware_proto_init() }
func file_proto_hardware_proto_init() {
	if File_proto_hardware_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hardware_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuDeviceMetrics); i {
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
		file_proto_hardware_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceMetrics); i {
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
		file_proto_hardware_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineMetrics); i {
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
			RawDescriptor: file_proto_hardware_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_hardware_proto_goTypes,
		DependencyIndexes: file_proto_hardware_proto_depIdxs,
		MessageInfos:      file_proto_hardware_proto_msgTypes,
	}.Build()
	File_proto_hardware_proto = out.File
	file_proto_hardware_proto_rawDesc = nil
	file_proto_hardware_proto_goTypes = nil
	file_proto_hardware_proto_depIdxs = nil
}
