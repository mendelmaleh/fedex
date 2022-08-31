package fedex

import "git.sr.ht/~mendelmaleh/track"

func (r *TrackResponse) Tracking() (t track.Tracking) {
	if len(r.Output.CompleteTrackResults) < 1 {
		return
	}

	ctr := r.Output.CompleteTrackResults[0]
	t.Number = ctr.TrackingNumber

	if len(ctr.TrackResults) < 1 {
		return
	}

	tr := ctr.TrackResults[0]
	t.Service = tr.ServiceDetail.Description

	for _, v := range tr.DateAndTimes {
		if v.Type == "ESTIMATED_DELIVERY" {
			t.Delivery = v.DateTime.Time
			break
		}
	}

	t.Events = make([]track.Event, len(tr.ScanEvents))
	for i, v := range tr.ScanEvents {
		t.Events[i] = track.Event{
			Status:   v.EventDescription,
			Location: v.ScanLocation.City,
			Time:     v.Date.Time,
		}
	}

	return
}

// generated with github.com/mendelmaleh/gojson

type TrackResponse struct {
	CustomerTransactionID string `json:"customerTransactionId"`
	Output                Output `json:"output"`
	TransactionID         string `json:"transactionId"`
}

type Output struct {
	Alerts               string                `json:"alerts"`
	CompleteTrackResults []CompleteTrackResult `json:"completeTrackResults"`
}

type CompleteTrackResult struct {
	TrackResults   []TrackResult `json:"trackResults"`
	TrackingNumber string        `json:"trackingNumber"`
}

type TrackResult struct {
	AdditionalTrackingInfo        AdditionalTrackingInfo `json:"additionalTrackingInfo"`
	AvailableImages               []AvailableImage       `json:"availableImages"`
	AvailableNotifications        []string               `json:"availableNotifications"`
	ConsolidationDetail           []ConsolidationDetail  `json:"consolidationDetail"`
	CustomDeliveryOptions         []CustomDeliveryOption `json:"customDeliveryOptions"`
	DateAndTimes                  []DateAndTime          `json:"dateAndTimes"`
	DeliveryDetails               DeliveryDetails        `json:"deliveryDetails"`
	DestinationLocation           Location               `json:"destinationLocation"`
	DistanceToDestination         DistanceToDestination  `json:"distanceToDestination"`
	Error                         Error                  `json:"error"`
	EstimatedDeliveryTimeWindow   TimeWindow             `json:"estimatedDeliveryTimeWindow"`
	GoodsClassificationCode       string                 `json:"goodsClassificationCode"`
	HoldAtLocation                HoldAtLocation         `json:"holdAtLocation"`
	InformationNotes              []InformationNote      `json:"informationNotes"`
	LastUpdatedDestinationAddress Address                `json:"lastUpdatedDestinationAddress"`
	LatestStatusDetail            LatestStatusDetail     `json:"latestStatusDetail"`
	MeterNumber                   string                 `json:"meterNumber"`
	OriginLocation                Location               `json:"originLocation"`
	PackageDetails                PackageDetails         `json:"packageDetails"`
	PieceCounts                   []PieceCount           `json:"pieceCounts"`
	ReasonDetail                  ReasonDetail           `json:"reasonDetail"`
	RecipientInformation          ContactAndAddress      `json:"recipientInformation"`
	ReturnDetail                  ReturnDetail           `json:"returnDetail"`
	ScanEvents                    []ScanEvent            `json:"scanEvents"`
	ServiceCommitMessage          ServiceCommitMessage   `json:"serviceCommitMessage"`
	ServiceDetail                 ServiceDetail          `json:"serviceDetail"`
	ShipmentDetails               ShipmentDetails        `json:"shipmentDetails"`
	ShipperInformation            ContactAndAddress      `json:"shipperInformation"`
	SpecialHandlings              []SpecialHandling      `json:"specialHandlings"`
	StandardTransitTimeWindow     TimeWindow             `json:"standardTransitTimeWindow"`
	TrackingNumberInfo            TrackingNumberInfo     `json:"trackingNumberInfo"`
}

