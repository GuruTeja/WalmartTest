### Prerequisite
 - Install Docker for Local MongoDb database

### Setup MongoDB locally
 - docker pull mongo
 - docker run -d --name localMongoDb -p 10001:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password mongo
 - check if docker container is running or not
   - docker container ls
 - Now you can connect to your local mongoDb database by using: mongodb://admin:password@0.0.0.0:10001/

### Running the Project
 - cd walmartTest/cmd/server
 - go run main.go
 - Now you can POST http://localhost:10000/Alerts

### Screenshots:
**CreateAlert:**
![Screenshot 2023-11-09 at 3.58.14 AM.png](..%2F..%2FDesktop%2FScreenshot%202023-11-09%20at%203.58.14%20AM.png)

**Mongo Database:**
![Screenshot 2023-11-09 at 4.04.30 AM.png](..%2F..%2FDesktop%2FScreenshot%202023-11-09%20at%204.04.30%20AM.png)

**GetAlerts:**
WithData
![Screenshot 2023-11-09 at 4.44.38 AM.png](..%2F..%2FDesktop%2FScreenshot%202023-11-09%20at%204.44.38%20AM.png)
NoData
![Screenshot 2023-11-09 at 4.00.19 AM.png](..%2F..%2FDesktop%2FScreenshot%202023-11-09%20at%204.00.19%20AM.png)
ValidationError
![Screenshot 2023-11-09 at 4.00.31 AM.png](..%2F..%2FDesktop%2FScreenshot%202023-11-09%20at%204.00.31%20AM.png)
ValidationError
![Screenshot 2023-11-09 at 3.59.08 AM.png](..%2F..%2FDesktop%2FScreenshot%202023-11-09%20at%203.59.08%20AM.png)