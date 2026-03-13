package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {

	tests := []struct {
		name    string
		size    int
		wantLen int
		wantNil bool
	}{
		{
			name:    "valid size",
			size:    10,
			wantLen: 10,
			wantNil: false,
		},
		{
			name:    "zero size",
			size:    0,
			wantNil: true,
		},
		{
			name:    "negative size",
			size:    -5,
			wantNil: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			data := generateRandomElements(tt.size)

			if tt.wantNil {
				require.Nil(t, data)
				return
			}

			require.NotNil(t, data)
			require.Len(t, data, tt.wantLen)
		})
	}
}

func TestMaximum(t *testing.T) {

	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "normal slice",
			data: []int{1, 5, 3, 9, 2, 11, 66, 77, 88},
			want: 88,
		},
		{
			name: "single element",
			data: []int{42},
			want: 42,
		},
		{
			name: "negative numbers",
			data: []int{-10, -3, -7, -1},
			want: -1,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := maximum(tt.data)

			require.Equal(t, tt.want, result)
		})
	}
}

func TestMaxChunks(t *testing.T) {

	tests := []struct {
		name string
		data []int
	}{
		{
			name: "small slice",
			data: []int{1, 5, 3, 9, 2, 11, 4, 99, 22, 11},
		},
		{
			name: "larger slice",
			data: []int{5, 12, 3, 44, 7, 19, 100, 2, 55, 78, 34, 90},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			expected := maximum(tt.data)

			result := maxChunks(tt.data)

			require.Equal(t, expected, result)
		})
	}
}

func TestMaxChunksLarge(t *testing.T) {

	data := generateRandomElements(10000)
	require.NotNil(t, data)

	expected := maximum(data)

	result := maxChunks(data)

	require.Equal(t, expected, result)
}
