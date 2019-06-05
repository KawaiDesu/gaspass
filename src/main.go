package main

import (
	"fmt"
	"github.com/KawaiDesu/gaspass/src/gaspass"
	"github.com/chzyer/readline"
	flags "github.com/jessevdk/go-flags"
	"os"
)

type Resource struct {
	PassLen int
	Serial  int
	Host    string
}

func checkOpts() bool {
	return true
}

func processFlags() {
	_, err := flags.Parse(&opts)
	if flags.WroteHelp(err) {
		os.Exit(1)
	}

}

var (
	opts struct {
		CharsLower    bool   `short:"l" long:"lower" description:"Use lower-case characters for generating password"`
		CharsUpper    bool   `short:"u" long:"upper" description:"Use upper-case characters for generating password"`
		CharsNumbers  bool   `short:"n" long:"numeric" description:"Use numeric characters for generating password"`
		CharsSpecials bool   `short:"s" long:"specials" description:"Use speacial (punctuation) characters for generating password"`
		Length        int    `short:"q" long:"quantity" default:"16" description:"Set number of characters in the password"`
		Salt          string `short:"r" long:"resource" description:"Resource name (url or some descriptive text) for which password will be generated"`
		Counter       string `short:"c" long:"counter" default:"0" description:"Serial number of the password for the same resource"`
		ActionAdd     bool   `short:"A" long:"add" description:"Add resource record to the database"`
		ActionRemove  bool   `short:"D" long:"delete" description:"Remove resource record from the database"`
		ActionUseRes  bool   `short:"R" long:"use-resource" description:"Use existing resource"`
		ActionList    bool   `short:"L" long:"list" description:"List resource records in the database"`
		ActionBench   bool   `short:"B" long:"bench" description:"Run benchmark"`
	}

	charSet string = ""
)

func main() {

	processFlags()

	privKey, err := readline.Password("Enter your key:")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	p := gaspass.Params{
		PrivKey:     privKey,
		Salt:        []byte(opts.Salt),
		Counter:     []byte(opts.Counter),
		PassLength:  opts.Length,
		UseLower:    opts.CharsLower,
		UseUpper:    opts.CharsUpper,
		UseNumber:   opts.CharsNumbers,
		UseSpecials: opts.CharsSpecials,
	}

	resultPass, err := p.GeneratePassword()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resultPass)
}
