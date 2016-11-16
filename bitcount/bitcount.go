package bitcount

import (
	"fmt"
	"math/rand"
)

func InitializeBitsPerBytesLookupTable() []uint8 {
	table := make([]uint8, 256)
	for i := 0; i < 256; i++ {
		table[i] = 0
		var j uint8
		for j = 0; j < 8; j++ {
			if i&(1<<j) != 0 {
				table[i]++
			}
		}
	}
	return table
}

func GenerateData(seed int64, count uint) []uint16 {
	data := make([]uint16, count)
	r := rand.New(rand.NewSource(seed))
	var (
		i uint
		x uint16
	)
	for i = 0; i < count; i++ {
		x = uint16(r.Uint32())
		data[i] = x
	}
	return data
}

func CountBitsLookupTable(table []uint8, data []uint16) uint {
	var sum uint = 0
	for _, number := range data {
		sum += uint(table[uint8(number)] + table[uint8(number>>8)])
	}
	return sum
}

func SprintfLookupTable(table []uint8) string {
	s := ""
	for i, j := range table {
		if i != 0 && i%8 == 0 {
			s += fmt.Sprintf("\n")
		}
		s += fmt.Sprintf("%6s", fmt.Sprintf("%d:%d", i, j))
	}
	return s
}

func CountBitsTrivial(data []uint16) uint {
	var (
		sum uint = 0
		i   uint8
	)
	for _, number := range data {
		for i = 0; i < 16; i++ {
			if uint(number&(1<<i)) != 0 {
				sum++
			}
		}
	}
	return sum
}
