// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/item"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/itemfield"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/predicate"
)

// ItemFieldUpdate is the builder for updating ItemField entities.
type ItemFieldUpdate struct {
	config
	hooks    []Hook
	mutation *ItemFieldMutation
}

// Where appends a list predicates to the ItemFieldUpdate builder.
func (ifu *ItemFieldUpdate) Where(ps ...predicate.ItemField) *ItemFieldUpdate {
	ifu.mutation.Where(ps...)
	return ifu
}

// SetUpdatedAt sets the "updated_at" field.
func (ifu *ItemFieldUpdate) SetUpdatedAt(t time.Time) *ItemFieldUpdate {
	ifu.mutation.SetUpdatedAt(t)
	return ifu
}

// SetName sets the "name" field.
func (ifu *ItemFieldUpdate) SetName(s string) *ItemFieldUpdate {
	ifu.mutation.SetName(s)
	return ifu
}

// SetDescription sets the "description" field.
func (ifu *ItemFieldUpdate) SetDescription(s string) *ItemFieldUpdate {
	ifu.mutation.SetDescription(s)
	return ifu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ifu *ItemFieldUpdate) SetNillableDescription(s *string) *ItemFieldUpdate {
	if s != nil {
		ifu.SetDescription(*s)
	}
	return ifu
}

// ClearDescription clears the value of the "description" field.
func (ifu *ItemFieldUpdate) ClearDescription() *ItemFieldUpdate {
	ifu.mutation.ClearDescription()
	return ifu
}

// SetType sets the "type" field.
func (ifu *ItemFieldUpdate) SetType(i itemfield.Type) *ItemFieldUpdate {
	ifu.mutation.SetType(i)
	return ifu
}

// SetTextValue sets the "text_value" field.
func (ifu *ItemFieldUpdate) SetTextValue(s string) *ItemFieldUpdate {
	ifu.mutation.SetTextValue(s)
	return ifu
}

// SetNillableTextValue sets the "text_value" field if the given value is not nil.
func (ifu *ItemFieldUpdate) SetNillableTextValue(s *string) *ItemFieldUpdate {
	if s != nil {
		ifu.SetTextValue(*s)
	}
	return ifu
}

// ClearTextValue clears the value of the "text_value" field.
func (ifu *ItemFieldUpdate) ClearTextValue() *ItemFieldUpdate {
	ifu.mutation.ClearTextValue()
	return ifu
}

// SetNumberValue sets the "number_value" field.
func (ifu *ItemFieldUpdate) SetNumberValue(i int) *ItemFieldUpdate {
	ifu.mutation.ResetNumberValue()
	ifu.mutation.SetNumberValue(i)
	return ifu
}

// SetNillableNumberValue sets the "number_value" field if the given value is not nil.
func (ifu *ItemFieldUpdate) SetNillableNumberValue(i *int) *ItemFieldUpdate {
	if i != nil {
		ifu.SetNumberValue(*i)
	}
	return ifu
}

// AddNumberValue adds i to the "number_value" field.
func (ifu *ItemFieldUpdate) AddNumberValue(i int) *ItemFieldUpdate {
	ifu.mutation.AddNumberValue(i)
	return ifu
}

// ClearNumberValue clears the value of the "number_value" field.
func (ifu *ItemFieldUpdate) ClearNumberValue() *ItemFieldUpdate {
	ifu.mutation.ClearNumberValue()
	return ifu
}

// SetBooleanValue sets the "boolean_value" field.
func (ifu *ItemFieldUpdate) SetBooleanValue(b bool) *ItemFieldUpdate {
	ifu.mutation.SetBooleanValue(b)
	return ifu
}

// SetNillableBooleanValue sets the "boolean_value" field if the given value is not nil.
func (ifu *ItemFieldUpdate) SetNillableBooleanValue(b *bool) *ItemFieldUpdate {
	if b != nil {
		ifu.SetBooleanValue(*b)
	}
	return ifu
}

// SetTimeValue sets the "time_value" field.
func (ifu *ItemFieldUpdate) SetTimeValue(t time.Time) *ItemFieldUpdate {
	ifu.mutation.SetTimeValue(t)
	return ifu
}

