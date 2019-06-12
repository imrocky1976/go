/*
ROT13 ("rotate by 13 places", sometimes hyphenated ROT-13) is a simple letter substitution cipher that
replaces a letter with the 13th letter after it, in the alphabet. ROT13 is a special case of the Caesar
cipher which was developed in ancient Rome.

Because there are 26 letters (2×13) in the basic Latin alphabet, ROT13 is its own inverse; that is,
to undo ROT13, the same algorithm is applied, so the same action can be used for encoding and decoding.
The algorithm provides virtually no cryptographic security, and is often cited as a canonical example
of weak encryption.

ROT13 is used in online forums as a means of hiding spoilers, punchlines, puzzle solutions,
and offensive materials from the casual glance. ROT13 has inspired a variety of letter and word
games on-line, and is frequently mentioned in newsgroup conversations.
*/

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (v rot13Reader) Read(b []byte) (sz int, err error) {
	n, err := v.r.Read(b)

	for i := 0; i < n; i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = b[i] + 13
			if b[i] > 'z' {
				b[i] = 'a' - 1 + b[i] - 'z'
			}
		} else if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = b[i] + 13
			if b[i] > 'Z' {
				b[i] = 'A' - 1 + b[i] - 'Z'
			}
		}
	}

	return n, err
}

/*
func (rot *rot13Reader) Read(p []byte) (n int, err error) {
    n, err = rot.r.Read(p)
    for i := 0; i < len(p); i++ {
        if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
            p[i] += 13
        } else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
            p[i] -= 13
        }
    }
    return
}
*/
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
