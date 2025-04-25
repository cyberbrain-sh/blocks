# Blocks

A powerful and flexible entity-based content management system developed by [Cyberbrain](https://github.com/cyberbrain-sh).

## Overview

Blocks is a Go library that provides a versatile foundation for managing structured content as interconnected entity blocks. Unlike traditional block-based systems that focus only on layout and formatting, Blocks represents meaningful real-world entities such as movies, TV series, YouTube videos, Instagram posts, emails, and more.

Each block encapsulates a specific entity type with its relevant metadata, properties, and relationships. This enables rich content modeling beyond simple text formatting, allowing applications to work with semantically meaningful content units that map directly to real-world objects and concepts.

The system supports a wide range of entity types within a unified data model, enabling complex applications that can intelligently process, organize, and present rich media and structured information.

## Features

- **Rich Content Types**: Support for multiple content types including pages, paragraphs, headers, lists, to-do items, links, emails, and media (movies, series, YouTube, Instagram).
- **Hierarchical Structure**: Blocks can contain other blocks, allowing for complex content organization.
- **Properties System**: Flexible properties system for storing metadata associated with each block.
- **Content References**: Ability to reference and link blocks using UUID annotations.
- **Content History**: Track the history of block movements and modifications.
- **Block Rendering**: Utilities for rendering blocks into various formats.
- **Lifecycle Management**: Support for different states in a block's lifecycle.

## Installation

```bash
go get github.com/cyberbrain-sh/blocks
```

## Usage

```go
import "github.com/cyberbrain-sh/blocks"

// Create a new empty block
block := blocks.NewEmptyBlock()

// Set properties
block.Properties.ReplaceValue(blocks.PropertyKeyTitle, "My Block")
block.Properties.ReplaceValue(blocks.PropertyKeyText, "This is a sample block")

// Add a child block
childID := uuid.New()
block.AppendChild(childID)

// Change block type
block.Type = blocks.TypeParagraph
```

## Block Types

The library supports two main categories of block types:

- **Structural Blocks**: Content blocks with complex structure and rich metadata
  - Movie, Series, Link, ToDo, Email, Page, Database, YouTube, Instagram, Fragment

- **Textual Blocks**: Simple content containers primarily for text
  - Paragraph, Headers (1-6), Bullet List Items, Numbered List Items, Image, Video, Audio, File

## Block Connectivity and Hierarchy

Blocks connect to each other in a hierarchical structure, creating a flexible and powerful content graph:

### Parent-Child Relationships

Each block maintains both upward and downward references:

- **Parent Reference**: Each block (except root blocks) contains a `ParentID` field that points to its parent block. This creates an upward link in the hierarchy.
- **Child References**: Blocks contain a `Content` array of UUIDs that points to their direct children, creating downward links in the hierarchy.

```go
// Example parent-child relationship
parentBlock := blocks.NewEmptyBlock()

childBlock := blocks.NewEmptyBlock()
childBlock.ParentID = &parentBlock.ID  // Child points to parent

parentBlock.AppendChild(childBlock.ID) // Parent includes child in its Content array
```

### Hierarchy Operations

The Block struct provides methods for managing these relationships:

- `AppendChild(id uuid.UUID)` - Adds a child to the end of the Content array
- `RemoveChild(id uuid.UUID)` - Removes a child from the Content array
- `InsertChild(id uuid.UUID, afterID uuid.UUID)` - Inserts a child at a specific position

### Recursive Relationships

For efficient traversal and operations, blocks also maintain:

- `ChildrenRecursive []uuid.UUID` - Contains all descendants (children, grandchildren, etc.)
- `RootParentID *uuid.UUID` - Points to the topmost ancestor in the hierarchy

This dual-reference system (parent references and child content arrays) allows for:
- Fast bidirectional traversal of the content tree
- Easy restructuring of content hierarchies
- Efficient querying of both ancestors and descendants

### Block Hierarchy Diagram

The following diagram illustrates the block hierarchy structure:

```mermaid
flowchart TD
    %% Define the root block with full details
    Root["ROOT BLOCK (Page/Database)
    ------------------
    ID: uuid-root
    ParentID: nil
    RootParentID: nil
    Content: [uuid-child1, uuid-child2, uuid-child3]
    ChildrenRecursive: [all descendant IDs...]"]
    
    %% Define child blocks
    Child1["CHILD BLOCK 1 (Structural)
    ------------------
    ID: uuid-child1
    ParentID: uuid-root
    RootParentID: uuid-root
    Content: [uuid-gc1, uuid-gc2]"]
    
    Child2["CHILD BLOCK 2 (Textual)
    ------------------
    ID: uuid-child2
    ParentID: uuid-root
    RootParentID: uuid-root
    Content: []"]
    
    Child3["CHILD BLOCK 3 (Structural)
    ------------------
    ID: uuid-child3
    ParentID: uuid-root
    RootParentID: uuid-root
    Content: [uuid-gc3]"]
    
    %% Define grandchild blocks
    GrandChild1["GRANDCHILD 1 (ToDo)
    ------------------
    ID: uuid-gc1
    ParentID: uuid-child1
    RootParentID: uuid-root"]
    
    GrandChild2["GRANDCHILD 2 (Email)
    ------------------
    ID: uuid-gc2
    ParentID: uuid-child1
    RootParentID: uuid-root"]
    
    GrandChild3["GRANDCHILD 3 (YouTube)
    ------------------
    ID: uuid-gc3
    ParentID: uuid-child3
    RootParentID: uuid-root"]
    
    %% Define relationships - downward through Content array
    Root -->|"Content[0]"| Child1
    Root -->|"Content[1]"| Child2
    Root -->|"Content[2]"| Child3
    
    Child1 -->|"Content[0]"| GrandChild1
    Child1 -->|"Content[1]"| GrandChild2
    Child3 -->|"Content[0]"| GrandChild3
    
    %% Define upward relationships - through ParentID
    GrandChild1 -.->|"ParentID"| Child1
    GrandChild2 -.->|"ParentID"| Child1
    GrandChild3 -.->|"ParentID"| Child3
    
    Child1 -.->|"ParentID"| Root
    Child2 -.->|"ParentID"| Root
    Child3 -.->|"ParentID"| Root
    
    %% Define RootParentID relationships with different style
    GrandChild1 -.-|"RootParentID"| Root
    GrandChild2 -.-|"RootParentID"| Root
    GrandChild3 -.-|"RootParentID"| Root
    
    %% Add styling
    classDef root fill:#f9f,stroke:#333,stroke-width:2px
    classDef child fill:#bbf,stroke:#333,stroke-width:1px
    classDef grandchild fill:#dfd,stroke:#333,stroke-width:1px
    
    class Root root
    class Child1,Child2,Child3 child
    class GrandChild1,GrandChild2,GrandChild3 grandchild
```

The diagram shows:
- **Root Block**: A page or database with no parent
- **Solid Arrows** (→): Represent downward references via the Content array
- **Dashed Arrows** (⇢): Represent upward references via ParentID
- **Dotted-Dashed Arrows** (-.-): Represent direct RootParentID references to the root
- Each block contains its key fields including ID, ParentID, RootParentID, and Content
- Different colored blocks represent different levels in the hierarchy

This structure enables both upward traversal (via ParentID) and downward traversal (via Content array) through the block hierarchy.

## Properties

Blocks have type-specific properties that can be accessed and modified:

```go
// Get a string property
if title, exists := block.Properties.GetString(blocks.PropertyKeyTitle); exists {
    fmt.Println("Title:", title)
}

// Set a property
block.Properties.ReplaceValue(blocks.PropertyKeyChecked, true)
```

## Entity-Specific Properties

Each block type represents a specific real-world entity and therefore has its own set of specialized properties reflecting the unique attributes of that entity type. The system provides helper functions for adding, retrieving, and rendering these type-specific properties.

### Movie Block Properties

Movie blocks represent film entities with properties like:

- **Basic Information**: title, description, image URL, original URL
- **Identifiers**: IMDB ID, TMDB ID
- **Film Details**: release year, rating, runtime, tagline
- **Financial Data**: budget, revenue
- **People**: directors, cast members
- **Categorization**: genres

```go
// Example: Creating a Movie block
movieBlock := blocks.NewEmptyBlock()
movieBlock.Type = blocks.TypeMovie

// Set movie-specific properties
title := "The Matrix"
year := 1999
rating := "8.7"
runtime := "136 min"
directors := []string{"Lana Wachowski", "Lilly Wachowski"}
genres := []string{"Action", "Sci-Fi"}

blocks.AddMovieProperties(&movieBlock, 
    &title,                  // title
    nil,                     // description
    nil,                     // imageURL
    nil,                     // url 
    nil,                     // imdbID
    nil,                     // tmdbID
    &year,                   // releaseYear
    &rating,                 // rating
    &runtime,                // runtime
    nil,                     // tagline
    nil,                     // budget
    nil,                     // revenue
    &genres,                 // genres
    &directors,              // directors
    nil,                     // cast
    nil,                     // checked
    true)                    // enriched

// Render the movie block to a human-readable format
renderedMovie := blocks.RenderMovieProperties(movieBlock)
// Output:
// # The Matrix (1999)
// **Rating: ⭐ 8.7 | Runtime: 136 min**
// **Genres:** Action, Sci-Fi
// **Directors:** Lana Wachowski, Lilly Wachowski
```

### Email Block Properties

Email blocks represent email messages with properties like:

- **Message Identifiers**: email ID, thread ID
- **Participants**: from, to
- **Content**: subject, body text
- **Timing**: date, received at
- **Organization**: attachments, labels

```go
// Example: Creating an Email block
emailBlock := blocks.NewEmptyBlock()
emailBlock.Type = blocks.TypeEmail

// Set email-specific properties
emailID := "msg_123456"
from := "sender@example.com"
to := "recipient@example.com"
subject := "Meeting Tomorrow"
text := "Hi there,\n\nLet's meet tomorrow at 2 PM.\n\nRegards,\nSender"
date := time.Now()
labels := []string{"Important", "Work"}

blocks.AddEmailProperties(&emailBlock,
    &emailID,               // emailID
    nil,                    // threadID
    &from,                  // from
    &to,                    // to
    &subject,               // subject
    &text,                  // text
    &date,                  // date
    nil,                    // receivedAt
    nil,                    // attachments
    &labels)                // labels

// Render the email block to a human-readable format
renderedEmail := blocks.RenderEmailProperties(emailBlock)
// Output:
// ## Meeting Tomorrow
// **From:** sender@example.com
// **To:** recipient@example.com
// **Date:** Jul 20, 2023 3:45 PM
// **Labels:** `Important` `Work`
// ---
// Hi there,
//
// Let's meet tomorrow at 2 PM.
//
// Regards,
// Sender
```

### Other Entity Types

Similar specialized property sets exist for other entity types:

- **YouTube Videos**: video ID, channel information, view counts, etc.
- **Series**: seasons, episodes, networks, etc.
- **ToDo Items**: checked status, due dates, etc.
- **Instagram Posts**: post ID, caption, like count, etc.

Each entity type has dedicated functions for:
- Adding properties (`Add[Type]Properties`)
- Rendering to human-readable format (`Render[Type]Properties`)
- Getting a list of available properties (`Get[Type]Properties`)

This type-specific approach enables the blocks system to accurately model diverse real-world entities while maintaining a consistent interface for working with all block types.

## License

[License Information]

## Contributing

[Contribution Guidelines]

## Links

- [Documentation](https://github.com/cyberbrain-sh/blocks)
- [Issues](https://github.com/cyberbrain-sh/blocks/issues)
