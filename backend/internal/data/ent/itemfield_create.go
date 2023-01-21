// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/item"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/itemfield"
)

// ItemFieldCreate is the builder for creating a ItemField entity.
type ItemFieldCreate struct {
	config
	mutation *ItemFieldMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ifc *ItemFieldCreate) SetCreatedAt(t time.Time) *ItemFieldCreate {
	ifc.mutation.SetCreatedAt(t)
	return ifc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ifc *ItemFieldCreate) SetNillableCreatedAt(t *time.Time) *ItemFieldCreate {
	if t != nil {
		ifc.SetCreatedAt(*t)
	}
	return ifc
}

// SetUpdatedAt sets the "updated_at" field.
func (ifc *ItemFieldCreate) SetUpdatedAt(t time.Time) *ItemFieldCreate {
	ifc.mutation.SetUpdatedAt(t)
	return ifc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ifc *ItemFieldCreate) SetNillableUpdatedAt(t *time.Time) *ItemFieldCreate {
	if t != nil {
		ifc.SetUpdatedAt(*t)
	}
	return ifc
}

// SetName sets the "name" field.
func (ifc *ItemFieldCreate) SetName(s string) *ItemFieldCreate {
	ifc.mutation.SetName(s)
	return ifc
}

// SetDescription sets the "description" field.
func (ifc *ItemFieldCreate) SetDescription(s string) *ItemFieldCreate {
	ifc.mutation.SetDescription(s)
	return ifc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ifc *ItemFieldCreate) SetNillableDescription(s *string) *ItemFieldCreate {
	if s != nil {
		ifc.SetDescription(*s)
	}
	return ifc
}

// SetType sets the "type" field.
func (ifc *ItemFieldCreate) SetType(i itemfield.Type) *ItemFieldCreate {
	ifc.mutation.SetType(i)
	return ifc
}

// SetTextValue sets the "text_value" field.
func (ifc *ItemFieldCreate) SetTextValue(s string) *ItemFieldCreate {
	ifc.mutation.SetTextValue(s)
	return ifc
}

// SetNillableTextValue sets the "text_value" field if the given value is not nil.
func (ifc *ItemFieldCreate) SetNillableTextValue(s *string) *ItemFieldCreate {
	if s != nil {
		ifc.SetTextValue(*s)
	}
	return ifc
}

// SetNumberValue sets the "number_value" field.
func (ifc *ItemFieldCreate) SetNumberValue(i int) *ItemFieldCreate {
	ifc.mutation.SetNumberValue(i)
	return ifc
}

// SetNillableNumberValue sets the "number_value" field if the given value is not nil.
func (ifc *ItemFieldCreate) SetNillableNumberValue(i *int) *ItemFieldCreate {
	if i != nil {
		ifc.SetNumberValue(*i)
	}
	return ifc
}

// SetBooleanValue sets the "boolean_value" field.
func (ifc *ItemFieldCreate) SetBooleanValue(b bool) *ItemFieldCreate {
	ifc.mutation.SetBooleanValue(b)
	return ifc
}

// SetNillableBooleanValue sets the "boolean_value" field if the given value is not nil.
func (ifc *ItemFieldCreate) SetNillableBooleanValue(b *bool) *ItemFieldCreate {
	if b != nil {
		ifc.SetBooleanValue(*b)
	}
	return ifc
}

// SetTimeValue sets the "time_value" field.
func (ifc *ItemFieldCreate) SetTimeValue(t time.Time) *ItemFieldCreate {
	ifc.mutation.SetTimeValue(t)
	return ifc
}

// SetNillableTimeValue sets the "time_value" field if the given value is not nil.
func (ifc *ItemFieldCreate) SetNillableTimeValue(t *time.Time) *ItemFieldCreate {
	if t != nil {
		ifc.SetTimeValue(*t)
	}
	return ifc
}

