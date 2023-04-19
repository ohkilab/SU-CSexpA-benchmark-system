// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/tagresult"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Contest is the client for interacting with the Contest builders.
	Contest *ContestClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Submit is the client for interacting with the Submit builders.
	Submit *SubmitClient
	// TagResult is the client for interacting with the TagResult builders.
	TagResult *TagResultClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Contest = NewContestClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Submit = NewSubmitClient(c.config)
	c.TagResult = NewTagResultClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Contest:   NewContestClient(cfg),
		Group:     NewGroupClient(cfg),
		Submit:    NewSubmitClient(cfg),
		TagResult: NewTagResultClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		Contest:   NewContestClient(cfg),
		Group:     NewGroupClient(cfg),
		Submit:    NewSubmitClient(cfg),
		TagResult: NewTagResultClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Contest.
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
	c.Contest.Use(hooks...)
	c.Group.Use(hooks...)
	c.Submit.Use(hooks...)
	c.TagResult.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Contest.Intercept(interceptors...)
	c.Group.Intercept(interceptors...)
	c.Submit.Intercept(interceptors...)
	c.TagResult.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ContestMutation:
		return c.Contest.mutate(ctx, m)
	case *GroupMutation:
		return c.Group.mutate(ctx, m)
	case *SubmitMutation:
		return c.Submit.mutate(ctx, m)
	case *TagResultMutation:
		return c.TagResult.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ContestClient is a client for the Contest schema.
type ContestClient struct {
	config
}

