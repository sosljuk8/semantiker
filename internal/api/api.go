package api

import "semantiker/dto"

type (
	// Deps provides dependencies interface for groups of handlers.
	CoreKeyBind interface {
		ReadKeyFile(name string) []*dto.Corekey
	}
)

//// JobsGroup provides interface for handlers of Jobs group.
//type JobsGroup interface {
//	JobStatus(context.Context, JobStatusCommand) (Job, error)
//}
//
//// ParsingGroup provides interface for handlers of Parsing group.
//type ParsingGroup interface {
//	Ping(context.Context) (string, error)
//}
//
//// Term provides DTO interface for Term type.
//type Term interface {
//	GetID() int64
//	GetName() string
//	SetID(id int64)
//	SetName(name string)
//	Validate() error
//	Zap() zap.Field
//}
//
//// Term provides DTO interface for Term type.
//type TermRead interface {
//	GetID() int64
//	GetName() string
//	Zap() zap.Field
//}
