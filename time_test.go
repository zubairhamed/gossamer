package gossamer

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimePeriod(t *testing.T) {
	var tp *TimePeriod
	var err error

	tp = NewDefaultTimePeriod()

	assert.NotNil(t, tp)
	assert.True(t, tp.From().IsZero())
	assert.True(t, tp.To().IsZero())

	tp.FromTime = time.Now()
	assert.False(t, tp.From().IsZero())

	tp.ToTime = time.Now()
	assert.False(t, tp.To().IsZero())

	tp = NewDefaultTimePeriod()
	tp.FromTime, err = time.Parse(STD_TIME_FORMAT_INSTANT, "2014-03-01T13:00:00+08:00")
	assert.Nil(t, err)

	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)

	err = encoder.Encode(tp)
	assert.Nil(t, err)
	assert.Equal(t, string(buf.String()), "\"2014-03-01T13:00:00+08:00\"\n")

	buf = new(bytes.Buffer)
	encoder = json.NewEncoder(buf)
	tp.ToTime, err = time.Parse(STD_TIME_FORMAT_INSTANT, "2015-03-01T13:00:00+08:00")
	err = encoder.Encode(tp)
	assert.Nil(t, err)
	assert.Equal(t, string(buf.String()), "\"2014-03-01T13:00:00.00Z/2015-03-01T13:00:00.00Z\"\n")
}

func TestTimeInstant(t *testing.T) {
	ti := NewTimeInstant(time.Now())
	assert.NotNil(t, ti)
}
