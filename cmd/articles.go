package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
	"github.com/Gealber/devto-cli/display"
)

type ArticlesCommand Command

func NewArticlesCmd() *ArticlesCommand {
	return &ArticlesCommand{
		Name:        "articles",
		Description: "Retrieve, Create and Update the articles",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrieve articles",
				Active:      false,
			},
			"retrieve_latest": {
				Description: "Retrieve latest articles",
				Active:      false,
			},
			"retrieve_videos": {
				Description: "Retrieve articles with videos",
				Active:      false,
			},
			"retrieve_id": {
				Description: "Retrieve article by ID",
				Active:      false,
			},
			"display_body": {
				Description: "Display the body markdown of the article with ID",
				Active:      false,
			},
			"latest_query": {
				Description: "Unable queries on retrieve latest articles",
				Active:      false,
			},
			"retrieve_query": {
				Description: "Retrieve articles with specific queries",
				Active:      false,
			},
			"retrieve_me": {
				Description: "Retrieve articles of authenticated user",
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
			body := c.Subcommands["display_body"].Active
			err := c.retrieveByID(body)
			if err != nil {
				return err
			}
			return nil
		}
		err = c.retrieve(queries)
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_latest"].Active {
		var queries *api.CommonQuery
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
	} else if c.Subcommands["retrieve_me"].Active {
		var queries *api.CommonQuery
		queries, err = processCommonQueries()
		if err != nil {
			return err
		}
		err = c.retrieveMe(queries)
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
	articles, err := api.RetrieveArticles(c.Data, queries)
	if err != nil {
		return err
	}
	display.RetrievedArticles(articles)
	return nil
}

//retrieveMe ...
func (c *ArticlesCommand) retrieveMe(queries *api.CommonQuery) CommandValidationError {
	articles, err := api.RetrieveMeArticles(queries, c.Data)
	if err != nil {
		return err
	}
	display.RetrievedMyArticles(articles)
	return nil
}

//retrieveByID ...
func (c *ArticlesCommand) retrieveByID(body bool) CommandValidationError {
	article, err := api.RetrieveArticleByID(c.Data)
	if err != nil {
		return err
	}
	if body {
		display.ModifiedArticleBody(article)
	} else {
		display.ModifiedArticle(article)
	}
	return nil
}

//retrieveArticlesVideo ...
func (c *ArticlesCommand) retrieveArticlesVideo() CommandValidationError {
	articles, err := api.RetrieveArticlesVideo(c.Data)
	if err != nil {
		return err
	}
	display.RetrievedArticlesVideos(articles)
	return nil
}

//retrieveLatest ...
func (c *ArticlesCommand) retrieveLatest(queries *api.CommonQuery) CommandValidationError {
	articles, err := api.RetrieveLatestArticles(queries)
	if err != nil {
		return err
	}
	display.RetrievedArticles(articles)
	return nil
}

//update ...
func (c *ArticlesCommand) update(data *api.ArticleEdit) CommandValidationError {
	article, err := api.UpdateArticle(c.Data, data)
	if err != nil {
		return err
	}
	display.ModifiedArticle(article)
	return nil
}

//processCommonQueries read the data from the User input and put
//that data inside an CommonQuery structure
func processCommonQueries() (*api.CommonQuery, error) {
	//to store field from CommonQuery
	queries := new(api.CommonQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//processLatestQueries read the data from the User input and put
//that data inside an GetArticleQuery structure
func processLatestQueries() (*api.CommonQuery, error) {
	//to store field from CommonQuery
	queries := new(api.CommonQuery)
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
	article, err := api.CreateArticle(data)
	if err != nil {
		return err
	}
	display.ModifiedArticle(article)
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
