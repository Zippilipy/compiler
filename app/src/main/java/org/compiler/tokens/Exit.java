package org.compiler.tokens;

public class Exit extends Token {
    public Exit(TokenWriter tokenWriter) {
        super(tokenWriter);
    }

    @Override
    protected String doTheThing() {
        return "Exit 60";
    }
}
