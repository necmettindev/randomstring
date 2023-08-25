package randomstring

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

type Charset string

const (
	Alphanumeric      = Charset("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	Lowercase         = Charset("abcdefghijklmnopqrstuvwxyz")
	Uppercase         = Charset("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Numeric           = Charset("0123456789")
	SpecialCharacters = Charset("!@#$%^&*()-_=+[]{}|;:'<>,.?/~")
)

var (
	GenerateString = generateString
)

type GenerationOptions struct {
	Length                 int
	DisableNumeric         bool
	DisableLowercase       bool
	DisableUppercase       bool
	EnableSpecialCharacter bool
	CustomCharset          Charset
}

func generateStringFromCharset(charset Charset, length int) (string, error) {
	if len(charset) == 0 || length <= 0 {
		return "", errors.New("invalid charset or length")
	}

	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result), nil
}
func modifyCharset(opts GenerationOptions, charsetMappings map[string]Charset, charset Charset) Charset {
	if opts.DisableNumeric {
		charset = Charset(strings.ReplaceAll(string(charset), string(charsetMappings["numeric"]), ""))
	}
	if opts.DisableLowercase {
		charset = Charset(strings.ReplaceAll(string(charset), string(charsetMappings["lowercase"]), ""))
	}
	if opts.DisableUppercase {
		charset = Charset(strings.ReplaceAll(string(charset), string(charsetMappings["uppercase"]), ""))
	}
	if opts.EnableSpecialCharacter {
		charset += charsetMappings["specialCharater"]
	}
	return charset
}

func generateString(opts GenerationOptions) (string, error) {
	charsetMappings := map[string]Charset{
		"numeric":         Numeric,
		"lowercase":       Lowercase,
		"uppercase":       Uppercase,
		"specialCharater": SpecialCharacters,
	}

	charset := Alphanumeric
	if opts.CustomCharset != "" {
		charset = opts.CustomCharset
	} else {
		charset = modifyCharset(opts, charsetMappings, charset)
	}

	if len(charset) == 0 {
		return "", errors.New("resulting charset is empty. adjust your options")
	}

	return generateStringFromCharset(charset, opts.Length)
}
