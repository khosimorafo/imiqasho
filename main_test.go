package imiqasho_test

import (
	"testing"
	"os"
	"github.com/khosimorafo/imiqasho"
	"encoding/json"
	"github.com/antonholmquist/jason"
)

var a imiqasho.App

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

		t.Errorf("Expected records but list is empty! d\n")
	}

	// Delete created tenant
	ten := entity
	imiqasho.Delete(*ten)
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
*/

/*
func TestCreateTenantFirstInvoice(t *testing.T)  {

	// Create tenant.
	tenant := imiqasho.Tenant{MoveInDate:"2017-05-13", FirstName: "ProRata", Surname:"Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
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

	t.Log(v_inv)

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

	result, invoices , err := ten.GetInvoices(map[string]string{})

	if err != nil {

		t.Errorf("Expected an invoice. No invoice found. %v", error)
	}

	if(len(*invoices) < 1){

		t.Errorf("Expected a single invoice. Found %v", len(*invoices))
	} else {

		t.Log("The number of invoices found is : ", len(*invoices))

		for _, invoice := range *invoices{

			// Must return an error
			_, _, error := tenant.CreateTenantInvoice(invoice.PeriodName)

			if error == nil{
				t.Errorf("Invoice created on an previously used period")
			} else {

				t.Log("Succesfully stopped duplicate invoice creation : ", error)
			}
		}
	}

	// Delete invoice
	imiqasho.Delete(*inv)
	// Delete tenant
	imiqasho.Delete(ten)
}
*/

/*
func TestCreateTenantNextInvoice(t *testing.T)  {

	// Create tenant.
	tenant := imiqasho.Tenant{MoveInDate:"2017-05-13", FirstName: "ProRata", Surname:"Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
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

	//t.Log(v_inv)

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

	t.Log("The first invoice id is ", id_inv)

	// Create second tenant invoice.

	result_nxt, inv_nxt, error_nxt := ten.CreateNextTenantInvoice()

	b_inv_nxt, _ := json.Marshal(inv_nxt)
	v_inv_nxt, _ := jason.NewObjectFromBytes(b_inv_nxt)
	id_inv_nxt, _ := v_inv_nxt.GetString("id")

	//t.Log(v_inv_nxt)

	if result != "success" {

		t.Errorf("Failed to create next invoice. Result = %v", result_nxt)
		// Delete tenant
		imiqasho.Delete(ten)
		return
	}

	if error != nil {

		t.Errorf("Failed to create next invoice %v", error_nxt)
		// Delete tenant
		imiqasho.Delete(ten) // May cause a test time error. But its unimportant for testing purposes
		return
	}

	if inv == nil{

		t.Errorf("Failed to create next invoice %v", error_nxt)
		// Delete tenant
		imiqasho.Delete(ten) // May cause a test time error. But its unimportant for testing purposes
		return
	}

	t.Log("The next invoice id is ", id_inv_nxt)

	filters := make(map[string]string)
	//filters["customer_id"] = ten.ID
	result, invoices , err := ten.GetInvoices(filters)

	if err != nil {

		t.Errorf("Expected invoice(s). No invoice found. %v", err)
	}

	if(len(*invoices) < 1){

		t.Errorf("Expected a single invoice. Found %v", len(*invoices))
	} else {

		t.Log("The number of invoices found is : ", len(*invoices))


		for _, invoice := range *invoices{

			// Must return an error
			_, _, error := ten.CreateTenantInvoice(invoice.PeriodName)

			if error == nil{

				t.Errorf("Invoice created on an previously used period")
			} else {

				t.Log("Succesfully stopped duplicate invoice creation : ", error)
			}
		}
	}

	// Delete next invoice
	imiqasho.Delete(*inv_nxt)
	// Delete first invoice
	imiqasho.Delete(*inv)
	// Delete tenant
	imiqasho.Delete(ten)
}

func TestMakePaymentExtensionRequestAndPay(t *testing.T)  {

	// Create tenant.
	tenant := imiqasho.Tenant{MoveInDate:"2017-05-13", FirstName: "ProRata", Surname:"Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
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
	//due_date, _ := v_inv.GetString("due_date")
	amount, _ := v_inv.GetFloat64("balance")


	t.Log(v_inv)

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


	//*********Make payment extension request***************

	invoice := imiqasho.Invoice{ID:id_inv}

	_, err_ext := invoice.MakePaymentExtensionRequest("2017-05-23")

	if err_ext != nil {

		t.Errorf("Extension request failure : %v", err_ext)
	}

	pay := imiqasho.PaymentPayload{InvoiceID:id_inv, PaymentAmount:amount, PaymentDate:"2017-06-23",PaymentMode:"Cash"}

	pay_result, payment, err := invoice.MakePayment(pay)

	if err != nil {

		t.Errorf("Failed with error >> ", err)
	}

	if pay_result != "success" {

		t.Errorf("Failed to make payment!", err)
	}

	// Delete payment
	imiqasho.Delete(*payment)
	// Delete invoice
	imiqasho.Delete(*inv)
	// Delete tenant
	imiqasho.Delete(ten)
}


func TestInvoiceAddLineItems(t *testing.T)  {

	// Create tenant.
	tenant := imiqasho.Tenant{MoveInDate:"2017-05-13", FirstName: "ProRata", Surname:"Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
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
	id_cust, _ := v.GetString("id")
	in_date, _ := v.GetString("move_in_date")


	ten := imiqasho.Tenant{ID: id_cust, MoveInDate: in_date}

	result, inv, error := ten.CreateFirstTenantInvoice()

	b_inv, _ := json.Marshal(inv)
	v_inv, _ := jason.NewObjectFromBytes(b_inv)
	id_inv, _ := v_inv.GetString("id")

	//t.Log(v_inv)

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

	t.Log("The first invoice id_cust is ", id_inv)


	item := imiqasho.GetRentalFineLineItem()

	//define items slice
	line_items := make([]imiqasho.LineItem, 0)
	line_items = append(line_items, item)

	invoice := imiqasho.Invoice{ID: id_inv, CustomerID: id_cust}

	_, update_invoice, error_upd := invoice.AddLineItems(line_items)

	if error_upd != nil {

		t.Errorf(error_upd.Error())
	}

	upd_inv, _ := json.Marshal(update_invoice)
	v_upd_inv, _ := jason.NewObjectFromBytes(upd_inv)
	upd_inv_line_items, _ := v_upd_inv.GetObjectArray("line_items")

	if len(upd_inv_line_items) != 2 {

		t.Errorf("Less than expected number of line items.")
	}

	// Delete first invoice
	imiqasho.Delete(*inv)
	// Delete tenant
	imiqasho.Delete(ten)
}*/


