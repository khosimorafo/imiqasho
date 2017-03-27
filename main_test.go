package imiqasho_test

import (
	"testing"
	"os"
	"github.com/khosimorafo/imiqasho"
	//"github.com/antonholmquist/jason"
	//"encoding/json"
	"encoding/json"
	"github.com/antonholmquist/jason"
)

var a imiqasho.App
var tenant_id string

func TestMain(m *testing.M) {

	a = imiqasho.App{}

	a.Initialize()

	//ensureTableExists()

	code := m.Run()

	//clearTable()

	os.Exit(code)
}

/*
func TestCreateAndDeleteTenant(t *testing.T) {

	tenant := imiqasho.Tenant{Name: "M Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
	var i imiqasho.EntityInterface
	i = tenant
	result, entity, error := imiqasho.Create(i)

	if result != "success" {

		t.Errorf("Failed to create tenant. Result = %v", error.Error())
	}

	if error != nil {

		t.Errorf("Failed to create tenant %v", error.Error())
	}

	if entity == nil {

		t.Errorf("Failed to create tenant. Entity is empty!")
	}

	ten := entity

	res, err := imiqasho.Delete(*ten)

	if res != "success" {

		t.Errorf("Failed to delete tenant. Result = %v", result)
	}

	if err != nil {

		t.Errorf("Failed to delete tenant %v", error)
	}
}
*/

func TestCreateAndUpdateTenant(t *testing.T) {

	tenant := imiqasho.Tenant{Name: "M Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
	var i imiqasho.EntityInterface
	i = tenant
	result, entity, error := imiqasho.Create(i)

	if result != "success" {

		t.Errorf("Failed to create tenant. Result = %v", error.Error())
	}

	if error != nil {

		t.Errorf("Failed to create tenant %v", error.Error())
	}

	if entity == nil {

		t.Errorf("Failed to create tenant. Entity is empty!")
	}

	b, _ := json.Marshal(entity)
	v, _ := jason.NewObjectFromBytes(b)

	id, err := v.GetString("id")

	ten := imiqasho.Tenant{ID:id, Name:"M Tenant - New Name"}

	res, ent, err := imiqasho.Update(ten)

	if res != "success" {

		t.Errorf("Failed to update tenant. Result = %v", result)
	}

	if err != nil {

		t.Errorf("Failed to update tenant %v", error)
	}

	b1, _ := json.Marshal(ent)
	v1, _ := jason.NewObjectFromBytes(b1)

	id1, _ := v1.GetString("name")

	t.Log("The update name is", id1)
}




func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
