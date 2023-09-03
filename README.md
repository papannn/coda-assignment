
# Coda Payment Technical Interview Assignment

Goal: Write an Round Robin API which receives HTTP POSTS and routes them to one of a list of
Application APIs

I'm using golang version `1.21.0`

# High Level Design

There are 3 services that I will use for the current requirement:
* API Gateway Service
* Discovery Service
* Service A

**API Gateway Service** purpose is to Route the **Client** request to the designated **Service A**

**Discovery Service** purpose is to store list of IP and Ports of the **Service A** that already connected to **Discovery Service**, and then it will return the IP + Ports by choosing one of the **Service A** IP + Ports using specific algorithm such as **Round Robin Algorithm** when the **API Gateway Service** call the `/api/lookup` API

**Service A** purpose is to return some Business Logic value that will be useful for the Client side

![service-A-connect-discovery-service](https://res.cloudinary.com/dzfbqsm39/image/upload/v1693311765/coda-payment/xq4owvnkofsngghelmmv.png)

At the first of **Service A** boot up the application, **Service A** will register itself to the **Discovery Service** by using `/api/register/` endpoint with this kind of payload:

    {
        "namespace": "service A",
        "ip": "localhost",
        "port": "8080"
    }

After that, **Discovery Service** will store the designated data to their variable.

![discovery-service-scheduler](https://res.cloudinary.com/dzfbqsm39/image/upload/v1693311765/coda-payment/ypp7oqco8vel24es7q6p.png)

Every `x` second (we can set this value on **Discovery Service** config on `health_check_time_interval` value json), **Discovery Service** will try to do the health checkup for every Service that already register itself to **Discovery Service**.

**Discovery Service** will check some metrics such as:
* Liveness of the service
*  ̶H̶o̶w̶ ̶l̶o̶n̶g̶ ̶t̶h̶e̶ ̶S̶e̶r̶v̶i̶c̶e̶ ̶A̶ ̶r̶e̶s̶p̶o̶n̶d̶s̶

If the metrics is bad, for example **Service A** can't return 200 success response on `/health_check` or the duration of returning back the response exceed the threshold that we already set in **Discovery Service** config on `timeout_threshold` json value. **Discovery Service** will update the status of that service into temporary unavailable, so the service will not get picked by Load Balancing Algorithm.

![client-hitting-api-flow](https://res.cloudinary.com/dzfbqsm39/image/upload/v1693311765/coda-payment/rhzv8oocctn91qqax4dm.png)

The flow is very straight-forward:
1. Client try to hit **Service A** via **API Gateway Service**
2. **API Gateway Service** will try to hit **Discovery Service**'s `/api/lookup` with this kind of payload. `{"namespace": "service A"}`
3. **Discovery Service** returns the value of the IP + Ports the can be hit.
4. Using the value that already got from Step 3, **API Gateway** try to hit **Service A** Port : 8080
5. (Edge case): It seems **Service A** Port : 8080 got timeout error
6. **API Gateway** will ask **Discovery Service** again which **Service A** port that available to hit
7. **Discovery Service** returns the value of the IP that can be hit, ex: Service A : 8081
8. Using the value that already got from Step 7, **API Gateway** try to hit **Service A** Port : 8081
9. **Service A** Port : 8081 now returns http 200 success with some payload from their Business Logic
10. **API Gateway** forward the result of **Service A** Port : 8081 to the client

# Folder Structure

    project
        |-api           # API Contract struct
        |-cmd           # for putting main.go file / func main() function
        |-config        # Config struct and config.json file 
        |-domain        # Struct that will be used for repository
        |-handler       # Endpoint list, putting the `main` class
        |-logic         # Business logic layer, only got called by `main` class
        |-repository    # Data persistence layer, only got called by logic layer
        |-scheduler     # Scheduler function
        |-server        # Initializing the Endpoint Registering, Config and Dependency injection

Basically, the flow of the application will be like this

`Client` -> `Handler` -> `Logic` -> `Repository` -> `Domain`

# How to run
### Discovery Service:

On `coda-assignment` folder
    
    cd discovery-service
    make run-local


### API Gateway Service

On `coda-assignment` folder

    cd api-gateway
    make run-local

### Service A

On `coda-assignment` folder

    cd service-a
    make run-local

**Note**: For `Service A`, you can run multiple instances with just running that above command on different terminal, don't worry, `Service A` will try to use another Port automatically if the current port on config already used.

    2023/09/04 00:30:43 Running on address: localhost:8000
    2023/09/04 00:30:43 Port: 8000 already been used, trying to use port 8001
    2023/09/04 00:30:43 Running on address: localhost:8001
    2023/09/04 00:30:43 Port: 8001 already been used, trying to use port 8002
    2023/09/04 00:30:43 Running on address: localhost:8002

**Note**: Be sure to run `Discovery Service` first before running `Service A`, because `Service A` can't hit `/api/register/` API if `Discovery Service` is not yet ready.

# Config

### API Gateway Service

    {
        "ip": "localhost",
        "port": "5555",
        
        # Change this if you change the IP & Port on discovery service config
        "discovery_service_base_url": "http://localhost:4444"
    }

### Discovery Service

    {
        "ip": "localhost",
        "port": "4444",
        
        # value in second, scheduler job interval to hit Service A health_check
        # lesser the value, will increase the accuracy but sacrifice performance
        # larger the value, will decrease API call but sacrificing accuracy
        "health_check_time_interval": 5,

        # value in millisecond, if Service A response time health check is more than
        # timeout_threshold value, it will temporary disable it
        "timeout_threshold": 300,

        # Currently we only have round_robin algorithm, 
        # but we can easily add another algorithm by creating a new class 
        # that implements ILoadBalancer interface
        "load_balancing_algorithm": "round_robin"
    }

### Service A

    {
        "namespace": "service-a",
        "ip": "localhost",
        "port": 8000,

        # Change this if you change the IP & Port on discovery service config
        "discovery_service_base_url": "http://localhost:4444"
    }

# Postman Collection
You can access the Postman Collection on this [link](https://www.postman.com/papannn/workspace/coda-assignment-api-collection/collection/6587816-91d2b026-23fd-4474-ad61-e74860cd1dc7?action=share&creator=6587816&active-environment=6587816-763289f2-4c69-4fe5-8ad9-170de20bdf54)

# Useful API for testing

### Discovery Service: `/api/status`

Currently, we have an endpoint to check registered service

    {
        "Service": {
            "service-a": { // Namespace
                "list": [
                    {
                        "ip": "localhost",
                        "port": "8000",
                        "is_active": true
                    },
                    {
                        "ip": "localhost",
                        "port": "8001",
                        "is_active": true
                    },
                    {
                        "ip": "localhost",
                        "port": "8002",
                        "is_active": true
                    }
                ],
                "index": 0
            }
        }
    }



# Design Flaws that need to be fixed in the future

### Single point of failure for API Gateway Service & Discovery Service
Currently only the `Service A` that got scaled horizontally, we can't achieve this on `API Gateway Service` and `Discovery Service` because there's no load balancer between `Client` to `API Gateway Service` and `API Gateway Service` to `Discovery Service`.

**Solution that I can propose** is to use **Client-Side Load Balancer** to tackle this problem, this will raises another problem we have to manually put the list of `API Gateway Service` and `Discovery Service` to the config.

### Race Condition
Because currently we are not using storage that support transaction lock such as RDBMS that have DB Transaction, when Race Condition happens, for example two client hit `/api/register` at the same time / when lookup API hit too fast, there's a chance for data anomaly. 

**Solution that I can propose** for current logic is to add Thread lock for every updating the data, this will raises performance issue to the `Discovery Service`. **The better solution** is using RDBMS and using DB Transaction.

### API Gateway retry mechanism
The API Gateway retry mechanism will works perfectly when there's only 1 Client / namespace that hit `/api/lookup` , as you can see in this flow
![client-hitting-api-flow](https://res.cloudinary.com/dzfbqsm39/image/upload/v1693311765/coda-payment/rhzv8oocctn91qqax4dm.png)

If for example:

Client 1 is on step 5, then there's Client 2 and 3 doing lookup, the current pointer will be pointing again to `Service A` Port : 8080 because Port 8081 and Port 8082 will be returned to Client 2 and 3, resulting Client 1 got timeout error again.

**Solution that I can propose** is, each client will have dedicated pointer for their Round Robin Algorithm. This will raises another problem such as adding the complexity of the application.

### No centralized monitoring
If we want to check the log of the application, we need to check the separate terminal one by one.

**Solution that I can propose** is to create another service named `Monitoring Service`, the service will have 1 DB to store the logs and can be searched using unique Correlation ID

### It's kinda scary for another developer to refactor the code
It's kinda scary to change the current existing code because if not tested properly, it will break the flow of code.

**Solution that I can propose** write unit test so when another developer tries to change the code, it can be validated by unit test if the current code doesn't break the existing flow or not.

# Time spent

Currently, roughly I've spent 15 hour of my time for this assignment