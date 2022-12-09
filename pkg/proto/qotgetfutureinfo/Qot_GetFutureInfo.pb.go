// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Qot_GetFutureInfo.proto

package qotgetfutureinfo

import (
	_ "futuq/pkg/proto/common"
	qotcommon "futuq/pkg/proto/qotcommon"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//交易时间
type TradeTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Begin *float64 `protobuf:"fixed64,1,opt,name=begin" json:"begin,omitempty"` // 开始时间,以分钟为单位
	End   *float64 `protobuf:"fixed64,2,opt,name=end" json:"end,omitempty"`     // 结束时间,以分钟为单位
}

func (x *TradeTime) Reset() {
	*x = TradeTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetFutureInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeTime) ProtoMessage() {}

func (x *TradeTime) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetFutureInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeTime.ProtoReflect.Descriptor instead.
func (*TradeTime) Descriptor() ([]byte, []int) {
	return file_Qot_GetFutureInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TradeTime) GetBegin() float64 {
	if x != nil && x.Begin != nil {
		return *x.Begin
	}
	return 0
}

func (x *TradeTime) GetEnd() float64 {
	if x != nil && x.End != nil {
		return *x.End
	}
	return 0
}

//期货合约资料的列表
type FutureInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               *string             `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`                               // 合约名称
	Security           *qotcommon.Security `protobuf:"bytes,2,req,name=security" json:"security,omitempty"`                       // 合约代码
	LastTradeTime      *string             `protobuf:"bytes,3,req,name=lastTradeTime" json:"lastTradeTime,omitempty"`             //最后交易日，只有非主连期货合约才有该字段
	LastTradeTimestamp *float64            `protobuf:"fixed64,4,opt,name=lastTradeTimestamp" json:"lastTradeTimestamp,omitempty"` //最后交易日时间戳，只有非主连期货合约才有该字段
	Owner              *qotcommon.Security `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`                             //标的股 股票期货和股指期货才有该字段
	OwnerOther         *string             `protobuf:"bytes,6,req,name=ownerOther" json:"ownerOther,omitempty"`                   //标的
	Exchange           *string             `protobuf:"bytes,7,req,name=exchange" json:"exchange,omitempty"`                       //交易所
	ContractType       *string             `protobuf:"bytes,8,req,name=contractType" json:"contractType,omitempty"`               //合约类型
	ContractSize       *float64            `protobuf:"fixed64,9,req,name=contractSize" json:"contractSize,omitempty"`             //合约规模
	ContractSizeUnit   *string             `protobuf:"bytes,10,req,name=contractSizeUnit" json:"contractSizeUnit,omitempty"`      //合约规模的单位
	QuoteCurrency      *string             `protobuf:"bytes,11,req,name=quoteCurrency" json:"quoteCurrency,omitempty"`            //报价货币
	MinVar             *float64            `protobuf:"fixed64,12,req,name=minVar" json:"minVar,omitempty"`                        //最小变动单位
	MinVarUnit         *string             `protobuf:"bytes,13,req,name=minVarUnit" json:"minVarUnit,omitempty"`                  //最小变动单位的单位
	QuoteUnit          *string             `protobuf:"bytes,14,opt,name=quoteUnit" json:"quoteUnit,omitempty"`                    //报价单位
	TradeTime          []*TradeTime        `protobuf:"bytes,15,rep,name=tradeTime" json:"tradeTime,omitempty"`                    //交易时间
	TimeZone           *string             `protobuf:"bytes,16,req,name=timeZone" json:"timeZone,omitempty"`                      //所在时区
	ExchangeFormatUrl  *string             `protobuf:"bytes,17,req,name=exchangeFormatUrl" json:"exchangeFormatUrl,omitempty"`    //交易所规格
}

func (x *FutureInfo) Reset() {
	*x = FutureInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetFutureInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FutureInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FutureInfo) ProtoMessage() {}

