package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	pattern := regexp.MustCompile(`\b[A-Za-z_.]{1,}@[A-Za-z.]*[A-Za-z]{1,3}\b`)
	scanner := bufio.NewScanner(os.Stdin)
	v, r := make(map[string]bool), make([]string, 0, 1000)
	if scanner.Scan() {
		for n, _ := strconv.Atoi(scanner.Text()); n > 0 && scanner.Scan(); n-- {
			l := scanner.Text()
			m := pattern.FindAllString(l, len(l))
			for _, i := range m {
				if !v[i] {
					v[i] = true
					r = append(r, i)
				}
			}
		}
		sort.Strings(r)
		wout := bufio.NewWriter(os.Stdout)
		wout.WriteString(strings.Join(r, ";"))
		wout.Flush()
	}

}
