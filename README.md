# ZipCode of Mexico Golang for the Bootcamp  
  
## Introduction  
  
The following project is designed to obtain the postal data of the official post office of Mexico in csv format and then store the information in a database in MongoDB and proceed to consult the colonies referring to a postal code as a result in json  
  
## Requirements  
  
You need to install MongoDB either with the normal [installer](https://docs.mongodb.com/manual/installation/) or in docker and run it on 

> port 27017

Docker

    docker pull mongo
    

## Run the code

On the path go to the folder /cmd.

`cd cmd`

Then run

    go install

And finally

     go run main.go

  

## About the project

To be able to consult the data, you must first fill the tables, so you must execute a get-type request to the method:

    localhost:8080/api/v1/direcciones/populate

*What happens behind the scenes is that the code makes a request to a government endpoint which returns a csv with all the zip codes, a 15mb file _(145,110 approximate zip codes)_    
    
Once you get the response of the inserted ids that will take approximately 15 seconds on complete, you can now check the colonies that belong to a postal code with the method

    localhost:8080/api/v1/direcciones/search/:zipCode

For example:

    localhost:8080/api/v1/direcciones/search/97306

And an example of the previous request would be the following:


 

    [
      {
        "Id": "56eb060711d84d67ab164e862c5eb0f1",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Chichi Suárez"
      },
      {
        "Id": "cd9c02f9983e4667ae3f730615263907",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Sitpach"
      },
      {
        "Id": "d24212569c7f42df9cb6b1b0f54f913d",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Villas de Oriente"
      },
      {
        "Id": "35467b546c21499e9870c4861fb8257d",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Los Héroes"
      },
      {
        "Id": "7e2386368a624f8c88e1ba827865ce7b",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Santa María Chí"
      },
      {
        "Id": "eee9a4cc733543d7a6d753b2ce40b5ab",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Chichi Díaz"
      }
    ]

## Unit test

Two unit tests were created that test these two end-point on the path

> /golang-bootcamp-2020/pkg/util/zipcodes_test.go

First run the unit test 

> Test_getCSVCodes()

And then 

> Test_searchZipCodes()

Another unit test was created to test the correct flushing of the zipcodes table, run it if you want

> Test_dropZipCodes()
