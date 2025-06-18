package where

import (
	"math"
	"testing"
)

func TestStatus_String(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{Active, "active"},
		{Deprecated, "deprecated"},
		{Preview, "preview"},
		{Status(99), "unknown"},
	}

	for _, test := range tests {
		if got := test.status.String(); got != test.expected {
			t.Errorf("Status(%d).String() = %q, want %q", test.status, got, test.expected)
		}
	}
}

func TestRegion_IsActive(t *testing.T) {
	tests := []struct {
		name     string
		region   Region
		expected bool
	}{
		{
			name:     "active region",
			region:   Region{Status: Active},
			expected: true,
		},
		{
			name:     "deprecated region",
			region:   Region{Status: Deprecated},
			expected: false,
		},
		{
			name:     "preview region",
			region:   Region{Status: Preview},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.region.IsActive(); got != test.expected {
				t.Errorf("Region.IsActive() = %v, want %v", got, test.expected)
			}
		})
	}
}

func TestRegion_Distance(t *testing.T) {
	// Test with known coordinates
	nyc := Region{Latitude: 40.7128, Longitude: -74.0060}
	london := Region{Latitude: 51.5074, Longitude: -0.1278}

	distance := nyc.Distance(london)

	// Approximate distance between NYC and London is ~5570 km
	expectedDistance := 5570.0
	tolerance := 50.0 // Allow 50km tolerance

	if math.Abs(distance-expectedDistance) > tolerance {
		t.Errorf("Distance between NYC and London = %v km, want approximately %v km", distance, expectedDistance)
	}

	// Test distance to self should be 0
	selfDistance := nyc.Distance(nyc)
	if selfDistance != 0 {
		t.Errorf("Distance to self = %v, want 0", selfDistance)
	}
}

func TestRegion_IsNear(t *testing.T) {
	region := Region{Latitude: 40.7128, Longitude: -74.0060} // NYC

	tests := []struct {
		name     string
		lat      float64
		lng      float64
		radiusKm float64
		expected bool
	}{
		{
			name:     "same location",
			lat:      40.7128,
			lng:      -74.0060,
			radiusKm: 1,
			expected: true,
		},
		{
			name:     "within radius",
			lat:      40.7589, // Manhattan
			lng:      -73.9851,
			radiusKm: 10,
			expected: true,
		},
		{
			name:     "outside radius",
			lat:      51.5074, // London
			lng:      -0.1278,
			radiusKm: 100,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := region.IsNear(test.lat, test.lng, test.radiusKm); got != test.expected {
				t.Errorf("Region.IsNear(%v, %v, %v) = %v, want %v",
					test.lat, test.lng, test.radiusKm, got, test.expected)
			}
		})
	}
}

func TestSet_Filter(t *testing.T) {
	regions := Set{
		{Code: "us-east-1", Provider: "aws", Status: Active},
		{Code: "us-west-1", Provider: "aws", Status: Deprecated},
		{Code: "eastus", Provider: "azure", Status: Active},
	}

	// Filter by provider
	awsRegions := regions.Filter(func(r Region) bool {
		return r.Provider == "aws"
	})

	if len(awsRegions) != 2 {
		t.Errorf("Expected 2 AWS regions, got %d", len(awsRegions))
	}

	// Filter by status
	activeRegions := regions.Filter(func(r Region) bool {
		return r.Status == Active
	})

	if len(activeRegions) != 2 {
		t.Errorf("Expected 2 active regions, got %d", len(activeRegions))
	}
}

func TestSet_ByProvider(t *testing.T) {
	regions := Set{
		{Code: "us-east-1", Provider: "aws"},
		{Code: "us-west-1", Provider: "AWS"}, // Test case insensitive
		{Code: "eastus", Provider: "azure"},
	}

	awsRegions := regions.OnProvider("aws")
	if len(awsRegions) != 2 {
		t.Errorf("Expected 2 AWS regions, got %d", len(awsRegions))
	}
}

func TestSet_ByCountry(t *testing.T) {
	regions := Set{
		{Code: "us-east-1", Country: "United States"},
		{Code: "us-west-1", Country: "united states"}, // Test case insensitive
		{Code: "eu-west-1", Country: "Ireland"},
	}

	usRegions := regions.ByCountry("United States")
	if len(usRegions) != 2 {
		t.Errorf("Expected 2 US regions, got %d", len(usRegions))
	}
}

func TestSet_ByCity(t *testing.T) {
	regions := Set{
		{Code: "us-east-1", City: "Ashburn"},
		{Code: "us-east-2", City: "Columbus"},
		{Code: "test", City: "ashburn"}, // Test case insensitive
	}

	ashburnRegions := regions.ByCity("Ashburn")
	if len(ashburnRegions) != 2 {
		t.Errorf("Expected 2 Ashburn regions, got %d", len(ashburnRegions))
	}
}

