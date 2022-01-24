# Mysql Demo Application

A demo application with rest apis to connect with a mysql cluster running in [KloverCloud](https://klovercloud.com) platform. The application dynamically loads required environment variables required to connect with mysql and and perform read/write actions.

## Application details
The application established connection with each endpoint on mysql endpoint. If it fails to connect with any instance, it will throw a fatal log and exit. All the write operations **(POST, PUT, DELETE)** are executed through the master instance and all the **GET** operations are executed through slave instances (round-robin)

## Deploying application and cache

create a new vpc with enough resources to deploy the go application and mysql server. Assuming you already have added personal access token of github/gitlab on klovercloud. OnBoard the application from the create new drop down bar and after that edit the DockerFile as necessary, *e.g.* exposing the port where application is running.

Create a mysql cache in the same vpc, and deploy the cache server. Yahoo! Nearly halfway there, with some mouse clicks right. No pesky terminal hassle. All we need to do is just inform the application about the cache serve. The application already expecting one, so lets do it.


## Secret information through environment variables
First of all, the authentication password. Which is required to connect with every mysql instance. You can easily add this through the secret tab on application page. Create a secret, give it a name and add the key-value as following

```
KEY                                       VALUE
MYSQL_PASSWORD                    {your_cache_password}
```

## Instance information through environment variables
When the application starts, it will look for the master endpoint from environment variable. And also look for slave instance count. If the slave instance is > 0, the application will try to load all the slave endpoints and will ping the slave instance. If it fails to get master / slave endpoints or ping returns error, the application will terminate with fatal error log. Add the environment variables as described below and double check. or what ? Your application will go on infinite crash loop (opps!)

Go to the cache [cache](https://console.klovercloud.com/cache) section, select your cache, go to overview section and click to service endpoint to see all the endpoints list. Add them to environment variables as described below.

```
KEY                VALUE
MASTER_ENDPOINT    {master_endpoint}
SLAVE_COUNT        {no_of_slaves}
SLAVE_ENDPOINT_0   {first_slave_endpoint}
SLAVE_ENDPOINT_1   {second_slave_endpoint}
.
.
.
.

```
deploy your application on desired deployment environment. Make sure you added the environment variable and secrets in the right deployment environment. if you already deployed (and obviously its in crashloop!), re-deploy your application after adding environment variables and secrets. If deployment fails, deployment logs should show the fatal log, its detailed enough to debug easily.
## REST endpoints
#### GET
From the CI/CD pipeline page, select the deployment, right click info. **External Default Endpoint** is your application url, so just add the path /api/v1/id. Where id is param. Without param it will return “method not found”.
```bash
Example:
https://your-application-prefix.eu-west-1.klovercloud.com/api/v1/10
```
If your requested data not exit in mysql database then it show “Data Not Found”

#### POST
in the request body, provide json containing data as the below example. Only strings are allowed. If id already exist it will show "Data Already Stored". 
Hit the same endPoint with POST method.\
https://your-application-prefix.eu-west-1.klovercloud.com/api/v1
```json
{
    "id"   : "10",
    "name" : "Tamim"
    "designation" : "pqr"
    "branch" : "7"
}
```
#### PUT
same as post. Oh, it checks if the key exist or not. Updating a key that isn't there yet? Nah that's not gonna happen. Just show "Data Not Found"

```json
{
    "id"   : "10",
    "name" : "Shahadath"
    "designation" : "xyz"
    "branch" : "6"
}
```
#### DELETE
Same as the GET api. Provide id into the params.

## Contributing
Pull requests for new features, bug fixes, and suggestions are welcome!


