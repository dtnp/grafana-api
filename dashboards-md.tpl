{{ $baseUrl := "https://pantheon.grafana.net" -}}
# Grafana Dashboard Status and Classification
This is an attempt to parse and display the simple taxonomy tags associated
with grafana dashboards.

---
{{ range . -}}
## {{ if .Name }}{{ .Name }}{{ else }}No Level 1 Taxonomy Tag{{ end }}

{{ range $taxName, $dashboards := .TaxL2 -}}
### {{ if $taxName }}{{ $taxName }}{{ else }}No Level 2 Taxonomy Tag{{ end }}
{{ range $dashboards -}}
* [{{ .Title }}]({{ $baseUrl }}{{ .Url }}) [owner: {{ if .Owner }}{{ .Owner }}{{ else }}None Provided{{ end }}]
  {{ if .Description }}{{ .Description }}{{ else }}No Description Provided{{ end }}
{{ end -}}
{{ end -}}
---
{{ end }}