// SetID sets the "id" field.
func (ifc *ItemFieldCreate) SetID(u uuid.UUID) *ItemFieldCreate {
	ifc.mutation.SetID(u)
	return ifc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ifc *ItemFieldCreate) SetNillableID(u *uuid.UUID) *ItemFieldCreate {
	if u != nil {
		ifc.SetID(*u)
	}
	return ifc
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (ifc *ItemFieldCreate) SetItemID(id uuid.UUID) *ItemFieldCreate {
	ifc.mutation.SetItemID(id)
	return ifc
}

// SetNillableItemID sets the "item" edge to the Item entity by ID if the given value is not nil.
func (ifc *ItemFieldCreate) SetNillableItemID(id *uuid.UUID) *ItemFieldCreate {
	if id != nil {
		ifc = ifc.SetItemID(*id)
	}
	return ifc
}

// SetItem sets the "item" edge to the Item entity.
func (ifc *ItemFieldCreate) SetItem(i *Item) *ItemFieldCreate {
	return ifc.SetItemID(i.ID)
}

// Mutation returns the ItemFieldMutation object of the builder.
func (ifc *ItemFieldCreate) Mutation() *ItemFieldMutation {
	return ifc.mutation
}

// Save creates the ItemField in the database.
func (ifc *ItemFieldCreate) Save(ctx context.Context) (*ItemField, error) {
	ifc.defaults()
	return withHooks[*ItemField, ItemFieldMutation](ctx, ifc.sqlSave, ifc.mutation, ifc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ifc *ItemFieldCreate) SaveX(ctx context.Context) *ItemField {
	v, err := ifc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ifc *ItemFieldCreate) Exec(ctx context.Context) error {
	_, err := ifc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifc *ItemFieldCreate) ExecX(ctx context.Context) {
	if err := ifc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ifc *ItemFieldCreate) defaults() {
	if _, ok := ifc.mutation.CreatedAt(); !ok {
		v := itemfield.DefaultCreatedAt()
		ifc.mutation.SetCreatedAt(v)
	}
	if _, ok := ifc.mutation.UpdatedAt(); !ok {
		v := itemfield.DefaultUpdatedAt()
		ifc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ifc.mutation.BooleanValue(); !ok {
		v := itemfield.DefaultBooleanValue
		ifc.mutation.SetBooleanValue(v)
	}
	if _, ok := ifc.mutation.TimeValue(); !ok {
		v := itemfield.DefaultTimeValue()
		ifc.mutation.SetTimeValue(v)
	}
	if _, ok := ifc.mutation.ID(); !ok {
		v := itemfield.DefaultID()
		ifc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ifc *ItemFieldCreate) check() error {
	if _, ok := ifc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ItemField.created_at"`)}
	}
	if _, ok := ifc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ItemField.updated_at"`)}
	}
	if _, ok := ifc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ItemField.name"`)}
	}
	if v, ok := ifc.mutation.Name(); ok {
		if err := itemfield.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ItemField.name": %w`, err)}
		}
	}
	if v, ok := ifc.mutation.Description(); ok {
		if err := itemfield.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "ItemField.description": %w`, err)}
		}
	}
	if _, ok := ifc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ItemField.type"`)}
	}
	if v, ok := ifc.mutation.GetType(); ok {
		if err := itemfield.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ItemField.type": %w`, err)}
		}
	}
	if v, ok := ifc.mutation.TextValue(); ok {
		if err := itemfield.TextValueValidator(v); err != nil {
			return &ValidationError{Name: "text_value", err: fmt.Errorf(`ent: validator failed for field "ItemField.text_value": %w`, err)}
		}
	}
	if _, ok := ifc.mutation.BooleanValue(); !ok {
		return &ValidationError{Name: "boolean_value", err: errors.New(`ent: missing required field "ItemField.boolean_value"`)}
	}
	if _, ok := ifc.mutation.TimeValue(); !ok {
		return &ValidationError{Name: "time_value", err: errors.New(`ent: missing required field "ItemField.time_value"`)}
	}
	return nil
}

func (ifc *ItemFieldCreate) sqlSave(ctx context.Context) (*ItemField, error) {
	if err := ifc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ifc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ifc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ifc.mutation.id = &_node.ID
	ifc.mutation.done = true
	return _node, nil
}

func (ifc *ItemFieldCreate) createSpec() (*ItemField, *sqlgraph.CreateSpec) {
	var (
		_node = &ItemField{config: ifc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: itemfield.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itemfield.FieldID,
			},
		}
	)
	if id, ok := ifc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ifc.mutation.CreatedAt(); ok {
		_spec.SetField(itemfield.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ifc.mutation.UpdatedAt(); ok {
		_spec.SetField(itemfield.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ifc.mutation.Name(); ok {
		_spec.SetField(itemfield.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ifc.mutation.Description(); ok {
		_spec.SetField(itemfield.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ifc.mutation.GetType(); ok {
		_spec.SetField(itemfield.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ifc.mutation.TextValue(); ok {
		_spec.SetField(itemfield.FieldTextValue, field.TypeString, value)
		_node.TextValue = value
	}
	if value, ok := ifc.mutation.NumberValue(); ok {
		_spec.SetField(itemfield.FieldNumberValue, field.TypeInt, value)
		_node.NumberValue = value
	}
	if value, ok := ifc.mutation.BooleanValue(); ok {
		_spec.SetField(itemfield.FieldBooleanValue, field.TypeBool, value)
		_node.BooleanValue = value
	}
	if value, ok := ifc.mutation.TimeValue(); ok {
		_spec.SetField(itemfield.FieldTimeValue, field.TypeTime, value)
		_node.TimeValue = value
	}
	if nodes := ifc.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemfield.ItemTable,
			Columns: []string{itemfield.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.item_fields = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ItemFieldCreateBulk is the builder for creating many ItemField entities in bulk.
type ItemFieldCreateBulk struct {
	config
	builders []*ItemFieldCreate
}

// Save creates the ItemField entities in the database.
func (ifcb *ItemFieldCreateBulk) Save(ctx context.Context) ([]*ItemField, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ifcb.builders))
	nodes := make([]*ItemField, len(ifcb.builders))
	mutators := make([]Mutator, len(ifcb.builders))
	for i := range ifcb.builders {
		func(i int, root context.Context) {
			builder := ifcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemFieldMutation)
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
					_, err = mutators[i+1].Mutate(root, ifcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ifcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ifcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ifcb *ItemFieldCreateBulk) SaveX(ctx context.Context) []*ItemField {
	v, err := ifcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ifcb *ItemFieldCreateBulk) Exec(ctx context.Context) error {
	_, err := ifcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifcb *ItemFieldCreateBulk) ExecX(ctx context.Context) {
	if err := ifcb.Exec(ctx); err != nil {
		panic(err)
	}
}
