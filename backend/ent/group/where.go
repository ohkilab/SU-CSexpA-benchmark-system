// Code generated by ent, DO NOT EDIT.

package group

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldName, v))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldYear, v))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldScore, v))
}

// EncryptedPassword applies equality check predicate on the "encrypted_password" field. It's identical to EncryptedPasswordEQ.
func EncryptedPassword(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldEncryptedPassword, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldName, v))
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

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldScore, v))
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
			sqlgraph.Edge(sqlgraph.O2M, false, SubmitsTable, SubmitsColumn),
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
