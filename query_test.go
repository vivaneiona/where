package where

import (
	"testing"
)

func TestNewQuery(t *testing.T) {
	q := NewQuery()
	if q == nil {
		t.Fatal("NewQuery() returned nil")
	}
	if len(q.regions) == 0 {
		t.Error("NewQuery() should start with all regions")
	}
	if len(q.errors) != 0 {
		t.Error("NewQuery() should start with no errors")
	}
}

func TestQueryChaining(t *testing.T) {
	t.Run("method chaining", func(t *testing.T) {
		q := NewQuery().
			InCountry("United States").
			ByProvider("aws").
			ActiveOnly()

		if q == nil {
			t.Fatal("Query chaining returned nil")
		}

		regions := q.Exec()
		// All returned regions should match criteria
		for _, region := range regions {
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
	})

	t.Run("geographic filters", func(t *testing.T) {
		q := NewQuery().InAsia()
		regions := q.Exec()

		// Verify all regions are in Asia
		for _, region := range regions {
			if region.Continent != "Asia" {
				t.Errorf("Expected Asian region, got %v from %v", region.Continent, region.Code)
			}
		}
	})

	t.Run("Europe filter", func(t *testing.T) {
		q := NewQuery().InEurope()
		regions := q.Exec()

		// Verify all regions are in Europe
		for _, region := range regions {
			if region.Continent != "Europe" {
				t.Errorf("Expected European region, got %v from %v", region.Continent, region.Code)
			}
		}
	})
}

func TestQueryFilters(t *testing.T) {
	tests := []struct {
		name   string
		filter func(*Query) *Query
		verify func(Set) error
	}{
		{
			name: "InCountry",
			filter: func(q *Query) *Query {
				return q.InCountry("United States")
			},
			verify: func(regions Set) error {
				for _, region := range regions {
					if region.Country != "United States" {
						t.Errorf("Expected US region, got %v", region.Country)
					}
				}
				return nil
			},
		},
		{
			name: "InCity",
			filter: func(q *Query) *Query {
				return q.InCity("Ashburn")
			},
			verify: func(regions Set) error {
				for _, region := range regions {
					if region.City != "Ashburn" {
						t.Errorf("Expected Ashburn region, got %v", region.City)
					}
				}
				return nil
			},
		},
		{
			name: "ByProvider",
			filter: func(q *Query) *Query {
				return q.ByProvider("aws")
			},
			verify: func(regions Set) error {
				for _, region := range regions {
					if region.Provider != "aws" {
						t.Errorf("Expected AWS region, got %v", region.Provider)
					}
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQuery()
			q = tt.filter(q)
			regions := q.Exec()

			if len(regions) == 0 {
				t.Error("Filter should return some regions")
			}

			tt.verify(regions)
		})
	}
}

func TestQuerySorting(t *testing.T) {
	t.Run("SortByName", func(t *testing.T) {
		regions := NewQuery().
			ByProvider("aws").
			SortByName().
			Exec()

		// Verify sorted order
		for i := 1; i < len(regions); i++ {
			if regions[i-1].Name > regions[i].Name {
				t.Error("Regions are not sorted by name")
				break
			}
		}
	})

	t.Run("SortByDistance", func(t *testing.T) {
		// Sort by distance from NYC
		regions := NewQuery().
			ByProvider("aws").
			SortByDistance(40.7128, -74.0060).
			Exec()

		if len(regions) == 0 {
			t.Skip("No AWS regions found for distance test")
		}

		// Verify sorted order (closest first)
		nycLat, nycLng := 40.7128, -74.0060
		lastDistance := -1.0
		for _, region := range regions {
			distance := region.Distance(Region{Latitude: nycLat, Longitude: nycLng})
			if lastDistance >= 0 && distance < lastDistance {
				t.Error("Regions are not sorted by distance")
				break
			}
			lastDistance = distance
		}
	})
}

func TestQueryLimits(t *testing.T) {
	t.Run("Limit", func(t *testing.T) {
		limit := 3
		regions := NewQuery().
			ByProvider("aws").
			Limit(limit).
			Exec()

		if len(regions) > limit {
			t.Errorf("Expected at most %d regions, got %d", limit, len(regions))
		}
	})

	t.Run("First", func(t *testing.T) {
		region, err := NewQuery().
			ByProvider("aws").
			First()

		if err != nil {
			t.Errorf("First() error = %v", err)
		}
		if region.Provider != "aws" {
			t.Errorf("Expected AWS region, got %v", region.Provider)
		}
	})

	t.Run("First on empty result", func(t *testing.T) {
		_, err := NewQuery().
			InCountry("NonexistentCountry").
			First()

		if err == nil {
			t.Error("First() should return error for empty result")
		}
	})
}

func TestQueryErrorHandling(t *testing.T) {
	t.Run("no errors on valid query", func(t *testing.T) {
		q := NewQuery().ByProvider("aws")
		regions, errors := q.ExecWithErrors()

		if len(errors) != 0 {
			t.Errorf("Expected no errors, got %d", len(errors))
		}

		if len(regions) == 0 {
			t.Error("Expected some AWS regions")
		}
	})
}

// Test edge cases and combinations
func TestQueryEdgeCases(t *testing.T) {
	t.Run("empty result set", func(t *testing.T) {
		regions := NewQuery().
			InCountry("NonexistentCountry").
			Exec()

		if len(regions) != 0 {
			t.Error("Query for nonexistent country should return empty set")
		}
	})

	t.Run("multiple providers", func(t *testing.T) {
		// This should work as providers are OR'd in typical implementations
		awsRegions := NewQuery().ByProvider("aws").Exec()
		azureRegions := NewQuery().ByProvider("azure").Exec()

		if len(awsRegions) == 0 && len(azureRegions) == 0 {
			t.Skip("No AWS or Azure regions found")
		}
	})

	t.Run("complex chaining", func(t *testing.T) {
		regions := NewQuery().
			InContinent("North America").
			ByProvider("aws").
			ActiveOnly().
			SortByName().
			Limit(5).
			Exec()

		// Verify all criteria
		for _, region := range regions {
			if region.Continent != "North America" {
				t.Errorf("Expected North American region, got %v", region.Continent)
			}
			if region.Provider != "aws" {
				t.Errorf("Expected AWS region, got %v", region.Provider)
			}
			if !region.IsActive() {
				t.Error("Expected active region")
			}
		}

		if len(regions) > 5 {
			t.Errorf("Expected at most 5 regions, got %d", len(regions))
		}
	})
}

// Benchmark query operations
func BenchmarkQueryChaining(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewQuery().
			InCountry("United States").
			ByProvider("aws").
			ActiveOnly().
			Exec()
	}
}

func BenchmarkQuerySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewQuery().
			ByProvider("aws").
			SortByDistance(40.7128, -74.0060).
			Exec()
	}
}
