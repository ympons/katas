package main

import (
	"bufio"
	"encoding/hex"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		key := scanner.Text()
		if scanner.Scan() {
			txt, _ := hex.DecodeString(scanner.Text())
			klen := len(key)
			for i := 0; i < len(txt); i++ {
				out.WriteString(string(txt[i] ^ key[i%klen]))
			}
		}
		out.Flush()
	}
}
