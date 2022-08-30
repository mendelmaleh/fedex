package main

// generated with github.com/mendelmaleh/gojson

type TrackResponse struct {
	CustomerTransactionID string `json:"customerTransactionId"`
	Output                Output `json:"output"`
	TransactionID         string `json:"transactionId"`
}

type AncillaryDetails struct {
	Action            string `json:"action"`
	ActionDescription string `json:"actionDescription"`
	Reason            string `json:"reason"`
	ReasonDescription string `json:"reasonDescription"`
}

type DeliveryDetails struct {
	ActualDeliveryAddress             ActualDeliveryAddress              `json:"actualDeliveryAddress"`
	DeliveryAttempts                  string                             `json:"deliveryAttempts"`
	DeliveryOptionEligibilityDetails  []DeliveryOptionEligibilityDetails `json:"deliveryOptionEligibilityDetails"`
	DeliveryToday                     bool                               `json:"deliveryToday"`
	DestinationServiceArea            string                             `json:"destinationServiceArea"`
	DestinationServiceAreaDescription string                             `json:"destinationServiceAreaDescription"`
	LocationDescription               string                             `json:"locationDescription"`
	LocationType                      string                             `json:"locationType"`
	OfficeOrderDeliveryMethod         string                             `json:"officeOrderDeliveryMethod"`
	ReceivedByName                    string                             `json:"receivedByName"`
	SignedByName                      string                             `json:"signedByName"`
}

type TrackResults struct {
	AdditionalTrackingInfo        AdditionalTrackingInfo          `json:"additionalTrackingInfo"`
	AvailableImages               []AvailableImages               `json:"availableImages"`
	AvailableNotifications        []string                        `json:"availableNotifications"`
	ConsolidationDetail           []ConsolidationDetail           `json:"consolidationDetail"`
	CustomDeliveryOptions         []CustomDeliveryOptions         `json:"customDeliveryOptions"`
	DateAndTimes                  []DateAndTimes                  `json:"dateAndTimes"`
	DeliveryDetails               DeliveryDetails                 `json:"deliveryDetails"`
	DestinationLocation           DestinationLocation             `json:"destinationLocation"`
	DistanceToDestination         DistanceToDestination           `json:"distanceToDestination"`
	Error                         Error                           `json:"error"`
	EstimatedDeliveryTimeWindow   Window_sub7                     `json:"estimatedDeliveryTimeWindow"`
	GoodsClassificationCode       string                          `json:"goodsClassificationCode"`
	HoldAtLocation                HoldAtLocation                  `json:"holdAtLocation"`
	InformationNotes              []InformationNotes              `json:"informationNotes"`
	LastUpdatedDestinationAddress ActualDeliveryAddress           `json:"lastUpdatedDestinationAddress"`
	LatestStatusDetail            LatestStatusDetail              `json:"latestStatusDetail"`
	MeterNumber                   string                          `json:"meterNumber"`
	OriginLocation                DestinationLocation             `json:"originLocation"`
	PackageDetails                PackageDetails                  `json:"packageDetails"`
	PieceCounts                   []PieceCounts                   `json:"pieceCounts"`
	ReasonDetail                  ReasonDetail                    `json:"reasonDetail"`
	RecipientInformation          LocationContactAndAddress_sub21 `json:"recipientInformation"`
	ReturnDetail                  ReturnDetail                    `json:"returnDetail"`
	ScanEvents                    []ScanEvents                    `json:"scanEvents"`
	ServiceCommitMessage          ServiceCommitMessage            `json:"serviceCommitMessage"`
	ServiceDetail                 ServiceDetail                   `json:"serviceDetail"`
	ShipmentDetails               ShipmentDetails                 `json:"shipmentDetails"`
	ShipperInformation            LocationContactAndAddress_sub21 `json:"shipperInformation"`
	SpecialHandlings              []SpecialHandlings              `json:"specialHandlings"`
	StandardTransitTimeWindow     Window_sub7                     `json:"standardTransitTimeWindow"`
	TrackingNumberInfo            TrackingNumberInfo              `json:"trackingNumberInfo"`
}

