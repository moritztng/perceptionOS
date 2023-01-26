package graph

import (
	"github.com/moritztng/perceptionOS/data"
	"github.com/moritztng/perceptionOS/messaging"
)

type Resolver struct {
	DB              *data.Database
	MessageProducer *messaging.Producer
}
