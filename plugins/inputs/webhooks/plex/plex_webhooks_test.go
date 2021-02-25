package plex

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/influxdata/telegraf/testutil"
	"github.com/stretchr/testify/require"
)

func postWebhooks(rb *PlexWebhook, eventBody string) *httptest.ResponseRecorder {
	values := map[string]io.Reader{
		"payload": strings.NewReader(eventBody),
	}
	writer, b, _ := BuildFormData(values)
	req, _ := http.NewRequest("POST", "/plex", &b)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	w.Code = 500

	rb.eventHandler(w, req)

	return w
}

func BuildFormData(values map[string]io.Reader) (w *multipart.Writer, b bytes.Buffer, err error) {
	// Prepare a form that you will submit to that URL.
	w = multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add other fields
		fw, err = w.CreateFormField(key)
		if err != nil {
			return nil, b, err
		}
		_, err = io.Copy(fw, r)
		if err != nil {
			return nil, b, err
		}

	}
	w.Close()
	return w, b, nil
}

func TestPlexMediaPlayEvent(t *testing.T) {
	var acc testutil.Accumulator
	p := &PlexWebhook{Path: "/plex", acc: &acc}
	resp := postWebhooks(p, MediaPlayEventJSON())
	require.Equal(t, http.StatusOK, resp.Code)

	expectedTags := map[string]string{
		"account_title":                  "elan",
		"event":                          "media.play",
		"is_owner":                       "true",
		"is_player_local":                "true",
		"is_user":                        "true",
		"metadata_grandparent_title":     "Stephen Stills",
		"metadata_library_section_title": "Music",
		"metadata_library_section_type":  "artist",
		"metadata_parent_title":          "Stephen Stills",
		"metadata_title":                 "Love The One You're With",
		"metadata_type":                  "track",
		"player_title":                   "Plex Web (Safari)",
		"server_title":                   "Office",
	}
	expectedFields := map[string]interface{}{
		"counter": int64(1),
	}
	acc.AssertContainsTaggedFields(t, "plex", expectedFields, expectedTags)
}

func TestRateEvent(t *testing.T) {
	var acc testutil.Accumulator
	p := &PlexWebhook{Path: "/plex", acc: &acc}
	resp := postWebhooks(p, RateEventJSON())
	require.Equal(t, http.StatusOK, resp.Code)

	expectedTags := map[string]string{
		"account_title":                  "username",
		"event":                          "media.rate",
		"is_owner":                       "true",
		"is_player_local":                "true",
		"is_user":                        "true",
		"metadata_content_rating":        "TVPG",
		"metadata_library_section_title": "Movies",
		"metadata_library_section_type":  "movie",
		"metadata_title":                 "The Harder They Fall",
		"metadata_type":                  "movie",
		"metadata_year":                  "1956",
		"player_title":                   "Chrome",
		"server_title":                   "Office",
	}
	expectedFields := map[string]interface{}{
		"counter": int64(1),
		"rating":  int64(10),
	}
	acc.AssertContainsTaggedFields(t, "plex", expectedFields, expectedTags)
}

func TestPlexMediaPhotoEvent(t *testing.T) {
	var acc testutil.Accumulator
	p := &PlexWebhook{Path: "/plex", acc: &acc}
	resp := postWebhooks(p, MediaPhotoEventJSON())
	require.Equal(t, http.StatusOK, resp.Code)

	expectedTags := map[string]string{
		"account_title":                  "username",
		"event":                          "media.play",
		"is_owner":                       "true",
		"is_player_local":                "true",
		"is_user":                        "true",
		"metadata_library_section_title": "Photos",
		"metadata_library_section_type":  "photo",
		"metadata_title":                 "Screen Shot 2021-02-24 at 3.26.14 PM",
		"metadata_type":                  "photo",
		"metadata_year":                  "2021",
		"player_title":                   "Chrome",
		"server_title":                   "Office",
	}
	expectedFields := map[string]interface{}{
		"counter": int64(1),
	}
	acc.AssertContainsTaggedFields(t, "plex", expectedFields, expectedTags)
}

