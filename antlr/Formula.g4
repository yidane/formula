grammar Formula;

options{language=go;}

@header {
import "github.com/yidane/formula/opt"
}

@members {
}

@init {
    numberFormatInfo.NumberDecimalSeparator = ".";
}
ncalc returns[*opt.LogicalExpression retValue]
    : expr EOF { $retValue = $expr.retValue}
    ;

//三元操作
 expr returns[*opt.LogicalExpression retValue]
    : first=orExpr '?' middle=expr ':' right=expr   { $retValue = opt.NewTernaryExpression($first.retValue, $middle.retValue, $right.retValue)}
    | orExpr                                        { $retValue = $orExpr.retValue}
    ;

orExpr returns[*opt.LogicalExpression retValue]
    : first=orExpr ('||'|'or') andExpr      { $retValue = opt.NewBinaryExpression("||", $first.retValue, $andExpr.retValue)}
    | andExpr                               { $retValue = $andExpr.retValue}
    ;

andExpr returns[*opt.LogicalExpression retValue]
    : first=andExpr ('&&'|'and') bitOrExpr  { $retValue = opt.NewBinaryExpression("&&", $first.retValue, $bitOrExpr.retValue)}
    | bitOrExpr                             { $retValue = $bitOrExpr.retValue}
    ;

bitOrExpr returns[*opt.LogicalExpression retValue]
    : first=bitOrExpr '|' bitXorExpr    { $retValue = opt.NewBinaryExpression("|", $first.retValue, $bitXorExpr.retValue)}
    | bitXorExpr                        { $retValue = $bitXorExpr.retValue}
    ;

bitXorExpr returns[*opt.LogicalExpression retValue]
    : first=bitXorExpr '^' bitAndExpr   { $retValue = opt.NewBinaryExpression("^", $first.retValue, $bitAndExpr.retValue)}
    | bitAndExpr                        { $retValue = $bitAndExpr.retValue}
    ;

bitAndExpr returns[*opt.LogicalExpression retValue]
    : first=bitAndExpr '&' eqExpr       { $retValue = opt.NewBinaryExpression("&", $first.retValue, $eqExpr.retValue)}
    | eqExpr                            { $retValue = $eqExpr.retValue}
    ;

eqExpr returns[*opt.LogicalExpression retValue]
    : first=eqExpr ('=='|'=')  relExpr  { $retValue = opt.NewBinaryExpression("==", $first.retValue, $relExpr.retValue)}
    | first=eqExpr ('!='|'<>') relExpr  { $retValue = opt.NewBinaryExpression("!=", $first.retValue, $relExpr.retValue)}
    | relExpr                           { $retValue = $relExpr.retValue}
    ;

relExpr returns[*opt.LogicalExpression retValue]
    : first=relExpr '<'  shiftExpr      { $retValue = opt.NewBinaryExpression("<", $first.retValue, $shiftExpr.retValue)}
    | first=relExpr '<=' shiftExpr      { $retValue = opt.NewBinaryExpression("<=", $first.retValue, $shiftExpr.retValue)}
    | first=relExpr '>'  shiftExpr      { $retValue = opt.NewBinaryExpression(">", $first.retValue, $shiftExpr.retValue)}
    | first=relExpr '>=' shiftExpr      { $retValue = opt.NewBinaryExpression(">=", $first.retValue, $shiftExpr.retValue)}
    | shiftExpr                         { $retValue = $shiftExpr.retValue}
    ;

shiftExpr returns[*opt.LogicalExpression retValue]
    : first=shiftExpr '<<' addExpr      { $retValue = opt.NewBinaryExpression("<<", $first.retValue, $addExpr.retValue)}
    | first=shiftExpr '>>' addExpr      { $retValue = opt.NewBinaryExpression(">>", $first.retValue, $addExpr.retValue)}
    | addExpr                           { $retValue = $addExpr.retValue}
    ;

