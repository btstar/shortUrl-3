package tools

import (
	"math"
	"errors"
)

var (
	// 64进制使用到的字符列表(编码使用)
	endCode = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/")
	// 64进制使用到的字符map(解码使用)
	deCode = map[rune]int{}

	// 64进制
	SYSTEM uint64 = 64
)

func init() {
	for k, v := range endCode {
		deCode[v] = k
	}
}

// 编码
func Encode(id uint64) (string, error) {
	var data []rune
	for {
		var r rune   // 下标指向的字符
		var k uint64 // 64进制字符数组下标
		if id < SYSTEM {
			k = id
			r = endCode[k]
			data = append([]rune{r}, data...)
			break
		} else {
			k = id % SYSTEM
			r = endCode[k]
			data = append([]rune{r}, data...)

			id = (id - k) / SYSTEM
		}
	}

	return string(data), nil
}

// 解码
func Decode(str string) (uint64, error) {
	strRune := []rune(str) // 字符串转字符数组

	l := len(strRune)
	zs := l - 1 // 当前位指数
	var value uint64
	for i := 0; i < l; i++ {
		number, err := searchV(strRune[i])
		if err != nil {
			return 0, err
		}

		value += uint64(math.Pow(float64(SYSTEM), float64(zs))) * number
		zs--
	}

	return value, nil
}

// 过去字符在定义好的字符数组中的位置
func searchV(rune rune) (uint64, error) {

	k, ok := deCode[rune]
	if !ok {
		return 0, errors.New("字符不存在")
	}

	return uint64(k), nil

}
