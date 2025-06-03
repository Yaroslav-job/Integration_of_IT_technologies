// Вариант 6.
// Основной исполняемый файл для демонстрации поиска медианы.
// Разработчик: [Ковалева Алина]
// Лицензия: GPLv3
// История изменений:
// - 03.04.2025: Первоначальная реализация
package main

import (
    "fmt"
    "median-finder/algo"
    //"github.com/yourusername/median-two-arrays/algo"
)

func main() {
    nums1 := []int{1, 3, 5}
    nums2 := []int{2, 4, 6}

    median, err := algo.FindMedianSortedArrays(nums1, nums2)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }

    fmt.Printf("Медиана массивов %v и %v: %.1f\n", nums1, nums2, median)
}