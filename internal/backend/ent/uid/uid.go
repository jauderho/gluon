// Code generated by entc, DO NOT EDIT.

package uid

const (
	// Label holds the string label denoting the uid type in the database.
	Label = "uid"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// FieldRecent holds the string denoting the recent field in the database.
	FieldRecent = "recent"
	// FieldInDeletionPool holds the string denoting the indeletionpool field in the database.
	FieldInDeletionPool = "in_deletion_pool"
	// EdgeMessage holds the string denoting the message edge name in mutations.
	EdgeMessage = "message"
	// EdgeMailbox holds the string denoting the mailbox edge name in mutations.
	EdgeMailbox = "mailbox"
	// Table holds the table name of the uid in the database.
	Table = "ui_ds"
	// MessageTable is the table that holds the message relation/edge.
	MessageTable = "ui_ds"
	// MessageInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessageInverseTable = "messages"
	// MessageColumn is the table column denoting the message relation/edge.
	MessageColumn = "uid_message"
	// MailboxTable is the table that holds the mailbox relation/edge.
	MailboxTable = "ui_ds"
	// MailboxInverseTable is the table name for the Mailbox entity.
	// It exists in this package in order to avoid circular dependency with the "mailbox" package.
	MailboxInverseTable = "mailboxes"
	// MailboxColumn is the table column denoting the mailbox relation/edge.
	MailboxColumn = "mailbox_ui_ds"
)

// Columns holds all SQL columns for uid fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldDeleted,
	FieldRecent,
	FieldInDeletionPool,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ui_ds"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"mailbox_ui_ds",
	"uid_message",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDeleted holds the default value on creation for the "Deleted" field.
	DefaultDeleted bool
	// DefaultRecent holds the default value on creation for the "Recent" field.
	DefaultRecent bool
	// DefaultInDeletionPool holds the default value on creation for the "InDeletionPool" field.
	DefaultInDeletionPool bool
)
