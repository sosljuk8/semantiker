package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Cores holds the schema definition for the Cores entity.
type Cores struct {
	ent.Schema
}

// Fields of the Cores.
func (Cores) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").StructTag(`json:"Name" validate:"required"`),
		field.String("Status").StructTag(`json:"Status"`),
		field.Int("Count").StructTag(`json:"Count"`).Default(0),
		field.String("Created").StructTag(`json:"Created"`).DefaultFunc(UTCNow),
		field.String("Updated").StructTag(`json:"Updated"`).DefaultFunc(UTCNow),
	}
}

// Edges of the Cores.
func (Cores) Edges() []ent.Edge {
	return nil
}
