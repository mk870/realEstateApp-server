# R-Estates  Server (Backend)
<img src="https://i.ibb.co/51PDVwk/gin.png" alt="gin" border="0" width="400" align="center"> 

## Project Summary 
* This  is a Restful Gin application with CRUD operations that allow users to create an account to the R-Estates app, save or add agents and properties , delete those agents and properties and login and out using Go jwt authentication.
* It uses  Gin with Tomcat server as a framework.
* The app has 14 endpoints namely : /home, /signup, /logout, /verifification, /refresh-token, /login, /property, /agent, /properties, /agents, /property/id, /notifications, /user, and /agent/id.
* Uses Go Jwt to secure these endpoints.
* Uses Gin GORM to persist data to a postgreSQL database.


### **Resources Used**
***
**Go Version**: 1.19.4

**Dependencies**: Jwt Token, GORM, Go Mail, GIN, postgreSQL-Driver, Go-Cors and Google uuid.  
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white) 	![JWT](https://img.shields.io/badge/JWT-black?style=flat&logo=JSON%20web%20tokens) 	![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=flat&logo=postgresql&logoColor=white)

**For Web Framework Requirements**: go.mod

**APIs**: None

### **EndPoints Building**
***
Built 6 Controllers, authController, agentController, userController, verificationController, notificationsController and propertyController.
#### **User Account Creation Endpoints:** 
* **/signup (PostMethod)**: Takes in firstname, lastname, password and email for user signup. A Jwt token is created as an authentication tool, its stored on the database and also sent by go mail to user email for verification. The password is encrypted using BCryptPasswordEncoder.

* **/verification  (GetMethod)**: validates the email token against the one on the database, once verified the account is enabled. 
* **/login  (GetMethod)**: A Jwt token is created and returned if user login credentials are valid. 


#### **UserAgents Endpoints:**  
* **/agent (PostMethod)**:  saves users' agents to the database with all the agent's properties like name, photo and agent_id. 
* **/agents (GetMethod)**:  retrieves all the saved agents of a client from the database.
* **/agent/id (DeleteMethod)** : deletes a specific agent by id from the database.

#### **UserNotifications Endpoints:**  
* **/notification (PostMethod)**:  saves users' notifications to the database with all the notification's properties like type, description, date and action. 
* **/notifications (GetMethod)**:  retrieves all the saved notifications of a client from the database.
* **/notification/id (DeleteMethod)** : deletes a specific notification by id from the database.

#### **UserProperties Endpoints:**  
* **/property (PostMethod)**:  saves users' properties to the database with all the property's properties like name, price, photo,bedrooms, bathrooms, size, location and status. 
* **/properties (GetMethod)**:  retrieves all the saved properties of a client from the database.
* **/property/id (DeleteMethod)** : deletes a specific property by id from the database.  

### **Data Storage**
Used GORM (ORM) to persist and retrieve data from a postgreSQL database.  
Built 4 models: 
* User Model to store app users.
* VerificationToken Model to store signup verification tokens.
* Agent Model.
* Property Model.
* Notifications Model. 



### **Productionization**
***
In this step I deployed the postgreSQL database to ElephantSQl via 3rd party and deployed the Gin app to Render Cloud.

**Live Implemantation:** [R-Estates](https://r-estates.vercel.app/)