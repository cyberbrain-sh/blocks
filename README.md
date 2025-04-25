# Blocks

A powerful and flexible block-based content management system developed by [Cyberbrain](https://github.com/cyberbrain-sh).

## Overview

Blocks is a Go library that provides a versatile foundation for managing structured content as interconnected blocks. It enables the creation, manipulation, and organization of various content types within a unified data model. The system supports a wide range of content types including text, to-do items, links, emails, media (YouTube, Instagram), movies, and series information.

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
import "github.com/cyberbrain-sh/blocks/pkg"

// Create a new empty block
block := pkg.NewEmptyBlock()

// Set properties
block.Properties.ReplaceValue(pkg.PropertyKeyTitle, "My Block")
block.Properties.ReplaceValue(pkg.PropertyKeyText, "This is a sample block")

// Add a child block
childID := uuid.New()
block.AppendChild(childID)

// Change block type
block.Type = pkg.TypeParagraph
```

## Block Types

The library supports various block types including:

- **Structural Blocks**: Page, Fragment, Database
- **Text Blocks**: Paragraph, Headers (1-6), Bullet List Items, Numbered List Items
- **Task Blocks**: To-Do Items
- **Media Blocks**: Image, Video, Audio, File, YouTube, Instagram
- **Reference Blocks**: Link, Email, Movie, Series

## Properties

Blocks have type-specific properties that can be accessed and modified:

```go
// Get a string property
if title, exists := block.Properties.GetString(pkg.PropertyKeyTitle); exists {
    fmt.Println("Title:", title)
}

// Set a property
block.Properties.ReplaceValue(pkg.PropertyKeyChecked, true)
```

## License

[License Information]

## Contributing

[Contribution Guidelines]

## Links

- [Documentation](https://github.com/cyberbrain-sh/blocks)
- [Issues](https://github.com/cyberbrain-sh/blocks/issues)
