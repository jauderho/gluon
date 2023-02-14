package command

import (
	"bytes"
	"github.com/ProtonMail/gluon/imap/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser_LSubCommand(t *testing.T) {
	input := toIMAPLine(`tag LSUB "" "*"`)
	s := parser.NewScanner(bytes.NewReader(input))
	p := NewParser(s)

	expected := Command{Tag: "tag", Payload: &LSubCommand{
		Mailbox:     "",
		LSubMailbox: "*",
	}}

	cmd, err := p.Parse()
	require.NoError(t, err)
	require.Equal(t, expected, cmd)
	require.Equal(t, "lsub", p.LastParsedCommand())
	require.Equal(t, "tag", p.LastParsedTag())
}

func TestParser_LSubCommandSpecialAsterisk(t *testing.T) {
	input := toIMAPLine(`tag LSUB "foo" *`)
	s := parser.NewScanner(bytes.NewReader(input))
	p := NewParser(s)

	expected := Command{Tag: "tag", Payload: &LSubCommand{
		Mailbox:     "foo",
		LSubMailbox: "*",
	}}

	cmd, err := p.Parse()
	require.NoError(t, err)
	require.Equal(t, expected, cmd)
	require.Equal(t, "lsub", p.LastParsedCommand())
	require.Equal(t, "tag", p.LastParsedTag())
}

func TestParser_LSubCommandSpecialPercentage(t *testing.T) {
	input := toIMAPLine(`tag LSUB "bar" %`)
	s := parser.NewScanner(bytes.NewReader(input))
	p := NewParser(s)

	expected := Command{Tag: "tag", Payload: &LSubCommand{
		Mailbox:     "bar",
		LSubMailbox: "%",
	}}

	cmd, err := p.Parse()
	require.NoError(t, err)
	require.Equal(t, expected, cmd)
	require.Equal(t, "lsub", p.LastParsedCommand())
	require.Equal(t, "tag", p.LastParsedTag())
}