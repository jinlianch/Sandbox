grammar Query;

query: query_item;

query_item:
	query_item AND query_item	# AndQueryItem
	| query_item OR query_item	# OrQueryItem
	| NOT query_item			# NotQueryItem
	| '(' query_item ')'		# BracketQueryItem
	| single_query_item			# SingleQueryItem;

single_query_item:
	ID ':' STRING		# HeaderQueryItem
	| HEADER ':' STRING	# KeyQueryItem
	| STRING			# ValueQueryItem;

AND: 'and' | '&' | ' ';
OR: 'or' | '|';
NOT: 'not' | '-';
// LPAR: '('; RPAR: ')'; COLON: ':';
ID: [a-zA-Z_\-][a-zA-Z_\-0-9]*;
VALUE: '~["]' (ESC | ~["\\])* '"';
STRING: '"' (ESC | ~["\\])* '"';
HEADER: '#' | 'header[' ID ']';
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];

WS: [ \t\n\r]+;