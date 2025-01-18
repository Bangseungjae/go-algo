package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	letterLogs := []string{}
	digitLogs := []string{}

	for _, log := range logs {
		parts := strings.Split(log, " ")
		if unicode.IsNumber(rune(parts[1][0])) {
			digitLogs = append(digitLogs, log)
		} else {
			letterLogs = append(letterLogs, log)
		}
	}

	sort.Slice(letterLogs, func(i, j int) bool {
		s1 := strings.SplitN(letterLogs[i], " ", 2)
		s2 := strings.SplitN(letterLogs[j], " ", 2)

		if s1[1] != s2[1] {
			return s1[1] < s2[1]
		}
		return s1[0] < s2[0]
	})

	letterLogs = append(letterLogs, digitLogs...)

	return letterLogs
}

/**
Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
        [let1 art can let3 art zero let2 own kit dig dig1 8 1 5 1 dig2 3 6]
*/

func main() {
	logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	rs := reorderLogFiles(logs)
	fmt.Println(rs)
}
