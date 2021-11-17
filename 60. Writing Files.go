package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	checkDirectors(err)

	f, err := os.Create("/tmp/dat2")
	checkDirectors(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkDirectors(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	checkDirectors(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	checkDirectors(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
