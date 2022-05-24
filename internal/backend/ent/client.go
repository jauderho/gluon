// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/ProtonMail/gluon/internal/backend/ent/migrate"

	"github.com/ProtonMail/gluon/internal/backend/ent/mailbox"
	"github.com/ProtonMail/gluon/internal/backend/ent/mailboxattr"
	"github.com/ProtonMail/gluon/internal/backend/ent/mailboxflag"
	"github.com/ProtonMail/gluon/internal/backend/ent/mailboxpermflag"
	"github.com/ProtonMail/gluon/internal/backend/ent/message"
	"github.com/ProtonMail/gluon/internal/backend/ent/messageflag"
	"github.com/ProtonMail/gluon/internal/backend/ent/uid"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Mailbox is the client for interacting with the Mailbox builders.
	Mailbox *MailboxClient
	// MailboxAttr is the client for interacting with the MailboxAttr builders.
	MailboxAttr *MailboxAttrClient
	// MailboxFlag is the client for interacting with the MailboxFlag builders.
	MailboxFlag *MailboxFlagClient
	// MailboxPermFlag is the client for interacting with the MailboxPermFlag builders.
	MailboxPermFlag *MailboxPermFlagClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// MessageFlag is the client for interacting with the MessageFlag builders.
	MessageFlag *MessageFlagClient
	// UID is the client for interacting with the UID builders.
	UID *UIDClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Mailbox = NewMailboxClient(c.config)
	c.MailboxAttr = NewMailboxAttrClient(c.config)
	c.MailboxFlag = NewMailboxFlagClient(c.config)
	c.MailboxPermFlag = NewMailboxPermFlagClient(c.config)
	c.Message = NewMessageClient(c.config)
	c.MessageFlag = NewMessageFlagClient(c.config)
	c.UID = NewUIDClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Mailbox:         NewMailboxClient(cfg),
		MailboxAttr:     NewMailboxAttrClient(cfg),
		MailboxFlag:     NewMailboxFlagClient(cfg),
		MailboxPermFlag: NewMailboxPermFlagClient(cfg),
		Message:         NewMessageClient(cfg),
		MessageFlag:     NewMessageFlagClient(cfg),
		UID:             NewUIDClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:             ctx,
		config:          cfg,
		Mailbox:         NewMailboxClient(cfg),
		MailboxAttr:     NewMailboxAttrClient(cfg),
		MailboxFlag:     NewMailboxFlagClient(cfg),
		MailboxPermFlag: NewMailboxPermFlagClient(cfg),
		Message:         NewMessageClient(cfg),
		MessageFlag:     NewMessageFlagClient(cfg),
		UID:             NewUIDClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Mailbox.
//		Query().
//		Count(ctx)
//
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
	c.Mailbox.Use(hooks...)
	c.MailboxAttr.Use(hooks...)
	c.MailboxFlag.Use(hooks...)
	c.MailboxPermFlag.Use(hooks...)
	c.Message.Use(hooks...)
	c.MessageFlag.Use(hooks...)
	c.UID.Use(hooks...)
}

// MailboxClient is a client for the Mailbox schema.
type MailboxClient struct {
	config
}

