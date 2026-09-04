"use strict";

const readline = require("node:readline/promises");
const { double } = require("./app");

async function main() {
  console.log("Hello, World!");

  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  });

  try {
    let answer;
    try {
      answer = await rl.question("Enter a number: ");
    } catch {
      console.log("Invalid input: please enter a valid integer.");
      return;
    }

    const trimmed = typeof answer === "string" ? answer.trim() : "";
    if (/^[+-]?\d+$/.test(trimmed)) {
      console.log(double(Number.parseInt(trimmed, 10)));
    } else {
      console.log("Invalid input: please enter a valid integer.");
    }
  } finally {
    rl.close();
  }
}

main().catch((err) => {
  console.error(`Error: ${err.message}`);
  process.exitCode = 1;
});
