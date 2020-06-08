package storage

import (
	"time"
)

//WebResource defines the prepresentation of a web resource
type WebResource struct {
	Content []byte
	Expiry  time.Time
}
