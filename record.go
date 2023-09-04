package doclize

import (
	"sort"
)

type record struct {
	Index     []int
	LastIndex int
}

func (r *record) copy(content []byte, new []byte) string {
	sort.Ints(r.Index)
	var i = 0
	for ; i < len(new); i++ {
		if i < len(r.Index) {
			content[r.Index[i]] = new[i]
		} else {
			return string(append(content[r.Index[len(r.Index)-1]:], append(new[i:], content[:r.Index[len(r.Index)-1]]...)...))
		}
	}
	if i < len(r.Index) {
		r.LastIndex = r.Index[i] + 1
		return string(append(content[:r.Index[i]], content[r.Index[len(r.Index)-1]+1:]...))
	}
	//var remove int
	//var ans = tools.Make()
	//if len(r.Index) <= new.Len() {
	//	remove = new.Len() - 1
	//	ans.Append(content[:r.Index[0]], new, content[r.Index[0]+1:])
	//} else {
	//
	//}
	//r.LastIndex = r.Index[0] + new.Len()
	//content = ans.Bytes()
	//for i := 1; i < len(r.Index); i++ {
	//	content[r.Index[i]+remove] = ' '
	//}
	return string(content)
}
