package idmapper


type MappingEntry struct {
	In      string `json:"in"`
	InType  string `json:"inType"`
	Matches map[string]interface{} `json:"matches"`
}


type MappingResult struct {
	Matched   [] MappingEntry `json:"matched"`
	Unmatched []string `json:"unmatched"`
}
