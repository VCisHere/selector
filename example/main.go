package main

import (
	"fmt"
	"github.com/VCisHere/selector"
	"time"
)

func main() {
	var SelectQuestionTemplate = `
{{- define "option"}}
   {{- if eq .SelectedIndex .CurrentIndex }}{{color .Config.Icons.SelectFocus.Format }}{{ .Config.Icons.SelectFocus.Text }} {{else}}{{color "default"}}  {{end}}
   {{- .CurrentOpt.Value}}{{ if ne ($.GetDescription .CurrentOpt) "" }} - {{color "cyan"}}{{ $.GetDescription .CurrentOpt }}{{end}}
   {{- color "reset"}}
{{end}}
{{- color "default"}}{{ .Message }}{{ .FilterMessage }}{{color "reset"}}
 {{- "  "}}
 {{- "\n"}}
 {{- range $ix, $option := .PageEntries}}
   {{- template "option" $.IterateOption $ix $option}}
 {{- end}}`

	selector.SelectQuestionTemplate = SelectQuestionTemplate
	prompt := &selector.Select{
		Message: "Choose a color:",
		Options: []string{"red", "blue", "green"},
	}

	go func() {
		i := 0
		for {
			prompt.Options = append(prompt.Options, fmt.Sprintf("color%v", i))
			i++
			time.Sleep(time.Second)
		}
	}()

	var color string
	err := selector.AskOne(prompt, &color, selector.WithValidator(selector.Required))
	if err != nil {
		panic(err)
	}
	fmt.Println(color)
}
