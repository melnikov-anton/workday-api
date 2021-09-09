# Is today a workday?
The API can tell you whether a date is a workday.  

## How to run
#### In Docker
1) Build image
```bash
docker build -t workday-api-image .
```
2) Start container
```bash
docker run -p 8000:80 --name workday-api-app workday-api-image
```
App will be available on [http://localhost:8000/](http://localhost:8000/).  
#### On Linux or MacOS
1) Build the app and distribution folder
```bash
./gobuild.sh
```
Script will create **dist** folder which contains app.  
2) Start app from **dist** folder
```bash
./workday-api -port 8000
```
App will be available on [http://localhost:8000/](http://localhost:8000/).  
Without specifying port, app will start on port 8080.  


## How to use
Use path **/api/{country_code}/workday/{date}**, where:  
**country_code** - two letters country code;  
**date** - date in format YYYY-MM-DD (or word **today**),  
and you get an answer like:
```json
{
  "date": "2021.09.06",
  "is_workday": true
}
```
