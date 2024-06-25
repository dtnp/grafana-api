{{ $baseUrl := "https://pantheon.grafana.net" -}}
{{ $tagUrl := "/dashboards?tag=" -}}
{{ range . -}}
# {{ if .Name }}{{ .Name }}{{ else }}No Level 1 Taxonomy Tag{{ end }}

{{ range $taxName, $dashboards := .TaxL2 -}}
## {{ if $taxName }}{{ $taxName }}{{ else }}No Level 2 Taxonomy Tag{{ end }}
{{ range $dashboards -}}
* **[{{ .Title }}]({{ $baseUrl }}{{ .Url }})**
  * {{ if .Description }}{{ .Description }}{{ else }}No Description Provided{{ end }}
  * Tags: {{ range $i, $t := .Tags }}{{ if $i }}, {{ end }}[{{ $t }}]({{ $baseUrl }}{{ $tagUrl }}{{ urlquery $t }}) {{ end }}
{{ end -}}{{ end -}}{{ end }}
