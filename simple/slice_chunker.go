package simple

func ChunkSlicer(slice []string, quantity int) [][]string {
	chunks := [][]string{}
	for i := 0; i < len(slice); i += quantity {
		end := i + quantity
		if end > len(slice) {
			end = len(slice) 
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}