func (x *FutureInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetFutureInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FutureInfo.ProtoReflect.Descriptor instead.
func (*FutureInfo) Descriptor() ([]byte, []int) {
	return file_Qot_GetFutureInfo_proto_rawDescGZIP(), []int{1}
}

func (x *FutureInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *FutureInfo) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *FutureInfo) GetLastTradeTime() string {
	if x != nil && x.LastTradeTime != nil {
		return *x.LastTradeTime
	}
	return ""
}

func (x *FutureInfo) GetLastTradeTimestamp() float64 {
	if x != nil && x.LastTradeTimestamp != nil {
		return *x.LastTradeTimestamp
	}
	return 0
}

func (x *FutureInfo) GetOwner() *qotcommon.Security {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *FutureInfo) GetOwnerOther() string {
	if x != nil && x.OwnerOther != nil {
		return *x.OwnerOther
	}
	return ""
}

func (x *FutureInfo) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *FutureInfo) GetContractType() string {
	if x != nil && x.ContractType != nil {
		return *x.ContractType
	}
	return ""
}

func (x *FutureInfo) GetContractSize() float64 {
	if x != nil && x.ContractSize != nil {
		return *x.ContractSize
	}
	return 0
}

func (x *FutureInfo) GetContractSizeUnit() string {
	if x != nil && x.ContractSizeUnit != nil {
		return *x.ContractSizeUnit
	}
	return ""
}

func (x *FutureInfo) GetQuoteCurrency() string {
	if x != nil && x.QuoteCurrency != nil {
		return *x.QuoteCurrency
	}
	return ""
}

func (x *FutureInfo) GetMinVar() float64 {
	if x != nil && x.MinVar != nil {
		return *x.MinVar
	}
	return 0
}

func (x *FutureInfo) GetMinVarUnit() string {
	if x != nil && x.MinVarUnit != nil {
		return *x.MinVarUnit
	}
	return ""
}

func (x *FutureInfo) GetQuoteUnit() string {
	if x != nil && x.QuoteUnit != nil {
		return *x.QuoteUnit
	}
	return ""
}

func (x *FutureInfo) GetTradeTime() []*TradeTime {
	if x != nil {
		return x.TradeTime
	}
	return nil
}

func (x *FutureInfo) GetTimeZone() string {
	if x != nil && x.TimeZone != nil {
		return *x.TimeZone
	}
	return ""
}

func (x *FutureInfo) GetExchangeFormatUrl() string {
	if x != nil && x.ExchangeFormatUrl != nil {
		return *x.ExchangeFormatUrl
	}
	return ""
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityList []*qotcommon.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"` //股票列表
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetFutureInfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetFutureInfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Qot_GetFutureInfo_proto_rawDescGZIP(), []int{2}
}

func (x *C2S) GetSecurityList() []*qotcommon.Security {
	if x != nil {
		return x.SecurityList
	}
	return nil
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FutureInfoList []*FutureInfo `protobuf:"bytes,1,rep,name=futureInfoList" json:"futureInfoList,omitempty"` //期货合约资料的列表
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetFutureInfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetFutureInfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Qot_GetFutureInfo_proto_rawDescGZIP(), []int{3}
}

