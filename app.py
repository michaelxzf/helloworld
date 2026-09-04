import os

API_KEY = os.getenv("API_KEY")  # credential loaded from environment variable

try:
    number = int(input("Enter a number: "))
except ValueError:
    print("Invalid input: please enter a valid integer.")
else:
    print(number * 2)
