// Code generated by capnpc-go. DO NOT EDIT.

package updatetool

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

// Constants defined in update-tool.capnp.
var (
	UpdatePublicKeys = PublicSigningKey_List(capnp.MustUnmarshalRoot(x_96c3fff3f4beb8fe[0:56]).List())
)

type PublicSigningKey capnp.Struct

// PublicSigningKey_TypeID is the unique identifier for the type PublicSigningKey.
const PublicSigningKey_TypeID = 0x9b54bbec5de99f09

func NewPublicSigningKey(s *capnp.Segment) (PublicSigningKey, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	return PublicSigningKey(st), err
}

func NewRootPublicSigningKey(s *capnp.Segment) (PublicSigningKey, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	return PublicSigningKey(st), err
}

func ReadRootPublicSigningKey(msg *capnp.Message) (PublicSigningKey, error) {
	root, err := msg.Root()
	return PublicSigningKey(root.Struct()), err
}

func (s PublicSigningKey) String() string {
	str, _ := text.Marshal(0x9b54bbec5de99f09, capnp.Struct(s))
	return str
}

func (s PublicSigningKey) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (PublicSigningKey) DecodeFromPtr(p capnp.Ptr) PublicSigningKey {
	return PublicSigningKey(capnp.Struct{}.DecodeFromPtr(p))
}

func (s PublicSigningKey) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s PublicSigningKey) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s PublicSigningKey) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s PublicSigningKey) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s PublicSigningKey) Key0() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s PublicSigningKey) SetKey0(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s PublicSigningKey) Key1() uint64 {
	return capnp.Struct(s).Uint64(8)
}

func (s PublicSigningKey) SetKey1(v uint64) {
	capnp.Struct(s).SetUint64(8, v)
}

func (s PublicSigningKey) Key2() uint64 {
	return capnp.Struct(s).Uint64(16)
}

func (s PublicSigningKey) SetKey2(v uint64) {
	capnp.Struct(s).SetUint64(16, v)
}

func (s PublicSigningKey) Key3() uint64 {
	return capnp.Struct(s).Uint64(24)
}

func (s PublicSigningKey) SetKey3(v uint64) {
	capnp.Struct(s).SetUint64(24, v)
}

// PublicSigningKey_List is a list of PublicSigningKey.
type PublicSigningKey_List = capnp.StructList[PublicSigningKey]

// NewPublicSigningKey creates a new list of PublicSigningKey.
func NewPublicSigningKey_List(s *capnp.Segment, sz int32) (PublicSigningKey_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0}, sz)
	return capnp.StructList[PublicSigningKey](l), err
}

// PublicSigningKey_Future is a wrapper for a PublicSigningKey promised by a client call.
type PublicSigningKey_Future struct{ *capnp.Future }

func (f PublicSigningKey_Future) Struct() (PublicSigningKey, error) {
	p, err := f.Future.Ptr()
	return PublicSigningKey(p.Struct()), err
}

type Signature capnp.Struct

// Signature_TypeID is the unique identifier for the type Signature.
const Signature_TypeID = 0xc4d0558d33210637

func NewSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0})
	return Signature(st), err
}

func NewRootSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0})
	return Signature(st), err
}

func ReadRootSignature(msg *capnp.Message) (Signature, error) {
	root, err := msg.Root()
	return Signature(root.Struct()), err
}

func (s Signature) String() string {
	str, _ := text.Marshal(0xc4d0558d33210637, capnp.Struct(s))
	return str
}

func (s Signature) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Signature) DecodeFromPtr(p capnp.Ptr) Signature {
	return Signature(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Signature) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Signature) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Signature) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Signature) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Signature) Sig0() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Signature) SetSig0(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s Signature) Sig1() uint64 {
	return capnp.Struct(s).Uint64(8)
}

func (s Signature) SetSig1(v uint64) {
	capnp.Struct(s).SetUint64(8, v)
}

func (s Signature) Sig2() uint64 {
	return capnp.Struct(s).Uint64(16)
}

func (s Signature) SetSig2(v uint64) {
	capnp.Struct(s).SetUint64(16, v)
}

func (s Signature) Sig3() uint64 {
	return capnp.Struct(s).Uint64(24)
}

