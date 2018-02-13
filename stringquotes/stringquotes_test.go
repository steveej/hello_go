package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

// execEscape uses Golang's string quoting for ", \, \n, and regex for special cases
func execEscape(i int, str string) string {
	escapeMap := map[string]string{
		`'`: `\`,
	}

	if i > 0 { // These are escaped only after the first argument
		escapeMap[`$`] = `$`
	}

	escArg := fmt.Sprintf("%q", str)
	keys, values := make([]string, len(escapeMap)), make([]string, len(escapeMap))
	for k, v := range escapeMap {
		keys = append(keys, k)
		values = append(values, v)

		reStr := `([` + regexp.QuoteMeta(k) + `])`
		re := regexp.MustCompile(reStr)
		escArg = re.ReplaceAllStringFunc(escArg, func(s string) string {
			escaped := escapeMap[s] + s
			fmt.Printf("Found '%s', returning '%s'\n", s, escaped)
			return escaped
		})
	}
	return escArg
}

// quoteExec returns an array of quoted strings appropriate for systemd execStart usage
func quoteExec(exec []string) string {
	fmt.Printf("Quoting %q\n", exec)
	if len(exec) == 0 {
		// existing callers prefix {"/appexec", "/app/root", "/work/dir", "/env/file"} so this shouldn't occur.
		panic("empty exec")
	}

	var qexec []string
	for i, arg := range exec {
		escArg := execEscape(i, arg)
		qexec = append(qexec, escArg)
	}
	return strings.Join(qexec, " ")
}

func TestQuoteExec(t *testing.T) {
	tests := []struct {
		input  []string
		output string
	}{
		{
			input:  []string{`path`, `"arg1"`, `"'arg2'"`, `'arg3'`},
			output: `"path" "\"arg1\"" "\"\'arg2\'\"" "\'arg3\'"`,
		}, {
			input:  []string{`path`},
			output: `"path"`,
		}, {
			input:  []string{`path`, ``, `arg2`},
			output: `"path" "" "arg2"`,
		}, {
			input: []string{`path`, `"foo\bar"`, `\`}, output: `"path" "\"foo\\bar\"" "\\"`,
		}, {
			input: []string{`path with spaces`, `"foo\bar"`, `\`}, output: `"path with spaces" "\"foo\\bar\"" "\\"`,
		}, {input: []string{`path with "quo't'es" and \slashes`, `"arg"`, `\`},
			output: `"path with \"quo\'t\'es\" and \\slashes" "\"arg\"" "\\"`,
		}, {input: []string{`$path$`, `$argument`},
			output: `"$path$" "$$argument"`,
		}, {
			input:  []string{`path`, `Args\nwith\nnewlines`},
			output: `"path" "Args\\nwith\\nnewlines"`}}

	for i, tt := range tests {
		o := quoteExec(tt.input)
		if o != tt.output {
			t.Errorf("#%d: expected `%v` got `%v`", i, tt.output, o)
		}
	}
}

func main() {}
