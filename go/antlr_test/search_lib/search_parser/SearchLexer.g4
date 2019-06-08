lexer grammar SearchLexer;

tokens {
	NOQUOTE_STRING,
	QUOTE_STRING,
	FIELD_START,
	HEADER_START,
	PAREN_START
}

OR: 'OR' | '|';
NOT: 'NOT' | '-';
AND: 'AND' | '&';
// LPAREN: '('; RPAREN: ')';
FIELD: [a-zA-Z_\-][a-zA-Z_\-0-9]*;
HEADER: '#' FIELD | H E A D E R '[' FIELD ']';

PAREN_START: '(' -> type(PAREN_START), pushMode(PARENMODE);
FIELD_START:
	FIELD ':' -> type(FIELD_START), pushMode(PAIRMODE);
HEADER_START:
	HEADER ':' -> type(HEADER_START), pushMode(PAIRMODE);
VALUE_START: ~[ \t\n\r] -> more, pushMode(PAIRMODE);

fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];
fragment A: [Aa];
fragment D: [Dd];
fragment E: [Ee];
fragment H: [Hh];
fragment R: [Rr];

WS: [ \t\n\r]+ -> skip;

mode PAIRMODE;
PQUOTE_STRING:
	'"' (ESC | .)*? '"' -> type(QUOTE_STRING), popMode;
PNOQUOTE_STRING:
	~[ "\t\n\r]+ -> type(NOQUOTE_STRING), popMode;
mode PARENMODE;
INSIDE_OPEN: '(' -> type(PAREN_START), pushMode(PARENMODE);
PAREN_FIELD_START:
	FIELD ':' -> type(FIELD_START), pushMode(PAREN_PAIRMODE);
PAREN_HEADER_START:
	HEADER ':' -> type(HEADER_START), pushMode(PAREN_PAIRMODE);
PAREN_VALUE_START:
	~[ (\t\n\r] -> more, pushMode(PAREN_PAIRMODE);
mode PAREN_PAIRMODE;
PPQUOTE_STRING: '"' (ESC | .)*? '"' -> type(QUOTE_STRING);
PPNOQUOTE_STRING: ~[ "\t\n\r]+ -> type(NOQUOTE_STRING);
CLOSE: ')' -> popMode;