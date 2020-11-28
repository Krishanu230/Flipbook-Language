# Flipbook-Language
a language for describing flipbooks and  a compiler  for this language that can convert a flipbook description into a print- able pdf

## Getting Started
0. Get the Dependecies: [gopdf](https://github.com/signintech/gopdf)
```
go get -u github.com/signintech/gopdf
```
1. Get the code
```
go get github.com/Krishanu230/Flipbook-Language
```
2. CD into the GOPATH src folder. It is usually:
```
cd ~/go/src/github.com/Krishanu230/Flipbook-Language
```
3. Go run it on any flipbook file.
```
go run main.go flipbook.flip 
```

## Language Design

### Rules:
1. There are two type of objects in this language: Book and Image.

2. Modifiable Properties of Objects:

    2.1 Book: it has no public properties that can be modified with keyframe statement
    
    2.2 Image: It has three public properties that can be modified via Keyframe: Scale, PositionX and PositionY
    
3. DataPrimitives: 

    3.1 Strings: Anything in double quotes is string.
    
    3.2 Integers: They are any positive integer.
    
4. There are four types of Statements implemented right now: new, insert, keyframe, save. Every Statement must end with a semicolon.

    4.1 ***New***: It creates a new object
    Syntax - 
    
    ***new [AnyObjType] *variable name* = ((dimesionX_INT, dimensionY_INT), AttributeOfTheObject)***
    
    where the attributeOfTheObject is page number count if the onject is a BookType or a StringType containing the filepath if the object is imageType.
    
    4.2 ***Insert***: It inserts a image in a book
    Syntax - 
    
    ***insert [Page ObjType variable] [Book ObjType variable] from page StartPage_Numeber_INT to EndPageNumber_INT  at (positionX_INT, positionY_INT)***
    
    4.3 ***Keyframe***: It modifies some modifiable property of some object over a range of pages in a liner way. Currently only images properties can be modified.
    Syntax - 
    
    ***keyframe [Page ObjType variable] [Book ObjType variable] [Property Name Keyword] (startPageNumber, startPropertyValue) to (endPageNumber, endPropertyValue)***
    Note that only the pages that contain the specified image will be changed.
    
    4.4  ***Save***: It takes a book and renders it as a pdf.
    Syntax - 
    ***save [Book ObjType variable] outputNameString***

## Example program
Lets animate apple falling over newton's head.
```
new book bookone = ((1600,1600), 25);
new image newton = ((100,100), "newton.png");
insert newton bookone from page 1 to 25 at (600,1270);
new image melon = ((100,100), "melon.png");
insert melon bookone from page 1 to 25 at (600,0);
keyframe melon bookone positionY (1,0) to (25, 1000);
save bookone "out.pdf";
```

Lets Run it:
[![Final output]](https://www.youtube.com/watch?v=ulpEuGnCMP8)
[![Walkthrough Demo]](https://www.youtube.com/watch?v=U1MMO5FGT9Q)

KeyWords= {	new, at, to, set, image, book, scale, insert, from, page, keyframe,, save, positionX, positionY}

##Points to note
1. Objective of the language: The language should be english like and intutive, animation properties should be flexible. Should easily be extensible.
2. Why not a parser and lexer generator?: I assumed that the point of this excercise is to see us code a framework. Parsing and lexical analysis are the most intresting part and using a generator would I think defeat the purpose of the excercise. Furthermore, I need to learn a different schema to write my specification if I want to use a generator. I'd first try to get my hands dirty with making parsers and lexical analiser by myself.

##Extending the language:
1. Adding more lexical rules like allowing characters etc: First register the toke in token.go and add the rules in lexer.go
2. Adding More Statements: First define the ast node for your new statement in ast.go, then add the parsing rules in parser.go and finally add the evaluation logic in evaluator.go. If required you can also create new object types in object.go

##ToDo:
1. Compelete the document: Due to the limitation of time, The documentation is not very resourceful. 
2.Clean Up the Code: Due to the same reason, I couldnt get the chance to properly refactor code and comment it. 
3. Add Expression pevaluation via Pratt's Parser technique
4. Add an analouge of Functions as Effects like swirl
5. Add more properties like opacity and more statements like 'set'.
