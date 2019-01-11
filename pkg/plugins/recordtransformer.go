package plugins

// RecordTransformerFilter CRD name
const RecordTransformerFilter = "recordtransformer"

// RecordTransformerFilterDefaultValues for recordtransformer plugin
var RecordTransformerFilterDefaultValues = map[string]string{
	"remove_keys": "namespace",
}

// RecordTransformerFilterTemplate for recordtransformer plugin
const RecordTransformerFilterTemplate = `
<filter **.{{ .pattern }}>
  @type record_transformer
  remove_keys {{ .remove_keys }}
</filter>
`
