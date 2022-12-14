// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProtonMail/gluon/imap"
	"github.com/ProtonMail/gluon/internal/db/ent/mailbox"
	"github.com/ProtonMail/gluon/internal/db/ent/mailboxattr"
	"github.com/ProtonMail/gluon/internal/db/ent/mailboxflag"
	"github.com/ProtonMail/gluon/internal/db/ent/mailboxpermflag"
	"github.com/ProtonMail/gluon/internal/db/ent/uid"
)

// MailboxCreate is the builder for creating a Mailbox entity.
type MailboxCreate struct {
	config
	mutation *MailboxMutation
	hooks    []Hook
}

// SetRemoteID sets the "RemoteID" field.
func (mc *MailboxCreate) SetRemoteID(ii imap.MailboxID) *MailboxCreate {
	mc.mutation.SetRemoteID(ii)
	return mc
}

// SetNillableRemoteID sets the "RemoteID" field if the given value is not nil.
func (mc *MailboxCreate) SetNillableRemoteID(ii *imap.MailboxID) *MailboxCreate {
	if ii != nil {
		mc.SetRemoteID(*ii)
	}
	return mc
}

// SetName sets the "Name" field.
func (mc *MailboxCreate) SetName(s string) *MailboxCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetUIDNext sets the "UIDNext" field.
func (mc *MailboxCreate) SetUIDNext(i imap.UID) *MailboxCreate {
	mc.mutation.SetUIDNext(i)
	return mc
}

// SetNillableUIDNext sets the "UIDNext" field if the given value is not nil.
func (mc *MailboxCreate) SetNillableUIDNext(i *imap.UID) *MailboxCreate {
	if i != nil {
		mc.SetUIDNext(*i)
	}
	return mc
}

// SetUIDValidity sets the "UIDValidity" field.
func (mc *MailboxCreate) SetUIDValidity(i imap.UID) *MailboxCreate {
	mc.mutation.SetUIDValidity(i)
	return mc
}

// SetNillableUIDValidity sets the "UIDValidity" field if the given value is not nil.
func (mc *MailboxCreate) SetNillableUIDValidity(i *imap.UID) *MailboxCreate {
	if i != nil {
		mc.SetUIDValidity(*i)
	}
	return mc
}

// SetSubscribed sets the "Subscribed" field.
func (mc *MailboxCreate) SetSubscribed(b bool) *MailboxCreate {
	mc.mutation.SetSubscribed(b)
	return mc
}

// SetNillableSubscribed sets the "Subscribed" field if the given value is not nil.
func (mc *MailboxCreate) SetNillableSubscribed(b *bool) *MailboxCreate {
	if b != nil {
		mc.SetSubscribed(*b)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MailboxCreate) SetID(imi imap.InternalMailboxID) *MailboxCreate {
	mc.mutation.SetID(imi)
	return mc
}

// AddUIDIDs adds the "UIDs" edge to the UID entity by IDs.
func (mc *MailboxCreate) AddUIDIDs(ids ...int) *MailboxCreate {
	mc.mutation.AddUIDIDs(ids...)
	return mc
}

// AddUIDs adds the "UIDs" edges to the UID entity.
func (mc *MailboxCreate) AddUIDs(u ...*UID) *MailboxCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mc.AddUIDIDs(ids...)
}

// AddFlagIDs adds the "flags" edge to the MailboxFlag entity by IDs.
func (mc *MailboxCreate) AddFlagIDs(ids ...int) *MailboxCreate {
	mc.mutation.AddFlagIDs(ids...)
	return mc
}

// AddFlags adds the "flags" edges to the MailboxFlag entity.
func (mc *MailboxCreate) AddFlags(m ...*MailboxFlag) *MailboxCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddFlagIDs(ids...)
}

// AddPermanentFlagIDs adds the "permanent_flags" edge to the MailboxPermFlag entity by IDs.
func (mc *MailboxCreate) AddPermanentFlagIDs(ids ...int) *MailboxCreate {
	mc.mutation.AddPermanentFlagIDs(ids...)
	return mc
}

