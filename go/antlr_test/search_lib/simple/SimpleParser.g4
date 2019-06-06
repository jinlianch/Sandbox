parser grammar SimpleParser;
options {
	tokenVocab = SimpleLexer;
}

query: pairs;
// pairs: pairs (WS+ pairs)+ | pair;
pairs: pairs AND? pairs # AndQueryItem | pair # SingleQueryItem;

pair: keypair # KeyItem | item # ValueItem;

// keypair: id ':' item;
keypair: kvpair | headerpair;
// headerpair: HEAD_START VSTRING | HEAD_START VNSTRING; kvpair: PAIR_START VSTRING | PAIR_START
// VNSTRING;
headerpair: HEAD_START value;
kvpair: PAIR_START value;
value: VSTRING | VNSTRING;
//  id: ID;

item: VSTRING | VNSTRING | KEY;

