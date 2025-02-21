package types

import (
	"path/filepath"
)

type Sqlite struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *Sqlite) Dsn() string {
	return filepath.Join(s.Path, s.Dbname+".db")
}
