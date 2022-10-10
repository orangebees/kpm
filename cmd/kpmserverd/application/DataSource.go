package application

type DataSource interface {
	SearchName(name string) string
	SearchSubPkgNames(SubPkgNames []string) string

	Publish(pkgtgz []byte) string
}
