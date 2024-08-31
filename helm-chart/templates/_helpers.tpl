{{/*
Return the proper app image name
*/}}
{{- define "app.image" -}}
{{ include "common.images.image" (dict "imageRoot" .Values.app.image "global" .Values.global) }}
{{- end -}}

{{/*
Return the proper Docker Image Registry Secret Names
*/}}
{{- define "app.imagePullSecrets" -}}
{{- include "common.images.renderPullSecrets" (dict "images" (list .Values.app.image) "context" $) -}}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "app.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "common.names.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}
