# Simple Image Gallery

A Go-based image generator and gallery service using Firebase.

## Prerequisites

- Go 1.21 or later
- Firebase project
- Firebase service account credentials

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/simple-image-gallery.git
cd simple-image-gallery
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up your environment variables in a `.env` file:
```env
PORT=8080
FIREBASE_PROJECT_ID=your-project-id
FIREBASE_PRIVATE_KEY="your-private-key"
FIREBASE_CLIENT_EMAIL=your-client-email
FIREBASE_CLIENT_ID=your-client-id
```

## Usage

1. Start the server:
```bash
go run main.go
```

2. Generate an image:
```bash
curl -X POST http://localhost:8080/generate \
-H "Content-Type: application/json" \
-d '{"width": 800, "height": 600, "text": "Hello World"}'
```

3. Retrieve an image:
```bash
curl http://localhost:8080/images/{image-id}
```

## API Endpoints

- `POST /generate` - Generate a new image
  - Request body:
    ```json
    {
        "width": 800,
        "height": 600,
        "text": "Hello World"
    }
    ```
  - Response:
    ```json
    {
        "url": "image-url"
    }
    ```

- `GET /images/{id}` - Retrieve an image
  - Returns the image file directly
