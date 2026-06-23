package main

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	digits = []string{"", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
	units  = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}
)

// BahtToThaiText converts a decimal currency value into its Thai language representation.
func BahtToThaiText(amount decimal.Decimal) string {
	if amount.IsZero() {
		return "ศูนย์บาท"
	}

	var result strings.Builder
	result.Grow(128)

	if amount.IsNegative() {
		result.WriteString("ลบ")
		amount = amount.Neg()
	}

	baht := amount.Truncate(0).IntPart()
	satang := amount.Mul(decimal.NewFromInt(100)).IntPart() % 100

	if baht > 0 {
		result.WriteString(numberToThai(baht))
		result.WriteString("บาท")
	}

	if satang > 0 {
		result.WriteString(numberToThai(satang))
		result.WriteString("สตางค์")
	} else {
		result.WriteString("ถ้วน")
	}

	return result.String()
}

func numberToThai(number int64) string {
	if number == 0 {
		return ""
	}

	var groups []int64
	for n := number; n > 0; n /= 1000000 {
		groups = append(groups, n%1000000)
	}

	var b strings.Builder
	b.Grow(128 * len(groups))

	for i := len(groups) - 1; i >= 0; i-- {
		if groups[i] == 0 {
			continue
		}
		writeSubMillion(&b, groups[i])
		for j := 0; j < i; j++ {
			b.WriteString("ล้าน")
		}
	}

	return b.String()
}

func writeSubMillion(b *strings.Builder, number int64) {
	numStr := fmt.Sprintf("%d", number)
	length := len(numStr)

	for i, digitChar := range numStr {
		digit := int(digitChar - '0')
		position := length - i - 1

		if digit == 0 {
			continue
		}

		if position == 1 {
			if digit == 2 {
				b.WriteString("ยี่")
			} else if digit != 1 {
				b.WriteString(digits[digit])
			}
			b.WriteString(units[1])
			continue
		}

		if position == 0 && digit == 1 && number >= 10 {
			b.WriteString("เอ็ด")
			continue
		}

		b.WriteString(digits[digit])
		if position > 0 {
			b.WriteString(units[position])
		}
	}
}

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
		decimal.NewFromFloat(1000001),
	}
	for _, input := range inputs {
		fmt.Println(input)
		fmt.Println(BahtToThaiText(input))
	}
}
