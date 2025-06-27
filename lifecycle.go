package blocks

type LifecycleStatus string

const (
	LifecycleStatusIngested             = LifecycleStatus("ingested")
	LifecycleStatusCreated              = LifecycleStatus("created")
	LifecycleStatusEnriching            = LifecycleStatus("enriching")
	LifecycleStatusEnriched             = LifecycleStatus("enriched")
	LifecycleStatusTransforming         = LifecycleStatus("transforming")
	LifecycleStatusTransformed          = LifecycleStatus("transformed")
	LifecycleStatusRouted               = LifecycleStatus("routed")
	LifecycleStatusRoutedFinal          = LifecycleStatus("routed_final")
	LifecycleStatusEditing              = LifecycleStatus("editing")
	LifecycleStatusEdited               = LifecycleStatus("edited")
	LifecycleStatusProcessing           = LifecycleStatus("processing")
	LifecycleStatusProcessed            = LifecycleStatus("processed")
	LifecycleStatusIndexing             = LifecycleStatus("indexing")
	LifecycleStatusIndexed              = LifecycleStatus("indexed")
	LifecycleStatusOnHold               = LifecycleStatus("on_hold")
	LifecycleStatusEnrichmentFailed     = LifecycleStatus("enrichment_failed")
	LifecycleStatusRoutingFailed        = LifecycleStatus("routing_failed")
	LifecycleStatusProcessingFailed     = LifecycleStatus("processing_failed")
	LifecycleStatusArchived             = LifecycleStatus("archived")
	LifecycleStatusPreProcessing        = LifecycleStatus("pre_processing")
	LifecycleStatusPreProcessingFailed  = LifecycleStatus("pre_processing_failed")
	LifecycleStatusPreProcessed         = LifecycleStatus("pre_processed")
	LifecycleStatusPostProcessing       = LifecycleStatus("post_processing")
	LifecycleStatusPostProcessingFailed = LifecycleStatus("post_processing_failed")
	LifecycleStatusPostProcessed        = LifecycleStatus("post_processed")
)

// String returns the string representation of the lifecycle status.
// This method is kept for backward compatibility, but it simply returns the string value.
func (s LifecycleStatus) String() string {
	return string(s)
}

func (s LifecycleStatus) Recordable() bool {
	switch s {
	case LifecycleStatusIngested,
		LifecycleStatusCreated,
		LifecycleStatusEnriched,
		LifecycleStatusTransformed,
		LifecycleStatusRouted,
		LifecycleStatusRoutedFinal,
		LifecycleStatusEdited,
		LifecycleStatusProcessed,
		LifecycleStatusPreProcessed,
		LifecycleStatusPostProcessed:
		return true
	default:
		return false
	}
}
