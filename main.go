brandsMap := map[string]string{
	"ford": "ford",
	"audi": "ford",
	"lada": "nil",
}
_, value := brandsMap["lada"]
fmt.Print(value) // true,
// value выдаст true или false, если ключ lada есть в brandsMap