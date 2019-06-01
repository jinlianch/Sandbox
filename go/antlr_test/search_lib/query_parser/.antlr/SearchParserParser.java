// Generated from /Users/jinlianchen/EasilyDo/gitProject/Go/src/Sandbox/go/antlr_test/search_lib/search_parser/SearchParser.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SearchParserParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, AND=4, OR=5, NOT=6, ID=7, STRING=8, HEADER=9, 
		WS=10;
	public static final int
		RULE_query = 0, RULE_query_item = 1, RULE_single_query_item = 2;
	public static final String[] ruleNames = {
		"query", "query_item", "single_query_item"
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

	@Override
	public String getGrammarFileName() { return "SearchParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SearchParserParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class QueryContext extends ParserRuleContext {
		public Query_itemContext query_item() {
			return getRuleContext(Query_itemContext.class,0);
		}
		public QueryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_query; }
	}

	public final QueryContext query() throws RecognitionException {
		QueryContext _localctx = new QueryContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_query);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(6);
			query_item(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Query_itemContext extends ParserRuleContext {
		public Query_itemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_query_item; }
	 
		public Query_itemContext() { }
		public void copyFrom(Query_itemContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class BracketQueryItemContext extends Query_itemContext {
		public Query_itemContext query_item() {
			return getRuleContext(Query_itemContext.class,0);
		}
		public BracketQueryItemContext(Query_itemContext ctx) { copyFrom(ctx); }
	}
	public static class SingleQueryItemContext extends Query_itemContext {
		public Single_query_itemContext single_query_item() {
			return getRuleContext(Single_query_itemContext.class,0);
		}
		public SingleQueryItemContext(Query_itemContext ctx) { copyFrom(ctx); }
	}
	public static class NotQueryItemContext extends Query_itemContext {
		public TerminalNode NOT() { return getToken(SearchParserParser.NOT, 0); }
		public Query_itemContext query_item() {
			return getRuleContext(Query_itemContext.class,0);
		}
		public NotQueryItemContext(Query_itemContext ctx) { copyFrom(ctx); }
	}
	public static class OrQueryItemContext extends Query_itemContext {
		public List<Query_itemContext> query_item() {
			return getRuleContexts(Query_itemContext.class);
		}
		public Query_itemContext query_item(int i) {
			return getRuleContext(Query_itemContext.class,i);
		}
		public TerminalNode OR() { return getToken(SearchParserParser.OR, 0); }
		public OrQueryItemContext(Query_itemContext ctx) { copyFrom(ctx); }
	}
	public static class AndQueryItemContext extends Query_itemContext {
		public List<Query_itemContext> query_item() {
			return getRuleContexts(Query_itemContext.class);
		}
		public Query_itemContext query_item(int i) {
			return getRuleContext(Query_itemContext.class,i);
		}
		public TerminalNode AND() { return getToken(SearchParserParser.AND, 0); }
		public AndQueryItemContext(Query_itemContext ctx) { copyFrom(ctx); }
	}

	public final Query_itemContext query_item() throws RecognitionException {
		return query_item(0);
	}

	private Query_itemContext query_item(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Query_itemContext _localctx = new Query_itemContext(_ctx, _parentState);
		Query_itemContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_query_item, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(16);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				{
				_localctx = new NotQueryItemContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(9);
				match(NOT);
				setState(10);
				query_item(3);
				}
				break;
			case T__0:
				{
				_localctx = new BracketQueryItemContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(11);
				match(T__0);
				setState(12);
				query_item(0);
				setState(13);
				match(T__1);
				}
				break;
			case ID:
			case STRING:
			case HEADER:
				{
				_localctx = new SingleQueryItemContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(15);
				single_query_item();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(26);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(24);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new AndQueryItemContext(new Query_itemContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_query_item);
						setState(18);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(19);
						match(AND);
						setState(20);
						query_item(6);
						}
						break;
					case 2:
						{
						_localctx = new OrQueryItemContext(new Query_itemContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_query_item);
						setState(21);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(22);
						match(OR);
						setState(23);
						query_item(5);
						}
						break;
					}
					} 
				}
				setState(28);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Single_query_itemContext extends ParserRuleContext {
		public Single_query_itemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_single_query_item; }
	 
		public Single_query_itemContext() { }
		public void copyFrom(Single_query_itemContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class HeaderQueryItemContext extends Single_query_itemContext {
		public TerminalNode ID() { return getToken(SearchParserParser.ID, 0); }
		public TerminalNode STRING() { return getToken(SearchParserParser.STRING, 0); }
		public HeaderQueryItemContext(Single_query_itemContext ctx) { copyFrom(ctx); }
	}
	public static class ValueQueryItemContext extends Single_query_itemContext {
		public TerminalNode STRING() { return getToken(SearchParserParser.STRING, 0); }
		public ValueQueryItemContext(Single_query_itemContext ctx) { copyFrom(ctx); }
	}
	public static class KeyQueryItemContext extends Single_query_itemContext {
		public TerminalNode HEADER() { return getToken(SearchParserParser.HEADER, 0); }
		public TerminalNode STRING() { return getToken(SearchParserParser.STRING, 0); }
		public KeyQueryItemContext(Single_query_itemContext ctx) { copyFrom(ctx); }
	}

	public final Single_query_itemContext single_query_item() throws RecognitionException {
		Single_query_itemContext _localctx = new Single_query_itemContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_single_query_item);
		try {
			setState(36);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				_localctx = new HeaderQueryItemContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(29);
				match(ID);
				setState(30);
				match(T__2);
				setState(31);
				match(STRING);
				}
				break;
			case HEADER:
				_localctx = new KeyQueryItemContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(32);
				match(HEADER);
				setState(33);
				match(T__2);
				setState(34);
				match(STRING);
				}
				break;
			case STRING:
				_localctx = new ValueQueryItemContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(35);
				match(STRING);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return query_item_sempred((Query_itemContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean query_item_sempred(Query_itemContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 5);
		case 1:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\f)\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3\23\n\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\7\3\33\n\3\f\3\16\3\36\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\5\4\'\n\4\3\4\2\3\4\5\2\4\6\2\2\2+\2\b\3\2\2\2\4\22\3\2\2\2\6&\3\2"+
		"\2\2\b\t\5\4\3\2\t\3\3\2\2\2\n\13\b\3\1\2\13\f\7\b\2\2\f\23\5\4\3\5\r"+
		"\16\7\3\2\2\16\17\5\4\3\2\17\20\7\4\2\2\20\23\3\2\2\2\21\23\5\6\4\2\22"+
		"\n\3\2\2\2\22\r\3\2\2\2\22\21\3\2\2\2\23\34\3\2\2\2\24\25\f\7\2\2\25\26"+
		"\7\6\2\2\26\33\5\4\3\b\27\30\f\6\2\2\30\31\7\7\2\2\31\33\5\4\3\7\32\24"+
		"\3\2\2\2\32\27\3\2\2\2\33\36\3\2\2\2\34\32\3\2\2\2\34\35\3\2\2\2\35\5"+
		"\3\2\2\2\36\34\3\2\2\2\37 \7\t\2\2 !\7\5\2\2!\'\7\n\2\2\"#\7\13\2\2#$"+
		"\7\5\2\2$\'\7\n\2\2%\'\7\n\2\2&\37\3\2\2\2&\"\3\2\2\2&%\3\2\2\2\'\7\3"+
		"\2\2\2\6\22\32\34&";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}