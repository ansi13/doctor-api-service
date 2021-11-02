# doctor-api-service

## Assignment
Create a REST API Web Service for doctors to create and manage their patients with their medical records. There can be multiple doctors, a doctor can create multiple patients. Each doctor can access their patients only.
A doctor can create, update and delete their patients’ medical records. Deploy this web service on the platform of your choice.

## Working

This app works based on the user authentication with username and password. Create an account using `signup` and login in to get the token `login`. Using the login
token as a `Bearer` token in the request header we will authenticate the user for retrieving all patients for that particualr doctor, creating a patient, retrieving a
single patient, updating a patient and deleting a patient.

## Documentation & Resources
This below postman collection contains details you can use to start working with this API.
https://www.getpostman.com/collections/5c3f1b8f9d7ec736f6b9