func (s Signature) SetSig3(v uint64) {
	capnp.Struct(s).SetUint64(24, v)
}

func (s Signature) Sig4() uint64 {
	return capnp.Struct(s).Uint64(32)
}

func (s Signature) SetSig4(v uint64) {
	capnp.Struct(s).SetUint64(32, v)
}

func (s Signature) Sig5() uint64 {
	return capnp.Struct(s).Uint64(40)
}

func (s Signature) SetSig5(v uint64) {
	capnp.Struct(s).SetUint64(40, v)
}

func (s Signature) Sig6() uint64 {
	return capnp.Struct(s).Uint64(48)
}

func (s Signature) SetSig6(v uint64) {
	capnp.Struct(s).SetUint64(48, v)
}

func (s Signature) Sig7() uint64 {
	return capnp.Struct(s).Uint64(56)
}

func (s Signature) SetSig7(v uint64) {
	capnp.Struct(s).SetUint64(56, v)
}

// Signature_List is a list of Signature.
type Signature_List = capnp.StructList[Signature]

// NewSignature creates a new list of Signature.
func NewSignature_List(s *capnp.Segment, sz int32) (Signature_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0}, sz)
	return capnp.StructList[Signature](l), err
}

// Signature_Future is a wrapper for a Signature promised by a client call.
type Signature_Future struct{ *capnp.Future }

func (f Signature_Future) Struct() (Signature, error) {
	p, err := f.Future.Ptr()
	return Signature(p.Struct()), err
}

type UpdateSignature capnp.Struct

// UpdateSignature_TypeID is the unique identifier for the type UpdateSignature.
const UpdateSignature_TypeID = 0x9c805f76ef46e6c0

func NewUpdateSignature(s *capnp.Segment) (UpdateSignature, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UpdateSignature(st), err
}

func NewRootUpdateSignature(s *capnp.Segment) (UpdateSignature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UpdateSignature(st), err
}

func ReadRootUpdateSignature(msg *capnp.Message) (UpdateSignature, error) {
	root, err := msg.Root()
	return UpdateSignature(root.Struct()), err
}

func (s UpdateSignature) String() string {
	str, _ := text.Marshal(0x9c805f76ef46e6c0, capnp.Struct(s))
	return str
}

func (s UpdateSignature) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (UpdateSignature) DecodeFromPtr(p capnp.Ptr) UpdateSignature {
	return UpdateSignature(capnp.Struct{}.DecodeFromPtr(p))
}

func (s UpdateSignature) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s UpdateSignature) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s UpdateSignature) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s UpdateSignature) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s UpdateSignature) Signatures() (Signature_List, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Signature_List(p.List()), err
}

