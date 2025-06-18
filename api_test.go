package where

import (
	"errors"
	"testing"
)

func TestIs(t *testing.T) {
	tests := []struct {
		name    string
		code    Code
		wantErr bool
	}{
		{
			name:    "valid region",
			code:    "us-east-1",
			wantErr: false,
		},
		{
			name:    "invalid region",
			code:    "invalid-region",
			wantErr: true,
		},
		{
			name:    "empty code",
			code:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := Is(tt.code)
			region, err := query.First()
			if (err != nil) != tt.wantErr {
				t.Errorf("Is() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && region.Code != tt.code {
				t.Errorf("Is() returned wrong region code: got %v, want %v", region.Code, tt.code)
			}
			if tt.wantErr && !errors.Is(err, ErrRegionNotFound) {
				t.Errorf("Is() should return ErrRegionNotFound, got %v", err)
			}
		})
	}
}

func TestMustIs(t *testing.T) {
	// Test successful case
	t.Run("valid region", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("MustIs() panicked unexpectedly: %v", r)
			}
		}()
		region := MustIs("us-east-1")
		if region.Code != "us-east-1" {
			t.Errorf("MustIs() returned wrong region")
		}
	})

	// Test panic case
	t.Run("invalid region should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustIs() should have panicked for invalid region")
			}
		}()
		MustIs("invalid-region")
	})
}

func TestAre(t *testing.T) {
	tests := []struct {
		name     string
		codes    []Code
		wantLen  int
		wantErr  bool
		errCount int
	}{
		{
			name:    "all valid regions",
			codes:   []Code{"us-east-1", "us-west-2"},
			wantLen: 3,
			wantErr: false,
		},
		{
			name:     "mixed valid and invalid",
			codes:    []Code{"us-east-1", "invalid", "us-west-2"},
			wantLen:  3,
			wantErr:  true,
			errCount: 1,
		},
		{
			name:    "empty input",
			codes:   []Code{},
			wantLen: 0,
			wantErr: false,
		},
		{
			name:     "all invalid",
			codes:    []Code{"invalid1", "invalid2"},
			wantLen:  0,
			wantErr:  true,
			errCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regions, err := Are(tt.codes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Are() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(regions) != tt.wantLen {
				t.Errorf("Are() returned %d regions, want %d", len(regions), tt.wantLen)
			}
			if tt.wantErr && !errors.Is(err, ErrRegionNotFound) {
				t.Errorf("Are() should return ErrRegionNotFound, got %v", err)
			}
		})
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		name string
		code string
		want bool
	}{
		{
			name: "valid region",
			code: "us-east-1",
			want: true,
		},
		{
			name: "invalid region",
			code: "invalid-region",
			want: false,
		},
		{
			name: "empty string",
			code: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Has(tt.code); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsActive(t *testing.T) {
	tests := []struct {
		name string
		code Code
		want bool
	}{
		{
			name: "active region",
			code: "us-east-1",
			want: true,
		},
		{
			name: "invalid region",
			code: "invalid-region",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsActive(tt.code); got != tt.want {
				t.Errorf("IsActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscoveryFunctions(t *testing.T) {
	t.Run("Providers", func(t *testing.T) {
		providers := Providers()
		if len(providers) == 0 {
			t.Error("Providers() returned empty slice")
		}
		// Verify contains expected providers
		found := false
		for _, p := range providers {
			if p == "aws" {
				found = true
				break
			}
		}
		if !found {
			t.Error("Providers() should include 'aws'")
		}
	})

	t.Run("Countries", func(t *testing.T) {
		countries := Countries()
		if len(countries) == 0 {
			t.Error("Countries() returned empty slice")
		}
	})

	t.Run("Cities", func(t *testing.T) {
		cities := Cities()
		if len(cities) == 0 {
			t.Error("Cities() returned empty slice")
		}
	})

	t.Run("Continents", func(t *testing.T) {
		continents := Continents()
		if len(continents) == 0 {
			t.Error("Continents() returned empty slice")
		}
	})
}

func TestDistance(t *testing.T) {
	tests := []struct {
		name    string
		from    Code
		to      Code
		wantErr bool
	}{
		{
			name:    "valid regions",
			from:    "us-east-1",
			to:      "us-west-2",
			wantErr: false,
		},
		{
			name:    "invalid from region",
			from:    "invalid",
			to:      "us-west-2",
			wantErr: true,
		},
		{
			name:    "invalid to region",
			from:    "us-east-1",
			to:      "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distance, err := Distance(tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && distance < 0 {
				t.Error("Distance() should return non-negative value")
			}
		})
	}
}

func TestClosest(t *testing.T) {
	tests := []struct {
		name    string
		to      Code
		wantErr bool
	}{
		{
			name:    "valid region",
			to:      "us-east-1",
			wantErr: false,
		},
		{
			name:    "invalid region",
			to:      "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closest, err := Closest(tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("Closest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && closest.Code == tt.to {
				t.Error("Closest() should not return the same region")
			}
		})
	}
}

func TestFilteredQueries(t *testing.T) {
	t.Run("InCountry", func(t *testing.T) {
		regions := InCountry("United States")
		if len(regions) == 0 {
			t.Error("InCountry('United States') should return regions")
		}
	})

	t.Run("ByProvider", func(t *testing.T) {
		regions := OnProvider("aws")
		if len(regions) == 0 {
			t.Error("ByProvider('aws') should return regions")
		}
	})

	t.Run("ActiveRegions", func(t *testing.T) {
		regions := ActiveRegions()
		if len(regions) == 0 {
			t.Error("ActiveRegions() should return regions")
		}
		// Verify all returned regions are active
		for _, region := range regions {
			if !region.IsActive() {
				t.Error("ActiveRegions() returned inactive region")
			}
		}
	})

	t.Run("Near", func(t *testing.T) {
		// Test near NYC coordinates
		regions := Near(40.7128, -74.0060, 1000) // 1000km radius
		if len(regions) == 0 {
			t.Error("Near() should return regions within radius")
		}
	})
}

// Benchmark critical API functions
func BenchmarkIs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		query := Is("us-east-1")
		_, _ = query.First()
	}
}

func BenchmarkAre(b *testing.B) {
	codes := []Code{"us-east-1", "us-west-2", "eu-west-1"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Are(codes...)
	}
}

func BenchmarkByProvider(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = OnProvider("aws")
	}
}
