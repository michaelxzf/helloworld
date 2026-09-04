"use strict";

const test = require("node:test");
const assert = require("node:assert/strict");
const fs = require("node:fs");
const path = require("node:path");

const { double, divide, readFile } = require("./app");

test("double doubles a positive number", () => {
  assert.equal(double(2), 4);
});

test("double doubles zero", () => {
  assert.equal(double(0), 0);
});

test("double doubles a negative number", () => {
  assert.equal(double(-3), -6);
});

test("divide divides positive numbers", () => {
  assert.equal(divide(6, 3), 2);
});

test("divide divides negative numbers", () => {
  assert.equal(divide(-6, 3), -2);
});

test("divide returns fractional results", () => {
  assert.equal(divide(1, 4), 0.25);
});

test("divide rejects division by zero", () => {
  assert.throws(() => divide(1, 0), RangeError);
});

function withTempFile(name, contents, run) {
  const filePath = path.join(process.cwd(), name);
  fs.writeFileSync(filePath, contents);
  try {
    run(filePath);
  } finally {
    fs.rmSync(filePath, { force: true });
  }
}

test("readFile returns file contents", () => {
  withTempFile("readfile-test-example.txt", "hello world", (filePath) => {
    assert.equal(readFile(filePath), "hello world");
  });
});

test("readFile returns an empty string for an empty file", () => {
  withTempFile("readfile-test-empty.txt", "", (filePath) => {
    assert.equal(readFile(filePath), "");
  });
});

test("readFile preserves multiline contents", () => {
  withTempFile("readfile-test-lines.txt", "first\nsecond\n", (filePath) => {
    assert.equal(readFile(filePath), "first\nsecond\n");
  });
});

test("readFile rejects paths outside the working directory", () => {
  assert.throws(() => readFile(path.join("..", "outside.txt")), Error);
});
