package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type CommandLine struct {
	Commands []CommandI
	tw       *tabwriter.Writer
}

func NewCli() *CommandLine {
	return &CommandLine{}
}

var (
	Cli             *CommandLine
	adminConfig     *AdminConfig
	articlesCmd     *ArticlesCommand
	authCmd         *AuthCommand
	commentsCmd     *CommentsCommand
	followersCmd    *FollowersCommand
	listingsCmd     *ListingsCommand
	organizationCmd *OrganizationsCommand
	podcastCmd      *PodcastsCommand
	profileImageCmd *ProfileImageCommand
	readingListCmd  *ReadingListsCommand
	webhookCmd      *WebhooksCommand
	tagsCmd         *TagsCommand
)

func init() {
	Cli = NewCli()
	//initialize tabwriter
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	Cli.tw = tw

	//commands
	adminConfig = NewAdminConfigCmd()
	Cli.Commands = append(Cli.Commands, adminConfig)

	articlesCmd = NewArticlesCmd()
	Cli.Commands = append(Cli.Commands, articlesCmd)

	authCmd = NewAuthCmd()
	Cli.Commands = append(Cli.Commands, authCmd)

	commentsCmd = NewCommentsCommand()
	Cli.Commands = append(Cli.Commands, commentsCmd)

	followersCmd = NewFollowersCommand()
	Cli.Commands = append(Cli.Commands, followersCmd)

	listingsCmd = NewListingsCommand()
	Cli.Commands = append(Cli.Commands, listingsCmd)

	organizationCmd = NewOrganizationsCommand()
	Cli.Commands = append(Cli.Commands, organizationCmd)

	podcastCmd = NewPodcastsCommand()
	Cli.Commands = append(Cli.Commands, podcastCmd)

	profileImageCmd = NewProfileImageCmd()
	Cli.Commands = append(Cli.Commands, profileImageCmd)

	readingListCmd = NewReadingListsCommand()
	Cli.Commands = append(Cli.Commands, readingListCmd)

	tagsCmd = NewTagsCommand()
	Cli.Commands = append(Cli.Commands, tagsCmd)

	webhookCmd = NewWebhooksCmd()
	Cli.Commands = append(Cli.Commands, webhookCmd)
}

func (cli *CommandLine) printUsage() {
	fmt.Println("\033[1;36mUsage:\033[0m")
	fmt.Printf("  \033[3;32mdevto\033[0m \033[3;33m<command> <data> <subcommand> <flags>\033[0m\n")
	fmt.Println()
	fmt.Println("\033[1;36mCore commands:\033[0m")
	for _, cmd := range cli.Commands {
		cmd.Helper(cli.tw)
	}
	cli.tw.Flush()
}

func (cli *CommandLine) validateArgs() bool {
	if len(os.Args) < 2 {
		cli.printUsage()
		return false
	}
	return true
}

