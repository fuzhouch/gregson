package cli

import "github.com/spf13/cobra"

type CLI struct {
	name       string
	usageShort string
	usageLong  string
	cmd        *cobra.Command
}

func New(name string) *CLI {
	s := &CLI{
		name: name,
	}
	return s
}

func (c *CLI) UsageShort(v string) *CLI {
	c.usageShort = v
	return c
}

func (c *CLI) UsageLong(v string) *CLI {
	c.usageLong = v
	return c
}

func (c *CLI) BuildCobra() *cobra.Command {
	c.cmd = newRootCmd(c.name, c.usageShort, c.usageLong)
	return c.cmd
}
