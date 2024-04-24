package rand

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
)

var maxDepth int

func RandomJSON(depath int) []byte {
	maxDepth = depath
	js := randomJson(maxDepth)
	data, err := json.Marshal(js)
	if err != nil {
		panic(err)
	}
	return data
}

func generateToMap(depth int) interface{} {
	if depth <= 0 {
		return nil
	}
	data := make(map[string]interface{})
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		data[key] = randomJson(depth - 1)
	}
	return data
}

func generateToArray(depth int) interface{} {
	if depth <= 0 {
		return nil
	}
	arr := make([]interface{}, randomInt(5)+1)
	for i := range arr {
		arr[i] = randomJson(depth - 1)
	}
	return arr
}

func randomJson(depth int) interface{} {
	if maxDepth == depth {
		switch randomInt(3) {
		case 0:
			return generateArray(depth, maxDepth)
		case 1:
			return generateMap(depth, maxDepth)
		}
	}
	if depth <= 0 {
		switch randomInt(9) {
		case 0, 1:
			return randomInt(100)
		case 2, 3:
			return randomFloat() * 100
		case 4, 5:
			return fmt.Sprintf("String:%d", randomInt(100))
		case 6, 7:
			return randomInt(2) == 1
		default:
			return nil
		}
	}

	switch randomInt(9) {
	case 0, 1:
		return randomInt(100)
	case 2, 3:
		return randomFloat() * 100
	case 4, 5:
		return fmt.Sprintf("String:%d", randomInt(100))
	case 6, 7:
		return randomInt(2) == 1
	default:
		if depth < maxDepth {
			if randomInt(2) == 0 {
				return generateToArray(depth)
			}
			return generateToMap(depth)
		}
		return nil
	}
}

func randomInt(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return n.Int64()
}

func randomFloat() float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		panic(err)
	}
	return float64(n.Int64()) / 100.0
}
