package imiqasho

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/antonholmquist/jason"
	"github.com/smallnest/goreq"
	"github.com/khosimorafo/imiqashoserver"
)

type App struct {}

func (a *App) Initialize() {
}

type EntityInterface interface {

	Create() (string, *EntityInterface, error)
	Read() (string, *EntityInterface, error)
	Update() (string, *EntityInterface, error)
	Delete() (string, error)
}

func Create(i EntityInterface) (string, *EntityInterface, error) {

	result, message, _ := i.Create()
	return result, message, nil
}

func Read(i EntityInterface) (string, *EntityInterface, error) {

	result, message, _ := i.Read()
	return result, message, nil
}

func Update(i EntityInterface) (string, *EntityInterface, error) {

	result, message, _ := i.Update()
	return result, message, nil
}

func Delete(i EntityInterface) (string, error) {

	result, err := i.Delete()
	return result, err
}

//****************************Tenants

type TenantInterface interface {

	CreateFirstTenantInvoice() (string, *EntityInterface, error)
	CreateInvoice() (string, *EntityInterface, error)
}

func CreateFirstTenantInvoice(t TenantInterface) (string, *EntityInterface, error){

	result, message, _ := t.CreateFirstTenantInvoice()
	return result, message, nil
}

func CreateInvoice(t TenantInterface) (string, *EntityInterface, error){

	result, message, _ := t.CreateInvoice()
	return result, message, nil
}

type TenantZoho struct {
	ID           string        `json:"contact_id,omitempty"`
	Name         string        `json:"contact_name,omitempty"`
	Telephone    string        `json:"telephone,omitempty"`
	Fax          string        `json:"fax,omitempty"`
	Mobile       string        `json:"mobile,omitempty"`
	Status       string        `json:"status,omitempty"`
	CustomFields []CustomField `json:"custom_fields,omitempty"`
}

type Tenant struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	ZAID        string  `json:"zaid"`
	Telephone   string  `json:"telephone"`
	Fax         string  `json:"fax"`
	Mobile      string  `json:"mobile"`
	Site        string  `json:"site"`
	Room        string  `json:"room"`
	MoveInDate  string  `json:"move_in_date"`
	MoveOutDate string  `json:"move_out_date"`
	Outstanding float64 `json:"outstanding"`
	Credits     float64 `json:"credit_available"`
	Status      string  `json:"status"`

}

//A method to create new tenant

func (tenant Tenant) Create() (string, *EntityInterface, error) {

	fmt.Printf("Creating tenant - %s with zar_id %s\n", tenant.Name, tenant.ZAID)

	cfs := make([]CustomField, 0)

	cfs = append(cfs, CustomField{Index: 4, Value: tenant.ZAID})
	cfs = append(cfs, CustomField{Index: 5, Value: tenant.Site})
	cfs = append(cfs, CustomField{Index: 6, Value: tenant.Room})
	cfs = append(cfs, CustomField{Index: 7, Value: tenant.MoveInDate})
	cfs = append(cfs, CustomField{Index: 8, Value: tenant.MoveOutDate})

	tenant_zoho := TenantZoho{ID: tenant.ID, Name: tenant.Name, Mobile: tenant.Mobile, Fax: tenant.Fax,
		Telephone: tenant.Telephone, CustomFields: cfs}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(tenant_zoho)

	resp, _, err := goreq.New().
		Post(postUrl("contacts")).
		SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SendRawString("JSONString=" + b.String()).End()

	result, entity, error := TenantResult(resp, err)

	return result, entity, error
}

func (tenant Tenant) Read() (string, *EntityInterface, error) {

	fmt.Printf("Retrieving tenant - %s \n", tenant.ID)

	resp, _, err := goreq.New().Get(readUrl("contacts", tenant.ID)).End()

	result, entity, error := TenantResult(resp, err)

	return result, entity, error
}

