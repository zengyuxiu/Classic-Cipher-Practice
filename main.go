package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var k1 int
	var k2 int
	var re_k2 int
	config, err := ioutil.ReadFile("config")
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	_, _ = fmt.Sscanf(string(config), "%d %d", &k1, &k2)
	re_k := func(k int) int {
		for i := 1; i <= 26; i++ {
			multi := (k * i) % 26
			if multi == 1 {
				return i
			}
		}
		return 0
	}
	gcd := func(k int) int {
		for i := 2; i <= k; i++ {
			if k%i == 0 && 26%i == 0 {
				return i
			}
		}
		return 1
	}

	k1 = (k1%26 + 26) % 26

	if gcd(k2) == 1 {
		re_k2 = re_k(k2)
	} else {
		fmt.Print("k2 illegal!")
		os.Exit(29)
	}

	encode_mapping := func(m rune) rune {
		switch {
		case m >= 'A' && m <= 'Z':
			char := m - 'A'
			m = rune((k1+k2*(int(char)))%26 + int('A'))
			return m
		case m >= 'a' && m <= 'z':
			char := m - 'a'
			m = rune((k1+k2*(int(char)))%26 + int('a'))
			return m
		}
		return m
	}
	decode_mapping := func(c rune) rune {
		switch {
		case c >= 'A' && c <= 'Z':
			char := c - 'A'
			c = rune((re_k2*((int(char))-k1+26))%26 + int('A'))
			return c
		case c >= 'a' && c <= 'z':
			char := c - 'a'
			c = rune((re_k2*((int(char))-k1+26))%26 + int('a'))
			return c
		}
		return c
	}

	plaintext, err := ioutil.ReadFile("plaintext.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	ciphertext := strings.Map(encode_mapping, string(plaintext))
	err = ioutil.WriteFile("ciphertest.txt", []byte(ciphertext), 0644)
	if err != nil {
		panic(err)
	}
	decode_text := strings.Map(decode_mapping, ciphertext)
	err = ioutil.WriteFile("decode_text.txt", []byte(decode_text), 0644)
	if err != nil {
		panic(err)
	}
}
