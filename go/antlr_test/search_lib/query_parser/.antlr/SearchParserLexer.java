// Generated from /Users/jinlianchen/EasilyDo/gitProject/Go/src/Sandbox/go/antlr_test/search_lib/search_parser/SearchParser.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SearchParserLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, AND=4, OR=5, NOT=6, ID=7, STRING=8, HEADER=9, 
		WS=10;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "AND", "OR", "NOT", "ID", "STRING", "HEADER", 
		"ESC", "UNICODE", "HEX", "WS"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'('", "')'", "':'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, "AND", "OR", "NOT", "ID", "STRING", "HEADER", 
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


	public SearchParserLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SearchParser.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\ff\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\5\3\5"+
		"\5\5(\n\5\3\6\3\6\3\6\5\6-\n\6\3\7\3\7\3\7\3\7\5\7\63\n\7\3\b\3\b\7\b"+
		"\67\n\b\f\b\16\b:\13\b\3\t\3\t\3\t\7\t?\n\t\f\t\16\tB\13\t\3\t\3\t\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\nS\n\n\3\13\3\13\3"+
		"\13\5\13X\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\16\6\16c\n\16\r\16\16"+
		"\16d\2\2\17\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\2\27\2\31\2\33"+
		"\f\3\2\t\4\2\"\"((\6\2//C\\aac|\7\2//\62;C\\aac|\4\2$$^^\n\2$$\61\61^"+
		"^ddhhppttvv\5\2\62;CHch\5\2\13\f\17\17\"\"\2k\2\3\3\2\2\2\2\5\3\2\2\2"+
		"\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3"+
		"\2\2\2\2\23\3\2\2\2\2\33\3\2\2\2\3\35\3\2\2\2\5\37\3\2\2\2\7!\3\2\2\2"+
		"\t\'\3\2\2\2\13,\3\2\2\2\r\62\3\2\2\2\17\64\3\2\2\2\21;\3\2\2\2\23R\3"+
		"\2\2\2\25T\3\2\2\2\27Y\3\2\2\2\31_\3\2\2\2\33b\3\2\2\2\35\36\7*\2\2\36"+
		"\4\3\2\2\2\37 \7+\2\2 \6\3\2\2\2!\"\7<\2\2\"\b\3\2\2\2#$\7c\2\2$%\7p\2"+
		"\2%(\7f\2\2&(\t\2\2\2\'#\3\2\2\2\'&\3\2\2\2(\n\3\2\2\2)*\7q\2\2*-\7t\2"+
		"\2+-\7~\2\2,)\3\2\2\2,+\3\2\2\2-\f\3\2\2\2./\7p\2\2/\60\7q\2\2\60\63\7"+
		"v\2\2\61\63\7/\2\2\62.\3\2\2\2\62\61\3\2\2\2\63\16\3\2\2\2\648\t\3\2\2"+
		"\65\67\t\4\2\2\66\65\3\2\2\2\67:\3\2\2\28\66\3\2\2\289\3\2\2\29\20\3\2"+
		"\2\2:8\3\2\2\2;@\7$\2\2<?\5\25\13\2=?\n\5\2\2><\3\2\2\2>=\3\2\2\2?B\3"+
		"\2\2\2@>\3\2\2\2@A\3\2\2\2AC\3\2\2\2B@\3\2\2\2CD\7$\2\2D\22\3\2\2\2EF"+
		"\7%\2\2FS\5\21\t\2GH\7j\2\2HI\7g\2\2IJ\7c\2\2JK\7f\2\2KL\7g\2\2LM\7t\2"+
		"\2MN\7]\2\2NO\3\2\2\2OP\5\21\t\2PQ\7_\2\2QS\3\2\2\2RE\3\2\2\2RG\3\2\2"+
		"\2S\24\3\2\2\2TW\7^\2\2UX\t\6\2\2VX\5\27\f\2WU\3\2\2\2WV\3\2\2\2X\26\3"+
		"\2\2\2YZ\7w\2\2Z[\5\31\r\2[\\\5\31\r\2\\]\5\31\r\2]^\5\31\r\2^\30\3\2"+
		"\2\2_`\t\7\2\2`\32\3\2\2\2ac\t\b\2\2ba\3\2\2\2cd\3\2\2\2db\3\2\2\2de\3"+
		"\2\2\2e\34\3\2\2\2\f\2\',\628>@RWd\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}