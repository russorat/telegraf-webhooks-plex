package plex

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hekmon/plexwebhooks"
	"github.com/influxdata/telegraf"
)

const measurement = "plex_webhooks"

type PlexWebhook struct {
	Path string
	acc  telegraf.Accumulator
}

func (p *PlexWebhook) Register(router *mux.Router, acc telegraf.Accumulator) {
	router.HandleFunc(p.Path, p.eventHandler).Methods("POST")
	log.Printf("I! Started the webhooks_plex on %s\n", p.Path)
	p.acc = acc
}

func (p *PlexWebhook) eventHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// Create the multi part reader
	multiPartReader, err := r.MultipartReader()
	if err != nil {
		// Detect error type for the http answer
		if err == http.ErrNotMultipart || err == http.ErrMissingBoundary {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		// Try to write the error as http body
		_, wErr := w.Write([]byte(err.Error()))
		if wErr != nil {
			err = fmt.Errorf("request error: %v | write error: %v", err, wErr)
		}
		// Log the error
		fmt.Println("can't create a multipart reader from request:", err)
		return
	}
	// Use the multipart reader to parse the request body
	payload, _, err := plexwebhooks.Extract(multiPartReader)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// Try to write the error as http body
		_, wErr := w.Write([]byte(err.Error()))
		if wErr != nil {
			log.Printf("request error: %v | write error: %v", err, wErr)
		}
		// Log the error
		log.Println("can't create a multipart reader from request:", err)
		return
	}
	if payload != nil {
		newMetric, err := NewMetric(payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p.acc.AddFields(measurement, newMetric.Fields(), newMetric.Tags(), newMetric.Time())
	}
	w.WriteHeader(http.StatusOK)
}
