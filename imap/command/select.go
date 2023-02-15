package command

import (
	"fmt"
	rfcparser "github.com/ProtonMail/gluon/rfcparser"
)

type SelectCommand struct {
	Mailbox string
}

func (l SelectCommand) String() string {
	return fmt.Sprintf("SELECT '%v'", l.Mailbox)
}

func (l SelectCommand) SanitizedString() string {
	return fmt.Sprintf("SELECT '%v'", sanitizeString(l.Mailbox))
}

type SelectCommandParser struct{}

func (SelectCommandParser) FromParser(p *rfcparser.Parser) (Payload, error) {
	// select          = "SELECT" SP mailbox
	if err := p.Consume(rfcparser.TokenTypeSP, "expected space after command"); err != nil {
		return nil, err
	}

	mailbox, err := ParseMailbox(p)
	if err != nil {
		return nil, err
	}

	return &SelectCommand{
		Mailbox: mailbox.Value,
	}, nil
}
