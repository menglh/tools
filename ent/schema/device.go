package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

func (Device) Config() ent.Config {
	return ent.Config{
		Table: "device",
	}
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.String("ip"),
		field.Int("model"),
		field.Int("seq"),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return nil
}
