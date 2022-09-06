package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	num, err := rot.r.Read(b)
	if err == io.EOF {
		return 0, err
	}
	for i := 0; i < num; i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = 'A' + (b[i]-'A'+13)%26
		}
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = 'a' + (b[i]-'a'+13)%26
		}
	}
	return num, err

}

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		// if there is any update on the element of the buffer
		// the elements will be overwritten
		// any remaining element not overwritten, will be kept as it is
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	// Exercise rot13Reader
	// You cracked the code!
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot := rot13Reader{s}
	io.Copy(os.Stdout, &rot)

	fmt.Println()
}
