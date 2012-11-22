// Code generated by protoc-gen-go from "pond.proto"
// DO NOT EDIT!

package protos

import proto "code.google.com/p/goprotobuf/proto"
import "encoding/json"
import "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Reply_Status int32

const (
	Reply_OK                         Reply_Status = 0
	Reply_PARSE_ERROR                Reply_Status = 1
	Reply_NO_REQUEST                 Reply_Status = 2
	Reply_INTERNAL_ERROR             Reply_Status = 3
	Reply_IDENTITY_ALREADY_KNOWN     Reply_Status = 10
	Reply_OVERLOAD                   Reply_Status = 11
	Reply_NO_SUCH_ADDRESS            Reply_Status = 12
	Reply_DELIVERY_SIGNATURE_INVALID Reply_Status = 13
	Reply_INCORRECT_GENERATION       Reply_Status = 14
	Reply_MAILBOX_FULL               Reply_Status = 15
	Reply_NO_ACCOUNT                 Reply_Status = 16
)

var Reply_Status_name = map[int32]string{
	0:  "OK",
	1:  "PARSE_ERROR",
	2:  "NO_REQUEST",
	3:  "INTERNAL_ERROR",
	10: "IDENTITY_ALREADY_KNOWN",
	11: "OVERLOAD",
	12: "NO_SUCH_ADDRESS",
	13: "DELIVERY_SIGNATURE_INVALID",
	14: "INCORRECT_GENERATION",
	15: "MAILBOX_FULL",
	16: "NO_ACCOUNT",
}
var Reply_Status_value = map[string]int32{
	"OK":                         0,
	"PARSE_ERROR":                1,
	"NO_REQUEST":                 2,
	"INTERNAL_ERROR":             3,
	"IDENTITY_ALREADY_KNOWN":     10,
	"OVERLOAD":                   11,
	"NO_SUCH_ADDRESS":            12,
	"DELIVERY_SIGNATURE_INVALID": 13,
	"INCORRECT_GENERATION":       14,
	"MAILBOX_FULL":               15,
	"NO_ACCOUNT":                 16,
}

func (x Reply_Status) Enum() *Reply_Status {
	p := new(Reply_Status)
	*p = x
	return p
}
func (x Reply_Status) String() string {
	return proto.EnumName(Reply_Status_name, int32(x))
}
func (x Reply_Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Reply_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Reply_Status_value, data, "Reply_Status")
	if err != nil {
		return err
	}
	*x = Reply_Status(value)
	return nil
}

type Message_Encoding int32

const (
	Message_RAW  Message_Encoding = 0
	Message_GZIP Message_Encoding = 1
)

var Message_Encoding_name = map[int32]string{
	0: "RAW",
	1: "GZIP",
}
var Message_Encoding_value = map[string]int32{
	"RAW":  0,
	"GZIP": 1,
}

func (x Message_Encoding) Enum() *Message_Encoding {
	p := new(Message_Encoding)
	*p = x
	return p
}
func (x Message_Encoding) String() string {
	return proto.EnumName(Message_Encoding_name, int32(x))
}
func (x Message_Encoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Message_Encoding) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_Encoding_value, data, "Message_Encoding")
	if err != nil {
		return err
	}
	*x = Message_Encoding(value)
	return nil
}

