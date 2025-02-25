<h1>Backend Proejct</h1>
<h2>Introduction</h2>

  In Spring 2023 I took my Processes of Object-Oriented Programming Class (POOSD) where we were tasked as a group to create a functional application. My group chose to design a workout application tailored to athletes to allow them to track their calories. I was tasked with developing the backend application for our project. I had no prior experience, but in reading the documentation and watching videos online, I was able to help my team to develop a functional application that was able to pass data from the front-end and store it into the MongoDB database.
  During my POOSD class I took sort of a brute force approach in getting the application to work because I had no prior experience, so when I came across this internship, I was excited at the opportunity to learn the “proper” way to design and develop the backend for an application. During this internship I was able to fill in gaps in the importance of topics such as data input validation, organization and file structure, documentation, and security to name a few. I created a mini project of my own to demonstrate some of the knowledge I gained from this internship. My approach is to show the flow of data from the front-end until it is stored. Two key aspects of this project are the file structure and testing.


<h2>File Structure</h2>

	The way the code is organized is very important because it determines how easy it is for new and current team members to follow. Before I participated in this internship, I stored all components necessary for the backend into one file. I learned the importance of file structure and directly witnessed how much the readability increases and time taken decreases for development with well-structured code. The key areas of the file structure that I would like to cover are the following folders:
-	models
-	routes
-	configs


<h3>Models</h3>
  Within the models folder, the schema is designed and defined. The schema is an outline of the data that is expected to be collected from the frontend and stored in the backend database. In my Courses file (for schemas), I outlined a few fields that define information that I wish to store about each course.

