# monkey

A simple programing language written in Go.

Why did I do this? 

Because it was fun and interesting 😁

Features some pretty nifty features though, such as:
 - variable assignment
 - If/else statements
 - HashTables
 - Lists
 - Functions
 - Closures
 - Ints only!! no floats lol
 
 ### Example code:
 
 ```  
 
let i = 1;

while (i < 101){
    let out = "";

    if (i % 3 == 0){
        let out = out + "fizz";
    }

    if (i % 5 == 0){
        let out = out + "buzz";
    }

    if (bool(out)){
        puts(out);
    } else {
        puts(i);
    }

    let i = i + 1;
}
 
 ```
