# go-receipt-processor

INSTRUCTIONS

1) Clone (or download) this repo to the machine where the app will be run. Open a terminal and change directory to the project home folder (go-receipt-processor/)

2) Build a docker image
Use below command to build a docker image which contains the code and associated dependencies to run the app.
```
docker build -t fetch-receipt-processor:latest .
```
3) Ensure port 8000 is free since the API endpoints will be exposed on port 8000.

4) Run the app
In the terminal, go to the project home folder (go-receipt-processor/) and run below command to start the app inside a docker container.
```
docker compose up
```