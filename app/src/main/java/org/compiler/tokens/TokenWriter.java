package org.compiler.tokens;

import java.io.BufferedWriter;
import java.io.FileWriter;
import java.io.IOException;

public class TokenWriter {
    private String output = "Testing";
    private BufferedWriter bufferedWriter;

    public TokenWriter() {
        try {
            new BufferedWriter(new FileWriter("src/main/resources/output.asm")).write("");
            bufferedWriter = new BufferedWriter(new FileWriter("src/main/resources/output.asm", true));
        } catch (Exception e) {
            System.out.println("Error");
        }
    }

    public void append(String string){
        try {
            bufferedWriter.append(string);
        } catch (IOException e) {
            System.out.println("Could not write to file!");
        }
    }

    public void close() {
        try {
            bufferedWriter.close();
        } catch (IOException e) {
            System.out.println("Could not close file!");
        }
    }
}
