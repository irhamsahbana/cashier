package mapper

func DeleteKeys(m map[string]interface{}, keys ...string) {
	for _, k := range keys {
		delete(m, k)
	}
}
