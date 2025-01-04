package org.compiler.tokens;

public abstract class Token {
    protected TokenWriter tokenWriter;

    public Token(TokenWriter tokenWriter) {
        this.tokenWriter = tokenWriter;
    }

    public void execute() {
        tokenWriter.append(doTheThing());
    }

    protected abstract String doTheThing();
}
