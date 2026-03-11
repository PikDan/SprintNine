package main

// Пишите тесты в этом файле
import (
	"slices"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {

	data, err := generateRandomElements(10)

	if err != nil {
		t.Fatalf("ошибка генерации 10 элементов: %v", err)
	}

	if len(data) != 10 {
		t.Fatalf("Ожидалась длина 10, получено %d", len(data))
	}
}

func TestGenerateRandomElementsError(t *testing.T) {

	_, err := generateRandomElements(0)

	if err == nil {
		t.Fatalf("Ожидалась ошибка, но получено значение nil")
	}
}

func TestMaxChunksLessThanChunks(t *testing.T) {

	data := []int{1, 5, 3}

	_, err := maxChunks(data)
	if err == nil {
		t.Fatalf("ожидалась ошибка при len(data) < CHUNKS: %v", err)
	}

}

func TestMaximum(t *testing.T) {

	data := []int{1, 5, 3, 9, 2, 11, 66, 77, 88}

	max, err := maximum(data)

	if err != nil {
		t.Fatalf("неожиданная ошибка поиска max: %v", err)
	}

	if max != 88 {
		t.Fatalf("Ожидалось 88, получено: %d", max)
	}
}
func TestMaximumEmptySlice(t *testing.T) {

	_, err := maximum([]int{})

	if err == nil {
		t.Fatalf("Ожидалась ошибка, но получено значение nil")
	}
}
func TestMaxChunks(t *testing.T) {

	data := []int{1, 5, 3, 9, 2, 11, 4, 99, 22, 11}

	expected := slices.Max(data)

	result, err := maxChunks(data)
	if err != nil {
		t.Fatalf("неожиданная ошибка поиска maxChunks: %v", err)
	}
	if result != expected {
		t.Fatalf("Ожидалось %d, получено %d", expected, result)
	}
}
func TestMaxChunksLarge(t *testing.T) {

	data, err := generateRandomElements(10000)

	if err != nil {
		t.Fatalf("ошибка генерации 10000 элементов: %v", err)
	}

	expected, err := maximum(data)
	if err != nil {
		t.Fatalf("неожиданная ошибка поиска max: %v", err)
	}
	result, err := maxChunks(data)
	if err != nil {
		t.Fatalf("неожиданная ошибка поиска maxChunks: %v", err)
	}

	if result != expected {
		t.Fatalf("Ожидалось %d, получено %d", expected, result)
	}
}
