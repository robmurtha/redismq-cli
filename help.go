package main

import (
	"fmt"
	"os"
	"text/template"
)

var cmdHelp = &Command{
	Usage: "help [topic]",
	Long:  `Help shows usage for a command or other topic.`,
}

func init() {
	cmdHelp.Run = runHelp // break init loop
}

func runHelp(cmd *Command, args []string) {
	if len(args) == 0 {
		printUsage()
		return // not os.Exit(2); success
	}
	if len(args) != 1 {
		fmt.Println("too many arguments")
		os.Exit(2)
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.printUsage()
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic: %q. Run 'redismq-cli help'.\n", args[0])
	os.Exit(2)
}

var usageTemplate = template.Must(template.New("usage").Parse(`
Usage: redismq-cli [command] [options] [arguments]

Commands:
{{range .Commands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-8s"}}  {{.Short}}{{end}}{{end}}{{end}}

General Options:
    -host    redis hostname    (default: localhost)
    -port    redis port        (default: 6371)
    -pass    redis password    (default: "")
    -db      redis db number   (default: 9)


Run 'redismq-cli help [command]' for details.

`[1:]))

func usage() {
	printUsage()
	os.Exit(2)
}

func printUsage() {
	usageTemplate.Execute(os.Stdout, struct {
		Commands []*Command
	}{
		commands,
	})
}
