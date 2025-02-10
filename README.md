# fetch-Receipt-processor-API-th-project
# Receipt Processor API

This is a Go-based web service that processes receipts and calculates points based on specific rules.

## Features
- **Process Receipts**: Submit a receipt and get a unique ID.
- **Get Points**: Retrieve points for a receipt using its ID.

## Rules for Points Calculation
- One point for every alphanumeric character in the retailer name.
- 50 points if the total is a round dollar amount.
- 25 points if the total is a multiple of 0.25.
- 5 points for every two items on the receipt.
- Additional points for item descriptions with lengths divisible by 3.
- 6 points if the purchase day is odd.
- 10 points if the purchase time is between 2:00 PM and 4:00 PM.

  
## How to Run
1. **Clone the repository:**
   ```bash
   git clone https://github.com/saichandreddy00/fetch-Receipt-processor-API-th-project.git
   cd receipt-processor
2. **Run the application:**
- go run main.go

3.**The API will be available at** 
- http://localhost:8080.

**API Endpoints:**
- POST /receipts/process: Submit a receipt.
- GET /receipts/{id}/points: Retrieve points for a receipt.





