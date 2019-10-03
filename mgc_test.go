package mgc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMGC(t *testing.T) {
	// fh, err := os.Open("/usr/share/misc/magic.mgc")
	fh, err := os.Open("fortran.mgc")
	if err != nil {
		t.Fatal(err)
	}

	// data, err := ioutil.ReadAll(fh)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// spew.Dump(data)
	// fh.Seek(0, 0)

	scanner := bufio.NewScanner(fh)
	var i int
	for scanner.Scan() {

		i++
		if i == 10 {
			break
		}
		data := scanner.Bytes()
		// ignore comments
		if bytes.HasPrefix(data, []byte("#")) {
			println("comment")
			continue
		}
		parts := bytes.Split(data, []byte("\t"))
		offset := parts[0]
		offset = bytes.TrimRight(offset, "\x00")
		spew.Dump(offset)
		offset = bytes.TrimLeft(offset, "<")
		spew.Dump(offset)
		println(binary.BigEndian.Uint16(offset))
		for _, part := range parts {
			//fmt.Println(string(part))
			fmt.Printf("Line: %s\n", bytes.Trim(part, "\x00")[:])
		}
		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}
	}
	t.Fatal("woot")
}
