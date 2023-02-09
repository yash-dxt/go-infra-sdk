package helpers

func ArrayChunk[I any](arr []I, chunkSize int) [][]I {
	chunked := [][]I{}

	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		chunked = append(chunked, arr[i:end])
	}

	return chunked
}
