package where_test

import (
	"fmt"
	"testing"

	"github.com/vivaneiona/where"
)

func TestIntegration(t *testing.T) {
	t.Run("basic region operations", func(t *testing.T) {
		r := where.Region{
			Code:      "test",
			Name:      "Test Region",
			Provider:  "test",
			Country:   "Test Country",
			City:      "Test City",
			Continent: "Test Continent",
			Latitude:  40.7128,
			Longitude: -74.0060,
			Status:    where.Active,
		}

		if !r.IsActive() {
			t.Error("Region should be active")
		}

		distance := r.Distance(r)
		if distance != 0 {
			t.Errorf("Distance to self should be 0, got %v", distance)
		}

		if !r.IsNear(40.7128, -74.0060, 1) {
			t.Error("Region should be near itself")
		}
	})

	t.Run("API workflow", func(t *testing.T) {
		if !where.Exists("us-east-1") {
			t.Skip("us-east-1 not available for integration test")
		}

		region, err := where.Is("us-east-1")
		if err != nil {
			t.Fatalf("Failed to get us-east-1: %v", err)
		}

		if !where.IsActive("us-east-1") {
			t.Error("us-east-1 should be active")
		}

		nearby := where.Near(region.Latitude, region.Longitude, 1000)
		found := false
		for _, r := range nearby {
			if r.Code == "us-east-1" {
				found = true
				break
			}
		}
		if !found {
			t.Error("us-east-1 should be in its own nearby results")
		}
	})

	t.Run("query builder workflow", func(t *testing.T) {
		// Test complex query workflow
		usAwsRegions := where.NewQuery().
			InCountry("United States").
			ByProvider("aws").
			ActiveOnly().
			SortByName().
			Exec()

		for _, region := range usAwsRegions {
			if region.Country != "United States" {
				t.Errorf("Expected US region, got %v", region.Country)
			}
			if region.Provider != "aws" {
				t.Errorf("Expected AWS region, got %v", region.Provider)
			}
			if !region.IsActive() {
				t.Error("Expected active region")
			}
		}

		// Verify sorted order
		for i := 1; i < len(usAwsRegions); i++ {
			if usAwsRegions[i-1].Name > usAwsRegions[i].Name {
				t.Error("Regions should be sorted by name")
				break
			}
		}
	})

	t.Run("provider constants", func(t *testing.T) {
		// Test that provider constants work
		if where.AWS.USEast1 != "us-east-1" {
			t.Errorf("AWS.USEast1 should be 'us-east-1', got %v", where.AWS.USEast1)
		}

		// Test using constants with API
		if where.Exists(string(where.AWS.USEast1)) {
			region := where.MustIs(where.AWS.USEast1)
			if region.Provider != "aws" {
				t.Errorf("Expected AWS provider, got %v", region.Provider)
			}
		}
	})

	t.Run("set operations", func(t *testing.T) {
		// Create test sets
		set1 := where.Set{
			{Code: "us-east-1", Provider: "aws"},
			{Code: "us-west-2", Provider: "aws"},
		}

		set2 := where.Set{
			{Code: "us-west-2", Provider: "aws"},
			{Code: "eu-west-1", Provider: "aws"},
		}

		// Test union
		union := set1.Union(set2)
		if len(union) != 3 {
			t.Errorf("Union should have 3 regions, got %d", len(union))
		}

		// Test intersection
		intersect := set1.Intersect(set2)
		if len(intersect) != 1 {
			t.Errorf("Intersection should have 1 region, got %d", len(intersect))
		}

		// Test difference
		diff := set1.Difference(set2)
		if len(diff) != 1 {
			t.Errorf("Difference should have 1 region, got %d", len(diff))
		}
	})

	t.Run("error scenarios", func(t *testing.T) {
		// Test error conditions
		_, err := where.Is("nonexistent-region")
		if err == nil {
			t.Error("Should get error for nonexistent region")
		}

		_, err = where.Are("us-east-1", "nonexistent", "us-west-2")
		if err == nil {
			t.Error("Should get error when some regions don't exist")
		}

		// Test with empty sets
		emptySet := where.Set{}
		_, err = emptySet.First()
		if err == nil {
			t.Error("Should get error for First() on empty set")
		}

		_, err = emptySet.Last()
		if err == nil {
			t.Error("Should get error for Last() on empty set")
		}
	})
}

