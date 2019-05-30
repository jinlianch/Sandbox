grammar SearchParser;

query expression;

expression:
	expression AND expression	# AndExpression
	| expression OR expression	# OrExpression
	| NOT expression			# NotExpression
	| LB expression RB			# BracketExpression
	| single					# SingleExpression;

header_match: HEADER COLON STRING;

single: ID COLON STRING # MatchSingle | STRING # SearchSingle;

AND: 'and' | '&' | ' ';
OR: 'or' | '|';
NOT: 'not' | '-';
LPAR: '(';
RPAR: ')';
COLON: ':';
ID: [a-zA-Z_\-][a-zA-Z_\-0-9]*;
STRING: '"' (ESC | ~["\\])* '"';
HEADER: '#' STRING | 'header[' STRING ']';
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];

WS: [ \t\n\r]+;