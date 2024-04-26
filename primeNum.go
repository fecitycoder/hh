func primeNum(min, max int) []int {
	var primes []int
	// Создаем Slice простых чисел от min До max
	for i := min; i <= max; i++ {
		count := 0
		for n := 2; n <= i; n++ { // На 1 делятся все числа (пропускаем)
			if i%n == 0 {
				count++ // счетчик количества деления без остатка
			}
		}
		if count == 1 { // Если делится только 1 раз (на себя же)
			primes = append(primes, i)
		}

	}
	return primes // Возвращаем Slice
}