package blocks

import "github.com/google/uuid"

type Origin struct {
	ConnectorSlug                 string    `json:"connector_slug"`
	ConnectorUniqSourceIdentifier string    `json:"connector_uniq_source_identifier"`
	ConnectorIdentifier           uuid.UUID `json:"connector_identifier"`
	ModifiedBy                    *string   `json:"modified_by,omitempty"`
}

func NewOriginGeneric() Origin {
	return Origin{
		ConnectorSlug:                 "generic",
		ConnectorUniqSourceIdentifier: uuid.New().String(),
		ConnectorIdentifier:           uuid.Nil,
	}
}

func NewOriginWebapp() Origin {
	return Origin{
		ConnectorSlug:                 "webapp",
		ConnectorUniqSourceIdentifier: uuid.New().String(),
		ConnectorIdentifier:           uuid.Nil,
	}
}

func NewOriginAI() Origin {
	return Origin{
		ConnectorSlug:                 "ai",
		ConnectorUniqSourceIdentifier: uuid.New().String(),
		ConnectorIdentifier:           uuid.Nil,
	}
}
