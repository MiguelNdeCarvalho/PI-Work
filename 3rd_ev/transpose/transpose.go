package transpose

func Transpose(input []string) []string {

	output := []string{}

	for linha1, linha2 := range input {

		for coluna, i := range linha2 {

			if len(output) <= coluna {

				output = append(output, "")
			}

			for len(output[coluna]) < linha1 {

				output[coluna] += " "
			}

			output[coluna] += string(i)
		}
	}

	return output
}