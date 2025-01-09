package app_models

type ItemsWeb struct {
	Itmid string `json:"itmid"`

	Description string `json:"description"`
	Price       string `json:"price"`
	Updtime     string `json:"updtime"`
	Serial      string `json:"serial"`

	Typ string `json:"typ"`
	Man string `json:"man"`
	Loc string `json:"loc"`
	Sta string `json:"sta"`

	Uid string `json:"uid"`
}
type ItmSearch struct {
	Itmid  string `json:"itmid"`
	Serial string `json:"serial"`

	Typid string `json:"typid"`
	Manid string `json:"manid"`
	Locid string `json:"locid"`

	Uid string `json:"uid"`
}

type ItmStatHistWeb struct {
	Staname string `json:"staname"`
	Updated string `json:"updated"`
	Uid     string `json:"uid"`
	Comment string `json:"comment"`
}

type DoSearch struct {
	Locid int `json:"locid"`
	Typid int `json:"typid"`
	Manid int `json:"manid"`
	Staid int `json:"staid"`
}

type Stats struct {
	Itm_count   int64   `json:"itm_count"`
	Price       float64 `json:"price"`
	Total_Price string  `json:"total_price"`

	FirstDate string `json:"firstdate"`
	LastDate  string `json:"lastdate"`

	Locs    []int64 `json:"locs"`
	LocType []Loc   `json:"loctype"`
}

type Loc struct {
	Locname         string  `json:"loc_name"`
	ItemsCount      int64   `json:"items_count"`
	ItemsTotalCount int64   `json:"items_totcount"`
	LocPrice        float64 `json:"loc_price"`
	LocTotPrice     string  `json:"loc_totprice"`
	Types           []int64 `json:"types"`
	Loctype         []Typ   `json:"loc_type"`
}

type Typ struct {
	Typname  string `json:"typname"`
	TypCount int64  `json:"typcount"`
}
