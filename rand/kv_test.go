package rand

import (
	"fmt"
	"testing"
)

func TestRandomKey(t *testing.T) {
	r := DefaultRandomGenerator(12)
	data1 := r.Generate()
	data2 := ByteToString(data1)
	fmt.Printf("%v %s %s\n", data1, data1, data2)
	data3 := r.SecureGenerate()
	data4 := ByteToString(data3)
	fmt.Printf("%v %s %s\n", data3, data3, data4)
	f := DefaultFormatTestKey(12)
	data5 := f.Generate()
	data6 := ByteToString(data5)
	fmt.Printf("%v %s %s\n", data5, data5, data6)

	// data := randomJSON(3, 3)
	// fmt.Println("data:", data)
	// data := RandomJson()
	// fmt.Println("data:", data, ByteToString(data))
	data0 := RandomJSON(5)
	fmt.Println("data:", data0, ByteToString(data0))
}
