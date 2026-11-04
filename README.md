# go-rest-app

# Event Booking REST API

A practice project for learning how to build a REST API in Go. The API models event booking, letting users create and browse events, register for events, and manage their registrations.

## What the project is for

Use this project to practice:

- Designing REST endpoints with HTTP methods and URL paths
- Creating, reading, updating, and deleting event records
- Signing users up and authenticating them at login
- Protecting routes with JWT authentication
- Restricting event updates and deletion to the event creator
- Registering for an event and cancelling a registration

## API overview

| Method | Endpoint | Purpose | Access |
| --- | --- | --- | --- |
| `GET` | `/events` | List available events | Public |
| `GET` | `/events/<id>` | Get an event by ID | Public |
| `POST` | `/events` | Create a bookable event | Authentication required |
| `PUT` | `/events/<id>` | Update an event | Authenticated event creator |
| `DELETE` | `/events/<id>` | Delete an event | Authenticated event creator |
| `POST` | `/signup` | Create a user account | Public |
| `POST` | `/login` | Authenticate a user and obtain a JWT | Public |
| `POST` | `/events/<id>/register` | Register for an event | Authentication required |
| `DELETE` | `/events/<id>/register` | Cancel an event registration | Authentication required |

Authenticated endpoints expect a valid JWT obtained through `/login`. Creating, updating, or deleting events requires authentication; only the event's creator can update or delete it.

## Learning goal

The goal is to build practical experience with Go, HTTP APIs, authentication, and authorization by implementing the core flows of an event booking service.
