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

func BenchmarkMapNil(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < b.N; i++ {
		m[nil] = struct{}{}
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
		case 2:
			m[nil] = struct{}{}
		}
	}
}

func BenchmarkMapIntGet(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < 20000; i++ {
		m[i%1000] = struct{}{}
	}

	b.ResetTimer()

	var result struct{}
	for i := 0; i < b.N; i++ {
		result = m[i%1000]
	}

	_ = result
}

func BenchmarkMapStringGet(b *testing.B) {
	m := make(map[string]struct{})
	for i := 0; i < 20000; i++ {
		m[randStrs[i%len(randStrs)]] = struct{}{}
	}

	b.ResetTimer()

	var result struct{}
	for i := 0; i < b.N; i++ {
		result = m[randStrs[i%len(randStrs)]]
	}

	_ = result
}

func BenchmarkMapInterfaceIntGet(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < 20000; i++ {
		m[i%1000] = struct{}{}
	}

	b.ResetTimer()

	var result struct{}
	for i := 0; i < b.N; i++ {
		result = m[i%1000]
	}

	_ = result
}

func BenchmarkMapInterfaceStringGet(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < 20000; i++ {
		m[randStrs[i%len(randStrs)]] = struct{}{}
	}

	b.ResetTimer()

	var result struct{}
	for i := 0; i < b.N; i++ {
		result = m[randStrs[i%len(randStrs)]]
	}

	_ = result
}

func BenchmarkMapNilGet(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < 20000; i++ {
		m[nil] = struct{}{}
	}

	b.ResetTimer()

	var result struct{}
	for i := 0; i < b.N; i++ {
		result = m[nil]
	}

	_ = result
}

func BenchmarkMapInterfaceBothGet(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < 20000; i++ {
		switch i % 2 {
		case 0:
			m[randStrs[i%len(randStrs)]] = struct{}{}
		case 1:
			m[i%1000] = struct{}{}
		case 2:
			m[nil] = struct{}{}
		}
	}

	b.ResetTimer()

	var result struct{}

	for i := 0; i < b.N; i++ {
		switch i % 2 {
		case 0:
			result = m[randStrs[i%len(randStrs)]]
		case 1:
			result = m[i%1000]
		case 2:
			result = m[nil]
		}
	}

	_ = result
}

func BenchmarkMapIntRange(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < 5; i++ {
		m[i%1000] = struct{}{}
	}

	b.ResetTimer()
	var keyResult int
	var valueResult struct{}
	for i := 0; i < b.N; i++ {
		for k, v := range m {
			keyResult = k
			valueResult = v
		}
	}
	_ = keyResult
	_ = valueResult
}

func BenchmarkMapStringRange(b *testing.B) {
	m := make(map[string]struct{})
	for i := 0; i < 5; i++ {
		m[randStrs[i%len(randStrs)]] = struct{}{}
	}

	b.ResetTimer()
	var keyResult string
	var valueResult struct{}
	for i := 0; i < b.N; i++ {
		for k, v := range m {
			keyResult = k
			valueResult = v
		}
	}
	_ = keyResult
	_ = valueResult
}

func BenchmarkMapInterfaceIntRange(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < 5; i++ {
		m[i%1000] = struct{}{}
	}

	b.ResetTimer()
	var keyResult interface{}
	var valueResult struct{}
	for i := 0; i < b.N; i++ {
		for k, v := range m {
			keyResult = k
			valueResult = v
		}
	}
	_ = keyResult
	_ = valueResult
}

func BenchmarkMapInterfaceStringRange(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < 5; i++ {
		m[randStrs[i%len(randStrs)]] = struct{}{}
	}

	b.ResetTimer()
	var keyResult interface{}
	var valueResult struct{}
	for i := 0; i < b.N; i++ {
		for k, v := range m {
			keyResult = k
			valueResult = v
		}
	}
	_ = keyResult
	_ = valueResult
}

func BenchmarkMapInterfaceBothRange(b *testing.B) {
	m := make(map[interface{}]struct{})
	for i := 0; i < 5; i++ {
		switch i % 2 {
		case 0:
			m[randStrs[i%len(randStrs)]] = struct{}{}
		case 1:
			m[i%1000] = struct{}{}
		case 2:
			m[nil] = struct{}{}
		}
	}

	b.ResetTimer()
	var keyResult interface{}
	var valueResult struct{}
	for i := 0; i < b.N; i++ {
		for k, v := range m {
			keyResult = k
			valueResult = v
		}
	}
	_ = keyResult
	_ = valueResult
}