func (cli *CommandLine) Execute() {
	if !cli.validateArgs() {
		return
	}

	//on validateArgs we make sure it has more than one arg
	argsCount := len(os.Args)
	switch os.Args[1] {
	case "auth":
		//api_key
		if argsCount < 3 {
			cli.printUsage()
			return
		}
		data := os.Args[2]
		authCmd.SetData(data)
		//need to find a better way to display the data
		err := authCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "admin-config":
		switch argsCount {
		case 2:
			err := adminConfig.ActivateSubcommand("retrieve")
			if err != nil {
				cli.printUsage()
				return
			}
		case 3:
			if os.Args[2] == "update" {
				err := adminConfig.ActivateSubcommand("update")
				if err != nil {
					cli.printUsage()
					return
				}
			}
		}
		err := adminConfig.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "articles":
		switch {
		case argsCount == 2:
			err := articlesCmd.ActivateSubcommand("retrieve")
			if err != nil {
				fmt.Fprintf(os.Stdout, "%v\n", err)
				cli.printUsage()
				return
			}
		case argsCount > 2:
			switch os.Args[2] {
			case "update":
				err := articlesCmd.ActivateSubcommand("update")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
				if argsCount == 4 {
					articleID := os.Args[3]
					articlesCmd.SetData(articleID)
				} else {
					cli.printUsage()
					return
				}
			case "create":
				err := articlesCmd.ActivateSubcommand("create")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
			case "latest":
				err := articlesCmd.ActivateSubcommand("retrieve_latest")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
				query := func() bool {
					if os.Args[argsCount-1] == "-q" {
						return true
					}
					return false
				}()
				//in case queries are unables
				if query {
					err := articlesCmd.ActivateSubcommand("latest_query")
					if err != nil {
						fmt.Fprintf(os.Stdout, "%v\n", err)
						cli.printUsage()
						return
					}
				}
			case "videos":
				err := articlesCmd.ActivateSubcommand("retrieve_videos")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
			case "me":
				if argsCount == 4 {
					if os.Args[3] == "-p" {
						articlesCmd.SetData("/published")
					} else if os.Args[3] == "-up" {
						articlesCmd.SetData("/unpublished")
					} else if os.Args[3] == "-all" {
						articlesCmd.SetData("/all")
					} else {
						cli.printUsage()
						return
					}
				}
				err := articlesCmd.ActivateSubcommand("retrieve_me")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
			default:
				err := articlesCmd.ActivateSubcommand("retrieve")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
				username := func() string {
					if argsCount > 2 && os.Args[2] != "-q" {
						if _, err := strconv.ParseInt(os.Args[2], 10, 32); err != nil {
							return os.Args[2]
						}
						return ""
					} else {
						return ""
					}
				}()
				query := func() bool {
					if os.Args[argsCount-1] == "-q" {
						return true
					}
					return false
				}()
				id := func() string {
					if argsCount > 2 && os.Args[2] != "-q" {
						if _, err := strconv.ParseInt(os.Args[2], 10, 32); err != nil {
							return ""
						}
						return os.Args[2]
					} else {
						return ""
					}
				}()
				//in case queries are unables
				if query {
					err := articlesCmd.ActivateSubcommand("retrieve_query")
					if err != nil {
						fmt.Fprintf(os.Stdout, "%v\n", err)
						cli.printUsage()
						return
					}
				}
				if len(username) > 0 {
					articlesCmd.SetData(username)
				}
				if len(id) > 0 {
					articlesCmd.SetData(id)
					err := articlesCmd.ActivateSubcommand("retrieve_id")
					if err != nil {
						fmt.Fprintf(os.Stdout, "%v\n", err)
						cli.printUsage()
						return
					}
				}
			}
		default:
			cli.printUsage()
		}
		err := articlesCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "comments":
		switch argsCount {
		case 4:
			commentsCmd.SetData(os.Args[3])
			if os.Args[2] == "-a_id" {
				err := commentsCmd.ActivateSubcommand("retrieve_aid")
				if err != nil {
					cli.printUsage()
					return
				}
			} else if os.Args[2] == "-p_id" {
				err := commentsCmd.ActivateSubcommand("retrieve_pid")
				if err != nil {
					cli.printUsage()
					return
				}
			} else if os.Args[2] == "-id" {
				err := commentsCmd.ActivateSubcommand("retrieve_id")
				if err != nil {
					cli.printUsage()
					return
				}
			} else {
				cli.printUsage()
				return
			}
		default:
			cli.printUsage()
			return
		}
		err := commentsCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "tags":
		switch argsCount {
		case 2:
			err := tagsCmd.ActivateSubcommand("retrieve")
			if err != nil {
				cli.printUsage()
				return
			}
		case 3:
			if os.Args[2] == "follows" {
				err := tagsCmd.ActivateSubcommand("retrieve_follows")
				if err != nil {
					cli.printUsage()
					return
				}
			} else {
				cli.printUsage()
				return
			}
		default:
			cli.printUsage()
			return
		}
		err := tagsCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "followers":
		switch argsCount {
		case 2:
			err := followersCmd.ActivateSubcommand("retrieve")
			if err != nil {
				cli.printUsage()
				return
			}
		default:
			cli.printUsage()
			return
		}
		err := followersCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "listings":
		switch argsCount {
		case 2:
			err := listingsCmd.ActivateSubcommand("retrieve")
			if err != nil {
				cli.printUsage()
				return
			}
		case 3:
			switch os.Args[2] {
			case "create":
				err := listingsCmd.ActivateSubcommand("create")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
			default:
				cli.printUsage()
				return
			}
		case 4:
			switch os.Args[2] {
			case "retrieve":
				err := listingsCmd.ActivateSubcommand("retrieve_id")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
				listingsCmd.SetData(os.Args[3])
			case "update":
				err := listingsCmd.ActivateSubcommand("update")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
				listingsCmd.SetData(os.Args[3])
			default:
				cli.printUsage()
				return
			}
		default:
			cli.printUsage()
			return
		}
		err := listingsCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "organizations":
		switch argsCount {
		case 3:
			err := organizationCmd.ActivateSubcommand("retrieve")
			if err != nil {
				cli.printUsage()
				return
			}
			//username of organization
			organizationCmd.SetData(os.Args[2])
		case 4:
			organizationCmd.SetData(os.Args[2])
			switch os.Args[3] {
			case "-u":
				err := organizationCmd.ActivateSubcommand("retrieve_users")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
			case "-l":
				err := organizationCmd.ActivateSubcommand("retrieve_listing")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
			case "-a":
				err := organizationCmd.ActivateSubcommand("retrieve_articles")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
			default:
				cli.printUsage()
				return
			}
		default:
			cli.printUsage()
			return
		}
		err := organizationCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "podcasts":
		err := podcastCmd.ActivateSubcommand("retrieve")
		if err != nil {
			cli.printUsage()
			return
		}
		err = podcastCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "reading_lists":
		err := readingListCmd.ActivateSubcommand("retrieve")
		if err != nil {
			cli.printUsage()
			return
		}
		err = readingListCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "webhooks":
		switch argsCount {
		case 2:
			err := webhookCmd.ActivateSubcommand("retrieve")
			if err != nil {
				cli.printUsage()
				return
			}
		case 3:
			switch os.Args[2] {
			case "create":
				err := webhookCmd.ActivateSubcommand("create")
				if err != nil {
					cli.printUsage()
					return
				}
			default:
				//refactor
				if os.Args[2] != "delete" {
					//assuming want to retrieve by id
					err := webhookCmd.ActivateSubcommand("retrieve_id")
					if err != nil {
						cli.printUsage()
						return
					}
					webhookCmd.SetData(os.Args[2])
				} else {
					cli.printUsage()
					return
				}
			}
		case 4:
			switch os.Args[2] {
			case "delete":
				err := webhookCmd.ActivateSubcommand("delete")
				if err != nil {
					cli.printUsage()
					return
				}
				webhookCmd.SetData(os.Args[3])
			default:
				cli.printUsage()
				return
			}
		default:
			cli.printUsage()
			return
		}
		err := webhookCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	case "profile_images":
		switch argsCount {
		case 3:
			err := profileImageCmd.ActivateSubcommand("retrieve")
			if err != nil {
				cli.printUsage()
				return
			}
			profileImageCmd.SetData(os.Args[2])
		default:
			cli.printUsage()
			return
		}
		err := profileImageCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	default:
		cli.printUsage()
	}
}
