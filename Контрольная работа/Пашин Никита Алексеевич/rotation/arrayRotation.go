// arrayRotation.go
// Модуль для циклического сдвига массива целых чисел
// Разработчик: Nikita Pashin
// Лицензия: GPLv3
// История изменений:
// 2025-04-03: Первая версия

package rotation

// RotateRight выполняет циклический сдвиг массива arr на n позиций вправо без использования дополнительного массива
func RotateRight(arr []int, n int) {
	length := len(arr)
	if length <= 1 || n == 0 {
		return
	}

	// Нормализуем n для случаев, когда n > length
	n = n % length

	// Переворачиваем весь массив
	reverse(arr, 0, length-1)

	// Переворачиваем первые n элементов
	reverse(arr, 0, n-1)

	// Переворачиваем оставшиеся элементы
	reverse(arr, n, length-1)
}

// reverse переворачивает элементы массива в диапазоне [start, end]
func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
