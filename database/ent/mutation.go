// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"portfolio-backend/database/ent/predicate"
	"portfolio-backend/database/ent/project"
	"portfolio-backend/database/ent/team"
	"portfolio-backend/database/ent/user"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeProject = "Project"
	TypeTeam    = "Team"
	TypeUser    = "User"
)

// ProjectMutation represents an operation that mutates the Project nodes in the graph.
type ProjectMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	team          *int
	clearedteam   bool
	done          bool
	oldValue      func(context.Context) (*Project, error)
	predicates    []predicate.Project
}

var _ ent.Mutation = (*ProjectMutation)(nil)

// projectOption allows management of the mutation configuration using functional options.
type projectOption func(*ProjectMutation)

// newProjectMutation creates new mutation for the Project entity.
func newProjectMutation(c config, op Op, opts ...projectOption) *ProjectMutation {
	m := &ProjectMutation{
		config:        c,
		op:            op,
		typ:           TypeProject,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withProjectID sets the ID field of the mutation.
func withProjectID(id int) projectOption {
	return func(m *ProjectMutation) {
		var (
			err   error
			once  sync.Once
			value *Project
		)
		m.oldValue = func(ctx context.Context) (*Project, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Project.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withProject sets the old Project of the mutation.
func withProject(node *Project) projectOption {
	return func(m *ProjectMutation) {
		m.oldValue = func(context.Context) (*Project, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProjectMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProjectMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ProjectMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ProjectMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Project.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *ProjectMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ProjectMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ProjectMutation) ResetName() {
	m.name = nil
}

// SetTeamID sets the "team" edge to the Team entity by id.
func (m *ProjectMutation) SetTeamID(id int) {
	m.team = &id
}

// ClearTeam clears the "team" edge to the Team entity.
func (m *ProjectMutation) ClearTeam() {
	m.clearedteam = true
}

// TeamCleared reports if the "team" edge to the Team entity was cleared.
func (m *ProjectMutation) TeamCleared() bool {
	return m.clearedteam
}

// TeamID returns the "team" edge ID in the mutation.
func (m *ProjectMutation) TeamID() (id int, exists bool) {
	if m.team != nil {
		return *m.team, true
	}
	return
}

// TeamIDs returns the "team" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// TeamID instead. It exists only for internal usage by the builders.
func (m *ProjectMutation) TeamIDs() (ids []int) {
	if id := m.team; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetTeam resets all changes to the "team" edge.
func (m *ProjectMutation) ResetTeam() {
	m.team = nil
	m.clearedteam = false
}

// Where appends a list predicates to the ProjectMutation builder.
func (m *ProjectMutation) Where(ps ...predicate.Project) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ProjectMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ProjectMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Project, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ProjectMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ProjectMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Project).
func (m *ProjectMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ProjectMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, project.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ProjectMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case project.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ProjectMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case project.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Project field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProjectMutation) SetField(name string, value ent.Value) error {
	switch name {
	case project.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Project field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ProjectMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ProjectMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProjectMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Project numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ProjectMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ProjectMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProjectMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Project nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ProjectMutation) ResetField(name string) error {
	switch name {
	case project.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Project field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ProjectMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.team != nil {
		edges = append(edges, project.EdgeTeam)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ProjectMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case project.EdgeTeam:
		if id := m.team; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ProjectMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ProjectMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ProjectMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedteam {
		edges = append(edges, project.EdgeTeam)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ProjectMutation) EdgeCleared(name string) bool {
	switch name {
	case project.EdgeTeam:
		return m.clearedteam
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ProjectMutation) ClearEdge(name string) error {
	switch name {
	case project.EdgeTeam:
		m.ClearTeam()
		return nil
	}
	return fmt.Errorf("unknown Project unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ProjectMutation) ResetEdge(name string) error {
	switch name {
	case project.EdgeTeam:
		m.ResetTeam()
		return nil
	}
	return fmt.Errorf("unknown Project edge %s", name)
}

// TeamMutation represents an operation that mutates the Team nodes in the graph.
type TeamMutation struct {
	config
	op             Op
	typ            string
	id             *int
	name           *string
	clearedFields  map[string]struct{}
	project        map[int]struct{}
	removedproject map[int]struct{}
	clearedproject bool
	users          map[int]struct{}
	removedusers   map[int]struct{}
	clearedusers   bool
	done           bool
	oldValue       func(context.Context) (*Team, error)
	predicates     []predicate.Team
}

var _ ent.Mutation = (*TeamMutation)(nil)

// teamOption allows management of the mutation configuration using functional options.
type teamOption func(*TeamMutation)

// newTeamMutation creates new mutation for the Team entity.
func newTeamMutation(c config, op Op, opts ...teamOption) *TeamMutation {
	m := &TeamMutation{
		config:        c,
		op:            op,
		typ:           TypeTeam,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTeamID sets the ID field of the mutation.
func withTeamID(id int) teamOption {
	return func(m *TeamMutation) {
		var (
			err   error
			once  sync.Once
			value *Team
		)
		m.oldValue = func(ctx context.Context) (*Team, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Team.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTeam sets the old Team of the mutation.
func withTeam(node *Team) teamOption {
	return func(m *TeamMutation) {
		m.oldValue = func(context.Context) (*Team, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TeamMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TeamMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TeamMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TeamMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Team.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *TeamMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TeamMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Team entity.
// If the Team object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TeamMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TeamMutation) ResetName() {
	m.name = nil
}

// AddProjectIDs adds the "project" edge to the Project entity by ids.
func (m *TeamMutation) AddProjectIDs(ids ...int) {
	if m.project == nil {
		m.project = make(map[int]struct{})
	}
	for i := range ids {
		m.project[ids[i]] = struct{}{}
	}
}

// ClearProject clears the "project" edge to the Project entity.
func (m *TeamMutation) ClearProject() {
	m.clearedproject = true
}

// ProjectCleared reports if the "project" edge to the Project entity was cleared.
func (m *TeamMutation) ProjectCleared() bool {
	return m.clearedproject
}

// RemoveProjectIDs removes the "project" edge to the Project entity by IDs.
func (m *TeamMutation) RemoveProjectIDs(ids ...int) {
	if m.removedproject == nil {
		m.removedproject = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.project, ids[i])
		m.removedproject[ids[i]] = struct{}{}
	}
}

// RemovedProject returns the removed IDs of the "project" edge to the Project entity.
func (m *TeamMutation) RemovedProjectIDs() (ids []int) {
	for id := range m.removedproject {
		ids = append(ids, id)
	}
	return
}

// ProjectIDs returns the "project" edge IDs in the mutation.
func (m *TeamMutation) ProjectIDs() (ids []int) {
	for id := range m.project {
		ids = append(ids, id)
	}
	return
}

// ResetProject resets all changes to the "project" edge.
func (m *TeamMutation) ResetProject() {
	m.project = nil
	m.clearedproject = false
	m.removedproject = nil
}

// AddUserIDs adds the "users" edge to the User entity by ids.
func (m *TeamMutation) AddUserIDs(ids ...int) {
	if m.users == nil {
		m.users = make(map[int]struct{})
	}
	for i := range ids {
		m.users[ids[i]] = struct{}{}
	}
}

// ClearUsers clears the "users" edge to the User entity.
func (m *TeamMutation) ClearUsers() {
	m.clearedusers = true
}

// UsersCleared reports if the "users" edge to the User entity was cleared.
func (m *TeamMutation) UsersCleared() bool {
	return m.clearedusers
}

// RemoveUserIDs removes the "users" edge to the User entity by IDs.
func (m *TeamMutation) RemoveUserIDs(ids ...int) {
	if m.removedusers == nil {
		m.removedusers = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.users, ids[i])
		m.removedusers[ids[i]] = struct{}{}
	}
}

// RemovedUsers returns the removed IDs of the "users" edge to the User entity.
func (m *TeamMutation) RemovedUsersIDs() (ids []int) {
	for id := range m.removedusers {
		ids = append(ids, id)
	}
	return
}

// UsersIDs returns the "users" edge IDs in the mutation.
func (m *TeamMutation) UsersIDs() (ids []int) {
	for id := range m.users {
		ids = append(ids, id)
	}
	return
}

// ResetUsers resets all changes to the "users" edge.
func (m *TeamMutation) ResetUsers() {
	m.users = nil
	m.clearedusers = false
	m.removedusers = nil
}

// Where appends a list predicates to the TeamMutation builder.
func (m *TeamMutation) Where(ps ...predicate.Team) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TeamMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TeamMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Team, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TeamMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TeamMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Team).
func (m *TeamMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TeamMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, team.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TeamMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case team.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TeamMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case team.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Team field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TeamMutation) SetField(name string, value ent.Value) error {
	switch name {
	case team.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Team field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TeamMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TeamMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TeamMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Team numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TeamMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TeamMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TeamMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Team nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TeamMutation) ResetField(name string) error {
	switch name {
	case team.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Team field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TeamMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.project != nil {
		edges = append(edges, team.EdgeProject)
	}
	if m.users != nil {
		edges = append(edges, team.EdgeUsers)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TeamMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case team.EdgeProject:
		ids := make([]ent.Value, 0, len(m.project))
		for id := range m.project {
			ids = append(ids, id)
		}
		return ids
	case team.EdgeUsers:
		ids := make([]ent.Value, 0, len(m.users))
		for id := range m.users {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TeamMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedproject != nil {
		edges = append(edges, team.EdgeProject)
	}
	if m.removedusers != nil {
		edges = append(edges, team.EdgeUsers)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TeamMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case team.EdgeProject:
		ids := make([]ent.Value, 0, len(m.removedproject))
		for id := range m.removedproject {
			ids = append(ids, id)
		}
		return ids
	case team.EdgeUsers:
		ids := make([]ent.Value, 0, len(m.removedusers))
		for id := range m.removedusers {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TeamMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedproject {
		edges = append(edges, team.EdgeProject)
	}
	if m.clearedusers {
		edges = append(edges, team.EdgeUsers)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TeamMutation) EdgeCleared(name string) bool {
	switch name {
	case team.EdgeProject:
		return m.clearedproject
	case team.EdgeUsers:
		return m.clearedusers
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TeamMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Team unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TeamMutation) ResetEdge(name string) error {
	switch name {
	case team.EdgeProject:
		m.ResetProject()
		return nil
	case team.EdgeUsers:
		m.ResetUsers()
		return nil
	}
	return fmt.Errorf("unknown Team edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	role          *user.Role
	clearedFields map[string]struct{}
	teams         map[int]struct{}
	removedteams  map[int]struct{}
	clearedteams  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetRole sets the "role" field.
func (m *UserMutation) SetRole(u user.Role) {
	m.role = &u
}

// Role returns the value of the "role" field in the mutation.
func (m *UserMutation) Role() (r user.Role, exists bool) {
	v := m.role
	if v == nil {
		return
	}
	return *v, true
}

// OldRole returns the old "role" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRole(ctx context.Context) (v user.Role, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRole is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRole requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRole: %w", err)
	}
	return oldValue.Role, nil
}

// ResetRole resets all changes to the "role" field.
func (m *UserMutation) ResetRole() {
	m.role = nil
}

// AddTeamIDs adds the "teams" edge to the Team entity by ids.
func (m *UserMutation) AddTeamIDs(ids ...int) {
	if m.teams == nil {
		m.teams = make(map[int]struct{})
	}
	for i := range ids {
		m.teams[ids[i]] = struct{}{}
	}
}

// ClearTeams clears the "teams" edge to the Team entity.
func (m *UserMutation) ClearTeams() {
	m.clearedteams = true
}

// TeamsCleared reports if the "teams" edge to the Team entity was cleared.
func (m *UserMutation) TeamsCleared() bool {
	return m.clearedteams
}

// RemoveTeamIDs removes the "teams" edge to the Team entity by IDs.
func (m *UserMutation) RemoveTeamIDs(ids ...int) {
	if m.removedteams == nil {
		m.removedteams = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.teams, ids[i])
		m.removedteams[ids[i]] = struct{}{}
	}
}

// RemovedTeams returns the removed IDs of the "teams" edge to the Team entity.
func (m *UserMutation) RemovedTeamsIDs() (ids []int) {
	for id := range m.removedteams {
		ids = append(ids, id)
	}
	return
}

// TeamsIDs returns the "teams" edge IDs in the mutation.
func (m *UserMutation) TeamsIDs() (ids []int) {
	for id := range m.teams {
		ids = append(ids, id)
	}
	return
}

// ResetTeams resets all changes to the "teams" edge.
func (m *UserMutation) ResetTeams() {
	m.teams = nil
	m.clearedteams = false
	m.removedteams = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.role != nil {
		fields = append(fields, user.FieldRole)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldRole:
		return m.Role()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldRole:
		return m.OldRole(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldRole:
		v, ok := value.(user.Role)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRole(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldRole:
		m.ResetRole()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.teams != nil {
		edges = append(edges, user.EdgeTeams)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTeams:
		ids := make([]ent.Value, 0, len(m.teams))
		for id := range m.teams {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedteams != nil {
		edges = append(edges, user.EdgeTeams)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTeams:
		ids := make([]ent.Value, 0, len(m.removedteams))
		for id := range m.removedteams {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedteams {
		edges = append(edges, user.EdgeTeams)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeTeams:
		return m.clearedteams
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeTeams:
		m.ResetTeams()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}