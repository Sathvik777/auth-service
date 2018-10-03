package request

import "encoding/json"

type BlobRequest struct {
	LogType string
	Logs *json.RawMessage
}
