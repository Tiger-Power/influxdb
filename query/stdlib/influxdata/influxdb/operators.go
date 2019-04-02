package influxdb

import (
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/values"
)

const ReadRangePhysKind = "ReadRangePhysKind"

type ReadRangePhysSpec struct {
	plan.DefaultCost

	Bucket   string
	BucketID string

	Bounds flux.Bounds
}

func (s *ReadRangePhysSpec) Kind() plan.ProcedureKind {
	return ReadRangePhysKind
}
func (s *ReadRangePhysSpec) Copy() plan.ProcedureSpec {
	ns := new(ReadRangePhysSpec)

	ns.Bucket = s.Bucket
	ns.BucketID = s.BucketID

	ns.Bounds = s.Bounds

	return ns
}

// TimeBounds implements plan.BoundsAwareProcedureSpec.
func (s *ReadRangePhysSpec) TimeBounds(predecessorBounds *plan.Bounds) *plan.Bounds {
	return &plan.Bounds{
		Start: values.ConvertTime(s.Bounds.Start.Time(s.Bounds.Now)),
		Stop:  values.ConvertTime(s.Bounds.Stop.Time(s.Bounds.Now)),
	}
}
