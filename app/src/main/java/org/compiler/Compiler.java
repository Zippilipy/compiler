package org.compiler;

import java.io.*;

import org.compiler.tokens.*;

public class Compiler {
    public static void main(String[] args) {
        String result = "";
        try(BufferedReader bufferedReader = new BufferedReader(new FileReader("src/main/resources/input.xd"))) {
            StringBuilder stringBuilder = new StringBuilder();
            String line = bufferedReader.readLine();
            while (line != null) {
                stringBuilder.append(line);
                stringBuilder.append(System.lineSeparator());
                line = bufferedReader.readLine();
            }
            result = stringBuilder.toString();
        } catch (IOException e) {
            System.out.println("Error");
        }
        TokenWriter tokenWriter = new TokenWriter();
        Exit exit = new Exit(tokenWriter);
        exit.execute();
        tokenWriter.close();
    }
}
