package snowflake

import (
	"testing"
)

func BenchmarkSnowFlake(b *testing.B) {
	test := make(map[uint64]uint64)
	for i := 0; i < b.N; i++ {
		id := GlobalSnowflake.Generate().UInt64()
		if _, ok := test[id]; ok {
			//fmt.Print(id)
		}
		test[id] = id
	}
}

func BenchmarkSnowFlake2(b *testing.B) {
	test := make(map[uint64]uint64)
	for i := 0; i < b.N; i++ {
		l, _ := New(0)
		id := l.Generate().UInt64()
		if _, ok := test[id]; ok {
			//fmt.Print(id)
		}
		test[id] = id
	}
}
