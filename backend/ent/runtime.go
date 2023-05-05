// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/schema"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/submit"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	contestFields := schema.Contest{}.Fields()
	_ = contestFields
	// contestDescYear is the schema descriptor for year field.
	contestDescYear := contestFields[5].Descriptor()
	// contest.YearValidator is a validator for the "year" field. It is called by the builders before save.
	contest.YearValidator = contestDescYear.Validators[0].(func(int) error)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescYear is the schema descriptor for year field.
	groupDescYear := groupFields[2].Descriptor()
	// group.YearValidator is a validator for the "year" field. It is called by the builders before save.
	group.YearValidator = groupDescYear.Validators[0].(func(int) error)
	// groupDescScore is the schema descriptor for score field.
	groupDescScore := groupFields[3].Descriptor()
	// group.ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	group.ScoreValidator = groupDescScore.Validators[0].(func(int) error)
	submitFields := schema.Submit{}.Fields()
	_ = submitFields
	// submitDescYear is the schema descriptor for year field.
	submitDescYear := submitFields[2].Descriptor()
	// submit.YearValidator is a validator for the "year" field. It is called by the builders before save.
	submit.YearValidator = submitDescYear.Validators[0].(func(int) error)
}
