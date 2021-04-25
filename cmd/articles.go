package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
)

type ArticlesCommand Command

func NewArticlesCmd() *ArticlesCommand {
	return &ArticlesCommand{
		Name:        "articles",
		Description: "Retrieve, Create and Update the articles",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrive articles",
				Active:      false,
			},
			"retrieve_latest": {
				Description: "Retrieve latest articles",
				Active:      false,
			},
			"retrieve_videos": {
				Description: "Retrive articles with videos",
				Active:      false,
			},
			"retrieve_id": {
				Description: "Retrieve article by ID",
				Active:      false,
			},
			"latest_query": {
				Description: "Unable queries on retrieve latest articles",
				Active:      false,
			},
			"retrieve_query": {
				Description: "Retrive articles with specific queries",
				Active:      false,
			},
			"create": {
				Description: "Create article",
				Active:      false,
			},
			"update": {
				Description: "Update and article",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *ArticlesCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve"].Active {
		var queries *api.GetArticleQuery
		if c.Subcommands["retrieve_query"].Active {
			queries, err = processQueries()
			if err != nil {
				return err
			}
		} else if c.Subcommands["retrieve_id"].Active {
			err := c.retrieveByID()
			if err != nil {
				return err
			}
		}
		err = c.retrieve(queries)
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_latest"].Active {
		var queries *api.GetLatestArticleQuery
		if c.Subcommands["latest_query"].Active {
			queries, err = processLatestQueries()
			if err != nil {
				return err
			}
		}
		err = c.retrieveLatest(queries)
		if err != nil {
			return err
		}
	} else if c.Subcommands["update"].Active {
		article, err := processUpdate()
		if err != nil {
			return err
		}
		err = c.update(article)
		if err != nil {
			return err
		}
	} else if c.Subcommands["create"].Active {
		article, err := processCreate()
		if err != nil {
			return err
		}
		err = c.create(article)
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_videos"].Active {
		err = c.retrieveArticlesVideo()
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *ArticlesCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *ArticlesCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *ArticlesCommand) SetData(data string) {
	c.Data = data
}

//retrieve ...
func (c *ArticlesCommand) retrieve(queries *api.GetArticleQuery) CommandValidationError {
	_, err := api.RetrieveArticles(c.Data, queries)
	if err != nil {
		return err
	}
	return nil
}

//retrieveByID ...
func (c *ArticlesCommand) retrieveByID() CommandValidationError {
	_, err := api.RetrieveArticleByID(c.Data)
	if err != nil {
		return err
	}
	return nil
}

//retrieveArticlesVideo ...
func (c *ArticlesCommand) retrieveArticlesVideo() CommandValidationError {
	_, err := api.RetrieveArticlesVideo(c.Data)
	if err != nil {
		return err
	}
	return nil
}

//retrieveLatest ...
func (c *ArticlesCommand) retrieveLatest(queries *api.GetLatestArticleQuery) CommandValidationError {
	_, err := api.RetrieveLatestArticles(queries)
	if err != nil {
		return err
	}
	return nil
}

//update ...
func (c *ArticlesCommand) update(data *api.ArticleEdit) CommandValidationError {
	_, err := api.UpdateArticle(c.Data, data)
	if err != nil {
		return err
	}
	return nil
}

//processLatestQueries read the data from the User input and put
//that data inside an GetArticleQuery structure
func processLatestQueries() (*api.GetLatestArticleQuery, error) {
	//to store field from GetLatestArticleQuery
	queries := new(api.GetLatestArticleQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//processQueries read the data from the User input and put
//that data inside an GetArticleQuery structure
func processQueries() (*api.GetArticleQuery, error) {
	//to store field from GetArticleQuery
	queries := new(api.GetArticleQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//processUpdate read the data from the User input and put
//that data inside an ArticleEdit structure
func processUpdate() (*api.ArticleEdit, error) {
	//to store field from ArticleEdit
	article := new(api.ArticleEditType)
	err := processInput(article)
	if err != nil {
		return nil, err
	}
	return &api.ArticleEdit{
		Article: article,
	}, nil
}

//processCreate read the data from the User input and put
//that data inside an ArticleCreate structure
func processCreate() (*api.ArticleCreate, error) {
	//to store field from ArticleEdit
	article := new(api.ArticleCreateType)
	err := processInput(article)
	if err != nil {
		return nil, err
	}
	return &api.ArticleCreate{
		Article: article,
	}, nil
}

//create ...
func (c *ArticlesCommand) create(data *api.ArticleCreate) CommandValidationError {
	_, err := api.CreateArticle(data)
	if err != nil {
		return err
	}
	return nil
}

//ActivateSubcommand ...
func (c *ArticlesCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
