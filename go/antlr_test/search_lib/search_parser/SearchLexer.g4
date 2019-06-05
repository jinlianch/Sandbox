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
VALUE_START: WS -> pushMode(PAIRMODE);

// NOQUOTE_STRING: ~[ "\t\n\r]+; QUOTE_STRING: '"' ( ESC | .)*? '"';
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];
// 这个地方还有问题，所有字段都会进入这里，为啥呢, 为啥不匹配其他字段，比如or and 这些
fragment WS: [ \t\n\r]+ -> skip;

fragment A: [Aa];
fragment D: [Dd];
fragment E: [Ee];
fragment H: [Hh];
fragment R: [Rr];

mode PAIRMODE;
NOQUOTE_STRING: ~[ "\t\n\r]+ -> popMode;
QUOTE_STRING: '"' ( ESC | .)*? '"' -> popMode;