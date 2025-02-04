// Code generated by ent, DO NOT EDIT.

package mailbox

import (
	"github.com/ProtonMail/gluon/imap"
)

const (
	// Label holds the string label denoting the mailbox type in the database.
	Label = "mailbox"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRemoteID holds the string denoting the remoteid field in the database.
	FieldRemoteID = "remote_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUIDNext holds the string denoting the uidnext field in the database.
	FieldUIDNext = "uid_next"
	// FieldUIDValidity holds the string denoting the uidvalidity field in the database.
	FieldUIDValidity = "uid_validity"
	// FieldSubscribed holds the string denoting the subscribed field in the database.
	FieldSubscribed = "subscribed"
	// EdgeUIDs holds the string denoting the uids edge name in mutations.
	EdgeUIDs = "UIDs"
	// EdgeFlags holds the string denoting the flags edge name in mutations.
	EdgeFlags = "flags"
	// EdgePermanentFlags holds the string denoting the permanent_flags edge name in mutations.
	EdgePermanentFlags = "permanent_flags"
	// EdgeAttributes holds the string denoting the attributes edge name in mutations.
	EdgeAttributes = "attributes"
	// Table holds the table name of the mailbox in the database.
	Table = "mailboxes"
	// UIDsTable is the table that holds the UIDs relation/edge.
	UIDsTable = "ui_ds"
	// UIDsInverseTable is the table name for the UID entity.
	// It exists in this package in order to avoid circular dependency with the "uid" package.
	UIDsInverseTable = "ui_ds"
	// UIDsColumn is the table column denoting the UIDs relation/edge.
	UIDsColumn = "mailbox_ui_ds"
	// FlagsTable is the table that holds the flags relation/edge.
	FlagsTable = "mailbox_flags"
	// FlagsInverseTable is the table name for the MailboxFlag entity.
	// It exists in this package in order to avoid circular dependency with the "mailboxflag" package.
	FlagsInverseTable = "mailbox_flags"
	// FlagsColumn is the table column denoting the flags relation/edge.
	FlagsColumn = "mailbox_flags"
	// PermanentFlagsTable is the table that holds the permanent_flags relation/edge.
	PermanentFlagsTable = "mailbox_perm_flags"
	// PermanentFlagsInverseTable is the table name for the MailboxPermFlag entity.
	// It exists in this package in order to avoid circular dependency with the "mailboxpermflag" package.
	PermanentFlagsInverseTable = "mailbox_perm_flags"
	// PermanentFlagsColumn is the table column denoting the permanent_flags relation/edge.
	PermanentFlagsColumn = "mailbox_permanent_flags"
	// AttributesTable is the table that holds the attributes relation/edge.
	AttributesTable = "mailbox_attrs"
	// AttributesInverseTable is the table name for the MailboxAttr entity.
	// It exists in this package in order to avoid circular dependency with the "mailboxattr" package.
	AttributesInverseTable = "mailbox_attrs"
	// AttributesColumn is the table column denoting the attributes relation/edge.
	AttributesColumn = "mailbox_attributes"
)

// Columns holds all SQL columns for mailbox fields.
var Columns = []string{
	FieldID,
	FieldRemoteID,
	FieldName,
	FieldUIDNext,
	FieldUIDValidity,
	FieldSubscribed,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUIDNext holds the default value on creation for the "UIDNext" field.
	DefaultUIDNext imap.UID
	// DefaultUIDValidity holds the default value on creation for the "UIDValidity" field.
	DefaultUIDValidity imap.UID
	// DefaultSubscribed holds the default value on creation for the "Subscribed" field.
	DefaultSubscribed bool
)
