package main

import (
	"fmt"
)

func main() {
	allMenus := map[int64]Menu{}
	initApiResp := getApiResponseByPage(1, &allMenus)
	pagesToIterate := getPaginationPages(initApiResp.Pagination)
	for i := 2; i <= pagesToIterate; i++ {
		getApiResponseByPage(i, &allMenus)
	}
	validResults, invalidResults := findCyclicAndNonCyclicMenus(&allMenus)

	fmt.Printf("%+v \n\n%v\n\n%v", *initApiResp, validResults, invalidResults)
}
