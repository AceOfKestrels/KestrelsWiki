package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Mirror holds the schema definition for the Mirror entity.
type Mirror struct {
	ent.Schema
}

// Fields of the Mirror.
func (Mirror) Fields() []ent.Field {
	return []ent.Field{
		field.String("originPath").Unique().NotEmpty().Immutable(),
		field.String("targetPath"),
	}
}

// Edges of the Mirror.
func (Mirror) Edges() []ent.Edge {
	return nil
}
