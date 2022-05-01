package tlv

import (
	"encoding/json"
	"io"
	"testing"
)

type MsgHelloXXX struct {
	Status string //'json:"status"'
}

var (
	MsgHelloTypeXXX = int8(1)
)

func initTestEvn() (*io.PipeReader, *io.PipeWriter, *Message) {
	m := &MsgHelloXXX{
		Status: "ok",
	}
	data, _ := json.Marshal(m)

	msg := &Message{
		Header: MsgHeader{
			MsgType: MsgHelloTypeXXX,
			Len:     int16(len(data)),
		},
		Data: data,
	}
	r, w := io.Pipe()
	return r, w, msg
}

func TestTlvOk(t *testing.T) {
	r, w, m := initTestEvn()
	go func() {
		for {
			rm, err := ReadMsg(r)
			if err != nil {
				return
			}
			data := &MsgHelloXXX{}
			err = json.Unmarshal(rm.Data, data)
			if err != nil {
				t.Fatal("read msg failed: ", err.Error())
			}

			if data.Status != "ok" {
				t.Fatal("read msg error: ", data.Status)
			}
		}
	}()
	err := SendMsg(w, m)
	if err != nil {
		t.Fatal("Send msg failed: ", err.Error())
	}

	errPrefix := []byte{'\xFE', '\xFF', '\xEE'}
	w.Write(errPrefix)
	err = SendMsg(w, m)
	if err != nil {
		t.Fatal("Send msg failed: ", err.Error())
	}
	w.Close()
}
