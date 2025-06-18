package where

// AWS provides direct access to all Amazon Web Services regions.
// Each field corresponds to the actual AWS region code.
var AWS = struct {
	// North America
	USEast1   Code // us-east-1 (N. Virginia)
	USEast2   Code // us-east-2 (Ohio)
	USWest1   Code // us-west-1 (N. California)
	USWest2   Code // us-west-2 (Oregon)
	CACanada1 Code // ca-central-1 (Canada Central)
	CAWest1   Code // ca-west-1 (Canada West)

	// South America
	SAEast1 Code // sa-east-1 (São Paulo)

	// Europe
	EUWest1    Code // eu-west-1 (Ireland)
	EUWest2    Code // eu-west-2 (London)
	EUWest3    Code // eu-west-3 (Paris)
	EUCentral1 Code // eu-central-1 (Frankfurt)
	EUCentral2 Code // eu-central-2 (Zurich)
	EUNorth1   Code // eu-north-1 (Stockholm)
	EUSouth1   Code // eu-south-1 (Milan)
	EUSouth2   Code // eu-south-2 (Spain)

	// Asia Pacific
	APSouth1     Code // ap-south-1 (Mumbai)
	APSouth2     Code // ap-south-2 (Hyderabad)
	APSoutheast1 Code // ap-southeast-1 (Singapore)
	APSoutheast2 Code // ap-southeast-2 (Sydney)
	APSoutheast3 Code // ap-southeast-3 (Jakarta)
	APSoutheast4 Code // ap-southeast-4 (Melbourne)
	APSoutheast5 Code // ap-southeast-5 (Kuala Lumpur)
	APEast1      Code // ap-east-1 (Hong Kong)
	APNortheast1 Code // ap-northeast-1 (Tokyo)
	APNortheast2 Code // ap-northeast-2 (Seoul)
	APNortheast3 Code // ap-northeast-3 (Osaka)

	// Middle East
	MESouth1   Code // me-south-1 (Bahrain)
	MECentral1 Code // me-central-1 (UAE)
	ILCentral1 Code // il-central-1 (Tel Aviv)

	// Africa
	AFSouth1 Code // af-south-1 (Cape Town)

	// China
	CNNorth1     Code // cn-north-1 (Beijing)
	CNNorthwest1 Code // cn-northwest-1 (Ningxia)

	// Government (US)
	USGovEast1 Code // us-gov-east-1
	USGovWest1 Code // us-gov-west-1
}{
	// Initialize with actual region codes
	USEast1:      "us-east-1",
	USEast2:      "us-east-2",
	USWest1:      "us-west-1",
	USWest2:      "us-west-2",
	CACanada1:    "ca-central-1",
	CAWest1:      "ca-west-1",
	SAEast1:      "sa-east-1",
	EUWest1:      "eu-west-1",
	EUWest2:      "eu-west-2",
	EUWest3:      "eu-west-3",
	EUCentral1:   "eu-central-1",
	EUCentral2:   "eu-central-2",
	EUNorth1:     "eu-north-1",
	EUSouth1:     "eu-south-1",
	EUSouth2:     "eu-south-2",
	APSouth1:     "ap-south-1",
	APSouth2:     "ap-south-2",
	APSoutheast1: "ap-southeast-1",
	APSoutheast2: "ap-southeast-2",
	APSoutheast3: "ap-southeast-3",
	APSoutheast4: "ap-southeast-4",
	APSoutheast5: "ap-southeast-5",
	APEast1:      "ap-east-1",
	APNortheast1: "ap-northeast-1",
	APNortheast2: "ap-northeast-2",
	APNortheast3: "ap-northeast-3",
	MESouth1:     "me-south-1",
	MECentral1:   "me-central-1",
	ILCentral1:   "il-central-1",
	AFSouth1:     "af-south-1",
	CNNorth1:     "cn-north-1",
	CNNorthwest1: "cn-northwest-1",
	USGovEast1:   "us-gov-east-1",
	USGovWest1:   "us-gov-west-1",
}

