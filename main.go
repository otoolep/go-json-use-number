package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
)

func main() {
	type Foo struct {
		Bar string `json:"bar,omitempty"`
		Baz uint64 `json:"baz,omitempty"`
		Qux any    `json:"qux,omitempty"`
		Qaz any    `json:"qaz,omitempty"`
	}

	b1 := `{"bar":"hello","baz":42}`
	b2 := `{"bar":"world","baz":1000000000000000000000000000000000000000000000000000000}`
	b3 := `{"bar":"big","qux":1234567890123456789012345678901234567890}`
	b4 := `{"bar":"float","qaz":389.29}`

	f1 := Foo{}
	f2 := Foo{}

	if err := json.Unmarshal([]byte(b1), &f1); err != nil {
		fmt.Printf("Error unmarshaling b1: %v\n", err)
	}
	if err := json.Unmarshal([]byte(b2), &f2); err != nil {
		fmt.Printf("Error unmarshaling b2: %v\n", err)
	}

	dec := json.NewDecoder(bytes.NewReader([]byte(b3)))
	dec.UseNumber()
	f3 := Foo{}
	if err := dec.Decode(&f3); err != nil {
		fmt.Printf("Error decoding b3 with UseNumber: %v\n", err)
	}

	if _, err := f3.Qux.(json.Number).Int64(); err != nil {
		fmt.Printf("Qux is too large for int64: %v\n", err)
	} else {
		fmt.Println("Qux fits in int64")
	}

	var bi big.Int
	if _, err := bi.SetString(f3.Qux.(json.Number).String(), 10); !err {
		fmt.Println("Failed to convert Qux to big.Int")
	} else {
		fmt.Printf("Qux as big.Int: %s\n", bi.String())
	}

	if bi.IsInt64() {
		fmt.Printf("big.Int fits in int64: %d\n", bi.Int64())
	} else {
		fmt.Println("big.Int does not fit in int64")
	}

	dec2 := json.NewDecoder(bytes.NewReader([]byte(b4)))
	dec2.UseNumber()
	f4 := Foo{}
	if err := dec2.Decode(&f4); err != nil {
		fmt.Printf("Error decoding b4 with UseNumber: %v\n", err)
	}

	var bi2 big.Int
	if _, err := bi2.SetString(f4.Qaz.(json.Number).String(), 10); !err {
		fmt.Println("Failed to convert Qaz to big.Int")
	} else {
		fmt.Printf("Qux as big.Int: %s\n", bi2.String())
	}

	if bi2.IsInt64() {
		fmt.Printf("big.Int bi2 fits in int64: %d\n", bi2.Int64())
	} else {
		fmt.Println("big.Int bi2 does not fit in int64")
	}
}
