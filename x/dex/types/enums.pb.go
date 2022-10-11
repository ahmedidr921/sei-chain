// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/enums.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type PositionDirection int32

const (
	PositionDirection_LONG  PositionDirection = 0
	PositionDirection_SHORT PositionDirection = 1
)

var PositionDirection_name = map[int32]string{
	0: "LONG",
	1: "SHORT",
}

var PositionDirection_value = map[string]int32{
	"LONG":  0,
	"SHORT": 1,
}

func (x PositionDirection) String() string {
	return proto.EnumName(PositionDirection_name, int32(x))
}

func (PositionDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8c5bb23c6eb0b88, []int{0}
}

type PositionEffect int32

const (
	PositionEffect_OPEN  PositionEffect = 0
	PositionEffect_CLOSE PositionEffect = 1
)

var PositionEffect_name = map[int32]string{
	0: "OPEN",
	1: "CLOSE",
}

var PositionEffect_value = map[string]int32{
	"OPEN":  0,
	"CLOSE": 1,
}

func (x PositionEffect) String() string {
	return proto.EnumName(PositionEffect_name, int32(x))
}

func (PositionEffect) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8c5bb23c6eb0b88, []int{1}
}

type OrderType int32

const (
<<<<<<< Updated upstream
	OrderType_LIMIT       OrderType = 0
	OrderType_MARKET      OrderType = 1
	OrderType_LIQUIDATION OrderType = 2
	OrderType_FOKMARKET   OrderType = 3
=======
	OrderType_LIMIT            OrderType = 0
	OrderType_MARKET           OrderType = 1
	OrderType_LIQUIDATION      OrderType = 2
	OrderType_FOKMARKET        OrderType = 3
	OrderType_FOKMARKETBYVALUE OrderType = 4
	OrderType_STOPLOSS         OrderType = 5
	OrderType_STOPLIMIT        OrderType = 6
>>>>>>> Stashed changes
)

var OrderType_name = map[int32]string{
	0: "LIMIT",
	1: "MARKET",
	2: "LIQUIDATION",
	3: "FOKMARKET",
<<<<<<< Updated upstream
}

var OrderType_value = map[string]int32{
	"LIMIT":       0,
	"MARKET":      1,
	"LIQUIDATION": 2,
	"FOKMARKET":   3,
=======
	4: "FOKMARKETBYVALUE",
	5: "STOPLOSS",
	6: "STOPLIMIT",
}

var OrderType_value = map[string]int32{
	"LIMIT":            0,
	"MARKET":           1,
	"LIQUIDATION":      2,
	"FOKMARKET":        3,
	"FOKMARKETBYVALUE": 4,
	"STOPLOSS":         5,
	"STOPLIMIT":        6,
>>>>>>> Stashed changes
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}

func (OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8c5bb23c6eb0b88, []int{2}
}

type Unit int32

const (
	Unit_STANDARD Unit = 0
	Unit_MILLI    Unit = 1
	Unit_MICRO    Unit = 2
	Unit_NANO     Unit = 3
)

var Unit_name = map[int32]string{
	0: "STANDARD",
	1: "MILLI",
	2: "MICRO",
	3: "NANO",
}

var Unit_value = map[string]int32{
	"STANDARD": 0,
	"MILLI":    1,
	"MICRO":    2,
	"NANO":     3,
}

func (x Unit) String() string {
	return proto.EnumName(Unit_name, int32(x))
}

func (Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8c5bb23c6eb0b88, []int{3}
}

type OrderStatus int32

const (
	OrderStatus_PLACED          OrderStatus = 0
	OrderStatus_FAILED_TO_PLACE OrderStatus = 1
	OrderStatus_CANCELLED       OrderStatus = 2
	OrderStatus_FULFILLED       OrderStatus = 3
)

var OrderStatus_name = map[int32]string{
	0: "PLACED",
	1: "FAILED_TO_PLACE",
	2: "CANCELLED",
	3: "FULFILLED",
}

