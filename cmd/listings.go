package cmd

import (
	"context"
	"errors"
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
	"github.com/Gealber/devto-cli/display"
)

type ListingsCommand Command

func NewListingsCommand() *ListingsCommand {
	return &ListingsCommand{
		Name:        "listings",
		Description: "Retrieve and Create listings",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrieve listings",
				Active:      false,
			},
			"retrieve_id": {
				Description: "Retrieve a listing by id",
				Active:      false,
			},
			"update": {
				Description: "Update listings",
				Active:      false,
			},
			"create": {
				Description: "Create a listing",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *ListingsCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve"].Active {
		queries, err := processListingsQueries()
		if err != nil {
			return err
		}
		err = c.retrieveListing(queries)
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_id"].Active {
		err = c.retrieveListingByID()
		if err != nil {
			return err
		}
	} else if c.Subcommands["create"].Active {
		listing, err := processListingCreate()
		if err != nil {
			return err
		}
		err = c.createListing(listing)
		if err != nil {
			return err
		}
	} else if c.Subcommands["update"].Active {
		listing, err := processListingUpdate()
		if err != nil {
			return err
		}
		err = c.updateListing(listing)
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *ListingsCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *ListingsCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *ListingsCommand) SetData(data string) {
	c.Data = data
}

//processListingsQueries read the data from the User input and put
//that data inside an ListingQuery structure
func processListingsQueries() (*api.ListingQuery, error) {
	//to store field from ListingQuery
	queries := new(api.ListingQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//retrieveListing ...
func (c *ListingsCommand) retrieveListing(queries *api.ListingQuery) CommandValidationError {
	ctx := context.Background()
	listings, err := api.RetrieveListings(ctx, queries)
	if err != nil {
		return err
	}
	display.ListingResponse(listings)
	return nil
}

//retrieveListingByID ...
func (c *ListingsCommand) retrieveListingByID() CommandValidationError {
	ctx := context.Background()
	listings, err := api.RetrieveListingsByID(ctx, c.Data)
	if err != nil {
		return err
	}
	display.ListingResponse(listings)
	return nil
}

//processListingCreate read the data from the User input and put
//that data inside an ListingCreateType structure
func processListingCreate() (*api.ListingCreate, error) {
	//to store field from ArticleEdit
	listing := new(api.ListingCreateType)
	err := processInput(listing)
	if err != nil {
		return nil, err
	}
	return &api.ListingCreate{
		Listing: listing,
	}, nil
}

//createListing ...
func (c *ListingsCommand) createListing(data *api.ListingCreate) CommandValidationError {
	//validate category
	allowed := []string{
		"cfp",
		"forhire",
		"collabs",
		"education",
		"jobs",
		"mentors",
		"products",
		"mentees",
		"forsale",
		"events",
		"misc",
	}
	valid := false
	for _, cat := range allowed {
		if data.Listing.Category == cat {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("Category is not allowed")
	}
	ctx := context.Background()
	listing, err := api.CreateListing(ctx, data)
	if err != nil {
		return err
	}
	display.CreatedListing(listing)
	return nil
}

//processListingUpdate read the data from the User input and put
//that data inside an ListingCreateType structure
func processListingUpdate() (*api.ListingUpdate, error) {
	//to store field from ListingUpdateType
	listing := new(api.ListingUpdateType)
	err := processInput(listing)
	if err != nil {
		return nil, err
	}
	return &api.ListingUpdate{
		Listing: listing,
	}, nil
}

//updateListing ...
func (c *ListingsCommand) updateListing(data *api.ListingUpdate) CommandValidationError {
	//validate category
	allowed := []string{
		"cfp",
		"forhire",
		"collabs",
		"education",
		"jobs",
		"mentors",
		"products",
		"mentees",
		"forsale",
		"events",
		"misc",
	}
	valid := false
	for _, cat := range allowed {
		if data.Listing.Category == cat {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("Category is not allowed")
	}
	ctx := context.Background()
	listing, err := api.UpdateListing(ctx, c.Data, data)
	if err != nil {
		return err
	}
	display.ModifiedArticle(listing)
	return nil
}

//ActivateSubcommand ...
func (c *ListingsCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
