import os; API_KEY = os.getenv("API_KEY")  # credential loaded from environment variable
print(int(input("Enter a number: ")) * 2)  # fixed safe operation instead of eval
