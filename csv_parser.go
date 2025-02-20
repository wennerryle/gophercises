package main

func csv_parser(data string) uint {
	amount_of_columns := get_columns_amount(data)

	return amount_of_columns
}

func get_columns_amount(data string) uint {
	var amount uint
	var in_quote bool

	for _, v := range data {
		if v == '"' && in_quote {
			in_quote = false
		}

		if v == '"' {
			in_quote = true
		}

		if !in_quote && v == ',' {
			amount++
		}

		if v == '\n' {
			break
		}
	}

	return amount
}
