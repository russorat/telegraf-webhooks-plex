package plex

import (
	"fmt"
	"time"

	"github.com/hekmon/plexwebhooks"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
)

const meas = "plex_webhooks"

func NewMetric(s *plexwebhooks.Payload) (telegraf.Metric, error) {
	t := map[string]string{
		"is_user":                 fmt.Sprintf("%v", s.User),
		"is_owner":                fmt.Sprintf("%v", s.Owner),
		"server_title":            s.Server.Title,
		"is_player_local":         fmt.Sprintf("%v", s.Player.Local),
		"library_selection_type":  string(s.Metadata.LibrarySectionType),
		"library_selection_title": s.Metadata.LibrarySectionTitle,
		"media_type":              string(s.Metadata.Type),
		"grandparent_title":       s.Metadata.GrandparentTitle,
		"parent_title":            s.Metadata.ParentTitle,
		"event":                   fmt.Sprintf("%v", s.Event),
		"user_name":               s.Account.Title,
		"player_title":            s.Player.Title,
		"title":                   s.Metadata.Title,
	}
	f := map[string]interface{}{
		"view_count": s.Metadata.ViewCount,
	}
	if s.Event == plexwebhooks.EventTypeRate {
		f["rating"] = s.Rating
	}
	if s.Metadata.Type == plexwebhooks.MediaTypeMovie {
		t["content_rating"] = s.Metadata.ContentRating
		t["chapter_source"] = s.Metadata.ChapterSource
		if len(s.Metadata.Genre) > 0 {
			t["genre"] = fmt.Sprintf("%v", s.Metadata.Genre[0].Tag)
		}
		if len(s.Metadata.Director) > 0 {
			t["director"] = fmt.Sprintf("%v", s.Metadata.Director[0].Tag)
		}
		if len(s.Metadata.Country) > 0 {
			t["country"] = fmt.Sprintf("%v", s.Metadata.Country[0].Tag)
		}
		if len(s.Metadata.Collection) > 0 {
			t["collection"] = fmt.Sprintf("%v", s.Metadata.Collection[0].Tag)
		}
		t["studio"] = s.Metadata.Studio
		t["year"] = fmt.Sprintf("%v", s.Metadata.Year)
	} else if s.Metadata.Type == plexwebhooks.MediaTypeEpisode {
		t["content_rating"] = s.Metadata.ContentRating
		t["year"] = fmt.Sprintf("%v", s.Metadata.Year)
		t["chapter_source"] = s.Metadata.ChapterSource
		if len(s.Metadata.Director) > 0 {
			t["director"] = fmt.Sprintf("%v", s.Metadata.Director[0].Tag)
		}
	} else if s.Metadata.Type == plexwebhooks.MediaTypeTrack {
		if len(s.Metadata.Mood) > 0 {
			t["mood"] = fmt.Sprintf("%v", s.Metadata.Mood[0].Tag)
		}

		f["rating_count"] = s.Metadata.RatingCount
	}
	return metric.New(meas, t, f, time.Now())
}
