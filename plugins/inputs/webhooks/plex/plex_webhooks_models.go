package plex

import (
	"fmt"
	"time"

	"github.com/hekmon/plexwebhooks"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
)

const meas = "plex_webhooks"

type PlexWebhookEvent struct {
	Event    string   `json:"event"`
	User     bool     `json:"user"`
	Owner    bool     `json:"owner"`
	Account  Account  `json:"Account"`
	Server   Server   `json:"Server"`
	Player   Player   `json:"Player"`
	Metadata Metadata `json:"Metadata"`
}

type Account struct {
	ID    int    `json:"id"`
	Thumb string `json:"thumb"`
	Title string `json:"title"`
}

type Server struct {
	Title string `json:"title"`
	UUID  string `json:"uuid"`
}

type Player struct {
	Local         bool   `json:"local"`
	PublicAddress string `json:"PublicAddress"`
	Title         string `json:"title"`
	UUID          string `json:"uuid"`
}

type Metadata struct {
	Directors            []Director `json:"Director"`
	Writers              []Writer   `json:"Writer"`
	LibrarySectionType   string     `json:"librarySectionType"`
	RatingKey            string     `json:"ratingKey"`
	Key                  string     `json:"key"`
	ParentRatingKey      string     `json:"parentRatingKey"`
	GrandparentRatingKey string     `json:"grandparentRatingKey"`
	GUID                 string     `json:"guid"`
	LibrarySectionID     int        `json:"librarySectionID"`
	MediaType            string     `json:"type"`
	Title                string     `json:"title"`
	GrandparentKey       string     `json:"grandparentKey"`
	ParentKey            string     `json:"parentKey"`
	GrandparentTitle     string     `json:"grandparentTitle"`
	ParentTitle          string     `json:"parentTitle"`
	Summary              string     `json:"summary"`
	Index                int        `json:"index"`
	ParentIndex          int        `json:"parentIndex"`
	RatingCount          int        `json:"ratingCount"`
	Thumb                string     `json:"thumb"`
	Art                  string     `json:"art"`
	ParentThumb          string     `json:"parentThumb"`
	GrandparentThumb     string     `json:"grandparentThumb"`
	GrandparentArt       string     `json:"grandparentArt"`
	AddedAt              int        `json:"addedAt"`
	UpdatedAt            int        `json:"updatedAt"`
}

type Director struct {
	Filter string `json:"filter"`
	ID     int    `json:"id"`
	Tag    string `json:"tag"`
}

type Writer struct {
	Filter string `json:"filter"`
	ID     int    `json:"id"`
	Tag    string `json:"tag"`
}

func NewMetric(s *plexwebhooks.Payload) (telegraf.Metric, error) {

	t := map[string]string{
		"is_user":                fmt.Sprintf("%v", s.User),
		"is_owner":               fmt.Sprintf("%v", s.Owner),
		"user_id":                fmt.Sprintf("%v", s.Account.ID),
		"user_name":              s.Account.Title,
		"server_title":           s.Server.Title,
		"is_player_local":        fmt.Sprintf("%v", s.Player.Local),
		"library_selection_type": string(s.Metadata.LibrarySectionType),
		"media_type":             string(s.Metadata.Type),
		"grandparent_key":        s.Metadata.GrandparentKey,
		"parent_key":             s.Metadata.ParentKey,
		"grandparent_title":      s.Metadata.GrandparentTitle,
		"parent_title":           s.Metadata.ParentTitle,
		"parent_index":           fmt.Sprintf("%v", s.Metadata.ParentIndex),
	}
	f := map[string]interface{}{
		"event":                s.Event,
		"player_title":         s.Player.Title,
		"player_public_ip":     s.Player.PublicAddress,
		"rating_count":         s.Metadata.RatingCount,
		"added_at":             s.Metadata.AddedAt,
		"updated_at":           s.Metadata.UpdatedAt,
		"summary":              s.Metadata.Summary,
		"title":                s.Metadata.Title,
		"index":                s.Metadata.Index,
		"library_selection_id": s.Metadata.LibrarySectionID,
	}
	return metric.New(meas, t, f, time.Now())
}
