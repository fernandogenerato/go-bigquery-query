
**Golang BigQuery Query Project**

This is a simple Golang project demonstrating how to perform queries against BigQuery, Google Cloud Platform's (GCP) data warehousing and analytics service. This project can serve as a starting point for integrating BigQuery queries into your own Go applications.

### Prerequisites

Before running this project, you'll need the following set up:

1.  Google Cloud Platform (GCP) account.
2.  Project created on GCP.
3.  BigQuery API enabled for the project.
4.  Service credentials set up for authentication.

### Configuration

1.  Clone this repository:
       
    `git clone https://github.com/fernandogenerato/go-bigquery-example.git` 
    
2.  Install dependencies:
        
    `go mod tidy` 
    
3.  Configure your GCP credentials. You can do this by setting the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to the path of your credentials JSON file:
        
    `export GOOGLE_APPLICATION_CREDENTIALS=/path/to/your/credentials.json` 
    

### Usage

The example in this project demonstrates how to make a simple query to BigQuery. To modify the query or add more functionalities, you can edit the `main.go` file.

To run the example, use the following command:

`go run main.go`