grammar Denol;

NUMBER : [0-9]+ | [0-9]+'.'[0-9]+;
PI     : 'pi';
POW    : '^';
NL     : '\n' | '\r' | '\r\n';
WS     : [ n\t\r]+;
TWOPOINS: ':';
ID     : [a-zA-Z_][a-zA-Z_0-9]*;

SPACE : ' ';

STRING : '"' ~('\r' | '\n' | '"')* '"' | '\'' ~('\r' | '\n' | '\'')* '\'';

PAREN_START: '(';
SQRT_START: '[';
KEYS_START: '{';

paren_expr
   : '(' expr ')'
   ;
sqrt_expr
   : '[' expr ']'
   ;
keys_expr
   : '{' expr '}'
   ;
expr
   : test
   | ID '==' expr ';' | ID '>=' expr ';' | ID '<=' expr ';' | ID '>' expr ';'| ID '<' expr ';'| ID '=' expr ';'| '&&' | '||'
   ;
test
   : sum
   | sum '<' sum
   ;
sum
   : term
   | sum '+' term
   | sum '-' term
   ;
term
   : ID
   | NUMBER
   | paren_expr
   ;

statement
   : 'salve' paren_expr statement
   | 'salve' paren_expr statement 'naosalvou' statement
   | 'zoeira' paren_expr statement
   | 'pamonhosa' paren_expr statement
   | 'day' statement 'zoeira' paren_expr ';'
   | '{' statement* '}'
   | expr ';'
   | ';'
   ;
