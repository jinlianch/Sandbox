antlr4 SearchParser.g4 SearchLexer.g4
javac Search*.java
antrun SearchParser query -tokens -diagnostics
