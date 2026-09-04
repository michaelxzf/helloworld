import os

API_KEY = os.getenv("API_KEY")  # credential loaded from environment variable


def double(number):
    return number * 2


def divide(a, b):
    return a / b


if __name__ == "__main__":
    print(double(int(input("Enter a number: "))))
