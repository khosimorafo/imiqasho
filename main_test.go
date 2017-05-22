package imiqasho_test

import (
	"testing"
	"os"
	"github.com/khosimorafo/imiqasho"

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

	tenant := imiqasho.Tenant{MoveInDate:"2017-05-01", Gender:"Male",  FirstName: "ProRata", Surname:"Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}
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

	id1, _ := v.GetString("name")
	gender, _ := v.GetString("gender")
	site, _ := v.GetString("site")
	mobile, _ := v.GetString("mobile")


	t.Log("The new name is", id1)
	t.Log("The new gender is", gender)
	t.Log("The new site is", site)
	t.Log("The new mobile is", mobile)


	ten := imiqasho.Tenant{ID:id, Name:"M Tenant - New Name", Gender:"Female"}

	res, ent, err := imiqasho.Update(ten)

	if res != "success" {

		t.Errorf("Failed to update tenant. Result = %v", result)
	}

	if err != nil {

		t.Errorf("Failed to update tenant %v", error)
	}

	b1, _ := json.Marshal(ent)
	v1, _ := jason.NewObjectFromBytes(b1)

	id_u, _ := v1.GetString("name")
	gender_u, _ := v1.GetString("gender")
	site_u, _ := v1.GetString("site")
	mobile_u, _ := v1.GetString("mobile")


	t.Log("The update name is", id_u)
	t.Log("The update gender is", gender_u)
	t.Log("The update site is", site_u)
	t.Log("The new mobile is", mobile_u)


	// Remove created tenant
	ten1 := entity
	res1, err1 := imiqasho.Delete(*ten1)

	if res1 != "success" {

		t.Errorf("Failed to delete tenant. Result = %v", res1)
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
	tenant := imiqasho.Tenant{CreateProRataInvoice:true, MoveInDate:"2017-05-13", FirstName: "ProRata",
		Surname:"Tenant", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}

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

/*
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

	var invoice imiqasho.Invoice
	invoice.ID = id_inv
	pay_result, payment, err := invoice.MakePayment(pay)//ten.CreatePayment(pay)

	if err != nil {

		t.Errorf("Failed with error >> ", err)
	}

	if pay_result != "success" {

		t.Errorf("Failed to make payment!", err)
	}

	t.Log(*payment)


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



func TestCreateTenantWITHFirstInvoice(t *testing.T)  {

	// Create tenant.
	tenant := imiqasho.Tenant{ MoveInDate:"2017-05-13", FirstName: "Kholel", CreateProRataInvoice:false, LastManualPeriod:"February-2017",
		Surname:"Futshnane", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}

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

	_, invoices , err := ten.GetInvoices(map[string]string{})

	if err != nil {

		t.Errorf("Expected an invoice. No invoice found. %v", err)
	}

	if(len(*invoices) < 1){

		t.Errorf("Expected a single invoice. Found %v", len(*invoices))
	} else {

		t.Log("The number of invoices found is : ", len(*invoices))

		for _, invoice := range *invoices{

			imiqasho.Delete(invoice)
		}
	}

	// Delete tenant
	imiqasho.Delete(ten)
}


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

func TestTenantRead(t *testing.T) {

	ten := imiqasho.Tenant{ID: "256831000000249281"}

	_, entity, error := imiqasho.Read(ten)

	if error != nil {

		t.Errorf("Server error")
	}

	if entity == nil {

		t.Errorf("No tenant found with id. %v", ten.ID)
		return
	}

	b, _ := json.Marshal(entity)
	p, _ := jason.NewObjectFromBytes(b)
	id_ten, _ := p.GetString("id")

	t.Log(p)

	if id_ten != ten.ID {

		t.Errorf("Expected tenant id of %v got %v", ten.ID, id_ten)
	}
}

func TestInvoiceRead(t *testing.T) {

	inv := imiqasho.Invoice{ID:"256831000000048033"}

	_, entity, error := imiqasho.Read(inv)

	if error != nil {

		t.Errorf("Server error")
	}

	if entity == nil {

		t.Errorf("No invoice found with id. %v", inv.ID)
		return
	}

	b, _ := json.Marshal(entity)
	p, _ := jason.NewObjectFromBytes(b)
	id_inv, _ := p.GetString("id")

	if id_inv != inv.ID {

		t.Errorf("Expected invoice id of %v got %v", inv.ID, id_inv)
	}
}

func TestPaymentRead(t *testing.T) {

	pay := imiqasho.Payment{ID:"256831000000048057"}

	_, entity, error := imiqasho.Read(pay)

	if error != nil {

		t.Errorf("Server error")
	}

	if entity == nil {

		t.Errorf("No payment found with id. %v", pay.ID)
		return
	}

	b, _ := json.Marshal(entity)
	p, _ := jason.NewObjectFromBytes(b)
	id_pay, _ := p.GetString("id")

	t.Log(p)

	if id_pay != pay.ID {

		t.Errorf("Expected payment id of %v got %v", pay.ID, id_pay)
	}
}
*/

/*
func TestCreateTenantWithLastManualPeriod(t *testing.T)  {

	// Create tenant.
	tenant := imiqasho.Tenant{ MoveInDate:"2017-05-13", FirstName: "Kholel", CreateProRataInvoice:false, LastManualPeriod:"February-2017",
		Surname:"Futshnane", Mobile: "0832345678", ZAID: "2222222222222", Site: "Mganka", Room: "3"}

	//var i imiqasho.EntityInterface
	//i = tenant
	_, entity, _ := tenant.CreateTenant()//imiqasho.Create(i)

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

	_, invoices , err := ten.GetInvoices(map[string]string{})

	if err != nil {

		t.Errorf("Expected an invoice. No invoice found. %v", err)
	}

	if(len(*invoices) < 1){

		t.Errorf("Expected a single invoice. Found %v", len(*invoices))
	} else {

		t.Log("The number of invoices found is : ", len(*invoices))

		//for _, invoice := range *invoices{
		//
		//	imiqasho.Delete(invoice)
		//}
	}

	// Delete tenant
	//imiqasho.Delete(ten)
}


*/

func TestDoMonthlyLatePaymentFines(t *testing.T) {

	i,j, err := imiqasho.DoMonthlyLatePaymentFines("May-2017")

	if err != nil{

		t.Errorf("Failure!")
	}

	t.Log("error strings %s", err)
	t.Log("no if invoice %", i)
	t.Log("no if success %", j)



}

/*
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
	//amount, _ := v_inv.GetFloat64("balance")


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

	var p imiqasho.PaymentExtension

	p.InvoiceID = id_inv
	p.PayByDate = "2017-05-23"

	_, err_ext := invoice.MakePaymentExtensionRequest(p)

	if err_ext != nil {

		t.Errorf("Extension request failure : %v", err_ext)
	}


	// Delete invoice
	imiqasho.Delete(*inv)
	// Delete tenant
	imiqasho.Delete(ten)
}

*/