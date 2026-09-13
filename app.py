import os

API_KEY = os.getenv("API_KEY")  # credential loaded from environment variable

WORKSPACE_ROOT = os.path.dirname(os.path.abspath(__file__))


def double(number):
    return number * 2


def divide(a, b):
    if b == 0:
        raise ValueError("division by zero is not allowed")
    return a / b


def read_file(path):
    root = os.path.realpath(WORKSPACE_ROOT)
    resolved = os.path.realpath(os.path.join(root, path))
    if resolved != root and not resolved.startswith(root + os.sep):
        raise ValueError("path is outside the workspace directory")
    with open(resolved) as f:
        return f.read()


if __name__ == "__main__":
    print(double(int(input("Enter a number: "))))
