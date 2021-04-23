package cmd

import "text/tabwriter"

type AdminConfig Command

func NewAdminConfigCmd() *AdminConfig {
	return &AdminConfig{
		Name:        "admin-config",
		Description: "Retrieve or update admin configuration from dev.to",
		Subcommands: []*Command{
			{
				Name:        "retrieve",
				Description: "Retrive the admin configuration from dev.to",
			},
			{
				Name:        "update",
				Description: "Update the admin configuration on dev.to",
			},
		},
	}
}

//Run execute the command
func (c *AdminConfig) Run() (*CommandResponse, CommandValidationError) {
	//Diferentiate two cases when is a retrieve and when is an update
	return nil, nil
}

//Validate execute the command
func (c *AdminConfig) Validate() CommandValidationError {
	return nil
}

//Helper display info about the command
func (c *AdminConfig) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//Retrieve ...
func (c *AdminConfig) Retrieve() (*CommandResponse, CommandValidationError) {
	return nil, nil
}

//SetData ...
func (c *AdminConfig) SetData(data string) {
	c.Data = data
}
