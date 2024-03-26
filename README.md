# go-receipt-processor

### INSTRUCTIONS

1) Clone (or download) this repo to the machine where the app will be run. Open a terminal and change directory to the project home folder (go-receipt-processor/)

2) Build a docker image\
Ensure docker daemon is runnning on the machine.\
Use below command to build a docker image which contains the code and associated dependencies to run the app.
```
docker build -t go-receipt-processor:latest .
```
3) Ensure port 8000 is free since the API endpoints will be exposed on port 8000.

4) Run the app\
In the terminal, go to the project home folder (go-receipt-processor/) and run below command to start the app inside a docker container.
```
docker compose up
```

5) Once the app starts, the below message will be printed on the terminal:\
Receipt server running on port:  :8000

6) The app is now ready to serve API requests.

### ASSUMPTIONS

- The purchase date is valid and in YYYY-MM-DD format
- The purchase time is valid in 24-hour format
- The prices are positive numbers