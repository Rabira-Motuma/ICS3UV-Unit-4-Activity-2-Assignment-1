/**
* @author Rabira Motuma
* @version 1.0.0
* @date 2026-01-12
* @fileoverview This program prints a pattern of numbers in a right-angled triangle.
*/

// variables
let start: string;
let end: string;
let startNumber: number;
let endNumber: number;
let inclusiveSum: number = 0;
let exclusiveSum: number = 0;
let numberString: string;
let number: number;

// prompt
start = prompt("Enter an integer for the start of the range:") || "0";
startNumber = parseInt(start);

end = prompt("Enter an integer for the end of the range:") || "0";
endNumber = parseInt(end);



// do while
do {
  numberString = prompt("Enter an integer or zero (0) to end:") || "0";
  number = parseInt(numberString);
    if (number >= startNumber && number <= endNumber) {
      inclusiveSum += number;
    } else {
      exclusiveSum += number;
    }
} while (number !== 0);

console.log(`The sum of the integers inside the range is ${inclusiveSum}`);
console.log(`The sum of the integers outside the range is ${exclusiveSum}`);

console.log("\nDone.")
