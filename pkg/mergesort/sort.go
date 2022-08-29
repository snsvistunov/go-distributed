package mergesort

import (
	"fmt"
	"sort"
	"sync"
)

func ParralelSort(s []string) {
	len := len(s)

	if len > 1 {
		quarter := len / 4
		middle := len / 2

		var wg sync.WaitGroup
		wg.Add(4)

		go func() {
			defer wg.Done()
			ParralelSort(s[:quarter])
		}()

		go func() {
			defer wg.Done()
			ParralelSort(s[quarter:middle])
		}()

		go func() {
			defer wg.Done()
			ParralelSort(s[middle : middle+quarter])
		}()

		go func() {
			defer wg.Done()
			ParralelSort(s[middle+quarter:])
		}()

		wg.Wait()

		fmt.Printf("Sorted slice %v\n", s)
		sort.Strings(s)
	}

}
