// +build fuzz

package dns

func Fuzz(data []byte) int {
	ret := 1
	msg := new(Msg)

	err := msg.Unpack(data)
	if err != nil {
		ret = 0
	}

	_, err = msg.Pack()
	if err != nil {
		ret = 0
	}

	return ret
}
