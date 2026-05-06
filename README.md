# TCP Chat in Go

A simple TCP-based chat application written in Go. This project demonstrates how to build a concurrent client-server chat system using Go’s networking and goroutines.

---

## Features

- TCP server handling multiple clients concurrently
- Real-time message broadcasting
- Lightweight and minimal dependencies
- Goroutine-based concurrency model
- Simple CLI-based client interaction

---

## Architecture

The application follows a basic client-server model:

### Server
- Listens for incoming TCP connections
- Handles each client in a separate goroutine
- Broadcasts messages to all connected clients

### Client
- Connects to the server via TCP
- Sends user input to the server
- Receives and displays messages from other users

---

## Commands

- `/nick <name>` – Set your username (default: anonymous)
- `/join <name>` – Join or create a chat room (one room at a time)
- `/rooms` – List available chat rooms
- `/msg <message>` – Send a message to the current room
- `/quit` – Disconnect from the server

---

## Getting Started

### Prerequisites

- Go installed (>= 1.18 recommended)
- Telnet client installed

---

### Step 1 – Clone the repository

```bash
git clone https://github.com/Taengsil/TCP-Chat-in-Go.git
cd TCP-Chat-in-Go
```

### Step 2 - Build
```
go build
```

## Step 3 - Run the server
On Linux
```
./TCP_Chat_In_Go
```
On Windows
```
TCP_Chat_In_Go.exe
```
## Step 4 - Connect via Telnet

Open a new terminal and run:
```
telnet localhost 8888
```

## Step 5 – Start chatting
You should now be connected and able to send/receive messages within chat rooms.

# Credit

Original video - [Building a TCP Chat in Go(video)](https://www.youtube.com/watch?v=Sphme0BqJiY) by Alex Pliutau, found in the [Github Project Repository](https://github.com/practical-tutorials/project-based-learning#go).

# Preview
![](https://github.com/Taengsil/TCP-Chat-in-Go/assets/75990982/640ff4ea-f28f-4499-b139-6cdd3272dbb4)
