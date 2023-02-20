package helpers

func Sum(numbers []float64) float64 {
	sum := float64(0)
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

type mapper[I any, O any] func(I) O

func ArrayMap[I any, O any](arr []I, mapper mapper[I, O]) []O {
	out := make([]O, len(arr))

	for i := range arr {
		out[i] = mapper((arr)[i])
	}

	return out
}

type testFunction[I any] func(I) bool

func ArrayFilter[I any](arr []I, testFunc testFunction[I]) []I {
	out := []I{}

	for i := range arr {
		if testFunc((arr)[i]) {
			out = append(out, (arr)[i])
		}
	}

	return out
}

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

func ArrayUniqueString(arr []string) []string {
	mapped := map[string]bool{}
	uarr := []string{}
	for _, s := range arr {
		if !mapped[s] {
			mapped[s] = true
			uarr = append(uarr, s)
		}
	}
	return uarr
}
