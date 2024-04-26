func primeNum(inArr []int) []int {
	var cDivisor []int
	for i := 2; i <= slices.Min(inArr); i++ { // Судя из условий 1 мы не учитываем, поэтому начнем с 2х и до самого минимального в массиве
		count := 0
		for _, val := range inArr {
			if val%i == 0 {
				count++
			}
		}
		if count == len(inArr) { // Если все перебранные числа массива поделились
			cDivisor = append(cDivisor, i)
		}
	}

	return cDivisor // Возвращаем Slice
}