// NewMailboxClient returns a client for the Mailbox from the given config.
func NewMailboxClient(c config) *MailboxClient {
	return &MailboxClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mailbox.Hooks(f(g(h())))`.
func (c *MailboxClient) Use(hooks ...Hook) {
	c.hooks.Mailbox = append(c.hooks.Mailbox, hooks...)
}

// Create returns a create builder for Mailbox.
func (c *MailboxClient) Create() *MailboxCreate {
	mutation := newMailboxMutation(c.config, OpCreate)
	return &MailboxCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Mailbox entities.
func (c *MailboxClient) CreateBulk(builders ...*MailboxCreate) *MailboxCreateBulk {
	return &MailboxCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Mailbox.
func (c *MailboxClient) Update() *MailboxUpdate {
	mutation := newMailboxMutation(c.config, OpUpdate)
	return &MailboxUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MailboxClient) UpdateOne(m *Mailbox) *MailboxUpdateOne {
	mutation := newMailboxMutation(c.config, OpUpdateOne, withMailbox(m))
	return &MailboxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MailboxClient) UpdateOneID(id int) *MailboxUpdateOne {
	mutation := newMailboxMutation(c.config, OpUpdateOne, withMailboxID(id))
	return &MailboxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Mailbox.
func (c *MailboxClient) Delete() *MailboxDelete {
	mutation := newMailboxMutation(c.config, OpDelete)
	return &MailboxDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MailboxClient) DeleteOne(m *Mailbox) *MailboxDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MailboxClient) DeleteOneID(id int) *MailboxDeleteOne {
	builder := c.Delete().Where(mailbox.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MailboxDeleteOne{builder}
}

// Query returns a query builder for Mailbox.
func (c *MailboxClient) Query() *MailboxQuery {
	return &MailboxQuery{
		config: c.config,
	}
}

// Get returns a Mailbox entity by its id.
func (c *MailboxClient) Get(ctx context.Context, id int) (*Mailbox, error) {
	return c.Query().Where(mailbox.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MailboxClient) GetX(ctx context.Context, id int) *Mailbox {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUIDs queries the UIDs edge of a Mailbox.
func (c *MailboxClient) QueryUIDs(m *Mailbox) *UIDQuery {
	query := &UIDQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mailbox.Table, mailbox.FieldID, id),
			sqlgraph.To(uid.Table, uid.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, mailbox.UIDsTable, mailbox.UIDsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFlags queries the flags edge of a Mailbox.
func (c *MailboxClient) QueryFlags(m *Mailbox) *MailboxFlagQuery {
	query := &MailboxFlagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mailbox.Table, mailbox.FieldID, id),
			sqlgraph.To(mailboxflag.Table, mailboxflag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, mailbox.FlagsTable, mailbox.FlagsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPermanentFlags queries the permanent_flags edge of a Mailbox.
func (c *MailboxClient) QueryPermanentFlags(m *Mailbox) *MailboxPermFlagQuery {
	query := &MailboxPermFlagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mailbox.Table, mailbox.FieldID, id),
			sqlgraph.To(mailboxpermflag.Table, mailboxpermflag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, mailbox.PermanentFlagsTable, mailbox.PermanentFlagsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAttributes queries the attributes edge of a Mailbox.
func (c *MailboxClient) QueryAttributes(m *Mailbox) *MailboxAttrQuery {
	query := &MailboxAttrQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mailbox.Table, mailbox.FieldID, id),
			sqlgraph.To(mailboxattr.Table, mailboxattr.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, mailbox.AttributesTable, mailbox.AttributesColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MailboxClient) Hooks() []Hook {
	return c.hooks.Mailbox
}

// MailboxAttrClient is a client for the MailboxAttr schema.
type MailboxAttrClient struct {
	config
}

// NewMailboxAttrClient returns a client for the MailboxAttr from the given config.
func NewMailboxAttrClient(c config) *MailboxAttrClient {
	return &MailboxAttrClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mailboxattr.Hooks(f(g(h())))`.
func (c *MailboxAttrClient) Use(hooks ...Hook) {
	c.hooks.MailboxAttr = append(c.hooks.MailboxAttr, hooks...)
}

// Create returns a create builder for MailboxAttr.
func (c *MailboxAttrClient) Create() *MailboxAttrCreate {
	mutation := newMailboxAttrMutation(c.config, OpCreate)
	return &MailboxAttrCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MailboxAttr entities.
func (c *MailboxAttrClient) CreateBulk(builders ...*MailboxAttrCreate) *MailboxAttrCreateBulk {
	return &MailboxAttrCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MailboxAttr.
func (c *MailboxAttrClient) Update() *MailboxAttrUpdate {
	mutation := newMailboxAttrMutation(c.config, OpUpdate)
	return &MailboxAttrUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MailboxAttrClient) UpdateOne(ma *MailboxAttr) *MailboxAttrUpdateOne {
	mutation := newMailboxAttrMutation(c.config, OpUpdateOne, withMailboxAttr(ma))
	return &MailboxAttrUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MailboxAttrClient) UpdateOneID(id int) *MailboxAttrUpdateOne {
	mutation := newMailboxAttrMutation(c.config, OpUpdateOne, withMailboxAttrID(id))
	return &MailboxAttrUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MailboxAttr.
func (c *MailboxAttrClient) Delete() *MailboxAttrDelete {
	mutation := newMailboxAttrMutation(c.config, OpDelete)
	return &MailboxAttrDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MailboxAttrClient) DeleteOne(ma *MailboxAttr) *MailboxAttrDeleteOne {
	return c.DeleteOneID(ma.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MailboxAttrClient) DeleteOneID(id int) *MailboxAttrDeleteOne {
	builder := c.Delete().Where(mailboxattr.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MailboxAttrDeleteOne{builder}
}

// Query returns a query builder for MailboxAttr.
func (c *MailboxAttrClient) Query() *MailboxAttrQuery {
	return &MailboxAttrQuery{
		config: c.config,
	}
}

// Get returns a MailboxAttr entity by its id.
func (c *MailboxAttrClient) Get(ctx context.Context, id int) (*MailboxAttr, error) {
	return c.Query().Where(mailboxattr.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MailboxAttrClient) GetX(ctx context.Context, id int) *MailboxAttr {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MailboxAttrClient) Hooks() []Hook {
	return c.hooks.MailboxAttr
}

// MailboxFlagClient is a client for the MailboxFlag schema.
type MailboxFlagClient struct {
	config
}

// NewMailboxFlagClient returns a client for the MailboxFlag from the given config.
func NewMailboxFlagClient(c config) *MailboxFlagClient {
	return &MailboxFlagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mailboxflag.Hooks(f(g(h())))`.
func (c *MailboxFlagClient) Use(hooks ...Hook) {
	c.hooks.MailboxFlag = append(c.hooks.MailboxFlag, hooks...)
}

// Create returns a create builder for MailboxFlag.
func (c *MailboxFlagClient) Create() *MailboxFlagCreate {
	mutation := newMailboxFlagMutation(c.config, OpCreate)
	return &MailboxFlagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MailboxFlag entities.
func (c *MailboxFlagClient) CreateBulk(builders ...*MailboxFlagCreate) *MailboxFlagCreateBulk {
	return &MailboxFlagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MailboxFlag.
func (c *MailboxFlagClient) Update() *MailboxFlagUpdate {
	mutation := newMailboxFlagMutation(c.config, OpUpdate)
	return &MailboxFlagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MailboxFlagClient) UpdateOne(mf *MailboxFlag) *MailboxFlagUpdateOne {
	mutation := newMailboxFlagMutation(c.config, OpUpdateOne, withMailboxFlag(mf))
	return &MailboxFlagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MailboxFlagClient) UpdateOneID(id int) *MailboxFlagUpdateOne {
	mutation := newMailboxFlagMutation(c.config, OpUpdateOne, withMailboxFlagID(id))
	return &MailboxFlagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MailboxFlag.
func (c *MailboxFlagClient) Delete() *MailboxFlagDelete {
	mutation := newMailboxFlagMutation(c.config, OpDelete)
	return &MailboxFlagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MailboxFlagClient) DeleteOne(mf *MailboxFlag) *MailboxFlagDeleteOne {
	return c.DeleteOneID(mf.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MailboxFlagClient) DeleteOneID(id int) *MailboxFlagDeleteOne {
	builder := c.Delete().Where(mailboxflag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MailboxFlagDeleteOne{builder}
}

// Query returns a query builder for MailboxFlag.
func (c *MailboxFlagClient) Query() *MailboxFlagQuery {
	return &MailboxFlagQuery{
		config: c.config,
	}
}

// Get returns a MailboxFlag entity by its id.
func (c *MailboxFlagClient) Get(ctx context.Context, id int) (*MailboxFlag, error) {
	return c.Query().Where(mailboxflag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MailboxFlagClient) GetX(ctx context.Context, id int) *MailboxFlag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MailboxFlagClient) Hooks() []Hook {
	return c.hooks.MailboxFlag
}

// MailboxPermFlagClient is a client for the MailboxPermFlag schema.
type MailboxPermFlagClient struct {
	config
}

// NewMailboxPermFlagClient returns a client for the MailboxPermFlag from the given config.
func NewMailboxPermFlagClient(c config) *MailboxPermFlagClient {
	return &MailboxPermFlagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mailboxpermflag.Hooks(f(g(h())))`.
func (c *MailboxPermFlagClient) Use(hooks ...Hook) {
	c.hooks.MailboxPermFlag = append(c.hooks.MailboxPermFlag, hooks...)
}

// Create returns a create builder for MailboxPermFlag.
func (c *MailboxPermFlagClient) Create() *MailboxPermFlagCreate {
	mutation := newMailboxPermFlagMutation(c.config, OpCreate)
	return &MailboxPermFlagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MailboxPermFlag entities.
func (c *MailboxPermFlagClient) CreateBulk(builders ...*MailboxPermFlagCreate) *MailboxPermFlagCreateBulk {
	return &MailboxPermFlagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MailboxPermFlag.
func (c *MailboxPermFlagClient) Update() *MailboxPermFlagUpdate {
	mutation := newMailboxPermFlagMutation(c.config, OpUpdate)
	return &MailboxPermFlagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MailboxPermFlagClient) UpdateOne(mpf *MailboxPermFlag) *MailboxPermFlagUpdateOne {
	mutation := newMailboxPermFlagMutation(c.config, OpUpdateOne, withMailboxPermFlag(mpf))
	return &MailboxPermFlagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MailboxPermFlagClient) UpdateOneID(id int) *MailboxPermFlagUpdateOne {
	mutation := newMailboxPermFlagMutation(c.config, OpUpdateOne, withMailboxPermFlagID(id))
	return &MailboxPermFlagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MailboxPermFlag.
func (c *MailboxPermFlagClient) Delete() *MailboxPermFlagDelete {
	mutation := newMailboxPermFlagMutation(c.config, OpDelete)
	return &MailboxPermFlagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MailboxPermFlagClient) DeleteOne(mpf *MailboxPermFlag) *MailboxPermFlagDeleteOne {
	return c.DeleteOneID(mpf.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MailboxPermFlagClient) DeleteOneID(id int) *MailboxPermFlagDeleteOne {
	builder := c.Delete().Where(mailboxpermflag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MailboxPermFlagDeleteOne{builder}
}

// Query returns a query builder for MailboxPermFlag.
func (c *MailboxPermFlagClient) Query() *MailboxPermFlagQuery {
	return &MailboxPermFlagQuery{
		config: c.config,
	}
}

// Get returns a MailboxPermFlag entity by its id.
func (c *MailboxPermFlagClient) Get(ctx context.Context, id int) (*MailboxPermFlag, error) {
	return c.Query().Where(mailboxpermflag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MailboxPermFlagClient) GetX(ctx context.Context, id int) *MailboxPermFlag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MailboxPermFlagClient) Hooks() []Hook {
	return c.hooks.MailboxPermFlag
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Create returns a create builder for Message.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id int) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MessageClient) DeleteOneID(id int) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id int) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id int) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFlags queries the flags edge of a Message.
func (c *MessageClient) QueryFlags(m *Message) *MessageFlagQuery {
	query := &MessageFlagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(messageflag.Table, messageflag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, message.FlagsTable, message.FlagsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUIDs queries the UIDs edge of a Message.
func (c *MessageClient) QueryUIDs(m *Message) *UIDQuery {
	query := &UIDQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(uid.Table, uid.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, message.UIDsTable, message.UIDsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	return c.hooks.Message
}

// MessageFlagClient is a client for the MessageFlag schema.
type MessageFlagClient struct {
	config
}

// NewMessageFlagClient returns a client for the MessageFlag from the given config.
func NewMessageFlagClient(c config) *MessageFlagClient {
	return &MessageFlagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `messageflag.Hooks(f(g(h())))`.
func (c *MessageFlagClient) Use(hooks ...Hook) {
	c.hooks.MessageFlag = append(c.hooks.MessageFlag, hooks...)
}

// Create returns a create builder for MessageFlag.
func (c *MessageFlagClient) Create() *MessageFlagCreate {
	mutation := newMessageFlagMutation(c.config, OpCreate)
	return &MessageFlagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MessageFlag entities.
func (c *MessageFlagClient) CreateBulk(builders ...*MessageFlagCreate) *MessageFlagCreateBulk {
	return &MessageFlagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MessageFlag.
func (c *MessageFlagClient) Update() *MessageFlagUpdate {
	mutation := newMessageFlagMutation(c.config, OpUpdate)
	return &MessageFlagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageFlagClient) UpdateOne(mf *MessageFlag) *MessageFlagUpdateOne {
	mutation := newMessageFlagMutation(c.config, OpUpdateOne, withMessageFlag(mf))
	return &MessageFlagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageFlagClient) UpdateOneID(id int) *MessageFlagUpdateOne {
	mutation := newMessageFlagMutation(c.config, OpUpdateOne, withMessageFlagID(id))
	return &MessageFlagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MessageFlag.
func (c *MessageFlagClient) Delete() *MessageFlagDelete {
	mutation := newMessageFlagMutation(c.config, OpDelete)
	return &MessageFlagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MessageFlagClient) DeleteOne(mf *MessageFlag) *MessageFlagDeleteOne {
	return c.DeleteOneID(mf.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MessageFlagClient) DeleteOneID(id int) *MessageFlagDeleteOne {
	builder := c.Delete().Where(messageflag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageFlagDeleteOne{builder}
}

// Query returns a query builder for MessageFlag.
func (c *MessageFlagClient) Query() *MessageFlagQuery {
	return &MessageFlagQuery{
		config: c.config,
	}
}

// Get returns a MessageFlag entity by its id.
func (c *MessageFlagClient) Get(ctx context.Context, id int) (*MessageFlag, error) {
	return c.Query().Where(messageflag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageFlagClient) GetX(ctx context.Context, id int) *MessageFlag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MessageFlagClient) Hooks() []Hook {
	return c.hooks.MessageFlag
}

// UIDClient is a client for the UID schema.
type UIDClient struct {
	config
}

// NewUIDClient returns a client for the UID from the given config.
func NewUIDClient(c config) *UIDClient {
	return &UIDClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `uid.Hooks(f(g(h())))`.
func (c *UIDClient) Use(hooks ...Hook) {
	c.hooks.UID = append(c.hooks.UID, hooks...)
}

// Create returns a create builder for UID.
func (c *UIDClient) Create() *UIDCreate {
	mutation := newUIDMutation(c.config, OpCreate)
	return &UIDCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UID entities.
func (c *UIDClient) CreateBulk(builders ...*UIDCreate) *UIDCreateBulk {
	return &UIDCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UID.
func (c *UIDClient) Update() *UIDUpdate {
	mutation := newUIDMutation(c.config, OpUpdate)
	return &UIDUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UIDClient) UpdateOne(u *UID) *UIDUpdateOne {
	mutation := newUIDMutation(c.config, OpUpdateOne, withUID(u))
	return &UIDUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UIDClient) UpdateOneID(id int) *UIDUpdateOne {
	mutation := newUIDMutation(c.config, OpUpdateOne, withUIDID(id))
	return &UIDUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UID.
func (c *UIDClient) Delete() *UIDDelete {
	mutation := newUIDMutation(c.config, OpDelete)
	return &UIDDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UIDClient) DeleteOne(u *UID) *UIDDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UIDClient) DeleteOneID(id int) *UIDDeleteOne {
	builder := c.Delete().Where(uid.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UIDDeleteOne{builder}
}

// Query returns a query builder for UID.
func (c *UIDClient) Query() *UIDQuery {
	return &UIDQuery{
		config: c.config,
	}
}

// Get returns a UID entity by its id.
func (c *UIDClient) Get(ctx context.Context, id int) (*UID, error) {
	return c.Query().Where(uid.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UIDClient) GetX(ctx context.Context, id int) *UID {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMessage queries the message edge of a UID.
func (c *UIDClient) QueryMessage(u *UID) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(uid.Table, uid.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, uid.MessageTable, uid.MessageColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMailbox queries the mailbox edge of a UID.
func (c *UIDClient) QueryMailbox(u *UID) *MailboxQuery {
	query := &MailboxQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(uid.Table, uid.FieldID, id),
			sqlgraph.To(mailbox.Table, mailbox.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, uid.MailboxTable, uid.MailboxColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UIDClient) Hooks() []Hook {
	return c.hooks.UID
}
