package models

//Company #1 wise database
type Company struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   int64  `json:"cid"`
	CompanyName string `json:"company_name,omitempty"`
	Website     string `json:"website,omitempty"`
	Status      int    `json:"status"`
}

//Access #2 login type admin,super,user,etc
type Access struct {
	ID         string `json:"aid"`
	Type       string `json:"type"`
	CompanyID  string `json:"cid"`    //foreign key
	Serial     int64  `json:"serial"` //company wise increase
	AccessName string `json:"access_name"`
	Status     int    `json:"status"`
}

//Account #3 ::cid::account_id
type Account struct {
	//ContactInfo []Contact `json:"contact_info,omitempty"`
	ID          string `json:"aid"`
	Type        string `json:"type"`                //account
	CompanyID   string `json:"cid"`                 //foreign key
	Serial      int64  `json:"serial"`              //company wise increase
	ParentID    string `json:"parent_id,omitempty"` //if any parent
	AccountType string `json:"account_type"`        //vendor,supplier,customer,consumer
	AccountName string `json:"account_name"`        //supplier business name
	Code        string `json:"code"`                //supplier or customer code
	LoginID     string `json:"login_id"`            //foreign key
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	DateOfBirth string `json:"dob,omitempty"`
	Gender      string `json:"gender,omitempty"` //female,male,other
	Mobile      string `json:"mobile"`           //
	Email       string `json:"email"`            //
	CreateDate  string `json:"create_date"`
	UpdateDate  string `json:"update_date,omitempty"`
	Status      int    `json:"status"`
}

//Login #4  all user account login table
type Login struct {
	ID         string `json:"aid"`
	Type       string `json:"type"`
	CompanyID  string `json:"cid"`        //foreign key
	Serial     int64  `json:"serial"`     //company wise increasing
	AccountID  string `json:"account_id"` //foreign key
	AccessID   string `json:"access_id"`  //foreign key
	AccessName string `json:"access_name"`
	UserName   string `json:"username"` //email or mobile as username
	Password   string `json:"passw"`
	CreateDate string `json:"create_date"`
	LastLogin  string `json:"last_login,omitempty"`
	Status     int    `json:"status"`
}

//Contact info for account and user
type Contact struct {
	ID        string `json:"aid"`
	Type      string `json:"type"`   //table_name
	CompanyID string `json:"cid"`    //foreign key
	Serial    int64  `json:"serial"` //companywise increase
	ContactID int64  `json:"contact_id"`
	//mobile,email,pager,phone,website,socialmedia
	ContactType string `json:"contact_type"`
	ContactData string `json:"contact_date"`
	Owner       string `json:"owner"`     ////owner table info, who owns this data
	Parameter   string `json:"parameter"` //docID
	Status      int    `json:"status"`
}

//Address ...
type Address struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   string `json:"cid"` //foreign key
	Serial      int64  `json:"serial"`
	AccountID   string `json:"account_id"` //foreign_key
	AddressID   string `json:"address_id"`
	AddressType string `json:"address_type"` //billing,shipping
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	Zip         string `json:"zip"`
	Status      int    `json:"status"`
}

//ActivityLog ***
type ActivityLog struct {
	ID           string `json:"aid"`
	Type         string `json:"type"`
	CompanyID    string `json:"cid"`    //foreign key
	Serial       int64  `json:"serial"` //company wise increase
	ActivityType string `json:"activity_type"`
	OwnerTable   string `json:"owner_table"` //table_name
	Parameter    string `json:"parameter"`   //key=val or aid
	LogDetails   string `json:"log_details"`
	IPAddress    string `json:"ip_address"`
	LoginID      string `json:"login_id"`      //optional foreign key
	LoginSession string `json:"login_session"` //foreign key
	CreateDate   string `json:"create_date"`
	Status       int    `json:"status"`
}

//DeviceLog tracks user device info (where they login to the system)
type DeviceLog struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   string `json:"cid"`      //foreign key
	Serial      int64  `json:"serial"`   //company wise increase
	LoginID     string `json:"login_id"` //foreign key
	Browser     string `json:"browser"`
	DeviceType  string `json:"device_type"`
	Os          string `json:"os"`
	Platform    string `json:"platform"`
	ScreenSize  string `json:"screen_size"`
	GeoLocation string `json:"geolocation,omitempty"`
	CreateDate  string `json:"create_date"`
	Status      int    `json:"status"`
}

