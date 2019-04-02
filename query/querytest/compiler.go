package querytest

import (
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/stdlib/influxdata/influxdb/v1"
	"github.com/influxdata/influxdb/query/influxql"
	"github.com/influxdata/influxdb/query/stdlib/influxdata/influxdb"
)

// FromInfluxJSONCompiler returns a compiler that replaces all From operations with FromJSON.
func FromInfluxJSONCompiler(c *influxql.Compiler, jsonFile string) flux.Compiler {
	c.WithLogicalPlannerOptions(plan.AddLogicalRules(ReplaceFromRule{Filename: jsonFile}))
	return c
}

type ReplaceFromRule struct {
	Filename string
}

func (ReplaceFromRule) Name() string {
	return "ReplaceFromRule"
}

func (ReplaceFromRule) Pattern() plan.Pattern {
	return plan.Pat(influxdb.FromKind)
}

func (r ReplaceFromRule) Rewrite(n plan.Node) (plan.Node, bool, error) {
	if err := n.ReplaceSpec(&v1.FromInfluxJSONProcedureSpec{
		File: r.Filename,
	}); err != nil {
		return nil, false, err
	}
	return n, true, nil
}