func (tenant Tenant) Update() (string, *EntityInterface, error) {

	fmt.Printf("Updating tenant - %s\n", tenant.ID)

	cfs := make([]CustomField, 0)

	cfs = append(cfs, CustomField{Index: 4, Value: tenant.ZAID})
	cfs = append(cfs, CustomField{Index: 5, Value: tenant.Site})
	cfs = append(cfs, CustomField{Index: 6, Value: tenant.Room})

	tenant_zoho := TenantZoho{ID: tenant.ID, Name: tenant.Name}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(tenant_zoho)

	resp, _, err := goreq.New().
		Put(putUrl("contacts", tenant.ID)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SendRawString("JSONString=" + b.String()).End()

	result, entity, error := TenantResult(resp, err)

	return result, entity, error
}

func (tenant Tenant) Delete() (string, error) {

	resp, _, err := goreq.New().Delete(deleteUrl("contacts", tenant.ID)).End()

	result, _ := jason.NewObjectFromReader(resp.Body)

	if err != nil {

		return "failure", errors.New("Failed to delete tenant. Api http error")
	} else {

		code, _ := result.GetInt64("code")

		if code == 0 {

			return "success", nil
		} else {

			fmt.Print(result)
			return "failure", errors.New("Failed to delete tenant. Api interface error")
		}
	}
}

func (tenant Tenant) CreateFirstTenantInvoice() (string, *EntityInterface, error) {

	layout := "2006-01-02"


	fmt.Printf("Move in date : %v",tenant.MoveInDate)
	//ti := "2017-04-12"
	t, err := time.Parse(layout, tenant.MoveInDate)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Actual Date : %v",t)


	p := imiqashoserver.P{t}

	period, err1 := p.GetPeriod()

	//fmt.Printf(period)

	if err1 != nil {

		return "failure", nil, err1
	}else {

		line_item := GetRentalLineItem()
		// Set pro rata item amount
		pr, _ := p.GetProRataDays()
		line_item.Rate = line_item.Rate * pr

		_, entity, error := tenant.CreateInvoice(period, line_item)

		return "success", entity, error
	}


}

func (tenant Tenant) CreateTenantInvoice() (string, *EntityInterface, error) {

	//1. Check tenant invoices and get the highest index among created invoices.




	//2. Query the next index to get the next period.

	//3. Create invoice based on the result of point 2.

	//4. Return invoice create at point 3.

	return "", nil, nil

}

func (tenant Tenant) CreateInvoice(period imiqashoserver.Period, item LineItem) (string, *EntityInterface, error) {

	date, due := generateInvoiceDates(period.Start)
	//line_item := GetRentalLineItem()

	//Set invoice number
	length := len(tenant.ID) - 6

	var reference bytes.Buffer
	reference.WriteString(period.Name)
	reference.WriteString("-")
	reference.WriteString(tenant.ID[length:])

	item.Description = item.Description + "  " + period.Name

	//define items slice
	line_items := make([]LineItem, 0)
	line_items = append(line_items, item)

	//Must remove this hack
	var index int64
	index = int64(period.Index)

	invoice := Invoice{CustomerID: tenant.ID, InvoiceDate: date, DueDate: due,LineItems:line_items,
		ReferenceNumber: reference.String(), PeriodIndex: index, PeriodName:period.Name}

	var i EntityInterface
	i = invoice
	result, entity, error := Create(i)

	return result, entity, error
}

func (tenant Tenant) GetInvoices(filters map[string]string) (string, *[]Invoice, error) {

	filters["customer_id"] = tenant.ID

	resp, _, _ := goreq.New().Get(listsUrl("invoices", filters)).End()

	result, error := jason.NewObjectFromReader(resp.Body)

	//define slice
	invoices := make([]Invoice, 0)

	if error != nil {

		return "failure", &invoices, errors.New("Invoice query failure. Api http error")
	} else {

		message, _ := result.GetString("message")

		if message == "success" {

			invs, _ := result.GetObjectArray("invoices")
			for _, inv := range invs {

				invoice_id, _ := inv.GetString("invoice_id")
				customer_id, _ := inv.GetString("customer_id")
				customer_name, _ := inv.GetString("customer_name")
				reference, _ := inv.GetString("reference_number")
				due_date, _ := inv.GetString("due_date")
				invoice_date, _ := inv.GetString("date")
				balance, _ := inv.GetFloat64("balance")
				total, _ := inv.GetFloat64("total")


				invoice := Invoice{ID: invoice_id, CustomerID: customer_id, CustomerName:customer_name,
					ReferenceNumber: reference, DueDate: due_date, InvoiceDate: invoice_date,
					Balance:balance, Total:total}

				invoices = append(invoices, invoice)
			}

			return "success", &invoices, nil
		}

		return "failure", &invoices, errors.New("Invoice query failure. Api http error")
	}
}

func (tenant Tenant) GetPayments(filters map[string]string) (string, *[]Payment, error) {

	filters["customer_id"] = tenant.ID

	resp, _, _ := goreq.New().Get(listsUrl("customerpayments", filters)).End()

	result, error := jason.NewObjectFromReader(resp.Body)

	//define slice
	payments := make([]Payment, 0)

	if error != nil {

		return "failure", &payments, errors.New("Payment query failure. Api http error")
	} else {

		message, _ := result.GetString("message")

		if message == "success" {

			pymnts, _ := result.GetObjectArray("payments")
			for _, pymnt := range pymnts {

				payment_id, _ := pymnt.GetString("payment_id")
				payment_number, _ := pymnt.GetString("payment_number")
				invoice_number, _ := pymnt.GetString("invoice_number")
				mode, _ := pymnt.GetString("payment_mode")
				customer_id := tenant.ID
				amount, _ := pymnt.GetFloat64("amount")
				date, _ := pymnt.GetString("date")

				payment := Payment{ID: payment_id, CustomerID: customer_id, PaymentAmount: amount,
					PaymentDate: date, PaymentMode: mode, PaymentNumber: payment_number,
					InvoiceNumber: invoice_number}

				payments = append(payments, payment)
			}

			return "success", &payments, nil
		}

		return "failure", &payments, errors.New("Payment query failure. Api http error")
	}
}

func GetTenants(filters map[string]string) (string, *[]Tenant, error) {

	//filters := map[string]string{}

	resp, _, _ := goreq.New().Get(listsUrl("contacts", filters)).End()

	result, error := jason.NewObjectFromReader(resp.Body)

	//define slice
	tenants := make([]Tenant, 0)

	if error != nil {

		return "failure", &tenants, errors.New("Tenant query failure. Api http error")
	} else {

		message, _ := result.GetString("message")

		if message == "success" {

			contacts, _ := result.GetObjectArray("contacts")
			for _, contact := range contacts {

				customer_id, _ := contact.GetString("contact_id")
				name, _ := contact.GetString("contact_name")
				zaid, _ := contact.GetString("cf_zar_id_no")
				telephone, _ := contact.GetString("telephone")
				mobile, _ := contact.GetString("mobile")

				site, _ := contact.GetString("cf_site")
				room, _ := contact.GetString("cf_room")

				outstanding, _ := contact.GetFloat64("outstanding_receivable_amount")
				credit_available, _ := contact.GetFloat64("unused_credits_receivable_amount")
				status, _ := contact.GetString("status")

				tenant := Tenant{ID: customer_id, Name: name, ZAID: zaid, Telephone: telephone, Mobile: mobile,
					Site: site, Room: room, Status: status, Outstanding: outstanding, Credits: credit_available}
				tenants = append(tenants, tenant)
			}

			return "success", &tenants, nil
		} else {

			return "failure", &tenants, errors.New("Tenant query failure. Api interface error")
		}
	}
}

func TenantResult(response goreq.Response, err []error) (string, *EntityInterface, error) {

	if err != nil {

		return "failure", nil, errors.New("Tenant operation failure. Api http error")
	} else {

		result, _ := jason.NewObjectFromReader(response.Body)

		code, _ := result.GetInt64("code")

		if code == 0 {

			contact, _ := result.GetObject("contact")

			customer_id, _ := contact.GetString("contact_id")
			name, _ := contact.GetString("contact_name")
			zaid, _ := contact.GetString("cf_zar_id_no")
			telephone, _ := contact.GetString("telephone")
			mobile, _ := contact.GetString("mobile")

			site, _ := contact.GetString("cf_site")
			room, _ := contact.GetString("cf_room")

			in_date, _ := contact.GetString("cf_moveindate")
			out_date, _ := contact.GetString("cf_moveoutdate")



			outstanding, _ := contact.GetFloat64("outstanding_receivable_amount")
			credit_available, _ := contact.GetFloat64("unused_credits_receivable_amount")
			status, _ := contact.GetString("status")

			tenant := Tenant{ID: customer_id, Name: name, ZAID: zaid, Telephone: telephone, Mobile: mobile,
				Site: site, Room: room, Status: status, Outstanding: outstanding,
				Credits: credit_available, MoveInDate:in_date, MoveOutDate:out_date}

			var i EntityInterface
			i = tenant
			return "success", &i, nil
		} else {

			fmt.Print(result)
			return "failure", nil, errors.New("Tenant operation failure. Api interface error")
		}
	}
}

//****************************Invoices

type InvoiceZoho struct {
	ID              string     `json:"invoice_id,omitempty"`
	CustomerID      string     `json:"customer_id"`
	ReferenceNumber string     `json:"reference_number"`
	InvoiceDate     string     `json:"date"`
	DueDate         string     `json:"due_date"`
	LineItems       []LineItem `json:"line_items"`
	CustomFields []CustomField `json:"custom_fields,omitempty"`
}

type Invoice struct {
	ID              string     	`json:"id,omitempty"`
	CustomerID      string     	`json:"customer_id"`
	CustomerName      string     	`json:"customer_name"`
	InvoiceNumber  	string     	`json:"invoice_number"`
	ReferenceNumber string     	`json:"reference_number"`
	Total		float64		`json:"total"`
	Balance		float64		`json:"balance"`
	InvoiceDate     string     	`json:"date"`
	DueDate         string     	`json:"due_date"`
	LineItems       []LineItem 	`json:"line_items,omitempty"`
	PeriodIndex	int64		`json:"period_index,omitempty"`
	PeriodName	string 		`json:"period_name,omitempty"`
	Status          string		`json:"status,omitempty"`
}

func (invoice Invoice) Create() (string, *EntityInterface, error) {

	fmt.Printf("Creating invoice for customer %s\n", invoice.CustomerID)

	cfs := make([]CustomField, 0)

	cfs = append(cfs, CustomField{Index: 2, Value: invoice.PeriodIndex})
	cfs = append(cfs, CustomField{Index: 3, Value: invoice.PeriodName})


	invoice_zoho := InvoiceZoho{ID: invoice.ID, CustomerID:invoice.CustomerID, ReferenceNumber:invoice.ReferenceNumber,
		InvoiceDate:invoice.InvoiceDate, DueDate:invoice.DueDate, LineItems:invoice.LineItems, CustomFields: cfs}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(invoice_zoho)


	resp, _, err := goreq.New().
		Post(postUrl("invoices")).
		SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SendRawString("JSONString=" + b.String()).End()


	result, entity, error := InvoiceResult(resp, err)

	return result, entity, error
}

func (invoice Invoice) Read() (string, *EntityInterface, error) {

	fmt.Printf("Retrieving invoice - %s \n", invoice.ID)

	resp, _, err := goreq.New().Get(readUrl("invoice", invoice.ID)).End()

	result, entity, error := InvoiceResult(resp, err)

	return result, entity, error
}

func (invoice Invoice) Update() (string, *EntityInterface, error) {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(invoice)

	fmt.Println(b)

	resp, _, err := goreq.New().
		Put(putUrl("invoice", invoice.ID)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SendRawString("JSONString=" + b.String()).End()

	result, entity, error := InvoiceResult(resp, err)

	return result, entity, error
}

func (invoice Invoice) Delete() (string, error) {

	resp, _, err := goreq.New().Delete(deleteUrl("invoices", invoice.ID)).End()

	result, _ := jason.NewObjectFromReader(resp.Body)

	if err != nil {

		return "failure", errors.New("Failed to delete invoice. Api http error")
	} else {

		code, _ := result.GetInt64("code")

		if code == 0 {

			return "success", nil
		} else {

			fmt.Print(result)
			return "failure", errors.New("Failed to delete invoice. Api interface error")
		}
	}
}

func GetInvoices(filters map[string]string) (string, *[]Invoice, error) {

	resp, _, _ := goreq.New().Get(listsUrl("invoices", filters)).End()

	result, error := jason.NewObjectFromReader(resp.Body)

	//define slice
	invoices := make([]Invoice, 0)

	if error != nil {

		return "failure", &invoices, errors.New("Invoice query failure. Api http error")
	} else {

		message, _ := result.GetString("message")

		if message == "success" {

			invs, _ := result.GetObjectArray("invoices")

			for _, inv := range invs {

				customer_id, _ := inv.GetString("customer_id")
				invoice_id, _ := inv.GetString("invoice_id")
				invoice_number, _ := inv.GetString("invoice_number")
				reference_number, _ := inv.GetString("reference_number")
				invoice_date, _ := inv.GetString("date")
				due_date, _ := inv.GetString("due_date")

				p_index, _ := inv.GetInt64("cf_periodindex")
				p_name, _ := inv.GetString("cf_periodname")

				total, _ := inv.GetFloat64("total")
				balance, _ := inv.GetFloat64("balance")

				status, _ := inv.GetString("status")

				invoice := Invoice{CustomerID: customer_id, ID:invoice_id, InvoiceNumber:invoice_number, Status: status,
					ReferenceNumber:reference_number,InvoiceDate:invoice_date, DueDate:due_date, PeriodIndex:p_index,
					PeriodName:p_name, Total:total, Balance:balance}

				invoices = append(invoices, invoice)
			}

			return "success", &invoices, nil
		} else {

			return "failure", &invoices, errors.New("Invoice query failure. Api interface error")
		}
	}
}

func InvoiceResult(response goreq.Response, err []error) (string, *EntityInterface, error) {

	if err != nil {

		return "failure", nil, errors.New("Invoice operation failure. Api http error")
	} else {

		result, _ := jason.NewObjectFromReader(response.Body)

		code, _ := result.GetInt64("code")

		//fmt.Printf(result.String())

		if code == 0 {

			inv, _ := result.GetObject("invoice")

			invoice_id, _ := inv.GetString("invoice_id")
			customer_id, _ := inv.GetString("customer_id")
			customer_name, _ := inv.GetString("customer_name")
			reference, _ := inv.GetString("reference_number")
			due_date, _ := inv.GetString("due_date")
			invoice_date, _ := inv.GetString("date")
			balance, _ := inv.GetFloat64("balance")
			total, _ := inv.GetFloat64("total")

			line_items, _ := inv.GetObjectArray("line_items")

			period_index, _ := inv.GetInt64("cf_periodindex")
			period_name, _ := inv.GetString("cf_periodname")


			items := make([]LineItem, 0)

			for _, item := range line_items {

				id, _ := item.GetString("item_id")
				name, _ := item.GetString("name")
				description, _ := item.GetString("description")
				rate, _ := item.GetFloat64("rate")
				quantity, _ := item.GetInt64("quantity")

				i := LineItem{ItemID: id, Name: name, Description: description, Rate: rate, Quantity: quantity}
				items = append(items, i)
			}

			invoice := Invoice{ID: invoice_id, CustomerID: customer_id, CustomerName:customer_name,
				ReferenceNumber: reference, DueDate:       due_date, InvoiceDate: invoice_date,
				Balance:balance, Total:total, LineItems: items, PeriodIndex: period_index,
				PeriodName: period_name}

			var i EntityInterface
			i = invoice

			return "success", &i, nil
		} else {
			fmt.Printf("error bottom")

			return "failure", nil, errors.New("Invoice operation failure. Api interface error")
		}
	}
}

//****************************Payment

type PaymentZoho struct {
	ID            string       `json:"id,omitempty"`
	CustomerID    string       `json:"customer_id"`
	InvoiceID     string       `json:"invoice_id"`
	PaymentAmount float64      `json:"amount"`
	PaymentMode   string       `json:"payment_mode"`
	Description   string       `json:"description"`
	Invoices      []PayInvoice `json:"invoices"`
}

type Payment struct {
	ID            string  `json:"id,omitempty"`
	CustomerID    string  `json:"customer_id"`
	InvoiceID     string  `json:"invoice_id,omitempty"`
	InvoiceNumber string  `json:"invoice_number"`
	PaymentNumber string  `json:"payment_number"`
	PaymentAmount float64 `json:"amount"`
	Balance       float64 `json:"balance,omitempty"`
	PaymentMode   string  `json:"payment_mode"`
	PaymentDate   string  `json:"payment_date"`
	Status        string  `json:"status,omitempty"`
	Description   string  `json:"description,omitempty"`
	CustomerName  string  `json:"customer_name,omitempty"`
}

type PayInvoice struct {
	InvoiceID     string  `json:"invoice_id"`
	AppliedAmount float64 `json:"amount_applied"`
}

func (payment Payment) Create() (string, *EntityInterface, error) {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(payment)

	fmt.Println(b)

	resp, _, err := goreq.New().
		Post(postUrl("invoices")).
		SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SendRawString("JSONString=" + b.String()).End()

	result, entity, error := PaymentResult(resp, err)

	return result, entity, error
}

func (payment Payment) Read() (string, *EntityInterface, error) {

	fmt.Printf("Retrieving payment - %s \n", payment.ID)

	resp, _, err := goreq.New().Get(readUrl("payment", payment.ID)).End()

	result, entity, error := PaymentResult(resp, err)

	return result, entity, error
}

func (payment Payment) Update() (string, *EntityInterface, error) {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(payment)

	fmt.Println(b)

	resp, _, err := goreq.New().
		Put(putUrl("payment", payment.ID)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SendRawString("JSONString=" + b.String()).End()

	result, entity, error := PaymentResult(resp, err)

	return result, entity, error
}

func (payment Payment) Delete() (string, error) {

	resp, _, err := goreq.New().Delete(deleteUrl("payment", payment.ID)).End()

	result, _ := jason.NewObjectFromReader(resp.Body)

	if err != nil {

		return "failure", errors.New("Failed to delete payment. Api http error")
	} else {

		code, _ := result.GetInt64("code")

		if code == 0 {

			return "success", nil
		} else {

			fmt.Print(result)
			return "failure", errors.New("Failed to delete payment. Api interface error")
		}
	}
}

func PaymentResult(response goreq.Response, err []error) (string, *EntityInterface, error) {

	if err != nil {

		return "failure", nil, errors.New("Payment operation failure. Api http error")
	} else {

		result, _ := jason.NewObjectFromReader(response.Body)

		code, _ := result.GetInt64("code")

		if code == 0 {

			record, _ := result.GetObject("payment")

			id, _ := record.GetString("payment_id")
			customer_id, _ := record.GetString("customer_id")
			invoice_id, _ := record.GetString("invoice_id")
			amount, _ := record.GetFloat64("amount")
			date, _ := record.GetString("date")
			mode, _ := record.GetString("payment_mode")
			status, _ := record.GetString("status")
			description, _ := record.GetString("description")
			customer_name, _ := record.GetString("customer_name")

			payment := Payment{ID: id, CustomerID: customer_id, InvoiceID: invoice_id, PaymentAmount: amount,
				PaymentDate: date, PaymentMode: mode, Status: status, Description: description,
				CustomerName: customer_name}

			var i EntityInterface
			i = payment
			return "success", &i, nil
		} else {

			fmt.Print(result)
			return "failure", nil, errors.New("Invoice operation failure. Api interface error")
		}
	}
}


//****************************Item

func GetRentalLineItem() LineItem {

	item_id, rate, _ := getRentalItemID()
	line := LineItem{ItemID: item_id, Rate: rate, Quantity: 1}

	return line
}

func getRentalItemID() (string, float64, error) {

	apiUrl := "https://invoice.zoho.com"
	resource := "/api/v3/items/"
	data := url.Values{}
	data.Set("authtoken", "23d96588d022f48fe2ce16dfd2b69c71")
	data.Add("organization_id", "163411778")
	data.Add("item_name", "Monthly Rental")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	resp, _, _ := goreq.New().Get(urlStr).End()

	//fmt.Println(body)

	result, error := jason.NewObjectFromReader(resp.Body)

	if error != nil {

		return "", 0, error
	} else {

		code, _ := result.GetInt64("code")
		if code == 0 {

			items, _ := result.GetObjectArray("items")
			for _, item := range items {

				id, _ := item.GetString("item_id")
				rate, _ := item.GetFloat64("rate")
				return id, rate, nil
			}
		}

		return "", 0, nil
	}
}

type Item struct {
	Name        string  `json:"name"`
	Description string  `json:"description                                                                  "`
	Rate        float64 `json:"rate"`
}

type LineItem struct {
	ItemID      string  `json:"item_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Rate        float64 `json:"rate,omitempty"`
	Quantity    int64   `json:"quantity,omitempty"`
}

//****************************Common

type CustomField struct {
	Index int64  `json:"index,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

func DoMonthlyInvoiceRun(m string) (string, string, error){

	filters := map[string]string{}

	period, _ := imiqashoserver.GetPeriod(m)

	fmt.Printf(period.Name)

	result, tenants, _ := GetTenants(filters)

	if result == "success"{

		var invoices_succesfully_created int
		invoices_succesfully_created = 0

		line_item := GetRentalLineItem()

		for _, tenant := range *tenants{

			_, _, err := tenant.CreateInvoice(period, line_item)

			if(err == nil) {

				invoices_succesfully_created++
			}
		}

		if len(*tenants) != invoices_succesfully_created{

			return "failure", "Not all tenants were processed", nil

		}else {

			return "success", "All valid tenants processed",  nil
		}

	}

	return "failure", "", nil
}

func generateInvoiceDates(cur string) (string, string) {

	layout := "2006-01-02"


	fmt.Printf("Move in date : %v",cur)
	ti := "2017-05-12"
	t, err := time.Parse(layout, ti)

	if err != nil {
		fmt.Println(err)
	}

	// Derive invoice date
	t.Format(time.RFC3339)
	current := t.Format("2006-01-02")

	// Derive 5th of the next month. Due date
	t2 := t.AddDate(0, 1, 0)
	d := time.Duration(-int(t2.Day())+5) * 24 * time.Hour

	due := t2.Add(d).Format("2006-01-02")

	return current, due
}

func generatePeriod() string {

	var buffer bytes.Buffer

	month := time.Now().AddDate(0, 1, 0).Month().String()
	year := time.Now().Year()

	buffer.WriteString(month)
	buffer.WriteString("-")
	buffer.WriteString(strconv.Itoa(year))

	return buffer.String()
}

func postUrl(entity string) string {

	apiUrl := "https://invoice.zoho.com"
	resource := fmt.Sprintf("/api/v3/%s/", entity)
	data := url.Values{}
	data.Set("authtoken", "23d96588d022f48fe2ce16dfd2b69c71")
	data.Add("organization_id", "163411778")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	return urlStr
}

func putUrl(entity string, id string) string {

	apiUrl := "https://invoice.zoho.com"
	resource := fmt.Sprintf("/api/v3/%s/%s", entity, id)
	data := url.Values{}
	data.Set("authtoken", "23d96588d022f48fe2ce16dfd2b69c71")
	data.Add("organization_id", "163411778")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	return urlStr
}

func readUrl(entity string, id string) string {

	apiUrl := "https://invoice.zoho.com"
	resource := fmt.Sprintf("/api/v3/%s/%s", entity, id)
	data := url.Values{}
	data.Set("authtoken", "23d96588d022f48fe2ce16dfd2b69c71")
	data.Add("organization_id", "163411778")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	return urlStr
}

func deleteUrl(entity string, id string) string {

	apiUrl := "https://invoice.zoho.com"
	resource := fmt.Sprintf("/api/v3/%s/%s", entity, id)
	data := url.Values{}
	data.Set("authtoken", "23d96588d022f48fe2ce16dfd2b69c71")
	data.Add("organization_id", "163411778")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	return urlStr
}

func listsUrl(entity string, filters map[string]string) string {

	apiUrl := "https://invoice.zoho.com"
	resource := fmt.Sprintf("/api/v3/%s/", entity)
	data := url.Values{}
	data.Set("authtoken", "23d96588d022f48fe2ce16dfd2b69c71")
	data.Add("organization_id", "163411778")

	for k, v := range filters {
		data.Add(k, v)
	}

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	return urlStr
}
