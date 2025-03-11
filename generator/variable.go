package generator

const (
	charset      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lenOfSymbols = 62
	lenOfNumbs   = 10
)

var (
	symbols = []rune(charset)
	numbs   = []rune("1234567890")
)
