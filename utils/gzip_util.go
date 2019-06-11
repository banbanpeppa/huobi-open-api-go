package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

func ParseGzip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	if err != nil {
		fmt.Println("[ParseGzip] NewReader error: , maybe data is ungzip: ", err, string(data))
		fmt.Println("[ParseGzip] NewReader error: , maybe data is ungzip: ", err, string(b.Bytes()))
		return nil, err
	} else {
		defer r.Close()
		undatas, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Println("[ParseGzip]  ioutil.ReadAll error: :", err, string(data))
			return nil, err
		}
		return undatas, nil
	}
}