func TestSet_ByContinent(t *testing.T) {
	regions := Set{
		{Code: "us-east-1", Continent: "North America"},
		{Code: "eu-west-1", Continent: "Europe"},
		{Code: "ca-central-1", Continent: "north america"}, // Test case insensitive
	}

	naRegions := regions.ByContinent("North America")
	if len(naRegions) != 2 {
		t.Errorf("Expected 2 North America regions, got %d", len(naRegions))
	}
}

func TestSet_ActiveOnly(t *testing.T) {
	regions := Set{
		{Code: "us-east-1", Status: Active},
		{Code: "us-west-1", Status: Deprecated},
		{Code: "test", Status: Preview},
		{Code: "eastus", Status: Active},
	}

	activeRegions := regions.ActiveOnly()
	if len(activeRegions) != 2 {
		t.Errorf("Expected 2 active regions, got %d", len(activeRegions))
	}
}

func TestSet_Near(t *testing.T) {
	regions := Set{
		{Code: "nyc", Latitude: 40.7128, Longitude: -74.0060},
		{Code: "boston", Latitude: 42.3601, Longitude: -71.0589},
		{Code: "london", Latitude: 51.5074, Longitude: -0.1278},
	}

	// Find regions near NYC within 500km
	nearRegions := regions.Near(40.7128, -74.0060, 500)

	// Should include NYC and Boston, but not London
	if len(nearRegions) != 2 {
		t.Errorf("Expected 2 regions near NYC, got %d", len(nearRegions))
	}
}

func TestSet_First(t *testing.T) {
	// Test with empty set
	emptySet := Set{}
	_, err := emptySet.First()
	if err == nil {
		t.Error("Expected error for empty set, got nil")
	}

	// Test with regions
	regions := Set{
		{Code: "first"},
		{Code: "second"},
	}

	first, err := regions.First()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if first.Code != "first" {
		t.Errorf("Expected first region code 'first', got %v", first.Code)
	}
}

func TestSet_Last(t *testing.T) {
	// Test with empty set
	emptySet := Set{}
	_, err := emptySet.Last()
	if err == nil {
		t.Error("Expected error for empty set, got nil")
	}

	// Test with regions
	regions := Set{
		{Code: "first"},
		{Code: "last"},
	}

	last, err := regions.Last()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if last.Code != "last" {
		t.Errorf("Expected last region code 'last', got %v", last.Code)
	}
}

func TestSet_Union(t *testing.T) {
	set1 := Set{
		{Code: "us-east-1"},
		{Code: "us-west-1"},
	}

	set2 := Set{
		{Code: "us-west-1"}, // Duplicate
		{Code: "eu-west-1"}, // New
	}

	union := set1.Union(set2)

	// Should have 3 unique regions
	if len(union) != 3 {
		t.Errorf("Expected 3 regions in union, got %d", len(union))
	}

	// Check for duplicates
	codes := make(map[Code]int)
	for _, region := range union {
		codes[region.Code]++
	}

	for code, count := range codes {
		if count > 1 {
			t.Errorf("Duplicate region %v found %d times", code, count)
		}
	}
}

func TestSet_Intersect(t *testing.T) {
	set1 := Set{
		{Code: "us-east-1"},
		{Code: "us-west-1"},
		{Code: "eu-west-1"},
	}

	set2 := Set{
		{Code: "us-west-1"},
		{Code: "eu-west-1"},
		{Code: "ap-south-1"},
	}

	intersect := set1.Intersect(set2)

	// Should have 2 common regions
	if len(intersect) != 2 {
		t.Errorf("Expected 2 regions in intersection, got %d", len(intersect))
	}
}

func TestSet_Difference(t *testing.T) {
	set1 := Set{
		{Code: "us-east-1"},
		{Code: "us-west-1"},
		{Code: "eu-west-1"},
	}

	set2 := Set{
		{Code: "us-west-1"},
		{Code: "ap-south-1"},
	}

	diff := set1.Difference(set2)

	// Should have regions in set1 but not in set2
	if len(diff) != 2 {
		t.Errorf("Expected 2 regions in difference, got %d", len(diff))
	}

	expectedCodes := map[Code]bool{
		"us-east-1": true,
		"eu-west-1": true,
	}

	for _, region := range diff {
		if !expectedCodes[region.Code] {
			t.Errorf("Unexpected region %v in difference", region.Code)
		}
	}
}

