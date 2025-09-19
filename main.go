package main

import (
	"fmt"
	"math"
	"sort"
)

func Median(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	copy_of_numbers := make([]int, len(numbers))
	copy(copy_of_numbers, numbers)
	sort.Ints(copy_of_numbers)
	if len(copy_of_numbers)%2 != 0 {
		return copy_of_numbers[len(copy_of_numbers)/2]
	} else {

		mid := len(copy_of_numbers) / 2

		return int(math.Floor(float64(copy_of_numbers[mid-1]+copy_of_numbers[mid]) / 2.0))
	}
}

func Mode(numbers []int) ([]int, int) {
	if len(numbers) == 0 {
		return []int{}, 1
	}
	map_of_numbers := make(map[int]int)
	for _, val := range numbers {
		map_of_numbers[val]++
		//fmt.Println("map", map_of_numbers)
	}
	max_count, mode := 0, []int{}
	for _, count := range map_of_numbers {
		if count > max_count {
			max_count = count

		}

	}
	if max_count == 1 {
		return []int{}, 1
	}
	//fmt.Println("max_count", max_count)

	for num, count := range map_of_numbers {
		if count == max_count {
			mode = append(mode, num)
			sort.Ints(mode)
		}
	}
	return mode, max_count

}

func Average(numbers []int) float64 {
	sum := 0
	if len(numbers) == 0 {
		return 0

	}
	for _, val := range numbers {

		sum += val
	}
	return float64(sum) / float64(len(numbers))
}

func Range(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min_number := numbers[0]
	max_number := numbers[0]
	for _, val := range numbers {

		if min_number > val {
			min_number = val

		}
		if max_number < val {
			max_number = val
		}

	}
	return max_number - min_number
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
	averages := []float64{
		0, 57, 36, 100, 4, 47, 44, 39938,
	}
	ranges := []int{
		0, 0, 85, 200, 1, 109, 123, 80000,
	}
	medians := []int{
		0, 57, 35, 99, 4, 44, 43, 22000,
	}
	modes := [][]int{
		{}, {}, {}, {},
		{4},
		{-7, 102},
		{},
		{20000},
	}
	mcount := []int{
		1, 1, 1, 1, 3, 2, 1, 3,
	}

	for i, list := range lists {
		if average := math.Round(Average(list)); average != averages[i] {
			fmt.Printf("average %d: %.2f != %.2f", i, averages[i], average)

		}
		if r := Range(list); r != ranges[i] {
			fmt.Printf("range %d: %d != %d", i, ranges[i], r)

		}
		if median := Median(list); median != medians[i] {
			fmt.Printf("median %d: %d != %d\n", i, medians[i], median)

		}
		mode, count := Mode(list)
		if len(mode) != len(modes[i]) {
			fmt.Printf("len mode %d: %v != %v'\n", i, modes[i], mode)

		} else {
			for j, v := range mode {
				if v != modes[i][j] {
					fmt.Printf("mode %d: %v != %v\n", i, modes[i], mode)

					break
				}
			}
		}
		if count != mcount[i] {
			fmt.Printf("mcount %d: %d != %d\n", i, mcount[i], count)

		}
	}

	fmt.Println("Тестирование завершено")
}
