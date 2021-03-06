// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dft.proto

/*
Package dft is a generated protocol buffer package.

It is generated from these files:
	dft.proto

It has these top-level messages:
	Task
	Job
	Client
	Result
*/
package dft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	TaskId      string `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Repetitions string `protobuf:"bytes,3,opt,name=repetitions" json:"repetitions,omitempty"`
	Destination string `protobuf:"bytes,4,opt,name=destination" json:"destination,omitempty"`
	Timeout     string `protobuf:"bytes,5,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *Task) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Task) GetRepetitions() string {
	if m != nil {
		return m.Repetitions
	}
	return ""
}

func (m *Task) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *Task) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

type Job struct {
	ClientId string  `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	JobId    string  `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Tasks    []*Task `protobuf:"bytes,3,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Job) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Job) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Job) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Client struct {
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Jobs     []*Job `protobuf:"bytes,2,rep,name=jobs" json:"jobs,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Client) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Client) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type Result struct {
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	JobId    string `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Results  string `protobuf:"bytes,3,opt,name=results" json:"results,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Result) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Result) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Result) GetResults() string {
	if m != nil {
		return m.Results
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "Task")
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*Client)(nil), "Client")
	proto.RegisterType((*Result)(nil), "Result")
}

func init() { proto.RegisterFile("dft.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x95, 0xc6, 0x71, 0xc8, 0x75, 0xb3, 0x84, 0x38, 0xa9, 0x4b, 0x95, 0xa9, 0x53, 0x06,
	0x78, 0x00, 0x06, 0xa6, 0x74, 0x8c, 0x50, 0x57, 0x14, 0x63, 0x23, 0x39, 0x2d, 0xb9, 0x28, 0xbe,
	0x0e, 0x3c, 0x06, 0x6f, 0x8c, 0x7c, 0x51, 0x44, 0x27, 0x06, 0x36, 0xff, 0xbe, 0xfb, 0xa3, 0xcf,
	0x36, 0x54, 0xee, 0x83, 0x9b, 0x69, 0x26, 0xa6, 0xfa, 0x3b, 0x03, 0xf5, 0xda, 0xc7, 0xb3, 0x79,
	0x80, 0x92, 0xfb, 0x78, 0x7e, 0x0b, 0x0e, 0xb3, 0x7d, 0x76, 0xa8, 0x3a, 0x9d, 0x62, 0xeb, 0x8c,
	0x01, 0xc5, 0x5f, 0x93, 0xc7, 0x8d, 0x50, 0x39, 0x9b, 0x3d, 0x6c, 0x67, 0x3f, 0x79, 0x0e, 0x1c,
	0x68, 0x8c, 0x98, 0x4b, 0xe9, 0x16, 0xa5, 0x0e, 0xe7, 0x23, 0x87, 0xb1, 0x4f, 0x19, 0xd5, 0xd2,
	0x71, 0x83, 0x0c, 0x42, 0xc9, 0xe1, 0xd3, 0xd3, 0x95, 0xb1, 0x90, 0xea, 0x1a, 0xeb, 0x13, 0xe4,
	0x47, 0xb2, 0x66, 0x07, 0xd5, 0xfb, 0x25, 0xf8, 0x91, 0x7f, 0x9d, 0xee, 0x16, 0xd0, 0x3a, 0x73,
	0x0f, 0x7a, 0x20, 0x9b, 0x2a, 0x8b, 0x57, 0x31, 0x90, 0x6d, 0x9d, 0xd9, 0x41, 0x91, 0xb4, 0x93,
	0x52, 0x7e, 0xd8, 0x3e, 0x16, 0x4d, 0xba, 0x5b, 0xb7, 0xb0, 0xfa, 0x19, 0xf4, 0x8b, 0xcc, 0xff,
	0xbd, 0x1a, 0x41, 0x0d, 0x64, 0x23, 0x6e, 0x64, 0x85, 0x6a, 0x8e, 0x64, 0x3b, 0x21, 0xf5, 0x09,
	0x74, 0xe7, 0xe3, 0xf5, 0xc2, 0xff, 0x72, 0x43, 0x28, 0x67, 0x99, 0x5e, 0x1f, 0x6c, 0x8d, 0x56,
	0xcb, 0x5f, 0x3c, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x21, 0xcd, 0x5d, 0x72, 0x98, 0x01, 0x00,
	0x00,
}
