# Is today a workday?
The API can tell you whether a date is a workday.  

### Endpoint
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
