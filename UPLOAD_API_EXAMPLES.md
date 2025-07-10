# Image Upload API Documentation

## Endpoint
```
POST /api/upload/image
```

## Description
Upload an image file to the server. The uploaded image will be stored in the `uploads/` directory and accessible via static URL.

## Request Format
- **Content-Type**: `multipart/form-data`
- **Field Name**: `image`
- **Supported Formats**: JPG, JPEG, PNG, GIF, WebP
- **Max File Size**: 5MB
- **Max Form Size**: 10MB

## Response Format

### Success Response (200 OK)
```json
{
  "message": "Image uploaded successfully",
  "url": "/uploads/1703123456_AbCdEfGh.jpg",
  "filename": "1703123456_AbCdEfGh.jpg",
  "size": 1048576
}
```

### Error Responses

#### 400 Bad Request - No file provided
```json
{
  "error": "No image file provided"
}
```

#### 400 Bad Request - Invalid file type
```json
{
  "error": "Invalid file type. Only JPG, JPEG, PNG, GIF, WebP are allowed"
}
```

#### 400 Bad Request - File too large
```json
{
  "error": "File size too large. Maximum 5MB allowed"
}
```

#### 400 Bad Request - Form parsing error
```json
{
  "error": "Failed to parse form"
}
```

#### 500 Internal Server Error
```json
{
  "error": "Failed to create uploads directory"
}
```

## cURL Examples

### Upload a JPG image
```bash
curl -X POST http://localhost:8080/api/upload/image \
  -F "image=@/path/to/your/image.jpg"
```

### Upload a PNG image
```bash
curl -X POST http://localhost:8080/api/upload/image \
  -F "image=@/path/to/your/image.png"
```

### Upload with verbose output
```bash
curl -X POST http://localhost:8080/api/upload/image \
  -F "image=@/path/to/your/image.jpg" \
  -v
```

### Save response to file
```bash
curl -X POST http://localhost:8080/api/upload/image \
  -F "image=@/path/to/your/image.jpg" \
  -o upload_response.json
```

## Access Uploaded Images

Once uploaded, images can be accessed via:
```
GET http://localhost:8080/uploads/{filename}
```

Example:
```bash
curl http://localhost:8080/uploads/1703123456_AbCdEfGh.jpg
```

## JavaScript/Fetch Example

```javascript
const uploadImage = async (file) => {
  const formData = new FormData();
  formData.append('image', file);

  try {
    const response = await fetch('http://localhost:8080/api/upload/image', {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Upload failed');
    }

    const result = await response.json();
    console.log('Upload successful:', result);
    return result;
  } catch (error) {
    console.error('Upload error:', error);
    throw error;
  }
};

// Usage
const fileInput = document.getElementById('file-input');
fileInput.addEventListener('change', async (e) => {
  const file = e.target.files[0];
  if (file) {
    await uploadImage(file);
  }
});
```

## Security Features

1. **File Type Validation**: Only image files are accepted
2. **File Size Limits**: Maximum 5MB per file
3. **Unique Filenames**: Prevents filename conflicts using timestamp + random string
4. **MIME Type Checking**: Validates both file extension and MIME type
5. **Directory Creation**: Automatically creates uploads directory if it doesn't exist

## Notes

- Uploaded files are stored in `backend/uploads/` directory
- Filenames are automatically generated to prevent conflicts: `{timestamp}_{random8chars}.{extension}`
- The API returns the relative URL path, prepend your server URL to get the full URL
- Make sure the `uploads/` directory has proper write permissions 