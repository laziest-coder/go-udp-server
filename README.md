## About the repository

This repository is about a very simple service that starts UDP server 
on port written in `.env` file. It receives raw data in bytes that has a certain
structure in json, decodes it and inserts data to an example table `couriers_location`.

## How to start the service

1. Make sure that `.env` is filled correctly
2. `go build` to build an executable file
3. Run executable command `./courier-location-tracker`

## How to check if it is working

1. Open UDP connection with command `nc -u 0.0.0.0 {port}` where `{port}` is filled in `.env` file
2. Send packet `{"courier_id":int,"latitude":float,"longitude":float,"speed":float,"accuracy":float,"azimuth":float}`
3. Check if the record is inserted in table `couriers_location`

> **_NOTE:_** If packet does not match format, service will return error message and record will not be saved!