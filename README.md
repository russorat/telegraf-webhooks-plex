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

Tags and fields are determined by the media type of the event (Movie, Show, or Music). All events contain the following tags and fields, additional tags and fields per type are mentioned below. 

- plex_webhooks
  - tags:
    - is_user
    - is_owner
    - server_title
    - is_player_local
    - library_selection_type
    - library_selection_title
    - media_type
    - grandparent_title
    - parent_title
    - event
    - user_name
    - player_title
    - title
  - fields:
    - view_count (int, count)

If the event is a rating event, then a `rating` field will be added of type `int`.

If the event is a Movie, the following tags are included:

- tags:
  - content_rating: For example PG-13, R
  - chapter_source
  - genre: Note, only the first in the list is stored
  - director: Note, only the first in the list is stored
  - country: Note, only the first in the list is stored
  - collection: Note, only the first in the list is stored
  - studio
  - year

If the event is a Show, the following tags are included:

- tags:
  - content_rating: For example TV-14, TV-MA
  - year
  - chapter_source
  - director: Note, only the first in the list is stored

If the event is Music, the following tags and fields are included:

- tags:
  - mood
- fields:
  - rating_count (int, count)

### Example Output

```
plex_webhooks,content_rating=Passed,country=USA,director=Phil\ Karlson,event=media.play,genre=Crime,is_owner=true,is_player_local=true,is_user=true,library_selection_title=Movies,library_selection_type=movie,media_type=movie,player_title=Chrome,server_title=BigSur,studio=Columbia\ Pictures,title=Scandal\ Sheet,user_name=russorat,year=1952 view_count=1i 1614017055347830000
plex_webhooks,content_rating=Passed,country=USA,director=Phil\ Karlson,event=media.stop,genre=Crime,is_owner=true,is_player_local=true,is_user=true,library_selection_title=Movies,library_selection_type=movie,media_type=movie,player_title=Chrome,server_title=BigSur,studio=Columbia\ Pictures,title=Scandal\ Sheet,user_name=russorat,year=1952 view_count=1i 1614017062286684000
plex_webhooks,content_rating=Passed,country=USA,director=Phil\ Karlson,event=media.rate,genre=Crime,is_owner=true,is_player_local=true,is_user=true,library_selection_title=Movies,library_selection_type=movie,media_type=movie,player_title=Chrome,server_title=BigSur,studio=Columbia\ Pictures,title=Scandal\ Sheet,user_name=russorat,year=1952 view_count=1i,rating=4i 1614017071892773000
plex_webhooks,content_rating=Passed,country=USA,director=Phil\ Karlson,event=media.rate,genre=Crime,is_owner=true,is_player_local=true,is_user=true,library_selection_title=Movies,library_selection_type=movie,media_type=movie,player_title=Chrome,server_title=BigSur,studio=Columbia\ Pictures,title=Scandal\ Sheet,user_name=russorat,year=1952 view_count=1i,rating=10i 1614017073035550000
plex_webhooks,content_rating=TV-14,event=media.play,grandparent_title=Saturday\ Night\ Live,is_owner=true,is_player_local=true,is_user=true,library_selection_title=TV\ Shows,library_selection_type=show,media_type=episode,parent_title=Season\ 46,player_title=Chrome,server_title=BigSur,title=Kristen\ Wiig\ /\ Dua\ Lipa,user_name=russorat,year=2020 view_count=0i 1614017148410286000
plex_webhooks,content_rating=TV-14,event=media.stop,grandparent_title=Saturday\ Night\ Live,is_owner=true,is_player_local=true,is_user=true,library_selection_title=TV\ Shows,library_selection_type=show,media_type=episode,parent_title=Season\ 46,player_title=Chrome,server_title=BigSur,title=Kristen\ Wiig\ /\ Dua\ Lipa,user_name=russorat,year=2020 view_count=0i 1614017155183018000
```
