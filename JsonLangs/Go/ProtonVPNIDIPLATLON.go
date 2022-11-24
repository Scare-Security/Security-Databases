package Implimentation

type ProtonVPNIPStructEncoded struct {
	Code           int `json:"Code"`
	LogicalServers []struct {
		Name         string      `json:"Name"`
		EntryCountry string      `json:"EntryCountry"`
		ExitCountry  string      `json:"ExitCountry"`
		Domain       string      `json:"Domain"`
		Tier         int         `json:"Tier"`
		Features     int         `json:"Features"`
		Region       interface{} `json:"Region"`
		City         string      `json:"City"`
		Score        float64     `json:"Score"`
		HostCountry  interface{} `json:"HostCountry"`
		ID           string      `json:"ID"`
		Location     struct {
			Lat  float64 `json:"Lat"`
			Long float64 `json:"Long"`
		} `json:"Location"`
		Status  int `json:"Status"`
		Servers []struct {
			EntryIP            string      `json:"EntryIP"`
			ExitIP             string      `json:"ExitIP"`
			Domain             string      `json:"Domain"`
			ID                 string      `json:"ID"`
			Label              string      `json:"Label"`
			X25519PublicKey    string      `json:"X25519PublicKey"`
			Generation         int         `json:"Generation"`
			Status             int         `json:"Status"`
			ServicesDown       int         `json:"ServicesDown"`
			ServicesDownReason interface{} `json:"ServicesDownReason"`
		} `json:"Servers"`
		Load int `json:"Load"`
	} `json:"LogicalServers"`
}
