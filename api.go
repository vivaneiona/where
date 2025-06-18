package where

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	// ErrRegionNotFound is returned when a region code is not recognized.
	ErrRegionNotFound = errors.New("region not found")
	// ErrProviderNotFound is returned when a provider name is not recognized.
	ErrProviderNotFound = errors.New("provider not found")
)

// Question-style API functions that read like natural English

// Is answers "where is {code}?" - returns a query that can be filtered by provider.
// Usage: Is("us-east-1").OnAWS() or Is("us-east-1").First()
func Is(code Code) RegionQuery {
	regions, exists := regionRegistry[code]
	if !exists {
		return RegionQuery{regions: []Region{}}
	}
	return RegionQuery{regions: regions}
}

// MustIs is like Is().First() but panics on error. Use when you're certain the region exists.
func MustIs(code Code) Region {
	query := Is(code)
	region, err := query.First()
	if err != nil {
		panic(err)
	}
	return region
}

// Are answers "where are {codes}?" - returns information about multiple regions.
func Are(codes ...Code) (Set, error) {
	regions := make(Set, 0, len(codes))
	var notFound []Code

	for _, code := range codes {
		if regionList, exists := regionRegistry[code]; exists {
			// Add all regions for this code to the result
			regions = append(regions, regionList...)
		} else {
			notFound = append(notFound, code)
		}
	}

	if len(notFound) > 0 {
		return regions, fmt.Errorf("%w: %v", ErrRegionNotFound, notFound)
	}

	return regions, nil
}

// InCountry answers "where in country {name}?" - returns all regions in a country.
func InCountry(name string) Set {
	return allRegions().ByCountry(name)
}

// InCity answers "where in city {name}?" - returns all regions in a city.
func InCity(name string) Set {
	return allRegions().ByCity(name)
}

// InContinent answers "where in continent {name}?" - returns all regions in a continent.
func InContinent(name string) Set {
	return allRegions().ByContinent(name)
}

// OnProvider answers "where by provider {name}?" - returns all regions from a provider.
func OnProvider(name string) Set {
	return allRegions().OnProvider(name)
}

// Near answers "where near {location} within {radius}km?" - requires coordinates.
func Near(lat, lng float64, radiusKm float64) Set {
	return allRegions().Near(lat, lng, radiusKm)
}

// ActiveRegions answers "where active?" - returns all currently active regions.
func ActiveRegions() Set {
	return allRegions().ActiveOnly()
}

// PreviewRegions answers "where preview?" - returns all preview/beta regions.
func PreviewRegions() Set {
	return allRegions().Filter(func(r Region) bool {
		return r.Status == Preview
	})
}

// DeprecatedRegions answers "where deprecated?" - returns all deprecated regions.
func DeprecatedRegions() Set {
	return allRegions().Filter(func(r Region) bool {
		return r.Status == Deprecated
	})
}

// Validation functions with simple yes/no answers

// Has answers "where valid {code}?" - checks if a region code exists.
func Has(code string) bool {
	_, exists := regionRegistry[Code(code)]
	return exists
}

// IsActive answers "where active {code}?" - true if region is currently active.
func IsActive(code Code) bool {
	if regionList, exists := regionRegistry[code]; exists {
		// Return true if any region with this code is active
		for _, region := range regionList {
			if region.IsActive() {
				return true
			}
		}
	}
	return false
}

// HasProvider answers "where has provider {name}?" - checks if a provider exists.
func HasProvider(name string) bool {
	for _, regionList := range regionRegistry {
		for _, region := range regionList {
			if strings.EqualFold(region.Provider, name) {
				return true
			}
		}
	}
	return false
}

// Discovery functions for exploration

// Providers answers "where providers?" - returns all unique provider names.
func Providers() []string {
	providerMap := make(map[string]bool)
	for _, regionList := range regionRegistry {
		for _, region := range regionList {
			providerMap[region.Provider] = true
		}
	}

	providers := make([]string, 0, len(providerMap))
	for provider := range providerMap {
		providers = append(providers, provider)
	}
	return providers
}

// Countries answers "where countries?" - returns all unique country names.
func Countries() []string {
	countryMap := make(map[string]bool)
	for _, regionList := range regionRegistry {
		for _, region := range regionList {
			countryMap[region.Country] = true
		}
	}

	countries := make([]string, 0, len(countryMap))
	for country := range countryMap {
		countries = append(countries, country)
	}
	return countries
}

// Cities answers "where cities?" - returns all unique city names.
func Cities() []string {
	cityMap := make(map[string]bool)
	for _, regionList := range regionRegistry {
		for _, region := range regionList {
			cityMap[region.City] = true
		}
	}

	cities := make([]string, 0, len(cityMap))
	for city := range cityMap {
		cities = append(cities, city)
	}
	return cities
}

// Continents answers "where continents?" - returns all unique continent names.
func Continents() []string {
	continentMap := make(map[string]bool)
	for _, regionList := range regionRegistry {
		for _, region := range regionList {
			continentMap[region.Continent] = true
		}
	}

	continents := make([]string, 0, len(continentMap))
	for continent := range continentMap {
		continents = append(continents, continent)
	}
	return continents
}

// Distance calculates the distance between two regions in kilometers.
// If multiple regions exist for a code, uses the first region.
func Distance(from, to Code) (float64, error) {
	fromQuery := Is(from)
	fromRegion, err := fromQuery.First()
	if err != nil {
		return 0, fmt.Errorf("source region %w", err)
	}

	toQuery := Is(to)
	toRegion, err := toQuery.First()
	if err != nil {
		return 0, fmt.Errorf("destination region %w", err)
	}

	return fromRegion.Distance(toRegion), nil
}

// Closest finds the closest region to the specified region.
// If multiple regions exist for a code, uses the first region.
func Closest(to Code) (Region, error) {
	targetQuery := Is(to)
	target, err := targetQuery.First()
	if err != nil {
		return Region{}, err
	}

	var closest Region
	minDistance := math.MaxFloat64

	for _, regionList := range regionRegistry {
		for _, region := range regionList {
			if region.Code == to {
				continue // Skip the target region itself
			}

			distance := target.Distance(region)
			if distance < minDistance {
				minDistance = distance
				closest = region
			}
		}
	}

	if closest.Code == "" {
		return Region{}, errors.New("no other regions found")
	}

	return closest, nil
}

// allRegions returns all regions as a Set.
func allRegions() Set {
	regions := make(Set, 0)
	for _, regionList := range regionRegistry {
		regions = append(regions, regionList...)
	}
	return regions
}
