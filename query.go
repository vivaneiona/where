package where

// Query provides a fluent builder pattern for complex region queries.
// Usage:
//
//	results := where.NewQuery().
//	  InCountry("Japan").
//	  ByProvider("aws").
//	  ActiveOnly().
//	  Exec()
type Query struct {
	regions Set
	errors  []error
}

// NewQuery creates a new query builder starting with all regions.
func NewQuery() *Query {
	return &Query{
		regions: allRegions(),
		errors:  make([]error, 0),
	}
}

// InCountry filters regions by country name.
func (q *Query) InCountry(name string) *Query {
	q.regions = q.regions.ByCountry(name)
	return q
}

// InCity filters regions by city name.
func (q *Query) InCity(name string) *Query {
	q.regions = q.regions.ByCity(name)
	return q
}

// InContinent filters regions by continent name.
func (q *Query) InContinent(name string) *Query {
	q.regions = q.regions.ByContinent(name)
	return q
}

// InAsia filters to only Asian regions.
func (q *Query) InAsia() *Query {
	return q.InContinent("Asia")
}

// InEurope filters to only European regions.
func (q *Query) InEurope() *Query {
	return q.InContinent("Europe")
}

// InAmericas filters to only American regions (North + South America).
func (q *Query) InAmericas() *Query {
	americas := make(Set, 0)
	for _, region := range q.regions {
		if region.Continent == "North America" || region.Continent == "South America" {
			americas = append(americas, region)
		}
	}
	q.regions = americas
	return q
}

// InOceania filters to only Oceania regions.
func (q *Query) InOceania() *Query {
	return q.InContinent("Oceania")
}

// InAfrica filters to only African regions.
func (q *Query) InAfrica() *Query {
	return q.InContinent("Africa")
}

// ByProvider filters regions by cloud provider name.
func (q *Query) ByProvider(name string) *Query {
	q.regions = q.regions.ByProvider(name)
	return q
}

// ByAWS filters to only AWS regions.
func (q *Query) ByAWS() *Query {
	return q.ByProvider("aws")
}

// ByAzure filters to only Azure regions.
func (q *Query) ByAzure() *Query {
	return q.ByProvider("azure")
}

// ByGCP filters to only Google Cloud Platform regions.
func (q *Query) ByGCP() *Query {
	return q.ByProvider("gcp")
}

// ByYandex filters to only Yandex Cloud regions.
func (q *Query) ByYandex() *Query {
	return q.ByProvider("yandex")
}

// ByVK filters to only VK Cloud regions.
func (q *Query) ByVK() *Query {
	return q.ByProvider("vk")
}

// ByAlibaba filters to only Alibaba Cloud regions.
func (q *Query) ByAlibaba() *Query {
	return q.ByProvider("alibaba")
}

// ActiveOnly filters to only active regions.
func (q *Query) ActiveOnly() *Query {
	q.regions = q.regions.ActiveOnly()
	return q
}

// PreviewOnly filters to only preview/beta regions.
func (q *Query) PreviewOnly() *Query {
	q.regions = q.regions.Filter(func(r Region) bool {
		return r.Status == Preview
	})
	return q
}

// DeprecatedOnly filters to only deprecated regions.
func (q *Query) DeprecatedOnly() *Query {
	q.regions = q.regions.Filter(func(r Region) bool {
		return r.Status == Deprecated
	})
	return q
}

// Near filters regions within the specified radius of a location.
func (q *Query) Near(lat, lng float64, radiusKm float64) *Query {
	q.regions = q.regions.Near(lat, lng, radiusKm)
	return q
}

// NearRegion filters regions within the specified radius of another region.
func (q *Query) NearRegion(code Code, radiusKm float64) *Query {
	region, err := Is(code)
	if err != nil {
		q.errors = append(q.errors, err)
		return q
	}
	return q.Near(region.Latitude, region.Longitude, radiusKm)
}

// NearCity filters regions within the specified radius of a city.
func (q *Query) NearCity(cityName string, radiusKm float64) *Query {
	// Find regions in the city first
	cityRegions := allRegions().ByCity(cityName)
	if len(cityRegions) == 0 {
		q.regions = Set{} // No regions found in city
		return q
	}
	// Use the first city region as reference point
	ref := cityRegions[0]
	return q.Near(ref.Latitude, ref.Longitude, radiusKm)
}

// Filter applies a custom predicate function.
func (q *Query) Filter(predicate func(Region) bool) *Query {
	q.regions = q.regions.Filter(predicate)
	return q
}

// SortByDistance sorts regions by distance from a location (closest first).
func (q *Query) SortByDistance(lat, lng float64) *Query {
	q.regions.SortByDistance(lat, lng)
	return q
}

// SortByName sorts regions alphabetically by name.
func (q *Query) SortByName() *Query {
	q.regions.SortByName()
	return q
}

// SortByProvider sorts regions by provider name.
func (q *Query) SortByProvider() *Query {
	q.regions.SortByProvider()
	return q
}

// SortByCountry sorts regions by country name.
func (q *Query) SortByCountry() *Query {
	q.regions.SortByCountry()
	return q
}

// Limit restricts the result set to the first N regions.
func (q *Query) Limit(n int) *Query {
	if n < len(q.regions) {
		q.regions = q.regions[:n]
	}
	return q
}

// First returns the first region in the result set.
func (q *Query) First() (Region, error) {
	if len(q.errors) > 0 {
		return Region{}, q.errors[0]
	}
	return q.regions.First()
}

// Last returns the last region in the result set.
func (q *Query) Last() (Region, error) {
	if len(q.errors) > 0 {
		return Region{}, q.errors[0]
	}
	return q.regions.Last()
}

// Count returns the number of regions in the result set.
func (q *Query) Count() int {
	return len(q.regions)
}

// Has returns true if any regions match the query.
func (q *Query) Has() bool {
	return len(q.regions) > 0
}

// Exec executes the query and returns the result set.
func (q *Query) Exec() Set {
	return q.regions
}

// ExecWithErrors executes the query and returns both results and any errors.
func (q *Query) ExecWithErrors() (Set, []error) {
	return q.regions, q.errors
}

// Codes returns just the region codes from the result set.
func (q *Query) Codes() []Code {
	codes := make([]Code, len(q.regions))
	for i, region := range q.regions {
		codes[i] = region.Code
	}
	return codes
}

// Names returns just the region names from the result set.
func (q *Query) Names() []string {
	names := make([]string, len(q.regions))
	for i, region := range q.regions {
		names[i] = region.Name
	}
	return names
}

// Providers returns unique provider names from the result set.
func (q *Query) Providers() []string {
	providerMap := make(map[string]bool)
	for _, region := range q.regions {
		providerMap[region.Provider] = true
	}

	providers := make([]string, 0, len(providerMap))
	for provider := range providerMap {
		providers = append(providers, provider)
	}
	return providers
}

// Countries returns unique country names from the result set.
func (q *Query) Countries() []string {
	countryMap := make(map[string]bool)
	for _, region := range q.regions {
		countryMap[region.Country] = true
	}

	countries := make([]string, 0, len(countryMap))
	for country := range countryMap {
		countries = append(countries, country)
	}
	return countries
}

// Cities returns unique city names from the result set.
func (q *Query) Cities() []string {
	cityMap := make(map[string]bool)
	for _, region := range q.regions {
		cityMap[region.City] = true
	}

	cities := make([]string, 0, len(cityMap))
	for city := range cityMap {
		cities = append(cities, city)
	}
	return cities
}
