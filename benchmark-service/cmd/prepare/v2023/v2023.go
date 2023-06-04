package v2023

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Tag struct {
	id   uint64
	name string
}

type Geotag struct {
	id        uint64
	elapsed   uint64
	latitude  float64
	longitude float64
	farmNum   int8   // http://farm{farmNum}.flickr...
	directory string // /8237/8520927781_4f86a7a3b1.jpg
}

type TagJson struct {
	TagName string    `json:"tag_name"`
	Geotags []Geotag2 `json:"geotags"`
}

type Geotag2 struct {
	Elapsed   uint64  `json:"elapsed"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	FarmNum   int8    `json:"farm_num"`  // http://farm{farmNum}.flickr...
	Directory string  `json:"directory"` // /8237/8520927781_4f86a7a3b1.jpg
}

func Build(tagPath, geotagPath, destPath string) error {
	if _, err := os.Stat(destPath); err == nil {
		log.Println("skip build")
		return nil
	}

	tags, err := LoadTags(tagPath)
	if err != nil {
		return err
	}
	geotags, err := LoadGeotags(geotagPath)
	if err != nil {
		return err
	}
	log.Println("load done")

	tagmap := map[string][]Geotag2{}
	for _, tag := range tags {
		geotagIndex := sort.Search(len(geotags), func(i2 int) bool {
			return geotags[i2].id >= tag.id
		})
		if _, ok := tagmap[tag.name]; !ok {
			tagmap[tag.name] = make([]Geotag2, 0)
		}
		if geotagIndex == len(geotags) {
			return fmt.Errorf("tag id %d not found", tag.id)
		}
		tagmap[tag.name] = append(tagmap[tag.name], Geotag2{
			Elapsed:   geotags[geotagIndex].elapsed,
			Latitude:  geotags[geotagIndex].latitude,
			Longitude: geotags[geotagIndex].longitude,
			FarmNum:   geotags[geotagIndex].farmNum,
			Directory: geotags[geotagIndex].directory,
		})
	}
	log.Println("search done")

	tagJsonRoot := []TagJson{}
	for tagName, geotags := range tagmap {
		sort.Slice(geotags, func(i, j int) bool {
			return geotags[i].Elapsed > geotags[j].Elapsed
		})
		geotags2 := geotags
		if len(geotags) > 100 {
			geotags2 = geotags2[:100]
		}
		tagJsonRoot = append(tagJsonRoot, TagJson{
			TagName: tagName,
			Geotags: geotags2,
		})
	}
	sort.Slice(tagJsonRoot, func(i, j int) bool {
		return tagJsonRoot[i].TagName < tagJsonRoot[j].TagName
	})
	log.Println("sort done")

	jsonFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	if err := json.NewEncoder(jsonFile).Encode(tagJsonRoot); err != nil {
		return err
	}

	return nil
}

func LoadTags(name string) ([]*Tag, error) {
	tagFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer tagFile.Close()

	tagsc := bufio.NewScanner(tagFile)
	tags := []*Tag{}
	for tagsc.Scan() {
		tokens := strings.Split(strings.TrimSpace(tagsc.Text()), ",")
		id, _ := strconv.ParseUint(tokens[0], 10, 64)
		tag := tokens[1]

		// 空白のタグを除去
		if tag == "" {
			continue
		}
		// マルチバイト文字が含まれるタグを除去
		// if len(tag) != len([]rune(tag)) {
		// 	continue
		// }

		tags = append(tags, &Tag{id, tag})
	}

	return tags, nil
}

func LoadGeotags(name string) ([]*Geotag, error) {
	geotagFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer geotagFile.Close()

	geotags := make([]*Geotag, 0)
	geotagsc := bufio.NewScanner(geotagFile)
	for geotagsc.Scan() {
		tokens := strings.Split(strings.TrimSpace(geotagsc.Text()), ",")
		id, _ := strconv.ParseUint(tokens[0], 10, 64)
		baseDate := time.Date(2012, time.January, 1, 0, 0, 0, 0, time.UTC)
		date, err := time.Parse("2006-01-02 15:04:05", strings.Trim(tokens[1], "\""))
		if err != nil {
			return nil, err
		}
		latitude, _ := strconv.ParseFloat(tokens[2], 64)
		longitude, _ := strconv.ParseFloat(tokens[3], 64)
		var farmNum int8
		var directory string
		fmt.Sscanf(tokens[4], "http://farm%d.static.flickr.com%s", &farmNum, &directory)
		geotags = append(geotags, &Geotag{
			id:        id,
			elapsed:   uint64(date.Sub(baseDate).Seconds()),
			latitude:  latitude,
			longitude: longitude,
			farmNum:   farmNum,
			directory: directory,
		})
	}

	sort.Slice(geotags, func(i, j int) bool {
		return geotags[i].id < geotags[j].id
	})

	return geotags, nil
}
