package plugins

// ElasticSearchOutput CRD name
const ElasticSearchOutput = "elasticsearch"

// ElasticSearchDefaultValues for ElasticSearch output plugin
var ElasticSearchDefaultValues = map[string]string{
	"type_name":  "logstash",
}

// ElasticSearchTemplate for ElasticSearch output plugin
const ElasticSearchTemplate = `
<match **.{{ .pattern }}>
  @type elasticsearch

  host {{ .host }}
  port {{ .port }}
  logstash_format {{ .logstash_format }}
  logstash_prefix {{ .logstash_prefix }}
  logstash_dateformat {{ .logstash_dateformat }}
  flush_interval {{ .flush_interval }}
  type_name {{ .type_name }}
  include_tag_key {{ .include_tag_key }}
</match>`
