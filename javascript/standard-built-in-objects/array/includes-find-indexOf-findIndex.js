let array = [];

// Determines whether an array includes a certain element, returning true or false as appropriate.
var included = array.includes(2);
// console.log(included);

// Returns the value of the first element in the array where predicate is true,
// and undefined otherwise.
var found = array.find(function(element) {
    return element > 1;
});
// console.log(found);

// Returns the index of the first occurrence of a value in an array.
// or -1
var index = array.indexOf(3);
// console.log(index);

// Returns the index of the first element in the array where predicate is true, 
// and -1 otherwise.
var index = array.findIndex(function(element) {
    return element > 1;
});
// console.log(index);