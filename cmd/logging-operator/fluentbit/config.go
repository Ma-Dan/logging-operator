package fluentbit

var fluentBitConfigTemplate = `
[SERVICE]
    Flush        1
    Daemon       Off
    Log_Level    info
    Parsers_File parsers.conf
    HTTP_Server  On
    HTTP_Listen  0.0.0.0
    HTTP_Port    {{ .Monitor.Port }}

[INPUT]
    Name             tail
    Path             /var/log/containers/*.log
    Parser           docker
    Tag              kubernetes.*
    Refresh_Interval 5
    Mem_Buf_Limit    5MB
    Skip_Long_Lines  On
    DB               /tail-db/tail-containers-state.db
    DB.Sync          Normal

[FILTER]
    Name                kubernetes
    Match               kubernetes.*
    Kube_URL            https://kubernetes.default.svc:443
    Kube_CA_File        /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    Kube_Token_File     /var/run/secrets/kubernetes.io/serviceaccount/token
    Merge_Log           On

[OUTPUT]
    Name  es
    Match kubernetes.*
    Host  elasticsearch-logging-data.{{ .Namespace }}.svc
    Port  9200
    Logstash_Format On
    Replace_Dots on
    Retry_Limit False
    Type  flb_type
    Time_Key @timestamp
    Logstash_Prefix logstash
    {{ if .TLS.Enabled }}
    tls           On
    tls.verify    Off
    tls.ca_file   /fluent-bit/tls/caCert
    tls.crt_file  /fluent-bit/tls/clientCert
    tls.key_file  /fluent-bit/tls/clientKey
    Shared_Key    {{ .TLS.SharedKey }}
    {{- end }}
`
