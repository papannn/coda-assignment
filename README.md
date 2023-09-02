
# Coda Payment Technical Interview Assignment

Goal: Write an Round Robin API which receives HTTP POSTS and routes them to one of a list of
Application APIs

I'm using golang version `1.21.0`

Library that I used:
* TBD


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
        "port": "8080"
    }

After that, **Discovery Service** will store the designated data to their variable.

![discovery-service-scheduler](https://res.cloudinary.com/dzfbqsm39/image/upload/v1693311765/coda-payment/ypp7oqco8vel24es7q6p.png)

Every `x` second (we can set this value on **Discovery Service** config on `health_check_time_interval` value json), **Discovery Service** will try to do the health checkup for every Service that already register itself to **Discovery Service**.

**Discovery Service** will check some metrics such as:
* Liveness of the service
* How long the Service A responds

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

# How to use
On terminal:
    
    TBD

# Useful Command / API
There's some useful command such as:
    
    TBD

# Future Improvement

    TBD

# Time spent

Currently, roughly I've spent 10 hours of my time for this assignment