// AddPermanentFlags adds the "permanent_flags" edges to the MailboxPermFlag entity.
func (mc *MailboxCreate) AddPermanentFlags(m ...*MailboxPermFlag) *MailboxCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddPermanentFlagIDs(ids...)
}

// AddAttributeIDs adds the "attributes" edge to the MailboxAttr entity by IDs.
func (mc *MailboxCreate) AddAttributeIDs(ids ...int) *MailboxCreate {
	mc.mutation.AddAttributeIDs(ids...)
	return mc
}

// AddAttributes adds the "attributes" edges to the MailboxAttr entity.
func (mc *MailboxCreate) AddAttributes(m ...*MailboxAttr) *MailboxCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddAttributeIDs(ids...)
}

// Mutation returns the MailboxMutation object of the builder.
func (mc *MailboxCreate) Mutation() *MailboxMutation {
	return mc.mutation
}

// Save creates the Mailbox in the database.
func (mc *MailboxCreate) Save(ctx context.Context) (*Mailbox, error) {
	var (
		err  error
		node *Mailbox
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MailboxMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Mailbox)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MailboxMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MailboxCreate) SaveX(ctx context.Context) *Mailbox {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MailboxCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MailboxCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MailboxCreate) defaults() {
	if _, ok := mc.mutation.UIDNext(); !ok {
		v := mailbox.DefaultUIDNext
		mc.mutation.SetUIDNext(v)
	}
	if _, ok := mc.mutation.UIDValidity(); !ok {
		v := mailbox.DefaultUIDValidity
		mc.mutation.SetUIDValidity(v)
	}
	if _, ok := mc.mutation.Subscribed(); !ok {
		v := mailbox.DefaultSubscribed
		mc.mutation.SetSubscribed(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MailboxCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "Mailbox.Name"`)}
	}
	if _, ok := mc.mutation.UIDNext(); !ok {
		return &ValidationError{Name: "UIDNext", err: errors.New(`ent: missing required field "Mailbox.UIDNext"`)}
	}
	if _, ok := mc.mutation.UIDValidity(); !ok {
		return &ValidationError{Name: "UIDValidity", err: errors.New(`ent: missing required field "Mailbox.UIDValidity"`)}
	}
	if _, ok := mc.mutation.Subscribed(); !ok {
		return &ValidationError{Name: "Subscribed", err: errors.New(`ent: missing required field "Mailbox.Subscribed"`)}
	}
	return nil
}

func (mc *MailboxCreate) sqlSave(ctx context.Context) (*Mailbox, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = imap.InternalMailboxID(id)
	}
	return _node, nil
}

func (mc *MailboxCreate) createSpec() (*Mailbox, *sqlgraph.CreateSpec) {
	var (
		_node = &Mailbox{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mailbox.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: mailbox.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.RemoteID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldRemoteID,
		})
		_node.RemoteID = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldName,
		})
		_node.Name = value
	}
	if value, ok := mc.mutation.UIDNext(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDNext,
		})
		_node.UIDNext = value
	}
	if value, ok := mc.mutation.UIDValidity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDValidity,
		})
		_node.UIDValidity = value
	}
	if value, ok := mc.mutation.Subscribed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: mailbox.FieldSubscribed,
		})
		_node.Subscribed = value
	}
	if nodes := mc.mutation.UIDsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.UIDsTable,
			Columns: []string{mailbox.UIDsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.FlagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.FlagsTable,
			Columns: []string{mailbox.FlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.PermanentFlagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.PermanentFlagsTable,
			Columns: []string{mailbox.PermanentFlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxpermflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.AttributesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.AttributesTable,
			Columns: []string{mailbox.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MailboxCreateBulk is the builder for creating many Mailbox entities in bulk.
type MailboxCreateBulk struct {
	config
	builders []*MailboxCreate
}

// Save creates the Mailbox entities in the database.
func (mcb *MailboxCreateBulk) Save(ctx context.Context) ([]*Mailbox, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Mailbox, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MailboxMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = imap.InternalMailboxID(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MailboxCreateBulk) SaveX(ctx context.Context) []*Mailbox {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MailboxCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MailboxCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
