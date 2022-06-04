package anagram

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type AnaChkResult struct {
	IsAnagram bool
	Distance  uint16
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CheckAnagram(str1 string, str2 string) AnaChkResult {
	if len(str1) != len(str2) {
		return AnaChkResult{false, 0}
	}

	hash := make(map[string]int)

	distanceHash := make(map[string]int)

	var distance = 0

	for index, r := range str1 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}

		distanceHash[string(r)] = index
	}

	for index, r := range str2 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}

		if distanceHash[string(r)] != index {
			distance++
		}
	}

	var isAnagram bool = true
	for _, value := range hash {
		if value%2 != 0 {
			isAnagram = false
			break
		}
	}

	return AnaChkResult{isAnagram, uint16(distance)}
}

func TwistWord(word string, letters []string) []string {
	if len(letters) == 0 {
		return []string{word}
	} else {
		var twistedWords []string
		for i, l := range letters {
			letters_ := make([]string, len(letters))
			copy(letters_, letters)
			twistedWords_ := TwistWord(word+l, append(letters_[:i], letters_[i+1:]...))
			twistedWords = append(twistedWords, twistedWords_...)
		}
		return twistedWords
	}
}

func FindAnagrams(lang string, word string) ([]string, string) {
	f, err := os.Open(fmt.Sprintf("./word-list/%v.txt", lang))
	if err != nil {
		return nil, fmt.Sprintf("No dictionary for the lang - %v", lang)
	}

	defer f.Close()

	words := strings.Split(word, " ")
	var anagrams []string

	if len(words) > 0 {
		scanner := bufio.NewScanner(f)

		var dictionaryWords []string
		for scanner.Scan() {
			word_ := scanner.Text()
			dictionaryWords = append(dictionaryWords, word_)
		}

		twistedWords := TwistWord("", strings.Split(word, ""))

		for _, twistedWord := range twistedWords {
			var notFound = false
			for _, phrasePiece := range strings.Split(twistedWord, " ") {
				var bFound = false
				for _, dictionaryWord := range dictionaryWords {
					if phrasePiece == dictionaryWord {
						bFound = true
						break
					}
				}
				if !bFound {
					notFound = true
					break
				}
			}
			if !notFound {
				anagrams = append(anagrams, twistedWord)
			}
		}
	} else {
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			word_ := scanner.Text()
			if word == word_ {
				continue
			}

			result := CheckAnagram(word, word_)
			if result.IsAnagram {
				anagrams = append(anagrams, word_)
			}
		}
	}
	return anagrams, ""
}

func FindLongest(lang string, word string) (string, uint16, string) {

	f, err := os.Open(fmt.Sprintf("./word-list/%v.txt", lang))
	if err != nil {
		return "", 0, fmt.Sprintf("No dictionary for the lang - %v", lang)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var anagram = ""
	var maxDistance = 0

	for scanner.Scan() {
		word_ := scanner.Text()
		if word == word_ {
			continue
		}

		result := CheckAnagram(word, word_)
		if result.IsAnagram && maxDistance <= int(result.Distance) {
			anagram = word_
			maxDistance = int(result.Distance)
		}
	}
	return anagram, uint16(maxDistance), ""
}
