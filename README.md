# Swagger generated server

Spring Boot Server 


## Overview  
This server was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  
By using the [OpenAPI-Spec](https://github.com/swagger-api/swagger-core), you can easily generate a server stub.  
This is an example of building a swagger-enabled server in Java using the SpringBoot framework.  

The underlying library integrating swagger to SpringBoot is [springfox](https://github.com/springfox/springfox)  

Start your server as an simple java application  

You can view the api documentation in swagger-ui by pointing to  
http://localhost:8080/  

Change default port value in application.properties

### AWS Lambda base event handlder

a sample Lambda...
It can be triggered either by SQS or API Gateway directly.

##### Event Message Flow
![ScreenShot](/handler/flow.png?raw=true "Message Flow")


#### Building

```
go build -o main . && upx main && zip main.zip main
```

#### Testing

```sh
export API_KEY=... CLIENT_ID=... CLIENT_SECRET=...
go test -v .
```

Or create **.env**-file with environment variables.

```ini

# Copy this file to '.env'-file and change values
# ORCID Hub API client credentials:
CLIENT_ID=...
CLIENT_SECRET=...
# UoA API Key:
API_KEY=...

```

