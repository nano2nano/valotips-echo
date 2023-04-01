package model

type ImgurResponse struct {
	Data struct {
		// ID          string `json:"id"`
		// Title       any    `json:"title"`
		// Description any    `json:"description"`
		// Datetime    int    `json:"datetime"`
		// Type        string `json:"type"`
		// Animated    bool   `json:"animated"`
		// Width       int    `json:"width"`
		// Height      int    `json:"height"`
		// Size        int    `json:"size"`
		// Views       int    `json:"views"`
		// Bandwidth   int    `json:"bandwidth"`
		// Vote        any    `json:"vote"`
		// Favorite    bool   `json:"favorite"`
		// Nsfw        any    `json:"nsfw"`
		// Section     any    `json:"section"`
		// AccountURL  any    `json:"account_url"`
		// AccountID   int    `json:"account_id"`
		// IsAd        bool   `json:"is_ad"`
		// InMostViral bool   `json:"in_most_viral"`
		// HasSound    bool   `json:"has_sound"`
		// Tags        []any  `json:"tags"`
		// AdType      int    `json:"ad_type"`
		// AdURL       string `json:"ad_url"`
		// Edited      string `json:"edited"`
		// InGallery   bool   `json:"in_gallery"`
		// Deletehash  string `json:"deletehash"`
		// Name        string `json:"name"`
		Link string `json:"link"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}
