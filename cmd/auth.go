package cmd

import (
	"os"
	"text/tabwriter"
)

type AuthCommand Command

func NewAuthCmd() *AuthCommand {
	return &AuthCommand{
		Name:        "auth",
		Description: "Store api_key on a config file. The default name of the config file is .devto",
	}
}

//Run execute the command
func (c *AuthCommand) Run() CommandValidationError {
	//big assumption for now, that .devto is on the same
	//folder from where the cli it is executed
	file, err := os.Create(".devto")
	if err != nil {
		return NewCommandError("Unable to create file")
	}
	defer file.Close()

	if c.Data == "" {
		return ApiKeyMissing
	}
	_, err = file.Write([]byte(c.Data))
	if err != nil {
		return NewCommandError("Unable to store api_key")
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *AuthCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *AuthCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *AuthCommand) SetData(data string) {
	c.Data = data
}
