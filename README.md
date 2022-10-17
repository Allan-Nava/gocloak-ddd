# GoCloak Domain Driven Design Api Gateway
[![Go Report Card](https://goreportcard.com/badge/github.com/Allan-Nava/gocloak-ddd)](https://goreportcard.com/report/github.com/Allan-Nava/gocloak-ddd)
[![GoDoc](https://godoc.org/github.com/Allan-Nava/gocloak-ddd?status.svg)](https://godoc.org/github.com/Allan-Nava/gocloak-ddd)



## api overview

every api (for each domain model) is structured like below:

    handler <----rest api endpoint
    |
    service <----service layer, main business logic of the api lives here
    |
    store <---- database logic, operation related to database fetching

