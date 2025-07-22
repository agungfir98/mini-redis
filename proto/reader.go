package proto

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

const (
	RespStatus = '+' // +<string>\r\n
	RespError  = '-' // -<string>\r\n
	RespString = '$' // $<length>\r\n<bytes>\r\n
	RespInt    = ':' // :<number>\r\n
	RespNil    = '_' // _\r\n
	RespArray  = '*' // *<len>\r\n... (same as resp2)

	cr = '\r'
	lf = '\n'
)

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	r := &Resp{reader: bufio.NewReader(rd)}

	return r
}

func (r *Resp) Read() (RespMessage, error) {
	_type, err := r.reader.ReadByte()
	if err != nil {
		return RespMessage{}, err
	}

	switch _type {
	case RespArray:
		return r.readArray()
	case RespString:
		return r.readString()
	default:
		return RespMessage{}, nil
	}

}

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == cr && line[len(line)-1] == lf {
			break
		}
	}
	if len(line) >= 2 && line[len(line)-2] == cr && line[len(line)-1] == lf {
		return line[:len(line)-2], n, nil
	}

	return line, n, nil
}

func (r *Resp) readArray() (RespMessage, error) {
	v := RespMessage{}
	v.Typ = "array"

	length, _, err := r.readInteger()
	if err != nil {
		return RespMessage{}, err
	}

	v.Array = make([]RespMessage, length)

	for i := range length {
		val, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return v, err
		}
		v.Array[i] = val
	}

	return v, nil
}

func (r *Resp) readString() (RespMessage, error) {
	v := RespMessage{}
	v.Typ = "string"

	length, _, err := r.readInteger()
	if err != nil {
		return RespMessage{}, err
	}

	// NOTE:
	// Need to handle edge case where the string length doesn't match actual bulk string.
	// currently it doesn't throw error if the bulk string length doesn't match the string length
	// $2\r\nfoo\r\n		<- this is fine
	// $3\r\nfo\r\n 		<- this should not be fine (maybe)
	// I'm not sure if it should error or not since it will only happen if user send raw RESP.
	StrBuf := make([]byte, length)
	r.reader.Read(StrBuf)
	v.String = string(bytes.TrimRight(StrBuf, "\r\n"))
	r.readLine()

	return v, nil
}

func (r *Resp) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, 0, nil
	}

	return int(i64), n, nil
}
