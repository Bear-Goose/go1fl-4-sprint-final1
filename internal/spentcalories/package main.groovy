package main

import (
	"fmt"
	"sort"
)

// Mode возвращает моды числовой последовательности.
func Mode(nums []int) ([]int, int) {
	if len(nums) == 0 {
		return []int{}, 1
	}

	// Подсчёт частот каждого числа
	freq := make(map[int]int)
	maxFreq := 0

	for _, num := range nums {
		freq[num]++
		if freq[num] > maxFreq {
			maxFreq = freq[num]
		}
	}

	// Если максимум равен 1 — моды как таковой нет, но по заданию нужно вернуть [] и 1
	if maxFreq == 1 {
		return []int{}, 1
	}

	// Сборка всех чисел с максимальной частотой
	var modes []int
	for num, count := range freq {
		if count == maxFreq {
			modes = append(modes, num)
		}
	}

	sort.Ints(modes)
	return modes, maxFreq
}

func main() {
    lists := [][]int{
        {},
        {57},
        {78, -7},
        {99, 200, 0},
        {4, 4, 4, 3},
        {102, -7, 44, -7, 102},
        {82, -23, 1, 5, 98, 100},
        {100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
    }
    modes := [][]int{
        {}, {}, {}, {},
        {4},
        {-7, 102}, {},
        {20000},
    }
    mcount := []int{
        1, 1, 1, 1, 3, 2, 1, 3,
    }
    for i, list := range lists {
        mode, count := Mode(list)
        if len(mode) != len(modes[i]) {
            fmt.Printf("len mode %d: %v != %v'\n", i, modes[i], mode)
        } else {
            for j, v := range mode {
                if v != modes[i][j] {
                    fmt.Printf("mode %d: %v != %v\n", i, modes[i], mode)
                }
            }
        }
        if count != mcount[i] {
            fmt.Printf("mcount %d: %d != %d\n", i, mcount[i], count)
        }
    }
    fmt.Println("Тестирование завершено")
} 