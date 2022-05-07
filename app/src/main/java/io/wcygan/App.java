package io.wcygan;

public class App {
    public String getGreeting() {
        return "Hello World!";
    }

    public static void main(String[] args) {
        for (int i = 0; i < 10; i++) {
            System.out.println(new App().getGreeting());
        }
    }
}
