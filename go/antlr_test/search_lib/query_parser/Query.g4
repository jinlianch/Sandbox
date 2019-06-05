grammar Query;

query: queryCriterias;

/*
 queryCriterias: andQueryCriteria | orQueryCriteria | bracketQueryCriteria | notQueryCriteria |
 queryCriteria;
 */
queryCriterias:
	queryCriterias AND queryCriterias	# AndQueryCriterias
	| queryCriterias OR queryCriterias	# OrQueryCriterias
	| NOT queryCriterias				# NotQueryCriterias
	| '(' queryCriterias ')'			# BracketQueryCriterias
	| queryCriteria						# SingleQueryCriteria;

/*
 andQueryCriteria: queryCriterias AND queryCriterias; orQueryCriteria: queryCriterias OR
 queryCriterias; bracketQueryCriteria: '(' queryCriterias ')'; notQueryCriteria: NOT queryCriterias;
 */
queryCriteria: keyCriteria | headerCriteria | valueCriteria;

/*
 ID ':' valueCriteria | HEADER ':' valueCriteria | value_item;
 */

keyCriteria: KEY ':' valueCriteria;
headerCriteria: HEADER ':' valueCriteria;
valueCriteria: KEY | QUOTE_STRING | NOQUOTE_STRING;

OR: 'or' | '|';
NOT: 'not' | '-';
AND: 'and' | '&';
// KEY: [a-zA-Z_\-][a-zA-Z_\-0-9]*;
KEY: ~[ :"\t\n\r]+;
HEADER: '#' KEY | 'header[' KEY ']';
NOQUOTE_STRING: ~[ "\t\n\r]+;
QUOTE_STRING: '"' ( ESC | .)*? '"';
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];

WS: [ \t\n\r]+ -> skip;