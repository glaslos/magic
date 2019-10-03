package mgc

import (
	"testing"

	"bitbucket.org/taruti/mimemagic"
	"github.com/davecgh/go-spew/spew"
)

func TestMIME(t *testing.T) {
	out := mimemagic.Match("banana", []byte("bananana"))
	spew.Dump(out)
}
