package cardvalid

import (
	"strconv"
	"strings"

	"github.com/RomanIkonnikov93/keeper/server/internal/models"
)

// CheckCard custom bank card data validator.
func CheckCard(card string) (bool, error) {

	if card == "" {
		return false, models.ErrNotValid
	}

	str := strings.Split(card, "")
	if len(str) != 28 {
		return false, models.ErrNotValid
	}

	arr := strings.Split(card, ",")

	for i := range arr {
		if len(arr[i]) < 1 {
			return false, models.ErrNotValid
		}
	}

	for i := range arr {
		_, err := strconv.Atoi(arr[i])
		if err != nil {
			return false, models.ErrNotValid
		}
	}

	if len(arr[1]) != 2 || len(arr[2]) != 4 || len(arr[3]) != 3 {
		return false, models.ErrNotValid
	}

	num, err := strconv.Atoi(arr[0])
	if err != nil {
		return false, models.ErrNotValid
	}

	res := LuhnValid(num)

	return res, nil
}

// LuhnValid check number is valid or not based on Luhn algorithm.
func LuhnValid(number int) bool {
	return (number%10+checksum(number/10))%10 == 0
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 {
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
