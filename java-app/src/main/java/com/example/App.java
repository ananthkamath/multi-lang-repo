package com.example;

import com.google.gson.Gson;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class App {
    private static final Logger logger = LogManager.getLogger(App.class);

    public static void main(String[] args) {
        Gson gson = new Gson();
        Greeting greeting = new Greeting("Hello, World! (Java)");
        logger.info("Starting java-app");
        System.out.println(gson.toJson(greeting));
    }

    static class Greeting {
        String message;

        Greeting(String message) {
            this.message = message;
        }
    }
}