//LoginSession keeps user login session for 24 hours or more
type LoginSession struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   string `json:"cid"`        //foreign key
	Serial      int64  `json:"serial"`     //company wise increase
	DeviceInfo  string `json:"device_log"` //foreign key device_log::1
	SessionCode string `json:"session_code"`
	LoginID     string `json:"login_id"` //foreign key
	IPAddress   string `json:"ip_address"`
	IPCity      string `json:"city"`
	IPCountry   string `json:"country"`
	LoginTime   string `json:"login_time"`
	LogoutTime  string `json:"logout_time"`
	CreateDate  string `json:"create_date"`
	Status      int    `json:"status"`
}

//Setting table keep only one file
type Setting struct {
	ID         string `json:"aid"`
	Type       string `json:"type"`
	CompanyID  string `json:"cid"` //foreign key
	FieldName  string `json:"field_name"`
	FieldValue string `json:"field_value"`
	Status     int    `json:"status"`
}

//Message table
type Message struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   string `json:"cid"`          //foreign key
	Serial      int64  `json:"serial"`       //company wise increase
	MessageType string `json:"message_type"` //contactus,inapp,email,sms,verify
	Sender      string `json:"sender"`       //sender login_id or email
	Receiver    string `json:"receiver"`     //receiver login_id
	Subject     string `json:"subject"`
	MessageBody string `json:"message_body"`
	CreateDate  string `json:"create_date"`
	Status      int    `json:"status"`
}

//Verification message
type Verification struct {
	ID                  string `json:"aid"`
	Type                string `json:"type"`
	CompanyID           string `json:"cid"`                  //foreign key
	Serial              int64  `json:"serial"`               //company wise increase
	MessageID           string `json:"message_id"`           //foreign key
	VerificationPurpose string `json:"verification_purpose"` //TFA/VERIFY_EMAIL/VERIFIFY_CELL
	VerificationCOde    string `json:"verification_code"`
	CreateDate          string `json:"create_date"`
	Status              int    `json:"status"`
}

//VisitorSession info
type VisitorSession struct {
	ID             string `json:"aid"`
	Type           string `json:"type"`
	CompanyID      string `json:"cid"`    //foreign key
	Serial         int64  `json:"serial"` //company wise increase
	SessionCode    string `json:"session_code"`
	Device         string `json:"device"`
	Platform       string `json:"platform"`
	ScreenSize     string `json:"screen_size"`
	BrowserVersion string `json:"browser_version"`
	OsVersion      string `json:"os_version"`
	IPAddress      string `json:"ip_address"`
	GeoLocation    string `json:"geo_location"`
	City           string `json:"city"`
	Country        string `json:"country"`
	VisitorCount   int    `json:"vcount"`
	CreateDate     string `json:"create_date"`
	UpdateDate     string `json:"update_date"`
	Status         int    `json:"status"`
}

//AccountGroup ..
type AccountGroup struct {
	ID         string `json:"aid"`
	Type       string `json:"type"`
	CompanyID  string `json:"cid"`        //foreign key
	Serial     int64  `json:"serial"`     //company wise increase
	Name       string `json:"name"`       //Group Name
	Code       string `json:"code"`       //group code
	ParentID   string `json:"parent_id"`  //Group Under same table foreign relation
	Restricted int    `json:"restricted"` //Yes, No
	GroupType  string `json:"group_type"` //Asset | Liability | Equity | Revenue | Expense
	Remarks    string `json:"remarks"`    //remarks preveius software code
	CreateDate string `json:"create_date"`
	Status     int    `json:"status"` //0=Inactive, 1=Active, 9=Deleted
}

//AccountHead accountLedger
type AccountHead struct {
	ID                   string  `json:"aid"`           //unique id
	Type                 string  `json:"type"`          //table
	CompanyID            string  `json:"cid"`           //foreign key
	Serial               int64   `json:"serial"`        //company wise increase
	AccountGroupID       string  `json:"account_group"` //foreign key AccountGroup
	Name                 string  `json:"name"`          //ledger name
	Description          string  `json:"description"`   //
	LedgerCode           string  `json:"ledger"`        //ledger number
	CurrentBalance       float64 `json:"balance"`       //ledger balance
	CurrentBalanceType   string  `json:"baltype"`       //ledger balance type Dr or Cr
	Restricted           int     `json:"restricted"`    //1=Yes, No=0
	CostCenterApplicable int     `json:"cost_center"`   //1=Yes, No=0
	CreateDate           string  `json:"create_date"`   //insert date
	Status               int     `json:"status"`        //0=Inactive, 1=Active, 9=Deleted
}

