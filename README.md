# vivaneiona/where

A Go library for discovering, querying, and working with cloud provider region names and codes across different cloud service providers.

```bash
go get github.com/vivaneiona/where@v0.202506.0
```

## Usage Examples

```go
package main

import (
	"fmt"
	"log"

	"github.com/vivaneiona/where"
)

func main() {
	fmt.Println("=== Where Package Usage Examples ===\n")

	// Example 1: Basic region lookup
	basicLookupExample()

	// Example 2: Provider constants (type-safe, zero-allocation)
	providerConstantsExample()

	// Example 3: Geographic queries
	geographicQueriesExample()

	// Example 4: Provider-specific queries
	providerQueriesExample()

	// Example 5: Complex builder pattern queries
	builderPatternExample()

	// Example 6: Distance calculations
	distanceCalculationsExample()

	// Example 7: Region discovery and filtering
	discoveryAndFilteringExample()

	// Example 8: Error handling
	errorHandlingExample()

	// Example 9: Region status and metadata
	regionMetadataExample()

	// Example 10: Practical use cases
	practicalUseCasesExample()
}

func basicLookupExample() {
	fmt.Println("🔍 Basic Region Lookup:")
	// Output:
	//   Region: US East (N. Virginia)
	//   City: Ashburn, United States
	//   Provider: aws
	//   Status: active
	//   Launched: 2006-08-25
	//   Found 3 regions

	// Single region lookup
	region, err := where.Is("us-east-1")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("  Region: %s\n", region.Name)
	fmt.Printf("  City: %s, %s\n", region.City, region.Country)
	fmt.Printf("  Provider: %s\n", region.Provider)
	fmt.Printf("  Status: %s\n", region.Status)
	fmt.Printf("  Launched: %s\n", region.LaunchDate.Format("2006-01-02"))

	// Multiple regions lookup
	regions, err := where.Are("us-east-1", "eu-west-1", "ap-southeast-1")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("  Found %d regions\n", len(regions))
	fmt.Println()
}

func providerConstantsExample() {
	fmt.Println("🏗️ Provider Constants (Type-Safe Access):")
	// Output:
	//   AWS US East 1: us-east-1
	//   AWS US West 2: us-west-2
	//   AWS EU West 1: eu-west-1
	//   GCP US Central 1: us-central1
	//   GCP Europe West 1: europe-west1
	//   GCP Asia South 1: asia-south1
	//   Azure East US: eastus
	//   Azure West Europe: westeurope
	//   Azure Southeast Asia: southeastasia

	// AWS regions
	fmt.Printf("  AWS US East 1: %s\n", where.AWS.USEast1)
	fmt.Printf("  AWS US West 2: %s\n", where.AWS.USWest2)
	fmt.Printf("  AWS EU West 1: %s\n", where.AWS.EUWest1)

	// Google Cloud Platform regions
	fmt.Printf("  GCP US Central 1: %s\n", where.GCP.USCentral1)
	fmt.Printf("  GCP Europe West 1: %s\n", where.GCP.EuropeWest1)
	fmt.Printf("  GCP Asia South 1: %s\n", where.GCP.AsiaSouth1)

	// Azure regions
	fmt.Printf("  Azure East US: %s\n", where.Azure.EastUS)
	fmt.Printf("  Azure West Europe: %s\n", where.Azure.WestEurope)
	fmt.Printf("  Azure Southeast Asia: %s\n", where.Azure.SoutheastAsia)
	fmt.Println()
}

func geographicQueriesExample() {
	fmt.Println("🌍 Geographic Queries:")
	// Output:
	//   Asia: 35 regions
	//   Europe: 31 regions
	//   Americas: 30 regions
	//   Japan: 6 regions
	//   United States: 21 regions
	//   Germany: 4 regions
	//   Tokyo: 3 regions

	// Continent-based queries using namespace
	asianRegions := where.In.Asia()
	europeanRegions := where.In.Europe()
	americanRegions := where.In.Americas()

	fmt.Printf("  Asia: %d regions\n", len(asianRegions))
	fmt.Printf("  Europe: %d regions\n", len(europeanRegions))
	fmt.Printf("  Americas: %d regions\n", len(americanRegions))

	// Country-specific queries using namespace
	japanRegions := where.In.Country("Japan")
	usRegions := where.In.Country("United States")
	germanyRegions := where.In.Country("Germany")

	fmt.Printf("  Japan: %d regions\n", len(japanRegions))
	fmt.Printf("  United States: %d regions\n", len(usRegions))
	fmt.Printf("  Germany: %d regions\n", len(germanyRegions))

	// City-specific queries using namespace
	tokyoRegions := where.In.City("Tokyo")
	fmt.Printf("  Tokyo: %d regions\n", len(tokyoRegions))
	fmt.Println()
}

func providerQueriesExample() {
	fmt.Println("☁️ Provider-Specific Queries:")
	// Output:
	//   AWS: 28 regions
	//   Azure: 31 regions
	//   GCP: 33 regions
	//   AWS (alternative): 28 regions

	// Provider-based queries using namespace
	awsRegions := where.By.AWS()
	azureRegions := where.By.Azure()
	gcpRegions := where.By.GCP()

	fmt.Printf("  AWS: %d regions\n", len(awsRegions))
	fmt.Printf("  Azure: %d regions\n", len(azureRegions))
	fmt.Printf("  GCP: %d regions\n", len(gcpRegions))

	// Alternative provider query
	awsRegions2 := where.By.Provider("aws")
	fmt.Printf("  AWS (alternative): %d regions\n", len(awsRegions2))
	fmt.Println()
}

func builderPatternExample() {
	fmt.Println("🔧 Builder Pattern Queries:")
	// Output:
	//   Active AWS regions in Japan: 2
	//     - ap-northeast-3 (Asia Pacific (Osaka)) in Osaka
	//     - ap-northeast-1 (Asia Pacific (Tokyo)) in Tokyo
	//   AWS regions within 1000km of Tokyo: 2
	//   Top 3 European GCP regions closest to Frankfurt:
	//     1. europe-west3 (Frankfurt) in Frankfurt
	//     2. europe-west6 (Zurich) in Zurich
	//     3. europe-west1 (Belgium) in St. Ghislain

	// Complex query: AWS regions in Japan that are active
	awsJapanActive := where.NewQuery().
		InCountry("Japan").
		ByProvider("aws").
		ActiveOnly().
		Exec()

	fmt.Printf("  Active AWS regions in Japan: %d\n", len(awsJapanActive))
	for _, region := range awsJapanActive {
		fmt.Printf("    - %s (%s) in %s\n", region.Code, region.Name, region.City)
	}

	// Geographic proximity query: AWS regions within 1000km of Tokyo
	tokyoLat, tokyoLng := 35.6762, 139.6503
	nearTokyo := where.NewQuery().
		ByProvider("aws").
		Near(tokyoLat, tokyoLng, 1000).
		Exec()

	fmt.Printf("  AWS regions within 1000km of Tokyo: %d\n", len(nearTokyo))

	// European GCP regions, sorted by distance from Frankfurt
	frankfurtLat, frankfurtLng := 50.1109, 8.6821
	europeanGCP := where.NewQuery().
		InContinent("Europe").
		ByProvider("gcp").
		SortByDistance(frankfurtLat, frankfurtLng).
		Limit(3).
		Exec()

	fmt.Printf("  Top 3 European GCP regions closest to Frankfurt:\n")
	for i, region := range europeanGCP {
		fmt.Printf("    %d. %s (%s) in %s\n", i+1, region.Code, region.Name, region.City)
	}
	fmt.Println()
}

func distanceCalculationsExample() {
	fmt.Println("📏 Distance Calculations:")
	// Output:
	//   Distance between us-east-1 and us-west-2: 3779.26 km
	//   Closest AWS region to San Francisco: us-west-1 (US West (N. California))
	//   Regions within 500km of NYC: 2

	// Calculate distance between two regions
	distance, err := where.Distance("us-east-1", "us-west-2")
	if err != nil {
		log.Printf("Error calculating distance: %v", err)
		return
	}
	fmt.Printf("  Distance between us-east-1 and us-west-2: %.2f km\n", distance)

	// Find closest AWS region to a specific location (San Francisco)
	sfLat, sfLng := 37.7749, -122.4194
	awsRegions := where.By.AWS()
	awsRegions.SortByDistance(sfLat, sfLng)
	if len(awsRegions) > 0 {
		closestToSF := awsRegions[0]
		fmt.Printf("  Closest AWS region to San Francisco: %s (%s)\n",
			closestToSF.Code, closestToSF.Name)
	}

	// Regions within 500km of New York City
	nycLat, nycLng := 40.7128, -74.0060
	nearNYC := where.Near(nycLat, nycLng, 500)
	fmt.Printf("  Regions within 500km of NYC: %d\n", len(nearNYC))
	fmt.Println()
}

func discoveryAndFilteringExample() {
	fmt.Println("🔎 Discovery and Filtering:")
	// Output:
	//   Available providers: [gcp alibaba yandex aws azure]
	//   Available countries: 31 total
	//   Available continents: [North America Oceania Asia Europe South America Africa]
	//   Active regions: 104
	//   Preview regions: 0
	//   Deprecated regions: 0

	// Discover available providers
	providers := where.Providers()
	fmt.Printf("  Available providers: %v\n", providers)

	// Discover available countries
	countries := where.Countries()
	fmt.Printf("  Available countries: %d total\n", len(countries))

	// Discover available continents
	continents := where.Continents()
	fmt.Printf("  Available continents: %v\n", continents)

	// Filter by status
	activeRegions := where.ActiveRegions()
	previewRegions := where.PreviewRegions()
	deprecatedRegions := where.DeprecatedRegions()

	fmt.Printf("  Active regions: %d\n", len(activeRegions))
	fmt.Printf("  Preview regions: %d\n", len(previewRegions))
	fmt.Printf("  Deprecated regions: %d\n", len(deprecatedRegions))
	fmt.Println()
}

func errorHandlingExample() {
	fmt.Println("⚠️ Error Handling:")
	// Output:
	//   Expected error for invalid region: region not found: invalid-region
	//   ✓ us-east-1 exists
	//   ✗ invalid-region does not exist
	//   ✓ AWS provider exists
	//   ✗ invalid-provider does not exist

	// Try to lookup an invalid region
	_, err := where.Is("invalid-region")
	if err != nil {
		fmt.Printf("  Expected error for invalid region: %v\n", err)
	}

	// Check if region exists
	if where.Exists("us-east-1") {
		fmt.Printf("  ✓ us-east-1 exists\n")
	}

	if !where.Exists("invalid-region") {
		fmt.Printf("  ✗ invalid-region does not exist\n")
	}

	// Check if provider exists
	if where.HasProvider("aws") {
		fmt.Printf("  ✓ AWS provider exists\n")
	}

	if !where.HasProvider("invalid-provider") {
		fmt.Printf("  ✗ invalid-provider does not exist\n")
	}
	fmt.Println()
}

func regionMetadataExample() {
	fmt.Println("📊 Region Metadata:")
	// Output:
	//   Region Code: us-east-1
	//   Name: US East (N. Virginia)
	//   Provider: aws
	//   Country: United States
	//   City: Ashburn
	//   Continent: North America
	//   Coordinates: 38.9047, -77.0164
	//   Status: active
	//   Launch Date: August 25, 2006
	//   Is Active: true
	//   Is within 1000km of NYC: true

	// Get detailed region information
	region, _ := where.Is("us-east-1")

	fmt.Printf("  Region Code: %s\n", region.Code)
	fmt.Printf("  Name: %s\n", region.Name)
	fmt.Printf("  Provider: %s\n", region.Provider)
	fmt.Printf("  Country: %s\n", region.Country)
	fmt.Printf("  City: %s\n", region.City)
	fmt.Printf("  Continent: %s\n", region.Continent)
	fmt.Printf("  Coordinates: %.4f, %.4f\n", region.Latitude, region.Longitude)
	fmt.Printf("  Status: %s\n", region.Status)
	fmt.Printf("  Launch Date: %s\n", region.LaunchDate.Format("January 2, 2006"))
	fmt.Printf("  Is Active: %t\n", region.IsActive())

	// Check if region is near a location
	nycLat, nycLng := 40.7128, -74.0060
	isNearNYC := region.IsNear(nycLat, nycLng, 1000) // within 1000km
	fmt.Printf("  Is within 1000km of NYC: %t\n", isNearNYC)
	fmt.Println()
}

func practicalUseCasesExample() {
	fmt.Println("💡 Practical Use Cases:")
	// Output:
	//   1. Multi-region deployment planning:
	//      Primary: us-east-1
	//      Backup candidates in other continents: 22
	//   2. EU data residency compliance:
	//      EU-compliant AWS regions: 8
	//      - eu-west-2 (Europe (London)) in United Kingdom
	//      - eu-south-1 (Europe (Milan)) in Italy
	//      - eu-west-1 (Europe (Ireland)) in Ireland
	//   3. Latency optimization for global users:
	//      Closest AWS to Singapore: ap-southeast-1 (Singapore)
	//      Closest AWS to Sydney: ap-southeast-2 (Sydney)
	//      Closest AWS to London: eu-west-2 (London)

	// Use case 1: Multi-region deployment planning
	fmt.Println("  1. Multi-region deployment planning:")
	primaryRegion := "us-east-1"

	// Find backup regions in different continents
	backupCandidates := where.NewQuery().
		ByProvider("aws").
		ActiveOnly().
		Exec().
		Filter(func(r where.Region) bool {
			primary, _ := where.Is(where.Code(primaryRegion))
			return r.Continent != primary.Continent
		})

	fmt.Printf("     Primary: %s\n", primaryRegion)
	fmt.Printf("     Backup candidates in other continents: %d\n", len(backupCandidates))

	// Use case 2: Compliance and data residency
	fmt.Println("  2. EU data residency compliance:")
	euRegions := where.NewQuery().
		InContinent("Europe").
		ByProvider("aws").
		ActiveOnly().
		Exec()

	fmt.Printf("     EU-compliant AWS regions: %d\n", len(euRegions))
	for _, region := range euRegions[:3] { // Show first 3
		fmt.Printf("     - %s (%s) in %s\n", region.Code, region.Name, region.Country)
	}

	// Use case 3: Latency optimization
	fmt.Println("  3. Latency optimization for global users:")
	userLocations := map[string][2]float64{
		"London":    {51.5074, -0.1278},
		"Singapore": {1.3521, 103.8198},
		"Sydney":    {-33.8688, 151.2093},
	}

	for city, coords := range userLocations {
		awsRegions := where.By.AWS()
		awsRegions.SortByDistance(coords[0], coords[1])
		if len(awsRegions) > 0 {
			closest := awsRegions[0]
			fmt.Printf("     Closest AWS to %s: %s (%s)\n",
				city, closest.Code, closest.City)
		}
	}
	fmt.Println()
}
```

## License

MIT License - see [LICENSE](LICENSE) file for details.
