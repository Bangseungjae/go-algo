package main

func main() {

}

func permute(nums []int) [][]int {
	var results [][]int
	var dfs func(prev []int, element []int)
	dfs = func(prev []int, element []int) {
		if len(element) == 0 {
			combination := make([]int, len(prev))
			copy(combination, prev)
			results = append(results, prev)
			return
		}

		for i, e := range element {
			nextElement := append([]int{}, element[:i]...)
			nextElement = append(nextElement, element[i+1:]...)
			newPrev := append(prev, e)
			dfs(newPrev, nextElement)
		}
	}

	dfs([]int{}, nums)
	return results
}