func TestProviderNamespaces(t *testing.T) {
	t.Run("AWS constants", func(t *testing.T) {
		tests := []struct {
			constant where.Code
			expected string
		}{
			{where.AWS.USEast1, "us-east-1"},
			{where.AWS.USWest2, "us-west-2"},
			{where.AWS.EUWest1, "eu-west-1"},
		}

		for _, test := range tests {
			if string(test.constant) != test.expected {
				t.Errorf("AWS constant mismatch: got %v, want %v", test.constant, test.expected)
			}
		}
	})

	t.Run("provider regions exist", func(t *testing.T) {
		// Test a few key regions exist in registry
		keyRegions := []where.Code{
			where.AWS.USEast1,
			where.AWS.USWest2,
			where.AWS.EUWest1,
		}

		for _, code := range keyRegions {
			if !where.Exists(string(code)) {
				t.Errorf("Provider constant %v should map to valid region", code)
			}
		}
	})
}

func TestRealWorldScenarios(t *testing.T) {
	t.Run("find closest region", func(t *testing.T) {
		// Skip if no regions available
		allRegions := where.ActiveRegions()
		if len(allRegions) == 0 {
			t.Skip("No regions available for test")
		}

		// Test finding closest to a specific location
		regions := allRegions
		if len(regions) < 2 {
			t.Skip("Need at least 2 regions for closest test")
		}

		firstRegion := regions[0]
		closest, err := where.Closest(firstRegion.Code)
		if err != nil {
			t.Errorf("Closest() failed: %v", err)
		}

		if closest.Code == firstRegion.Code {
			t.Error("Closest region should not be the same region")
		}
	})

	t.Run("performance search", func(t *testing.T) {
		// Find best regions for specific criteria
		performantRegions := where.NewQuery().
			ActiveOnly().
			SortByDistance(37.7749, -122.4194). // SF coordinates
			Limit(3).
			Exec()

		if len(performantRegions) > 3 {
			t.Error("Limit(3) should return at most 3 regions")
		}

		// Verify all are active
		for _, region := range performantRegions {
			if !region.IsActive() {
				t.Error("All regions should be active")
			}
		}
	})

	t.Run("discovery workflow", func(t *testing.T) {
		// Test discovery functions
		providers := where.Providers()
		if len(providers) == 0 {
			t.Skip("No providers available")
		}

		countries := where.Countries()
		if len(countries) == 0 {
			t.Skip("No countries available")
		}

		// Use discovered data
		firstProvider := providers[0]
		providerRegions := where.ByProvider(firstProvider)

		if len(providerRegions) == 0 {
			t.Errorf("Provider %v should have regions", firstProvider)
		}
	})
}

func Example_integration() {
	// Check if a region exists
	if where.Exists("us-east-1") {
		// Get region details
		region, _ := where.Is("us-east-1")
		fmt.Printf("Region: %s in %s\n", region.Name, region.Country)
	}

	// Find all AWS regions in the US
	usAws := where.NewQuery().
		InCountry("United States").
		ByProvider("aws").
		ActiveOnly().
		Exec()

	fmt.Printf("Found %d AWS regions in the US\n", len(usAws))

}

func TestPackageInitialization(t *testing.T) {
	t.Run("regions loaded", func(t *testing.T) {
		all := where.ActiveRegions()
		if len(all) == 0 {
			t.Error("No regions loaded - package initialization failed")
		}
	})

	t.Run("basic operations work", func(_ *testing.T) {
		// Ensure basic operations don't panic
		_ = where.Providers()
		_ = where.Countries()
		_ = where.Cities()
		_ = where.Continents()
		_ = where.ActiveRegions()
		_ = where.PreviewRegions()
		_ = where.DeprecatedRegions()
	})
}
