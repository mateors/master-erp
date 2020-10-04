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
	CompanyID  string `json:"cid"` //foreign key
	AccessID   int    `json:"access_id"`
	AccessName string `json:"access_name"`
	Status     int    `json:"status"`
}

//Login #3  all user account login table
type Login struct {
	ID         string `json:"aid"`
	Type       string `json:"type"`
	CompanyID  string `json:"cid"` //foreign key
	LoginID    int64  `json:"login_id"`
	AccessID   string `json:"access_id"` //foreign key
	AccessName string `json:"access_name"`
	UserName   string `json:"username"` //email as username
	Password   string `json:"passw"`
	CreateDate string `json:"create_date"`
	LastLogin  string `json:"last_login,omitempty"`
	Status     int    `json:"status"`
}

//Contact info for account and user
type Contact struct {
	ID        string `json:"aid"`
	Type      string `json:"type"` //table_name
	CompanyID string `json:"cid"`  //foreign key
	ContactID int64  `json:"contact_id"`
	//mobile,email,pager,phone,website,socialmedia
	ContactType string `json:"contact_type"`
	ContactData string `json:"contact_date"`
	//owner table info
	TableName string `json:"table_name"` //account
	PkField   string `json:"pkfield"`    //account_id
	PkValue   string `json:"pkvalue"`    //
	Status    int    `json:"status"`
}

//Address ...
type Address struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   string `json:"cid"`        //foreign key
	AccountID   int64  `json:"account_id"` //foreign_key
	Serial      int64  `json:"serial"`
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

//Account #4 ::cid::account_id
type Account struct {
	//ContactInfo []Contact `json:"contact_info,omitempty"`
	ID          string `json:"aid"`
	Type        string `json:"type"` //account
	CompanyID   string `json:"cid"`  //foreign key
	ParentID    string `json:"parent_id,omitempty"`
	AccountID   int64  `json:"account_id"`
	AccountType string `json:"account_type"` //vendor,supplier,customer,consumer
	AccountName string `json:"account_name"`
	LoginID     string `json:"login_id"` //foreign key
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	DateOfBirth string `json:"dob,omitempty"`
	Gender      string `json:"gender"` //female,male,other
	Mobile      string `json:"mobile"` //
	Email       string `json:"email"`  //
	CreateDate  string `json:"create_date"`
	UpdateDate  string `json:"update_date,omitempty"`
	Status      int    `json:"status"`
}

//ActivityLog ***
type ActivityLog struct {
	ID           string `json:"aid"`
	Type         string `json:"type"`
	CompanyID    string `json:"cid"` //foreign key
	LogID        int64  `json:"logid"`
	ActivityType string `json:"activity_type"`
	OwnerTable   string `json:"owner_table"` //table_name
	Parameter    string `json:"parameter"`   //key=val or aid
	LogDetails   string `json:"log_details"`
	IPAddress    string `json:"ip_address"`
	LoginID      string `json:"login_id"` //foreign key
	CreateDate   string `json:"create_date"`
	Status       int    `json:"status"`
}

//DeviceLog tracks user device info (where they login to the system)
type DeviceLog struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   string `json:"cid"`      //foreign key
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
	CompanyID   string `json:"cid"` //foreign key
	MessageID   int64  `json:"message_id"`
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
	CompanyID           string `json:"cid"` //foreign key
	Serial              int64  `json:"serial"`
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
	CompanyID      string `json:"cid"` //foreign key
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

//Warehouse table
type Warehouse struct {
	ID               string `json:"aid"`
	Type             string `json:"type"`
	CompanyID        string `json:"cid"` //foreign key
	WarehouseID      int64  `json:"warehouse_id"`
	WarehouseName    string `json:"warehouse_name"`
	WarehouseDetails string `json:"warehouse_details"`
	IsDefault        bool   `json:"isdefault"` //true|false == 0|1
	Status           int    `json:"status"`
}

//Department itemDepartment -> ItemLine or department are same just one under another
type Department struct {
	ID                 string `json:"aid"`
	Type               string `json:"type"`
	CompanyID          string `json:"cid"`             //foreign key
	DepartmentID       string `json:"department_id"`   //company wise increase
	ParentDepartmentID string `json:"parent_id"`       //same table = parent_id="" parent department id
	DeaprtmentName     string `json:"department_name"` //child department
	Owner              string `json:"owner"`           //item | office
	Status             int    `json:"status"`
}

//ItemCategory or ItemGroup keeps product group/category info
type ItemCategory struct {
	ID                  string `json:"aid"`
	Type                string `json:"type"`
	CompanyID           string `json:"cid"`         //foreign key
	CategoryID          int64  `json:"category_id"` //company wise increase
	DepartmentID        string `json:"department_id"`
	Position            int    `json:"position"`
	CategoryName        string `json:"category_name"`
	CategoryDescription string `json:"category_description"`
	CategoryImage       string `json:"category_image"`
	CategoryFor         string `json:"category_for"` //item / product | service
	Status              int    `json:"status"`
}

