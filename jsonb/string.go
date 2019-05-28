package jsonb

var escapeCharMap = map[byte][]byte{
	9:  []byte{92, 116},
	10: []byte{92, 110},
	34: []byte{92, 34},
	92: []byte{92, 92},
}

func escapeString(w Writer, str string) {
	length := len(str)
	for i := 0; i < length; i++ {
		b := str[0]
		str = str[1:]
		if x, isOk := escapeCharMap[b]; isOk {
			w.Write(x)
			continue
		}
		w.WriteByte(b)
	}
}

func unescapeString(w Writer, b string) {
	length := len(b)
	for i := 0; i < length; {
		c := b[i]

		if c == '\\' {
			switch b[i+1] {
			case '"':
				w.WriteRune('"')
				i += 2
			case '\\':
				w.WriteRune('\\')
				i += 2
			case 'b':
				w.WriteRune('\b')
				i += 2
			case 'f':
				w.WriteRune('\f')
				i += 2
			case 'n':
				w.WriteRune('\n')
				i += 2
			case 'r':
				w.WriteRune('\r')
				i += 2
			case 't':
				w.WriteRune('\t')
				i += 2
			case '/':
				w.WriteRune('/')
				i += 2
			case 'u':
				i += 2
			}
		} else {
			w.WriteByte(c)
			i++
		}
	}
}
