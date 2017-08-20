package bench_test

import (
	"crypto/rand"
	"os"
	"testing"
)

var randStrs []string

func TestMain(m *testing.M) {
	buf := make([]byte, 25)
	for i := 0; i < 1000; i++ {
		rand.Read(buf)
		randStrs = append(randStrs, string(buf))
	}

	os.Exit(m.Run())
}

func BenchmarkMapInt(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < b.N; i++ {
		m[i%1000] = struct{}{}
	}
}

func BenchmarkMapString(b *testing.B) {
	m := make(map[string]struct{})
	for i := 0; i < b.N; i++ {
		m[randStrs[i%len(randStrs)]] = struct{}{}
	}
}

func BenchmarkMapInterfaceInt(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < b.N; i++ {
		m[i%1000] = struct{}{}
	}
}

func BenchmarkMapInterfaceString(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < b.N; i++ {
		m[randStrs[i%len(randStrs)]] = struct{}{}
	}
}

func BenchmarkMapInterfaceBoth(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < b.N; i++ {
		switch i % 2 {
		case 0:
			m[randStrs[i%len(randStrs)]] = struct{}{}
		case 1:
			m[i%1000] = struct{}{}
		}
	}
}
