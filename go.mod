module github.com/russorat/telegraf-webhooks-plex

go 1.14

replace github.com/hekmon/plexwebhooks => github.com/russorat/plexwebhooks v0.1.1-0.20210225195901-9fb16538071b

require (
	github.com/gorilla/mux v1.6.2
	github.com/hekmon/plexwebhooks v0.1.0
	github.com/influxdata/telegraf v1.17.3
	github.com/stretchr/testify v1.6.1
)
