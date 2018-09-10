package cmd

import (
	"github.com/c-bata/go-prompt"
)

var suggests []prompt.Suggest

func ignorePrompt(ignores []string) string {
	for _, ignore := range ignores {
		suggests = append(suggests, prompt.Suggest{Text: ignore})
	}

	return prompt.Input(
		"Choose a .gitignore template >>> ",
		completer,
		prompt.OptionTitle("gitignore-prompt"),
		prompt.OptionPrefixTextColor(prompt.Green),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.Blue),
	)
}

func completer(in prompt.Document) []prompt.Suggest {
	return prompt.FilterContains(suggests, in.GetWordBeforeCursor(), true)
}
