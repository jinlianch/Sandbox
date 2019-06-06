lexer grammar SimpleLexer;

AND: 'and' | 'AND';
// ID: [a-zA-Z_\-][a-zA-Z_\-0-9]*;
KEY: [a-zA-Z_\-][a-zA-Z_\-0-9]*;
HEADER: '#' KEY | H E A D E R '[' KEY ']';
HEAD_START: HEADER ':' -> pushMode(VALUEMODE);
PAIR_START: KEY ':' -> pushMode(VALUEMODE);
VALUE_START: ~[ \t\n\r] -> more, pushMode (VALUEMODE);
// NOQUOTE_STRING: ~[ "\t\n\r]+; QUOTE_STRING: '"' ( ESC | .)*? '"';
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];
fragment SPACE: [ \t\n\r]+;
fragment A: [Aa];
fragment D: [Dd];
fragment E: [Ee];
fragment H: [Hh];
fragment R: [Rr];

WS: [ \t\n\r]+ -> skip;

mode VALUEMODE;
VSTRING: '"' ( ESC | .)*? '"' -> popMode;
VNSTRING: ~[ "\t\n\r]+ -> popMode;
VWS: [ \t\n\r]+ -> skip;
// CLOSE: [ \t\n\r]+ -> popMode;

