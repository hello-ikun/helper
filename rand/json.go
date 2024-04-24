package rand

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	// 使用当前时间作为种子来初始化随机数生成器
	seed := time.Now().UnixNano()
	r = rand.New(rand.NewSource(seed))
}
func RandomJson() []byte {
	js := randomJSON(5, 5)
	// data, err := json.MarshalIndent(js, "\n", "\t")
	data, err := json.Marshal(js)
	if err != nil {
		panic(err)
	}
	return data
}

func generateMap(depth, maxDepth int) interface{} {
	if depth <= 0 || depth > maxDepth {
		return nil
	}
	data := make(map[string]interface{})
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		data[key] = randomJSON(depth-1, maxDepth)
	}
	return data
}

func generateArray(depth, maxDepth int) interface{} {
	if depth <= 0 || depth > maxDepth {
		return nil
	}
	arr := make([]interface{}, rand.Intn(5)+1)
	for i := range arr {
		arr[i] = randomJSON(depth-1, maxDepth)
	}
	return arr
}

// 随机生成字符串
func randomJSON(depth, maxDepth int) interface{} {
	if maxDepth == depth {
		switch rand.Intn(2) {
		case 0:
			return generateArray(depth, maxDepth)
		case 1:
			return generateMap(depth, maxDepth)
		}
	}
	if depth <= 0 {
		switch rand.Intn(9) {
		case 0, 1:
			return rand.Intn(100)
		case 2, 3:
			return rand.Float64() * 100
		case 4, 5:
			return fmt.Sprintf("String:%d", rand.Intn(100))
		case 6, 7:
			return rand.Intn(2) == 1
		default:
			return nil
		}
	}
	if depth > 0 {
		switch rand.Intn(15) {
		case 0, 1:
			return rand.Intn(100)
		case 2, 3:
			return rand.Float64() * 100
		case 4, 5:
			return fmt.Sprintf("String:%d", rand.Intn(100))
		case 6, 7:
			return rand.Intn(2) == 1
		case 8:
			return nil
		case 9, 10, 11:
			return generateArray(depth, maxDepth)
		default:
			return generateMap(depth, maxDepth)
		}
	}
	return nil
}
