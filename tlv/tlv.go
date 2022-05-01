package tlv

import (
	"encoding/binary"
	"io"
)

var (
	magicCode = [4]byte{'\xFE', '\xFF', '\xEF', '\xFF'}
	MAGL      = len(magicCode)
	MHL       = 3 // byte of MsgHeader
)

type MsgHeader struct {
	MsgType int8
	Len     int16
}

type Message struct {
	Header MsgHeader
	Data   []byte //msg body
}

func byte2MsgHeader(header []byte) *MsgHeader {
	if len(header) != MHL {
		return nil
	}
	mh := MsgHeader{}
	mh.MsgType = int8(header[0])
	mh.Len = int16(binary.BigEndian.Uint16(header[1:3]))
	return &mh
}

func Msg2byte(m *Message) []byte {
	msg := make([]byte, MAGL)
	copy(msg, magicCode[:])
	msg = append(msg, byte(m.Header.MsgType))
	bb := make([]byte, 2)
	binary.BigEndian.PutUint16(bb, uint16(m.Header.Len))
	msg = append(msg, bb...)
	msg = append(msg, m.Data...)
	return msg
}

func readMagicCode(c io.Reader) error {
	b := make([]byte, 1)

header:
	//find the first byte for magicCode
	_, err := io.ReadFull(c, b)
	if err != nil {
		return err
	}

	if b[0] != magicCode[0] {
		goto header
	}

	for i := 1; i < MAGL; i++ {
		_, err := io.ReadFull(c, b)
		if err != nil {
			return err
		}

		if b[0] != magicCode[i] {
			goto header
		}
	}

	return nil
}

//before user ReadMsgFormConn,we would like set the c read timeout
func ReadMsg(c io.Reader) (*Message, error) {
	err := readMagicCode(c)
	if err != nil {
		return nil, err
	}

	header := make([]byte, MHL)
	_, err = io.ReadFull(c, header)
	if err != nil {
		return nil, err
	}
	mh := byte2MsgHeader(header)
	if mh == nil {
		return nil, err
	}

	ms := Message{
		Header: *mh,
	}

	//read the body
	if mh.Len > 0 {
		body := make([]byte, mh.Len)
		_, err = io.ReadFull(c, body)
		if err != nil {
			return nil, err
		}

		ms.Data = body[:]
	}

	return &ms, nil
}

//before user ReadMsgFormConn,we would like set the c write timeout
func SendMsg(c io.Writer, m *Message) error {
	b := Msg2byte(m)
	_, err := c.Write(b)
	return err
}
