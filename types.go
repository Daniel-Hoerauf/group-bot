package main

type GroupmePost struct {
	Id 			string 		`json:"bot_id"`
	Text 		string 		`json:"text"`
	Attachment	[]PostImg 	`json:"attachments"`
}

type PostImg struct {
	Typ 	string 		`json:"type"`
	Url 	string		`json:"url"`	
}

type GroupmeInput struct {
	File 	string		`json:"file"`
}

type GroupmeResponse struct {
	Payld	GroupData	`json:"payload"`
}

type GroupData struct {
	Url 	string		`json:"url"`
	Pic		string		`json:"picture_url"`
}

type GiphyResponse struct {
	Data	GiphyData	`json:"data"`
	Meta	GiphyMeta 	`json:"meta"` 
}

type GiphyMeta struct {
	Status		uint		`json:"status"`
	Msg			string 		`json:"msg"`
}

type GiphyData struct {
	Typ			string		`json:"type"`
	Id			string		`json:"id"`
	Slug		string		`json:"slug"`
	Url			string		`json:"url"`
	Bitly_gif	string		`json:"bitly_gif_url"`
	Bitly_url 	string		`json:"bitly_url"`
	Embeded		string		`json:"embeded_url"`
	Username	string		`json:"username"`
	Source		string		`json:"source"`
	Rating		string		`json:"rating"`
	Content		string		`json:"content_url"`
	Source_tld	string 		`json:"source_tld"`
	Source_url	string		`json:"source_post_url"`
	Indexable	uint		`json:"is_indexable,omitempty"`
	Import		string		`json:"import_datetime"`
	Trending	string		`json:"trending_datetime"`
	Images		GiphyImages	`json:"images"`
}

type GiphyImages struct {
	Height		Fixed		`json:"fixed_height"`
	H_still		Still 		`json:"fixed_height_still"`
	H_down		Down 		`json:"fixed_height_downsampled"`
	Width 		Fixed 		`json:"fixed_width"`
	W_still 	Still 		`json:"fixed_width_still"`
	W_down		Down 		`json:"fixed_width_downsampled"`
	H_small		Fixed		`json:"fixed_height_small"`
	Hs_still	Still 		`json:"fixed_height_small_still"`
	W_small 	Fixed 		`json:"fixed_width_small"`
	Ws_still 	Still 		`json:"fixed_width_small_still"`
	Down 		DownSize	`json:"downsized"`
	D_still 	Still 		`json:"downsized_still"`
	D_large 	DownSize 	`json:"downsized_large"`
	D_med 		DownSize 	`json:"downsized_medium"`
	Original 	Fixed 		`json:"original"`
	O_still 	Still 		`json:"original_still"`
	Loop 		Looping 	`json:"looping"`

}

type DownSize struct {
	Url 		string		`json:"url"`
	Width		string		`json:"width"`
	Height 		string		`json:"height"`
	Size 		string		`json:"size"`
}

type Looping struct {
	Mp4 		string		`json:"mp4"`
}

type Fixed struct {
	Url 		string		`json:"url"`
	Width		string		`json:"width"`
	Height 		string		`json:"height"`
	Size 		string		`json:"size"`
	Frames 		string 		`json:"frames,omitempty"`
	Mp4			string		`json:"mp4"`
	Mp4_size	string		`json:"mp4_size"`
	Webp		string 		`json:"webp"`
	Webp_size	string		`json:"webp_size"`
}

type Still struct {
	Url 		string		`json:"url"`
	Width		string		`json:"width"`
	Height 		string		`json:"height"`
}

type Down struct {
	Url 		string		`json:"url"`
	Width		string		`json:"width"`
	Height 		string		`json:"height"`
	Size 		string		`json:"size"`
	Webp		string 		`json:"webp"`
	Webp_size	string		`json:"webp_size"`
}