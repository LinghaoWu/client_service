package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MemberSchema holds the schema definition for the MemberSchema entity.
type MemberSchema struct {
	ent.Schema
}

// Fields of the MemberSchema.
func (MemberSchema) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("MemberID").StorageKey("member_id").Immutable().Unique(),
		field.String("Name").StorageKey("name").NotEmpty(),
		field.String("Phone").StorageKey("phone").Optional(),
		field.String("Email").StorageKey("email").Optional(),
		field.String("Role").StorageKey("role").Optional(),
	}
}

// Edges of the MemberSchema.
func (MemberSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clients", ClientSchema.Type).Ref("members"),
	}
}
