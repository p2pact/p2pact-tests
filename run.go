package bench

import (
	"fmt"
	"io"
)

// Execute loads the scenario's data, runs it and returns the report.
func Execute(sc *Scenario, log io.Writer) (*Report, error) {
	sc.defaults()
	msgs, credit, err := LoadData(sc.Data, log)
	if err != nil {
		return nil, err
	}
	if msgs != nil {
		sc.fill(msgs)
		if len(sc.Sizes) == 0 {
			return nil, fmt.Errorf("%s: no messages fall into the size bins", sc.Data)
		}
		for _, s := range sc.Sizes {
			fmt.Fprintf(log, "%s: %d messages, median %d chars\n", s.Name, len(s.msgs), s.Chars)
		}
	}
	res, err := Run(sc, func(s string) { fmt.Fprintln(log, "done:", s) })
	if err != nil {
		return nil, err
	}
	return &Report{Scenario: sc, Data: credit, Env: CurrentEnv(), Results: res}, nil
}
