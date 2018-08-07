# Customer Profiling Engine Demo

## Faker
This project generates seed data accessible via REST-based API.

### Instructions
```bash
# Get package
go get github.com/epointpayment/mloc-tools/cpe/generator
cd $GOPATH/src/github.com/epointpayment/mloc-tools/cpe/generator

# Get dependencies
go get

# Build exec
go build -o faker

# Create sqlite3 database
touch faker.db

# Seed the database
./faker -seed

# Start Server
./faker
```

### Endpoints

#### Get Customer List
```
/customers/list?start=0&limit=2
```

Output
```JSON
[
  {
    "ID": 1,
    "Email": "james.taylor@example.com",
    "Gender": "male",
    "FirstName": "James",
    "LastName": "Taylor"
  },
  {
    "ID": 2,
    "Email": "ava.johnson@example.com",
    "Gender": "female",
    "FirstName": "Ava",
    "LastName": "Johnson"
  }
]
```


#### Get a Customer's Information
```
/customer/1/info
```

Output
```JSON
{
  "ID": 1,
  "Email": "james.taylor@example.com",
  "Gender": "male",
  "FirstName": "James",
  "LastName": "Taylor"
}
```


#### Get a Customer's Transactions
```
/customer/1/transactions?startDate=20160305&endDate=20160325
```

Output
```JSON
[
   {
      "ID":217,
      "CustomersID":1,
      "DateTime":"2016-03-19T22:27:49-07:00",
      "Description":"Amazon",
      "Credit":451.87,
      "Debit":0,
      "Balance":451.87
   },
   {
      "ID":218,
      "CustomersID":1,
      "DateTime":"2016-03-23T18:07:58-07:00",
      "Description":"J.C. Penny",
      "Credit":0,
      "Debit":212.77,
      "Balance":239.10
   },
   {
      "ID":219,
      "CustomersID":1,
      "DateTime":"2016-03-23T19:11:53-07:00",
      "Description":"Taco Bell",
      "Credit":0,
      "Debit":239.10,
      "Balance":0
   }
]
```

### TOOO
[ ] Simplify installation  
[x] Rounding errors due to floats  
[ ] Better error handling  
[ ] Cleanup and comment code  