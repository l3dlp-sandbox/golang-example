package domain

type RetryEvent struct {
	Value []byte `json:"value"`
	Meta  struct {
		Error      string `json:"error"`
		RetryCount int    `json:"retry_count"`
	} `json:"meta"`
}
type DriverTripTransaction struct {
	Body struct {
		Amount                  float64 `json:"amount"`
		CreatedBy               int     `json:"created_by"`
		CreatedTime             int64   `json:"created_time"`
		Description             string  `json:"description"`
		DriverID                int     `json:"driver_id"`
		TransactionCategory     int     `json:"transaction_category"`
		TransactionCategoryName string  `json:"transaction_category_name"`
		TransactionID           int     `json:"transaction_id"`
		TransactionType         string  `json:"transaction_type"`
		TripID                  int     `json:"trip_id"`
	} `json:"body"`
	CreatedAt int64  `json:"created_at"`
	Expiry    int    `json:"expiry"`
	ID        string `json:"id"`
	TraceInfo struct {
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
		SpanID   int  `json:"span_id"`
		TraceID  struct {
			High int `json:"high"`
			Low  int `json:"low"`
		} `json:"trace_id"`
	} `json:"trace_info"`
	Type    string `json:"type"`
	Version int    `json:"version"`
}

type DriverTripTransactionRequest struct {
	Body struct {
		Amount                  float64 `json:"amount"`
		CreatedBy               int     `json:"created_by"`
		CreatedTime             int64   `json:"created_time"`
		Description             string  `json:"description"`
		DriverID                int     `json:"driver_id"`
		TransactionCategory     int     `json:"transaction_category"`
		TransactionCategoryName string  `json:"transaction_category_name"`
		TransactionType         string  `json:"transaction_type"`
		TripID                  int     `json:"trip_id"`
	} `json:"body"`
	CreatedAt int64  `json:"created_at"`
	Expiry    int64  `json:"expiry"`
	ID        string `json:"id"`
	TraceInfo struct {
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
		SpanID   int  `json:"span_id"`
		TraceID  struct {
			High int `json:"high"`
			Low  int `json:"low"`
		} `json:"trace_id"`
	} `json:"trace_info"`
	Type    string `json:"type"`
	Version int    `json:"version"`
}

//**************driver_location*****************************************

type DriverLocation struct {
	DriverId int      `json:"driver_id"`
	Location Location `json:"location"`
}
type Location struct {
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Accuracy  float64 `json:"accuracy"`
	Speed     float64 `json:"speed"`
	Timestamp int64   `json:"timestamp"`
	Bearing   float64 `json:"bearing"`
}

type DriverLocationChanged struct {
	Id        string         `json:"id"`
	Type      string         `json:"type"`
	Body      DriverLocation `json:"body"`
	CreatedAt int64          `json:"created_at"`
	Expiry    int64          `json:"expiry"`
	Version   int            `json:"version"`
	TraceInfo TraceInfo      `json:"trace_info"`
}

type TraceInfo struct {
	TraceId  int64 `json:"trace_id"`
	SpanId   int64 `json:"span_id"`
	ParentId int64 `json:"parent_id"`
	Sampled  bool  `json:"sampled"`
}

type TraceId struct {
	High int64 `json:"high"`
	Low  int64 `json:"low"`
}

//*****************bank_status_change**********************************

type BankStatusChanged struct {
	Body      BankStatusChangedBody `json:"body"`
	TraceInfo struct {
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
		SpanID   int  `json:"span_id"`
		TraceID  struct {
			High int `json:"high"`
			Low  int `json:"low"`
		} `json:"trace_id"`
	} `json:"trace_info"`
	Type      string `json:"type"`
	Version   int    `json:"version"`
	CreatedAt int64  `json:"created_at"`
	Expiry    int64  `json:"expiry"`
	ID        string `json:"id"`
}

type BankStatusChangedBody struct {
	Id                     int64  `json:"id"`
	Bank                   int    `json:"bank"`
	TransactionRefId       string `json:"transaction_reference_id"`
	PaymentType            int    `json:"payment_type"`
	PaymentTypeReferenceId int64  `json:"payment_type_reference_id"`
	Status                 int    `json:"status"`
	UpdatedDatetime        int64  `json:"updated_datetime"`
	CreatedDatetime        int64  `json:"created_datetime"`
}