//ItemReward ... invoice qty 1 == 50 points
type ItemReward struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`         //item_reward
	CompanyID   string `json:"cid"`          //foreign key
	ItemID      string `json:"item_id"`      //foreign key
	RewardUnit  string `json:"reward_unit"`  //qty,amount,amountPercentAsPoint, ProfitSharePercent
	RewardValue int64  `json:"reward_value"` //10 pcs
	RewardPoint int64  `json:"reward_point"` //50 points
	Status      int    `json:"status"`
}

//Item keeps product/Item info
type Item struct {
	ID                          string  `json:"aid"`
	Type                        string  `json:"type"`
	CompanyID                   string  `json:"cid"`         //foreign key
	ItemID                      int64   `json:"item_id"`     //item_id company wise increase
	ItemCode                    string  `json:"item_code"`   //item_code = 12 character(4+4+4) ItemLine.Supplier.ItemSerial
	CategoryID                  string  `json:"category_id"` //foreign key
	ItemType                    string  `json:"item_type"`   //raw_material,stockable,consumable,service
	ItemName                    string  `json:"item_name"`
	ItemDescription             string  `json:"item_description"`
	ItemURL                     string  `json:"item_url"`
	ItemImage                   string  `json:"item_image"`
	Barcode                     string  `json:"barcode"`
	AssetAccount                string  `json:"asset_account,omitempty"`                  //AccountHead = ledger number
	COGSAccount                 string  `json:"cogs_account,omitempty"`                   //AccountHead = ledger number
	SalesAccount                string  `json:"sales_account,omitempty"`                  //AccountHead = ledger number
	OpeningBalanceEquityAccount string  `json:"opening_balance_equity_account,omitempty"` //AccountHead = ledger number
	VatID                       string  `json:"tax_id"`                                   //foreign key
	BuyPrice                    float64 `json:"buy_price"`                                //cost price, Trade Price
	SalePrice                   float64 `json:"sale_price"`                               //MRP
	VatPercent                  string  `json:"vat"`                                      //vat percent
	Tags                        string  `json:"tags,omitempty"`                           //?? department
	Supplier                    string  `json:"supplier,omitempty"`                       //Account*account_type = supplier
	UnitOfMeasure               string  `json:"uom,omitempty"`
	TrackingBy                  string  `json:"tracking,omitempty"` //tracking by unique_serial, lot_number, no_tracking
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
	ItemID       string `json:"item_id"`       //item_id foreign key
	AttributeKey string `json:"attribute_key"` //attr key
	Position     int    `json:"position"`      //serial / position
	KeyType      string `json:"key_type"`      //select, text, radio, color
	DefaultValue string `json:"default_value"`
	Status       int    `json:"status"`
}

//ItemAttributeValue ..
type ItemAttributeValue struct {
	ID              string `json:"aid"`
	Type            string `json:"type"`
	CompanyID       string `json:"cid"`             //foreign key
	ItemID          string `json:"item_id"`         //foreign key item_id
	ItemAttributeID string `json:"attribute_id"`    //foreign key ItemAttribute
	AttributeValue  string `json:"attribute_value"` //select, text, radio, color
	Status          int    `json:"status"`
}

//Rateplan ..
type Rateplan struct {
	ID               string  `json:"aid"`
	Type             string  `json:"type"`
	CompanyID        string  `json:"cid"`           //foreign key
	Serial           int64   `json:"serial"`        //cid wise auto increment
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
	FileID       int64  `json:"file_id"`
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
	DocID          int64   `json:"doc_id"`                 //doc serial
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

//TransactionRecord keeps all transaction info
type TransactionRecord struct {
	ID             string  `json:"aid"`
	Type           string  `json:"type"`
	CompanyID      string  `json:"cid"` //foreign key
	TrxID          int64   `json:"trx_id"`
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

//StockMovement table
type StockMovement struct {
	ID            string `json:"aid"`
	Type          string `json:"type"`
	CompanyID     string `json:"cid"` //foreign key
	Serial        int64  `json:"serial"`
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

//AccountGroup ..
type AccountGroup struct {
	ID                 string `json:"aid"`
	Type               string `json:"type"`
	CompanyID          string `json:"cid"` //foreign key
	Serial             int64  `json:"serial"`
	Name               string `json:"name"`
	Restricted         string `json:"restricted"` //Yes, No
	ParentAccountGroup string `json:"parent_id"`  //same table foreign relation
	GroupType          string `json:"group_type"` //Asset | Liability | Equity | Revenue | Expense
	Status             int    `json:"status"`     //0=Inactive, 1=Active, 9=Deleted
}

//AccountHeader ..
type AccountHeader struct {
	ID                   string  `json:"aid"`
	Type                 string  `json:"type"`
	CompanyID            string  `json:"cid"`             //foreign key
	Serial               int64   `json:"serial"`          //id
	Name                 string  `json:"name"`            //ledger name
	Code                 string  `json:"code"`            //ledger number
	Baltype              string  `json:"baltype"`         //balance type
	Opening              float64 `json:"opening"`         //opening balance
	CurrentBalance       float64 `json:"current_balance"` //ledger balance
	CurrentBalanceType   string  `json:"curbal_type"`     //ledger balance type Dr or Cr
	Restricted           string  `json:"restricted"`      //Yes, No
	Description          string  `json:"description"`
	AccountGroupID       string  `json:"account_group"`          //foreign key
	CostCenterApplicable string  `json:"cost_center_applicable"` //Yes, No
	Status               int     `json:"status"`                 //0=Inactive, 1=Active, 9=Deleted
}

//LedgerTransaction table
type LedgerTransaction struct {
	ID           string  `json:"aid"`
	Type         string  `json:"type"`
	CompanyID    string  `json:"cid"` //foreign key
	Serial       int64   `json:"serial"`
	DocNumber    string  `json:"doc_number"` //foriegn key
	VoucherName  string  `json:"voucher_name"`
	LedgerNumber string  `json:"ledger_number"`
	Description  string  `json:"description"`
	Debit        float64 `json:"debit"`
	Credit       float64 `json:"credit"`
	Balance      float64 `json:"balance"`
	BalanceType  string  `json:"baltype"` //balance type Dr/Cr/Eq
	Status       int     `json:"status"`  //0=Inactive, 1=Active, 9=Deleted
}