// SetNillableTimeValue sets the "time_value" field if the given value is not nil.
func (ifu *ItemFieldUpdate) SetNillableTimeValue(t *time.Time) *ItemFieldUpdate {
	if t != nil {
		ifu.SetTimeValue(*t)
	}
	return ifu
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (ifu *ItemFieldUpdate) SetItemID(id uuid.UUID) *ItemFieldUpdate {
	ifu.mutation.SetItemID(id)
	return ifu
}

// SetNillableItemID sets the "item" edge to the Item entity by ID if the given value is not nil.
func (ifu *ItemFieldUpdate) SetNillableItemID(id *uuid.UUID) *ItemFieldUpdate {
	if id != nil {
		ifu = ifu.SetItemID(*id)
	}
	return ifu
}

// SetItem sets the "item" edge to the Item entity.
func (ifu *ItemFieldUpdate) SetItem(i *Item) *ItemFieldUpdate {
	return ifu.SetItemID(i.ID)
}

// Mutation returns the ItemFieldMutation object of the builder.
func (ifu *ItemFieldUpdate) Mutation() *ItemFieldMutation {
	return ifu.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (ifu *ItemFieldUpdate) ClearItem() *ItemFieldUpdate {
	ifu.mutation.ClearItem()
	return ifu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ifu *ItemFieldUpdate) Save(ctx context.Context) (int, error) {
	ifu.defaults()
	return withHooks[int, ItemFieldMutation](ctx, ifu.sqlSave, ifu.mutation, ifu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ifu *ItemFieldUpdate) SaveX(ctx context.Context) int {
	affected, err := ifu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ifu *ItemFieldUpdate) Exec(ctx context.Context) error {
	_, err := ifu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifu *ItemFieldUpdate) ExecX(ctx context.Context) {
	if err := ifu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ifu *ItemFieldUpdate) defaults() {
	if _, ok := ifu.mutation.UpdatedAt(); !ok {
		v := itemfield.UpdateDefaultUpdatedAt()
		ifu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ifu *ItemFieldUpdate) check() error {
	if v, ok := ifu.mutation.Name(); ok {
		if err := itemfield.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ItemField.name": %w`, err)}
		}
	}
	if v, ok := ifu.mutation.Description(); ok {
		if err := itemfield.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "ItemField.description": %w`, err)}
		}
	}
	if v, ok := ifu.mutation.GetType(); ok {
		if err := itemfield.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ItemField.type": %w`, err)}
		}
	}
	if v, ok := ifu.mutation.TextValue(); ok {
		if err := itemfield.TextValueValidator(v); err != nil {
			return &ValidationError{Name: "text_value", err: fmt.Errorf(`ent: validator failed for field "ItemField.text_value": %w`, err)}
		}
	}
	return nil
}

func (ifu *ItemFieldUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ifu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemfield.Table,
			Columns: itemfield.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itemfield.FieldID,
			},
		},
	}
	if ps := ifu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ifu.mutation.UpdatedAt(); ok {
		_spec.SetField(itemfield.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ifu.mutation.Name(); ok {
		_spec.SetField(itemfield.FieldName, field.TypeString, value)
	}
	if value, ok := ifu.mutation.Description(); ok {
		_spec.SetField(itemfield.FieldDescription, field.TypeString, value)
	}
	if ifu.mutation.DescriptionCleared() {
		_spec.ClearField(itemfield.FieldDescription, field.TypeString)
	}
	if value, ok := ifu.mutation.GetType(); ok {
		_spec.SetField(itemfield.FieldType, field.TypeEnum, value)
	}
	if value, ok := ifu.mutation.TextValue(); ok {
		_spec.SetField(itemfield.FieldTextValue, field.TypeString, value)
	}
	if ifu.mutation.TextValueCleared() {
		_spec.ClearField(itemfield.FieldTextValue, field.TypeString)
	}
	if value, ok := ifu.mutation.NumberValue(); ok {
		_spec.SetField(itemfield.FieldNumberValue, field.TypeInt, value)
	}
	if value, ok := ifu.mutation.AddedNumberValue(); ok {
		_spec.AddField(itemfield.FieldNumberValue, field.TypeInt, value)
	}
	if ifu.mutation.NumberValueCleared() {
		_spec.ClearField(itemfield.FieldNumberValue, field.TypeInt)
	}
	if value, ok := ifu.mutation.BooleanValue(); ok {
		_spec.SetField(itemfield.FieldBooleanValue, field.TypeBool, value)
	}
	if value, ok := ifu.mutation.TimeValue(); ok {
		_spec.SetField(itemfield.FieldTimeValue, field.TypeTime, value)
	}
	if ifu.mutation.ItemCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ifu.mutation.ItemIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ifu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemfield.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ifu.mutation.done = true
	return n, nil
}

// ItemFieldUpdateOne is the builder for updating a single ItemField entity.
type ItemFieldUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemFieldMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ifuo *ItemFieldUpdateOne) SetUpdatedAt(t time.Time) *ItemFieldUpdateOne {
	ifuo.mutation.SetUpdatedAt(t)
	return ifuo
}

