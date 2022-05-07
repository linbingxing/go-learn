package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_learn/src/chatroom/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (ts *Transfer) ReadPkg() (msg message.Message, err error) {
	_, err = ts.Conn.Read(ts.Buf[:4])
	if err != nil {
		fmt.Println("conn read err =", err)
		return
	}

	var pkgLen uint32

	pkgLen = binary.BigEndian.Uint32(ts.Buf[:4])
	n, err := ts.Conn.Read(ts.Buf[:pkgLen])

	if n != int(pkgLen) || err != nil {
		fmt.Println("conn read fail err=", err)
		return
	}

	err = json.Unmarshal(ts.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json unmarshal err =", err)
		return
	}
	return
}

func (ts *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(ts.Buf[:4], pkgLen)
	n, err := ts.Conn.Write(ts.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println(" conn write len err ", err)
		return
	}
	n, err = ts.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write context err", err)
		return
	}
	return
}
