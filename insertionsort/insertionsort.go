package insertionsort

// MoveLeft bewegt das Element an Stelle i so lange nach links,
// bis es an der richtigen Stelle im bereits sortierten Teil der Liste ist.
func MoveLeft(arr []int, i int) {
	for i > 0 && arr[i] < arr[i-1] {
		j := arr[i]
		arr[i] = arr[i-1]
		arr[i-1] = j
		i--
	}
}

// InsertionSort sortiert die übergebene Liste mittels des Insertion-Sort-Algorithmus.
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		MoveLeft(arr, i)
	}
}
