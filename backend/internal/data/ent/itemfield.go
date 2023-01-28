// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/item"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/itemfield"
)

// ItemField is the model entity for the ItemField schema.
type ItemField struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Type holds the value of the "type" field.
	Type itemfield.Type `json:"type,omitempty"`
	// TextValue holds the value of the "text_value" field.
	TextValue string `json:"text_value,omitempty"`
	// NumberValue holds the value of the "number_value" field.
	NumberValue int `json:"number_value,omitempty"`
	// BooleanValue holds the value of the "boolean_value" field.
	BooleanValue bool `json:"boolean_value,omitempty"`
	// TimeValue holds the value of the "time_value" field.
	TimeValue time.Time `json:"time_value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemFieldQuery when eager-loading is set.
	Edges       ItemFieldEdges `json:"edges"`
	item_fields *uuid.UUID
}

// ItemFieldEdges holds the relations/edges for other nodes in the graph.
type ItemFieldEdges struct {
	// Item holds the value of the item edge.
	Item *Item `json:"item,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemFieldEdges) ItemOrErr() (*Item, error) {
	if e.loadedTypes[0] {
		if e.Item == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: item.Label}
		}
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ItemField) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case itemfield.FieldBooleanValue:
			values[i] = new(sql.NullBool)
		case itemfield.FieldNumberValue:
			values[i] = new(sql.NullInt64)
		case itemfield.FieldName, itemfield.FieldDescription, itemfield.FieldType, itemfield.FieldTextValue:
			values[i] = new(sql.NullString)
		case itemfield.FieldCreatedAt, itemfield.FieldUpdatedAt, itemfield.FieldTimeValue:
			values[i] = new(sql.NullTime)
		case itemfield.FieldID:
			values[i] = new(uuid.UUID)
		case itemfield.ForeignKeys[0]: // item_fields
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ItemField", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ItemField fields.
func (_if *ItemField) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case itemfield.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				_if.ID = *value
			}
		case itemfield.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				_if.CreatedAt = value.Time
			}
		case itemfield.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				_if.UpdatedAt = value.Time
			}
		case itemfield.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				_if.Name = value.String
			}
		case itemfield.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				_if.Description = value.String
			}
		case itemfield.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				_if.Type = itemfield.Type(value.String)
			}
		case itemfield.FieldTextValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text_value", values[i])
			} else if value.Valid {
				_if.TextValue = value.String
			}
		case itemfield.FieldNumberValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number_value", values[i])
			} else if value.Valid {
				_if.NumberValue = int(value.Int64)
			}
		case itemfield.FieldBooleanValue:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field boolean_value", values[i])
			} else if value.Valid {
				_if.BooleanValue = value.Bool
			}
		case itemfield.FieldTimeValue:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time_value", values[i])
			} else if value.Valid {
				_if.TimeValue = value.Time
			}
		case itemfield.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field item_fields", values[i])
			} else if value.Valid {
				_if.item_fields = new(uuid.UUID)
				*_if.item_fields = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryItem queries the "item" edge of the ItemField entity.
func (_if *ItemField) QueryItem() *ItemQuery {
	return NewItemFieldClient(_if.config).QueryItem(_if)
}

// Update returns a builder for updating this ItemField.
// Note that you need to call ItemField.Unwrap() before calling this method if this ItemField
// was returned from a transaction, and the transaction was committed or rolled back.
func (_if *ItemField) Update() *ItemFieldUpdateOne {
	return NewItemFieldClient(_if.config).UpdateOne(_if)
}

// Unwrap unwraps the ItemField entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_if *ItemField) Unwrap() *ItemField {
	_tx, ok := _if.config.driver.(*txDriver)
	if !ok {
		panic("ent: ItemField is not a transactional entity")
	}
	_if.config.driver = _tx.drv
	return _if
}

// String implements the fmt.Stringer.
func (_if *ItemField) String() string {
	var builder strings.Builder
	builder.WriteString("ItemField(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _if.ID))
	builder.WriteString("created_at=")
	builder.WriteString(_if.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(_if.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(_if.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(_if.Description)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", _if.Type))
	builder.WriteString(", ")
	builder.WriteString("text_value=")
	builder.WriteString(_if.TextValue)
	builder.WriteString(", ")
	builder.WriteString("number_value=")
	builder.WriteString(fmt.Sprintf("%v", _if.NumberValue))
	builder.WriteString(", ")
	builder.WriteString("boolean_value=")
	builder.WriteString(fmt.Sprintf("%v", _if.BooleanValue))
	builder.WriteString(", ")
	builder.WriteString("time_value=")
	builder.WriteString(_if.TimeValue.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ItemFields is a parsable slice of ItemField.
type ItemFields []*ItemField

func (_if ItemFields) config(cfg config) {
	for _i := range _if {
		_if[_i].config = cfg
	}
}
