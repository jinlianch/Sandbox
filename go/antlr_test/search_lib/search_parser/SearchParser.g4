parser grammar SearchParser;

options {
	tokenVocab = SearchLexer;
}

query: queryCriterias;

queryCriterias:
	'(' queryCriterias ')'					# BracketQueryCriterias
	| NOT queryCriterias					# NotQueryCriterias
	| queryCriterias AND? queryCriterias	# AndQueryCriterias
	| queryCriterias OR queryCriterias		# OrQueryCriterias
	| '{' queryCriterias+ '}'				# OrBraceQueryCriterias
	| queryCriteria							# SingleQueryCriteria;

queryCriteria: fieldCriteria | headerCriteria | valueCriteria;

headerCriteria: HEADER_START valueCriteria;
fieldCriteria: FIELD_START valueCriteria;
valueCriteria: FIELD | QUOTE_STRING | NOQUOTE_STRING;
