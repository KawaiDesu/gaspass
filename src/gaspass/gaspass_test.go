package gaspass

import "fmt"

func ExampleGeneratePassword(){
    p := Params{
        PrivKey:     []byte("asdfghjkl123"),
        Salt:        []byte(""),
        Counter:     []byte("0"),
        PassLength:  16,
        UseLower:    true,
        UseUpper:    true,
        UseNumbers:  true,
        UseSpecials: true,
    }

	password, err := p.GeneratePassword()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*password)
	// Output:
	// Esq8)G5q|NqXqt$T
}

