package application

import (
	"math"
	"strings"
	"unsafe"
)

func S2B(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s)
	return
}

// B2S []byte转string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

const hextable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Int64ToBytes62 int转bytes62
func Int64ToBytes62(num int64) []byte {
	var bytes []byte
	for num > 0 {
		bytes = append(bytes, hextable[num%62])
		num = num / 62
	}
	reverse := func(a []byte) {
		for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
			a[left], a[right] = a[right], a[left]
		}
	}
	reverse(bytes)
	return bytes
}

// Bytes62ToInt64  bytes62转int
func Bytes62ToInt64(str []byte) int64 {
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(hextable, str[i])
		num += int64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}
func RandBytes16() []byte {
	b := make([]byte, 16)
	for i := 0; i < 16; i++ {
		b[i] = hextable[RandIntn(62)]

	}
	return b
}
func RandBytes32() []byte {
	b := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b[i] = hextable[RandIntn(62)]

	}
	return b
}

// RandIntn  快速生成num范围内的随机数
func RandIntn(num uint32) int {
	return int(FastRand() % num)
}
