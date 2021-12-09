package dlms

import (
	"encoding/binary"
	"fmt"

	"github.com/Circutor/gosem/pkg/axdr"
)

type AttributeDescriptor struct {
	ClassID     uint16
	InstanceID  Obis
	AttributeID int8
}

func CreateAttributeDescriptor(c uint16, i string, a int8) *AttributeDescriptor {
	var ob Obis = *CreateObis(i)

	return &AttributeDescriptor{ClassID: c, InstanceID: ob, AttributeID: a}
}

func (ad AttributeDescriptor) Encode() (out []byte, err error) {
	var output []byte
	var c [2]byte
	binary.BigEndian.PutUint16(c[:], ad.ClassID)
	output = append(output, c[:]...)
	output = append(output, ad.InstanceID.Bytes()...)
	output = append(output, byte(ad.AttributeID))

	out = output
	return
}

func DecodeAttributeDescriptor(ori *[]byte) (out AttributeDescriptor, err error) {
	src := append([]byte(nil), (*ori)...)

	if len(src) < 9 {
		err = fmt.Errorf("byte slice length must be at least 9 bytes")
		return
	}

	_, out.ClassID, err = axdr.DecodeLongUnsigned(&src)
	if err != nil {
		return
	}
	out.InstanceID, err = DecodeObis(&src)
	if err != nil {
		return
	}
	out.AttributeID = int8(src[0])
	src = src[1:]

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