// AdditionalTrackingInfo
type AdditionalTrackingInfo struct {
	HasAssociatedShipments bool                `json:"hasAssociatedShipments"`
	Nickname               string              `json:"nickname"`
	PackageIdentifiers     []PackageIdentifier `json:"packageIdentifiers"`
	ShipmentNotes          string              `json:"shipmentNotes"`
}

type PackageIdentifier struct {
	TrackingNumberUniqueID string `json:"trackingNumberUniqueId"`
	Type                   string `json:"type"`
	Value                  string `json:"value"`
}

// AvailableImage
type AvailableImage struct {
	Size string `json:"size"`
	Type string `json:"type"`
}

// ConsolidationDetail
type ConsolidationDetail struct {
	ConsolidationID string       `json:"consolidationID"`
	EventType       string       `json:"eventType"`
	PackageCount    int64        `json:"packageCount"`
	ReasonDetail    ReasonDetail `json:"reasonDetail"`
	TimeStamp       string       `json:"timeStamp"`
}

type ReasonDetail struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

// CustomDeliveryOption
type CustomDeliveryOption struct {
	Description                string                     `json:"description"`
	RequestedAppointmentDetail RequestedAppointmentDetail `json:"requestedAppointmentDetail"`
	Status                     string                     `json:"status"`
	Type                       string                     `json:"type"`
}

type RequestedAppointmentDetail struct {
	Date   string       `json:"date"`
	Window []TimeWindow `json:"window"`
}

type TimeWindow struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Window      Window `json:"window"`
}

type Window struct {
	Begins string `json:"begins"`
	Ends   string `json:"ends"`
}

// DateAndTime
type DateAndTime struct {
	DateTime DateISO `json:"dateTime"`
	Type     string  `json:"type"`
}

// DeliveryDetails
type DeliveryDetails struct {
	ActualDeliveryAddress             Address                           `json:"actualDeliveryAddress"`
	DeliveryAttempts                  string                            `json:"deliveryAttempts"`
	DeliveryOptionEligibilityDetails  []DeliveryOptionEligibilityDetail `json:"deliveryOptionEligibilityDetails"`
	DeliveryToday                     bool                              `json:"deliveryToday"`
	DestinationServiceArea            string                            `json:"destinationServiceArea"`
	DestinationServiceAreaDescription string                            `json:"destinationServiceAreaDescription"`
	LocationDescription               string                            `json:"locationDescription"`
	LocationType                      string                            `json:"locationType"`
	OfficeOrderDeliveryMethod         string                            `json:"officeOrderDeliveryMethod"`
	ReceivedByName                    string                            `json:"receivedByName"`
	SignedByName                      string                            `json:"signedByName"`
}

type Address struct {
	City                string   `json:"city"`
	Classification      string   `json:"classification"`
	CountryCode         string   `json:"countryCode"`
	CountryName         string   `json:"countryName,omitempty"` // modified, not necessarily always present
	PostalCode          string   `json:"postalCode"`
	Residential         bool     `json:"residential"`
	StateOrProvinceCode string   `json:"stateOrProvinceCode"`
	StreetLines         []string `json:"streetLines"`
	UrbanizationCode    string   `json:"urbanizationCode"`
}

type DeliveryOptionEligibilityDetail struct {
	Eligibility string `json:"eligibility"`
	Option      string `json:"option"`
}

// Location
type Location struct {
	ContactAndAddress ContactAndAddress `json:"locationContactAndAddress"`
	ID                string            `json:"locationId"`
	Type              string            `json:"locationType"`
}

type ContactAndAddress struct {
	Address Address `json:"address"`
	Contact Contact `json:"contact"`
}

type Contact struct {
	CompanyName string `json:"companyName"`
	PersonName  string `json:"personName"`
	PhoneNumber string `json:"phoneNumber"`
}

