# vivaneiona/where

A Go library for discovering, querying, and working with cloud provider region names and codes across different cloud service providers.

```bash
go get github.com/vivaneiona/where@v0.202506.0
```

## Project Status

This library is currently in **active development**.

```go
package main

import (
    "fmt"
    "log"

    "github.com/vivaneiona/where"
)

func main() {
    // Provider constants available for all major clouds
    fmt.Printf("%s\n", where.GCP.USCentral1)
    fmt.Printf("%s\n", where.AWS.USEast1)
 
    // Use provider constants for type-safe region access
    region, err := where.Is("us-east-1")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Region: %s in %s\n", region.Name, region.City)

    // Provider-specific queries
    awsRegions := where.By.AWS()
    azureRegions := where.By.Azure()
    gcpRegions := where.By.GCP()

    fmt.Printf("AWS: %d, Azure: %d, GCP: %d regions\n", 
        len(awsRegions), len(azureRegions), len(gcpRegions))

    // Geographic queries
    asianRegions := where.In.Asia()
    usRegions := where.In.US()
    fmt.Printf("Asia: %d, US: %d regions\n", len(asianRegions), len(usRegions))

    // Builder pattern for complex queries
    results := where.NewQuery().
        InCountry("Japan").
        ByProvider("aws").
        ActiveOnly().
        Exec()
    
    fmt.Printf("AWS regions in Japan: %d\n", len(results))
}
```

## License

MIT License - see [LICENSE](LICENSE) file for details.
