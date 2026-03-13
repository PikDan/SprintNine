package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}
	data := make([]int, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {

		data[i] = r.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) < 1 {
		return 0
	}
	maxNumber := slices.Max(data)
	return maxNumber
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) < CHUNKS {
		return 0
	}
	var wg sync.WaitGroup
	totalMaxVals := make([]int, CHUNKS)

	chopData := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {

		firstIndx := i * chopData
		lastIndx := firstIndx + chopData

		if i == CHUNKS-1 {
			lastIndx = len(data)
		}
		chunk := data[firstIndx:lastIndx]

		wg.Add(1)

		go func(i int, chunk []int) {
			defer wg.Done()

			localMaxNum := maximum(chunk)
			totalMaxVals[i] = localMaxNum
		}(i, chunk)

	}

	wg.Wait()

	return maximum(totalMaxVals)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел ", SIZE)
	// ваш код здесь
	genetatedData := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток ")
	// ваш код здесь
	start := time.Now()

	max := maximum(genetatedData)

	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков ", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(genetatedData)

	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())
}
