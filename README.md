# grafana-api

This project is a wrapper around the grafana API and allows us to programmatically and quickly find details about dashboards.

## How to use

1. You need a grafana API key
	a. You can create one via a [service account](https://grafana.com/docs/grafana/latest/administration/service-accounts/).
    b. This key should be set as an ENV var.

If you run this program with no arguments it will try to return the information 
about the service account you used to setup the API key.

This is a good way to test if you set things up correctly.

For instance running:
```sh
$ GRAFANA_TOKEN=$GRAFANA_TOKEN go run main.go | jq
```
_Where `$GRAFANA_TOKEN` is the the token you obtained through the service account._

If successful, should yeild something like:
```sh
{
  "id": 0,
  "uid": "service-account:188",
  "email": "sa-spi-squad",
  "name": "SPI-squad",
  "login": "sa-1-1-spi-squad",
  "theme": "",
  "orgId": 1,
  "isGrafanaAdmin": false,
  "isDisabled": false,
  "isExternal": false,
  "isExternallySynced": false,
  "isGrafanaAdminExternallySynced": false,
  "authLabels": null,
  "updatedAt": "0001-01-01T00:00:00Z",
  "createdAt": "0001-01-01T00:00:00Z",
  "avatarUrl": ""
}
```
_The information here will reflect the service account you used to create the token._


## Get by Dashboard:

Example of usage from CLI

`$ GRAFANA_TOKEN=$GRAFANA_TOKEN go run main.go QJ62QK14k | jq`

The `QJ62QK14k` is from any dashboard you see in grafana.  You can change that out for something else to see that dashboards details

## Search with Optional Params

This is the primary function of this application. The search function will
filter, categorize, and output a formatted list of dashboards to the terminal.

Futher it will generate a markdown file with the name `dashboard-taxonomy.md`
as a useful export. The contents of this file can be pasted directly into something like
confluence, and should render nicely.

Example of usage from CLI

`$ GRAFANA_TOKEN=$GRAFANA_TOKEN go run main.go search`

or for a specific "fuzzy" search (like folder names)

`$ GRAFANA_TOKEN=$GRAFANA_TOKEN go run main.go search decoupled`

**NOTE:** For testing, i would use the second option (`search decoupled`) as it returns enough information without spamming the grafana api's

