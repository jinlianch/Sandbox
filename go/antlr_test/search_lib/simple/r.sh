antlr4 SimpleParser.g4 SimpleLexer.g4
javac Simple*.java
antrun Simple query -tokens -diagnostics t.txt
echo "header[xxx]: 345a:32   123  234aa" |  antrun Simple query -tokens -diagnostics
