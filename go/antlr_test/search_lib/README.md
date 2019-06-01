Base on antlr4(https://github.com/antlr/antlr4)

search_parser:

antlr4 -Dlanguage=Go Query.g4 -package query_parser -visitor
