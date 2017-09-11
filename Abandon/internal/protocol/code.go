package protocol

//"bytes"
//"encoding/binary"
//"errors"

//Pack 对指定结构按照protocol进行序列化
func Pack(data interface{}, protocol int) []byte {
	var b []byte
	switch protocol {
	case PROTOCOL_TYPE_BINARY:
		b = packBinary(data)
	case PROTOCOL_TYPE_JSON:
		b = packJson(data)
	}
	return b
}

//CodecEncode 对消息体进行编码
func CodecEncode(b []byte, length int, encode int) {
	switch encode {
	case ENCODE_DEFAULT:
	case ENCODE_BIT_NOT:
		for i := 0; i < length; i++ {
			b[i] = ^b[i]
		}
	case ENCODE_BYTE_RVS:
		if length%2 != 0 {
			length = length - 1
		}
		for i := 0; i < length; i = i + 2 {
			b[i], b[i+1] = b[i+1], b[i]
		}
		/*
			case ENCODE_HT_RVS:
				b[0], b[length-1] = b[length-1], b[0]*/

	case ENCODE_LOOP_XOR:
		if length == 1 {
			return
		}
		for i := 0; i < length-1; i++ {
			b[i+1] ^= b[i]
		}
		b[0] ^= b[length-1]
	}
}

//CodecDecode 对消息体进行编码
func CodecDecode(b []byte, length int, encode int) {
	switch encode {
	case ENCODE_DEFAULT:
	case ENCODE_BIT_NOT:
		for i := 0; i < length; i++ {
			b[i] = ^b[i]
		}
	case ENCODE_BYTE_RVS:
		if length%2 != 0 {
			length = length - 1
		}
		for i := 0; i < length; i = i + 2 {
			b[i], b[i+1] = b[i+1], b[i]
		}
		/*
			case ENCODE_HT_RVS:
				b[0], b[length-1] = b[length-1], b[0]*/

	case ENCODE_LOOP_XOR:
		if length == 1 {
			return
		}
		b[0] ^= b[length-1]
		for i := length - 1; i > 0; i-- {
			b[i] ^= b[i-1]
		}
	}
}
