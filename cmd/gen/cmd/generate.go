package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type generateOptions struct {
	context string
	output  string
}

var generateOpts = &generateOptions{}

var generateCmd = &cobra.Command{
	Use:   "generate [--context .]",
	Short: "Generate the static webiste",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGenerate(generateOpts)
	},
}

func init() {
	generateCmd.PersistentFlags().StringVar(&generateOpts.context, "context", ".", "base path for the proposals repository in your local file system")
	generateCmd.PersistentFlags().StringVar(&generateOpts.output, "output", "./site", "where the generate website will be stored")

	rootCmd.AddCommand(generateCmd)
}

func runGenerate(opts *generateOptions) error {
	proposalsDir := path.Join(opts.context, "proposals")
	info, err := os.Stat(proposalsDir)
	if os.IsNotExist(err) {
		return errors.Wrap(err, "we expect a proposals directory inside the repository.")
	}
	if info.IsDir() == false {
		return errors.New("the expected proposal directory has to be a directory, not a file")
	}

	files, err := ioutil.ReadDir(proposalsDir)
	if err != nil {
		return err
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	proposalsTable := map[int]struct {
		Title   string
		Status  string
		Authors string
	}{}
	for _, f := range files {
		readmeFile, err := os.Open(path.Join(proposalsDir, f.Name(), "README.md"))
		if err != nil {
			return errors.Wrap(err, "error reading the README.md proposal")
		}
		readme, err := ioutil.ReadAll(readmeFile)
		if err != nil {
			return errors.Wrap(err, "error reading the README.md proposal")
		}
		var buf bytes.Buffer
		context := parser.NewContext()
		if err := md.Convert(readme, &buf, parser.WithContext(context)); err != nil {
			panic(err)
		}
		metaData := meta.Get(context)
		id, err := strconv.Atoi(f.Name())
		if err != nil {
			return err
		}
		proposalsTable[id] = struct {
			Title   string
			Status  string
			Authors string
		}{
			Title:   metaData["title"].(string),
			Status:  metaData["status"].(string),
			Authors: metaData["authors"].(string),
		}
	}
	t, err := template.New("site").Parse(t)
	err = t.ExecuteTemplate(os.Stdout, "site", proposalsTable)
	return nil
}

var t = `<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, height=device-height, initial-scale=1.0"/>
        <title>Tinkerbell's Proposals</title>
        <style>
            body {margin: 5% auto; background: #f2f2f2; color: #444444; font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif; font-size: 16px; line-height: 1.8; text-shadow: 0 1px 0 #ffffff; max-width: 73%;}
            code {background: white;}
            a {border-bottom: 1px solid #444444; color: #444444; text-decoration: none;}
            a:hover {border-bottom: 0;}
        </style>
    </head>

    <body>
        <div>
            <h1>Tinkerbell's Proposals</h1>
        </div>
        <div>
            <p>Tinkerbell uses an open forms of proposals that communities and
            contributors can use to discuss evolution of the project. This is a
            list of them representing their current state.</p>
        </div>
        <div>
            <ul>
                {{ range $key, $value := . }}
                <li>[<a href="https://github.com/tinkerbell/proposals/blob/master/proposals/{{ printf "%04d" $key }}/README.md">{{ printf "%04d" $key }}</a>]: {{ $value.Title }} - <b>status:{{ $value.Status }}</b> - authored by: {{ $value.Authors }}</li>
                {{ end }}
            </ul>
        </div>
    </body>
</html>
`
