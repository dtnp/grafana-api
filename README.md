# grafana-api

This project is a wrapper around the grafana API and allows us to programmatically and quickly find details about dashboards.

## How to use

1. You need a grafana API key
	a. You can create one via a service account

## Get by Dashboard:

Example of usage from CLI

`$ GRAFANA_TOKEN=$GRAFANA_TOKEN go run main.go QJ62QK14k | jq`

The `QJ62QK14k` is from any dashboard you see in grafana.  You can change that out for something else to see that dashboards details

## Search with Optional Params

Example of usage from CLI

`$ GRAFANA_TOKEN=$GRAFANA_TOKEN go run main.go search`

or for a specific "fuzzy" search (like folder names)

`$ GRAFANA_TOKEN=$GRAFANA_TOKEN go run main.go search decoupled`

NOTE: For testing, i would use the second option (`search decoupled`) as it returns enough information without spamming the grafana api's

