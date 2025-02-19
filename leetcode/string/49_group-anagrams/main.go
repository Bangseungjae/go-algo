package main

import "slices"

func main() {

}

func groupAnagrams(strs []string) [][]string {
	results := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		s := string(runes)
		if _, ok := results[s]; !ok {
			results[s] = []string{}
		}
		results[s] = append(results[s], str)
	}

	anagramGroups := make([][]string, 0, len(results))
	for _, group := range results {
		anagramGroups = append(anagramGroups, group)
	}
	return anagramGroups
}
