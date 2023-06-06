package mock

//go:generate mockery --name animal --output ./
type animal interface {
	run() bool
	eat() bool
}
