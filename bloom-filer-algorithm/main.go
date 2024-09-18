package main

import (
	"fmt"
	"gobloom/bloom"
	"log"
)

func main() {
	hasher := bloom.NewXXHasher()
	bf, err := bloom.New(1000, 0.01, hasher)
	if err != nil {
		log.Fatalf("Failed to create Bloom filter: %v", err)
	}

	// Add a large number of items
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("item-%d", i)
		bf.Add([]byte(key))
		if i%1000 == 0 {
			fmt.Printf("After adding %d items: %s\n", i+1, bf.Info())
		}
	}

	// Test some items
	fmt.Println(bf.Test([]byte("item-42")))     // true
	fmt.Println(bf.Test([]byte("item-9999")))   // true
	fmt.Println(bf.Test([]byte("item-10000")))  // false (probably)
	fmt.Println(bf.Test([]byte("item-763541"))) // true
	fmt.Println(bf.Test([]byte("akljefgiw")))   // false

	fmt.Printf("Final state: %s\n", bf.Info())
}
