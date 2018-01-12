package main

import "math"

func getPaginationPages(pagination Pagination) int {
	pages := float64(pagination.Total / pagination.PerPage)
	return int(math.Ceil(pages))
}
