package cli

import (
	"bytes"
	"testing"
)

func TestCreateDefaultCLI(t *testing.T) {
	// it should be app := cli.New() in client code.
	args := []string{"--help"}
	buf := new(bytes.Buffer)
	cmd := New("testcli").BuildCobra()

	cmd.SetOut(buf)
	cmd.SetArgs(args)

	err := cmd.Execute()
	if err != nil {
		t.Errorf("RunHelpFail:%s", err.Error())
		return
	}
}
