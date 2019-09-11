package main

import (
	"bufio"
	"os"
	"strconv"
)

func checkMagazine(magazine []string, note []string) bool {
	hash := make(map[string]int)
	for _, m := range magazine {
		hash[m]++
	}
	for _, n := range note {
		if v, ok := hash[n]; !ok || v == 0 {
			return false
		}
		hash[n]--
	}
	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)

	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	magazine := make([]string, m)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	note := make([]string, n)

	for i := 0; i < m && sc.Scan(); i++ {
		magazine[i] = sc.Text()
	}

	for i := 0; i < n && sc.Scan(); i++ {
		note[i] = sc.Text()
	}

	result := checkMagazine(magazine, note)
	if result {
		out.WriteString("Yes")
	} else {
		out.WriteString("No")
	}
	out.Flush()
}
