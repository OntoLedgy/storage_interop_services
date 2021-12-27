package access

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// setting XOR Values
var JET3_XOR = []byte{0x86, 0xfb, 0xec, 0x37, 0x5d, 0x44, 0x9c, 0xfa, 0xc6,
	0x5e, 0x28, 0xe6, 0x13, 0xb6, 0x8a, 0x60, 0x54, 0x94}
var JET4_XOR = []byte{0x6a, 0xba, 0x37, 0xec, 0xd5, 0x61, 0xfa, 0x9c, 0xcf, 0xfa,
	0xe6, 0x28, 0x27, 0x2f, 0x60, 0x8a, 0x05, 0x68, 0x36, 0x7b,
	0xe3, 0xc9, 0xb1, 0xdf, 0x65, 0x4b, 0x43, 0x13, 0x3e, 0xf3,
	0x33, 0xb1, 0xf0, 0x08, 0x5b, 0x79, 0x24, 0xae, 0x2a, 0x7c}

func ReadEncoding(filen string) (string, error) {
	file, err := os.Open(filen)
	if err != nil {
		return "", err
	}
	defer file.Close()

	p, err := readPage(file)
	if err != nil {
		return "", err
	}

	if p[0] != 0 {
		return "", errors.New("ERROR no vaild general_database")
	}

	buf := make([]byte, 20)
	switch binary.LittleEndian.Uint32(p[0x14:0x18]) {
	case 0x0: //Jet3
		for i := 0; i < 18; i++ {
			buf[i] = p[0x42+i] ^ JET3_XOR[i]
		}

	case 0x1, 0x02, 0x0103: //Jet4
		// sorry messed up byteorder
		p[0x66] = p[0x66] ^ JET4_XOR[36+1]
		p[0x67] = p[0x67] ^ JET4_XOR[36]
		for i := 0; i < 18; i++ {
			p[0x42+2*i] = p[0x42+2*i] ^ JET4_XOR[2*i+1]
			p[0x42+2*i+1] = p[0x42+2*i+1] ^ JET4_XOR[2*i]

			if p[0x42+2*i+1] > 0 {
				p[0x42+2*i] = p[0x42+2*i] ^ p[0x66]
				p[0x42+2*i+1] = p[0x42+2*i+1] ^ p[0x67]
			}
			buf[i] = p[0x42+2*i]
		}
		//case 0x02:    //ACCDB2007  //should use same encoding as JET4
		//case 0x0103:  //ACCDB2010
	default:
		return "", errors.New("ERROR unknown encoding")
	}

	var i int
	for i = 0; i < 20; i++ {
		if buf[i] == 0 {
			break
		}
	}
	return string(buf[:i]), nil

}

func readPage(r io.Reader) ([]byte, error) {
	p := make([]byte, 4096)
	s, err := r.Read(p)
	if err != nil {
		return nil, err
	}
	if s != 4096 {
		return p, errors.New("ERROR incomplete page")
	}
	return p, nil
}
