const http = require("http");

// Simple HTTP server that handles /danica and returns some message.
http.createServer(function (request, response) {
    if (request.url.startsWith("/danica")) {
        response.writeHead(200, {'Content-Type': 'text/plain'});
        response.end('Hello from Danica!\n');
    } else {
        // not found
        response.writeHead(404, {'Content-Type': 'text/plain'});
        response.end('Not found\n');
    }
}).listen(8080);