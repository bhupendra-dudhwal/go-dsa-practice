package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/bhupendra-dudhwal/go-dsa-practice/problems/maps"
)

func main() {
	// fmt.Println(str.IsPalindrome("abccbaZabccba"))

	var (
		key      string        = "test"
		value    string        = "value"
		timespan time.Duration = 5 * time.Minute
	)
	keydbPorts := maps.NewKeydb()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			keydbPorts.Add(fmt.Sprintf("%s-%d", key, i), i, timespan)

			fmt.Printf("\nAdd: key - %s and success - %t\n", fmt.Sprintf("%s-%d", key, i), keydbPorts.Add(fmt.Sprintf("%s-%d", key, i), fmt.Sprintf("%s-%d", value, i), timespan))

			val, err := keydbPorts.Get(fmt.Sprintf("%s-%d", key, i))
			fmt.Printf("\nGet: key - %s and val - %v, err - %+v\n", fmt.Sprintf("%s-%d", key, i), val, err)

			// fmt.Printf("\nDelete: key - %s and succeess - %t\n", fmt.Sprintf("%s-%d", key, i), keydbPorts.Delete(fmt.Sprintf("%s-%d", key, i)))

			// val, err = keydbPorts.Get(fmt.Sprintf("%s-%d", key, i))
			// fmt.Printf("\nGet After Delete: key - %s and val - %v, err - %+v\n", fmt.Sprintf("%s-%d", key, i), val, err)
		}(i)
	}
	wg.Wait()

	// brackets := "))"
	// fmt.Printf("\n brackets '%s' - %t\n", brackets, str.NewBalancedParentheses().IsBalanced(brackets))
}
