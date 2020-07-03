package main

import (
	"sync"
	"testing"
)

var structSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]Basic, 100)
	},
}

var structPointerSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]*Basic, 100)
	},
}

type Basic struct {
	Id, N1, N2, N3, N4, N5 int
	Name                   string
}

func BenchmarkStructSliceWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var list []Basic
		for j := 0; j < 101; j++ {
			var data = Basic{Id: j, Name: "Name"}
			list = append(list, data)
		}
	}
}

func BenchmarkStructSliceWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := structSlicePool.Get().([]Basic)
		initLen := len(list) // 初始化长度为100
		for j := 0; j < 101; j++ {
			var data = Basic{Id: j, Name: "Name"}
			if j < initLen {
				list[j] = data
			} else {
				list = append(list, data)
			}
		}
		structSlicePool.Put(list)
	}
}

func BenchmarkStructPointerSliceWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var list []*Basic
		for j := 0; j < 101; j++ {
			var data = Basic{Id: j, Name: "Name"}
			list = append(list, &data)
		}
	}
}

func BenchmarkStructPointerSliceWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := structPointerSlicePool.Get().([]*Basic)
		initLen := len(list)
		for j := 0; j < 101; j++ {
			var data = Basic{Id: j, Name: "Name"}
			if j < initLen {
				list[j] = &data
			} else {
				list = append(list, &data)
			}
		}
		structPointerSlicePool.Put(list)
	}

}

/*
go test sync_pool_test.go -bench=. -benchmem
goos: darwin
goarch: amd64
BenchmarkStructSliceWithoutPool-4          	  199736	      5191 ns/op	   16320 B/op	       8 allocs/op
BenchmarkStructPointerSliceWithoutPool-4   	  172722	      6182 ns/op	    8504 B/op	     109 allocs/op
BenchmarkStructSliceWithPool-4             	 2017465	       594 ns/op	      32 B/op	       1 allocs/op
BenchmarkStructPointerSliceWithPool-4      	  208605	      5108 ns/op	    6496 B/op	     102 allocs/op
PASS
ok  	command-line-arguments	5.177s
*/