func TestSet_SortByDistance(t *testing.T) {
	regions := Set{
		{Code: "far", Latitude: 0, Longitude: 0},
		{Code: "near", Latitude: 40.7, Longitude: -74.0},
		{Code: "medium", Latitude: 42.3, Longitude: -71.0},
	}

	// Sort by distance from NYC (40.7128, -74.0060)
	regions.SortByDistance(40.7128, -74.0060)

	// "near" should be first (closest to NYC)
	if regions[0].Code != "near" {
		t.Errorf("Expected 'near' to be closest, got %v", regions[0].Code)
	}
}

func TestSet_SortByName(t *testing.T) {
	regions := Set{
		{Code: "test", Name: "Zebra"},
		{Code: "test", Name: "Alpha"},
		{Code: "test", Name: "Beta"},
	}

	regions.SortByName()

	expectedOrder := []string{"Alpha", "Beta", "Zebra"}
	for i, expected := range expectedOrder {
		if regions[i].Name != expected {
			t.Errorf("Expected %v at position %d, got %v", expected, i, regions[i].Name)
		}
	}
}

func TestSet_SortByProvider(t *testing.T) {
	regions := Set{
		{Code: "test", Provider: "gcp"},
		{Code: "test", Provider: "aws"},
		{Code: "test", Provider: "azure"},
	}

	regions.SortByProvider()

	expectedOrder := []string{"aws", "azure", "gcp"}
	for i, expected := range expectedOrder {
		if regions[i].Provider != expected {
			t.Errorf("Expected %v at position %d, got %v", expected, i, regions[i].Provider)
		}
	}
}

func TestSet_SortByCountry(t *testing.T) {
	regions := Set{
		{Code: "test", Country: "United States"},
		{Code: "test", Country: "Canada"},
		{Code: "test", Country: "Germany"},
	}

	regions.SortByCountry()

	expectedOrder := []string{"Canada", "Germany", "United States"}
	for i, expected := range expectedOrder {
		if regions[i].Country != expected {
			t.Errorf("Expected %v at position %d, got %v", expected, i, regions[i].Country)
		}
	}
}

func TestSet_Codes(t *testing.T) {
	regions := Set{
		{Code: "us-east-1"},
		{Code: "us-west-1"},
		{Code: "eu-west-1"},
	}

	codes := regions.Codes()

	if len(codes) != 3 {
		t.Errorf("Expected 3 codes, got %d", len(codes))
	}

	expectedCodes := map[Code]bool{
		"us-east-1": true,
		"us-west-1": true,
		"eu-west-1": true,
	}

	for _, code := range codes {
		if !expectedCodes[code] {
			t.Errorf("Unexpected code %v", code)
		}
	}
}

func TestSet_Len(t *testing.T) {
	regions := Set{
		{Code: "us-east-1"},
		{Code: "us-west-1"},
	}

	if regions.Len() != 2 {
		t.Errorf("Expected length 2, got %d", regions.Len())
	}

	emptySet := Set{}
	if emptySet.Len() != 0 {
		t.Errorf("Expected length 0 for empty set, got %d", emptySet.Len())
	}
}

func TestHaversineDistance(t *testing.T) {
	// Test known distance: NYC to London
	nycLat, nycLng := 40.7128, -74.0060
	londonLat, londonLng := 51.5074, -0.1278

	distance := haversineDistance(nycLat, nycLng, londonLat, londonLng)

	// Expected distance is approximately 5570 km
	expectedDistance := 5570.0
	tolerance := 50.0

	if math.Abs(distance-expectedDistance) > tolerance {
		t.Errorf("Haversine distance = %v km, want approximately %v km", distance, expectedDistance)
	}

	// Test zero distance
	zeroDistance := haversineDistance(nycLat, nycLng, nycLat, nycLng)
	if zeroDistance != 0 {
		t.Errorf("Distance to same point = %v, want 0", zeroDistance)
	}
}

// Benchmark tests
func BenchmarkRegion_Distance(b *testing.B) {
	region1 := Region{Latitude: 40.7128, Longitude: -74.0060}
	region2 := Region{Latitude: 51.5074, Longitude: -0.1278}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = region1.Distance(region2)
	}
}

func BenchmarkSet_Filter(b *testing.B) {
	regions := make(Set, 1000)
	for i := 0; i < 1000; i++ {
		regions[i] = Region{
			Code:     Code("region-" + string(rune(i))),
			Provider: "aws",
			Status:   Active,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = regions.Filter(func(r Region) bool {
			return r.Provider == "aws"
		})
	}
}

func BenchmarkSet_SortByDistance(b *testing.B) {
	regions := make(Set, 100)
	for i := 0; i < 100; i++ {
		regions[i] = Region{
			Code:      Code("region-" + string(rune(i))),
			Latitude:  float64(i),
			Longitude: float64(i),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		regionsCopy := make(Set, len(regions))
		copy(regionsCopy, regions)
		regionsCopy.SortByDistance(40.7128, -74.0060)
	}
}
