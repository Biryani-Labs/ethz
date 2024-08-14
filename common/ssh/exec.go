package ssh

import (
	"fmt"

	"github.com/Biryani-Labs/ethz/pkg/schema"
	"golang.org/x/crypto/ssh"
)

func ExecSSHCommandInOrder(client *ssh.Client, commands []string) []schema.CommandResult {
	results := make([]schema.CommandResult, 0, len(commands))

	for _, cmd := range commands {
		session, err := client.NewSession()
		if err != nil {
			results = append(results, schema.CommandResult{Command: cmd, Error: fmt.Errorf("failed to create session: %v", err)})
			return results
		}
		defer session.Close()

		output, err := session.CombinedOutput(cmd)
		result := schema.CommandResult{
			Command: cmd,
			Output:  string(output),
			Error:   err,
		}
		results = append(results, result)

		if err != nil {
			return results
		}
	}

	return results
}