func (x *S2C) GetFutureInfoList() []*FutureInfo {
	if x != nil {
		return x.FutureInfoList
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetFutureInfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetFutureInfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Qot_GetFutureInfo_proto_rawDescGZIP(), []int{4}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetFutureInfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetFutureInfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Qot_GetFutureInfo_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Qot_GetFutureInfo_proto protoreflect.FileDescriptor

var file_Qot_GetFutureInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x51, 0x6f, 0x74, 0x5f, 0x47,
	0x65, 0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x09,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x22, 0x86, 0x05, 0x0a, 0x0a, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0d, 0x6c,
	0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f,
	0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x06, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x07, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x02, 0x28, 0x01, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x55, 0x6e, 0x69, 0x74,
	0x18, 0x0a, 0x20, 0x02, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x72, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x01, 0x52, 0x06,
	0x6d, 0x69, 0x6e, 0x56, 0x61, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x72,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x56,
	0x61, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x10, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x11,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x55, 0x72,
	0x6c, 0x18, 0x11, 0x20, 0x02, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x3f, 0x0a, 0x03, 0x43, 0x32,
	0x53, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x03, 0x53,
	0x32, 0x43, 0x12, 0x45, 0x0a, 0x0e, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x47, 0x65, 0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46,
	0x75, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x66, 0x75, 0x74, 0x75, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x46, 0x75, 0x74, 0x75, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x86,
	0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34,
	0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74,
	0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a,
	0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x47, 0x65, 0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53,
	0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x42, 0x4b, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66,
	0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72, 0x69, 0x73,
	0x68, 0x65, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x62, 0x2f, 0x71, 0x6f, 0x74, 0x67, 0x65, 0x74, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65,
	0x69, 0x6e, 0x66, 0x6f,
}

var (
	file_Qot_GetFutureInfo_proto_rawDescOnce sync.Once
	file_Qot_GetFutureInfo_proto_rawDescData = file_Qot_GetFutureInfo_proto_rawDesc
)

func file_Qot_GetFutureInfo_proto_rawDescGZIP() []byte {
	file_Qot_GetFutureInfo_proto_rawDescOnce.Do(func() {
		file_Qot_GetFutureInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetFutureInfo_proto_rawDescData)
	})
	return file_Qot_GetFutureInfo_proto_rawDescData
}

var file_Qot_GetFutureInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Qot_GetFutureInfo_proto_goTypes = []interface{}{
	(*TradeTime)(nil),          // 0: Qot_GetFutureInfo.TradeTime
	(*FutureInfo)(nil),         // 1: Qot_GetFutureInfo.FutureInfo
	(*C2S)(nil),                // 2: Qot_GetFutureInfo.C2S
	(*S2C)(nil),                // 3: Qot_GetFutureInfo.S2C
	(*Request)(nil),            // 4: Qot_GetFutureInfo.Request
	(*Response)(nil),           // 5: Qot_GetFutureInfo.Response
	(*qotcommon.Security)(nil), // 6: Qot_Common.Security
}
var file_Qot_GetFutureInfo_proto_depIdxs = []int32{
	6, // 0: Qot_GetFutureInfo.FutureInfo.security:type_name -> Qot_Common.Security
	6, // 1: Qot_GetFutureInfo.FutureInfo.owner:type_name -> Qot_Common.Security
	0, // 2: Qot_GetFutureInfo.FutureInfo.tradeTime:type_name -> Qot_GetFutureInfo.TradeTime
	6, // 3: Qot_GetFutureInfo.C2S.securityList:type_name -> Qot_Common.Security
	1, // 4: Qot_GetFutureInfo.S2C.futureInfoList:type_name -> Qot_GetFutureInfo.FutureInfo
	2, // 5: Qot_GetFutureInfo.Request.c2s:type_name -> Qot_GetFutureInfo.C2S
	3, // 6: Qot_GetFutureInfo.Response.s2c:type_name -> Qot_GetFutureInfo.S2C
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_Qot_GetFutureInfo_proto_init() }
func file_Qot_GetFutureInfo_proto_init() {
	if File_Qot_GetFutureInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetFutureInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeTime); i {
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
		file_Qot_GetFutureInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FutureInfo); i {
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
		file_Qot_GetFutureInfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_Qot_GetFutureInfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_Qot_GetFutureInfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Qot_GetFutureInfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_Qot_GetFutureInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetFutureInfo_proto_goTypes,
		DependencyIndexes: file_Qot_GetFutureInfo_proto_depIdxs,
		MessageInfos:      file_Qot_GetFutureInfo_proto_msgTypes,
	}.Build()
	File_Qot_GetFutureInfo_proto = out.File
	file_Qot_GetFutureInfo_proto_rawDesc = nil
	file_Qot_GetFutureInfo_proto_goTypes = nil
	file_Qot_GetFutureInfo_proto_depIdxs = nil
}