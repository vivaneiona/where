package where

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Code represents a strongly-typed region identifier.
type Code string

// Status represents the operational status of a cloud region.
type Status uint8

const (
	// Active indicates the region is fully operational and available for general use.
	Active Status = iota
	// Deprecated indicates the region is being phased out and should not be used for new resources.
	Deprecated
	// Preview indicates the region has limited availability and may not have all services.
	Preview
)

// String returns the human-readable status name.
func (s Status) String() string {
	switch s {
	case Active:
		return "active"
	case Deprecated:
		return "deprecated"
	case Preview:
		return "preview"
	default:
		return "unknown"
	}
}

// Region represents a cloud provider region with comprehensive metadata.
type Region struct {
	Code       Code      `json:"code"`
	Name       string    `json:"name"`
	Provider   string    `json:"provider"`
	Country    string    `json:"country"`
	City       string    `json:"city"`
	Continent  string    `json:"continent"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Status     Status    `json:"status"`
	LaunchDate time.Time `json:"launch_date"`
	Zones      []string  `json:"zones"`
}

// Distance calculates the great-circle distance to another region in kilometers.
func (r Region) Distance(other Region) float64 {
	return haversineDistance(r.Latitude, r.Longitude, other.Latitude, other.Longitude)
}

// IsActive returns true if the region is currently active and available.
func (r Region) IsActive() bool {
	return r.Status == Active
}

// IsNear returns true if the region is within the specified radius (km) of the target location.
func (r Region) IsNear(lat, lng float64, radiusKm float64) bool {
	return haversineDistance(r.Latitude, r.Longitude, lat, lng) <= radiusKm
}

// Set represents a collection of regions with high-performance operations.
type Set []Region

// Filter applies a predicate function to create a new filtered set.
func (s Set) Filter(predicate func(Region) bool) Set {
	result := make(Set, 0, len(s))
	for _, region := range s {
		if predicate(region) {
			result = append(result, region)
		}
	}
	return result
}

// OnProvider filters regions by cloud provider name.
func (s Set) OnProvider(name string) Set {
	return s.Filter(func(r Region) bool {
		return strings.EqualFold(r.Provider, name)
	})
}

// ByCountry filters regions by country name.
func (s Set) ByCountry(name string) Set {
	return s.Filter(func(r Region) bool {
		return strings.EqualFold(r.Country, name)
	})
}

// ByCity filters regions by city name.
func (s Set) ByCity(name string) Set {
	return s.Filter(func(r Region) bool {
		return strings.EqualFold(r.City, name)
	})
}

// ByContinent filters regions by continent name.
func (s Set) ByContinent(name string) Set {
	return s.Filter(func(r Region) bool {
		return strings.EqualFold(r.Continent, name)
	})
}

// ActiveOnly filters to only active regions.
func (s Set) ActiveOnly() Set {
	return s.Filter(func(r Region) bool {
		return r.Status == Active
	})
}

// Near filters regions within the specified radius of a location.
func (s Set) Near(lat, lng float64, radiusKm float64) Set {
	return s.Filter(func(r Region) bool {
		return r.IsNear(lat, lng, radiusKm)
	})
}

// First returns the first region in the set, or an error if empty.
func (s Set) First() (Region, error) {
	if len(s) == 0 {
		return Region{}, fmt.Errorf("no regions found")
	}
	return s[0], nil
}

// Last returns the last region in the set, or an error if empty.
func (s Set) Last() (Region, error) {
	if len(s) == 0 {
		return Region{}, fmt.Errorf("no regions found")
	}
	return s[len(s)-1], nil
}

// Union returns a new set containing all regions from both sets (no duplicates).
func (s Set) Union(other Set) Set {
	seen := make(map[Code]bool)
	result := make(Set, 0, len(s)+len(other))

	for _, region := range s {
		if !seen[region.Code] {
			result = append(result, region)
			seen[region.Code] = true
		}
	}

	for _, region := range other {
		if !seen[region.Code] {
			result = append(result, region)
			seen[region.Code] = true
		}
	}

	return result
}

// Intersect returns a new set containing only regions present in both sets.
func (s Set) Intersect(other Set) Set {
	otherMap := make(map[Code]bool)
	for _, region := range other {
		otherMap[region.Code] = true
	}

	result := make(Set, 0)
	for _, region := range s {
		if otherMap[region.Code] {
			result = append(result, region)
		}
	}

	return result
}

// Difference returns a new set containing regions in this set but not in the other.
func (s Set) Difference(other Set) Set {
	otherMap := make(map[Code]bool)
	for _, region := range other {
		otherMap[region.Code] = true
	}

	result := make(Set, 0)
	for _, region := range s {
		if !otherMap[region.Code] {
			result = append(result, region)
		}
	}

	return result
}

// SortByDistance sorts regions by distance from a point (closest first).
func (s Set) SortByDistance(lat, lng float64) {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			dist1 := haversineDistance(lat, lng, s[i].Latitude, s[i].Longitude)
			dist2 := haversineDistance(lat, lng, s[j].Latitude, s[j].Longitude)
			if dist1 > dist2 {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

// SortByName sorts regions alphabetically by name.
func (s Set) SortByName() {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i].Name > s[j].Name {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

// SortByProvider sorts regions by provider name.
func (s Set) SortByProvider() {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i].Provider > s[j].Provider {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

// SortByCountry sorts regions by country name.
func (s Set) SortByCountry() {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i].Country > s[j].Country {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

// Codes returns a slice of region codes from the set.
func (s Set) Codes() []Code {
	codes := make([]Code, len(s))
	for i, region := range s {
		codes[i] = region.Code
	}
	return codes
}

// Len returns the number of regions in the set.
func (s Set) Len() int {
	return len(s)
}

// haversineDistance calculates the great-circle distance between two points on Earth.
func haversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadiusKm = 6371.0

	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLng := (lng2 - lng1) * math.Pi / 180.0

	lat1Rad := lat1 * math.Pi / 180.0
	lat2Rad := lat2 * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLng/2)*math.Sin(dLng/2)*math.Cos(lat1Rad)*math.Cos(lat2Rad)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}

// RegionQuery allows filtering regions by provider after lookup by code.
// This enables patterns like: where.Is("us-east-1").OnAWS()
type RegionQuery struct {
	regions []Region
}

// OnAWS filters the query results to only AWS regions.
func (rq RegionQuery) OnAWS() (Region, error) {
	return rq.OnProvider("aws")
}

// OnAzure filters the query results to only Azure regions.
func (rq RegionQuery) OnAzure() (Region, error) {
	return rq.OnProvider("azure")
}

// OnGCP filters the query results to only GCP regions.
func (rq RegionQuery) OnGCP() (Region, error) {
	return rq.OnProvider("gcp")
}

// OnAlibaba filters the query results to only Alibaba regions.
func (rq RegionQuery) OnAlibaba() (Region, error) {
	return rq.OnProvider("alibaba")
}

// OnYandex filters the query results to only Yandex regions.
func (rq RegionQuery) OnYandex() (Region, error) {
	return rq.OnProvider("yandex")
}

// OnProvider filters the query results to only regions from the specified provider.
func (rq RegionQuery) OnProvider(provider string) (Region, error) {
	for _, region := range rq.regions {
		if strings.EqualFold(region.Provider, provider) {
			return region, nil
		}
	}
	return Region{}, fmt.Errorf("%w: no region found for provider %s", ErrRegionNotFound, provider)
}

// All returns all matching regions for the code (useful when multiple providers have the same code).
func (rq RegionQuery) All() []Region {
	return rq.regions
}

// First returns the first matching region.
func (rq RegionQuery) First() (Region, error) {
	if len(rq.regions) == 0 {
		return Region{}, fmt.Errorf("%w", ErrRegionNotFound)
	}
	return rq.regions[0], nil
}