// Azure provides direct access to all Microsoft Azure regions.
var Azure = struct {
	// North America
	EastUS         Code // eastus
	EastUS2        Code // eastus2
	WestUS         Code // westus
	WestUS2        Code // westus2
	WestUS3        Code // westus3
	CentralUS      Code // centralus
	NorthCentralUS Code // northcentralus
	SouthCentralUS Code // southcentralus
	CanadaCentral  Code // canadacentral
	CanadaEast     Code // canadaeast

	// Europe
	NorthEurope        Code // northeurope
	WestEurope         Code // westeurope
	UKSouth            Code // uksouth
	UKWest             Code // ukwest
	FranceCentral      Code // francecentral
	FranceSouth        Code // francesouth
	GermanyNorth       Code // germanynorth
	GermanyWestCentral Code // germanywestcentral
	NorwayEast         Code // norwayeast
	NorwayWest         Code // norwaywest
	SwitzerlandNorth   Code // switzerlandnorth
	SwitzerlandWest    Code // switzerlandwest
	SwedenCentral      Code // swedencentral
	SwedenSouth        Code // swedensouth
	PolandCentral      Code // polandcentral
	SpainCentral       Code // spaincentral
	ItalyNorth         Code // italynorth
	IsraelCentral      Code // israelcentral
	AustriaEast        Code // austriaeast

	// Asia Pacific
	EastAsia           Code // eastasia
	SoutheastAsia      Code // southeastasia
	AustraliaEast      Code // australiaeast
	AustraliaSoutheast Code // australiasoutheast
	AustraliaCentral   Code // australiacentral
	AustraliaCentral2  Code // australiacentral2
	JapanEast          Code // japaneast
	JapanWest          Code // japanwest
	KoreaCentral       Code // koreacentral
	KoreaSouth         Code // koreasouth
	CentralIndia       Code // centralindia
	SouthIndia         Code // southindia
	WestIndia          Code // westindia

	// Middle East & Africa
	UAENorth         Code // uaenorth
	UAECentral       Code // uaecentral
	QatarCentral     Code // qatarcentral
	SouthAfricaNorth Code // southafricanorth
	SouthAfricaWest  Code // southafricawest

	// South America
	BrazilSouth     Code // brazilsouth
	BrazilSoutheast Code // brazilsoutheast
	ChileCentral    Code // chilecentral

	// North America
	MexicoCentral Code // mexicocentral

	// Oceania
	NewZealandNorth Code // newzealandnorth
}{
	EastUS:             "eastus",
	EastUS2:            "eastus2",
	WestUS:             "westus",
	WestUS2:            "westus2",
	WestUS3:            "westus3",
	CentralUS:          "centralus",
	NorthCentralUS:     "northcentralus",
	SouthCentralUS:     "southcentralus",
	CanadaCentral:      "canadacentral",
	CanadaEast:         "canadaeast",
	NorthEurope:        "northeurope",
	WestEurope:         "westeurope",
	UKSouth:            "uksouth",
	UKWest:             "ukwest",
	FranceCentral:      "francecentral",
	FranceSouth:        "francesouth",
	GermanyNorth:       "germanynorth",
	GermanyWestCentral: "germanywestcentral",
	NorwayEast:         "norwayeast",
	NorwayWest:         "norwaywest",
	SwitzerlandNorth:   "switzerlandnorth",
	SwitzerlandWest:    "switzerlandwest",
	SwedenCentral:      "swedencentral",
	SwedenSouth:        "swedensouth",
	PolandCentral:      "polandcentral",
	SpainCentral:       "spaincentral",
	ItalyNorth:         "italynorth",
	IsraelCentral:      "israelcentral",
	AustriaEast:        "austriaeast",
	EastAsia:           "eastasia",
	SoutheastAsia:      "southeastasia",
	AustraliaEast:      "australiaeast",
	AustraliaSoutheast: "australiasoutheast",
	AustraliaCentral:   "australiacentral",
	AustraliaCentral2:  "australiacentral2",
	JapanEast:          "japaneast",
	JapanWest:          "japanwest",
	KoreaCentral:       "koreacentral",
	KoreaSouth:         "koreasouth",
	CentralIndia:       "centralindia",
	SouthIndia:         "southindia",
	WestIndia:          "westindia",
	UAENorth:           "uaenorth",
	UAECentral:         "uaecentral",
	QatarCentral:       "qatarcentral",
	SouthAfricaNorth:   "southafricanorth",
	SouthAfricaWest:    "southafricawest",
	BrazilSouth:        "brazilsouth",
	BrazilSoutheast:    "brazilsoutheast",
	ChileCentral:       "chilecentral",
	MexicoCentral:      "mexicocentral",
	NewZealandNorth:    "newzealandnorth",
}

