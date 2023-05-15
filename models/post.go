package models

type (
	StoragePost struct {
		ID       uint64 `json:"id"`
		Text     string `json:"text"`
		Images   string `json:"images"`
		IsActive bool   `json:"is_active"`
	}

	ResponsePost struct {
		ID       uint64   `json:"id"`
		Text     string   `json:"text"`
		Images   []string `json:"images"`
		IsActive bool     `json:"is_active"`
	}
)
