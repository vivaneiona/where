/*
Package where provides a unified Go library for discovering, querying, and working
with cloud provider regions across AWS, Azure, and Google Cloud Platform.

# Overview

The where package offers multiple intuitive APIs for region discovery:

  - Question-style API: where.Is("us-east-1"), where.Are("us-east-1", "eu-west-1")
  - Provider constants: where.AWS.USEast1, where.GCP.USCentral1, where.Azure.EastUS
  - Geographic queries: where.In.Asia(), where.In.Country("Japan")
  - Provider filtering: where.On.AWS(), where.On.Azure(), where.On.GCP()
  - Builder pattern: where.NewQuery().InCountry("Japan").ByProvider("aws").Exec()

# Quick Start

Basic region lookup:

	region, err := where.Is("us-east-1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Region: %s in %s, %s\n", region.Name, region.City, region.Country)

Provider constants for type-safe, zero-allocation access:

	// Direct region code access
	fmt.Printf("AWS: %s\n", where.AWS.USEast1)    // "us-east-1"
	fmt.Printf("GCP: %s\n", where.GCP.USCentral1) // "us-central1"
	fmt.Printf("Azure: %s\n", where.Azure.EastUS) // "eastus"

Geographic and provider-based queries:

	// Get all regions in Asia
	asianRegions := where.In.Asia()

	// Get all AWS regions
	awsRegions := where.On.AWS()

	// Complex query with builder pattern
	results := where.NewQuery().
		InCountry("Japan").
		ByProvider("aws").
		ActiveOnly().
		Exec()

# API Design Philosophy

The where package is designed around natural language patterns that read like questions:

  - "Where is us-east-1?" → where.Is("us-east-1")
  - "Where are these regions?" → where.Are("us-east-1", "eu-west-1")
  - "What's in Asia?" → where.In.Asia()
  - "What does AWS have?" → where.On.AWS()

# Provider Constants

Each cloud provider has a dedicated namespace with type-safe region constants:

AWS Regions:

	where.AWS.USEast1      // us-east-1 (N. Virginia)
	where.AWS.USWest2      // us-west-2 (Oregon)
	where.AWS.EUWest1      // eu-west-1 (Ireland)
	where.AWS.APSoutheast1 // ap-southeast-1 (Singapore)

Google Cloud Platform:

	where.GCP.USCentral1    // us-central1 (Iowa)
	where.GCP.USWest1       // us-west1 (Oregon)
	where.GCP.EuropeWest1   // europe-west1 (Belgium)
	where.GCP.AsiaSouth1    // asia-south1 (Mumbai)

Microsoft Azure:

	where.Azure.EastUS      // eastus (Virginia)
	where.Azure.WestUS2     // westus2 (Washington)
	where.Azure.WestEurope  // westeurope (Netherlands)
	where.Azure.SoutheastAsia // southeastasia (Singapore)

# Geographic Queries

The In namespace provides geographic-based region discovery:

Continent-based queries:

	where.In.Asia()      // All Asian regions
	where.In.Europe()    // All European regions
	where.In.Americas()  // All North & South American regions
	where.In.Oceania()   // All Oceania regions
	where.In.Africa()    // All African regions

Country and city-based queries:

	where.In.Country("Japan")     // All regions in Japan
	where.In.Country("Germany")   // All regions in Germany
	where.In.City("Tokyo")        // All regions in Tokyo
	where.In.City("Frankfurt")    // All regions in Frankfurt

# Provider Queries

The By namespace provides provider-based region discovery:

	where.On.AWS()      // All Amazon Web Services regions
	where.On.Azure()    // All Microsoft Azure regions
	where.On.GCP()      // All Google Cloud Platform regions
	where.On.Provider("aws")    // Same as where.On.AWS()

# Builder Pattern Queries

For complex filtering requirements, use the Query builder:

	query := where.NewQuery()

	// Chain filters
	results := query.
		InCountry("United States").
		ByProvider("aws").
		ActiveOnly().
		WithinRadius(40.7128, -74.0060, 500). // 500km from NYC
		Exec()

Available filters:

  - InCountry(name string) - Filter by country
  - InCity(name string) - Filter by city
  - InContinent(name string) - Filter by continent
  - InAsia(), InEurope(), InAmericas() - Continent shortcuts
  - ByProvider(name string) - Filter by provider
  - ByAWS(), ByAzure(), ByGCP() - Provider shortcuts
  - ActiveOnly() - Only active regions
  - WithStatus(status Status) - Filter by specific status
  - WithinRadius(lat, lng, km float64) - Geographic proximity
  - SortByDistance(lat, lng float64) - Sort by distance from point
  - Limit(n int) - Limit result count

# Region Data Structure

Each region contains comprehensive metadata:

	type Region struct {
		Code       Code      // Unique region identifier
		Name       string    // Human-readable name
		Provider   string    // Cloud provider name
		Country    string    // Country name
		City       string    // City name
		Continent  string    // Continent name
		Latitude   float64   // Geographic latitude
		Longitude  float64   // Geographic longitude
		Status     Status    // Operational status
		LaunchDate time.Time // When the region became available
	}

Region status values:

	Active     // Fully operational and available
	Deprecated // Being phased out
	Preview    // Limited availability

# Set Operations

Region sets support common operations:

	regions := where.In.Asia()

	// Filtering
	awsRegions := regions.ByProvider("aws")
	japanRegions := regions.ByCountry("Japan")
	activeRegions := regions.ByStatus(where.Active)

	// Geographic operations
	nearbyRegions := regions.WithinRadius(35.6762, 139.6503, 1000) // 1000km from Tokyo
	sortedRegions := regions.SortByDistance(35.6762, 139.6503)

	// Set operations
	combined := set1.Union(set2)
	common := set1.Intersect(set2)
	different := set1.Difference(set2)

	// Utilities
	codes := regions.Codes()           // Extract region codes
	providers := regions.Providers()   // Get unique providers
	countries := regions.Countries()   // Get unique countries

# Error Handling

The package defines standard errors for common failure cases:

	var (
		ErrRegionNotFound   = errors.New("region not found")
		ErrProviderNotFound = errors.New("provider not found")
	)

Use errors.Is() for error checking:

	region, err := where.Is("invalid-region")
	if errors.Is(err, where.ErrRegionNotFound) {
		// Handle unknown region
	}

# Performance Characteristics

  - Provider constants: Zero allocations, sub-nanosecond access
  - Region lookups: O(1) hash table lookups
  - Geographic queries: Pre-indexed for fast filtering
  - Distance calculations: Optimized haversine formula
  - Memory footprint: Efficient data structures with minimal overhead

# Thread Safety

All package functions and types are safe for concurrent use across multiple goroutines.
The internal region registry is read-only after package initialization.

# Examples

Find the closest AWS region to a specific location:

	lat, lng := 35.6762, 139.6503 // Tokyo coordinates
	closest := where.On.AWS().
		SortByDistance(lat, lng).
		First()
	fmt.Printf("Closest AWS region to Tokyo: %s\n", closest.Name)

Get all active regions in Europe:

	europeanRegions := where.In.Europe().ByStatus(where.Active)
	for _, region := range europeanRegions {
		fmt.Printf("%s: %s in %s\n", region.Code, region.Name, region.City)
	}

Check if a region exists and get its details:

	if region, err := where.Is("us-west-2"); err == nil {
		fmt.Printf("Provider: %s\n", region.Provider)
		fmt.Printf("Location: %s, %s\n", region.City, region.Country)
		fmt.Printf("Status: %s\n", region.Status)
		fmt.Printf("Launched: %s\n", region.LaunchDate.Format("2006-01-02"))
	}

Find all GCP regions within 1000km of San Francisco:

	sfLat, sfLng := 37.7749, -122.4194
	nearbyGCP := where.On.GCP().WithinRadius(sfLat, sfLng, 1000)
	fmt.Printf("Found %d GCP regions within 1000km of SF\n", len(nearbyGCP))

# Version Compatibility

This package follows semantic versioning. The current API is stable and
backwards compatibility will be maintained within major versions.

For the latest region data and updates, see: https://github.com/vivaneiona/where
*/
package where
