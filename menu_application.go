package main

import (
	"fmt"
	"net/http"
)

func main() {
	allMenus := map[int64]Menu{}
	initApiResp := getApiResponseByPage(1, &allMenus)
	pagesToIterate := getPaginationPages(initApiResp.Pagination)
	for i := 2; i <= pagesToIterate; i++ {
		getApiResponseByPage(i, &allMenus)
	}
	validResults, invalidResults := findCyclicAndNonCyclicMenus(&allMenus)
	w := http.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.Write(
		map[string]interface{}{
			"valid_menus":   validResults,
			"invalid_menus": invalidResults,
		},
	)

	fmt.Printf("%+v \n\n%v\n\n%v", *initApiResp, validResults, invalidResults)
}