//Warehouse table
type Warehouse struct {
	ID               string `json:"aid"`
	Type             string `json:"type"`
	CompanyID        string `json:"cid"`    //foreign key
	Serial           int64  `json:"serial"` //company wise increase
	WarehouseID      int64  `json:"warehouse_id"`
	WarehouseName    string `json:"warehouse_name"`
	WarehouseDetails string `json:"warehouse_details"`
	IsDefault        bool   `json:"isdefault"` //true|false == 0|1
	Status           int    `json:"status"`
}

//UOM = Unit of Measurement
type UOM struct {
	ID         string `json:"aid"`
	Type       string `json:"type"`
	CompanyID  string `json:"cid"`    //foreign key
	Serial     int64  `json:"serial"` //company wise increase
	UnitName   string `json:"unit_name"`
	UnitSymbol string `json:"symbol"`
	Status     int    `json:"status"`
}

//Tax ..
type Tax struct {
	ID                string  `json:"aid"`
	Type              string  `json:"type"`
	CompanyID         string  `json:"cid"`            //foreign key
	Serial            int64   `json:"serial"`         //company wise increase
	Name              string  `json:"name"`           //tax name
	DisplayName       string  `json:"display_name"`   //display title
	TaxRegNumber      string  `json:"tax_reg_number"` //taxid provided by govt
	TaxType           string  `json:"tax_type"`       //vat | tax | income_tax
	Rate              float64 `json:"rate"`           //tax rate applicable in percentage
	AccountHeaderCode string  `json:"account_code"`   //ledger number attached to it
	Remarks           string  `json:"remarks"`        //remarks=default auto selected from item creation
	Status            int     `json:"status"`
}

//Department itemDepartment -> ItemLine or department are same just one under another
type Department struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   string `json:"cid"`         //foreign key
	Serial      int64  `json:"serial"`      //company wise increase
	ParentID    string `json:"parent_id"`   //same table = parent_id="" parent department id
	Name        string `json:"name"`        //child department
	Code        string `json:"code"`        //unique code
	Description string `json:"description"` //default = when add from item page
	Owner       string `json:"owner"`       //item | item_line | office
	Status      int    `json:"status"`
}

//Category or ItemCategory/ItemGroup keeps product group/category info
type Category struct {
	ID            string `json:"aid"`
	Type          string `json:"type"`
	CompanyID     string `json:"cid"`           //foreign key
	Serial        int64  `json:"serial"`        //company wise increase
	DepartmentID  string `json:"department_id"` //foreign key itemLine
	Name          string `json:"name"`
	Description   string `json:"description"`
	Code          string `json:"code"`
	Position      int64  `json:"position"`
	CategoryImage string `json:"image"`
	Owner         string `json:"owner"` //item / product | service
	Status        int    `json:"status"`
}

