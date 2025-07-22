package proto

type RespMessage struct {
	Typ    string
	Status string
	Error  string
	Num    int
	String string
	Array  []RespMessage
}

func (v *RespMessage) Marshal() []byte {
	switch v.Typ {
	case "status":
		return v.marshalStatus()
	case "integer":
		return v.marshalInteger()
	case "null":
		return v.marshalBulkNull()
	case "nil":
		return v.marshalNil()
	case "error":
		return v.marshalError()
	default:
		return []byte{}
	}
}

// status marshal format: +<status>\r\n
func (v *RespMessage) marshalStatus() []byte {
	var b []byte

	b = append(b, RespStatus)
	b = append(b, []byte(v.Status)...)
	b = append(b, cr, lf)

	return b
}

// integer marshal format: :[<+|->]<value>\r\n
func (v *RespMessage) marshalInteger() []byte {
	var b []byte
	// TODO:
	// it supposed to return with format bellow, [ haven't figure a good way to indicate (+|-) ] maybe made a separate handler.
	// :[<+|->]<value>\r\n
	b = append(b, RespInt)
	b = append(b, byte(v.Num))
	b = append(b, cr, lf)

	return b
}

// bulk null format: $-1\r\n
func (v *RespMessage) marshalBulkNull() []byte {
	return []byte{RespString, RespError, '1', cr, lf}
}

// format: _\r\n
func (v *RespMessage) marshalNil() []byte {
	return []byte{RespNil, cr, lf}
}

// format: -<error message>\r\n
func (v *RespMessage) marshalError() []byte {
	var b []byte

	b = append(b, RespError)
	b = append(b, []byte(v.Error)...)
	b = append(b, cr, lf)

	return b
}
