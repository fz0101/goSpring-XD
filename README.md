# goSpring-XD
Gets around WebHDFS, by created a sink API

Use Case:

Need to access data from a data lake, that lives in a kerberized hadoop system. 
Spring XD ingest data from sources, in the processor step, an HTTP client can post to an external API. 
This API will take a Post with a JSON file. 
It stores the file inside the API using a virtual, and in memory file system. 
A get command will be added to expose the JSON file
HTTPS and Oatuh are baked in and supported for accessing the API. 
Includes stats and internal logging that are baked in. 


Features to be added. 
Support for other data formats that are supported in PHD. 
Docker support
Pivotal Cloud Foundry support. 


API - Built in GoLang