// SetName sets the "name" field.
func (ifuo *ItemFieldUpdateOne) SetName(s string) *ItemFieldUpdateOne {
	ifuo.mutation.SetName(s)
	return ifuo
}

// SetDescription sets the "description" field.
func (ifuo *ItemFieldUpdateOne) SetDescription(s string) *ItemFieldUpdateOne {
	ifuo.mutation.SetDescription(s)
	return ifuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ifuo *ItemFieldUpdateOne) SetNillableDescription(s *string) *ItemFieldUpdateOne {
	if s != nil {
		ifuo.SetDescription(*s)
	}
	return ifuo
}

// ClearDescription clears the value of the "description" field.
func (ifuo *ItemFieldUpdateOne) ClearDescription() *ItemFieldUpdateOne {
	ifuo.mutation.ClearDescription()
	return ifuo
}

// SetType sets the "type" field.
func (ifuo *ItemFieldUpdateOne) SetType(i itemfield.Type) *ItemFieldUpdateOne {
	ifuo.mutation.SetType(i)
	return ifuo
}

// SetTextValue sets the "text_value" field.
func (ifuo *ItemFieldUpdateOne) SetTextValue(s string) *ItemFieldUpdateOne {
	ifuo.mutation.SetTextValue(s)
	return ifuo
}

// SetNillableTextValue sets the "text_value" field if the given value is not nil.
func (ifuo *ItemFieldUpdateOne) SetNillableTextValue(s *string) *ItemFieldUpdateOne {
	if s != nil {
		ifuo.SetTextValue(*s)
	}
	return ifuo
}

// ClearTextValue clears the value of the "text_value" field.
func (ifuo *ItemFieldUpdateOne) ClearTextValue() *ItemFieldUpdateOne {
	ifuo.mutation.ClearTextValue()
	return ifuo
}

// SetNumberValue sets the "number_value" field.
func (ifuo *ItemFieldUpdateOne) SetNumberValue(i int) *ItemFieldUpdateOne {
	ifuo.mutation.ResetNumberValue()
	ifuo.mutation.SetNumberValue(i)
	return ifuo
}

// SetNillableNumberValue sets the "number_value" field if the given value is not nil.
func (ifuo *ItemFieldUpdateOne) SetNillableNumberValue(i *int) *ItemFieldUpdateOne {
	if i != nil {
		ifuo.SetNumberValue(*i)
	}
	return ifuo
}

// AddNumberValue adds i to the "number_value" field.
func (ifuo *ItemFieldUpdateOne) AddNumberValue(i int) *ItemFieldUpdateOne {
	ifuo.mutation.AddNumberValue(i)
	return ifuo
}

// ClearNumberValue clears the value of the "number_value" field.
func (ifuo *ItemFieldUpdateOne) ClearNumberValue() *ItemFieldUpdateOne {
	ifuo.mutation.ClearNumberValue()
	return ifuo
}

// SetBooleanValue sets the "boolean_value" field.
func (ifuo *ItemFieldUpdateOne) SetBooleanValue(b bool) *ItemFieldUpdateOne {
	ifuo.mutation.SetBooleanValue(b)
	return ifuo
}

// SetNillableBooleanValue sets the "boolean_value" field if the given value is not nil.
func (ifuo *ItemFieldUpdateOne) SetNillableBooleanValue(b *bool) *ItemFieldUpdateOne {
	if b != nil {
		ifuo.SetBooleanValue(*b)
	}
	return ifuo
}

// SetTimeValue sets the "time_value" field.
func (ifuo *ItemFieldUpdateOne) SetTimeValue(t time.Time) *ItemFieldUpdateOne {
	ifuo.mutation.SetTimeValue(t)
	return ifuo
}

