package wordGame

func getWordList(letterSet []string) {
	var whereClause []string
	lengthClause := "length <= "
	lengthClause += string(len(letterSet))
	whereClause = append(whereClause, lengthClause)
	// iterate through alphabet
	// count occurences of letter in slice
	// if > 0, letter + "count <= letterCount"
	// else letter + "count = 0"
}
