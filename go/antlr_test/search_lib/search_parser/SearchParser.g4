parser grammar SearchParser;

options {
	tokenVocab = SearchLexer;
}

query: queryCriterias;

queryCriterias:
	LPAREN queryCriterias RPAREN			# BracketQueryCriterias
	| NOT queryCriterias					# NotQueryCriterias
	| queryCriterias AND? queryCriterias	# AndQueryCriterias
	| queryCriterias OR queryCriterias		# OrQueryCriterias
	| queryCriteria							# SingleQueryCriteria;
// | '{' queryCriterias* '}'				# OrBraceQueryCriterias # Not support now

queryCriteria: fieldCriteria | headerCriteria | valueCriteria;

headerCriteria: HEADER_START valueCriteria;
fieldCriteria: FIELD_START valueCriteria;

/*
 unquoteParts: UNQUOTE_PART+; unquoteText: unquoteParts | (FIELD unquoteParts)+ | (unquoteParts
 FIELD)+ | (FIELD unquoteParts FIELD)+;
 */

// unquoteFT: FIELD_UNQUOTE_START UNQUOTE_PART? EXIT?; unquoteTT: UNQOTE_STRING_START UNQUOTE_PART?
// EXIT?;

// unquoteText: UNQUOTE_PART UNQUOTE_PART? EXIT?;

// unquoteTextPart: UNQUOTE_PART? EXIT?; unquoteText: unquoteFT | unquoteTT;
unquoteText: (FIELD_UNQUOTE_START | UNQOTE_STRING_START) UNQUOTE_PART? EXIT?;

valueCriteria:
	QUOTE_STRING
	| UNQUOTE_STRING
	| unquoteText
	| FIELD;