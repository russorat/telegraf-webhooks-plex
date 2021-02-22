package webhooks

import (
	"fmt"
	"net"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/russorat/telegraf-webhooks-plex/plugins/inputs/webhooks/plex"
)

type Webhook interface {
	Register(router *mux.Router, acc telegraf.Accumulator)
}

func init() {
	inputs.Add("external_webhooks", func() telegraf.Input { return NewWebhooks() })
}

type Webhooks struct {
	ServiceAddress string `toml:"service_address"`

	Plex *plex.PlexWebhook `toml:"plex"`

	Log telegraf.Logger `toml:"-"`

	srv *http.Server
}

func NewWebhooks() *Webhooks {
	return &Webhooks{}
}

func (wb *Webhooks) SampleConfig() string {
	return `
  ## Address and port to host Webhook listener on
  service_address = ":1619"

  [inputs.webhooks.plex]
    path = "/plex"
`
}

func (wb *Webhooks) Description() string {
	return "A Webhooks Event collector"
}

func (wb *Webhooks) Gather(_ telegraf.Accumulator) error {
	return nil
}

// Looks for fields which implement Webhook interface
func (wb *Webhooks) AvailableWebhooks() []Webhook {
	webhooks := make([]Webhook, 0)
	s := reflect.ValueOf(wb).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		if !f.CanInterface() {
			continue
		}

		if wbPlugin, ok := f.Interface().(Webhook); ok {
			if !reflect.ValueOf(wbPlugin).IsNil() {
				webhooks = append(webhooks, wbPlugin)
			}
		}
	}

	return webhooks
}

func (wb *Webhooks) Start(acc telegraf.Accumulator) error {
	r := mux.NewRouter()

	for _, webhook := range wb.AvailableWebhooks() {
		webhook.Register(r, acc)
	}

	wb.srv = &http.Server{Handler: r}

	ln, err := net.Listen("tcp", fmt.Sprintf("%s", wb.ServiceAddress))
	if err != nil {
		return fmt.Errorf("error starting server: %v", err)
	}

	go func() {
		if err := wb.srv.Serve(ln); err != nil {
			if err != http.ErrServerClosed {
				acc.AddError(fmt.Errorf("error listening: %v", err))
			}
		}
	}()

	wb.Log.Infof("Started the webhooks service on %s", wb.ServiceAddress)

	return nil
}

func (wb *Webhooks) Stop() {
	wb.srv.Close()
	wb.Log.Infof("Stopping the Webhooks service")
}
