package prompt

import (
	"errors"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/lithammer/fuzzysearch/fuzzy"

	"github.com/jotadrilo/cookify/internal/logger"
)

func FuzzyFilter(filterValue string, optValue string, optIndex int) bool {
	return fuzzy.Match(filterValue, optValue)
}

func DefaultFilter(filterValue string, optValue string, optIndex int) bool {
	return strings.Contains(strings.ToLower(optValue), strings.ToLower(filterValue))
}

type AskOptions struct {
	Prompt  survey.Prompt
	AskOpts []survey.AskOpt
}

type AskOption func(*AskOptions)

func makeAskOptions(opts ...AskOption) *AskOptions {
	var v = &AskOptions{}

	for _, o := range opts {
		o(v)
	}

	return v
}

func WithAskPrompt(prompt survey.Prompt) AskOption {
	return func(o *AskOptions) {
		o.Prompt = prompt
	}
}

func WithSimpleFilter() AskOption {
	return func(o *AskOptions) {
		o.AskOpts = append(o.AskOpts, survey.WithFilter(DefaultFilter))
	}
}

func WithFuzzyFilter() AskOption {
	return func(o *AskOptions) {
		o.AskOpts = append(o.AskOpts, survey.WithFilter(FuzzyFilter))
	}
}

func WithAskKeepFilter() AskOption {
	return func(o *AskOptions) {
		o.AskOpts = append(o.AskOpts, survey.WithKeepFilter(true))
	}
}

func Ask(v any, opts ...AskOption) error {
	var o = makeAskOptions(opts...)
	return survey.AskOne(o.Prompt, v, o.AskOpts...)
}

func MustAsk(v any, opts ...AskOption) {
	if err := Ask(v, opts...); err != nil {
		if errors.Is(err, terminal.InterruptErr) {
			logger.Fatal("Execution cancelled")
		}
		logger.Fatal(err)
	}
}

func AskFloat32(v any, opts ...AskOption) error {
	var o = makeAskOptions(opts...)
	var s string

	if err := survey.AskOne(o.Prompt, &s, o.AskOpts...); err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	f, err := strconv.ParseFloat(strings.ReplaceAll(s, ",", "."), 32)
	if err != nil {
		return err
	}

	fptr := v.(*float32)
	*fptr = float32(f)

	return nil
}

func MustAskFloat32(v any, opts ...AskOption) {
	if err := AskFloat32(v, opts...); err != nil {
		if errors.Is(err, terminal.InterruptErr) {
			logger.Fatal("Execution cancelled")
		}
		logger.Fatal(err)
	}
}

type AskConfirmOptions struct {
	Message string
}

type AskConfirmOption func(*AskConfirmOptions)

func makeAskConfirmOptions(opts ...AskConfirmOption) *AskConfirmOptions {
	var v = AskConfirmOptions{
		Message: "Are you OK?",
	}
	for _, o := range opts {
		o(&v)
	}
	return &v
}

func WithAskConfirmMessage(message string) AskConfirmOption {
	return func(o *AskConfirmOptions) {
		o.Message = message
	}
}

func AskConfirm(opts ...AskConfirmOption) (bool, error) {
	var o = makeAskConfirmOptions(opts...)

	var ok bool

	if err := survey.AskOne(&survey.Confirm{
		Message: o.Message,
	}, &ok); err != nil {
		return false, err
	}

	return ok, nil
}

func MustAskConfirm(opts ...AskConfirmOption) bool {
	ok, err := AskConfirm(opts...)
	if err != nil {
		if errors.Is(err, terminal.InterruptErr) {
			logger.Fatal("Execution cancelled")
		}
		logger.Fatal(err)
	}
	return ok
}
