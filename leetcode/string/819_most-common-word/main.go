package main

import (
	"fmt"
	"regexp"
	"strings"
)

// mostCommonWord 함수는 주어진 문단(paragraph)에서 금지된 단어(banned)를 제외한 가장 흔하게 사용된 단어를 반환합니다.
func mostCommonWord(paragraph string, banned []string) string {
	// 금지된 단어들을 집합(set)으로 변환하여 빠른 조회가 가능하도록 함
	bannedSet := make(map[string]struct{})
	for _, word := range banned {
		bannedSet[word] = struct{}{}
	}

	// 정규표현식을 사용하여 모든 비단어 문자를 공백으로 대체하고, 소문자로 변환
	re := regexp.MustCompile(`\W+`)
	processedParagraph := re.ReplaceAllString(paragraph, " ")
	processedParagraph = strings.ToLower(processedParagraph)

	// 공백을 기준으로 단어들을 분리
	words := strings.Fields(processedParagraph)

	// 단어의 빈도를 세기 위한 맵 생성
	wordCount := make(map[string]int)
	for _, word := range words {
		// 금지된 단어가 아니면 카운트 증가
		if _, isBanned := bannedSet[word]; !isBanned {
			wordCount[word]++
		}
	}

	// 가장 빈도가 높은 단어를 찾기
	maxCount := 0
	var mostCommon string
	for word, count := range wordCount {
		if count > maxCount {
			maxCount = count
			mostCommon = word
		}
	}

	return mostCommon
}

func main() {
	// 예제 사용
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}

	result := mostCommonWord(paragraph, banned)
	fmt.Println(result) // 출력: "ball"
}
