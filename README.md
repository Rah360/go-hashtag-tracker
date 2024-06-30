# Hashtag Tracking System

This project is a simple hashtag tracking system built using Go. It allows users to create posts containing hashtags and tracks the count of each hashtag used. The system provides APIs to create posts, retrieve the count of specific hashtags, and list all posts.

## Features

- Create posts with content containing hashtags.
- Track and update hashtag counts in real-time.
- Retrieve the count of a specific hashtag.
- List all posts.

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/Rah360/go-hashtag-tracker
    cd go-hashtag-tracker
    ```

2. **Install Gorilla Mux:**

    ```bash
    go get -u github.com/gorilla/mux
    ```

## Usage

1. **Run the application:**

    ```bash
    go run main.go
    ```

2. **API Endpoints:**

    - **Create a new post:**

        ```bash
        POST /posts
        ```

        **Request Body:**

        ```json
        {
            "content": "This is a test post with #travel and #vacation hashtags"
        }
        ```

        **Response:**

        ```json
        {
            "id": 1,
            "content": "This is a test post with #travel and #vacation hashtags"
        }
        ```

    - **Get the count of a specific hashtag:**

        ```bash
        GET /hashtags/{hashtag}
        ```

        **Example:**

        ```bash
        curl http://localhost:8080/hashtags/%23travel
        ```

        **Response:**

        ```json
        {
            "#travel": 1
        }
        ```

    - **Get all posts:**

        ```bash
        GET /posts
        ```

        **Response:**

        ```json
        [
            {
                "id": 1,
                "content": "This is a test post with #travel and #vacation hashtags"
            }
        ]
        ```

## Code Overview

### Models

`models/post.go`
```go
package models

// Post represents a user's post containing content and an ID
type Post struct {
    ID      int    `json:"id"`
    Content string `json:"content"`
}


