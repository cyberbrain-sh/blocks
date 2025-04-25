package blocks

import (
	"time"

	"github.com/google/uuid"
)

type MoveReason int64

const MoveReasonManual MoveReason = 0
const MoveReasonRouter MoveReason = 1

type DestinationType int64

const DestinationTypeSpace DestinationType = 0
const DestinationTypeNote DestinationType = 1
const DestinationTypeBlock DestinationType = 2

type Move struct {
	FromType          DestinationType `json:"from_type"`
	FromID            uuid.UUID       `json:"from_id"`
	ToType            DestinationType `json:"to_type"`
	ToID              uuid.UUID       `json:"to_id"`
	Timestamp         time.Time       `json:"timestamp"`
	Reason            MoveReason      `json:"reason"`
	Accuracy          float64         `json:"accuracy"`
	Reasoning         string          `json:"reasoning"`
	ReasoningKeywords []string        `json:"reasoning_keywords"`
	SpaceKeywords     []string        `json:"space_keywords"`
}
