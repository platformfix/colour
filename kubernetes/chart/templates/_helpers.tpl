{{- define "colour.name" -}}
colour
{{- end -}}

{{- define "colour.labels" -}}
app.kubernetes.io/name: {{ include "colour.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}
