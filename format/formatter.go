package format

type Status uint8

const (
	Read  Status = 0
	Ok    Status = 1
	Error Status = 2
)

type Formatter interface {
	Format(status Status, output string) string
}
