package main

import "time"

func emit(wordChannel chan string, done chan bool) {
	defer close(wordChannel)

	words := []string{"The", "quick", "brown", "fox"}

	i := 0

	t := time.NewTimer(3 * time.Second)

	for {
		select {
		case wordChannel <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}

		case <-done:
			done <- true
			return

		case <-t.C:
			return
		}
	}
}

// func main() {
// 	wordChannel := make(chan string)
// 	doneChannel := make(chan bool)

// 	go emit(wordChannel, doneChannel)

// 	for word := range wordChannel {
// 		fmt.Printf("%s ", word)
// 	}
// }
