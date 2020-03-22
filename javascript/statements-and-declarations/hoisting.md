# Hoisting

## Reference

- https://developer.mozilla.org/en-US/docs/Glossary/Hoisting
- https://blog.techbridge.cc/2018/11/10/javascript-hoisting/

## Summary

- 宣告會提升
- 賦值不會提升

## Example 1

```javascript
console.log(data); // ReferenceError: data is not defined
```

## Example 2

```javascript
console.log(data); // ReferenceError: Cannot access 'data' before initialization
let data = '';
```

## Example 3

```javascript
console.log(data); // undefined
var data = '';
```

## Example 4

```javascript
var x = 1; // Initialize x
console.log(x, y); // '1 undefined'
var y = 2; // Initialize y
// JavaScript only hoists declarations
```