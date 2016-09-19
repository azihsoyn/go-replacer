package replacer

type replacer struct {
	tree tree
}

type node struct {
	value     byte
	next      tree
	replaceTo *string
}

type tree map[byte]*node

func NewReplacer(dict map[string]string) *replacer {
	root := make(tree)
	for kw, rp := range dict {
		current := root
		for i, bkw := range []byte(kw) {
			if current[bkw] == nil {
				var replaceTo *string
				// last
				if i == len([]byte(kw))-1 {
					rep := rp
					replaceTo = &rep
				}
				n := &node{
					value:     bkw,
					next:      make(tree),
					replaceTo: replaceTo,
				}
				current[bkw] = n
				current = n.next
				continue
			}
			var replaceTo *string
			// last
			if i == len([]byte(kw))-1 {
				rep := rp
				replaceTo = &rep
				current[bkw].replaceTo = replaceTo
			}
			current = current[bkw].next
		}
	}
	return &replacer{tree: root}
}

func (r *replacer) Replace(content string) string {
	current := r.tree
	str := make([]byte, 0, len(content))
	var replaceTo *string
	for _, c := range []byte(content) {
		if current[c] == nil {
			if replaceTo != nil {
				str = append(str, []byte(*replaceTo)...)
				replaceTo = nil
			}
			str = append(str, c)
			current = r.tree
			continue
		}
		if current[c].replaceTo != nil {
			replaceTo = current[c].replaceTo
		}
		current = current[c].next
	}
	if replaceTo != nil {
		str = append(str, []byte(*replaceTo)...)
	}
	return string(str)
}
