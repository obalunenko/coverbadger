package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/obalunenko/version"
	logging "github.com/sirupsen/logrus"
)

func printVersion(_ context.Context) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	_, err := fmt.Fprintf(w, `
| app_name:	%s	|
| version:	%s	|
| short_commit:	%s	|
| commit:	%s	|
| build_date:	%s	|
| goversion:	%s	|
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

`,
		version.GetAppName(),
		version.GetVersion(),
		version.GetShortCommit(),
		version.GetCommit(),
		version.GetBuildDate(),
		version.GetGoVersion())
	if err != nil {
		logging.WithError(err).Error("print version")
	}
}
