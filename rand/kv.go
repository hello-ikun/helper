package rand

import (
	"bytes"
	cryptoRand "crypto/rand"
	"math/rand"
	"time"
	"unsafe"
)

const (
	Base64Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"
	Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	HexChars    = "0123456789abcdef"
	DecChars    = "0123456789"
)

type RandomStringGenerator struct {
	charset string
	length  int
}

func NewRandomStringGenerator(charset string, length int) *RandomStringGenerator {
	return &RandomStringGenerator{
		charset: charset,
		length:  length,
	}
}
func DefaultRandomGenerator(length int) *RandomStringGenerator {
	return &RandomStringGenerator{charset: Base62Chars, length: length}
}
func Base64CharsRandomGenerator(length int) *RandomStringGenerator {
	return &RandomStringGenerator{charset: Base64Chars, length: length}
}
func HexCharsRandomGenerator(length int) *RandomStringGenerator {
	return &RandomStringGenerator{charset: HexChars, length: length}
}
func DecCharsRandomGenerator(length int) *RandomStringGenerator {
	return &RandomStringGenerator{charset: DecChars, length: length}
}

// 生成随机字符 []byte 这样是不安全的 但是速度是很快的
func (r *RandomStringGenerator) Generate() []byte {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, r.length)
	for i := range b {
		b[i] = r.charset[randGen.Intn(len(r.charset))]
	}
	return b
}

func (r *RandomStringGenerator) SecureGenerate() []byte {
	randomBytes := make([]byte, r.length)

	if _, err := cryptoRand.Read(randomBytes); err != nil {
		panic(err)
	}

	for i := 0; i < r.length; i++ {
		randomBytes[i] = r.charset[int(randomBytes[i])%len(r.charset)]
	}
	return randomBytes
}

type FormatTestKey struct {
	prefix     string
	suffix     string
	bodyLength int
}

func NewFormatTestKey(prefix, suffix string, length int) *FormatTestKey {
	return &FormatTestKey{prefix: prefix, suffix: suffix, bodyLength: length}
}

const (
	defaultPrefix = "ikunHelper"
	defaultSuffix = "end"
)

func DefaultFormatTestKey(length int) *FormatTestKey {
	return &FormatTestKey{prefix: defaultPrefix, suffix: defaultSuffix, bodyLength: length}
}
func (f *FormatTestKey) Generate() []byte {
	var buffer bytes.Buffer

	if f.prefix != "" {
		buffer.WriteString(f.prefix)
		buffer.WriteByte('_')
	}
	r := DefaultRandomGenerator(f.bodyLength)
	buffer.Write(r.Generate())
	if f.prefix != "" {
		buffer.WriteByte('_')
		buffer.WriteString(f.suffix)
	}
	return buffer.Bytes()
}

func ByteToString(data []byte) string {
	return unsafe.String(&data[0], len(data))
}
