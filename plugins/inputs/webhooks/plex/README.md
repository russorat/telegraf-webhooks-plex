# plex webhooks

You should configure Plex's Webhooks to point at the `webhooks` service. To do this, follow the instructions in the [Plex webhooks support documentation](https://support.plex.tv/articles/115002267687-webhooks/). Webhooks are a premium feature and require an active Plex Pass Subscription for the Plex Media Server account.

## Events

The titles of the following sections are links to the full payloads and details for each event. The body contains what information from the event is persisted. The format is as follows:
```
# TAGS
* 'tagKey' = `tagValue` type
# FIELDS
* 'fieldKey' = `fieldValue` type
```
The tag values and field values show the place on the incoming JSON object where the data is sourced from.

**Tags:**
* 'is_user_webhook' = `event.user` bool
* 'is_owner_webhook' = `event.owner` bool
* 'server_title' = `event.Server.title` string
* 'library_selection_type' = `event.Metadata.librarySectionType` string
* 'media_type' = `event.Metadata.type` string
* 'grandparent_key' = `event.Metadata.grandparentKey` string
* 'parent_key' = `event.Metadata.parentKey` string
* 'grandparent_title' = `event.Metadata.grandparentTitle` string
* 'parent_title' = `event.Metadata.parentTitle` string
* 'parent_index' = `event.Metadata.parentIndex` string

**Fields:**		
* 'event' = `event.event` string
* 'user_id' = `event.Account.id` string
* 'user_name' = `event.Account.title` string
* 'is_player_local' = `event.Player.local` bool
* 'player_public_ip' = `event.Player.publicAddress` string
* 'player_title' = `event.Player.title` string
* 'rating_count' = `event.Metadata.ratingCount` int
* 'added_at_ns' = `event.Metadata.addedAt` int
* 'updated_at_ns' = `event.Metadata.updatedAt` int
* 'summary' = `event.Metadata.summary` string
* 'title' = `event.Metadata.title` string
* 'index' = `event.Metadata.index` int
* 'library_selection_id' = `event.Metadata.librarySectionID` int

