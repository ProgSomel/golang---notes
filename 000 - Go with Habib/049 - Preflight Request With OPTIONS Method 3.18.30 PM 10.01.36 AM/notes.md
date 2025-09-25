# 049 - Preflight Request With OPTIONS Method
## What is a preflight request?
When a browser wants to call your API **from another origin** (e.g., UI at http://localhost:3000 → API at http://localhost:8080) and the request is “non-simple”, it first sends a **preflight** request:
- Method: OPTIONS
- Purpose: Ask permission to make the real request
- Sent automatically by the browser (you never code it on the client)

A request is **non-simple** (thus needs preflight) if **any** of these are true:
- Method is not GET, HEAD, or POST
- POST but **Content-Type** is not one of: **application/x-www-form-urlencoded**, **multipart/form-data**, **text/plain**
- You send **custom headers** (e.g., **Authorization**, **X-Requested-With**, **Content-Type: application/json)**
- You send **credentials** (cookies/Authorization) + certain other conditions

## What the preflight looks like
Browser → Server (Preflight):
```bash
OPTIONS /v1/users
Origin: http://localhost:3000
Access-Control-Request-Method: POST
Access-Control-Request-Headers: Content-Type, Authorization
```

Server → Browser (Preflight response):
```bash
HTTP/1.1 200 OK
Access-Control-Allow-Origin: http://localhost:3000
Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: Content-Type, Authorization
Access-Control-Max-Age: 600
```
If the headers “allow” it, the browser proceeds to send the actual request (e.g., POST /v1/users).

## The key CORS headers (server side)
- **Access-Control-Allow-Origin**: which origin may call you (* or a specific origin)
- **Access-Control-Allow-Methods**: allowed HTTP methods
- **Access-Control-Allow-Headers**: which request headers are allowed
- **Access-Control-Max-Age**: how long (seconds) the browser may cache the preflight result
- **Access-Control-Allow-Credentials**: true if you allow cookies/credentials (then cannot use * for origin)

## Quick checklist
- Add middleware that sets CORS headers
- Short-circuit OPTIONS with 200 OK
- Decide origin policy (* vs exact origin)
- If credentials needed: Allow-Credentials: true + specific origin + Vary: Origin
- Set Allow-Methods, Allow-Headers, Max-Age
- Test with curl and DevTools (Network tab)