package Implementation

type CountryData []struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Capital  string `json:"capital"`
	Region   string `json:"region"`
	Currency struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currency"`
	Language struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"language,omitempty"`
	Flag      string `json:"flag"`
	Language0 struct {
		Code       string `json:"code"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"language,omitempty"`
	Demonym   string `json:"demonym,omitempty"`
	Language1 struct {
		Code       string `json:"code"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"language,omitempty"`
	Language2 struct {
		Code       string `json:"code"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"language,omitempty"`
	Language3 struct {
		Code       string `json:"code"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"language,omitempty"`
	Language4 struct {
		Code       string `json:"code"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"language,omitempty"`
	Language5 struct {
		Code       string `json:"code"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"language,omitempty"`
	Language6 struct {
		Code       string `json:"code"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"language,omitempty"`
	Language7 struct {
		Code       string `json:"code"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"language,omitempty"`
	Language8 struct {
		Code       string `json:"code"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"language,omitempty"`
}
