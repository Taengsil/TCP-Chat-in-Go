# TCP Chat in Go

A TCP-based chat server written in Go that supports multiple clients, chat rooms, and basic command handling. This project demonstrates concurrent server design using goroutines and channels.

---

## Features

- Concurrent TCP server using goroutines
- Multiple client support
- Chat room system (create/join dynamically)
- Nickname system
- Message broadcasting within rooms
- Command-based interaction
- Lightweight (standard library only)

---

## Architecture

The application follows a channel-driven, concurrent client-server model.

### Core Components

#### Server
- Maintains:
  - Active chat rooms (`map[string]*room`)
  - Command channel (`chan command`)
- Central event loop (`run`) processes all client commands
- Handles:
  - Nickname changes
  - Room creation/joining
  - Message broadcasting
  - Client disconnection

#### Client
- Represents a single TCP connection
- Reads input from the socket
- Parses commands and sends them to the server via a channel
- Receives messages and writes them back to the connection

#### Room
- Contains:
  - Room name
  - Connected members (`map[net.Addr]*client`)

#### Command System
- Commands are parsed client-side and sent as structured data, processed centrally by the server
---

## Commands

- `/nick <name>` – Set your username (default: anonymous)
- `/join <name>` – Join or create a chat room (one room at a time)
- `/rooms` – List available chat rooms
- `/msg <message>` – Send a message to the current room
- `/quit` – Disconnect from the server

---
## How It Works
1. Client connects via TCP
2. A client instance is created
3. Input is continuously read and parsed into commands
4. Commands are sent through a channel to the server
5. The server processes commands in a single event loop
6. Messages are broadcast within rooms
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
The server will start on:
```
localhost:8888
```

## Step 4 - Connect via Telnet

Open a new terminal and connect using telnet:
```
telnet localhost 8888
```
You can open multiple terminals to simulate multiple users.

## Step 5 – Start chatting
You should now be connected and able to send/receive messages within chat rooms.

## Example Usage

```
/nick Maddie
/join general
/msg Hello everyone
```

## Notes / Limitations
- A user can only be in one room at a time
- No persistence (rooms and messages are in-memory)
- Minimal input validation (commands assume correct format)
- No authentication or encryption

# Credit

Original video - [Building a TCP Chat in Go(video)](https://www.youtube.com/watch?v=Sphme0BqJiY) by Alex Pliutau, found in the [Github Project Repository](https://github.com/practical-tutorials/project-based-learning#go).

# Preview
![](https://github.com/Taengsil/TCP-Chat-in-Go/assets/75990982/640ff4ea-f28f-4499-b139-6cdd3272dbb4)