//***************external payment request received
type ExternalPaymentRequestStatus struct {
	Body struct {
		CreatedDateTime int64  `json:"created_date_time"`
		DriverID        int    `json:"driver_id"`
		ID              int    `json:"id"`
		OrderID         string `json:"order_id"`
		PaymentTry      int    `json:"payment_try"`
		PaymentMethod   int    `json:"payment_method"`
		PaymentStatus   int    `json:"payment_status"`
	} `json:"body"`
	CreatedAt int64  `json:"created_at"`
	Expiry    int64  `json:"expiry"`
	ID        string `json:"id"`
	TraceInfo struct {
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
		SpanID   int  `json:"span_id"`
		TraceID  struct {
			High int `json:"high"`
			Low  int `json:"low"`
		} `json:"trace_id"`
	} `json:"trace_info"`
	Type    string `json:"type"`
	Version int    `json:"version"`
}
type TripCompleted struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Body struct {
		PassengerID  int    `json:"passenger_id"`
		DriverID     int    `json:"driver_id"`
		CurrencyCode string `json:"currency_code"`
		Trip         struct {
			ID       int     `json:"id"`
			BookedBy int     `json:"booked_by"`
			Distance int     `json:"distance"`
			TripCost float64 `json:"trip_cost"`
			Flags    struct {
				Itc bool `json:"itc"`
			} `json:"flags"`
			Payment   []Payment `json:"payment"`
			Discount  float64   `json:"discount"`
			PromoCode string    `json:"promo_code"`
			Corporate struct {
				CompanyID    int `json:"company_id"`
				DepartmentID int `json:"department_id"`
			} `json:"corporate"`
			ActualPickup struct {
				Address string  `json:"address"`
				Lat     float64 `json:"lat"`
				Lng     float64 `json:"lng"`
			} `json:"actual_pickup"`
			ActualDrop struct {
				Address string  `json:"address"`
				Lat     float64 `json:"lat"`
				Lng     float64 `json:"lng"`
			} `json:"actual_drop"`
		} `json:"trip"`
	} `json:"body"`
	CreatedAt int64 `json:"created_at"`
	Expiry    int64 `json:"expiry"`
	Version   int   `json:"version"`
	TraceInfo struct {
		TraceID struct {
			High int64 `json:"high"`
			Low  int64 `json:"low"`
		} `json:"trace_id"`
		SpanID   int  `json:"span_id"`
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
	} `json:"trace_info"`
}
type Payment struct {
	Method int     `json:"method"`
	Amount float64 `json:"amount"`
}

