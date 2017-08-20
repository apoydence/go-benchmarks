package bench_test

import (
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"testing"
)

func BenchmarkHashAdl32(b *testing.B) {
	defer b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		adler32.Checksum([]byte(randStrs[i%len(randStrs)]))
	}
}

func BenchmarkHashCrc32IEEE(b *testing.B) {
	defer b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		crc32.ChecksumIEEE([]byte(randStrs[i%len(randStrs)]))
	}
}

var tableECMA = crc64.MakeTable(crc64.ECMA)

func BenchmarkHashCrc64ECMA(b *testing.B) {
	defer b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		crc64.Checksum([]byte(randStrs[i%len(randStrs)]), tableECMA)
	}
}

var tableISO = crc64.MakeTable(crc64.ISO)

func BenchmarkHashCrc64ISO(b *testing.B) {
	defer b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		crc64.Checksum([]byte(randStrs[i%len(randStrs)]), tableISO)
	}
}

func BenchmarkHashFnv32(b *testing.B) {
	defer b.ReportAllocs()
	f := fnv.New32()

	var hash uint32
	for i := 0; i < b.N; i++ {
		f.Reset()
		f.Write([]byte(randStrs[i%len(randStrs)]))
		hash = f.Sum32()
		_ = hash
	}
}

func BenchmarkHashFnv64(b *testing.B) {
	defer b.ReportAllocs()
	f := fnv.New64()

	var hash uint64
	for i := 0; i < b.N; i++ {
		f.Reset()
		f.Write([]byte(randStrs[i%len(randStrs)]))
		hash = f.Sum64()
		_ = hash
	}
}

func BenchmarkHashFnv32A(b *testing.B) {
	defer b.ReportAllocs()
	f := fnv.New32a()

	var hash uint32
	for i := 0; i < b.N; i++ {
		f.Reset()
		f.Write([]byte(randStrs[i%len(randStrs)]))
		hash = f.Sum32()
		_ = hash
	}
}

func BenchmarkHashFnv64A(b *testing.B) {
	defer b.ReportAllocs()
	f := fnv.New64a()

	var hash uint64
	for i := 0; i < b.N; i++ {
		f.Reset()
		f.Write([]byte(randStrs[i%len(randStrs)]))
		hash = f.Sum64()
		_ = hash
	}
}
