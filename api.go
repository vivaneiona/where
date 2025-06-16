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

// Is answers "where is {code}?" - returns detailed information about a specific region.
func Is(code Code) (Region, error) {
	region, exists := regionRegistry[code]
	if !exists {
		return Region{}, fmt.Errorf("%w: %s", ErrRegionNotFound, code)
	}
	return region, nil
}

// MustIs is like Is but panics on error. Use when you're certain the region exists.
func MustIs(code Code) Region {
	region, err := Is(code)
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
		if region, exists := regionRegistry[code]; exists {
			regions = append(regions, region)
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

// ByProvider answers "where by provider {name}?" - returns all regions from a provider.
func ByProvider(name string) Set {
	return allRegions().ByProvider(name)
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
	if region, exists := regionRegistry[code]; exists {
		return region.IsActive()
	}
	return false
}

// HasProvider answers "where has provider {name}?" - checks if a provider exists.
func HasProvider(name string) bool {
	for _, region := range regionRegistry {
		if strings.EqualFold(region.Provider, name) {
			return true
		}
	}
	return false
}

// Discovery functions for exploration

// Providers answers "where providers?" - returns all unique provider names.
func Providers() []string {
	providerMap := make(map[string]bool)
	for _, region := range regionRegistry {
		providerMap[region.Provider] = true
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
	for _, region := range regionRegistry {
		countryMap[region.Country] = true
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
	for _, region := range regionRegistry {
		cityMap[region.City] = true
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
	for _, region := range regionRegistry {
		continentMap[region.Continent] = true
	}

	continents := make([]string, 0, len(continentMap))
	for continent := range continentMap {
		continents = append(continents, continent)
	}
	return continents
}

// Distance calculates the distance between two regions in kilometers.
func Distance(from, to Code) (float64, error) {
	fromRegion, err := Is(from)
	if err != nil {
		return 0, fmt.Errorf("source region %w", err)
	}

	toRegion, err := Is(to)
	if err != nil {
		return 0, fmt.Errorf("destination region %w", err)
	}

	return fromRegion.Distance(toRegion), nil
}

// Closest finds the closest region to the specified region.
func Closest(to Code) (Region, error) {
	target, err := Is(to)
	if err != nil {
		return Region{}, err
	}

	var closest Region
	minDistance := math.MaxFloat64

	for _, region := range regionRegistry {
		if region.Code == to {
			continue // Skip the target region itself
		}

		distance := target.Distance(region)
		if distance < minDistance {
			minDistance = distance
			closest = region
		}
	}

	if closest.Code == "" {
		return Region{}, errors.New("no other regions found")
	}

	return closest, nil
}

// allRegions returns all regions as a Set.
func allRegions() Set {
	regions := make(Set, 0, len(regionRegistry))
	for _, region := range regionRegistry {
		regions = append(regions, region)
	}
	return regions
}
