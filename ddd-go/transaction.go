package tavern

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a payment between two parties
type Transaction struct {
	// all values lowercase since they are immutable
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
