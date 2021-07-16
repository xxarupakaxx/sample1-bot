package domain


type Weather struct {
	PublicTimeFormatted string    `json:"publicTimeFormatted"`
	Title               string    `json:"title"`
	Link                string    `json:"link"`
	Description			Description `json:"description"`
	Forecasts 			[]Forecasts `json:"forecasts"`
}
type Description struct {
	Text                string    `json:"text"`
}
type Forecasts struct {
	Date      string `json:"date"`
	DateLabel string `json:"dateLabel"`
	Telop     string `json:"telop"`
	Detail    Detail `json:"detail"`
	Temperature Temperature `json:"temperature"`
	ChanceOfRain ChanceOfRain `json:"chanceOfRain"`
}
type Detail struct {
	Weather string `json:"weather"`
	Wind    string `json:"wind"`
	Wave    string `json:"wave"`
}
type Temperature struct{
	Max Max `json:"max"`
	Min Min `json:"min"`
}
type Min struct {
		Celsius    string `json:"celsius"`
}
type Max struct {
	Celsius    string `json:"celsius"`
}
type ChanceOfRain struct {
	T0006 string `json:"T00_06"`
	T0612 string `json:"T06_12"`
	T1218 string `json:"T12_18"`
	T1824 string `json:"T18_24"`
}
