package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func main() {
	root := flag.String("p", ".", "path to scan as root")
	key := flag.String("k", "", "key to match")
	matchType := flag.String("t", "go", "file type to match")
	reserve := flag.Bool("r", false, "reserve mode")
	verbose := flag.Bool("v", false, "verbose display")
	limit := flag.Int("l", 3, "overview limit for single match")
	flag.Parse()
	var matched []string
	scanned := Walk(DerefString(root))
	var list []string
	for _, s := range scanned {
		for _, t := range strings.Split(*matchType, ",") {
			if strings.Contains(filepath.Ext(s), t) {
				list = append(list, s)
				break
			}
		}
	}
OUTER:
	for _, v := range list {
		//fmt.Println(vv)
		match := false
		var related []string
		lines := Split(v)
		if len(lines) == 0 {
			continue
		}
		for _, vv := range lines {
			if len(vv) > 0 {
				for _, vvv := range strings.Split(*key, ",") {
					if len(related) >= 1 && !match {
						match = true
						matched = append(matched, v)
					}
					if len(related) >= *limit {
						continue OUTER
					}
					if strings.Contains(vv, vvv) {
						if !*reserve {
							match = true
							related = append(related, vv)
							fmt.Print(colorCyan, "--------------------------------", "\n")
							if *verbose {
								fmt.Print(colorCyan, "Path: ", v, "\n")
								fmt.Print(colorPurple, "Count: ", strconv.Itoa(len(related)), "\n")
								fmt.Print(colorGreen, "text matched:\n")
							} else {
								fmt.Print(colorPurple, v, ": ")
							}
							index := strings.Index(vv, vvv)
							if index == -1 {
								fmt.Println(vv)
								continue
							}
							before := vv[:index]
							after := vv[index+len(vvv):]
							var tmp []string
							tmp = append(tmp, colorWhite)
							tmp = append(tmp, before)
							tmp = append(tmp, colorRed)
							tmp = append(tmp, *key)
							tmp = append(tmp, colorWhite)
							tmp = append(tmp, after)
							fmt.Print(strings.Join(tmp, "") + "\n")
							continue
						}
					}
				}
			} else {
				continue
			}
		}
	}
	if *reserve {
		var relist []string
	ROUTER:
		for _, l := range list {
			for _, m := range matched {
				if m == l {
					continue ROUTER
				}
			}
			relist = append(relist, l)
		}
		for _, rf := range relist {
			fmt.Println(colorPurple + rf)
		}
	}
}