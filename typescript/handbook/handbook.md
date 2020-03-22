# handbook

## Basic Types

- Boolean
- Number
- String
- Array
- Tuple
- Enum
- Any
- Void
- Null
- Undefined
- Never
- Object

```typescript
let isDone: boolean = false;

let decimal: number = 6;
let hex: number = 0xf00d;
let octal: number = 0o744;
let binary: number = 0b1010;

let color: string = "blue";
color = "red";

let fullName: string = `Bob Bobbington`;
let age: number = 37;
let sentence: string = `Hello, my name is ${fullName}.
I'll be ${age + 1} years old next month.`;

let list: number[] = [1, 2, 3];
let list: Array<number> = [1, 2, 3];

let x: [string, number];
x = ["hello", 10];
console.log(x[0].substring(1));

enum Color {
  Red,
  Green,
  Blue = 4
}
let c: Color = Color.Green;
let colorName: string = Color[2];
console.log(colorName);
console.log(Color[colorName]);

let notSure: any = 4;
notSure = "maybe a string instead";
notSure = false;

let notSure: any = 4;
let prettySure: Object = 4;
notSure.ifItExists(); // okay, ifItExists might exist at runtime
notSure.toFixed(); // okay, toFixed exists (but the compiler doesn't check)
prettySure.toFixed(); // Error: Property 'toFixed' doesn't exist on type 'Object'.

let list: any[] = [1, true, "free"];
list[1] = 100;

https://www.typescriptlang.org/docs/handbook/basic-types.html#void
```

### Type assertions
