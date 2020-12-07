# Golang Bootcamp

First Deliverable (due November 22th 23:59PM)
Based on the self-study material and mentorship covered until this deliverable, I have perform the following:

 - Select architecture
 - Read a CSV
 - Handle Errors for CSV  (not valid, missing)
 - Use best practices
 - Use lint, error check

## Run

Clone the repository and execute the following command on the root folder.

``` bash
git clone https://github.com/ruvaz/golang-bootcamp-2020.git
cd golang-bootcamp-2020
go build . 
./golang-bootcamp-2020
``` 

## Selected architecture

Since the applications made in Go must be applications that are characterized by the speed and simplicity in their code as well as the low level of depth, I found that the clean architecture fits very well with the way of working with Go with its separation of each layer. By layering the software and adhering to the dependency rule, I will create a system that is easily testable, with all the benefits that come with such as when some of the external parts of the system become obsolete, such as the database or the web framework. , you can replace those outdated items with a minimum of effort.

```text
    .  
    ├── domain  
    │ └── model  
    │ │ └── student.go  
    ├── infrastructure  
    │ ├──  datastore  
    │ │ └── dataFile.csv  
    ├── interface  
    │ ├── controller  
    │ │ ├── student_controller.go 
    │ ├── presenter  
    │ │ └── user_presenter.go   
    ├── main.go  
```

## Files

 - dataFile.csv  Sample file to be displayed  in a console by Go.