module example.com/assessment

go 1.18

replace example.com/anagram => ./anagram

require (
	example.com/anagram v0.0.0-00010101000000-000000000000
	example.com/dictionary v0.0.0-00010101000000-000000000000
)

replace example.com/dictionary => ./word-list