//ItemReward ... invoice qty 1 == 50 points
type ItemReward struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`         //item_reward
	CompanyID   string `json:"cid"`          //foreign key
	Serial      int64  `json:"serial"`       //company wise increase
	ItemID      string `json:"item_id"`      //foreign key
	RewardField string `json:"reward_field"` //qty,amount,amountPercentAsPoint, ProfitSharePercent
	RewardValue int64  `json:"reward_value"` //10 pcs
	RewardPoint int64  `json:"reward_point"` //50 points
	Status      int    `json:"status"`
}

//Item keeps product/Item info
type Item struct {
	ID                          string  `json:"aid"`
	Type                        string  `json:"type"`
	CompanyID                   string  `json:"cid"`         //foreign key
	Serial                      int64   `json:"serial"`      //company wise increase
	ItemCode                    string  `json:"item_code"`   //item_code = 12 character(4+4+4) ItemLine.Supplier.ItemSerial
	CategoryID                  string  `json:"category_id"` //foreign key as item_group
	ItemType                    string  `json:"item_type"`   //raw_material,stockable,consumable,service
	ItemName                    string  `json:"item_name"`
	ItemDescription             string  `json:"item_description"`
	ItemURL                     string  `json:"item_url"`
	ItemLine                    string  `json:"item_line,omitempty"` //*** put extra to avoid complexity in join
	ItemImage                   string  `json:"item_image"`
	Barcode                     string  `json:"barcode"`
	AssetAccount                string  `json:"asset_account"`                  //AccountHead = ledger number
	COGSAccount                 string  `json:"cogs_account"`                   //AccountHead = ledger number
	SalesAccount                string  `json:"sales_account"`                  //AccountHead = ledger number
	OpeningBalanceEquityAccount string  `json:"opening_balance_equity_account"` //AccountHead = ledger number
	VatID                       string  `json:"tax_id"`                         //foreign key
	VatPercent                  string  `json:"vat"`                            //vat percent
	BuyPrice                    float64 `json:"buy_price"`                      //cost price, Trade Price
	SalePrice                   float64 `json:"sale_price"`                     //MRP
	Tags                        string  `json:"tags,omitempty"`                 //?? department
	Supplier                    string  `json:"supplier,omitempty"`             //Account*account_type = supplier
	UnitOfMeasure               string  `json:"uom,omitempty"`
	TrackingBy                  string  `json:"tracking"` //tracking by unique_serial, lot_number, no_tracking
	ReorderLevel                int64   `json:"reorder_level"`
	ReorderQty                  int64   `json:"reorder_qty"`
	BatchNo                     string  `json:"batch"`
	ExpireDate                  string  `json:"expire_date"`
	StockQty                    int64   `json:"stock_qty"`
	PublishOnWebsite            int     `json:"publish_on_web"` //published | unpublished
	DisplayOnSales              int     `json:"display_sales"`
	DisplayOnPurchase           int     `json:"display_purchase"`
	Availability                string  `json:"availability,omitempty"` //coming soon | available | for website
	Status                      int     `json:"status"`
}

//ItemAttribute ..
type ItemAttribute struct {
	ID           string `json:"aid"`
	Type         string `json:"type"`
	CompanyID    string `json:"cid"`           //foreign key
	Serial       int64  `json:"serial"`        //company wise increase
	ItemID       string `json:"item_id"`       //item_id foreign key
	AttributeKey string `json:"attribute_key"` //attr key
	Position     int64  `json:"position"`      //serial / position
	KeyType      string `json:"key_type"`      //select, text, radio, color
	DefaultValue string `json:"default_value"`
	Status       int    `json:"status"`
}

//ItemAttributeValue ..
type ItemAttributeValue struct {
	ID              string `json:"aid"`
	Type            string `json:"type"`
	CompanyID       string `json:"cid"`             //foreign key
	Serial          int64  `json:"serial"`          //company wise increase
	ItemID          string `json:"item_id"`         //foreign key item_id
	ItemAttributeID string `json:"attribute_id"`    //foreign key ItemAttribute
	AttributeValue  string `json:"attribute_value"` //select, text, radio, color
	Position        int64  `json:"position"`        //serial / position
	Status          int    `json:"status"`
}

//Rateplan ..
type Rateplan struct {
	ID               string  `json:"aid"`
	Type             string  `json:"type"`
	CompanyID        string  `json:"cid"`           //foreign key
	Serial           int64   `json:"serial"`        //company wise increase
	CustomerID       string  `json:"account_id"`    //connect with account receivable/customer/supplier
	ItemID           int64   `json:"item_id"`       //item_id
	ItemLoyaltyPoint int64   `json:"loyalty_point"` //bonus point
	Rate             float64 `json:"rate"`          //1 tk = ???
	Rebate           float64 `json:"rebate"`        //discount percentage
	Charge           float64 `json:"charge"`        //charge percentage
	Owner            string  `json:"owner"`         //customer, item
	Status           int     `json:"status"`
}

//FileStore general table ***
type FileStore struct {
	ID           string `json:"aid"`
	Type         string `json:"type"`
	CompanyID    string `json:"cid"`
	Serial       int64  `json:"serial"`      //company wise increase
	ItemID       string `json:"item_id"`     //foreign key
	OwnerTable   string `json:"owner_table"` //table_name
	Parameter    string `json:"parameter"`
	FileType     string `json:"file_type"`
	FileLocation string `json:"file_location"`
	Remarks      string `json:"remarks"`
	Status       int    `json:"status"`
}

//DocKeeper keeps all document info
type DocKeeper struct {
	ID             string  `json:"aid"`
	Type           string  `json:"type"`
	CompanyID      string  `json:"cid"`                    //foreign key
	WarehouseID    string  `json:"warehouse_id,omitempty"` //foreign key
	Serial         int64   `json:"serial"`                 //company wise increase
	DocName        string  `json:"doc_name"`
	DocType        string  `json:"doc_type"`
	DocRef         string  `json:"doc_ref"`
	DocNumber      string  `json:"doc_number"`
	DocDescription string  `json:"doc_description"`
	PostingDate    string  `json:"posting_date"`
	TaxRule        string  `json:"tax_rule"`
	LoginID        string  `json:"login_id"`   //foreign key
	AccountID      string  `json:"account_id"` //foreign key
	TotalPayable   float64 `json:"total_payable"`
	TotalTax       float64 `json:"total_tax"`
	TotalDiscount  float64 `json:"total_discount"`
	PaymentStatus  string  `json:"payment_status"`
	DocStatus      string  `json:"doc_status"`
	CreateDate     string  `json:"create_date"`
	Status         int     `json:"status"`
}

//TransactionRecord keeps all transaction info like sales, purchase, online oreder
type TransactionRecord struct {
	ID             string  `json:"aid"`
	Type           string  `json:"type"`
	CompanyID      string  `json:"cid"`    //foreign key
	Serial         int64   `json:"serial"` //company wise increase
	TrxType        string  `json:"trx_type"`
	DocNumber      string  `json:"doc_number"` //foriegn key
	ItemID         string  `json:"item_id"`
	StockInfo      string  `json:"stock_info"`
	ProductSerial  string  `json:"item_serial"` //SKU or barcode
	Quantity       int     `json:"quantity"`
	Rate           float64 `json:"rate"`  //per unit price, GAAP Compliance
	Price          float64 `json:"price"` //price = qty x rate'
	DiscountRate   float64 `json:"discount_rate"`
	TaxRate        float64 `json:"tax_rate"`
	DiscountAmount float64 `json:"discount_amount"` //(price x discount_rate)/100
	TaxableAmount  float64 `json:"taxable_amount"`  //price-total_discount
	TaxAmount      float64 `json:"tax_amount"`
	PayableAmount  float64 `json:"payable_amount"` //taxable_amount + total_tax | price-total_discount+total_tax
	CreateDate     string  `json:"create_date"`
	Status         int     `json:"status"` //0=Inactive, 1=Active, 9=Deleted
}

//StockMovement keeps stock transfer from warehouse to warehouse or other
type StockMovement struct {
	ID            string `json:"aid"`
	Type          string `json:"type"`
	CompanyID     string `json:"cid"`        //foreign key
	Serial        int64  `json:"serial"`     //company wise increase
	DocNumber     string `json:"doc_number"` //foriegn key
	StockNote     string `json:"stock_note"`
	ItemID        string `json:"item_id"`      //foreign key
	ProductSerial string `json:"item_serial"`  //SKU or barcode
	WarehouseID   string `json:"warehouse_id"` //foreign key
	MovementType  string `json:"mtype"`        //movement type = IN/OUT
	Quantity      int    `json:"quantity"`
	StockBalance  int    `json:"stock_balance"`
	Status        int    `json:"status"` //0=Inactive, 1=Active, 9=Deleted

}

//LedgerTransaction table stores accounting transaction between ledgers/accounting_header
type LedgerTransaction struct {
	ID           string  `json:"aid"`
	Type         string  `json:"type"`
	CompanyID    string  `json:"cid"`          //foreign key
	Serial       int64   `json:"serial"`       //company wise increase
	DocNumber    string  `json:"doc_number"`   //foriegn key
	VoucherName  string  `json:"voucher_name"` //
	LedgerNumber string  `json:"ledger"`       //ledger code on AccountHeader
	Description  string  `json:"description"`
	Debit        float64 `json:"debit"`
	Credit       float64 `json:"credit"`
	Balance      float64 `json:"balance"`
	BalanceType  string  `json:"baltype"` //balance type Dr/Cr/Eq
	CreateDate   string  `json:"create_date"`
	Status       int     `json:"status"` //0=Inactive, 1=Active, 9=Deleted
}
