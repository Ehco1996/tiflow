// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tmp.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Record struct {
	StartTs []byte `protobuf:"bytes,1,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Tid     int32  `protobuf:"varint,3,opt,name=tid,proto3" json:"tid,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_096f24f1bc44e34c, []int{0}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetStartTs() []byte {
	if m != nil {
		return m.StartTs
	}
	return nil
}

func (m *Record) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Record) GetTid() int32 {
	if m != nil {
		return m.Tid
	}
	return 0
}

type Request struct {
	MaxTid int32 `protobuf:"varint,1,opt,name=max_tid,json=maxTid,proto3" json:"max_tid,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_096f24f1bc44e34c, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMaxTid() int32 {
	if m != nil {
		return m.MaxTid
	}
	return 0
}

func init() {
	proto.RegisterType((*Record)(nil), "pb.Record")
	proto.RegisterType((*Request)(nil), "pb.Request")
}

func init() { proto.RegisterFile("tmp.proto", fileDescriptor_096f24f1bc44e34c) }

var fileDescriptor_096f24f1bc44e34c = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xc9, 0x2d, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe7, 0x62, 0x0b, 0x4a, 0x4d,
	0xce, 0x2f, 0x4a, 0x11, 0x92, 0xe4, 0xe2, 0x28, 0x2e, 0x49, 0x2c, 0x2a, 0x89, 0x2f, 0x29, 0x96,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x62, 0x07, 0xf3, 0x43, 0x8a, 0x85, 0x24, 0xb8, 0xd8, 0x0b,
	0x12, 0x2b, 0x73, 0xf2, 0x13, 0x53, 0x24, 0x98, 0x20, 0x32, 0x50, 0xae, 0x90, 0x00, 0x17, 0x73,
	0x49, 0x66, 0x8a, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x88, 0xa9, 0xa4, 0xc4, 0xc5, 0x1e,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xce, 0xc5, 0x9e, 0x9b, 0x58, 0x11, 0x0f, 0x52,
	0xc0, 0x08, 0x56, 0xc0, 0x96, 0x9b, 0x58, 0x11, 0x92, 0x99, 0x62, 0x64, 0xc2, 0xc5, 0x15, 0x92,
	0x5b, 0x10, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xa4, 0xc6, 0xc5, 0xe9, 0x5a, 0x96, 0x9a,
	0x57, 0xe2, 0x96, 0x9a, 0x9a, 0x22, 0xc4, 0xad, 0x57, 0x90, 0xa4, 0x07, 0x35, 0x40, 0x8a, 0x0b,
	0xc2, 0x01, 0x39, 0xcf, 0x80, 0xd1, 0x49, 0xe2, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0x92, 0xd8, 0xc0, 0xfe, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x2a, 0xbd, 0xb7,
	0xdc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TmpServiceClient is the client API for TmpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TmpServiceClient interface {
	EventFeed(ctx context.Context, in *Request, opts ...grpc.CallOption) (TmpService_EventFeedClient, error)
}

type tmpServiceClient struct {
	cc *grpc.ClientConn
}

func NewTmpServiceClient(cc *grpc.ClientConn) TmpServiceClient {
	return &tmpServiceClient{cc}
}

func (c *tmpServiceClient) EventFeed(ctx context.Context, in *Request, opts ...grpc.CallOption) (TmpService_EventFeedClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TmpService_serviceDesc.Streams[0], "/pb.TmpService/EventFeed", opts...)
	if err != nil {
		return nil, err
	}
	x := &tmpServiceEventFeedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TmpService_EventFeedClient interface {
	Recv() (*Record, error)
	grpc.ClientStream
}

type tmpServiceEventFeedClient struct {
	grpc.ClientStream
}

func (x *tmpServiceEventFeedClient) Recv() (*Record, error) {
	m := new(Record)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TmpServiceServer is the server API for TmpService service.
type TmpServiceServer interface {
	EventFeed(*Request, TmpService_EventFeedServer) error
}

// UnimplementedTmpServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTmpServiceServer struct {
}

func (*UnimplementedTmpServiceServer) EventFeed(req *Request, srv TmpService_EventFeedServer) error {
	return status.Errorf(codes.Unimplemented, "method EventFeed not implemented")
}

func RegisterTmpServiceServer(s *grpc.Server, srv TmpServiceServer) {
	s.RegisterService(&_TmpService_serviceDesc, srv)
}

func _TmpService_EventFeed_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TmpServiceServer).EventFeed(m, &tmpServiceEventFeedServer{stream})
}

type TmpService_EventFeedServer interface {
	Send(*Record) error
	grpc.ServerStream
}

type tmpServiceEventFeedServer struct {
	grpc.ServerStream
}

func (x *tmpServiceEventFeedServer) Send(m *Record) error {
	return x.ServerStream.SendMsg(m)
}

var _TmpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TmpService",
	HandlerType: (*TmpServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventFeed",
			Handler:       _TmpService_EventFeed_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tmp.proto",
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Tid != 0 {
		i = encodeVarintTmp(dAtA, i, uint64(m.Tid))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTmp(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StartTs) > 0 {
		i -= len(m.StartTs)
		copy(dAtA[i:], m.StartTs)
		i = encodeVarintTmp(dAtA, i, uint64(len(m.StartTs)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxTid != 0 {
		i = encodeVarintTmp(dAtA, i, uint64(m.MaxTid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTmp(dAtA []byte, offset int, v uint64) int {
	offset -= sovTmp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StartTs)
	if l > 0 {
		n += 1 + l + sovTmp(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTmp(uint64(l))
	}
	if m.Tid != 0 {
		n += 1 + sovTmp(uint64(m.Tid))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxTid != 0 {
		n += 1 + sovTmp(uint64(m.MaxTid))
	}
	return n
}

func sovTmp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTmp(x uint64) (n int) {
	return sovTmp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTmp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTmp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTmp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartTs = append(m.StartTs[:0], dAtA[iNdEx:postIndex]...)
			if m.StartTs == nil {
				m.StartTs = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTmp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTmp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tid", wireType)
			}
			m.Tid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTmp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTmp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTmp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTid", wireType)
			}
			m.MaxTid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTmp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTmp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTmp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTmp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTmp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTmp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTmp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTmp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTmp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTmp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTmp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTmp = fmt.Errorf("proto: unexpected end of group")
)
