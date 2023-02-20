package main

import "fmt"

func helper(chunks [][]string, chunkIndex int, prev []string, res [][]string) [][]string {
	chunk := chunks[chunkIndex]
	isLast := chunkIndex == len(chunks)-1
	for _, val := range chunk {
		if cur := append(prev, val); isLast {
			res = append(res, cur)
		} else {
			res = helper(chunks, chunkIndex+1, cur, res)
		}
	}
	return res
}
func combine(chunks ...[]string) (res [][]string) {
	return helper(chunks, 0, []string{}, res)
}
func main() {
	names := []string{"iPhone X", "iPhone XS"}
	colors := []string{"黑色", "白色"}
	storage := []string{"64g", "256g"}
	fmt.Println(combine(names, colors, storage))
}
