lexer grammar SearchLexer;

tokens {
	QUOTE_STRING,
	UNQUOTE_STRING
}

OR: 'OR' | '|';
NOT: 'NOT' | '-';
AND: 'AND' | '&';
LPAREN: '(';
RPAREN: ')';
// LBRACE: '{'; RBRACE: '}';
COLON: ':';
FIELD: [a-zA-Z_\-][a-zA-Z_\-0-9]*;
HEADER: '#' FIELD | H E A D E R '.' FIELD;

FIELD_START: FIELD ':' -> pushMode(PAIRMODE);
HEADER_START: HEADER ':' -> pushMode(PAIRMODE);
QOUTE_STRING_START: '"' -> more, pushMode(QUOTEMODE);
UNQOUTE_STRING_START:
	UNQUOTECHAR -> more, pushMode(UNQUOTEMODE);

fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];
fragment SAFECODEPOINT: ~ ["\\\u0000-\u001F];
fragment UNQUOTECHAR: ~[" \t\n\r(){}];
fragment A: [Aa];
fragment D: [Dd];
fragment E: [Ee];
fragment H: [Hh];
fragment R: [Rr];

WS: [ \t\n\r]+ -> skip;

mode PAIRMODE;
PAIR_QS:
	'"' (ESC | SAFECODEPOINT)* '"' -> type(QUOTE_STRING), popMode;
PAIR_UNS: UNQUOTECHAR+ -> type(UNQUOTE_STRING), popMode;

mode QUOTEMODE;
QUOTE_S: (ESC | SAFECODEPOINT)* '"' -> type(QUOTE_STRING), popMode;

mode UNQUOTEMODE;
UNQUOTE_S: UNQUOTECHAR+ -> type(UNQUOTE_STRING);
// END: [" \t\n\r(){}] -> type(UNQUOTE_STRING), popMode;
END: . -> type(UNQUOTE_STRING), popMode;