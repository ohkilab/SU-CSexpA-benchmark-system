// Code generated by ent, DO NOT EDIT.

package group

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldID, id))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldYear, v))
}

// EncryptedPassword applies equality check predicate on the "encrypted_password" field. It's identical to EncryptedPasswordEQ.
func EncryptedPassword(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldEncryptedPassword, v))
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldYear, v))
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldYear, v))
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldYear, vs...))
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldYear, vs...))
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldYear, v))
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldYear, v))
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldYear, v))
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldYear, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldRole, vs...))
}

// EncryptedPasswordEQ applies the EQ predicate on the "encrypted_password" field.
func EncryptedPasswordEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldEncryptedPassword, v))
}

// EncryptedPasswordNEQ applies the NEQ predicate on the "encrypted_password" field.
func EncryptedPasswordNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldEncryptedPassword, v))
}

// EncryptedPasswordIn applies the In predicate on the "encrypted_password" field.
func EncryptedPasswordIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldEncryptedPassword, vs...))
}

// EncryptedPasswordNotIn applies the NotIn predicate on the "encrypted_password" field.
func EncryptedPasswordNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldEncryptedPassword, vs...))
}

// EncryptedPasswordGT applies the GT predicate on the "encrypted_password" field.
func EncryptedPasswordGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldEncryptedPassword, v))
}

// EncryptedPasswordGTE applies the GTE predicate on the "encrypted_password" field.
func EncryptedPasswordGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldEncryptedPassword, v))
}

// EncryptedPasswordLT applies the LT predicate on the "encrypted_password" field.
func EncryptedPasswordLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldEncryptedPassword, v))
}

// EncryptedPasswordLTE applies the LTE predicate on the "encrypted_password" field.
func EncryptedPasswordLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldEncryptedPassword, v))
}

// EncryptedPasswordContains applies the Contains predicate on the "encrypted_password" field.
func EncryptedPasswordContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldEncryptedPassword, v))
}

// EncryptedPasswordHasPrefix applies the HasPrefix predicate on the "encrypted_password" field.
func EncryptedPasswordHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldEncryptedPassword, v))
}

// EncryptedPasswordHasSuffix applies the HasSuffix predicate on the "encrypted_password" field.
func EncryptedPasswordHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldEncryptedPassword, v))
}

// EncryptedPasswordEqualFold applies the EqualFold predicate on the "encrypted_password" field.
func EncryptedPasswordEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldEncryptedPassword, v))
}

// EncryptedPasswordContainsFold applies the ContainsFold predicate on the "encrypted_password" field.
func EncryptedPasswordContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldEncryptedPassword, v))
}

// HasSubmits applies the HasEdge predicate on the "submits" edge.
func HasSubmits() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SubmitsTable, SubmitsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubmitsWith applies the HasEdge predicate on the "submits" edge with a given conditions (other predicates).
func HasSubmitsWith(preds ...predicate.Submit) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := newSubmitsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		p(s.Not())
	})
}
