package mergesort

// Partition teilt das Array arr in zwei Teile, basierend auf dem ersten Element
// (dem Pivot). Alle Elemente kleiner oder gleich dem Pivot werden nach links,
// alle größeren nach rechts einsortiert. Gibt die neue Position des Pivotelements zurück.
func Partition(arr []int) int {
	pivot := arr[0]

	left := []int{}
	right := []int{}

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	for i := 0; i < len(left); i++ {
		arr[i] = left[i]
	}
	arr[len(left)] = pivot
	for i := 0; i < len(right); i++ {
		arr[i+len(left)+1] = right[i]
	}
	return 0
}

// QuickSort sortiert die übergebene Liste mittels des Quick-Sort-Algorithmus.
func QuickSort(arr []int) {
	// TODO
}
