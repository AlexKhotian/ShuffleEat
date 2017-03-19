package Tests

import (
	"ShuffleEat/Model/DataBase"
	"testing"
)

func TestDataBase(t *testing.T) {
	con := new(DataBase.Connector)
	con.InitDatabase()
}
