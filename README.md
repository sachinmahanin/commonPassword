# commonpassword service 
Responsible for searching the password provided in a file containing all the common password list placed at https://pwlist.cfapps.eu10.hana.ondemand.com/passwords.txt.
It downloads the file from the web on the application startup and stores in locally.

## Developer Setup
  Add launch.json to provide the environment variable
```
  {
    "version": "0.2.0",
    "configurations": [        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "remotePath": "",
            "HOSTNAME": "127.0.0.1",
            "program": "${workspaceRoot}",
            "env": {
                "HOSTNAME": "127.0.0.1",
                "APP_PORT": 18605,
                "APP_VERSION": "0.15.25" ,  
                "PASSWORD_URL":"https://pwlist.cfapps.eu10.hana.ondemand.com/passwords.txt",
                "PASSWORD_PATH":"./download/PasswordList.txt"
            },
            "args": [],
            "showLog": true
        }
    ]
  }
```

## API Request
How to send API Request to the code running locally through tool like Postman etc.?

Method : POST

URL : http://localhost:18605/CommonPassword

Body :

```
{
    "password": "JamesBond#007forYou"
}
```


Response 

Status : 200

Body : "Your password is not in the common password list "

## Code structure

### passwordstrength

  .github        -CI workflow
  
  config         -Application level configuration eg. AppName,AppVersion etc
  
  customization  -Customization related to bootstrapping,logging,hosting,handlers and web requests
  
  deployment     -k8 yaml files like deployment.yaml,service.yaml
  
  handler        -hanlers for API request
  
  middleware     -middleware related to logging or other
  
  sessionutil    -Used for Session related operations
  
  timeutil       -Used for time related operations
  
  vendor         -stored dependent libraries
  
  web            - RegisteredStatics Static,Business & Utility routes

## CI
As soon as commit is pushed to the master branch, CI workflow is triggered. Which include
1. Build
2. Running Test cases with code coverage
3. Building docker image
4. push the image to the docker hub account https://hub.docker.com/r/sachinmahanin/commonpassword

## Running the service in minikube
1. Run the minikube service on your local box - minkube start
2. Run following command from the root commonpassword will create deployment+service+configMap in the minikube
```
kubectl apply -f deployment
```

