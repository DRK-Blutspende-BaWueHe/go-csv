# go-csv
CSV and generic text parser lib

Golang library for handling csv and text files

###### Install
`go get github.com/DRK-Blutspende-BaWueHe/go-csv`

## Features
  - Provides a Marshal/Unmarshal function
  - Supports Encoding (UTF8, ASCI, WIN1254....)
  - Supports Timezone Encoding

## Quick start
Given an input file like 
```
Lorem;ipsum;;dolor;2022-01-53 15:00;15
```
And a target structure
```
type record struct {
    field1 string `csv:"1"`
    field2 string `csv:"2"`
    field3 string `csv:"3"`
    field4 string `csv:"4"`
    field5 time.Time `csv:"5,longdate"`
    field6 int `csv:"6"`
} 
```


### Reference
#### Annotations
  - longdate Date in the format YYYY-MM-DD
  - shortdate Date in the format YYYY-MM-DD hh:mm:ss
  - datefmt Provide a date format