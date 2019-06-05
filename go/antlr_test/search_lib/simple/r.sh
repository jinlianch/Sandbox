antlr4 SimpleParser.g4 SimpleLexer.g4
javac Simple*.java
antrun Simple query -tokens -diagnostics t.txt
