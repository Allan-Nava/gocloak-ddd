# GoCloak Domain Driven Design Api Gateway


## api overview

every api (for each domain model) is structured like below:

    handler <----rest api endpoint
    |
    service <----service layer, main business logic of the api lives here
    |
    store <---- database logic, operation related to database fetching

