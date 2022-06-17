## About the service

`Courier Location Tracker` is a very simple service that starts UDP server 
on port written in `.env` file. It receives data in raw bytes that has a certain
structure in json, decodes it and inserts data to table `ex24_drivers_location`.
Other details can be found in [ClickUp Docs](https://app.clickup.com/3609340/v/dc/3e4qw-1174/3e4qw-14663).

## How to start the service

1. Make sure that `.env` is filled correctly
2. `go build` to build an executable file
3. Run executable command `./courier-location-tracker`

## How to check if it is working

1. Open UDP connection with command `nc -u 0.0.0.0 {port}` where `{port}` is filled in `.env` file
2. Send packet `{"courier_id":int,"latitude":float,"longitude":float,"speed":float,"accuracy":float,"azimuth":float}`
3. Check if the record is inserted in table `ex24_drivers_location`

> **_NOTE:_** If packet does not match format, service will return error message and record will not be saved!