// SetNillableTimeValue sets the "time_value" field if the given value is not nil.
func (ifuo *ItemFieldUpdateOne) SetNillableTimeValue(t *time.Time) *ItemFieldUpdateOne {
	if t != nil {
		ifuo.SetTimeValue(*t)
	}
	return ifuo
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (ifuo *ItemFieldUpdateOne) SetItemID(id uuid.UUID) *ItemFieldUpdateOne {
	ifuo.mutation.SetItemID(id)
	return ifuo
}

// SetNillableItemID sets the "item" edge to the Item entity by ID if the given value is not nil.
func (ifuo *ItemFieldUpdateOne) SetNillableItemID(id *uuid.UUID) *ItemFieldUpdateOne {
	if id != nil {
		ifuo = ifuo.SetItemID(*id)
	}
	return ifuo
}

// SetItem sets the "item" edge to the Item entity.
func (ifuo *ItemFieldUpdateOne) SetItem(i *Item) *ItemFieldUpdateOne {
	return ifuo.SetItemID(i.ID)
}

// Mutation returns the ItemFieldMutation object of the builder.
func (ifuo *ItemFieldUpdateOne) Mutation() *ItemFieldMutation {
	return ifuo.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (ifuo *ItemFieldUpdateOne) ClearItem() *ItemFieldUpdateOne {
	ifuo.mutation.ClearItem()
	return ifuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ifuo *ItemFieldUpdateOne) Select(field string, fields ...string) *ItemFieldUpdateOne {
	ifuo.fields = append([]string{field}, fields...)
	return ifuo
}

// Save executes the query and returns the updated ItemField entity.
func (ifuo *ItemFieldUpdateOne) Save(ctx context.Context) (*ItemField, error) {
	ifuo.defaults()
	return withHooks[*ItemField, ItemFieldMutation](ctx, ifuo.sqlSave, ifuo.mutation, ifuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ifuo *ItemFieldUpdateOne) SaveX(ctx context.Context) *ItemField {
	node, err := ifuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ifuo *ItemFieldUpdateOne) Exec(ctx context.Context) error {
	_, err := ifuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifuo *ItemFieldUpdateOne) ExecX(ctx context.Context) {
	if err := ifuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ifuo *ItemFieldUpdateOne) defaults() {
	if _, ok := ifuo.mutation.UpdatedAt(); !ok {
		v := itemfield.UpdateDefaultUpdatedAt()
		ifuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ifuo *ItemFieldUpdateOne) check() error {
	if v, ok := ifuo.mutation.Name(); ok {
		if err := itemfield.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ItemField.name": %w`, err)}
		}
	}
	if v, ok := ifuo.mutation.Description(); ok {
		if err := itemfield.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "ItemField.description": %w`, err)}
		}
	}
	if v, ok := ifuo.mutation.GetType(); ok {
		if err := itemfield.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ItemField.type": %w`, err)}
		}
	}
	if v, ok := ifuo.mutation.TextValue(); ok {
		if err := itemfield.TextValueValidator(v); err != nil {
			return &ValidationError{Name: "text_value", err: fmt.Errorf(`ent: validator failed for field "ItemField.text_value": %w`, err)}
		}
	}
	return nil
}

func (ifuo *ItemFieldUpdateOne) sqlSave(ctx context.Context) (_node *ItemField, err error) {
	if err := ifuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemfield.Table,
			Columns: itemfield.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itemfield.FieldID,
			},
		},
	}
	id, ok := ifuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ItemField.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ifuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itemfield.FieldID)
		for _, f := range fields {
			if !itemfield.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != itemfield.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ifuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ifuo.mutation.UpdatedAt(); ok {
		_spec.SetField(itemfield.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ifuo.mutation.Name(); ok {
		_spec.SetField(itemfield.FieldName, field.TypeString, value)
	}
	if value, ok := ifuo.mutation.Description(); ok {
		_spec.SetField(itemfield.FieldDescription, field.TypeString, value)
	}
	if ifuo.mutation.DescriptionCleared() {
		_spec.ClearField(itemfield.FieldDescription, field.TypeString)
	}
	if value, ok := ifuo.mutation.GetType(); ok {
		_spec.SetField(itemfield.FieldType, field.TypeEnum, value)
	}
	if value, ok := ifuo.mutation.TextValue(); ok {
		_spec.SetField(itemfield.FieldTextValue, field.TypeString, value)
	}
	if ifuo.mutation.TextValueCleared() {
		_spec.ClearField(itemfield.FieldTextValue, field.TypeString)
	}
	if value, ok := ifuo.mutation.NumberValue(); ok {
		_spec.SetField(itemfield.FieldNumberValue, field.TypeInt, value)
	}
	if value, ok := ifuo.mutation.AddedNumberValue(); ok {
		_spec.AddField(itemfield.FieldNumberValue, field.TypeInt, value)
	}
	if ifuo.mutation.NumberValueCleared() {
		_spec.ClearField(itemfield.FieldNumberValue, field.TypeInt)
	}
	if value, ok := ifuo.mutation.BooleanValue(); ok {
		_spec.SetField(itemfield.FieldBooleanValue, field.TypeBool, value)
	}
	if value, ok := ifuo.mutation.TimeValue(); ok {
		_spec.SetField(itemfield.FieldTimeValue, field.TypeTime, value)
	}
	if ifuo.mutation.ItemCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ifuo.mutation.ItemIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ItemField{config: ifuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ifuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemfield.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ifuo.mutation.done = true
	return _node, nil
}