// GCP provides direct access to all Google Cloud Platform regions.
var GCP = struct {
	// North America
	USCentral1             Code // us-central1 (Iowa)
	USEast1                Code // us-east1 (South Carolina)
	USEast4                Code // us-east4 (Northern Virginia)
	USEast5                Code // us-east5 (Columbus)
	USWest1                Code // us-west1 (Oregon)
	USWest2                Code // us-west2 (Los Angeles)
	USWest3                Code // us-west3 (Salt Lake City)
	USWest4                Code // us-west4 (Las Vegas)
	USSouth1               Code // us-south1 (Texas)
	NorthamericaNortheast1 Code // northamerica-northeast1 (Montreal)
	NorthamericaNortheast2 Code // northamerica-northeast2 (Toronto)

	// South America
	SouthamericaEast1 Code // southamerica-east1 (São Paulo)
	SouthamericaWest1 Code // southamerica-west1 (Santiago)

	// Europe
	EuropeNorth1     Code // europe-north1 (Finland)
	EuropeWest1      Code // europe-west1 (Belgium)
	EuropeWest2      Code // europe-west2 (London)
	EuropeWest3      Code // europe-west3 (Frankfurt)
	EuropeWest4      Code // europe-west4 (Netherlands)
	EuropeWest6      Code // europe-west6 (Zurich)
	EuropeWest8      Code // europe-west8 (Milan)
	EuropeWest9      Code // europe-west9 (Paris)
	EuropeWest10     Code // europe-west10 (Berlin)
	EuropeWest12     Code // europe-west12 (Turin)
	EuropeCentral2   Code // europe-central2 (Warsaw)
	EuropeSouthwest1 Code // europe-southwest1 (Madrid)

	// Asia
	AsiaEast1      Code // asia-east1 (Taiwan)
	AsiaEast2      Code // asia-east2 (Hong Kong)
	AsiaNortheast1 Code // asia-northeast1 (Tokyo)
	AsiaNortheast2 Code // asia-northeast2 (Osaka)
	AsiaNortheast3 Code // asia-northeast3 (Seoul)
	AsiaSouth1     Code // asia-south1 (Mumbai)
	AsiaSouth2     Code // asia-south2 (Delhi)
	AsiaSoutheast1 Code // asia-southeast1 (Singapore)
	AsiaSoutheast2 Code // asia-southeast2 (Jakarta)

	// Australia
	AustraliaSoutheast1 Code // australia-southeast1 (Sydney)
	AustraliaSoutheast2 Code // australia-southeast2 (Melbourne)

	// Middle East
	MECentral1 Code // me-central1 (Doha)
	MECentral2 Code // me-central2 (Dammam)
	MEWest1    Code // me-west1 (Tel Aviv)

	// Africa
	AfricaSouth1 Code // africa-south1 (Johannesburg)
}{
	USCentral1:             "us-central1",
	USEast1:                "us-east1",
	USEast4:                "us-east4",
	USEast5:                "us-east5",
	USWest1:                "us-west1",
	USWest2:                "us-west2",
	USWest3:                "us-west3",
	USWest4:                "us-west4",
	USSouth1:               "us-south1",
	NorthamericaNortheast1: "northamerica-northeast1",
	NorthamericaNortheast2: "northamerica-northeast2",
	SouthamericaEast1:      "southamerica-east1",
	SouthamericaWest1:      "southamerica-west1",
	EuropeNorth1:           "europe-north1",
	EuropeWest1:            "europe-west1",
	EuropeWest2:            "europe-west2",
	EuropeWest3:            "europe-west3",
	EuropeWest4:            "europe-west4",
	EuropeWest6:            "europe-west6",
	EuropeWest8:            "europe-west8",
	EuropeWest9:            "europe-west9",
	EuropeWest10:           "europe-west10",
	EuropeWest12:           "europe-west12",
	EuropeCentral2:         "europe-central2",
	EuropeSouthwest1:       "europe-southwest1",
	AsiaEast1:              "asia-east1",
	AsiaEast2:              "asia-east2",
	AsiaNortheast1:         "asia-northeast1",
	AsiaNortheast2:         "asia-northeast2",
	AsiaNortheast3:         "asia-northeast3",
	AsiaSouth1:             "asia-south1",
	AsiaSouth2:             "asia-south2",
	AsiaSoutheast1:         "asia-southeast1",
	AsiaSoutheast2:         "asia-southeast2",
	AustraliaSoutheast1:    "australia-southeast1",
	AustraliaSoutheast2:    "australia-southeast2",
	MECentral1:             "me-central1",
	MECentral2:             "me-central2",
	MEWest1:                "me-west1",
	AfricaSouth1:           "africa-south1",
}

