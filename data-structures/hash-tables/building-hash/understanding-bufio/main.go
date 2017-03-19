package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "The harpoon was darted; the stricken whale flew forward; with igniting velocity the line ran through the grooves;--ran foul.  Ahab stooped to clear it; he did clear it; but the flying turn caught him  round the neck, and voicelessly as Turkish mutes bowstring their victim, he was shot out of the boat, ere the crew knew he was gone. Next instant, the heavy eye-splice in the rope's final end flew out of the stark-empty tub knocked down an oarsman, and smiting the sea, disappeared in its depths."

	// The input don't implement the Reader itself, so calling strings.NewReader returns the string with the Reader
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Count the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
