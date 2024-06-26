{{ $baseUrl := "https://pantheon.grafana.net" -}}
{{ $tagUrl := "/dashboards?tag=" -}}
{{ range $ti, $taxL1 := . -}}
{{ with $ti }}
# {{ if $taxL1.Name }}{{ $taxL1.Name }}{{ else }}No Level 1 Taxonomy Tag{{ end }}

{{ range $taxName, $dashboards := $taxL1.TaxL2 -}}
## {{ if $taxName }}{{ $taxName }}{{ else }}No Level 2 Taxonomy Tag{{ end }}
{{ range $dashboards -}}
* **[{{ .Title }}]({{ $baseUrl }}{{ .Url }})**
  * {{ if .Description }}{{ .Description }}{{ else }}No Description Provided{{ end }}
  * Tags: {{ range $i, $t := .Tags }}{{ if $i }}, {{ end }}[{{ $t }}]({{ $baseUrl }}{{ $tagUrl }}{{ urlquery $t }}) {{ end }}
{{ end }}{{ end }}{{ end }}{{ end }}

{{ range $ti, $taxL1 := . -}}
{{ with not $ti }}
# No Level 1 Taxonomy Tag

{{ range $taxName, $dashboards := $taxL1.TaxL2 -}}
## {{ if $taxName }}{{ $taxName }}{{ else }}No Level 2 Taxonomy Tag{{ end }}
{{ range $dashboards -}}
* **[{{ .Title }}]({{ $baseUrl }}{{ .Url }})**
  * {{ if .Description }}{{ .Description }}{{ else }}No Description Provided{{ end }}
  * Tags: {{ range $i, $t := .Tags }}{{ if $i }}, {{ end }}[{{ $t }}]({{ $baseUrl }}{{ $tagUrl }}{{ urlquery $t }}) {{ end }}
{{ end }}{{ end }}{{ end }}{{ end }}
