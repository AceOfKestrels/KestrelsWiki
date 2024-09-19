package models

type SearchContext struct {
	SearchInContent     bool   `json:"searchInContent"`
	SearchInSubHeadings bool   `json:"searchInSubHeadings"`
	SearchString        string `json:"searchString"`
	CurrentPage         string `json:"currentPage"`
}