func (s UpdateSignature) HasSignatures() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s UpdateSignature) SetSignatures(v Signature_List) error {
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewSignatures sets the signatures field to a newly
// allocated Signature_List, preferring placement in s's segment.
func (s UpdateSignature) NewSignatures(n int32) (Signature_List, error) {
	l, err := NewSignature_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Signature_List{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}

// UpdateSignature_List is a list of UpdateSignature.
type UpdateSignature_List = capnp.StructList[UpdateSignature]

// NewUpdateSignature creates a new list of UpdateSignature.
func NewUpdateSignature_List(s *capnp.Segment, sz int32) (UpdateSignature_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[UpdateSignature](l), err
}

// UpdateSignature_Future is a wrapper for a UpdateSignature promised by a client call.
type UpdateSignature_Future struct{ *capnp.Future }

func (f UpdateSignature_Future) Struct() (UpdateSignature, error) {
	p, err := f.Future.Ptr()
	return UpdateSignature(p.Struct()), err
}

const schema_96c3fff3f4beb8fe = "x\xda|\x93OH\x14Q\x1c\xc7\x7f\xdf\xf7{\xbb\xab" +
	"\x90\xe9\xdb\x99K\x9e\xec\x94\x04\x86:\x95\xd0\xa5?\x87" +
	".^\xf6a^\xa4\xc8M\xc7u\xc8tpv\x13\x09" +
	"\"\xe9\x9aA\x97(\xfaG\x97\xa0C\x87\xa0C\x91H" +
	"\x91\x10\x81\x95A\x81\x91\xd1&Z\x0aJQ\xde\x82&" +
	"\xden\xed\xbe :~>\xfc\xe6\xf7\xbe\xf3\x9d7\xad" +
	"\x03\xd8'\xdb\xea\xc6\x98\x84nN$\xe3\xda\x1b\xabG" +
	"\xd6\x1e\x1e\xbaB:\x0d\x19\xff\xbc?\xbd\xf1=~r" +
	"\x91d\x8a\xc8\xbb\x89F8w\x91\"r\xee\xe03!" +
	"~\xf4\xe9\xe0\x97\x93G\xcf\\%\x95Fu6a&" +
	"\xbc\x09\x91\x86sA\x98\xe1I\xb1\x97\x10w$\xb7z" +
	"\x93\xdds3fs\xcd\xdf\x9b\x9d\xc7b\xd1ynf" +
	"\xbdg\xe2)\x08\xf1\x8f\xf3\xbd\xcbSM\x0f\xbe\x91J" +
	"[)\x08\xde\xa4l\x84s\xcd<\xd4uI2\xba\xee" +
	"I\x01:\x10\x17\xc2\xfel\xdeo\xc9\xf3\xc8\xc8\xd0\x8e" +
	"\xbel8\x1c\xee\xc9\x14\x8e\x0d\x05}]An8\x18" +
	"\xceu\xb2?\x9e\x01t\x03K\"\x09\"\x95\xddN\xa4" +
	"\x0f3\xf4\xa0\x80\x02\\\x18\xe9\x1b\xd9\xcb\xd0C\x02J" +
	"\x08\x17\x82H\x05F\xf63t(\xa0\x98]0\x91:" +
	"a\xe4 C\xe7\x05\xea\x8f\xfb\xe3\xad\xa8%\x81Z*" +
	"A\x9b\x0d\xed6x\x7f\xe0_\x89\xbbK\xca$\xce\xe6" +
	"\x0b\xa3\xf0M`Y\x09\\\xd7C\xa471t\xb3@" +
	"\x1c\xfd\x1e\"\xf6#l&d\x18h\xa8VL0\xb2" +
	"r\x86\xa8\x9eQ\xde\x9e*\x8c\x96\xb67U\xb6\xbf2" +
	"\xef3\xcb\xd0\xf3V\x1do\x8c\x9cc\xe8\x05\xab\x8e\xb7" +
	"F\xbef\xe8\xa2U\xc7{#\xe7\x19zI@I\xe9" +
	"B\x12\xa9\x8fF.0\xf4\x8a\x80J$\\$\x88\xd4" +
	"\xb2\x91E\x86^\x13P\xc9\xa4\x8b$\x91Z5r\x89" +
	"\xa1\xbf\x0a\xa8T\xca57H\xad\x1b\xb9\xc2\xd0\x1b\x02" +
	"\xf5Q\x90\xabV\x1c\x05\xb96\x1b\xdam\xf0l\xd8i" +
	"\xc3.\x1bv\xdb\xd0\xf1\xbf\xcfRV\xe5\xeb\xd4\xe9\xf3" +
	"x\x94\x01\xaa\x9dW~\x98R\xe7\xa4\xb0MI\x19\x8f" +
	"\xdd>5z\xfdrK\x0f\xf7\xbd<\x17\x16o\xbd8" +
	";\xb5ezf}6\xfc0\x10\xed\x7f\xb7x\xbaf" +
	"\xe2W\x00\x00\x00\xff\xff\x8a\xd6\xdd\x10"

func init() {
	schemas.Register(schema_96c3fff3f4beb8fe,
		0x9b54bbec5de99f09,
		0x9c805f76ef46e6c0,
		0xc4d0558d33210637,
		0xf2b920bce5608efb)
}

var x_96c3fff3f4beb8fe = []byte{
	0, 0, 0, 0, 6, 0, 0, 0,
	1, 0, 0, 0, 39, 0, 0, 0,
	4, 0, 0, 0, 4, 0, 0, 0,
	119, 169, 123, 114, 158, 153, 45, 90,
	99, 207, 140, 112, 224, 166, 206, 131,
	188, 25, 190, 196, 237, 204, 112, 223,
	102, 115, 65, 219, 226, 126, 8, 129,
}
