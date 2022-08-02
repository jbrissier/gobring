package gobring

import "github.com/jbrissier/gobring/model"

type GoBringConfig struct {
	DBpath string
}

func (g *GoBringConfig) GetDB() *model.BringDB {
	return model.NewBringDB(g.DBpath)

}
