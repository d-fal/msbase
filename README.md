# General Rest Framework 


## Introduction

This structure invokes microservice paradigms to deliver a restful service.

* ### *Caveat:*
    * This project uses an identity.yaml file that determines the function of the app. Developers should use this file to work with this framework.
    * The project structure _should not be changed_ under any circumstances.

### Project Structure

* Project skeleton is as below:
    * `<codeBaseName>/`
        * assets
            * <sub>This folder hosts the static files needed for this project</sub>
        * cmd
            * <sub>main package lies here. You will find the start point of the project here.</sub> 
        * internal
            * <sub>internal package that includes fundamental packages of the framework.</sub>
        * pkg
            * <sub>this package encompasses common packages</sub>
        * test
        * vendor
        * config
            * This folder aggregates all the configs of the project.
### Microservice Architecture
![Microservice architecture](/assets/ref/xenonstack-what-are-microservices.png)
### Project Identity
> The identity.yaml file determines the future functionality of the project.
```
App: # Do not change this
  Name: Billing Inquiry # sample name
  ID: Billing 
  Version: 1.000
  BaseURL: /api/{{.version}}/billing/
  Routes:
    -
      URL: facade
      ID: internal_rec_billing_facade
      Description: prepares facade
      Method: GET
      Middleware:
        -
          ID: cors
    -
      URL: debtinquiry
      ID: internal_rec_billing_debtinquiry
      Description:
      Method: POST
      Middleware:
        -
          ID: cors
        -
          ID: authenticator

```
* App
    * **Name** : This is what you call this app.
    * **ID** : The ID of the app that will be used in the future.
    * **Version** : Tell the framework on which version of the app you are working on. The versioning should be in full compliance with the following convention:
    > Version: **Framework Version** - **Contribution Version**.*Major Version*.*Minor Version*

    For instance: ```v1-1.0.1``` means the framework version is 1 and the contribution is on 1.0.1
    * **Middleware** : this array tells the framework from which middleware should it traverse the incoming requests. This block holds the following parameters:
        * ***ID*** : The middleware ID that framework understands it.
        * ***URL*** : The url that the framework serve.
    > Recommended URL : `/<frameworkVersion>/<middlewareID>/<AppID>`
    * **Routes** :
        * You can add route blocks as your microservice needs.
            * **URL** : the address that specifies your receptor.
            * **ReceiverID** : The ID that signifies microservice a.k.a. **Receiver**.
            * **Description** : Any commentaries conducive grasp receiver's functionality.
            * **Method** : _http_ method that receiver responds to. It can be one of the following: `POST, GET, PUT, INSERT, DELETE`.
### Middlewares:
The available middlewares implemented in this framework are as follows.
### Contribute to this project :


### Developers can add microservices to this project as follows:
    1. Creating App Identity file
        This file defines the way app behaves. 
    2. Choosing the right moiddleware amongst the tuned ones.
    3. Creating microservice of ineterest.
    4. Setup microservice handlers.
    5. Set project router to use the handlers.

To help understand the abovesaid steps, a sample microservice called **sample** is being added to this project.

1. Creating identity file under **config/app_params.yaml**. In This example, we are going to make a project that serves the bellow URL:
> /api/v1/sample/testhandler

The **APP** part of **app_params.yaml** would look like this:
```
App: # Do not change this
  Name: Hello Microservice # sample name
  ID: Microservice # No whitespaces are allowed
  Version: 1.000
  BaseURL: /api/{{.version}}/hello/
  Routes:
    -
      URL: test
      HandlerID: internal_rec_hello_test # Go to internal/routing/handlers.go and pair this ID with the existing one in the code
      Description: test handler
      Method: GET
      Middleware:
        -
          ID: cors
        -
          ID: authenticator
    -
      URL: token/refresh
      HandlerID: internal_rec_get_auth_token
      Description: token generator
      Method: GET
      Middleware:
        -
          ID: cors
```

2. Choosing the proper middlewares from the existing ones. As you can see from the above example, we have chosen **cors**.

3. Register your handler(s). To do so, open ```internal/routing/init.go``` and add the *ID* of your handler. It is recommended to name your hanlder ID in the below convention and refrain from using camelcase letters.
   
    ```internal_rec_<AppID>_<hanlder Name>``` 

  > example: internal_rec_sample_samplehandler

for instance, see the below example of *internal/routes/init.go* 

    package routing

    type HandlersSet struct {
      HandlerSampleID          string
    }

    func init() {
      handlersSet = HandlersSet{       
        "internal_rec_sample_samplehandler",
      }
    }


4. Create your microservice package in ```internal/microservice/<microservice name>/<version descriptor>/<microsservice name>/services/services.go```
  
  > example: ```internal/microservice/sample/v1/services/services.go```
    
    package services 

5. Create your handlers 

> example: ```internal/microservice/sample/v1/sample/services/handlers.go```

    pacakge sample

    import (
    "net/http"
    "github.com/gin-gonic/gin"
    )

    func (handler *ServiceHandler) SampleHandler(c *gin.Context) {
        c.JSON(
          http.StatusOK,
          gin.H{
            "result": "Hello world!",
          },
        )
      }

6. Set your handler in ```internal/routing/handlers.go``` and add your newly registered handler. This file should look like as follows:

        package routing

        import (
          "errors"
          "fmt"

          sample "780.com/msbase/internal/microservice/sample/services"
          "github.com/gin-gonic/gin"
        )

        func getRequestHandler(handlerID string) (gin.HandlerFunc, error) {
        
          sampleHandler := sample.GetHandler()

          switch handlerID {

          case sample.HandlerSampleID:
            return sample.SampleHandler, nil
          }

          return nil, errors.New("An error")
        }

Fin!