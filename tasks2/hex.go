package main

func hex(s string) string {
	mapping := map[int32]string{
		'a': "10",
		'b': "11",
		'c': "12",
		'd': "13",
		'e': "14",
		'f': "15",
	}
	res := ""
	for _, c := range s {
		res += mapping[c]
	}

	return res
}
