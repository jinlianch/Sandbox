lexer grammar SearchLexer;

tokens {
	QUOTE_STRING,
	UNQUOTE_STRING,
	UNQUOTE_PART,
	LPAREN,
	RPAREN
}

OR: 'OR' | '|';
NOT: 'NOT' | '-' | '!';
AND: 'AND' | '&';
LPAREN_CHAR: '(' -> type(LPAREN);
RPAREN_CHAR: ')' -> type(RPAREN);
// LBRACE: '{'; RBRACE: '}';
COLON: ':';
FIELD: [a-zA-Z_\-][a-zA-Z_\-0-9]*;
HEADER: '#' FIELD | H E A D E R '.' FIELD;
// UNQUOTECHAR: SAFECODEPOINT -> type(UNQUOTE_STRING);

FIELD_START: FIELD ':' -> pushMode(PAIR_MODE);
HEADER_START: HEADER ':' -> pushMode(PAIR_MODE);
QOUTE_STRING_START: '"' -> more, pushMode(QUOTE_MODE);
/*
 FIELD_UNQUOTE_START: FIELD UNQUOTECHAR -> type(UNQUOTE_PART), pushMode(UNQUOTE_MODE); //
 UNQOTE_STRING_START: ~[" \t\n\r()] -> pushMode(UNQUOTE_MODE); UNQOTE_STRING_START: UNQUOTECHAR ->
 type(UNQUOTE_PART), pushMode(UNQUOTE_MODE); // ~[" \t\n\r()] -> type(UNQUOTE_STRING),
 pushMode(UNQUOTE_MODE);
 */
FIELD_UNQUOTE_START:
	FIELD UNQUOTECHAR -> pushMode(UNQUOTE_MODE);
UNQOTE_STRING_START: UNQUOTECHAR -> pushMode(UNQUOTE_MODE);
// UNQOTE_STRING_START: ~[" \t\n\r()] -> pushMode(UNQUOTE_MODE);

fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];
fragment SAFECODEPOINT: ~ ["\\\u0000-\u001F];
fragment UNQUOTECHAR: ~[" \t\n\r()];
fragment A: [Aa];
fragment D: [Dd];
fragment E: [Ee];
fragment H: [Hh];
fragment R: [Rr];

WS: [ \t\n\r]+ -> skip;

mode PAIR_MODE;
PAIR_QS:
	'"' (ESC | SAFECODEPOINT)* '"' -> type(QUOTE_STRING), popMode;
PAIR_UNQS: UNQUOTECHAR+ -> type(UNQUOTE_STRING), popMode;

mode QUOTE_MODE;
QUOTE_S: (ESC | SAFECODEPOINT)* '"' -> type(QUOTE_STRING), popMode;

mode UNQUOTE_MODE;
RP: RPAREN_CHAR -> type(RPAREN), popMode;
EXIT: [" \t\n\r(] -> popMode;
UNQUOTE_S: UNQUOTECHAR+ -> type(UNQUOTE_PART), popMode;
// UNQOTE_S: ~[" \t\n\r()]+ -> type(UNQUOTE_STRING), popMode;