addExpr returns[*opt.LogicalExpression retValue]
    : first=addExpr '+' multExpr        { $retValue = opt.NewBinaryExpression("+", $first.retValue, $multExpr.retValue)}
    | first=addExpr '-' multExpr        { $retValue = opt.NewBinaryExpression("-", $first.retValue, $multExpr.retValue)}
    | multExpr                          { $retValue = $multExpr.retValue}
    ;

multExpr returns[*opt.LogicalExpression retValue]
    : first=multExpr '*' unaryExpr      { $retValue = opt.NewBinaryExpression("*", $first.retValue, $unaryExpr.retValue)}
    | first=multExpr '/' unaryExpr      { $retValue = opt.NewBinaryExpression("/", $first.retValue, $unaryExpr.retValue)}
    | first=multExpr '%' unaryExpr      { $retValue = opt.NewBinaryExpression("%", $first.retValue, $unaryExpr.retValue)}
    | unaryExpr                         { $retValue = $unaryExpr.retValue}
    ;
    
//一元操作
unaryExpr returns[*opt.LogicalExpression retValue]
    : ('!' | 'not') primaryExpr         { $retValue = opt.NewNotUnaryExpression("!", $primaryExpr.retValue)}
    | '~' primaryExpr                   { $retValue = opt.NewBitwiseNotUnaryExpression("~", $primaryExpr.retValue)}
    | '-' primaryExpr                   { $retValue = opt.NewNegateUnaryExpression("-", $primaryExpr.retValue)}
    | primaryExpr                       { $retValue = $primaryExpr.retValue}
    ;

primaryExpr returns[*opt.LogicalExpression retValue]
    @init { var args = make([]*opt.LogicalExpression,0)}
    : '(' expr ')'                      { $retValue = $expr.retValue}
    | value                             { $retValue = $value.retValue}
    | id 
        '(' expr                        { args = append(args, $expr.retValue)}
            (',' expr                   { args = append(args, $expr.retValue)}
            )*
        ')'                             { $retValue = opt.NewFunctionExpression(opt.NewIdentifier($id.text), args)}
    | id                                { $retValue = $id.retValue}
    ;

//值类型
value returns[*opt.LogicalExpression retValue]
    : INTEGER     { $retValue = opt.NewIntegerValueExpression($INTEGER.text)}
    | FLOAT       { $retValue = opt.NewFloatExpression($FLOAT.text)}
    | STRING      { $retValue = opt.NewStringValueExpression($STRING.text)}
    | DATETIME    { $retValue = opt.NewDateTimeExpression($DATETIME.text)}
    | TRUE        { $retValue = opt.NewBooleanValueExpression(true)}
    | FALSE       { $retValue = opt.NewBooleanValueExpression(false)}
    ;

id returns[*opt.LogicalExpression retValue]
    : NAME { $retValue = opt.NewIdentifierExpression($NAME.text)}
    | VAR  { $retValue = opt.NewVarIdentifierExpression($VAR.text)}
    ;

TRUE    : 'true'  ;
FALSE   : 'false' ;
NAME    : LETTER (LETTER | DIGIT)* ;
INTEGER : DIGIT+  ;
DATETIME: '#' (~('#')*) '#' ;
VAR     : '[' (~(']')*) ']' ;
E       : ('E'|'e') ('+'|'-')? DIGIT+ ;

FLOAT
    : DIGIT* '.' DIGIT+ E?
    | DIGIT+ E
    ;
    
STRING
    : '\'' 
      ( EscapeSequence 
      | (~('\u0000'..'\u001f' | '\\' | '\''))
      )*
      '\''
    ;

fragment LETTER
    : 'a'..'z'
    | 'A'..'Z'
    | '_'
    ;

fragment DIGIT : '0'..'9';

fragment EscapeSequence 
    : '\\'
        ( 'n'
        | 'r'
        | 't'
        | '\''
        | '\\'
        | UnicodeEscape
        )
    ;

fragment HexDigit : ('0'..'9'|'a'..'f'|'A'..'F');

fragment UnicodeEscape : 'u' HexDigit HexDigit HexDigit HexDigit;

/* Ignore white spaces */	
WS : (' '|'\r'|'\t'|'\u000C'|'\n') -> skip;
