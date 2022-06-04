package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func AddDictionary(lang string, words []string) {
	f, err := os.OpenFile(fmt.Sprintf("./word-list/%v.txt", lang), os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		f_, err := os.Create(fmt.Sprintf("./word-list/%v.txt", lang))
		check(err)
		f = f_
	}

	scanner := bufio.NewScanner(f)

	hash := make(map[string]int)

	for scanner.Scan() {
		word := scanner.Text()
		hash[string(word)] = 1
	}

	for _, word := range words {
		if hash[word] == 0 {
			_, err := f.WriteString(fmt.Sprintf("%v\n", word))
			check(err)
		}
	}

	f.Close()
}

func RemoveDictionary(lang string, words []string) string {
	f, err := os.Open(fmt.Sprintf("./word-list/%v.txt", lang))
	if err != nil {
		f.Close()
		return fmt.Sprintf("No dictionary for the lang - %v", lang)
	}

	scanner := bufio.NewScanner(f)

	hash := make(map[string]int)
	for _, word := range words {
		hash[word] = 1
	}

	var newWords []string

	for scanner.Scan() {
		word := scanner.Text()
		if hash[string(word)] == 0 {
			newWords = append(newWords, word)
		}
	}

	err = os.Truncate(fmt.Sprintf("./word-list/%v.txt", lang), 0)
	check(err)

	f.Close()

	f, err = os.OpenFile(fmt.Sprintf("./word-list/%v.txt", lang), os.O_RDWR, 0660)
	check(err)

	_, err = f.WriteString(strings.Join(newWords, "\n"))
	check(err)

	f.Close()
	return ""
}
