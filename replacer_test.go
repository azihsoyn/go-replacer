package replacer

import "testing"

func TestGoReplacer(t *testing.T) {
	content := "1x10x100x1000"
	expect := "<1>x<10>x<100>x<1000>"
	dict := map[string]string{
		"1":    "<1>",
		"10":   "<10>",
		"100":  "<100>",
		"1000": "<1000>",
	}
	replacer := NewReplacer(dict)
	actual := replacer.Replace(content)
	if expect != actual {
		t.Errorf("expect %s but got %s\n", expect, actual)
	}
}

/* this fails
func TestStringsReplacer(t *testing.T) {
	content := "1x10x100x1000"
	expect := "<1>x<10>x<100>x<1000>"
	dict := []string{
		"1", "<1>",
		"10", "<10>",
		"100", "<100>",
		"1000", "<1000>",
	}
	replacer := strings.NewReplacer(dict...)
	actual := replacer.Replace(content)
	if expect != actual {
		t.Errorf("expect %s but got %s\n", expect, actual)
	}
}
*/

func BenchmarkNewGoReplacer(b *testing.B) {
	dict := map[string]string{
		"1":    "<1>",
		"10":   "<10>",
		"100":  "<100>",
		"1000": "<1000>",
	}
	for i := 0; i < b.N; i++ {
		_ = NewReplacer(dict)
	}
}

func BenchmarkReplaceGoReplacer(b *testing.B) {
	content := "1x10x100x1000"
	dict := map[string]string{
		"1":    "<1>",
		"10":   "<10>",
		"100":  "<100>",
		"1000": "<1000>",
	}
	replacer := NewReplacer(dict)
	for i := 0; i < b.N; i++ {
		_ = replacer.Replace(content)
	}
}
