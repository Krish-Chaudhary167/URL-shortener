# Go URL Shortener

A simple URL shortener built in **Go (Golang)** using HTTP handlers and in-memory storage.


##  Features

1. Generate short URL from a long URL
2. Store mapping of short URL → original URL
3. Redirect using the short URL
4. JSON request/response handling


## How It Works

1. Client sends a POST request with a long URL to `/shorten`
2. Server generates an 8-character hash
3. Mapping is stored in memory
4. Server returns the short URL as JSON
5. Visiting `/redirect/{id}` redirects to the original URL


##  API Endpoints

### ➤ Shorten URL

**POST** `/shorten`

**Request Body**
```json
{
  "url": "https://github.com/Krish-Chaudhary167"
}
