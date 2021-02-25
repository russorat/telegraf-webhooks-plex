package plex

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/hekmon/plexwebhooks"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
)

const MetricName = "plex"

func NewMetric(s *plexwebhooks.Payload) (telegraf.Metric, error) {
	f := map[string]interface{}{
		"counter": 1,
	}
	t := map[string]string{
		"is_user":  fmt.Sprintf("%v", s.User),
		"is_owner": fmt.Sprintf("%v", s.Owner),
		"event":    fmt.Sprintf("%v", s.Event),
	}

	if s.Server.Title != "" {
		t["server_title"] = s.Server.Title
	}

	if s.Account.Title != "" {
		t["account_title"] = s.Account.Title
	}

	if s.Player.Title != "" {
		t["is_player_local"] = fmt.Sprintf("%v", s.Player.Local)
		t["player_title"] = s.Player.Title
	}

	if s.Event == plexwebhooks.EventTypeRate {
		f["rating"] = s.Rating
	}

	if s.Metadata.Type != "" {
		if s.Metadata.ContentRating != "" {
			t["metadata_content_rating"] = s.Metadata.ContentRating
		}
		if len(s.Metadata.Director) > 0 {
			t["metadata_director"] = s.Metadata.Director[0].Tag
		}
		if s.Metadata.GrandparentTitle != "" {
			t["metadata_grandparent_title"] = string(s.Metadata.GrandparentTitle)
		}
		if s.Metadata.LibrarySectionTitle != "" {
			t["metadata_library_section_title"] = string(s.Metadata.LibrarySectionTitle)
			t["metadata_library_section_type"] = string(s.Metadata.LibrarySectionType)
		}
		if s.Metadata.OriginalTitle != "" {
			t["metadata_original_title"] = s.Metadata.OriginalTitle
		}
		if s.Metadata.ParentTitle != "" {
			t["metadata_parent_title"] = s.Metadata.ParentTitle
		}
		if s.Metadata.Studio != "" {
			t["metadata_studio"] = s.Metadata.Studio
		}
		if s.Metadata.Title != "" {
			t["metadata_title"] = s.Metadata.Title
		}
		t["metadata_type"] = string(s.Metadata.Type)
		if s.Metadata.SubType != "" {
			t["metadata_subtype"] = string(s.Metadata.SubType)
		}
		if s.Metadata.Year != 0 {
			t["metadata_year"] = fmt.Sprintf("%v", s.Metadata.Year)
		}
		if s.Metadata.ParentYear != 0 {
			t["metadata_parent_year"] = fmt.Sprintf("%v", s.Metadata.ParentYear)
		}

		// v := reflect.ValueOf(s.Metadata)
		// typeOfS := v.Type()
		// for i := 0; i < v.NumField(); i++ {
		// 	var fieldType = reflect.TypeOf(v.Field(i))
		// 	if fieldType.String() == "string" {
		// 		if v.Field(i).Interface() != "" {
		// 			f["metadata_"+ToSnakeCase(typeOfS.Field(i).Name)] = v.Field(i).Interface()
		// 		}
		// 	} else if fieldType == reflect.TypeOf([]plexwebhooks.MetadataItem{}) {
		// 		print("something")
		// 	}
		// 	fmt.Printf("Field: %s\tValue: %v\tType: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface(), v.Field(i).Type())
		// }
	}
	return metric.New(MetricName, t, f, time.Now())
}

var matchFirstCap = regexp.MustCompile("([A-Z])([A-Z][a-z])")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase converts the provided string to snake_case.
// Based on https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
func ToSnakeCase(input string) string {
	output := matchFirstCap.ReplaceAllString(input, "${1}_${2}")
	output = matchAllCap.ReplaceAllString(output, "${1}_${2}")
	output = strings.ReplaceAll(output, "-", "_")
	return strings.ToLower(output)
}
