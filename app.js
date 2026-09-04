"use strict";

const fs = require("node:fs");
const path = require("node:path");

// credential loaded from environment variable
const apiKey = process.env.API_KEY;

function double(number) {
  return number * 2;
}

function divide(a, b) {
  if (b === 0) {
    throw new RangeError("division by zero");
  }
  return a / b;
}

// Reads a file, restricting access to paths inside the current working
// directory to prevent path traversal.
function readFile(filePath) {
  const root = path.resolve(process.cwd());
  const resolved = path.resolve(root, filePath);
  if (!resolved.startsWith(root + path.sep)) {
    throw new Error(`path is outside the working directory: ${filePath}`);
  }
  return fs.readFileSync(resolved, "utf8");
}

module.exports = { double, divide, readFile };