func TestPlexMediaTrailerEvent(t *testing.T) {
	var acc testutil.Accumulator
	p := &PlexWebhook{Path: "/plex", acc: &acc}
	resp := postWebhooks(p, MediaPlayTrailerEventJSON())
	require.Equal(t, http.StatusOK, resp.Code)

	expectedTags := map[string]string{
		"account_title":    "username",
		"event":            "media.play",
		"is_owner":         "true",
		"is_player_local":  "true",
		"is_user":          "true",
		"metadata_subtype": "trailer",
		"metadata_title":   "Tenet",
		"metadata_type":    "clip",
		"metadata_year":    "2019",
		"player_title":     "Chrome",
		"server_title":     "Office",
	}
	expectedFields := map[string]interface{}{
		"counter": int64(1),
	}
	acc.AssertContainsTaggedFields(t, "plex", expectedFields, expectedTags)
}

func TestPlexMediaStopEvent(t *testing.T) {
	var acc testutil.Accumulator
	p := &PlexWebhook{Path: "/plex", acc: &acc}
	resp := postWebhooks(p, MediaStopEventJSON())
	require.Equal(t, http.StatusOK, resp.Code)

	expectedTags := map[string]string{
		"account_title":                  "username",
		"event":                          "media.stop",
		"is_owner":                       "true",
		"is_player_local":                "true",
		"is_user":                        "true",
		"metadata_content_rating":        "R",
		"metadata_director":              "Gore Verbinski",
		"metadata_library_section_title": "Movies",
		"metadata_library_section_type":  "movie",
		"metadata_studio":                "20th Century Fox",
		"metadata_title":                 "EVO",
		"metadata_type":                  "movie",
		"metadata_year":                  "2017",
		"player_title":                   "Chrome",
		"server_title":                   "Office",
	}
	expectedFields := map[string]interface{}{
		"counter": int64(1),
	}
	acc.AssertContainsTaggedFields(t, "plex", expectedFields, expectedTags)
}

func TestPlexMediaScrobbleEvent(t *testing.T) {
	var acc testutil.Accumulator
	p := &PlexWebhook{Path: "/plex", acc: &acc}
	resp := postWebhooks(p, MediaScrobbleEvent())
	require.Equal(t, http.StatusOK, resp.Code)

	expectedTags := map[string]string{
		"account_title":                  "username",
		"event":                          "media.scrobble",
		"is_owner":                       "true",
		"is_player_local":                "true",
		"is_user":                        "true",
		"metadata_content_rating":        "TVPG",
		"metadata_library_section_title": "Movies",
		"metadata_library_section_type":  "movie",
		"metadata_title":                 "3:10 to Yuma",
		"metadata_type":                  "movie",
		"metadata_year":                  "1957",
		"player_title":                   "Chrome",
		"server_title":                   "Office",
	}
	expectedFields := map[string]interface{}{
		"counter": int64(1),
	}
	acc.AssertContainsTaggedFields(t, "plex", expectedFields, expectedTags)
}

func TestPlexMediaMusicEvent(t *testing.T) {
	var acc testutil.Accumulator
	p := &PlexWebhook{Path: "/plex", acc: &acc}
	resp := postWebhooks(p, MediaMusicEventJSON())
	require.Equal(t, http.StatusOK, resp.Code)

	expectedTags := map[string]string{
		"account_title":                  "username",
		"event":                          "media.play",
		"is_owner":                       "true",
		"is_player_local":                "true",
		"is_user":                        "true",
		"metadata_grandparent_title":     "Various Artists",
		"metadata_library_section_title": "Music",
		"metadata_library_section_type":  "artist",
		"metadata_original_title":        "Britney Spears",
		"metadata_parent_title":          "RTL: I Like the 90’s",
		"metadata_parent_year":           "2014",
		"metadata_title":                 "Oops!…I Did It Again (album version)",
		"metadata_type":                  "track",
		"player_title":                   "Chrome",
		"server_title":                   "Office",
	}
	expectedFields := map[string]interface{}{
		"counter": int64(1),
	}
	acc.AssertContainsTaggedFields(t, "plex", expectedFields, expectedTags)
}
