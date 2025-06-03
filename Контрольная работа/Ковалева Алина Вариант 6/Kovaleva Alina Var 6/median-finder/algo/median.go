// Вариант 6.
// Package algo реализует алгоритмы поиска медианы в двух отсортированных массивах.
// Алгоритм имеет сложность O(log(min(n,m))), где n и m - длины массивов.
// Разработчик: Ковалева Алина
// Лицензия: GPLv3
// История изменений:
// - 03.04.2025: Первоначальная реализация

package algo

import "errors"

// FindMedianSortedArrays находит медиану двух отсортированных массивов.
// Возвращает ошибку если массивы пусты.
func FindMedianSortedArrays(nums1 []int, nums2 []int) (float64, error) {
    // Убедимся, что nums1 - это более короткий массив
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    m, n := len(nums1), len(nums2)
    if n == 0 {
        return 0, errors.New("оба массива пусты")
    }

    imin, imax, halfLen := 0, m, (m+n+1)/2

    for imin <= imax {
        i := (imin + imax) / 2
        j := halfLen - i

        if i < m && nums2[j-1] > nums1[i] {
            // i слишком маленький, нужно увеличить
            imin = i + 1
        } else if i > 0 && nums1[i-1] > nums2[j] {
            // i слишком большой, нужно уменьшить
            imax = i - 1
        } else {
            // i найден
            maxLeft := 0
            if i == 0 {
                maxLeft = nums2[j-1]
            } else if j == 0 {
                maxLeft = nums1[i-1]
            } else {
                maxLeft = max(nums1[i-1], nums2[j-1])
            }

            if (m+n)%2 == 1 {
                return float64(maxLeft), nil
            }

            minRight := 0
            if i == m {
                minRight = nums2[j]
            } else if j == n {
                minRight = nums1[i]
            } else {
                minRight = min(nums1[i], nums2[j])
            }

            return float64(maxLeft+minRight) / 2.0, nil
        }
    }

    return 0, errors.New("не удалось найти медиану")
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}