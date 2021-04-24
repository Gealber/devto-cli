package cmd

import (
	"fmt"
	"os"
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
	Cli         *CommandLine
	adminConfig *AdminConfig
	authCmd     *AuthCommand
	articlesCmd *ArticlesCommand
)

func init() {
	Cli = NewCli()
	//initialize tabwriter
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	Cli.tw = tw

	//commands
	adminConfig = NewAdminConfigCmd()
	Cli.Commands = append(Cli.Commands, adminConfig)
	authCmd = NewAuthCmd()
	Cli.Commands = append(Cli.Commands, authCmd)
	articlesCmd = NewArticlesCmd()
	Cli.Commands = append(Cli.Commands, articlesCmd)
}

func (cli *CommandLine) printUsage() {
	fmt.Println("\033[1;36mUsage:\033[0m")
	fmt.Printf("  \033[3;32mdevto\033[0m \033[3;33m<command> <subcommand>\033[0m\n")
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

	//devto articles create
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
			default:
				err := articlesCmd.ActivateSubcommand("retrieve")
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err)
					cli.printUsage()
					return
				}
				username := func() string {
					if argsCount > 2 && os.Args[2] != "-q" {
						return os.Args[2]
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
			}
		default:
			cli.printUsage()
		}
		err := articlesCmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
			cli.printUsage()
		}
	default:
		cli.printUsage()
	}
}
