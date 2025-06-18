package where

// Continent names
const (
	ContinentAsia         = "Asia"
	ContinentEurope       = "Europe"
	ContinentNorthAmerica = "North America"
	ContinentSouthAmerica = "South America"
	ContinentOceania      = "Oceania"
	ContinentAfrica       = "Africa"
)

// Provider names
const (
	ProviderAWS     = "aws"
	ProviderAzure   = "azure"
	ProviderGCP     = "gcp"
	ProviderYandex  = "yandex"
	ProviderVK      = "vk"
	ProviderAlibaba = "alibaba"
)

// InNamespace provides geographic-based region queries.
// Usage: where.In.Asia(), where.In.Europe(), where.In.Country("Japan")
type InNamespace struct{}

// Asia returns all regions in Asia.
func (InNamespace) Asia() Set {
	return InContinent(ContinentAsia)
}

// Europe returns all regions in Europe.
func (InNamespace) Europe() Set {
	return InContinent(ContinentEurope)
}

// Americas returns all regions in North and South America.
func (InNamespace) Americas() Set {
	americas := make(Set, 0)
	americas = append(americas, InContinent(ContinentNorthAmerica)...)
	americas = append(americas, InContinent(ContinentSouthAmerica)...)
	return americas
}

// Oceania returns all regions in Oceania.
func (InNamespace) Oceania() Set {
	return InContinent(ContinentOceania)
}

// Africa returns all regions in Africa.
func (InNamespace) Africa() Set {
	return InContinent(ContinentAfrica)
}

// Country returns all regions in the specified country.
func (InNamespace) Country(name string) Set {
	return InCountry(name)
}

// City returns all regions in the specified city.
func (InNamespace) City(name string) Set {
	return InCity(name)
}

// Continent returns all regions in the specified continent.
func (InNamespace) Continent(name string) Set {
	return InContinent(name)
}

// OnNamespace provides provider-based region queries.
// Usage: where.On.AWS(), where.On.Azure(), where.On.Provider("gcp")
type OnNamespace struct {
}

// AWS returns all AWS regions.
func (OnNamespace) AWS() Set {
	return OnProvider(ProviderAWS)
}

// Azure returns all Azure regions.
func (OnNamespace) Azure() Set {
	return OnProvider(ProviderAzure)
}

// GCP returns all Google Cloud Platform regions.
func (OnNamespace) GCP() Set {
	return OnProvider(ProviderGCP)
}

// Yandex returns all Yandex Cloud regions.
func (OnNamespace) Yandex() Set {
	return OnProvider(ProviderYandex)
}

// VK returns all VK Cloud regions.
func (OnNamespace) VK() Set {
	return OnProvider(ProviderVK)
}

// Alibaba returns all Alibaba Cloud regions.
func (OnNamespace) Alibaba() Set {
	return OnProvider(ProviderAlibaba)
}

// Provider returns all regions from the specified provider.
func (OnNamespace) Provider(name string) Set {
	return OnProvider(name)
}

// IsNamespace provides validation and status-based queries.
// Usage: where.Is.Active(), where.Is.Valid("us-east-1")
type IsNamespace struct{}

// Active returns all active regions.
func (IsNamespace) Active() Set {
	return ActiveRegions()
}

// Preview returns all preview/beta regions.
func (IsNamespace) Preview() Set {
	return PreviewRegions()
}

// Deprecated returns all deprecated regions.
func (IsNamespace) Deprecated() Set {
	return DeprecatedRegions()
}

// Valid returns true if the region code exists.
func (IsNamespace) Valid(code Code) bool {
	return Has(string(code))
}

// Has returns true if the region code exists (alias for Valid).
func (IsNamespace) Has(code Code) bool {
	return Has(string(code))
}

// NearNamespace provides proximity-based region queries.
// Usage: where.Near.Location(lat, lng, radius), where.Near.Region("us-east-1", radius)
type NearNamespace struct{}

// Location returns regions within the specified radius of coordinates.
func (NearNamespace) Location(lat, lng, radiusKm float64) Set {
	return Near(lat, lng, radiusKm)
}

// Region returns regions within the specified radius of another region.
func (NearNamespace) Region(code Code, radiusKm float64) (Set, error) {
	region, err := Is(code)
	if err != nil {
		return nil, err
	}
	return Near(region.Latitude, region.Longitude, radiusKm), nil
}

// City returns regions within the specified radius of a city.
func (NearNamespace) City(cityName string, radiusKm float64) Set {
	// Find regions in the city first, then find regions near those coordinates
	cityRegions := InCity(cityName)
	if len(cityRegions) == 0 {
		return Set{}
	}
	// Use the first city region as reference point
	ref := cityRegions[0]
	return Near(ref.Latitude, ref.Longitude, radiusKm)
}

// Global namespace instances
var (
	// In provides geographic-based region queries.
	In InNamespace

	// On provides provider-based region queries.
	On OnNamespace

	// Validation provides validation and status-based queries.
	Validation IsNamespace

	// Proximity provides proximity-based region queries.
	Proximity NearNamespace
)
