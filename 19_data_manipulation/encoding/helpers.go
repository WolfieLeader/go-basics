package encoding

import "bytes"

const (
	HELLO_WORLD = "ĤèlĬϴ ₩órł⫒🌍"
	JAPANESE    = "日本語🏯"
)

func areAllEqual(data ...[]byte) bool {
	if len(data) < 2 {
		return true
	}

	for i := 1; i < len(data); i++ {
		if !bytes.Equal(data[0], data[i]) {
			return false
		}
	}
	return true
}
