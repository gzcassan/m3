	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/ts"
	Open(namespace ts.ID, shard uint32, start time.Time) error
	// Write will write the id and data pair and returns an error on a write error
	Write(id ts.ID, data []byte) error
	// WriteAll will write the id and all byte slices and returns an error on a write error
	WriteAll(id ts.ID, data [][]byte) error
	Open(namespace ts.ID, shard uint32, start time.Time) error
	// Read returns the next id and data pair or error, will return io.EOF at end of volume
	Read() (id ts.ID, data []byte, err error)
// Options represents the options for filesystem persistence
	// SetClockOptions sets the clock options
	SetClockOptions(value clock.Options) Options

	// ClockOptions returns the clock options
	ClockOptions() clock.Options


	// SetThroughputCheckInterval sets the filesystem throughput check interval
	SetThroughputCheckInterval(value time.Duration) Options

	// ThroughputCheckInterval returns the filesystem throughput check interval
	ThroughputCheckInterval() time.Duration

	// SetThroughputLimitMbps sets the filesystem throughput limit
	SetThroughputLimitMbps(value float64) Options

	// ThroughputLimitMbps returns the filesystem throughput limit
	ThroughputLimitMbps() float64