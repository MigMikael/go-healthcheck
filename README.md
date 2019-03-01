# HealthCheck
Line Challenge for Golang Developer Candidate


## Getting Started
recommended to turn on DEBUG mode by go to /go-healthcheck/main.go and change value of DEBUG variable to `true`
```
$ cd go-healthcheck
$ go install -v .
$ go-healthcheck <path to csv file>
```


this will begin check URL contain in csv file 
### Example summary report

```
Checked webistes: 25
Successful websites: 21
Failure websites: 1
Total times to finished checking website: 2.464820383s
```


## Running The Tests

```
$ go test

Perform website checking...
https://blognone.com => 200 OK
https://httpstat.us/400 => 400 Bad Request
http://yfrog.com/17toj => 200 OK
http://facebook.com => 200 OK
http://yfrog.com/5b59688defbb5f1f3f8ce04bcj => 200 OK
Done!

PASS
ok  	go-healthcheck	3.434s
```

## Test Set 
I've created csv test set in vairous size 
* test.csv => normal test set consist of 5 URL
* tiny-test.csv => vary small set consist of 25 Record, some record is not real URL
* small-test.csv => small set consist of 40 URL
* medium-test.csv => medium set consis of 100 URL
* large-test.csv => consis of 1,721 URL 

## Performance
I've test performance by using each test set above on my Macbook Pro 2015 (Dual Core i5 CPU Ram 8 GB)
| Test Set | Time |
| :---: | :---: |
| test.csv | 2.54 s |
| tiny-test.csv | 3.50 s|
| small-test.csv | 30.00 s |
| medium-test.csv | 4m31s |
| large-test.csv | 00 |


## Task List
- [x] Check website available with http library
- [x] Send the report via Healcheck Report API
- [x] Use `Line Login` to request API token
- [x] Show the summary as the example above at the end
- [x] Handle errors without stopping the entire process
- [x] Write readable code
- [x] Write automated tests


## Bonus List 
- [x] Run program as fast as possible on multi-core CPU