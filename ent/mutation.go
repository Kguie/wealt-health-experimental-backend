// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"wealth-health-backend/ent/employee"
	"wealth-health-backend/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEmployee = "Employee"
)

// EmployeeMutation represents an operation that mutates the Employee nodes in the graph.
type EmployeeMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	firstName     *string
	lastName      *string
	dateOfBirth   *time.Time
	startDate     *time.Time
	department    *employee.Department
	street        *string
	city          *string
	state         *employee.State
	zipCode       *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Employee, error)
	predicates    []predicate.Employee
}

var _ ent.Mutation = (*EmployeeMutation)(nil)

// employeeOption allows management of the mutation configuration using functional options.
type employeeOption func(*EmployeeMutation)

// newEmployeeMutation creates new mutation for the Employee entity.
func newEmployeeMutation(c config, op Op, opts ...employeeOption) *EmployeeMutation {
	m := &EmployeeMutation{
		config:        c,
		op:            op,
		typ:           TypeEmployee,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEmployeeID sets the ID field of the mutation.
func withEmployeeID(id uuid.UUID) employeeOption {
	return func(m *EmployeeMutation) {
		var (
			err   error
			once  sync.Once
			value *Employee
		)
		m.oldValue = func(ctx context.Context) (*Employee, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Employee.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEmployee sets the old Employee of the mutation.
func withEmployee(node *Employee) employeeOption {
	return func(m *EmployeeMutation) {
		m.oldValue = func(context.Context) (*Employee, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EmployeeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EmployeeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Employee entities.
func (m *EmployeeMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EmployeeMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EmployeeMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Employee.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetFirstName sets the "firstName" field.
func (m *EmployeeMutation) SetFirstName(s string) {
	m.firstName = &s
}

// FirstName returns the value of the "firstName" field in the mutation.
func (m *EmployeeMutation) FirstName() (r string, exists bool) {
	v := m.firstName
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "firstName" field's value of the Employee entity.
// If the Employee object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmployeeMutation) OldFirstName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ResetFirstName resets all changes to the "firstName" field.
func (m *EmployeeMutation) ResetFirstName() {
	m.firstName = nil
}

// SetLastName sets the "lastName" field.
func (m *EmployeeMutation) SetLastName(s string) {
	m.lastName = &s
}

// LastName returns the value of the "lastName" field in the mutation.
func (m *EmployeeMutation) LastName() (r string, exists bool) {
	v := m.lastName
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "lastName" field's value of the Employee entity.
// If the Employee object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmployeeMutation) OldLastName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ResetLastName resets all changes to the "lastName" field.
func (m *EmployeeMutation) ResetLastName() {
	m.lastName = nil
}

// SetDateOfBirth sets the "dateOfBirth" field.
func (m *EmployeeMutation) SetDateOfBirth(t time.Time) {
	m.dateOfBirth = &t
}

// DateOfBirth returns the value of the "dateOfBirth" field in the mutation.
func (m *EmployeeMutation) DateOfBirth() (r time.Time, exists bool) {
	v := m.dateOfBirth
	if v == nil {
		return
	}
	return *v, true
}

// OldDateOfBirth returns the old "dateOfBirth" field's value of the Employee entity.
// If the Employee object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmployeeMutation) OldDateOfBirth(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDateOfBirth is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDateOfBirth requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDateOfBirth: %w", err)
	}
	return oldValue.DateOfBirth, nil
}

// ResetDateOfBirth resets all changes to the "dateOfBirth" field.
func (m *EmployeeMutation) ResetDateOfBirth() {
	m.dateOfBirth = nil
}

// SetStartDate sets the "startDate" field.
func (m *EmployeeMutation) SetStartDate(t time.Time) {
	m.startDate = &t
}

// StartDate returns the value of the "startDate" field in the mutation.
func (m *EmployeeMutation) StartDate() (r time.Time, exists bool) {
	v := m.startDate
	if v == nil {
		return
	}
	return *v, true
}

// OldStartDate returns the old "startDate" field's value of the Employee entity.
// If the Employee object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmployeeMutation) OldStartDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStartDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStartDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStartDate: %w", err)
	}
	return oldValue.StartDate, nil
}

// ResetStartDate resets all changes to the "startDate" field.
func (m *EmployeeMutation) ResetStartDate() {
	m.startDate = nil
}

// SetDepartment sets the "department" field.
func (m *EmployeeMutation) SetDepartment(e employee.Department) {
	m.department = &e
}

// Department returns the value of the "department" field in the mutation.
func (m *EmployeeMutation) Department() (r employee.Department, exists bool) {
	v := m.department
	if v == nil {
		return
	}
	return *v, true
}

// OldDepartment returns the old "department" field's value of the Employee entity.
// If the Employee object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmployeeMutation) OldDepartment(ctx context.Context) (v employee.Department, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDepartment is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDepartment requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDepartment: %w", err)
	}
	return oldValue.Department, nil
}

// ResetDepartment resets all changes to the "department" field.
func (m *EmployeeMutation) ResetDepartment() {
	m.department = nil
}

// SetStreet sets the "street" field.
func (m *EmployeeMutation) SetStreet(s string) {
	m.street = &s
}

// Street returns the value of the "street" field in the mutation.
func (m *EmployeeMutation) Street() (r string, exists bool) {
	v := m.street
	if v == nil {
		return
	}
	return *v, true
}

// OldStreet returns the old "street" field's value of the Employee entity.
// If the Employee object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmployeeMutation) OldStreet(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStreet is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStreet requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStreet: %w", err)
	}
	return oldValue.Street, nil
}

// ResetStreet resets all changes to the "street" field.
func (m *EmployeeMutation) ResetStreet() {
	m.street = nil
}

// SetCity sets the "city" field.
func (m *EmployeeMutation) SetCity(s string) {
	m.city = &s
}

// City returns the value of the "city" field in the mutation.
func (m *EmployeeMutation) City() (r string, exists bool) {
	v := m.city
	if v == nil {
		return
	}
	return *v, true
}

// OldCity returns the old "city" field's value of the Employee entity.
// If the Employee object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmployeeMutation) OldCity(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCity is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCity requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCity: %w", err)
	}
	return oldValue.City, nil
}

// ResetCity resets all changes to the "city" field.
func (m *EmployeeMutation) ResetCity() {
	m.city = nil
}

// SetState sets the "state" field.
func (m *EmployeeMutation) SetState(e employee.State) {
	m.state = &e
}

// State returns the value of the "state" field in the mutation.
func (m *EmployeeMutation) State() (r employee.State, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old "state" field's value of the Employee entity.
// If the Employee object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmployeeMutation) OldState(ctx context.Context) (v employee.State, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldState is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// ResetState resets all changes to the "state" field.
func (m *EmployeeMutation) ResetState() {
	m.state = nil
}

// SetZipCode sets the "zipCode" field.
func (m *EmployeeMutation) SetZipCode(s string) {
	m.zipCode = &s
}

// ZipCode returns the value of the "zipCode" field in the mutation.
func (m *EmployeeMutation) ZipCode() (r string, exists bool) {
	v := m.zipCode
	if v == nil {
		return
	}
	return *v, true
}

// OldZipCode returns the old "zipCode" field's value of the Employee entity.
// If the Employee object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmployeeMutation) OldZipCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldZipCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldZipCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldZipCode: %w", err)
	}
	return oldValue.ZipCode, nil
}

