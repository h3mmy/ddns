package models

type IfconfigJSONResponseCo struct {
	IP         string     `json:"ip"`
	IPDecimal  interface{}      `json:"ip_decimal,omitempty"`
	Country    string     `json:"country,omitempty"`
	CountryISO string     `json:"country_iso,omitempty"`
	CountryEU  bool       `json:"country_eu,omitempty"`
	RegionName string     `json:"region_name,omitempty"`
	RegionCode string     `json:"region_code,omitempty"`
	MetroCode  int        `json:"metro_code,omitempty"`
	ZipCode    string     `json:"zip_code,omitempty"`
	City       string     `json:"city,omitempty"`
	Latitude   float32    `json:"latitude,omitempty"`
	Longitude  float32    `json:"longitude,omitempty"`
	Timezone   string     `json:"time_zone,omitempty"`
	ASN        string     `json:"asn,omitempty"`
	ASNOrg     string     `json:"asn_org,omitempty"`
	UserAgent  *UserAgent `json:"user_agent,omitempty"`
}

type IfconfigJSONResponseMe struct {

		IP string `json:"ip_addr"`
		RemoteHost string `json:"remote_host,omitempty"`
		UserAgent string `json:"user_agent,omitempty"`
		Port int `json:"port,omitempty"`
		Method string `json:"method,omitempty"`
		Mime string `json:"mime,omitempty"`
		Via string `json:"via,omitempty"`
		Forwarded string `json:"forwarded,omitempty"`
}

type UserAgent struct {
	Product  string
	Version  string
	Comment  string
	RawValue string
}
