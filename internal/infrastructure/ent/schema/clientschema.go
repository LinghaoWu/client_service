package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClientSchema holds the schema definition for the ClientSchema entity.
type ClientSchema struct {
	ent.Schema
}

// Fields of the ClientSchema.
func (ClientSchema) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("ClientID").StorageKey("client_id").Immutable().Unique(),
		field.String("Name").StorageKey("name").NotEmpty(),
		field.String("USCC").StorageKey("uscc").NotEmpty().Unique(),
		field.String("Contact").StorageKey("contact").NotEmpty(),
		field.String("Phone").StorageKey("phone").NotEmpty(),
		field.String("Email").StorageKey("email").NotEmpty(),
		field.String("Address").StorageKey("address").NotEmpty(),
		field.Bool("IsActive").StorageKey("is_active").Default(true),
		field.Time("CreatedAt").StorageKey("created_at").Default(time.Now).Immutable(),
		field.Time("UpdatedAt").StorageKey("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ClientSchema.
func (ClientSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", MemberSchema.Type),
	}
}
