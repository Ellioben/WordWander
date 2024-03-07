package handle

func UniqueResource(collection []string, allCollection []string) []string {
	var uniqueWords []string

	for _, x := range collection {
		found := false
		for _, s := range allCollection {
			if x == s {
				found = true
				break
			}
		}
		if !found {
			uniqueWords = append(uniqueWords, x)
		}
	}

	collection = uniqueWords
	return collection
}
