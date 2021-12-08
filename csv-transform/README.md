# Lambda-based S3 -> DB CSV transformation

Bear in mind, I had exactly 0 serverless experience prior to this exercise. I'm going with Python for this one, as it
has a very nice CSV support, I can write a very succint lambda, and AWS considers it a first-class citizen.

I will expect only proper CSV files to be uploaded to the S3 bucket this lambda will be attached to. I will also expect
all of them to conform to the proper format of `first_name,last_name,age,title`. If these would fail, I would implement
the error handling - checking if the extension is proper, if file is readable, etc. I'm also very well aware of the fact
that I'm reading the entire file contents into the memory, but I also did not want to spend time optimizing that. One
very simple optimization would be to move the DB item submission pretty much next to the parsing, but that will entangle
code a bit more then I'm comfortable with.

I'm also picking DynamoDB as the target database, due to ease of setup -- I'll try to use CDK to set all of the needed
things up.

I did stumble a bit on:

- tempfiles in python - constantly forget how those work, unforutantely. Settled on just having a hardcoded target, as
  the environment is anyway transient
- policy definition is inlined for the lambda, though I would create a more durable solution in a production scenario
