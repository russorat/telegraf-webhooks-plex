# Webhooks

This is a Telegraf service plugin that start an http server and register multiple webhook listeners.

```sh
$ telegraf config -input-filter webhooks -output-filter influxdb > config.conf.new
```

Change the config file to point to the InfluxDB server you are using and adjust the settings to match your environment. Once that is complete:

```sh
$ cp config.conf.new /etc/telegraf/telegraf.conf
$ sudo service telegraf start
```


### Configuration:

```toml
[[inputs.webhooks]]
  ## Address and port to host Webhook listener on
  service_address = ":1619"

  [inputs.webhooks.plex]
    path = "/plex"
```


### Available webhooks

- [Plex](plugins/inputs/webhooks/plex/)


### Adding new webhooks plugin

1. Add your webhook plugin inside the `webhooks` folder
1. Your plugin must implement the `Webhook` interface
1. Import your plugin in the `webhooks.go` file and add it to the `Webhooks` struct

