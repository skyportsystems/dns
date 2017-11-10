// +build fuzz

package dns

func Fuzz(data []byte) int {
	ret := 0
	msg := new(Msg)

	err := msg.Unpack(data)
	if err != nil {
		ret = 1
	}

	_, err = msg.Pack()
	if err != nil {
		ret = 1
	}

	return ret
}
