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
func generateRandomElements(size int) ([]int, error) {
	// ваш код здесь
	if size <= 0 {
		return nil, fmt.Errorf("Размер должен быть > 0")
	}
	data := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < size; i++ {

		data[i] = r.Intn(9999)
	}
	return data, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	// ваш код здесь
	if len(data) < 1 {
		return 0, fmt.Errorf("Размер должен быть > 0")
	}
	maxNumber := slices.Max(data)
	return maxNumber, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	// ваш код здесь
	if len(data) < CHUNKS {
		return 0, fmt.Errorf("Размер должен быть > числа чанков")
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

		wg.Add(1)

		go func(i, firstIndx, lastIndx int) {
			defer wg.Done()

			localMaxNum := slices.Max(data[firstIndx:lastIndx])
			totalMaxVals[i] = localMaxNum
		}(i, firstIndx, lastIndx)

	}

	wg.Wait()

	return slices.Max(totalMaxVals), nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел ", SIZE)
	// ваш код здесь
	genetatedData, err := generateRandomElements(SIZE)
	if err != nil {
		panic(err)
	}

	fmt.Println("Ищем максимальное значение в один поток ")
	// ваш код здесь
	start := time.Now()

	max, err := maximum(genetatedData)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков ", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max, err = maxChunks(genetatedData)
	if err != nil {
		panic(err)
	}

	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())
}
