package gossamer
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQueryOptions(t *testing.T) {
	var opts QueryOptions
	var err error

	opts, err = CreateQueryOptions("$expand=Datastreams")
	assert.Nil(t, err)
	assert.True(t, opts.ExpandSet())
	assert.Equal(t,1, len(opts.GetExpandOption().GetValue()))
	assert.Equal(t, "Datastreams", opts.GetExpandOption().GetValue()[0])

	opts, err = CreateQueryOptions("$expand=Observations,ObservedProperty")
	assert.Nil(t, err)
	assert.True(t, opts.ExpandSet())
	assert.Equal(t,2, len(opts.GetExpandOption().GetValue()))
	assert.Equal(t, "Observations", opts.GetExpandOption().GetValue()[0])
	assert.Equal(t, "ObservedProperty", opts.GetExpandOption().GetValue()[1])

	opts, err = CreateQueryOptions("$expand=Observations($select=result)")
	assert.Nil(t, err)
	assert.True(t, opts.ExpandSet())
	assert.Equal(t, 1, len(opts.GetExpandOption().GetValue()))
	assert.Equal(t, opts.GetExpandOption().GetValue()[0], "Observations($select=result)")

	opts, err = CreateQueryOptions("$select=id,Observations&$expand=Observations/FeatureOfInterest")
	assert.Nil(t, err)
	assert.True(t, opts.SelectSet())
	assert.Equal(t, 2, len(opts.GetSelectOption().GetValue()))
	assert.Equal(t, opts.GetSelectOption().GetValue()[0], "id")
	assert.Equal(t, opts.GetSelectOption().GetValue()[1], "Observations")
	assert.True(t, opts.ExpandSet())
	assert.Equal(t, 1, len(opts.GetExpandOption().GetValue()))
	assert.Equal(t, opts.GetExpandOption().GetValue()[0], "Observations/FeatureOfInterest")

	opts, err = CreateQueryOptions("$top=5")
	assert.Nil(t, err)
	assert.True(t, opts.TopSet())
	assert.Equal(t, opts.GetTopOption().GetValue(), 5)

	opts, err = CreateQueryOptions("$skip=5")
	assert.Nil(t, err)
	assert.True(t, opts.SkipSet())
	assert.Equal(t, opts.GetSkipOption().GetValue(), 5)

	opts, err = CreateQueryOptions("$expand=Datastream&$orderby=Datastreams/id desc, phenomenonTime")
	assert.Nil(t, err)
	assert.True(t, opts.ExpandSet())
	assert.True(t, opts.OrderBySet())

	opts, err = CreateQueryOptions("$expand=Datastreams/Observations/FeatureOfInterest&$filter=Datastreams/Observations/FeatureOfInterest/id eq 'FOI_1' and Datastreams/Observations/resultTime ge 2010-06-01T00:00:00Z and Datastreams/Observations/resultTime le 2010-07-01T00:00:00Z")
	assert.Nil(t, err)
	assert.True(t, opts.ExpandSet())
	assert.Equal(t, 1, len(opts.GetExpandOption().GetValue()))
	assert.Equal(t, "Datastreams/Observations/FeatureOfInterest", opts.GetExpandOption().GetValue()[0])
	assert.True(t, opts.FilterSet())


	/*
	$filter=
		Datastreams/Observations/FeatureOfInterest/id eq 'FOI_1'
		and
		Datastreams/Observations/resultTime ge 2010-06-01T00:00:00Z
		and
		Datastreams/Observations/resultTime le 2010-07-01T00:00:00Z

		- Target
		- Operation
		- OpValue

	- Operand
	-

	 */

	opts, err = CreateQueryOptions("$count=true")
	assert.Nil(t, err)
	assert.True(t, opts.CountSet())
	assert.True(t, opts.GetCountOption().GetValue())

	opts, err = CreateQueryOptions("$filter=geo.distance(Locations/location,￼geography'POINT(-122, 43)') gt 1")
	assert.Nil(t, err)
	assert.True(t, opts.FilterSet())

	opts, err = CreateQueryOptions("$select=result,resultTime")
	assert.Nil(t, err)
	assert.True(t, opts.SelectSet())

	opts, err = CreateQueryOptions("$orderby=result")
	assert.Nil(t, err)
	assert.True(t, opts.OrderBySet())

	opts, err = CreateQueryOptions("$top=5&$orderby=phenomenonTime desc")
	assert.Nil(t, err)
	assert.True(t, opts.TopSet())
	assert.True(t, opts.OrderBySet())

	opts, err = CreateQueryOptions("$skip=2&$top=2&$orderby=resultTime")
	assert.Nil(t, err)
	assert.True(t, opts.SkipSet())
	assert.True(t, opts.TopSet())
	assert.True(t, opts.OrderBySet())

	opts, err = CreateQueryOptions("$filter=result lt 10.00")
	assert.Nil(t, err)
	assert.True(t, opts.FilterSet())

	opts, err = CreateQueryOptions("$filter=Datastream/id eq '1'")
	assert.Nil(t, err)
	assert.True(t, opts.FilterSet())
}

