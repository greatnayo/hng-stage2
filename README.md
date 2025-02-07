**Introduction**
This is api project that performs number classification and also fetches some intresting fun fact about the given number from the number API .

**Resources Used**
Fun fact API: http://numbersapi.com/#42
https://en.wikipedia.org/wiki/Parity_(mathematics)

**Given Task**
Create an API that takes a number and returns interesting mathematical properties about it, along with a fun fact.

**API Specification**
Endpoint: **GET** <your-url>/api/classify-number?number=371
Required JSON Response Format (200 OK):
{
"number": 371,
"is_prime": false,
"is_perfect": false,
"properties": ["armstrong", "odd"],
"digit_sum": 11, // sum of its digits
"fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
Required JSON Response Format (400 Bad Request)
{
"number": "alphabet",
"error": true
}
