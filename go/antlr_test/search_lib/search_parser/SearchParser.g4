parser grammar SearchParser;

options {
	tokenVocab = SearchLexer;
}

query: queryCriterias;

queryCriterias:
	queryCriterias AND? queryCriterias	# AndQueryCriterias
	| queryCriterias OR queryCriterias	# OrQueryCriterias
	| NOT queryCriterias				# NotQueryCriterias
	| PAREN_START queryCriterias CLOSE	# BracketQueryCriterias
	// | '(' queryCriterias ')'			# BracketQueryCriterias
	| queryCriteria # SingleQueryCriteria;

queryCriteria: fieldCriteria | headerCriteria | valueCriteria;

headerCriteria: HEADER_START valueCriteria;
fieldCriteria: FIELD_START valueCriteria;
valueCriteria: QUOTE_STRING | NOQUOTE_STRING | FIELD;