type BodyRestaurant struct {
	MerchantId    int64  `json:"merchant_id"`
	BankName      string `json:"bank_name"`
	BankBranch    string `json:"bank_branch"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	UpdatedBy     int    `json:"updated_by"`
	CompanyId     int    `json:"company_id"`
}
type RestaurantBankAccountCreated struct {
	Body      BodyRestaurant `json:"body"`
	CreatedAt int64          `json:"created_at"`
	Expiry    int            `json:"expiry"`
	ID        string         `json:"id"`
	TraceInfo struct {
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
		SpanID   int  `json:"span_id"`
		TraceID  struct {
			High int `json:"high"`
			Low  int `json:"low"`
		} `json:"trace_id"`
	} `json:"trace_info"`
	Type    string `json:"type"`
	Version int    `json:"version"`
}

type PayRecord struct {
	TransactionId   int64   `json:"transaction_id"`
	TripId          int     `json:"trip_id"`
	DriverId        int     `json:"driver_id"`
	BookingFromCid  int     `json:"company_id"`
	Promocode       string  `json:"promo_code"`
	PaymentMethod   int     `json:"payment_method"`
	AccountHolder   string  `json:"account_holder"`
	AccountNumber   string  `json:"account_number"`
	BankName        int     `json:"bank_name"`
	Amount          float64 `json:"amount"`
	IsCredit        bool    `json:"is_credit"`
	DbDescription   string  `json:"db_description"`
	DriverTelephone string  `json:"driver_phone"`
	TxnCategory     int64   `json:"txn_category"`
	StakeHolderType string  `json:"stakeholder_type"`
	StakeHolderId   int64   `json:"stakeholder_id"`
}
type Event struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Body      []byte    `json:"body"`
	CreatedAt int64     `json:"created_at"`
	Expiry    int64     `json:"expiry"`
	Version   int       `json:"version"`
	TraceId   TraceInfo `json:"trace_info"`
}
type AutoSettlementEvent struct {
	PaymentRecords []PayRecord `json:"payment_records"`
}

type PaymentRecordsNew struct {
	TransactionID   int64  `json:"transaction_id"`
	TripID          int    `json:"trip_id"`
	DriverID        int    `json:"driver_id"`
	CompanyID       int    `json:"company_id"`
	PromoCode       string `json:"promo_code"`
	PaymentMethod   int    `json:"payment_method"`
	AccountHolder   string `json:"account_holder"`
	AccountNumber   string `json:"account_number"`
	BankName        int    `json:"bank_name"`
	Amount          int    `json:"amount"`
	IsCredit        bool   `json:"is_credit"`
	DbDescription   string `json:"db_description"`
	DriverPhone     string `json:"driver_phone"`
	TxnCategory     int    `json:"txn_category"`
	StakeholderType string `json:"stakeholder_type"`
	StakeholderID   int64  `json:"stakeholder_id"`
}
type AutoSettlementNew struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Body struct {
		PaymentRecords []PayRecord `json:"payment_records"`
	} `json:"body"`
	CreatedAt int `json:"created_at"`
	Expiry    int `json:"expiry"`
	Version   int `json:"version"`
	TraceInfo struct {
		TraceID struct {
			High int64 `json:"high"`
			Low  int64 `json:"low"`
		} `json:"trace_id"`
		SpanID   int  `json:"span_id"`
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
	} `json:"trace_info"`
}
type TripCreated struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Body struct {
		Module           int    `json:"module"`
		ServiceGroupCode string `json:"service_group_code"`
		BookedBy         int    `json:"booked_by"`
		TripID           int    `json:"trip_id"`
		VehicleType      int    `json:"vehicle_type"`
		PreBooking       bool   `json:"pre_booking"`
		Passenger        struct {
			ID int `json:"id"`
		} `json:"passenger"`
		Driver struct {
			ID int `json:"id"`
		} `json:"driver"`
		Corporate struct {
			ID    int `json:"id"`
			DepID int `json:"dep_id"`
		} `json:"corporate"`
		Pickup struct {
			Time     int         `json:"time"`
			Location []Location2 `json:"location"`
		} `json:"pickup"`
		Drop struct {
			Location []Location2 `json:"location"`
		} `json:"drop"`
		Promotion struct {
			Code string `json:"code"`
		} `json:"promotion"`
		Region struct {
			Ids []int `json:"ids"`
		} `json:"region"`
		Payment struct {
			PrimaryMethod   int `json:"primary_method"`
			SecondaryMethod int `json:"secondary_method"`
		} `json:"payment"`
		Comments struct {
			Remark      string `json:"remark"`
			DriverNotes string `json:"driver_notes"`
		} `json:"comments"`
		Filters struct {
			Driver struct {
				LanguageID int `json:"language_id"`
			} `json:"driver"`
			Vehicle struct {
				CompanyID int `json:"company_id"`
				BrandID   int `json:"brand_id"`
				ColorID   int `json:"color_id"`
			} `json:"vehicle"`
		} `json:"filters"`
		Surge struct {
			RegionID int `json:"region_id"`
			Value    int `json:"value"`
		} `json:"surge"`
		FareDetails struct {
			FareType         string  `json:"fare_type"`
			MinKm            int     `json:"min_km"`
			MinFare          float64 `json:"min_fare"`
			AdditionalKmFare float64 `json:"additional_km_fare"`
			WaitingTimeFare  float64 `json:"waiting_time_fare"`
			FreeWaitingTime  int     `json:"free_waiting_time"`
			NightFare        int     `json:"night_fare"`
			RideHours        int     `json:"ride_hours"`
			ExtraRideFare    int     `json:"extra_ride_fare"`
			DriverBata       float64 `json:"driver_bata"`
			TripType         int     `json:"trip_type"`
		} `json:"fare_details"`
	} `json:"body"`
	CreatedAt int `json:"created_at"`
	Expiry    int `json:"expiry"`
	Version   int `json:"version"`
	TraceInfo struct {
		TraceID struct {
			High int64 `json:"high"`
			Low  int64 `json:"low"`
		} `json:"trace_id"`
		SpanID   int  `json:"span_id"`
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
	} `json:"trace_info"`
}

type Location2 struct {
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type TripCancelled struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Body struct {
		TripID        int    `json:"trip_id"`
		CancelledFrom int    `json:"cancelled_from"`
		CancelledBy   int    `json:"cancelled_by"`
		ReasonID      int    `json:"reason_id"`
		Note          string `json:"note"`
		CancelType    int    `json:"cancel_type"`
	} `json:"body"`
	CreatedAt int64 `json:"created_at"`
	Expiry    int64 `json:"expiry"`
	Version   int64 `json:"version"`
	TraceInfo struct {
		TraceID struct {
			High int64 `json:"high"`
			Low  int64 `json:"low"`
		} `json:"trace_id"`
		SpanID   int  `json:"span_id"`
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
	} `json:"trace_info"`
}
