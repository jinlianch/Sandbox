antlr4 SearchParser.g4 SearchLexer.g4
javac Search*.java
antrun SearchParser query -tokens -diagnostics
echo "subject:test AND header[to]:jinlian AND test (SUMMARY OR user) label:inbox OR (from)" |  antrun Search query -tokens -diagnostics -gui