type LocationContactAndAddress_sub21 struct {
	Address ActualDeliveryAddress `json:"address"`
	Contact Contact               `json:"contact"`
}

type LocationContactAndAddress struct {
	Address Address `json:"address"`
	Contact Contact `json:"contact"`
}

type Output struct {
	Alerts               string                 `json:"alerts"`
	CompleteTrackResults []CompleteTrackResults `json:"completeTrackResults"`
}

type LatestStatusDetail struct {
	AncillaryDetails []AncillaryDetails    `json:"ancillaryDetails"`
	Code             string                `json:"code"`
	DelayDetail      DelayDetail           `json:"delayDetail"`
	DerivedCode      string                `json:"derivedCode"`
	Description      string                `json:"description"`
	ScanLocation     ActualDeliveryAddress `json:"scanLocation"`
	StatusByLocale   string                `json:"statusByLocale"`
}

type ReturnDetail struct {
	AuthorizationName string         `json:"authorizationName"`
	ReasonDetail      []ReasonDetail `json:"reasonDetail"`
}

type ShipmentDetails struct {
	BeforePossessionStatus bool             `json:"beforePossessionStatus"`
	ContentPieceCount      string           `json:"contentPieceCount"`
	Contents               []Contents       `json:"contents"`
	SplitShipments         []SplitShipments `json:"splitShipments"`
	Weight                 []Weight         `json:"weight"`
}

type Window struct {
	Begins string `json:"begins"`
	Ends   string `json:"ends"`
}

type TrackingNumberInfo struct {
	CarrierCode            string `json:"carrierCode"`
	TrackingNumber         string `json:"trackingNumber"`
	TrackingNumberUniqueID string `json:"trackingNumberUniqueId"`
}

type ActualDeliveryAddress struct {
	City                string   `json:"city"`
	Classification      string   `json:"classification"`
	CountryCode         string   `json:"countryCode"`
	CountryName         string   `json:"countryName"`
	PostalCode          string   `json:"postalCode"`
	Residential         bool     `json:"residential"`
	StateOrProvinceCode string   `json:"stateOrProvinceCode"`
	StreetLines         []string `json:"streetLines"`
	UrbanizationCode    string   `json:"urbanizationCode"`
}

type Address struct {
	City                string   `json:"city"`
	Classification      string   `json:"classification"`
	CountryCode         string   `json:"countryCode"`
	PostalCode          string   `json:"postalCode"`
	Residential         bool     `json:"residential"`
	StateOrProvinceCode string   `json:"stateOrProvinceCode"`
	StreetLines         []string `json:"streetLines"`
	UrbanizationCode    string   `json:"urbanizationCode"`
}

type InformationNotes struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Error struct {
	Code          string          `json:"code"`
	Message       string          `json:"message"`
	ParameterList []ParameterList `json:"parameterList"`
}

type Contact struct {
	CompanyName string `json:"companyName"`
	PersonName  string `json:"personName"`
	PhoneNumber string `json:"phoneNumber"`
}

type ConsolidationDetail struct {
	ConsolidationID string       `json:"consolidationID"`
	EventType       string       `json:"eventType"`
	PackageCount    int64        `json:"packageCount"`
	ReasonDetail    ReasonDetail `json:"reasonDetail"`
	TimeStamp       string       `json:"timeStamp"`
}

type PackageDetails struct {
	ContentPieceCount     string              `json:"contentPieceCount"`
	Count                 string              `json:"count"`
	DeclaredValue         DeclaredValue       `json:"declaredValue"`
	PackageContent        []string            `json:"packageContent"`
	PackagingDescription  ReasonDetail        `json:"packagingDescription"`
	PhysicalPackagingType string              `json:"physicalPackagingType"`
	SequenceNumber        string              `json:"sequenceNumber"`
	UndeliveredCount      string              `json:"undeliveredCount"`
	WeightAndDimensions   WeightAndDimensions `json:"weightAndDimensions"`
}

