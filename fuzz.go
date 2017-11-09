// +build fuzz

package dns

import (
	"bytes"
	"encoding/hex"
	"os"
)

func Fuzz(data []byte) int {
	var (
		msg         = &Msg{}
		buf, bufOne = make([]byte, 100000), make([]byte, 100000)
		res, resOne []byte

		unpackErr, packErr error
	)

	if unpackErr = msg.Unpack(data); unpackErr != nil {
		return 0
	}

	if res, packErr = msg.PackBuffer(buf); packErr != nil {
		return 0
	}

	for i := range res {
		bufOne[i] = 1
	}

	resOne, packErr = msg.PackBuffer(bufOne)
	if packErr != nil {
		println("pack failed only with a filled buffer")
		panic(packErr)
	}

	if !bytes.Equal(res, resOne) {
		println("buffer bits leaked into the packed message")
		println(hex.Dump(res))
		println(hex.Dump(resOne))
		os.Exit(1)
	}

	return 1
}
