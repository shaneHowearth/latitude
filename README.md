# latitude
Small Profit searcher

# Install
Clone this repository with `git clone` or `go get -u`
`make` will build the application assuming you have all the dependencies.

# Usage
The built application will accept a space or comma sparated list of numbers and will produce appropriate output.
```
e.g. ./latitude 1,2,3,4,5
Latitude max sharetrading profit calculator
Calculating for: [1 2 3 4 5]
Buy at:  1
Sell at:  5
Profit:  4

or
./latitude 1 2 3 4 5
Latitude max sharetrading profit calculator
Calculating for: [1 2 3 4 5]
Buy at:  1
Sell at:  5
Profit:  4
```
