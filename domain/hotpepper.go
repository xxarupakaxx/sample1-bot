package domain

// response APIレスポンス
type Response struct {
	Results Results `json:"results"`
}

// results APIレスポンスの内容
type Results struct {
	Shop []Shop `json:"shop"`
}

// shop レストラン一覧
type Shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Photo   Photo  `json:"photo"`
	URLS    Urls   `json:"urls"`
}

// Photo 写真URL一覧
type Photo struct {
	Mobile Mobile `json:"mobile"`
}

// mobile モバイル用の写真URL
type Mobile struct {
	L string `json:"l"`
}

// urls URL一覧
type Urls struct {
	PC string `json:"pc"`
}

