# Sneat Go Server Architecture

This document describes the architecture of the Sneat Go Server, which provides a Firebase-based implementation for the Sneat backend.

## Overview

The `sneat-go-server` is a thin wrapper around the vendor-agnostic `sneat-go-backend` module. It provides Firebase-specific implementations for authentication, database access, and other infrastructure concerns while keeping the core business logic independent of any particular vendor.

## High-Level Architecture

```mermaid
graph TB
    subgraph "sneat-go-server"
        Main[main.go]
        Server[sneatserver]
        FB4Sneat[firebase4sneat]
        SneatFB[sneatfb]
        Facade2FB[facade2firebase]
    end
    
    subgraph "External Dependencies"
        Backend[sneat-go-backend<br/>Business Logic]
        Core[sneat-go-core<br/>Interfaces & Utilities]
        Firebase[Firebase Admin SDK]
        Firestore[Firestore Client]
        Dalgo[dalgo2firestore<br/>Data Abstraction Layer]
    end
    
    Main --> Server
    Server --> FB4Sneat
    FB4Sneat --> SneatFB
    FB4Sneat --> Facade2FB
    
    Server -.implements.-> Backend
    Facade2FB -.implements.-> Core
    SneatFB --> Firebase
    SneatFB --> Firestore
    FB4Sneat --> Dalgo
    
    Backend -.uses interfaces from.-> Core
    
    style Main fill:#e1f5ff
    style Server fill:#b3e5fc
    style FB4Sneat fill:#81d4fa
    style SneatFB fill:#4fc3f7
    style Facade2FB fill:#29b6f6
```

## Package Structure

```mermaid
graph LR
    subgraph "Core Packages"
        SS[sneatserver<br/>HTTP Server Setup]
        F4S[firebase4sneat<br/>Firebase Orchestration]
        SFB[sneatfb<br/>Firebase Integration]
        F2F[facade2firebase<br/>Firebase API Impl]
    end
    
    SS -->|initializes| F4S
    F4S -->|uses| SFB
    F4S -->|registers| F2F
    SFB -->|provides| FirebaseClient[Firebase Client]
    SFB -->|provides| AuthToken[Auth Token Parser]
    SFB -->|provides| FSClient[Firestore Context]
    F2F -->|implements| UserMgmt[User Management]
    F2F -->|implements| TokenAuth[Token Issuance]
    
    style SS fill:#ffecb3
    style F4S fill:#ffe082
    style SFB fill:#ffd54f
    style F2F fill:#ffca28
```

## Initialization Flow

```mermaid
sequenceDiagram
    participant M as main.go
    participant S as sneatserver.InitServer()
    participant F4S as firebase4sneat.InitFirebase()
    participant SFB as sneatfb
    participant F2F as facade2firebase
    participant App as sneatgaeapp.Start()
    
    M->>S: Initialize Server
    S->>F4S: InitFirebase()
    F4S->>SFB: GetFirebaseProjectID()
    SFB-->>F4S: Project ID
    F4S->>F4S: InitFirebaseForSneat(projectID, "sneat")
    F4S->>F2F: Register facade implementations
    Note over F2F: token4auth.IssueAuthToken<br/>facade.GetSneatDB
    F4S-->>S: Firebase Ready
    S->>App: Start(":8080")
    Note over App: HTTP server listening<br/>Business logic handlers active
```

## Component Interaction

```mermaid
graph TB
    subgraph "HTTP Layer"
        Router[HTTP Router]
        Handlers[Request Handlers<br/>from sneat-go-backend]
    end
    
    subgraph "Firebase Integration Layer"
        AuthParser[Auth Token Parser<br/>sneatfb]
        DBContext[Firestore Context<br/>sneatfb]
        UserFacade[User Management<br/>facade2firebase]
        TokenFacade[Token Issuance<br/>facade2firebase]
    end
    
    subgraph "Firebase Services"
        FirebaseAuth[Firebase Auth]
        FirestoreDB[Firestore Database]
    end
    
    subgraph "Data Abstraction"
        Dalgo[DAL Interface<br/>dalgo]
        Dalgo2FS[Firestore Adapter<br/>dalgo2firestore]
    end
    
    Router --> Handlers
    Handlers --> AuthParser
    Handlers --> DBContext
    Handlers --> UserFacade
    Handlers --> TokenFacade
    
    AuthParser --> FirebaseAuth
    DBContext --> Dalgo
    Dalgo --> Dalgo2FS
    Dalgo2FS --> FirestoreDB
    
    UserFacade --> FirebaseAuth
    TokenFacade --> FirebaseAuth
    
    style Router fill:#e8f5e9
    style Handlers fill:#c8e6c9
    style AuthParser fill:#ffecb3
    style DBContext fill:#ffe082
    style UserFacade fill:#ffd54f
    style TokenFacade fill:#ffca28
    style FirebaseAuth fill:#ffcdd2
    style FirestoreDB fill:#ef9a9a
    style Dalgo fill:#ce93d8
    style Dalgo2FS fill:#ba68c8
```

## Key Design Principles

### 1. Dependency Injection

The server uses dependency injection to provide Firebase-specific implementations to the vendor-agnostic backend:

