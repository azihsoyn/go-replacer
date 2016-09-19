package main

import (
	"fmt"

	replacer "github.com/azihsoyn/go-replacer"
)

func main() {
	content := "1x10x100x1000"
	expect := "<1>x<10>x<100>x<1000>"
	dict := map[string]string{
		"1":    "<1>",
		"10":   "<10>",
		"100":  "<100>",
		"1000": "<1000>",
	}
	replacer := replacer.NewReplacer(dict)
	actual := replacer.Replace(content)
	if expect != actual {
		fmt.Println("NG")
		return
	}
	fmt.Println("OK")
}
