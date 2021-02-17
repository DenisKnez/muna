## **Project**

- This is what I was able to do in ~5  hours and after a brake changed the stat methods and small errors that didn't allow the project to build (not sure how long that took)
- Was not able to test it to see if everything works (because of time), but it's buildable
- I used regex (website) to build and test my regex, so that is "tested"

## **What should be done/improved***

### **Logging**
- add logging into file or database

### **Migrations**
- instead of just sql files like I did here there should be real migrations either with 3rd party library or custom created migration logic
- with down and up migrations

### **Redis**
- setup redis for improved performance (I have never personaly used redis but I know what it is and why it is used, and have seen it being used)

### **Config file**

- There should be a config file for project configurations like the database connection string (in the project I just placed a connection string in the sql.Open())


### **Server config**

- The server configuration in golang should be implemented
  - max header bytes
  - read timeout
  - write timeout
  - ...

### **Dependency injection**
- In project everything is instantiated in the main **(bad)**
- There should be dependecy injection container (3rd party or custom made) and dependency injection should be used in the whole project

### **Routing**
- In this project the routing was done with the standard library which is very limiting (if not custom modified to have more features, like the ability to add parameters in the url path)


### **Testing**
- There should be tests

### **Concurrency**
- Make the program concurrent with go routines and channels

### **Style/Usage markdown file**
- Markdown file that explains the conventions/aggred ways used in this project for this like:
  - database constaint naming
  - migration files naming
  - project versioning
  - data that needs to present in every log
  - ...

### **Probably some configuration for staging/production in separate file or in the config file with the other configurations**
