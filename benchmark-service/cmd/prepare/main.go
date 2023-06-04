package main

import (
	"flag"

	v2023 "github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/cmd/prepare/v2023"
)

var (
	v2023TagCSVPath    = flag.String("v2023-csv-path", "csv/tag.csv", "path to v2023 csv file")
	v2023GeotagCSVPath = flag.String("v2023-geotag-csv-path", "csv/geotag.csv", "path to v2023 geotag csv file")
	v2023DestPath      = flag.String("v2023-dest-path", "data/v2023.json", "path to v2023 dest file")
)

func main() {
	flag.Parse()
	if err := v2023.Build(*v2023TagCSVPath, *v2023GeotagCSVPath, *v2023DestPath); err != nil {
		panic(err)
	}
}
