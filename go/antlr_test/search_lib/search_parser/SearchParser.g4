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
	| queryCriteria							# SingleQueryCriteria;
// | '{' queryCriterias* '}'				# OrBraceQueryCriterias # Not support now

queryCriteria: fieldCriteria | headerCriteria | valueCriteria;

headerCriteria: HEADER_START valueCriteria;
fieldCriteria: FIELD_START valueCriteria;
valueCriteria: QUOTE_STRING | UNQUOTE_STRING | FIELD;

