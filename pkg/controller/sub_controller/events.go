package sub_controller

// define event type for sub controller, Type  can be one of Normal, Warning.
type EventType string

// only Normal Warning, not add new type.
var (
	EventNormal  EventType = "Normal"
	EventWarning EventType = "Warning"
)

// 'reason' should be short and unique; it should be in UpperCamelCase format (starting with a capital letter).
const (
	StatefulSetNotExist     = "StatefulSetNotExist"
	AutoScalerDeleteFailed  = "AutoScalerDeleteFailed"
	ComponentImageUpdate    = "ComponentImageUpdate"
	PVCListFailed           = "PVCListFailed"
	PVCUpdate               = "PVCUpdated"
	PVCUpdateFailed         = "PVCUpdateFailed"
	PVCDeleteFailed         = "PVCDeleteFailed"
	PVCCreate               = "PVCCreate"
	PVCCreateFailed         = "PVCCreateFailed"
	FollowerScaleDownFailed = "FollowerScaleDownFailed"
	ConfigMapPathRepeated   = "ConfigMapPathRepeated"
)

type EventReason string

var (
	ImageFormatError                  EventReason = "ImageFormatError"
	FDBSpecEmpty                      EventReason = "SpecEmpty"
	ComputeGroupsEmpty                EventReason = "CGsEmpty"
	CGUniqueIdentifierDuplicate       EventReason = "CGUniqueIdentifierDuplicate"
	CGCreateResourceFailed            EventReason = "CGCreateResourceFailed"
	CGApplyResourceFailed             EventReason = "CGApplyResourceFailed"
	CGStatusUpdateFailed              EventReason = "CGStatusUpdatedFailed"
	CGStatefulsetDeleteFailed         EventReason = "CGStatefulsetDeleteFailed"
	DisaggregatedMetaServiceGetFailed EventReason = "DisaggregatedMetaServiceGetFailed"
	ObjectInfoInvalid                 EventReason = "ObjectInfoInvalid"
	ConfigMapGetFailed                EventReason = "ConfigMapGetFailed"
	ObjectConfigError                 EventReason = "ObjectConfigError"
	MSInteractError                   EventReason = "MSInteractError"
	InstanceMetaCreated               EventReason = "InstanceMetaCreated"
)

type Event struct {
	Type    EventType
	Reason  EventReason
	Message string
}

func EventString(event *Event) string {
	return string(event.Reason) + "," + event.Message
}