func TestCreateInvoiceAndMakePayment(t *testing.T)  {

	// Create tenant.
	tenant := imiqasho.Tenant{MoveInDate:"2017-05-01", FirstName: "ProRata", Surname:"Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
	var i imiqasho.EntityInterface
	i = tenant
	_, entity, _ := imiqasho.Create(i)

	if entity == nil {

		t.Errorf("Failed to create tenant. Entity is empty!")
		return
	}


	b, _ := json.Marshal(entity)
	v, _ := jason.NewObjectFromBytes(b)
	id, _ := v.GetString("id")
	in_date, _ := v.GetString("move_in_date")

	ten := imiqasho.Tenant{ID:id, MoveInDate:in_date}

	// Create first tenant invoice.
	result, inv, error := ten.CreateFirstTenantInvoice()

	b_inv, _ := json.Marshal(inv)
	v_inv, _ := jason.NewObjectFromBytes(b_inv)
	id_inv, _ := v_inv.GetString("id")

	if result != "success" {

		t.Errorf("Failed to create invoice. Result = %v", result)
	}

	if error != nil {

		t.Errorf("Failed to create invoice %v", error)
		return
	}

	if inv == nil{

		t.Errorf("Failed to create invoice %v", error)
		return
	}

	pay := imiqasho.PaymentPayload{InvoiceID:id_inv, PaymentAmount:300.0, PaymentDate:"2017-04-25",PaymentMode:"Cash"}

	pay_result, payment, err := ten.CreatePayment(pay)

	if err != nil {

		t.Errorf("Failed with error >> ", err)
	}

	if pay_result != "success" {

		t.Errorf("Failed to make payment!", err)
	}

	_, payments, err_pay := ten.GetPayments(map[string]string{})

	if err_pay != nil {

		t.Errorf("Failed with error >> ", err_pay)
	}

	if len(*payments) != 1 {

		t.Errorf("Expected 1 payment record, received ", len(*payments))
	}

	// Delete payment
	imiqasho.Delete(*payment)
	// Delete invoice
	imiqasho.Delete(*inv)
	// Delete tenant
	imiqasho.Delete(ten)
}

/*
func TestDiscountInvoice(t *testing.T)  {

	// Create tenant.
	tenant := imiqasho.Tenant{MoveInDate:"2017-05-01", FirstName: "ProRata", Surname:"Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
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

	//t.Log(v_inv)

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

	//t.Log("The invoice id is ", id_inv)

	pay := imiqasho.PaymentPayload{InvoiceID:id_inv, PaymentAmount:300, PaymentDate:"2017-04-23",PaymentMode:"Cash"}

	pay_result, payment, err := ten.CreatePayment(pay)

	if err != nil {

		t.Errorf("Failed with error >> ", err)
	}

	if pay_result != "success" {

		t.Errorf("Failed to make payment!", err)
		return
	}

	inv_ := imiqasho.Invoice{ID:id_inv}

	_, discounted_invoice, _ := inv_.Read()
	disc_inv, _ := json.Marshal(discounted_invoice)
	disc_v_inv, _ := jason.NewObjectFromBytes(disc_inv)
	discounted_invoice_status, _ := disc_v_inv.GetString("status")

	if discounted_invoice_status != "paid"{

		t.Errorf("Expected invoice status paid. Got '%v'", discounted_invoice_status)
	}

	// Delete payment
	imiqasho.Delete(*payment)
	// Delete invoice
	imiqasho.Delete(*inv)
	// Delete tenant
	imiqasho.Delete(ten)
}
*/