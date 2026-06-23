# baht-to-text

Converts a decimal currency value to Thai text (บาท/สตางค์).

```
1234       → หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน
33333.75   → สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์
```

## Requirements

- Go 1.22+
- [github.com/shopspring/decimal](https://github.com/shopspring/decimal)

## Getting Started

```bash
go mod tidy
go run main.go
```

## Usage

```go
import (
    "github.com/shopspring/decimal"
)

amount := decimal.NewFromFloat(1234)
text := BahtToThaiText(amount)
// "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"
```

## Function Signature

```go
func BahtToThaiText(amount decimal.Decimal) string
```

| Input | Output |
|-------|--------|
| `0` | `"ศูนย์บาท"` |
| `1234` | `"หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"` |
| `33333.75` | `"สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"` |
| `-100` | `"ลบหนึ่งร้อยบาทถ้วน"` |

**Rules:**
- No fractional part → append `ถ้วน`
- Fractional part → convert to สตางค์ (2 decimal places)
- Negative values → prepend `ลบ`
- Supports values up to hundreds of trillions (หลักแสนล้าน+)

## Running Tests

```bash
go test ./...
```

```bash
go test -v ./...                      # verbose
go test -bench=. -benchmem ./...      # benchmarks
```
