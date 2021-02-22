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
	t.Parallel()
	var acc testutil.Accumulator
	p := &PlexWebhook{Path: "/plex", acc: &acc}
	resp := postWebhooks(p, PlexMediaPlayEventJSON())
	require.Equal(t, http.StatusOK, resp.Code)

	expectedTags := map[string]string{
		"event":                 "media.play",
		"grandparent_title":     "Stephen Stills",
		"is_owner":              "true",
		"is_player_local":       "true",
		"is_user":               "true",
		"library_section_title": "Music",
		"library_section_type":  "artist",
		"media_type":            "track",
		"parent_title":          "Stephen Stills",
		"player_title":          "Plex Web (Safari)",
		"server_title":          "Office",
		"title":                 "Love The One You're With",
		"user_name":             "elan",
	}
	expectedFields := map[string]interface{}{
		"rating_count": int64(6794),
		"view_count":   int64(0),
	}
	acc.AssertContainsTaggedFields(t, "plex_webhooks", expectedFields, expectedTags)
}
