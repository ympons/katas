package main

import (
  "bufio"
  "os"
  "strconv"
)

func main() {
  s := bufio.NewScanner(os.Stdin)
  s.Split(bufio.ScanWords)
  o := bufio.NewWriter(os.Stdout)
  s.Scan()
  n, _ := strconv.Atoi(s.Text())
  s.Scan()
  m, _ := strconv.Atoi(s.Text())
  h := make(map[string]int)
  for t := n; t>0 && s.Scan(); t--{
    h[s.Text()]++    
  }
  flag := true
  for t := m; t>0 && s.Scan(); t--{
    if v, ok := h[s.Text()]; ok && v >= 0 {
      h[s.Text()]--
    } else {
      flag = false
      break
    }
  }
  if flag {
    o.WriteString("Yes\n")
  } else {
    o.WriteString("No\n")
  }
  o.Flush()
}