// DistanceToDestination
type DistanceToDestination struct {
	Units string  `json:"units"`
	Value float64 `json:"value"`
}

// Error
type Error struct {
	Code          string      `json:"code"`
	Message       string      `json:"message"`
	ParameterList []Parameter `json:"parameterList"`
}

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// HoldAtLocation
type HoldAtLocation struct {
	LocationContactAndAddress ContactAndAddress `json:"locationContactAndAddress"`
	LocationID                string            `json:"locationId"`
	LocationType              string            `json:"locationType"`
}

// InformationNote
type InformationNote struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// LatestStatusDetail
type LatestStatusDetail struct {
	AncillaryDetails []AncillaryDetail `json:"ancillaryDetails"`
	Code             string            `json:"code"`
	DelayDetail      DelayDetail       `json:"delayDetail"`
	DerivedCode      string            `json:"derivedCode"`
	Description      string            `json:"description"`
	ScanLocation     Address           `json:"scanLocation"`
	StatusByLocale   string            `json:"statusByLocale"`
}

type AncillaryDetail struct {
	Action            string `json:"action"`
	ActionDescription string `json:"actionDescription"`
	Reason            string `json:"reason"`
	ReasonDescription string `json:"reasonDescription"`
}

type DelayDetail struct {
	Status  string `json:"status"`
	SubType string `json:"subType"`
	Type    string `json:"type"`
}

// PackageDetails
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

type DeclaredValue struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type WeightAndDimensions struct {
	Dimensions []Dimensions `json:"dimensions"`
	Weight     []Weight     `json:"weight"`
}

type Dimensions struct {
	Height int64  `json:"height"`
	Length int64  `json:"length"`
	Units  string `json:"units"`
	Width  int64  `json:"width"`
}

type Weight struct {
	Unit  string `json:"unit"`
	Value string `json:"value"`
}

// PieceCount
type PieceCount struct {
	Count       string `json:"count"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// ReturnDetail
type ReturnDetail struct {
	AuthorizationName string         `json:"authorizationName"`
	ReasonDetail      []ReasonDetail `json:"reasonDetail"`
}

// ScanEvent
type ScanEvent struct {
	Date                 DateISO     `json:"date"`
	DelayDetail          DelayDetail `json:"delayDetail"`
	DerivedStatus        string      `json:"derivedStatus"`
	DerivedStatusCode    string      `json:"derivedStatusCode"`
	EventDescription     string      `json:"eventDescription"`
	EventType            string      `json:"eventType"`
	ExceptionCode        string      `json:"exceptionCode"`
	ExceptionDescription string      `json:"exceptionDescription"`
	ScanLocation         Address     `json:"scanLocation"`
}

// ServiceCommitMessage
type ServiceCommitMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// ServiceDetail
type ServiceDetail struct {
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	Type             string `json:"type"`
}

// ShipmentDetails
type ShipmentDetails struct {
	BeforePossessionStatus bool            `json:"beforePossessionStatus"`
	ContentPieceCount      string          `json:"contentPieceCount"`
	Contents               []Content       `json:"contents"`
	SplitShipments         []SplitShipment `json:"splitShipments"`
	Weight                 []Weight        `json:"weight"`
}

type Content struct {
	Description      string `json:"description"`
	ItemNumber       string `json:"itemNumber"`
	PartNumber       string `json:"partNumber"`
	ReceivedQuantity string `json:"receivedQuantity"`
}

type SplitShipment struct {
	PieceCount        string `json:"pieceCount"`
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	Timestamp         string `json:"timestamp"`
}

// SpecialHandling
type SpecialHandling struct {
	Description string `json:"description"`
	PaymentType string `json:"paymentType"`
	Type        string `json:"type"`
}

// TrackingNumberInfo
type TrackingNumberInfo struct {
	CarrierCode            string `json:"carrierCode"`
	TrackingNumber         string `json:"trackingNumber"`
	TrackingNumberUniqueID string `json:"trackingNumberUniqueId"`
}
