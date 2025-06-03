// Вариант 6.
// Основной исполняемый файл для демонстрации модульных тестов, покрывающих основной функционал модулей.
// Разработчик: [Ковалева Алина]
// Лицензия: GPLv3
// История изменений:
// - 03.04.2025: Первоначальная реализация

package tests

import (
	"median-finder/algo"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name    string
		nums1   []int
		nums2   []int
		want    float64
		wantErr bool
	}{
		{
			name:  "empty second array",
			nums1: []int{1, 2},
			nums2: []int{},
			want:  1.5,
		},
		{
			name:  "all elements equal",
			nums1: []int{1, 1, 1},
			nums2: []int{1, 1, 1},
			want:  1.0,
		},
		{
			name:  "non-overlapping arrays 1",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  3.5,
		},
		{
			name:  "non-overlapping arrays 2",
			nums1: []int{4, 5, 6},
			nums2: []int{1, 2, 3},
			want:  3.5,
		},
		{
			name:  "single element in each",
			nums1: []int{1},
			nums2: []int{2},
			want:  1.5,
		},
		{
			name:  "boundary case 1",
			nums1: []int{1},
			nums2: []int{2, 3, 4, 5, 6},
			want:  3.5,
		},
		{
			name:  "boundary case 2",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{6},
			want:  3.5,
		},
		{
			name:  "duplicates at boundary",
			nums1: []int{1, 2, 3},
			nums2: []int{3, 4, 5},
			want:  3.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := algo.FindMedianSortedArrays(tt.nums1, tt.nums2)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMedianSortedArrays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("FindMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
