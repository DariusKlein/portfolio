// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"portfolio/backend/database/ent/migrate"

	"portfolio/backend/database/ent/project"
	"portfolio/backend/database/ent/team"
	"portfolio/backend/database/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Project is the client for interacting with the Project builders.
	Project *ProjectClient
	// Team is the client for interacting with the Team builders.
	Team *TeamClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Project = NewProjectClient(c.config)
	c.Team = NewTeamClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Project: NewProjectClient(cfg),
		Team:    NewTeamClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Project: NewProjectClient(cfg),
		Team:    NewTeamClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Project.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Project.Use(hooks...)
	c.Team.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Project.Intercept(interceptors...)
	c.Team.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ProjectMutation:
		return c.Project.mutate(ctx, m)
	case *TeamMutation:
		return c.Team.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ProjectClient is a client for the Project schema.
type ProjectClient struct {
	config
}

// NewProjectClient returns a client for the Project from the given config.
func NewProjectClient(c config) *ProjectClient {
	return &ProjectClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `project.Hooks(f(g(h())))`.
func (c *ProjectClient) Use(hooks ...Hook) {
	c.hooks.Project = append(c.hooks.Project, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `project.Intercept(f(g(h())))`.
func (c *ProjectClient) Intercept(interceptors ...Interceptor) {
	c.inters.Project = append(c.inters.Project, interceptors...)
}

// Create returns a builder for creating a Project entity.
func (c *ProjectClient) Create() *ProjectCreate {
	mutation := newProjectMutation(c.config, OpCreate)
	return &ProjectCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Project entities.
func (c *ProjectClient) CreateBulk(builders ...*ProjectCreate) *ProjectCreateBulk {
	return &ProjectCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ProjectClient) MapCreateBulk(slice any, setFunc func(*ProjectCreate, int)) *ProjectCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ProjectCreateBulk{err: fmt.Errorf("calling to ProjectClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ProjectCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ProjectCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Project.
func (c *ProjectClient) Update() *ProjectUpdate {
	mutation := newProjectMutation(c.config, OpUpdate)
	return &ProjectUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProjectClient) UpdateOne(pr *Project) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProject(pr))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProjectClient) UpdateOneID(id int) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProjectID(id))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Project.
func (c *ProjectClient) Delete() *ProjectDelete {
	mutation := newProjectMutation(c.config, OpDelete)
	return &ProjectDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProjectClient) DeleteOne(pr *Project) *ProjectDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProjectClient) DeleteOneID(id int) *ProjectDeleteOne {
	builder := c.Delete().Where(project.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProjectDeleteOne{builder}
}

// Query returns a query builder for Project.
func (c *ProjectClient) Query() *ProjectQuery {
	return &ProjectQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProject},
		inters: c.Interceptors(),
	}
}

// Get returns a Project entity by its id.
func (c *ProjectClient) Get(ctx context.Context, id int) (*Project, error) {
	return c.Query().Where(project.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProjectClient) GetX(ctx context.Context, id int) *Project {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTeam queries the team edge of a Project.
func (c *ProjectClient) QueryTeam(pr *Project) *TeamQuery {
	query := (&TeamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, project.TeamTable, project.TeamColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProjectClient) Hooks() []Hook {
	return c.hooks.Project
}

// Interceptors returns the client interceptors.
func (c *ProjectClient) Interceptors() []Interceptor {
	return c.inters.Project
}

func (c *ProjectClient) mutate(ctx context.Context, m *ProjectMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProjectCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProjectUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProjectDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Project mutation op: %q", m.Op())
	}
}

// TeamClient is a client for the Team schema.
type TeamClient struct {
	config
}

// NewTeamClient returns a client for the Team from the given config.
func NewTeamClient(c config) *TeamClient {
	return &TeamClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `team.Hooks(f(g(h())))`.
func (c *TeamClient) Use(hooks ...Hook) {
	c.hooks.Team = append(c.hooks.Team, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `team.Intercept(f(g(h())))`.
func (c *TeamClient) Intercept(interceptors ...Interceptor) {
	c.inters.Team = append(c.inters.Team, interceptors...)
}

// Create returns a builder for creating a Team entity.
func (c *TeamClient) Create() *TeamCreate {
	mutation := newTeamMutation(c.config, OpCreate)
	return &TeamCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Team entities.
func (c *TeamClient) CreateBulk(builders ...*TeamCreate) *TeamCreateBulk {
	return &TeamCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TeamClient) MapCreateBulk(slice any, setFunc func(*TeamCreate, int)) *TeamCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TeamCreateBulk{err: fmt.Errorf("calling to TeamClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TeamCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TeamCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Team.
func (c *TeamClient) Update() *TeamUpdate {
	mutation := newTeamMutation(c.config, OpUpdate)
	return &TeamUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TeamClient) UpdateOne(t *Team) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeam(t))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TeamClient) UpdateOneID(id int) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeamID(id))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Team.
func (c *TeamClient) Delete() *TeamDelete {
	mutation := newTeamMutation(c.config, OpDelete)
	return &TeamDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TeamClient) DeleteOne(t *Team) *TeamDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TeamClient) DeleteOneID(id int) *TeamDeleteOne {
	builder := c.Delete().Where(team.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TeamDeleteOne{builder}
}

// Query returns a query builder for Team.
func (c *TeamClient) Query() *TeamQuery {
	return &TeamQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTeam},
		inters: c.Interceptors(),
	}
}

// Get returns a Team entity by its id.
func (c *TeamClient) Get(ctx context.Context, id int) (*Team, error) {
	return c.Query().Where(team.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TeamClient) GetX(ctx context.Context, id int) *Team {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProject queries the project edge of a Team.
func (c *TeamClient) QueryProject(t *Team) *ProjectQuery {
	query := (&ProjectClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, team.ProjectTable, team.ProjectColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a Team.
func (c *TeamClient) QueryUsers(t *Team) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, team.UsersTable, team.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TeamClient) Hooks() []Hook {
	return c.hooks.Team
}

// Interceptors returns the client interceptors.
func (c *TeamClient) Interceptors() []Interceptor {
	return c.inters.Team
}

func (c *TeamClient) mutate(ctx context.Context, m *TeamMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TeamCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TeamUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TeamDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Team mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTeams queries the teams edge of a User.
func (c *UserClient) QueryTeams(u *User) *TeamQuery {
	query := (&TeamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.TeamsTable, user.TeamsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Project, Team, User []ent.Hook
	}
	inters struct {
		Project, Team, User []ent.Interceptor
	}
)