type Request struct {
	NewAccount       *NewAccount `protobuf:"bytes,1,opt,name=new_account" json:"new_account,omitempty"`
	Deliver          *Delivery   `protobuf:"bytes,2,opt,name=deliver" json:"deliver,omitempty"`
	Fetch            *Fetch      `protobuf:"bytes,3,opt,name=fetch" json:"fetch,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (this *Request) Reset()         { *this = Request{} }
func (this *Request) String() string { return proto.CompactTextString(this) }
func (*Request) ProtoMessage()       {}

func (this *Request) GetNewAccount() *NewAccount {
	if this != nil {
		return this.NewAccount
	}
	return nil
}

func (this *Request) GetDeliver() *Delivery {
	if this != nil {
		return this.Deliver
	}
	return nil
}

func (this *Request) GetFetch() *Fetch {
	if this != nil {
		return this.Fetch
	}
	return nil
}

type Reply struct {
	Status           *Reply_Status   `protobuf:"varint,1,opt,name=status,enum=protos.Reply_Status,def=0" json:"status,omitempty"`
	AccountCreated   *AccountCreated `protobuf:"bytes,2,opt,name=account_created" json:"account_created,omitempty"`
	Fetched          *Fetched        `protobuf:"bytes,3,opt,name=fetched" json:"fetched,omitempty"`
	Announce         *ServerAnnounce `protobuf:"bytes,4,opt,name=announce" json:"announce,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *Reply) Reset()         { *this = Reply{} }
func (this *Reply) String() string { return proto.CompactTextString(this) }
func (*Reply) ProtoMessage()       {}

const Default_Reply_Status Reply_Status = Reply_OK

func (this *Reply) GetStatus() Reply_Status {
	if this != nil && this.Status != nil {
		return *this.Status
	}
	return Default_Reply_Status
}

func (this *Reply) GetAccountCreated() *AccountCreated {
	if this != nil {
		return this.AccountCreated
	}
	return nil
}

func (this *Reply) GetFetched() *Fetched {
	if this != nil {
		return this.Fetched
	}
	return nil
}

func (this *Reply) GetAnnounce() *ServerAnnounce {
	if this != nil {
		return this.Announce
	}
	return nil
}

type NewAccount struct {
	Generation       *uint32 `protobuf:"fixed32,1,req,name=generation" json:"generation,omitempty"`
	Group            []byte  `protobuf:"bytes,2,req,name=group" json:"group,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *NewAccount) Reset()         { *this = NewAccount{} }
func (this *NewAccount) String() string { return proto.CompactTextString(this) }
func (*NewAccount) ProtoMessage()       {}

func (this *NewAccount) GetGeneration() uint32 {
	if this != nil && this.Generation != nil {
		return *this.Generation
	}
	return 0
}

func (this *NewAccount) GetGroup() []byte {
	if this != nil {
		return this.Group
	}
	return nil
}

type AccountDetails struct {
	Queue            *uint32 `protobuf:"varint,1,req,name=queue" json:"queue,omitempty"`
	MaxQueue         *uint32 `protobuf:"varint,2,req,name=max_queue" json:"max_queue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *AccountDetails) Reset()         { *this = AccountDetails{} }
func (this *AccountDetails) String() string { return proto.CompactTextString(this) }
func (*AccountDetails) ProtoMessage()       {}

func (this *AccountDetails) GetQueue() uint32 {
	if this != nil && this.Queue != nil {
		return *this.Queue
	}
	return 0
}

func (this *AccountDetails) GetMaxQueue() uint32 {
	if this != nil && this.MaxQueue != nil {
		return *this.MaxQueue
	}
	return 0
}

type AccountCreated struct {
	Details          *AccountDetails `protobuf:"bytes,1,req,name=details" json:"details,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *AccountCreated) Reset()         { *this = AccountCreated{} }
func (this *AccountCreated) String() string { return proto.CompactTextString(this) }
func (*AccountCreated) ProtoMessage()       {}

func (this *AccountCreated) GetDetails() *AccountDetails {
	if this != nil {
		return this.Details
	}
	return nil
}

type Delivery struct {
	To               []byte  `protobuf:"bytes,1,req,name=to" json:"to,omitempty"`
	Signature        []byte  `protobuf:"bytes,2,req,name=signature" json:"signature,omitempty"`
	Generation       *uint32 `protobuf:"fixed32,3,req,name=generation" json:"generation,omitempty"`
	Message          []byte  `protobuf:"bytes,4,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Delivery) Reset()         { *this = Delivery{} }
func (this *Delivery) String() string { return proto.CompactTextString(this) }
func (*Delivery) ProtoMessage()       {}

func (this *Delivery) GetTo() []byte {
	if this != nil {
		return this.To
	}
	return nil
}

func (this *Delivery) GetSignature() []byte {
	if this != nil {
		return this.Signature
	}
	return nil
}

func (this *Delivery) GetGeneration() uint32 {
	if this != nil && this.Generation != nil {
		return *this.Generation
	}
	return 0
}

func (this *Delivery) GetMessage() []byte {
	if this != nil {
		return this.Message
	}
	return nil
}

type Fetch struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *Fetch) Reset()         { *this = Fetch{} }
func (this *Fetch) String() string { return proto.CompactTextString(this) }
func (*Fetch) ProtoMessage()       {}

type Fetched struct {
	Signature        []byte          `protobuf:"bytes,1,req,name=signature" json:"signature,omitempty"`
	Generation       *uint32         `protobuf:"fixed32,2,req,name=generation" json:"generation,omitempty"`
	Message          []byte          `protobuf:"bytes,3,req,name=message" json:"message,omitempty"`
	Details          *AccountDetails `protobuf:"bytes,4,req,name=details" json:"details,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *Fetched) Reset()         { *this = Fetched{} }
func (this *Fetched) String() string { return proto.CompactTextString(this) }
func (*Fetched) ProtoMessage()       {}

func (this *Fetched) GetSignature() []byte {
	if this != nil {
		return this.Signature
	}
	return nil
}

func (this *Fetched) GetGeneration() uint32 {
	if this != nil && this.Generation != nil {
		return *this.Generation
	}
	return 0
}

func (this *Fetched) GetMessage() []byte {
	if this != nil {
		return this.Message
	}
	return nil
}

func (this *Fetched) GetDetails() *AccountDetails {
	if this != nil {
		return this.Details
	}
	return nil
}

type ServerAnnounce struct {
	Time             *int64  `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	Msg              *string `protobuf:"bytes,2,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ServerAnnounce) Reset()         { *this = ServerAnnounce{} }
func (this *ServerAnnounce) String() string { return proto.CompactTextString(this) }
func (*ServerAnnounce) ProtoMessage()       {}

func (this *ServerAnnounce) GetTime() int64 {
	if this != nil && this.Time != nil {
		return *this.Time
	}
	return 0
}

func (this *ServerAnnounce) GetMsg() string {
	if this != nil && this.Msg != nil {
		return *this.Msg
	}
	return ""
}

type KeyExchange struct {
	PublicKey        []byte  `protobuf:"bytes,1,req,name=public_key" json:"public_key,omitempty"`
	IdentityPublic   []byte  `protobuf:"bytes,2,req,name=identity_public" json:"identity_public,omitempty"`
	Server           *string `protobuf:"bytes,3,req,name=server" json:"server,omitempty"`
	Dh               []byte  `protobuf:"bytes,4,req,name=dh" json:"dh,omitempty"`
	Group            []byte  `protobuf:"bytes,5,req,name=group" json:"group,omitempty"`
	GroupKey         []byte  `protobuf:"bytes,6,req,name=group_key" json:"group_key,omitempty"`
	Generation       *uint32 `protobuf:"varint,7,req,name=generation" json:"generation,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *KeyExchange) Reset()         { *this = KeyExchange{} }
func (this *KeyExchange) String() string { return proto.CompactTextString(this) }
func (*KeyExchange) ProtoMessage()       {}

func (this *KeyExchange) GetPublicKey() []byte {
	if this != nil {
		return this.PublicKey
	}
	return nil
}

func (this *KeyExchange) GetIdentityPublic() []byte {
	if this != nil {
		return this.IdentityPublic
	}
	return nil
}

func (this *KeyExchange) GetServer() string {
	if this != nil && this.Server != nil {
		return *this.Server
	}
	return ""
}

func (this *KeyExchange) GetDh() []byte {
	if this != nil {
		return this.Dh
	}
	return nil
}

func (this *KeyExchange) GetGroup() []byte {
	if this != nil {
		return this.Group
	}
	return nil
}

func (this *KeyExchange) GetGroupKey() []byte {
	if this != nil {
		return this.GroupKey
	}
	return nil
}

func (this *KeyExchange) GetGeneration() uint32 {
	if this != nil && this.Generation != nil {
		return *this.Generation
	}
	return 0
}

type SignedKeyExchange struct {
	Signed           []byte `protobuf:"bytes,1,req,name=signed" json:"signed,omitempty"`
	Signature        []byte `protobuf:"bytes,2,req,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *SignedKeyExchange) Reset()         { *this = SignedKeyExchange{} }
func (this *SignedKeyExchange) String() string { return proto.CompactTextString(this) }
func (*SignedKeyExchange) ProtoMessage()       {}

func (this *SignedKeyExchange) GetSigned() []byte {
	if this != nil {
		return this.Signed
	}
	return nil
}

func (this *SignedKeyExchange) GetSignature() []byte {
	if this != nil {
		return this.Signature
	}
	return nil
}

type Message struct {
	Id               *uint64           `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	Time             *int64            `protobuf:"varint,2,req,name=time" json:"time,omitempty"`
	Body             []byte            `protobuf:"bytes,3,req,name=body" json:"body,omitempty"`
	BodyEncoding     *Message_Encoding `protobuf:"varint,4,opt,name=body_encoding,enum=protos.Message_Encoding" json:"body_encoding,omitempty"`
	MyNextDh         []byte            `protobuf:"bytes,5,req,name=my_next_dh" json:"my_next_dh,omitempty"`
	InReplyTo        *uint64           `protobuf:"varint,6,opt,name=in_reply_to" json:"in_reply_to,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (this *Message) Reset()         { *this = Message{} }
func (this *Message) String() string { return proto.CompactTextString(this) }
func (*Message) ProtoMessage()       {}

func (this *Message) GetId() uint64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Message) GetTime() int64 {
	if this != nil && this.Time != nil {
		return *this.Time
	}
	return 0
}

func (this *Message) GetBody() []byte {
	if this != nil {
		return this.Body
	}
	return nil
}

func (this *Message) GetBodyEncoding() Message_Encoding {
	if this != nil && this.BodyEncoding != nil {
		return *this.BodyEncoding
	}
	return 0
}

func (this *Message) GetMyNextDh() []byte {
	if this != nil {
		return this.MyNextDh
	}
	return nil
}

func (this *Message) GetInReplyTo() uint64 {
	if this != nil && this.InReplyTo != nil {
		return *this.InReplyTo
	}
	return 0
}

func init() {
	proto.RegisterEnum("protos.Reply_Status", Reply_Status_name, Reply_Status_value)
	proto.RegisterEnum("protos.Message_Encoding", Message_Encoding_name, Message_Encoding_value)
}