// Yandex provides direct access to all Yandex Cloud regions.
var Yandex = struct {
	RUCentral1 Code // ru-central1 (Moscow)
	KZ1        Code // kz1 (Kazakhstan)
}{
	RUCentral1: "ru-central1",
	KZ1:        "kz1",
}

// Alibaba provides direct access to all Alibaba Cloud regions.
var Alibaba = struct {
	// China
	CNQingdao     Code // cn-qingdao
	CNBeijing     Code // cn-beijing
	CNZhangjiakou Code // cn-zhangjiakou
	CNHuhehaote   Code // cn-huhehaote
	CNWulanchabu  Code // cn-wulanchabu
	CNHangzhou    Code // cn-hangzhou
	CNShanghai    Code // cn-shanghai
	CNNanjing     Code // cn-nanjing
	CNFuzhou      Code // cn-fuzhou
	CNWuhanLR     Code // cn-wuhan-lr
	CNShenzhen    Code // cn-shenzhen
	CNHeyuan      Code // cn-heyuan
	CNGuangzhou   Code // cn-guangzhou
	CNChengdu     Code // cn-chengdu
	CNHongkong    Code // cn-hongkong

	// Asia Pacific
	APSoutheast1 Code // ap-southeast-1 (Singapore)
	APSoutheast3 Code // ap-southeast-3 (Kuala Lumpur)
	APSoutheast5 Code // ap-southeast-5 (Jakarta)
	APSoutheast6 Code // ap-southeast-6 (Manila)
	APSoutheast7 Code // ap-southeast-7 (Bangkok)
	APSouth1     Code // ap-south-1 (Mumbai)
	APNortheast1 Code // ap-northeast-1 (Tokyo)
	APNortheast2 Code // ap-northeast-2 (Seoul)

	// Europe & Americas
	EUCentral1 Code // eu-central-1 (Frankfurt)
	EUWest1    Code // eu-west-1 (London)
	USEast1    Code // us-east-1 (Virginia)
	USWest1    Code // us-west-1 (Silicon Valley)

	// Middle East & Africa
	MEEast1    Code // me-east-1 (Dubai)
	MECentral1 Code // me-central-1 (Riyadh)

	// North America
	NASouth1 Code // na-south-1 (Mexico)
}{
	CNQingdao:     "cn-qingdao",
	CNBeijing:     "cn-beijing",
	CNZhangjiakou: "cn-zhangjiakou",
	CNHuhehaote:   "cn-huhehaote",
	CNWulanchabu:  "cn-wulanchabu",
	CNHangzhou:    "cn-hangzhou",
	CNShanghai:    "cn-shanghai",
	CNNanjing:     "cn-nanjing",
	CNFuzhou:      "cn-fuzhou",
	CNWuhanLR:     "cn-wuhan-lr",
	CNShenzhen:    "cn-shenzhen",
	CNHeyuan:      "cn-heyuan",
	CNGuangzhou:   "cn-guangzhou",
	CNChengdu:     "cn-chengdu",
	CNHongkong:    "cn-hongkong",
	APSoutheast1:  "ap-southeast-1",
	APSoutheast3:  "ap-southeast-3",
	APSoutheast5:  "ap-southeast-5",
	APSoutheast6:  "ap-southeast-6",
	APSoutheast7:  "ap-southeast-7",
	APSouth1:      "ap-south-1",
	APNortheast1:  "ap-northeast-1",
	APNortheast2:  "ap-northeast-2",
	EUCentral1:    "eu-central-1",
	EUWest1:       "eu-west-1",
	USEast1:       "us-east-1",
	USWest1:       "us-west-1",
	MEEast1:       "me-east-1",
	MECentral1:    "me-central-1",
	NASouth1:      "na-south-1",
}
