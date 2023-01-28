// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/authroles"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/authtokens"
)

// AuthRoles is the model entity for the AuthRoles schema.
type AuthRoles struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Role holds the value of the "role" field.
	Role authroles.Role `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthRolesQuery when eager-loading is set.
	Edges             AuthRolesEdges `json:"edges"`
	auth_tokens_roles *uuid.UUID
}

// AuthRolesEdges holds the relations/edges for other nodes in the graph.
type AuthRolesEdges struct {
	// Token holds the value of the token edge.
	Token *AuthTokens `json:"token,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TokenOrErr returns the Token value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuthRolesEdges) TokenOrErr() (*AuthTokens, error) {
	if e.loadedTypes[0] {
		if e.Token == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: authtokens.Label}
		}
		return e.Token, nil
	}
	return nil, &NotLoadedError{edge: "token"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthRoles) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authroles.FieldID:
			values[i] = new(sql.NullInt64)
		case authroles.FieldRole:
			values[i] = new(sql.NullString)
		case authroles.ForeignKeys[0]: // auth_tokens_roles
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type AuthRoles", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthRoles fields.
func (ar *AuthRoles) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authroles.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ar.ID = int(value.Int64)
		case authroles.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				ar.Role = authroles.Role(value.String)
			}
		case authroles.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field auth_tokens_roles", values[i])
			} else if value.Valid {
				ar.auth_tokens_roles = new(uuid.UUID)
				*ar.auth_tokens_roles = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryToken queries the "token" edge of the AuthRoles entity.
func (ar *AuthRoles) QueryToken() *AuthTokensQuery {
	return NewAuthRolesClient(ar.config).QueryToken(ar)
}

// Update returns a builder for updating this AuthRoles.
// Note that you need to call AuthRoles.Unwrap() before calling this method if this AuthRoles
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AuthRoles) Update() *AuthRolesUpdateOne {
	return NewAuthRolesClient(ar.config).UpdateOne(ar)
}

// Unwrap unwraps the AuthRoles entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *AuthRoles) Unwrap() *AuthRoles {
	_tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuthRoles is not a transactional entity")
	}
	ar.config.driver = _tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AuthRoles) String() string {
	var builder strings.Builder
	builder.WriteString("AuthRoles(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ar.ID))
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", ar.Role))
	builder.WriteByte(')')
	return builder.String()
}

// AuthRolesSlice is a parsable slice of AuthRoles.
type AuthRolesSlice []*AuthRoles

func (ar AuthRolesSlice) config(cfg config) {
	for _i := range ar {
		ar[_i].config = cfg
	}
}
