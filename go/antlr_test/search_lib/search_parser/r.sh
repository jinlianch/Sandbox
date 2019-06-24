antlr4 SearchParser.g4 SearchLexer.g4
javac Search*.java
antrun SearchParser query -tokens -diagnostics
echo "subject:test AND header[to]:jinlian AND test (SUMMARY OR user) label:inbox OR (from)" |  antrun Search query -tokens -diagnostics -gui
echo "\"hello" | antrun Search valueCriteria -tokens -diagnostics -gui
echo "subject:\"hello" | antrun Search query -tokens -diagnostics -gui
echo "subject:test AND header.to:jinlian AND test (SUMMARY OR user) NOT label:inbox OR (from)" |  antrun Search query -tokens -diagnostics -gui
echo "subject:test AND header.to:jinlian AND {test (SUMMARY OR user) NOT label:inbox OR ("from}")}" |  antrun Search query -tokens -diagnostics -gui
