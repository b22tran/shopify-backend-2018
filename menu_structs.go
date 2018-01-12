package main

type MerchantApi struct {
	Menu       []Menu     `json:"menus"`
	Pagination Pagination `json:"pagination"`
}

type Menu struct {
	ID       int64   `json:"id"`
	Data     string  `json:"data"`
	ParentID int64   `json:"parent_id,omitempty"`
	ChildIDs []int64 `json:"child_ids"`
}

type Pagination struct {
	CurrentPage int64 `json:"current_page"`
	PerPage     int64 `json:"per_page"`
	Total       int64 `json:"total"`
}

type MenuResult struct {
	RootID   int64   `json:"root_id"`
	Children []int64 `json:"children"`
}