// NewContestClient returns a client for the Contest from the given config.
func NewContestClient(c config) *ContestClient {
	return &ContestClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contest.Hooks(f(g(h())))`.
func (c *ContestClient) Use(hooks ...Hook) {
	c.hooks.Contest = append(c.hooks.Contest, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `contest.Intercept(f(g(h())))`.
func (c *ContestClient) Intercept(interceptors ...Interceptor) {
	c.inters.Contest = append(c.inters.Contest, interceptors...)
}

// Create returns a builder for creating a Contest entity.
func (c *ContestClient) Create() *ContestCreate {
	mutation := newContestMutation(c.config, OpCreate)
	return &ContestCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Contest entities.
func (c *ContestClient) CreateBulk(builders ...*ContestCreate) *ContestCreateBulk {
	return &ContestCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Contest.
func (c *ContestClient) Update() *ContestUpdate {
	mutation := newContestMutation(c.config, OpUpdate)
	return &ContestUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContestClient) UpdateOne(co *Contest) *ContestUpdateOne {
	mutation := newContestMutation(c.config, OpUpdateOne, withContest(co))
	return &ContestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContestClient) UpdateOneID(id int) *ContestUpdateOne {
	mutation := newContestMutation(c.config, OpUpdateOne, withContestID(id))
	return &ContestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Contest.
func (c *ContestClient) Delete() *ContestDelete {
	mutation := newContestMutation(c.config, OpDelete)
	return &ContestDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContestClient) DeleteOne(co *Contest) *ContestDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ContestClient) DeleteOneID(id int) *ContestDeleteOne {
	builder := c.Delete().Where(contest.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContestDeleteOne{builder}
}

// Query returns a query builder for Contest.
func (c *ContestClient) Query() *ContestQuery {
	return &ContestQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeContest},
		inters: c.Interceptors(),
	}
}

// Get returns a Contest entity by its id.
func (c *ContestClient) Get(ctx context.Context, id int) (*Contest, error) {
	return c.Query().Where(contest.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContestClient) GetX(ctx context.Context, id int) *Contest {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ContestClient) Hooks() []Hook {
	return c.hooks.Contest
}

// Interceptors returns the client interceptors.
func (c *ContestClient) Interceptors() []Interceptor {
	return c.inters.Contest
}

func (c *ContestClient) mutate(ctx context.Context, m *ContestMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ContestCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ContestUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ContestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ContestDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Contest mutation op: %q", m.Op())
	}
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `group.Intercept(f(g(h())))`.
func (c *GroupClient) Intercept(interceptors ...Interceptor) {
	c.inters.Group = append(c.inters.Group, interceptors...)
}

// Create returns a builder for creating a Group entity.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id string) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GroupClient) DeleteOneID(id string) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGroup},
		inters: c.Interceptors(),
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id string) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id string) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySubmits queries the submits edge of a Group.
func (c *GroupClient) QuerySubmits(gr *Group) *SubmitQuery {
	query := (&SubmitClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(submit.Table, submit.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.SubmitsTable, group.SubmitsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// Interceptors returns the client interceptors.
func (c *GroupClient) Interceptors() []Interceptor {
	return c.inters.Group
}

func (c *GroupClient) mutate(ctx context.Context, m *GroupMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GroupCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GroupDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Group mutation op: %q", m.Op())
	}
}

// SubmitClient is a client for the Submit schema.
type SubmitClient struct {
	config
}

// NewSubmitClient returns a client for the Submit from the given config.
func NewSubmitClient(c config) *SubmitClient {
	return &SubmitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `submit.Hooks(f(g(h())))`.
func (c *SubmitClient) Use(hooks ...Hook) {
	c.hooks.Submit = append(c.hooks.Submit, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `submit.Intercept(f(g(h())))`.
func (c *SubmitClient) Intercept(interceptors ...Interceptor) {
	c.inters.Submit = append(c.inters.Submit, interceptors...)
}

// Create returns a builder for creating a Submit entity.
func (c *SubmitClient) Create() *SubmitCreate {
	mutation := newSubmitMutation(c.config, OpCreate)
	return &SubmitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Submit entities.
func (c *SubmitClient) CreateBulk(builders ...*SubmitCreate) *SubmitCreateBulk {
	return &SubmitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Submit.
func (c *SubmitClient) Update() *SubmitUpdate {
	mutation := newSubmitMutation(c.config, OpUpdate)
	return &SubmitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubmitClient) UpdateOne(s *Submit) *SubmitUpdateOne {
	mutation := newSubmitMutation(c.config, OpUpdateOne, withSubmit(s))
	return &SubmitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubmitClient) UpdateOneID(id string) *SubmitUpdateOne {
	mutation := newSubmitMutation(c.config, OpUpdateOne, withSubmitID(id))
	return &SubmitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Submit.
func (c *SubmitClient) Delete() *SubmitDelete {
	mutation := newSubmitMutation(c.config, OpDelete)
	return &SubmitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SubmitClient) DeleteOne(s *Submit) *SubmitDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SubmitClient) DeleteOneID(id string) *SubmitDeleteOne {
	builder := c.Delete().Where(submit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubmitDeleteOne{builder}
}

// Query returns a query builder for Submit.
func (c *SubmitClient) Query() *SubmitQuery {
	return &SubmitQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSubmit},
		inters: c.Interceptors(),
	}
}

// Get returns a Submit entity by its id.
func (c *SubmitClient) Get(ctx context.Context, id string) (*Submit, error) {
	return c.Query().Where(submit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubmitClient) GetX(ctx context.Context, id string) *Submit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTagResults queries the tagResults edge of a Submit.
func (c *SubmitClient) QueryTagResults(s *Submit) *TagResultQuery {
	query := (&TagResultClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(submit.Table, submit.FieldID, id),
			sqlgraph.To(tagresult.Table, tagresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, submit.TagResultsTable, submit.TagResultsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroup queries the group edge of a Submit.
func (c *SubmitClient) QueryGroup(s *Submit) *GroupQuery {
	query := (&GroupClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(submit.Table, submit.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, submit.GroupTable, submit.GroupPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubmitClient) Hooks() []Hook {
	return c.hooks.Submit
}

// Interceptors returns the client interceptors.
func (c *SubmitClient) Interceptors() []Interceptor {
	return c.inters.Submit
}

func (c *SubmitClient) mutate(ctx context.Context, m *SubmitMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SubmitCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SubmitUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SubmitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SubmitDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Submit mutation op: %q", m.Op())
	}
}

// TagResultClient is a client for the TagResult schema.
type TagResultClient struct {
	config
}

// NewTagResultClient returns a client for the TagResult from the given config.
func NewTagResultClient(c config) *TagResultClient {
	return &TagResultClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tagresult.Hooks(f(g(h())))`.
func (c *TagResultClient) Use(hooks ...Hook) {
	c.hooks.TagResult = append(c.hooks.TagResult, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `tagresult.Intercept(f(g(h())))`.
func (c *TagResultClient) Intercept(interceptors ...Interceptor) {
	c.inters.TagResult = append(c.inters.TagResult, interceptors...)
}

// Create returns a builder for creating a TagResult entity.
func (c *TagResultClient) Create() *TagResultCreate {
	mutation := newTagResultMutation(c.config, OpCreate)
	return &TagResultCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TagResult entities.
func (c *TagResultClient) CreateBulk(builders ...*TagResultCreate) *TagResultCreateBulk {
	return &TagResultCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TagResult.
func (c *TagResultClient) Update() *TagResultUpdate {
	mutation := newTagResultMutation(c.config, OpUpdate)
	return &TagResultUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagResultClient) UpdateOne(tr *TagResult) *TagResultUpdateOne {
	mutation := newTagResultMutation(c.config, OpUpdateOne, withTagResult(tr))
	return &TagResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagResultClient) UpdateOneID(id int) *TagResultUpdateOne {
	mutation := newTagResultMutation(c.config, OpUpdateOne, withTagResultID(id))
	return &TagResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TagResult.
func (c *TagResultClient) Delete() *TagResultDelete {
	mutation := newTagResultMutation(c.config, OpDelete)
	return &TagResultDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TagResultClient) DeleteOne(tr *TagResult) *TagResultDeleteOne {
	return c.DeleteOneID(tr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TagResultClient) DeleteOneID(id int) *TagResultDeleteOne {
	builder := c.Delete().Where(tagresult.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagResultDeleteOne{builder}
}

// Query returns a query builder for TagResult.
func (c *TagResultClient) Query() *TagResultQuery {
	return &TagResultQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTagResult},
		inters: c.Interceptors(),
	}
}

// Get returns a TagResult entity by its id.
func (c *TagResultClient) Get(ctx context.Context, id int) (*TagResult, error) {
	return c.Query().Where(tagresult.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagResultClient) GetX(ctx context.Context, id int) *TagResult {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TagResultClient) Hooks() []Hook {
	return c.hooks.TagResult
}

// Interceptors returns the client interceptors.
func (c *TagResultClient) Interceptors() []Interceptor {
	return c.inters.TagResult
}

func (c *TagResultClient) mutate(ctx context.Context, m *TagResultMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TagResultCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TagResultUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TagResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TagResultDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TagResult mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Contest, Group, Submit, TagResult []ent.Hook
	}
	inters struct {
		Contest, Group, Submit, TagResult []ent.Interceptor
	}
)