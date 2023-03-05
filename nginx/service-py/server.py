from http.server import BaseHTTPRequestHandler, HTTPServer
import time

hostname = "0.0.0.0"
serverPort = 8091

class MyServer(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-Type", "application/json")
        self.end_headers()
        self.wfile.write(bytes('{"message": "success python"}', 'utf-8'))

if __name__ == "__main__":
    webServer = HTTPServer((hostname, serverPort), MyServer)
    print(f'Server started http://{hostname}:{serverPort}')

    try:
        webServer.serve_forever()
    except KeyboardInterrupt:
        pass

    webServer.server_close()
    print('Server stopped.')
