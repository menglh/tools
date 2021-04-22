// Code generated by entc, DO NOT EDIT.

package device

import (
	"entgo.io/ent/dialect/sql"
	"github.com/menglh/tools/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// Seq applies equality check predicate on the "seq" field. It's identical to SeqEQ.
func Seq(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeq), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIP), v))
	})
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIP), v...))
	})
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIP), v...))
	})
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIP), v))
	})
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIP), v))
	})
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIP), v))
	})
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIP), v))
	})
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIP), v))
	})
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIP), v))
	})
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIP), v))
	})
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIP), v))
	})
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIP), v))
	})
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModel), v))
	})
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...int) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldModel), v...))
	})
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...int) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldModel), v...))
	})
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModel), v))
	})
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModel), v))
	})
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModel), v))
	})
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModel), v))
	})
}

// SeqEQ applies the EQ predicate on the "seq" field.
func SeqEQ(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeq), v))
	})
}

// SeqNEQ applies the NEQ predicate on the "seq" field.
func SeqNEQ(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeq), v))
	})
}

// SeqIn applies the In predicate on the "seq" field.
func SeqIn(vs ...int) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeq), v...))
	})
}

// SeqNotIn applies the NotIn predicate on the "seq" field.
func SeqNotIn(vs ...int) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeq), v...))
	})
}

// SeqGT applies the GT predicate on the "seq" field.
func SeqGT(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeq), v))
	})
}

// SeqGTE applies the GTE predicate on the "seq" field.
func SeqGTE(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeq), v))
	})
}

// SeqLT applies the LT predicate on the "seq" field.
func SeqLT(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeq), v))
	})
}

// SeqLTE applies the LTE predicate on the "seq" field.
func SeqLTE(v int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeq), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
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
func Not(p predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		p(s.Not())
	})
}
