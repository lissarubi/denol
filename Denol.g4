grammar Denol;

INT    : [0-9]+;
DOUBLE : [0-9]+'.'[0-9]+;
PI     : 'pi';
POW    : '^';
NL     : '\n';
WS     : [ \t\r]+ -> skip;
ID     : [a-zA-Z_][a-zA-Z_0-9]*;

SPACE : ' ';

STRING : '"' ~('\r' | '\n' | '"')* '"' | '\'' ~('\r' | '\n' | '\'')* '\'';


paren_expr
   : '(' expr ')'
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
   | INT
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