var OrderStatus_value = map[string]int32{
	"PLACED":          0,
	"FAILED_TO_PLACE": 1,
	"CANCELLED":       2,
	"FULFILLED":       3,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8c5bb23c6eb0b88, []int{4}
}

type CancellationInitiator int32

const (
	CancellationInitiator_USER       CancellationInitiator = 0
	CancellationInitiator_LIQUIDATED CancellationInitiator = 1
)

var CancellationInitiator_name = map[int32]string{
	0: "USER",
	1: "LIQUIDATED",
}

var CancellationInitiator_value = map[string]int32{
	"USER":       0,
	"LIQUIDATED": 1,
}

func (x CancellationInitiator) String() string {
	return proto.EnumName(CancellationInitiator_name, int32(x))
}

func (CancellationInitiator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8c5bb23c6eb0b88, []int{5}
}

func init() {
	proto.RegisterEnum("seiprotocol.seichain.dex.PositionDirection", PositionDirection_name, PositionDirection_value)
	proto.RegisterEnum("seiprotocol.seichain.dex.PositionEffect", PositionEffect_name, PositionEffect_value)
	proto.RegisterEnum("seiprotocol.seichain.dex.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("seiprotocol.seichain.dex.Unit", Unit_name, Unit_value)
	proto.RegisterEnum("seiprotocol.seichain.dex.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterEnum("seiprotocol.seichain.dex.CancellationInitiator", CancellationInitiator_name, CancellationInitiator_value)
}

func init() { proto.RegisterFile("dex/enums.proto", fileDescriptor_b8c5bb23c6eb0b88) }

var fileDescriptor_b8c5bb23c6eb0b88 = []byte{
<<<<<<< Updated upstream
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x91, 0xcf, 0x8e, 0xda, 0x30,
	0x10, 0xc6, 0xe3, 0x65, 0xbb, 0x5a, 0x66, 0xdb, 0xc5, 0x75, 0x55, 0xa9, 0xa7, 0xdc, 0x2a, 0x55,
	0x91, 0x36, 0x51, 0xd5, 0xbe, 0x80, 0x37, 0x31, 0x5b, 0x0b, 0x13, 0xd3, 0xfc, 0xb9, 0xf4, 0x82,
	0x42, 0x30, 0xc5, 0x12, 0x24, 0x28, 0x31, 0x12, 0xbc, 0x45, 0x1f, 0xab, 0x47, 0x8e, 0x3d, 0x56,
	0xf0, 0x22, 0x2b, 0x07, 0xb8, 0x7d, 0x33, 0xf3, 0x8d, 0x7e, 0x33, 0xfa, 0x60, 0x30, 0x57, 0xbb,
	0x40, 0x55, 0xdb, 0x75, 0xeb, 0x6f, 0x9a, 0xda, 0xd4, 0xe4, 0x53, 0xab, 0x74, 0xa7, 0xca, 0x7a,
	0xe5, 0xb7, 0x4a, 0x97, 0xcb, 0x42, 0x57, 0xfe, 0x5c, 0xed, 0xbc, 0x2f, 0xf0, 0x7e, 0x52, 0xb7,
	0xda, 0xe8, 0xba, 0x8a, 0x74, 0xa3, 0x4a, 0x2b, 0xc8, 0x3d, 0xdc, 0x0a, 0x19, 0xbf, 0x60, 0x87,
	0xf4, 0xe1, 0x4d, 0xfa, 0x43, 0x26, 0x19, 0x46, 0xde, 0x67, 0x78, 0xbc, 0x3a, 0xd9, 0x62, 0xa1,
	0x4a, 0x63, 0x6d, 0x72, 0xc2, 0xe2, 0xb3, 0x2d, 0x14, 0x32, 0x65, 0x18, 0x79, 0xcf, 0xd0, 0x97,
	0xcd, 0x5c, 0x35, 0xd9, 0x7e, 0xa3, 0x6c, 0x5f, 0xf0, 0x31, 0xcf, 0xb0, 0x43, 0x00, 0xee, 0xc6,
	0x34, 0x19, 0xb1, 0x0c, 0x23, 0x32, 0x80, 0x07, 0xc1, 0x7f, 0xe6, 0x3c, 0xa2, 0x19, 0x97, 0x31,
	0xbe, 0x21, 0xef, 0xa0, 0x3f, 0x94, 0xa3, 0xcb, 0xbc, 0xe7, 0x7d, 0x87, 0xdb, 0xbc, 0xd2, 0x86,
	0xbc, 0x85, 0xfb, 0x34, 0xa3, 0x71, 0x44, 0x93, 0xe8, 0x0c, 0x19, 0x73, 0x21, 0x38, 0x46, 0x67,
	0x19, 0x26, 0x12, 0xdf, 0xd8, 0x23, 0x62, 0x1a, 0x4b, 0xdc, 0xf3, 0x04, 0x3c, 0x74, 0xe4, 0xd4,
	0x14, 0x66, 0xdb, 0x5a, 0xe0, 0x44, 0xd0, 0x90, 0xd9, 0xd5, 0x0f, 0x30, 0x18, 0x52, 0x2e, 0x58,
	0x34, 0xcd, 0xe4, 0xb4, 0xeb, 0x62, 0x64, 0xa1, 0x21, 0x8d, 0x43, 0x26, 0x04, 0x8b, 0x2e, 0x37,
	0xe4, 0x62, 0xc8, 0xbb, 0xb2, 0xe7, 0x7d, 0x85, 0x8f, 0x61, 0x51, 0x95, 0x6a, 0xb5, 0x2a, 0xec,
	0xcb, 0xbc, 0xd2, 0x46, 0x17, 0xa6, 0x6e, 0x2c, 0x30, 0x4f, 0x59, 0x82, 0x1d, 0xf2, 0x08, 0x70,
	0x7d, 0x83, 0x45, 0x18, 0x3d, 0xbf, 0xfc, 0x3d, 0xba, 0xe8, 0x70, 0x74, 0xd1, 0xff, 0xa3, 0x8b,
	0xfe, 0x9c, 0x5c, 0xe7, 0x70, 0x72, 0x9d, 0x7f, 0x27, 0xd7, 0xf9, 0xf5, 0xf4, 0x5b, 0x9b, 0xe5,
	0x76, 0xe6, 0x97, 0xf5, 0x3a, 0x68, 0x95, 0x7e, 0xba, 0x66, 0xd1, 0x15, 0x5d, 0x18, 0xc1, 0x2e,
	0xb0, 0xa1, 0x99, 0xfd, 0x46, 0xb5, 0xb3, 0xbb, 0x6e, 0xfe, 0xed, 0x35, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0x50, 0x23, 0x73, 0xc8, 0x01, 0x00, 0x00,
=======
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x51, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x8d, 0xd7, 0xae, 0x5a, 0xef, 0x60, 0x35, 0x06, 0x24, 0x9e, 0xf2, 0x86, 0x84, 0x22, 0xad,
	0x15, 0x82, 0x3f, 0xe0, 0x25, 0xee, 0xb0, 0xe6, 0xc6, 0x25, 0x1f, 0x48, 0xf0, 0x32, 0x65, 0xa9,
	0xc7, 0x2c, 0x75, 0x49, 0x95, 0x38, 0x52, 0xf7, 0x2f, 0xf8, 0x59, 0x3c, 0xee, 0x91, 0x47, 0xd4,
	0xfe, 0x11, 0x64, 0xa7, 0xe5, 0xed, 0x9c, 0x7b, 0xcf, 0xf5, 0x3d, 0xbe, 0x07, 0x26, 0x2b, 0xb5,
	0x9d, 0xa9, 0xaa, 0x7b, 0x6c, 0xa7, 0x9b, 0xa6, 0x36, 0x35, 0x79, 0xd7, 0x2a, 0xed, 0x50, 0x59,
	0xaf, 0xa7, 0xad, 0xd2, 0xe5, 0x43, 0xa1, 0xab, 0xe9, 0x4a, 0x6d, 0x83, 0x0f, 0xf0, 0x6a, 0x59,
	0xb7, 0xda, 0xe8, 0xba, 0x8a, 0x74, 0xa3, 0x4a, 0x0b, 0xc8, 0x19, 0x0c, 0x85, 0x8c, 0xaf, 0xb1,
	0x47, 0xc6, 0x70, 0x9a, 0x7e, 0x91, 0x49, 0x86, 0x51, 0xf0, 0x1e, 0x2e, 0x8e, 0x4a, 0x76, 0x7f,
	0xaf, 0x4a, 0x63, 0x65, 0x72, 0xc9, 0xe2, 0x5e, 0x16, 0x0a, 0x99, 0x32, 0x8c, 0x82, 0x0e, 0xc6,
	0xb2, 0x59, 0xa9, 0x26, 0x7b, 0xda, 0x28, 0x5b, 0x17, 0x7c, 0xc1, 0x33, 0xec, 0x11, 0x80, 0xd1,
	0x82, 0x26, 0x37, 0x2c, 0xc3, 0x88, 0x4c, 0xe0, 0x5c, 0xf0, 0xaf, 0x39, 0x8f, 0x68, 0xc6, 0x65,
	0x8c, 0x4f, 0xc8, 0x4b, 0x18, 0xcf, 0xe5, 0xcd, 0xa1, 0x3f, 0x20, 0x6f, 0x00, 0xff, 0xa7, 0x57,
	0xdf, 0xbf, 0x51, 0x91, 0x33, 0x3c, 0x24, 0x2f, 0xe0, 0x2c, 0xcd, 0xe4, 0x52, 0xc8, 0x34, 0xc5,
	0xa7, 0x76, 0xc4, 0x31, 0xf7, 0xfc, 0x28, 0xf8, 0x0c, 0xc3, 0xbc, 0xd2, 0xa6, 0x17, 0xd1, 0x38,
	0xa2, 0x49, 0xd4, 0xfb, 0x5a, 0x70, 0x21, 0x38, 0x46, 0x3d, 0x0c, 0x13, 0x89, 0x4f, 0xac, 0xef,
	0x98, 0xc6, 0x12, 0x0f, 0x02, 0x01, 0xe7, 0xce, 0x6c, 0x6a, 0x0a, 0xd3, 0xb5, 0xd6, 0xe3, 0x52,
	0xd0, 0x90, 0xd9, 0xd1, 0xd7, 0x30, 0x99, 0x53, 0x2e, 0x58, 0x74, 0x9b, 0xc9, 0x5b, 0x57, 0xc5,
	0xc8, 0x2e, 0x0d, 0x69, 0x1c, 0x32, 0x21, 0x58, 0x74, 0xb0, 0x9d, 0x8b, 0x39, 0x77, 0x74, 0x10,
	0x7c, 0x84, 0xb7, 0x61, 0x51, 0x95, 0x6a, 0xbd, 0x2e, 0xec, 0x95, 0x78, 0xa5, 0x8d, 0x2e, 0x4c,
	0xdd, 0xd8, 0x85, 0x79, 0xca, 0x12, 0xec, 0x91, 0x0b, 0x80, 0xe3, 0xcf, 0x59, 0x84, 0xd1, 0xd5,
	0xf5, 0xef, 0x9d, 0x8f, 0x9e, 0x77, 0x3e, 0xfa, 0xbb, 0xf3, 0xd1, 0xaf, 0xbd, 0xef, 0x3d, 0xef,
	0x7d, 0xef, 0xcf, 0xde, 0xf7, 0x7e, 0x5c, 0xfe, 0xd4, 0xe6, 0xa1, 0xbb, 0x9b, 0x96, 0xf5, 0xe3,
	0xac, 0x55, 0xfa, 0xf2, 0x18, 0x9f, 0x23, 0x2e, 0xbf, 0xd9, 0x76, 0x66, 0x73, 0x36, 0x4f, 0x1b,
	0xd5, 0xde, 0x8d, 0x5c, 0xff, 0xd3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0xc8, 0xd2, 0x69,
	0xfb, 0x01, 0x00, 0x00,
>>>>>>> Stashed changes
}
