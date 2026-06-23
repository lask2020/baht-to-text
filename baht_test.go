package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestBahtToThaiText(t *testing.T) {
	tests := []struct {
		name     string
		input    decimal.Decimal
		expected string
	}{

		{
			name:     "1234",
			input:    decimal.NewFromFloat(1234),
			expected: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน",
		},
		{
			name:     "33333.75",
			input:    decimal.NewFromFloat(33333.75),
			expected: "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์",
		},
		{
			name:     "0 - ศูนย์บาท",
			input:    decimal.NewFromFloat(0),
			expected: "ศูนย์บาท",
		},
		{
			name:     "1 - หนึ่งบาทถ้วน",
			input:    decimal.NewFromFloat(1),
			expected: "หนึ่งบาทถ้วน",
		},
		{
			name:     "10 - สิบบาทถ้วน",
			input:    decimal.NewFromFloat(10),
			expected: "สิบบาทถ้วน",
		},
		{
			name:     "11 - สิบเอ็ดบาทถ้วน",
			input:    decimal.NewFromFloat(11),
			expected: "สิบเอ็ดบาทถ้วน",
		},
		{
			name:     "20 - ยี่สิบบาทถ้วน",
			input:    decimal.NewFromFloat(20),
			expected: "ยี่สิบบาทถ้วน",
		},
		{
			name:     "21 - ยี่สิบเอ็ดบาทถ้วน",
			input:    decimal.NewFromFloat(21),
			expected: "ยี่สิบเอ็ดบาทถ้วน",
		},
		{
			name:     "100 - หนึ่งร้อยบาทถ้วน",
			input:    decimal.NewFromFloat(100),
			expected: "หนึ่งร้อยบาทถ้วน",
		},
		{
			name:     "1000 - หนึ่งพันบาทถ้วน",
			input:    decimal.NewFromFloat(1000),
			expected: "หนึ่งพันบาทถ้วน",
		},
		{
			name:     "10000 - หนึ่งหมื่นบาทถ้วน",
			input:    decimal.NewFromFloat(10000),
			expected: "หนึ่งหมื่นบาทถ้วน",
		},
		{
			name:     "100000 - หนึ่งแสนบาทถ้วน",
			input:    decimal.NewFromFloat(100000),
			expected: "หนึ่งแสนบาทถ้วน",
		},
		{
			name:     "1000000 - หนึ่งล้านบาทถ้วน",
			input:    decimal.NewFromFloat(1000000),
			expected: "หนึ่งล้านบาทถ้วน",
		},
		{
			name:     "0.25 - ไม่มีส่วนบาท มีเพียงสตางค์",
			input:    decimal.NewFromFloat(0.25),
			expected: "ยี่สิบห้าสตางค์",
		},
		{
			name:     "0.01 - หนึ่งสตางค์",
			input:    decimal.NewFromFloat(0.01),
			expected: "หนึ่งสตางค์",
		},
		{
			name:     "0.50 - ห้าสิบสตางค์",
			input:    decimal.NewFromFloat(0.50),
			expected: "ห้าสิบสตางค์",
		},
		{
			name:     "1.50 - หนึ่งบาทห้าสิบสตางค์",
			input:    decimal.NewFromFloat(1.50),
			expected: "หนึ่งบาทห้าสิบสตางค์",
		},
		{
			name:     "123.50 - จาก function doc",
			input:    decimal.NewFromFloat(123.50),
			expected: "หนึ่งร้อยยี่สิบสามบาทห้าสิบสตางค์",
		},
		{
			name:     "-100 - ลบหนึ่งร้อยบาทถ้วน",
			input:    decimal.NewFromFloat(-100),
			expected: "ลบหนึ่งร้อยบาทถ้วน",
		},
		{
			name:     "-50.25 - ลบห้าสิบบาทยี่สิบห้าสตางค์",
			input:    decimal.NewFromFloat(-50.25),
			expected: "ลบห้าสิบบาทยี่สิบห้าสตางค์",
		},
		{
			name:     "1000000.01 - หนึ่งล้านบาทหนึ่งสตางค์",
			input:    decimal.NewFromFloat(1000000.01),
			expected: "หนึ่งล้านบาทหนึ่งสตางค์",
		},
		{
			name:     "9999999.99 - เก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทเก้าสิบเก้าสตางค์",
			input:    decimal.NewFromFloat(9999999.99),
			expected: "เก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทเก้าสิบเก้าสตางค์",
		},
		{
			name:     "1000000000 - หนึ่งพันล้านบาทถ้วน",
			input:    decimal.NewFromInt(1000000000),
			expected: "หนึ่งพันล้านบาทถ้วน",
		},
		{
			name:     "1000000000.50 - หนึ่งพันล้านบาทห้าสิบสตางค์",
			input:    decimal.NewFromFloat(1000000000.50),
			expected: "หนึ่งพันล้านบาทห้าสิบสตางค์",
		},
		{
			name:     "1000000000000 - หนึ่งล้านล้านบาทถ้วน",
			input:    decimal.NewFromInt(1000000000000),
			expected: "หนึ่งล้านล้านบาทถ้วน",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BahtToThaiText(tt.input)
			if got != tt.expected {
				t.Errorf("BahtToThaiText(%v)\n  got:  %q\n  want: %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestNumberToThai(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, ""},
		{1, "หนึ่ง"},
		{9, "เก้า"},
		{10, "สิบ"},
		{11, "สิบเอ็ด"},
		{20, "ยี่สิบ"},
		{21, "ยี่สิบเอ็ด"},
		{25, "ยี่สิบห้า"},
		{100, "หนึ่งร้อย"},
		{101, "หนึ่งร้อยเอ็ด"},
		{110, "หนึ่งร้อยสิบ"},
		{111, "หนึ่งร้อยสิบเอ็ด"},
		{1000, "หนึ่งพัน"},
		{10000, "หนึ่งหมื่น"},
		{100000, "หนึ่งแสน"},
		{1000000, "หนึ่งล้าน"},
		{1000001, "หนึ่งล้านหนึ่ง"},
		{1234567, "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ด"},
		{10000000, "สิบล้าน"},
		{100000000, "หนึ่งร้อยล้าน"},
		{1000000000, "หนึ่งพันล้าน"},
		{10000000000, "หนึ่งหมื่นล้าน"},
		{100000000000, "หนึ่งแสนล้าน"},
		{1000000000000, "หนึ่งล้านล้าน"},
		{1000001000000, "หนึ่งล้านล้านหนึ่งล้าน"},
		{9999999999999, "เก้าล้านล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้า"},
	}

	for _, tt := range tests {
		got := numberToThai(tt.input)
		if got != tt.expected {
			t.Errorf("numberToThai(%d)\n  got:  %q\n  want: %q", tt.input, got, tt.expected)
		}
	}
}

var benchCases = []decimal.Decimal{
	decimal.NewFromFloat(1),
	decimal.NewFromFloat(1234),
	decimal.NewFromFloat(33333.75),
	decimal.NewFromFloat(9999999.99),
	decimal.NewFromFloat(1000000.01),
}

var benchInts = []int64{1, 1234, 33333, 9999999, 1000000}

func BenchmarkBahtToThaiText(b *testing.B) {
	for b.Loop() {
		for _, v := range benchCases {
			BahtToThaiText(v)
		}
	}
}

func BenchmarkNumberToThai(b *testing.B) {
	for b.Loop() {
		for _, v := range benchInts {
			numberToThai(v)
		}
	}
}