type PieceCounts struct {
	Count       string `json:"count"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type DeclaredValue struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type ScanEvents struct {
	Date                 string              `json:"date"`
	DelayDetail          DelayDetail         `json:"delayDetail"`
	DerivedStatus        string              `json:"derivedStatus"`
	DerivedStatusCode    string              `json:"derivedStatusCode"`
	EventDescription     string              `json:"eventDescription"`
	EventType            string              `json:"eventType"`
	ExceptionCode        string              `json:"exceptionCode"`
	ExceptionDescription string              `json:"exceptionDescription"`
	ScanLocation         DestinationLocation `json:"scanLocation"`
}

type RequestedAppointmentDetail struct {
	Date   string        `json:"date"`
	Window []Window_sub7 `json:"window"`
}

type DateAndTimes struct {
	DateTime string `json:"dateTime"`
	Type     string `json:"type"`
}

type Contents struct {
	Description      string `json:"description"`
	ItemNumber       string `json:"itemNumber"`
	PartNumber       string `json:"partNumber"`
	ReceivedQuantity string `json:"receivedQuantity"`
}

type SpecialHandlings struct {
	Description string `json:"description"`
	PaymentType string `json:"paymentType"`
	Type        string `json:"type"`
}

type CustomDeliveryOptions struct {
	Description                string                     `json:"description"`
	RequestedAppointmentDetail RequestedAppointmentDetail `json:"requestedAppointmentDetail"`
	Status                     string                     `json:"status"`
	Type                       string                     `json:"type"`
}

type ServiceDetail struct {
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	Type             string `json:"type"`
}

type Window_sub7 struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Window      Window `json:"window"`
}

type ReasonDetail struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type WeightAndDimensions struct {
	Dimensions []Dimensions `json:"dimensions"`
	Weight     []Weight     `json:"weight"`
}

type DeliveryOptionEligibilityDetails struct {
	Eligibility string `json:"eligibility"`
	Option      string `json:"option"`
}

type AdditionalTrackingInfo struct {
	HasAssociatedShipments bool                 `json:"hasAssociatedShipments"`
	Nickname               string               `json:"nickname"`
	PackageIdentifiers     []PackageIdentifiers `json:"packageIdentifiers"`
	ShipmentNotes          string               `json:"shipmentNotes"`
}

type Dimensions struct {
	Height int64  `json:"height"`
	Length int64  `json:"length"`
	Units  string `json:"units"`
	Width  int64  `json:"width"`
}

type ParameterList struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DestinationLocation struct {
	LocationContactAndAddress LocationContactAndAddress `json:"locationContactAndAddress"`
	LocationID                string                    `json:"locationId"`
	LocationType              string                    `json:"locationType"`
}

type HoldAtLocation struct {
	LocationContactAndAddress LocationContactAndAddress_sub21 `json:"locationContactAndAddress"`
	LocationID                string                          `json:"locationId"`
	LocationType              string                          `json:"locationType"`
}

type ServiceCommitMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type SplitShipments struct {
	PieceCount        string `json:"pieceCount"`
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	Timestamp         string `json:"timestamp"`
}

type AvailableImages struct {
	Size string `json:"size"`
	Type string `json:"type"`
}

type DelayDetail struct {
	Status  string `json:"status"`
	SubType string `json:"subType"`
	Type    string `json:"type"`
}

type CompleteTrackResults struct {
	TrackResults   []TrackResults `json:"trackResults"`
	TrackingNumber string         `json:"trackingNumber"`
}

type PackageIdentifiers struct {
	TrackingNumberUniqueID string `json:"trackingNumberUniqueId"`
	Type                   string `json:"type"`
	Value                  string `json:"value"`
}

type Weight struct {
	Unit  string `json:"unit"`
	Value string `json:"value"`
}

type DistanceToDestination struct {
	Units string  `json:"units"`
	Value float64 `json:"value"`
}
