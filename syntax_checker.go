// Mattia Callegari, 2024 - spell_check.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	_ "embed"

	"github.com/bits-and-blooms/bloom/v3"
)

//go:embed words_alpha.txt
var data string

// bloom filter (https://en.wikipedia.org/wiki/Bloom_filter)
// construct a filter with 5 hashing functions on 2^30 bits
// which has a fprate of ~0.001
var filter = bloom.New(1, 1)

type fn func(string) uint32

func main() {
	time, num := benchmark(load_dict1)
	fmt.Printf("#1: Loaded %v words in %v\n", num, time)
	filter.BitSet().ClearAll()

	time, num = benchmark(load_dict2)
	fmt.Printf("#2: Loaded %v words in %v\n", num, time)

	cli()
}

func benchmark(f fn) (elapsed time.Duration, num uint32) {
	start := time.Now()
	num = f(data)
	elapsed = time.Since(start)
	return
}

func load_dict1(data string) (n uint32) {
	start := 0
	for i, v := range data {
		// '\n' == 10 in ASCII
		if v == 10 {
			filter.AddString(data[start : i-1])
			start = i + 1
			n += 1
		}
	}
	return
}

func load_dict2(data string) (n uint32) {
	return
}

func cli() {
	// Check words inserted in Stdin, exit on ""
	sc := bufio.NewScanner(os.Stdin)
	for word := read_word(*sc); word != ""; word = read_word(*sc) {
		check_word(word)
		sc.Scan()
		clear()
	}
}

func read_word(sc bufio.Scanner) string {
	fmt.Println("Enter a word: ")
	sc.Scan()
	return sc.Text()
}

func check_word(word string) {
	if filter.TestString(word) {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT A WORD")
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
