package chapter4

type CompositeKey struct {
	OriginalKey string
	key string
}

type CompositeKeys []CompositeKey

func (compositeKey CompositeKeys) Len() int {
	return len(compositeKey)
}

func (compositeKey CompositeKeys) Less(i, j int) bool {
	return compositeKey[i].OriginalKey < compositeKey[j].OriginalKey
}

func (compositeKey CompositeKeys) Swap(i, j int) {
	compositeKey[i], compositeKey[j] = compositeKey[j], compositeKey[i]
}

func (compositeKey CompositeKeys) Append(element CompositeKey) CompositeKeys {
	keys := append(compositeKey, element)

	return keys
}

