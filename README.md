# Chat Room Server

This is a simple chat room server written in Go. It uses a WebSocket to handle real-time communication between clients.

## Prerequisites

- Go 1.16 or later

## Installation

Clone the repository to your local machine:

## Usage

To start the server, use the following command: Go run cmd/main.go

By default, the server will run on port 8080. You can change this by setting the `PORT` environment variable.

The server uses the following routes:

- `/users`: to handle user-related requests.
- `/ws`: to handle WebSocket connections.
¿
