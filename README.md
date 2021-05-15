# DataPipe

    Repo for Data PipeLining and ETL Process 

**Intro to Project Repo**

This data pipeline ingest raw data from different sources(for now, only one upload folder)
and move the data to a destination for storage and analysis. This pipeline will include filtering 
and features that provide resiliency against failure.

---------------------------------------------------------
 
**Technologies Used**
    
    FrontEnd    :   Angular 11
    BackEnd     :   Golang, Python
    DataBase    :   MongoDB
    Other       :   Apache AirFlow, Docker

_Need to install above for setup of the project. Setup steps are mentioned inside each components._

---------------------------------------------------------

**Components in the Repo**     

**_1. AngForDocker_**

This is the FrontEnd GUI for Client/Users. This is built on Angular 11. User will be able to see the files uploaded,
processed, analyzed along with its status. The Metadata of files will be stored in the 
MongoDB database.

**_2. BackSideServer_**

This is the BackEnd for the DataPipe repo. This server is developed on Golang and will be developed to
serve as a service in microservice architecture(Need to work on it later). This server connects to MongoDB 
Database and updates/creates/delete/insert the metadata of files uploaded/sent.

**_3. PipeLine_**

This is also the backend portion of the DataPipe repo. This component will extract/transform/load the uploaded files 
or any data on the source. The sole purpose of this component will be to perform data engineering on data/files and 
load it to the Destination. 

----------------------------------------------------------
