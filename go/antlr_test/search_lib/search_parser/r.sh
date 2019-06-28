antlr4 SearchParser.g4 SearchLexer.g4
javac Search*.java
antrun Search query -tokens -diagnostics -gui
echo "subject:test AND header.to:jinlian AND test (SUMMARY OR user) label:inbox OR (from)" |  antrun Search query -tokens -diagnostics -gui
echo "subject:test AND header.to:jinlian AND test (SUMMARY OR user) NOT label:inbox OR (from)" |  antrun Search query -tokens -diagnostics -gui
echo "subject:test AND (header.to:jinlian AND test (SUMMARY OR user)) NOT label:inbox OR (from)" | antrun Search query -tokens -diagnostics -gui
echo "subject:test AND (header.to:jinlian AND test (SUMMARY OR user) NOT label:inbox) OR (from)" | antrun Search query -tokens -diagnostics -gui
# echo "subject:test AND header.to:jinlian AND {test (SUMMARY OR user) NOT label:inbox OR ("from}")}" |  antrun Search query -tokens -diagnostics -gui
echo "\"hello" | antrun Search query -tokens -diagnostics -gui
echo "subject:\"hello" | antrun Search query -tokens -diagnostics -gui
#subject:test AND header.to:jinlian AND {test (SUMMARY OR user) NOT label:inbox OR ("from}")}
subject:"hello world"
subject:"hello \"world" AND header.to:"hello"
{subject:"hello world" bcc:jinlian}
"from}"
"test\"}"
 (subject:"tes)t")
cc:hello AND bcc:"w3r 23" "hello eee" hello
"subject"
subject:\"test
 \"hello

