package rrtype

type RespUPInfo struct {
	Code    int        `json:"code"`
	Data    UpinfoData `json:"data"`
	Message string     `json:"message"`
	TTL     int        `json:"ttl"`
}

type UpinfoData struct {
	EpisodicButton `json:"episodic_button"`
	List           `json:"list"`
	Page           `json:"page"`
}

type EpisodicButton struct {
	Text string `json:"text"`
	URI  string `json:"uri"`
}

type List struct {
	TList []TableELE `json:"tlist"`
	VList []VideoELE `json:"vlist"`
}

type TableELE struct {
	Count int `json:"count"`
	Name  int `json:"name"`
	TID   int `json:"tid"`
}

type VideoELE struct {
	AID          int    `json:"aid"`
	Author       string `json:"author"`
	BVID         string `json:"bvid"`
	Comment      int    `json:"comment"`
	CopyRight    string `json:"copyright"`
	Created      int    `json:"created"`
	Description  string `json:"description"`
	HideClick    bool   `json:"hide_click"`
	ISPAY        int    `json:"is_pay"`
	ISSteinsGate int    `json:"is_steins_gate"`
	ISUnionVideo int    `json:"is_union_video"`
	Length       string `json:"length"`
	MID          int    `json:"mid"`
	PIC          string `json:"pic"`
	Play         int    `json:"play"`
	Review       int    `json:"review"`
	Subtitle     string `json:"subtitle"`
	Title        string `json:"title"`
	Typeid       int    `json:"typeid"`
	VideoReview  int    `json:"review"`
}

type Page struct {
	Count int `json:"count"`
	PN    int `json:"pn"`
	PS    int `json:"ps"`
}
