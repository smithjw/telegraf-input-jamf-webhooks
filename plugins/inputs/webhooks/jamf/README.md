# github webhooks

You should configure your Organization's Webhooks to point at the `webhooks`
service. To do this go to `github.com/{my_organization}` and click
`Settings > Webhooks > Add webhook`. In the resulting menu set `Payload URL` to
`http://<my_ip>:1619/github`, `Content type` to `application/json` and under
the section `Which events would you like to trigger this webhook?` select
'Send me **everything**'. By default all of the events will write to the
`github_webhooks` measurement, this is configurable by setting the
`measurement_name` in the config file.

You can also add a secret that will be used by telegraf to verify the
authenticity of the requests.

## Metrics

The titles of the following sections are links to the full payloads and details
for each event. The body contains what information from the event is persisted.
The format is as follows:

```toml
# TAGS
* 'tagKey' = `tagValue` type
# FIELDS
* 'fieldKey' = `fieldValue` type
```

The tag values and field values show the place on the incoming JSON object
where the data is sourced from.


### [`team_add` event](https://developer.github.com/v3/activity/events/types/#teamaddevent)

**Tags:**

* 'event' = `headers[X-Github-Event]` string
* 'repository' = `event.repository.full_name` string
* 'private' = `event.repository.private` bool
* 'user' = `event.sender.login` string
* 'admin' = `event.sender.site_admin` bool

**Fields:**

* 'stars' = `event.repository.stargazers_count` int
* 'forks' = `event.repository.forks_count` int
* 'issues' = `event.repository.open_issues_count` int
* 'teamName' = `event.team.name` string

