package imiqasho_test

import (
	"testing"
	"os"
	"github.com/khosimorafo/imiqasho"

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


/*
func TestCreateUpdateAndDeleteTenant(t *testing.T) {

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

	// Update created tenant

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

	// Remove created tenant
	ten1 := entity
	res1, err1 := imiqasho.Delete(*ten1)

	if res1 != "success" {

		t.Errorf("Failed to delete tenant. Result = %v", res)
	}

	if err1 != nil {

		t.Errorf("Failed to delete tenant %v", error)
	}
}
*/

/*

func TestCreateReadAndDeleteTenant(t *testing.T) {

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

	// Update created tenant

	b, _ := json.Marshal(entity)
	v, _ := jason.NewObjectFromBytes(b)
	id, err := v.GetString("id")
	ten := imiqasho.Tenant{ID:id}

	res, ent, err := imiqasho.Read(ten)

	if res != "success" {

		t.Errorf("Failed to read tenant. Result = %v", result)
	}

	if err != nil {

		t.Errorf("Failed to read tenant %v", error)
	}

	b1, _ := json.Marshal(ent)
	v1, _ := jason.NewObjectFromBytes(b1)

	status, _ := v1.GetString("status")

	t.Log("The read tenant's status is ", status)

	// Remove created tenant
	ten1 := entity
	res1, err1 := imiqasho.Delete(*ten1)

	if res1 != "success" {

		t.Errorf("Failed to delete tenant. Result = %v", res)
	}

	if err1 != nil {

		t.Errorf("Failed to delete tenant %v", error)
	}
}

*/

/*
func TestGetTenants(t *testing.T) {

	// Create tenant
	tenant := imiqasho.Tenant{Name: "M Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
	var i imiqasho.EntityInterface
	i = tenant
	_, entity, _ := imiqasho.Create(i)

	// Query tenant list
	filters := map[string]string{}
	result, tenants, _ := imiqasho.GetTenants(filters)

	if result != "success" {

		t.Errorf("Failed to quiry tenants. d\n")
	}

	t.Log("Length is : ", len(*tenants))

	if len(*tenants) < 1{

		t.Errorf("Tenant list is empty! d\n")
	}

	// Delete created tenant
	ten := entity
	imiqasho.Delete(*ten)
}

*/


func TestCreateTenantFirstInvoice(t *testing.T)  {

	// Create tenant.
	tenant := imiqasho.Tenant{MoveInDate:"2017-05-13", Name: "M Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
	var i imiqasho.EntityInterface
	i = tenant
	_, entity, _ := imiqasho.Create(i)

	if entity == nil {

		t.Errorf("Failed to create tenant. Entity is empty!")
		return
	}

	// Create first tenant invoice.
	b, _ := json.Marshal(entity)
	v, _ := jason.NewObjectFromBytes(b)
	id, _ := v.GetString("id")
	in_date, _ := v.GetString("move_in_date")


	ten := imiqasho.Tenant{ID:id, MoveInDate:in_date}

	result, inv, error := ten.CreateFirstTenantInvoice()

	b_inv, _ := json.Marshal(inv)
	v_inv, _ := jason.NewObjectFromBytes(b_inv)
	id_inv, _ := v_inv.GetString("id")

	if result != "success" {

		t.Errorf("Failed to create invoice. Result = %v", result)
		// Delete tenant
		imiqasho.Delete(ten)
		return
	}

	if error != nil {

		t.Errorf("Failed to create invoice %v", error)
		// Delete tenant
		imiqasho.Delete(ten) // May cause a test time error. But its unimportant for testing purposes
		return
	}

	if inv == nil{

		t.Errorf("Failed to create invoice %v", error)
		// Delete tenant
		imiqasho.Delete(ten) // May cause a test time error. But its unimportant for testing purposes
		return
	}

	t.Log("The invoice id is ", id_inv)

	// Delete invoice
	//imiqasho.Delete(*inv)
	// Delete tenant
	//imiqasho.Delete(ten)
}


func TestDoMonthlyInvoiceRun(t *testing.T) {

	p := "June-2017"

	result, message, err := imiqasho.DoMonthlyInvoiceRun(p)

	if err != nil {
		t.Error("Failed to create invoices")
	}

	if result != "success" {

		t.Error(message)
	}
}