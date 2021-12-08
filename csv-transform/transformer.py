"""Basic lambda to transform csv file data to DB records"""

import csv
import tempfile
import urllib

import boto3

s3 = boto3.client('s3')
dynamodb = boto3.resource('dynamodb')

def transform_csv(event, context):
    """Main lambda logic"""
    bucket = event['Records'][0]['s3']['bucket']['name']
    key = urllib.parse.unquote_plus(event['Records'][0]['s3']['object']['key'], encoding='utf-8')

    try:
        items = list()
        with tempfile.NamedTemporaryFile() as temp_file:
            # Download and parse the csv
            s3.download_file(bucket, key, temp_file.name)
            reader = csv.reader(temp_file)
            for row in reader:
                first_name, last_name, age, title = row
                record = {'first_name': first_name, 'last_name': last_name, 'title': title, 'age': age}
                items.append(record)

        # Now that we have all of the stuff in memory, we can just push that to the table.

        people_table = dynamodb.Table("People")

        for record in items:
            people_table.put_item(Table="People", Item=record)

    except Exception as exception:
        print(exception)
        raise exception

    return 1
