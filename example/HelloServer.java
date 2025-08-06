import com.sun.net.httpserver.HttpServer;

import java.io.IOException;
import java.net.InetSocketAddress;

public class HelloServer {
    public static void main(String[] args) throws IOException {
        int serverPort = 8080;
        var server = HttpServer.create(new InetSocketAddress(serverPort), 0);

        server.createContext("/", exchange -> {
            var response = "Hello, World!";

            exchange.sendResponseHeaders(200, response.getBytes().length);

            try(var os = exchange.getResponseBody()) {
                os.write(response.getBytes());
            }
        });

        server.start();
        System.out.println("Listening :" + serverPort);
    }
}

