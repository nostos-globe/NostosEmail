# Nostos Email Service

## Description

The Nostos Email Service is a dedicated microservice responsible for handling all email communications within the Nostos platform. It processes user-related events through NATS messaging and sends appropriate email notifications such as welcome emails and password reset instructions.

---

## Features

- Event-driven architecture using NATS for message processing
- Templated HTML emails for consistent branding
- Support for multiple email types (welcome, password reset)
- Embedded HTML templates for easy deployment
- Environment-based configuration for SMTP settings
- Containerized deployment with Docker

---

## Technologies Used

- **Language**: Go
- **Messaging**: NATS for event subscription
- **Email**: SMTP for email delivery
- **Templates**: Go's html/template package
- **Configuration**: Environment variables with godotenv support
- **Containerization**: Docker with multi-stage builds

---

## Architecture

The service follows a clean architecture pattern with the following components:

- **Event Listener**: Subscribes to NATS topics and processes incoming events
- **Mailer**: Handles email template rendering and SMTP delivery
- **Templates**: HTML email templates with styling

---

## Email Templates

The service includes professionally designed HTML email templates:

- **Welcome Email**: Sent when a new user registers
- **Password Reset Email**: Sent when a user requests a password reset

All templates feature responsive design and consistent branding.

---

## Event Processing

The service listens for the following NATS events:

- `user.registered`: Triggers welcome email
- `user.password_reset_requested`: Triggers password reset email

---

## Security

- Secure SMTP connections
- Environment-based configuration for sensitive information
- No storage of user data beyond processing requirements

---

