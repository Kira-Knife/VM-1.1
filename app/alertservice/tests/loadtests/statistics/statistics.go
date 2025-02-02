package statistics

import (
	"fmt"
	"loadtests/entity"
	"sort"
	"time"
)

// Функция для вычисления минимального времени запроса для каждого из StatusCode
func minExecutionTimeByStatus(reqRes []entity.RequestResult) map[int]time.Duration {
	minTimes := make(map[int]time.Duration)
	for _, result := range reqRes {
		if currentMin, exists := minTimes[result.StatusCode]; !exists || result.ExecutionTime < currentMin {
			minTimes[result.StatusCode] = result.ExecutionTime
		}
	}
	return minTimes
}

// Функция для вычисления максимального времени запроса для каждого из StatusCode
func maxExecutionTimeByStatus(reqRes []entity.RequestResult) map[int]time.Duration {
	maxTimes := make(map[int]time.Duration)
	for _, result := range reqRes {
		if currentMax, exists := maxTimes[result.StatusCode]; !exists || result.ExecutionTime > currentMax {
			maxTimes[result.StatusCode] = result.ExecutionTime
		}
	}
	return maxTimes
}

// Функция для вычисления числа ответов по каждому из StatusCode
func countByStatus(reqRes []entity.RequestResult) map[int]int {
	counts := make(map[int]int)
	for _, result := range reqRes {
		counts[result.StatusCode]++
	}
	return counts
}

// Функция для вычисления медианного времени по каждому из StatusCode
func medianExecutionTimeByStatus(reqRes []entity.RequestResult) map[int]time.Duration {
	medians := make(map[int]time.Duration)
	grouped := make(map[int][]time.Duration)

	for _, result := range reqRes {
		grouped[result.StatusCode] = append(grouped[result.StatusCode], result.ExecutionTime)
	}

	for statusCode, times := range grouped {
		sort.Slice(times, func(i, j int) bool { return times[i] < times[j] })
		mid := len(times) / 2
		if len(times)%2 == 0 {
			medians[statusCode] = (times[mid-1] + times[mid]) / 2
		} else {
			medians[statusCode] = times[mid]
		}
	}

	return medians
}

// Функция для вычисления среднего арифметического времени по каждому из StatusCode
func averageExecutionTimeByStatus(reqRes []entity.RequestResult) map[int]time.Duration {
	averages := make(map[int]time.Duration)
	totals := make(map[int]time.Duration)
	counts := make(map[int]int)

	for _, result := range reqRes {
		totals[result.StatusCode] += result.ExecutionTime
		counts[result.StatusCode]++
	}

	for statusCode, total := range totals {
		averages[statusCode] = total / time.Duration(counts[statusCode])
	}

	return averages
}

func PrintStatistics(title string, totalTime time.Duration, reqRes []entity.RequestResult) {
	fmt.Println(title)

	fmt.Println("Число ответов:")
	for status, count := range countByStatus(reqRes) {
		fmt.Printf("\tStatus Code %d: %d / %d\n", status, count, len(reqRes))
	}

	fmt.Println("Минимальное время запроса:")
	for status, minTime := range minExecutionTimeByStatus(reqRes) {
		fmt.Printf("\tStatus Code %d: %v\n", status, minTime)
	}

	fmt.Println("Максимальное время:")
	for status, maxTime := range maxExecutionTimeByStatus(reqRes) {
		fmt.Printf("\tStatus Code %d: %v\n", status, maxTime)
	}

	fmt.Println("Медианное время:")
	for status, medianTime := range medianExecutionTimeByStatus(reqRes) {
		fmt.Printf("\tStatus Code %d: %v\n", status, medianTime)
	}

	fmt.Println("Среднее арифметическое время:")
	for status, avgTime := range averageExecutionTimeByStatus(reqRes) {
		fmt.Printf("\tStatus Code %d: %v\n", status, avgTime)
	}

	fmt.Println()
}
