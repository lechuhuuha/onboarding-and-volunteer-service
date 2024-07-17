### Onboarding and Volunteer Service

This repository contains the codebase for the Onboarding and Volunteer Service application. The application is built using the Go programming language with the Gin web framework, GORM ORM, and PostgreSQL database. It aims to streamline the process of onboarding and managing volunteers, providing functionality for managing 2 kinds of user: applicant and volunteer.  

## Table of Contents
Project Structure
Installation
Configuration
Database Migration
Usage
API Endpoints
Contributing
License

# Project Structure
The project follows a modular structure with clearly defined folders:

# Installation
To get started with the Onboarding and Volunteer Service application, follow these steps:

Clone the repository:
git clone https://github.com/cesc1802/onboarding-and-volunteer-service.git
cd onboarding-and-volunteer-service

Install dependencies:
go mod download
Set up your environment variables. Copy the .env.example to .env and fill in the necessary values.

cp .env.example .env

# Configuration
Make sure to configure the following environment variables in the .env file:

DB.HOST: Database host
DB.PORT: Database port
DB.USER: Database user
DB.PASS: Database password
DB.NAME: Database name

Database Migration
Run the database migrations to set up the required tables:
go run cmd/migration/main.go

# Usage
To start the application, run:
go run cmd/main.go

The server will start on the port specified in the .env file.

# API Endpoints "/api/v1"
Admin Endpoints: "/admin"
GET "/list-request": Get the request list
GET "/request/:id" : Get a specific request
POST "/approve-request/:id": Approve a request, change status of request
POST "/reject-request/:id": Reject a request, change status of request
POST "/add-reject-notes/:id": Add reject notes to a request
DELETE "/delete-request/:id": Delete a request

User Endpoints: "/applicant"
POST "/:" Create a new user
PUT "/:id" : Update an existing user from register form
DELETE "/:id" : Delete an user with id
GET "/:id" : Get information about an user with id

Application Request Endpoints: "/applicant-request"
POST "/" : Create a record request

User Identity Endpoints: "/applicant-identity"
POST "/" : Create a user identity record
GET "/:id": Find a user identity
PUT "/:id": Update a user identity record  


# Contributing
We welcome contributions to enhance the features and functionality of this project. Please follow these steps:

Fork the repository.
Create a new branch (git checkout -b feature/your-feature).
Commit your changes (git commit -am 'Add new feature').
Push to the branch (git push origin feature/your-feature).
Create a new Pull Request.
License
This project is licensed under the MIT License. See the LICENSE file for details.

