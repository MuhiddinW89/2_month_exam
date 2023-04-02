package models

type Staff struct {
	StaffId     int    `json:"staff_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Active      int    `json:"active"`
	StoreId     int    `json:"store_id"`
	StoreData   *Store `json:"store_data"`
	ManagerId   int    `json:"manager_id"`
	ManagerData *Staff `json:"manager_data"`
}

type StaffPrimaryKey struct {
	StaffId int `json:"staff_id"`
}

type CreateStaff struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Active    int    `json:"active"`
	StoreId   int    `json:"store_id"`
	ManagerId int    `json:"manager_id"`
}

type UpdateStaff struct {
	StaffId   int    `json:"staff_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Active    int    `json:"active"`
	StoreId   int    `json:"store_id"`
	ManagerId int    `json:"manager_id"`
}

type GetListStaffRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStaffResponse struct {
	Count  int      `json:"count"`
	Staffs []*Staff `json:"staffs"`
}


type GetStaffSellsById struct {
	FirstName string `json:"first_name"`
	CategoryData *Category `json:"category_data"`
	ProductData  *Product `json:"product_data"`
	OrderItemData *OrderItem `json:"order_item_data"`
	OrderItemPrice *OrderItem `json:"order_item_price"`
	OrderData   *Order `json:"orders_data"`
}