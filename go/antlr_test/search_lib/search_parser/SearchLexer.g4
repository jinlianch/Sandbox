lexer grammar SearchLexer;

OR: 'OR' | '|';
NOT: 'NOT' | '-';
AND: 'AND' | '&';
LP: '(';
RP: ')';
FIELD: [a-zA-Z_\-][a-zA-Z_\-0-9]*;
HEADER: '#' FIELD | H E A D E R '[' FIELD ']';

FIELD_START: FIELD ':' -> pushMode(PAIRMODE);
HEADER_START: HEADER ':' -> pushMode(PAIRMODE);
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
NOQUOTE_STRING: ~[ "\t\n\r]+ -> popMode;
QUOTE_STRING: '"' ( ESC | .)*? '"' -> popMode;