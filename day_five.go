package main

func Split(min int, max int, character string) (int, int) {
	if character == "F" || character == "L" {
		return min, min + (max-min)/2 
	}

	return min+(max-min+2-1)/2, max
}
