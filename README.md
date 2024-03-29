[![Go](https://github.com/russorat/telegraf-webhooks-plex/actions/workflows/go.yml/badge.svg)](https://github.com/russorat/telegraf-webhooks-plex/actions/workflows/go.yml)

# Plex Webhooks Input Plugin

This plugin will listen to events published from [Plex webhooks](https://support.plex.tv/articles/115002267687-webhooks/). Webhooks are a premium feature and require an active Plex Pass Subscription for the Plex Media Server account.

It is modeled after the [Telegraf Webhooks plugin](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/webhooks).

Telegraf minimum version: Telegraf 17.0
Plugin minimum tested version: 17.3

### Build and Run

To build this plugin, just run:

```sh
make
```

Which will build the binary `./bin/webhook-plex`

You can run it with `./bin/webhook-plex --config plugin.conf`

### Configuration

This is the plugin configuration that is expected via the `--config` param

```toml
[[inputs.custom_webhooks]]
  ## Address and port to host Webhook listener on
  service_address = ":1619"

  [inputs.custom_webhooks.plex]
    path = "/plex"
```

And this is an example of how to configure this with the Telegraf execd plugin that you would run with `telegraf --config telegraf.conf`

```toml
[[inputs.execd]]
  ## One program to run as daemon.
  ## NOTE: process and each argument should each be their own string
  command = ["/path/to/webhook-plex", "--config", "/path/to/plugin.conf"]

  ## Define how the process is signaled on each collection interval.
  ## Valid values are:
  ##   "none"    : Do not signal anything. (Recommended for service inputs)
  ##               The process must output metrics by itself.
  ##   "STDIN"   : Send a newline on STDIN. (Recommended for gather inputs)
  ##   "SIGHUP"  : Send a HUP signal. Not available on Windows. (not recommended)
  ##   "SIGUSR1" : Send a USR1 signal. Not available on Windows.
  ##   "SIGUSR2" : Send a USR2 signal. Not available on Windows.
  signal = "none"

  ## Delay before the process is restarted after an unexpected termination
  restart_delay = "10s"

  ## Data format to consume.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_INPUT.md
  data_format = "influx"
  ```

### Metrics

Tags can vary based on the Metadata Type of the event. The field is always a counter with a value of 1. If the event is a "rating" event, an additional field of "rating" is added. 

- plex
  - tags:
    - is_user
    - is_owner
    - event
    - server_title
    - account_title
    - is_player_local
    - player_title
    - metadata_content_rating
    - metadata_director: Note, only the first director in the list is stored
    - metadata_grandparent_title
    - metadata_library_section_title
    - metadata_library_section_type
    - metadata_original_title
    - metadata_parent_title
    - metadata_studio
    - metadata_title
    - metadata_type
    - metadata_subtype
    - metadata_year
    - metadata_parent_year
  - fields:
    - counter (int, count)
    - rating (int, 1 - 10, only available if the event type is rating)

### Example Output

```
plex,account_title=username,event=media.play,host=plexserver.lan,is_owner=true,is_player_local=true,is_user=true,metadata_content_rating=TV-14,metadata_grandparent_title=30\ Rock,metadata_library_section_title=TV\ Shows,metadata_library_section_type=show,metadata_parent_title=Season\ 2,metadata_title=Jack\ Gets\ in\ the\ Game,metadata_type=episode,metadata_year=2007,player_title=Chrome,server_title=Office counter=1i 1614285785465333000
plex,account_title=username,event=media.stop,host=plexserver.lan,is_owner=true,is_player_local=true,is_user=true,metadata_content_rating=TV-14,metadata_grandparent_title=30\ Rock,metadata_library_section_title=TV\ Shows,metadata_library_section_type=show,metadata_parent_title=Season\ 2,metadata_title=Jack\ Gets\ in\ the\ Game,metadata_type=episode,metadata_year=2007,player_title=Chrome,server_title=Office counter=1i 1614285787351854000
plex,account_title=username,event=media.play,host=plexserver.lan,is_owner=true,is_player_local=true,is_user=true,metadata_grandparent_title=Various\ Artists,metadata_library_section_title=Music,metadata_library_section_type=artist,metadata_original_title=Britney\ Spears,metadata_parent_title=RTL:\ I\ Like\ the\ 90’s,metadata_parent_year=2014,metadata_title=Oops!…I\ Did\ It\ Again\ (album\ version),metadata_type=track,player_title=Chrome,server_title=Office counter=1i 1614285797801001000
plex,account_title=username,event=media.stop,host=plexserver.lan,is_owner=true,is_player_local=true,is_user=true,metadata_grandparent_title=Various\ Artists,metadata_library_section_title=Music,metadata_library_section_type=artist,metadata_original_title=Britney\ Spears,metadata_parent_title=RTL:\ I\ Like\ the\ 90’s,metadata_parent_year=2014,metadata_title=Oops!…I\ Did\ It\ Again\ (album\ version),metadata_type=track,player_title=Chrome,server_title=Office counter=1i 1614285800510840000
plex,account_title=username,event=media.rate,host=plexserver.lan,is_owner=true,is_player_local=true,is_user=true,metadata_grandparent_title=Various\ Artists,metadata_library_section_title=Music,metadata_library_section_type=artist,metadata_original_title=Britney\ Spears,metadata_parent_title=RTL:\ I\ Like\ the\ 90’s,metadata_parent_year=2014,metadata_title=Oops!…I\ Did\ It\ Again\ (album\ version),metadata_type=track,player_title=Chrome,server_title=Office counter=1i,rating=10i 1614285816914533000
```