// ResetZipCode resets all changes to the "zipCode" field.
func (m *EmployeeMutation) ResetZipCode() {
	m.zipCode = nil
}

// Where appends a list predicates to the EmployeeMutation builder.
func (m *EmployeeMutation) Where(ps ...predicate.Employee) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the EmployeeMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *EmployeeMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Employee, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *EmployeeMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *EmployeeMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Employee).
func (m *EmployeeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EmployeeMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.firstName != nil {
		fields = append(fields, employee.FieldFirstName)
	}
	if m.lastName != nil {
		fields = append(fields, employee.FieldLastName)
	}
	if m.dateOfBirth != nil {
		fields = append(fields, employee.FieldDateOfBirth)
	}
	if m.startDate != nil {
		fields = append(fields, employee.FieldStartDate)
	}
	if m.department != nil {
		fields = append(fields, employee.FieldDepartment)
	}
	if m.street != nil {
		fields = append(fields, employee.FieldStreet)
	}
	if m.city != nil {
		fields = append(fields, employee.FieldCity)
	}
	if m.state != nil {
		fields = append(fields, employee.FieldState)
	}
	if m.zipCode != nil {
		fields = append(fields, employee.FieldZipCode)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EmployeeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case employee.FieldFirstName:
		return m.FirstName()
	case employee.FieldLastName:
		return m.LastName()
	case employee.FieldDateOfBirth:
		return m.DateOfBirth()
	case employee.FieldStartDate:
		return m.StartDate()
	case employee.FieldDepartment:
		return m.Department()
	case employee.FieldStreet:
		return m.Street()
	case employee.FieldCity:
		return m.City()
	case employee.FieldState:
		return m.State()
	case employee.FieldZipCode:
		return m.ZipCode()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EmployeeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case employee.FieldFirstName:
		return m.OldFirstName(ctx)
	case employee.FieldLastName:
		return m.OldLastName(ctx)
	case employee.FieldDateOfBirth:
		return m.OldDateOfBirth(ctx)
	case employee.FieldStartDate:
		return m.OldStartDate(ctx)
	case employee.FieldDepartment:
		return m.OldDepartment(ctx)
	case employee.FieldStreet:
		return m.OldStreet(ctx)
	case employee.FieldCity:
		return m.OldCity(ctx)
	case employee.FieldState:
		return m.OldState(ctx)
	case employee.FieldZipCode:
		return m.OldZipCode(ctx)
	}
	return nil, fmt.Errorf("unknown Employee field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmployeeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case employee.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case employee.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case employee.FieldDateOfBirth:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDateOfBirth(v)
		return nil
	case employee.FieldStartDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStartDate(v)
		return nil
	case employee.FieldDepartment:
		v, ok := value.(employee.Department)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDepartment(v)
		return nil
	case employee.FieldStreet:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStreet(v)
		return nil
	case employee.FieldCity:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCity(v)
		return nil
	case employee.FieldState:
		v, ok := value.(employee.State)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	case employee.FieldZipCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetZipCode(v)
		return nil
	}
	return fmt.Errorf("unknown Employee field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EmployeeMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EmployeeMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmployeeMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Employee numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EmployeeMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EmployeeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EmployeeMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Employee nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EmployeeMutation) ResetField(name string) error {
	switch name {
	case employee.FieldFirstName:
		m.ResetFirstName()
		return nil
	case employee.FieldLastName:
		m.ResetLastName()
		return nil
	case employee.FieldDateOfBirth:
		m.ResetDateOfBirth()
		return nil
	case employee.FieldStartDate:
		m.ResetStartDate()
		return nil
	case employee.FieldDepartment:
		m.ResetDepartment()
		return nil
	case employee.FieldStreet:
		m.ResetStreet()
		return nil
	case employee.FieldCity:
		m.ResetCity()
		return nil
	case employee.FieldState:
		m.ResetState()
		return nil
	case employee.FieldZipCode:
		m.ResetZipCode()
		return nil
	}
	return fmt.Errorf("unknown Employee field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EmployeeMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EmployeeMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EmployeeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EmployeeMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EmployeeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EmployeeMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EmployeeMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Employee unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EmployeeMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Employee edge %s", name)
}
