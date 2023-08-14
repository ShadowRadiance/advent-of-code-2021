import run from "aocrunner";
import { test } from "node:test";

const parseInput = (rawInput) => rawInput.split("\n").map((line)=>parseInt(line));

const parseElves = (numbers) => {
  const elves = [];
  let currentSum = 0;

  for (const number of numbers) {
    if (isNaN(number)) {
      elves.push(currentSum);
      currentSum = 0;
    } else {
      currentSum += number;
    }
  }
  elves.push(currentSum);
  return elves;
};

const part1 = (rawInput) => {
  const elves = parseElves(parseInput(rawInput));
  elves.sort((a, b) => a - b);
  return `${elves.pop()}`;
};

const part2 = (rawInput) => {
  const elves = parseElves(parseInput(rawInput));
  elves.sort((a, b) => a - b);
  const best_three = elves.slice(-3);
  const sum = best_three.reduce((prev,curr,currIdx,arr)=>{ return prev + curr; });
  return `${sum}`;
};

const testInput = `
  1000
  2000
  3000

  4000

  5000
  6000

  7000
  8000
  9000

  10000
`;
run({
  part1: {
    tests: [
      {
        input: testInput,
        expected: "24000",
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: testInput,
        expected: "45000",
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
