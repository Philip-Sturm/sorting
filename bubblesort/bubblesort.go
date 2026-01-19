package bubblesort

// BubbleUp geht einmal durch das Array arr und tauscht benachbarte Elemente,
// wenn das linke größer ist als das rechte.
// Gibt true zurück, wenn mindestens ein Tausch stattgefunden hat.
func BubbleUp(arr []int) bool {
	check := false
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] > arr[j+1] {
			zwischenspeicher := arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = zwischenspeicher
			check = true
		}
	}
	return check
}

// BubbleSort sortiert die übergebene Liste mittels des Bubble-Sort-Algorithmus.
func BubbleSort(arr []int) {
	for i := len(arr); i > 0; i-- {
		BubbleUp(arr)
	}
}
