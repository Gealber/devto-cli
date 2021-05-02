package cmd

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
	"github.com/Gealber/devto-cli/display"
)

type WebhooksCommand Command

func NewWebhooksCmd() *WebhooksCommand {
	return &WebhooksCommand{
		Name:        "webhooks",
		Description: "Retrieve, Create and Delete the webhooks",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrieve webhooks",
				Active:      false,
			},
			"retrieve_id": {
				Description: "Retrieve webhook by ID",
				Active:      false,
			},
			"create": {
				Description: "Create a webhook",
				Active:      false,
			},
			"delete": {
				Description: "Delete a webhook",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *WebhooksCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve"].Active {
		err = c.retrieveWebHooks()
		if err != nil {
			return err
		}
	} else if c.Subcommands["create"].Active {
		webhook, err := processWebhookCreate()
		if err != nil {
			return err
		}
		err = c.createWebhook(webhook)
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_id"].Active {
		err = c.retrieveWebhookByID()
		if err != nil {
			return err
		}
	} else if c.Subcommands["delete"].Active {
		err = c.deleteWebhooks()
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *WebhooksCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *WebhooksCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *WebhooksCommand) SetData(data string) {
	c.Data = data
}

//retrieveWebhooks ...
func (c *WebhooksCommand) retrieveWebHooks() CommandValidationError {
	ctx := context.Background()
	wbhooks, err := api.RetrieveWebhooks(ctx)
	if err != nil {
		return err
	}
	display.WebhooksResponse(wbhooks)
	return nil
}

//retrieveWebhookByID ...
func (c *WebhooksCommand) retrieveWebhookByID() CommandValidationError {
	ctx := context.Background()
	webhook, err := api.RetrieveWebhookByID(ctx, c.Data)
	if err != nil {
		return err
	}
	display.WebhooksTypeBasic(webhook)
	return nil
}

//deleteWebhooks ...
func (c *WebhooksCommand) deleteWebhooks() CommandValidationError {
	ctx := context.Background()
	webhook, err := api.DeleteWebhook(ctx, c.Data)
	if err != nil {
		return err
	}
	display.WebhooksTypeBasic(webhook)
	return nil
}

//processWebhookCreate read the data from the User input and put
//that data inside an WebhooksCreateType structure
func processWebhookCreate() (*api.WebhooksCreateType, error) {
	//to store field from WebhooksCreateType
	webhook := new(api.WebhookDataCreate)
	err := processInput(webhook)
	if err != nil {
		return nil, err
	}
	return &api.WebhooksCreateType{
		WebhookEndpoint: webhook,
	}, nil
}

//createWebhook ...
func (c *WebhooksCommand) createWebhook(data *api.WebhooksCreateType) CommandValidationError {
	ctx := context.Background()
	webhook, err := api.CreateWebhook(ctx, data)
	if err != nil {
		return err
	}
	display.WebhooksCreated(webhook)
	return nil
}

//ActivateSubcommand ...
func (c *WebhooksCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
