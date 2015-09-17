package main

import (
	"bufio"
	"os"
	"regexp"
)

func main() {
	idre := regexp.MustCompile(`id="question-summary-\d+`)
	h3re := regexp.MustCompile(`<a.*class="question-hyperlink">(.*?)</a>`)
	tre := regexp.MustCompile(`<span.*class="relativetime">(.*?)</span>`)

	idr := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(os.Stdin)
	wout := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		l := scanner.Text()
		id := idre.FindAllString(l, 1)
		for len(id) < 0 && scanner.Scan() {
			l = scanner.Text()
			id = idre.FindAllString(l, 1)
		}
		if len(id) > 0 {
			wout.WriteString(idr.FindAllString(id[0], 1)[0])
		}
		h3 := h3re.FindAllString(l, 1)
		for len(h3) < 0 && scanner.Scan() {
			l = scanner.Text()
			h3 = h3re.FindAllString(l, 1)
		}
		if len(h3) > 0 {
			wout.WriteString(";")
			wout.WriteString(h3re.ReplaceAllString(h3[0], "$1"))
		}
		t := tre.FindAllString(l, 1)
		for len(t) < 0 && scanner.Scan() {
			l = scanner.Text()
			t = tre.FindAllString(l, 1)
		}
		if len(t) > 0 {
			wout.WriteString(";")
			wout.WriteString(tre.ReplaceAllString(t[0], "$1"))
			wout.WriteString("\n")
		}
	}

	wout.Flush()
}
