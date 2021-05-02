package cmd

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
	"github.com/Gealber/devto-cli/display"
)

type OrganizationsCommand Command

func NewOrganizationsCommand() *OrganizationsCommand {
	return &OrganizationsCommand{
		Name:        "organizations",
		Description: "Retrieve  organizations",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrieve organizations",
				Active:      false,
			},
			"retrieve_users": {
				Description: "Retrieve users on an organization",
				Active:      false,
			},
			"retrieve_listing": {
				Description: "Retrieve listings on an organization",
				Active:      false,
			},
			"retrieve_articles": {
				Description: "Retrieve articles belonging to an organization",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *OrganizationsCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve"].Active {
		err = c.retrieveOrganization()
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_users"].Active {
		queries, err := processOrganizationsQueries()
		if err != nil {
			return err
		}
		err = c.retrieveUsersOnOrganization(queries)
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_listing"].Active {
		queries, err := processListingOrganizationsQueries()
		if err != nil {
			return err
		}
		err = c.retrieveListingOnOrganization(queries)
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_articles"].Active {
		queries, err := processOrganizationsQueries()
		if err != nil {
			return err
		}
		err = c.retrieveArticlesOnOrganization(queries)
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *OrganizationsCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *OrganizationsCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *OrganizationsCommand) SetData(data string) {
	c.Data = data
}

//processOrganizationsQueries read the data from the User input and put
//that data inside an CommonQuery structure
func processOrganizationsQueries() (*api.CommonQuery, error) {
	//to store field from CommonQuery
	queries := new(api.CommonQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//processListingOrganizationsQueries read the data from the User input and put
//that data inside an OrganizationQuery structure
func processListingOrganizationsQueries() (*api.OrganizationListingQuery, error) {
	//to store field from OrganizationListingQuery
	queries := new(api.OrganizationListingQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//retrieveOrganization ...
func (c *OrganizationsCommand) retrieveOrganization() CommandValidationError {
	ctx := context.Background()
	organization, err := api.RetrieveOrganizationByUsername(ctx, c.Data)
	if err != nil {
		return err
	}
	display.OrganizationResponse(organization)
	return nil
}

//retrieveUsersOnOrganization ...
func (c *OrganizationsCommand) retrieveUsersOnOrganization(query *api.CommonQuery) CommandValidationError {
	ctx := context.Background()
	users, err := api.RetrieveUsersOnOrganization(ctx, c.Data, query)
	if err != nil {
		return err
	}
	display.UserOnOrganizationResponse(users)
	return nil
}

//retrieveListingOnOrganization ...
func (c *OrganizationsCommand) retrieveListingOnOrganization(query *api.OrganizationListingQuery) CommandValidationError {
	ctx := context.Background()
	listings, err := api.RetrieveListingOnOrganization(ctx, c.Data, query)
	if err != nil {
		return err
	}
	display.ListingResponse(listings)
	return nil
}

//retrieveArticlesOnOrganization ...
func (c *OrganizationsCommand) retrieveArticlesOnOrganization(query *api.CommonQuery) CommandValidationError {
	ctx := context.Background()
	articles, err := api.RetrieveArticlesOnOrganization(ctx, c.Data, query)
	if err != nil {
		return err
	}
	display.RetrievedArticles(articles)
	return nil
}

//ActivateSubcommand ...
func (c *OrganizationsCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
