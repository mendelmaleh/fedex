# fedex

Go bindings for the fedex tracking api.

[![Go Documentation](https://godocs.io/git.sr.ht/~mendelmaleh/fedex?status.svg)](https://godocs.io/git.sr.ht/~mendelmaleh/fedex)

Example output:
```go
track.Tracking{
    Number:   "XXXXXXXXXXXX",
    Service:  "FedEx Home Delivery",
    Status:   "Delivered",
    Delivery: time.Date(2022, time.August, 31, 11, 10, 27, 0, time.Local),
    Events:   {
        {
            Status:   "Delivered",
            Location: "Astoria",
            Time:     time.Date(2022, time.August, 31, 11, 10, 27, 0, time.Local),
        },
        {
            Status:   "On FedEx vehicle for delivery",
            Location: "WOODSIDE",
            Time:     time.Date(2022, time.August, 31, 4, 34, 0, 0, time.Local),
        },
        {
            Status:   "At local FedEx facility",
            Location: "WOODSIDE",
            Time:     time.Date(2022, time.August, 31, 4, 21, 0, 0, time.Local),
        },
        {
            Status:   "Left FedEx origin facility",
            Location: "LONG ISLAND CITY",
            Time:     time.Date(2022, time.August, 31, 2, 40, 45, 0, time.Local),
        },
        {
            Status:   "Shipment arriving On-Time",
            Location: "LONG ISLAND CITY",
            Time:     time.Date(2022, time.August, 30, 19, 22, 21, 0, time.Local),
        },
        {
            Status:   "Arrived at FedEx location",
            Location: "LONG ISLAND CITY",
            Time:     time.Date(2022, time.August, 30, 19, 7, 0, 0, time.Local),
        },
        {
            Status:   "Picked up",
            Location: "LONG ISLAND CITY",
            Time:     time.Date(2022, time.August, 30, 12, 42, 0, 0, time.Local),
        },
        {
            Status:   "Dropped off at FedEx OnSite-XXXXXXXXXXXXXXX",
            Location: "BROOKLYN",
            Time:     time.Date(2022, time.August, 29, 16, 9, 0, 0, time.Local),
        },
        {
            Status:   "Shipment information sent to FedEx",
            Location: "",
            Time:     time.Date(2022, time.August, 29, 12, 35, 0, 0, time.Local),
        },
    },
}
```
