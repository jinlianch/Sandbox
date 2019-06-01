// Generated from /Users/jinlianchen/EasilyDo/gitProject/Go/src/Sandbox/go/antlr_test/search_lib/query_parser/Query.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class QueryLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, AND=4, OR=5, NOT=6, ID=7, VALUE=8, STRING=9, HEADER=10, 
		WS=11;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "AND", "OR", "NOT", "ID", "VALUE", "STRING", "HEADER", 
		"ESC", "UNICODE", "HEX", "WS"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'('", "')'", "':'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, "AND", "OR", "NOT", "ID", "VALUE", "STRING", "HEADER", 
		"WS"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public QueryLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Query.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\ru\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3"+
		"\5\3\5\3\5\5\5*\n\5\3\6\3\6\3\6\5\6/\n\6\3\7\3\7\3\7\3\7\5\7\65\n\7\3"+
		"\b\3\b\7\b9\n\b\f\b\16\b<\13\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\tE\n\t\f"+
		"\t\16\tH\13\t\3\t\3\t\3\n\3\n\3\n\7\nO\n\n\f\n\16\nR\13\n\3\n\3\n\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13b\n\13\3\f"+
		"\3\f\3\f\5\fg\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\17\6\17r\n\17\r"+
		"\17\16\17s\2\2\20\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\2\31"+
		"\2\33\2\35\r\3\2\t\4\2\"\"((\6\2//C\\aac|\7\2//\62;C\\aac|\4\2$$^^\n\2"+
		"$$\61\61^^ddhhppttvv\5\2\62;CHch\5\2\13\f\17\17\"\"\2|\2\3\3\2\2\2\2\5"+
		"\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2"+
		"\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\35\3\2\2\2\3\37\3\2\2\2\5!\3"+
		"\2\2\2\7#\3\2\2\2\t)\3\2\2\2\13.\3\2\2\2\r\64\3\2\2\2\17\66\3\2\2\2\21"+
		"=\3\2\2\2\23K\3\2\2\2\25a\3\2\2\2\27c\3\2\2\2\31h\3\2\2\2\33n\3\2\2\2"+
		"\35q\3\2\2\2\37 \7*\2\2 \4\3\2\2\2!\"\7+\2\2\"\6\3\2\2\2#$\7<\2\2$\b\3"+
		"\2\2\2%&\7c\2\2&\'\7p\2\2\'*\7f\2\2(*\t\2\2\2)%\3\2\2\2)(\3\2\2\2*\n\3"+
		"\2\2\2+,\7q\2\2,/\7t\2\2-/\7~\2\2.+\3\2\2\2.-\3\2\2\2/\f\3\2\2\2\60\61"+
		"\7p\2\2\61\62\7q\2\2\62\65\7v\2\2\63\65\7/\2\2\64\60\3\2\2\2\64\63\3\2"+
		"\2\2\65\16\3\2\2\2\66:\t\3\2\2\679\t\4\2\28\67\3\2\2\29<\3\2\2\2:8\3\2"+
		"\2\2:;\3\2\2\2;\20\3\2\2\2<:\3\2\2\2=>\7\u0080\2\2>?\7]\2\2?@\7$\2\2@"+
		"A\7_\2\2AF\3\2\2\2BE\5\27\f\2CE\n\5\2\2DB\3\2\2\2DC\3\2\2\2EH\3\2\2\2"+
		"FD\3\2\2\2FG\3\2\2\2GI\3\2\2\2HF\3\2\2\2IJ\7$\2\2J\22\3\2\2\2KP\7$\2\2"+
		"LO\5\27\f\2MO\n\5\2\2NL\3\2\2\2NM\3\2\2\2OR\3\2\2\2PN\3\2\2\2PQ\3\2\2"+
		"\2QS\3\2\2\2RP\3\2\2\2ST\7$\2\2T\24\3\2\2\2Ub\7%\2\2VW\7j\2\2WX\7g\2\2"+
		"XY\7c\2\2YZ\7f\2\2Z[\7g\2\2[\\\7t\2\2\\]\7]\2\2]^\3\2\2\2^_\5\17\b\2_"+
		"`\7_\2\2`b\3\2\2\2aU\3\2\2\2aV\3\2\2\2b\26\3\2\2\2cf\7^\2\2dg\t\6\2\2"+
		"eg\5\31\r\2fd\3\2\2\2fe\3\2\2\2g\30\3\2\2\2hi\7w\2\2ij\5\33\16\2jk\5\33"+
		"\16\2kl\5\33\16\2lm\5\33\16\2m\32\3\2\2\2no\t\7\2\2o\34\3\2\2\2pr\t\b"+
		"\2\2qp\3\2\2\2rs\3\2\2\2sq\3\2\2\2st\3\2\2\2t\36\3\2\2\2\16\2).\64:DF"+
		"NPafs\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}