```mermaid
graph LR
    Backend[sneat-go-backend] -.depends on interfaces.-> Interfaces[sneat-go-core interfaces]
    Server[sneat-go-server] -->|injects implementations| Backend
    Server -->|implements| Interfaces
    
    style Backend fill:#b3e5fc
    style Interfaces fill:#ffe082
    style Server fill:#c8e6c9
```

### 2. Abstraction Layers

Multiple abstraction layers keep the codebase maintainable and testable:

```mermaid
graph TB
    BL[Business Logic Layer<br/>sneat-go-backend]
    AL[Abstraction Layer<br/>dalgo interfaces]
    IL[Implementation Layer<br/>dalgo2firestore]
    VL[Vendor Layer<br/>Firestore SDK]
    
    BL --> AL
    AL --> IL
    IL --> VL
    
    style BL fill:#e1bee7
    style AL fill:#ce93d8
    style IL fill:#ba68c8
    style VL fill:#ab47bc
```

### 3. Separation of Concerns

Each package has a clear, single responsibility:

```mermaid
mindmap
  root((sneat-go-server))
    sneatserver
      HTTP routing
      Static files
      Server initialization
    firebase4sneat
      Firebase setup
      Facade registration
      Dependency wiring
    sneatfb
      Auth token parsing
      Firestore context
      Firebase client mgmt
    facade2firebase
      User management API
      Token issuance API
      Firebase App wrapper
```

## Module Dependencies

```mermaid
graph TB
    Server[sneat-go-server<br/>This Repository]
    
    subgraph "Core Modules"
        Backend[sneat-go-backend<br/>Business Logic & HTTP Handlers]
        Core[sneat-go-core<br/>Interfaces & Utilities]
    end
    
    subgraph "Data Layer"
        Dalgo[dal-go/dalgo<br/>Database Abstraction]
        D2F[dal-go/dalgo2firestore<br/>Firestore Implementation]
    end
    
    subgraph "Firebase SDKs"
        FBAdmin[firebase.google.com/go/v4<br/>Firebase Admin SDK]
        FSClient[cloud.google.com/go/firestore<br/>Firestore Client]
    end
    
    subgraph "Infrastructure"
        HttpRouter[julienschmidt/httprouter<br/>HTTP Routing]
    end
    
    Server --> Backend
    Server --> Core
    Server --> D2F
    Server --> FBAdmin
    Server --> FSClient
    Server --> HttpRouter
    
    Backend --> Core
    Backend --> Dalgo
    D2F --> Dalgo
    D2F --> FSClient
    
    style Server fill:#81d4fa
    style Backend fill:#b3e5fc
    style Core fill:#e1f5ff
    style Dalgo fill:#ce93d8
    style D2F fill:#ba68c8
    style FBAdmin fill:#ffcdd2
    style FSClient fill:#ef9a9a
    style HttpRouter fill:#c8e6c9
```

## Request Flow

```mermaid
sequenceDiagram
    participant Client
    participant Router
    participant Handler as Handler<br/>(sneat-go-backend)
    participant Auth as Auth Parser<br/>(sneatfb)
    participant DB as Database<br/>(dalgo → Firestore)
    participant Firebase
    
    Client->>Router: HTTP Request + Auth Token
    Router->>Handler: Route to handler
    Handler->>Auth: Parse & validate token
    Auth->>Firebase: Verify Firebase token
    Firebase-->>Auth: User info
    Auth-->>Handler: Authenticated context
    Handler->>DB: Query/Update via DAL
    DB->>Firebase: Firestore operations
    Firebase-->>DB: Results
    DB-->>Handler: Data
    Handler-->>Router: Response
    Router-->>Client: HTTP Response
    
    Note over Auth: sneatfb.ParseAuthToken
    Note over DB: dalgo interface<br/>dalgo2firestore impl
```

## Testing Strategy

The architecture enables multiple testing strategies:

```mermaid
graph TB
    subgraph "Test Levels"
        Unit[Unit Tests<br/>Individual functions]
        Integration[Integration Tests<br/>With test Firestore]
        E2E[End-to-End Tests<br/>Full request flow]
    end
    
    subgraph "Test Facilities"
        Mock[Mock Implementations<br/>For interfaces]
        TestFS[Test Firestore Emulator]
        TestFB[Test Firebase Project]
    end
    
    Unit --> Mock
    Integration --> TestFS
    Integration --> TestFB
    E2E --> TestFS
    E2E --> TestFB
    
    style Unit fill:#c8e6c9
    style Integration fill:#a5d6a7
    style E2E fill:#81c784
    style Mock fill:#ffecb3
    style TestFS fill:#ffe082
    style TestFB fill:#ffd54f
```

## Summary

The Sneat Go Server architecture achieves several key goals:

1. **Vendor Independence**: Core business logic in `sneat-go-backend` remains vendor-agnostic
2. **Clean Abstractions**: DAL pattern allows swapping database implementations
3. **Testability**: Interface-based design enables mocking and testing at multiple levels
4. **Maintainability**: Clear separation of concerns makes code easy to understand and modify
5. **Extensibility**: New implementations can be added by implementing existing interfaces

The architecture follows Go best practices and creates a solid foundation for building a scalable, maintainable backend service.