![image](https://github.com/user-attachments/assets/e2d43bb7-6757-41db-8ef4-a052e6a6aa7e)<br>

<h3>Routes</h3>

  Within the routes folder, a few things happen. Functions are defined to determine what happens with the incoming data from the frontend. The Course schema is imported for use in these functions. This is necessary because within this subdirectory, communication with the database occurs and so data has to travel to the database following a certain format (which is defined by the schema).
  The most basic database query operations are known as the CRUD operations, and this acronym translates to Create, Read, Update and Delete. For any object, you must be able to create a new one, read (or get) it from the database, update its contents in the database, and delete (or remove it from the database). These are the core functions that I have outlined in the scope of this project, but with respect to managing courses additional functions can be created such as favoriting/disliking a course, listing a course as a pre/corequisite, grouping courses, etc.; and all of these functions would be defined under the routes subfolder. Below a snippet of my courses file (for functions) can be seen.

![image](https://github.com/user-attachments/assets/f53ae207-487a-475b-b9e2-ad6ed9cf3967)<br>

All CRUD operations were implemented in this process. The following is a pseudocode overview of the CreateCourse function. It is as follows:
<br>```// Create a new course, and save it to the database```
<br>```// Parse the request body to create the course object```
<br>```// Validate the data```
<br>```// Insert the new course into MongoDB```
<br>```// Log the created course to the terminal```
<br>```// Return a response with the course created and a status code```
<br>Results for all functions will be displayed in a later section. It is not shown in the image above, nor in the pseudocode, but another crucial step is error handling. This is done after almost every process to ensure that valid data is being transferred to the next stage. For error handling I often use if statements that check if an error has occurred and return the appropriate status code and a brief message describing which attempted action led to the error. This has made debugging much simpler and also made it easier for other developers (who were not the initial authors of the code), to also debug and integrate their code.

<h3>Configs</h3>

  The final folder of significance is the configs folder. Within this folder, the database connection is configured and established.

![image](https://github.com/user-attachments/assets/9c95a8ed-9abd-49f6-89b2-f833328f768c)<br>
The pseudocode can be provided for the ConnectDB function as well:
<br>```// Set the connection string, imported from the .env file```
<br>```// Attempt to connect to MongoDB```
<br>```// Ping the database to ensure connection```
<br>```// Log success or failure message```
<br>```// Return the connection```

  The .env file is a hidden file (denoted by the “.” at the beginning of the file name). It is short for the environment, and it stores environmental variables. An environment variable is a name/value pair that is set outside of the program and accessed during runtime by referencing the name in the pair. If changes are made in the configuration, environment variables reduce the need to modify the application as a result. In this instance, my environment variable is called “MONGO_URI,” and the value is the connection string. MongoDB is an external application that is integrated into the program, so it requires external credentials for access. This connection string can be used to access the MongoDB application.
  The function in my program that is responsible for establishing a database connection is called “ConnectDB.” The database.go file within the configs folder is referenced by the courses.go file within the routes folder during the creation of functions because access to the database is needed to implement the CRUD operations. Once a method to establish the database has been implemented, everything can be placed together in main.go, which is the file that is compiled at execution time.


<h2>Main.go</h2>

  This is an overview of my main.go file

![image](https://github.com/user-attachments/assets/74aff9c4-fed9-47ff-b660-2416eda80ba9)<br>

  This main.go file involves two processes, creating routes and setting up a server.

![image](https://github.com/user-attachments/assets/7ece7b1d-fb6a-4d2d-abf3-b08779b0db1d)<br>
  A route is the name you can use to access an API endpoint, which is a medium for connecting our application to the database. It is a URL that acts as the point of contact between the client and server. Four routes were created for each operation, a Post method for the CreateCourse function, a Get method for the Read function, a Put method for the Update function, and a Delete method for the Delete function. These methods are commonly used for each operation within the industry.

![image](https://github.com/user-attachments/assets/b644585d-8ea0-44cd-a3d6-e49441a290d8)<br>
  The above image is responsible for starting a server that can be reached through an API tester at port 3000 in the localhost. Localhost is good for development as it establishes a temporary environment that lasts for the length of program’s execution, but more permanent servers are commonly established for production environments. After running the command “go run main.go,” the server is set up and running. 


<h2>Testing</h2>

  The MongoDB database is initially empty

![image](https://github.com/user-attachments/assets/9b6242b0-d1aa-45eb-a2c8-e1a98c981f55)<br>

<h3>Creating A Course</h3>

![image](https://github.com/user-attachments/assets/f477c3d8-3f1e-48ac-8db0-ad3e158ce383)<br>
![image](https://github.com/user-attachments/assets/7a531c1c-6dd6-48c5-bbc6-699235dd0d67)<br>
  The first image shows the JSON Object being defined in the API tester Insomnia. The object contains all the fields expected and outline in the Course schema. The second image shows the successful creation of the new course in the MongoDB database. A new ObjectId is issued for the new course.

<h3>Reading A Course</h3>

![image](https://github.com/user-attachments/assets/4a8bcff2-0de8-458a-9881-ff914916f6d5)<br>
 
  In Insomnia, the ObjectId created for the new course has to be provided in the form of a URL parameter, in order to identify the correct course. The correct course is displayed on the right.

<h3>Updating The Course</h3>

![image](https://github.com/user-attachments/assets/f05e5350-95d6-4473-99c2-aa00e9dcbe27)<br>
![image](https://github.com/user-attachments/assets/bacaebf9-7fd9-4d65-9a27-b1138e17a70a)<br>
![image](https://github.com/user-attachments/assets/71c391cc-dd04-4667-8a33-c32b67097e8c)<br>
  The first image follows the same steps as Reading the course. The second image follows the same steps as Creating the course. The Course Id must be supplied in the URL param, and a JSON object has to be created with the same initial fields but with updated information. The third image shows the updated information in MongoDB.

<h3>Deleting A Course</h3>

![image](https://github.com/user-attachments/assets/5b570374-dad4-4230-adc3-fecb0977b83a)<br>
![image](https://github.com/user-attachments/assets/679f570b-5ea4-4d62-a050-e4385ba62091)<br>
  The first image follows the same steps as Reading a course such that the ObjectId has to be supplied as a URL param. The second image shows that the Course has been deleted from MongoDB.
  A common pattern seen when executing the CRUD operations in Insomnia is the use of the Params and Body tabs, but both Insomnia and MongoDB’s operations can be expanded to cover functionalities beyond the basic implementation of CRUD operations. 



<h2>Conclusion</h2>

  In this project, I aimed to map the complete journey of data from the front-end to the back end of an application. Initially, my understanding of back-end development was limited to establishing database connections and integrating external services. However, this internship has broadened my perspective, emphasizing the importance of organization, design, and security - especially when handling sensitive data. My role evolved from renaming files and variables to gain familiarity with our codebase, to actively contributing to development, maintenance, documentation, and testing. I have honed my skills in technologies like Golang, JavaScript, NodeJS, MongoDB, Insomnia, Postman, and Figma.
  My future goals include mastering automated testing and database management through cron jobs, such as automatically purging inactive accounts. I also plan to expand my experience into front-end development to transition into a full-stack role and pursue professional certifications. I am eager to apply the practical insights gained from this internship to personal projects and further integrate my learning in both theoretical and applied aspects of software development. I am grateful for this opportunity, which allowed me to gain a lot of practical skills and allowed me to work alongside hardworking, inspiring developers.
