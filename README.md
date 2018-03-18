# Spain BOE reader

 Request the [boe.es](http://boe.es/) page to retrieve the info for the publication of the specified date

 ## Build
 To build it you need at least `go 1.8` and `GOPATH` set
 ```sh
 go get https://github.com/publicgov/spain-boe-reader
 ```
 When done, go inside the folder and install it
 ```sh
 cd spain-boe-reader && go install
 ```

 ## Execute
 You can view the list of default options passing the `-h` flag
 ```sh
spain-boe-reader -h
Usage of spain-boe-reader:
  -date string
        BOE publication date in format YYYYMMDD (default "20180318")
```
As you can see, if no arguments found, the program use for default the current date