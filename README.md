# In-Memory Resource Access Control System

(Authorization Simulation in Go)

* * *

## 1\. Project Title

Secure In-Memory Resource Access Control System (Go)

* * *

## 2\. Project Overview

This project implements a tiny backend-like authorization system using Golang.

The system simulates how backend APIs protect resources using:

*   Ownership-based access control  
      
    
*   Role-based authorization (user vs admin)  
      
    

All data is stored in memory, and no actual HTTP server is used.  
The program runs in the terminal and demonstrates authorization logic clearly.

* * *

## 3\. Objective of the Assignment

The goal of this assignment is to understand:

*   How protected resources are accessed  
      
    
*   How authorization decisions are made  
      
    
*   How ownership and roles affect access  
      
    
*   How backend systems deny or allow requests  
      
    

This project focuses on authorization, not authentication.

* * *

## 4\. Technologies Used

*   Language: Go (Golang)  
      
    
*   Data Storage: In-memory structures  
      
    
*   Interface: Command Line (CLI)  
      
    

* * *

## 5\. Data Model Used

### Document Structure

Each document has:

*   id  
      
    
*   owner  
      
    
*   content  
      
    

Example:

{ id: 1, owner: "userA", content: "A's secret" }

  

### User Structure

Each user has:

*   username  
      
    
*   role (user or admin)  
      
    

* * *

## 6\. Core Features Implemented

### Mandatory Features

*   In-memory document storage  
      
    
*   User input for username  
      
    
*   User input for document ID  
      
    
*   Ownership-based access check  
      
    
*   Access granted or denied message  
      
    

### Bonus Features

*   Role-based authorization (admin vs user)  
      
    
*   Admin can access all documents  
      
    
*   Menu-driven program using switch-case  
      
    
*   User switching during runtime  
      
    
*   Session timeout simulation  
      
    
*   Audit logging of access attempts  
      
    
*   Structured API-like responses  
      
    

* * *

## 7\. Program Flow (High-Level)

1.  User logs in with username and role  
      
    
2.  Program displays a menu  
      
    
3.  User selects an action  
      
    
4.  Authorization logic is applied  
      
    
5.  Result is printed  
      
    
6.  Program continues until exit  
      
    

* * *

## 8\. Menu Options Explained

1\. View documents

2\. Access document

3\. View audit logs (admin only)

4\. Exit

5\. Switch user

  

### Option 1: View Documents

*   Displays available document IDs and owners  
      
    
*   Does not reveal content  
      
    

### Option 2: Access Document

*   User enters document ID  
      
    
*   Authorization rules are applied  
      
    

### Option 3: View Audit Logs

*   Admin-only feature  
      
    
*   Displays all access attempts  
      
    

### Option 4: Exit

*   Ends program  
      
    
*   Prints total access attempts  
      
    

### Option 5: Switch User

*   Ends current session  
      
    
*   Starts a new user session  
      
    
*   Audit logs are preserved  
      
    

* * *

## 9\. Authorization Logic

### Normal User

*   Can access only documents they own  
      
    
*   Access denied for other documents  
      
    

### Admin User

*   Can access all documents  
      
    
*   Can view audit logs  
      
    

* * *

## 10\. Audit Logging (Bonus Feature)

Every access attempt is logged with:

*   Timestamp  
      
    
*   Username  
      
    
*   Role  
      
    
*   Document ID  
      
    
*   Result (granted / denied / not found)  
      
    

This simulates real backend security logging.

* * *

## 11\. Session Timeout Simulation

*   Each session allows limited actions  
      
    
*   After limit is reached, user must log in again  
      
    
*   Demonstrates session management concept  
      
    

* * *

## 12\. Structured Output

Responses are printed in API-like format:

{"status":"granted","content":"A's secret"}

{"status":"denied","reason":"not owner"}

  

This makes the program resemble a backend API.

* * *

## 13\. Sample Dry Run

### Case 1: Owner Access

Input:

Username: userA

Role: user

Document ID: 1

  

Output:

Access Granted

A's secret

  

* * *

### Case 2: Unauthorized Access

Input:

Username: userA

Role: user

Document ID: 2

  

Output:

Access Denied

  

* * *

### Case 3: Admin Access

Input:

Username: admin

Role: admin

Document ID: 2

  

Output:

Access Granted (Admin)

B's secret

  

* * *

## 14\. Commands Used

Initialize module:

go mod init in-memory-access-control

  

Run program:

go run main.go

  

Fix module issues:

go mod edit -module=in-memory-access-control

go mod tidy

  

* * *

## 15\. Why This Design Is Correct

*   Uses least privilege principle  
      
    
*   Separates ownership and role-based access  
      
    
*   Mimics real backend authorization logic  
      
    
*   Clean control flow using switch-case  
      
    
*   Easy to extend into HTTP APIs  
      
    

* * *

## 16\. Learning Outcomes

*   Understanding of authorization concepts  
      
    
*   Difference between authentication and authorization  
      
    
*   Ownership-based access control  
      
    
*   Role-based authorization  
      
    
*   Secure backend design thinking  
      
    

* * *
