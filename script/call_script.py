import requests
import json
import time
from datetime import datetime, timedelta
import random

# Define variations for each field
levels = ["info", "warning", "error", "critical", "debug"]
messages = ["Connection established", "Failed to connect to DB", "Resource not found", "User not authenticated", "User not authorized", "Invalid request", "Invalid response", "Request timed out"]
resource_ids = ["server-1234", "server-5678", "server-91011", "server-121314", "server-151617", "server-181920"]
trace_ids = ["abc-xyz-123", "def-uvw-456", "ghi-rst-789", "jkl-mno-101112", "pqr-stu-131415", "vwx-efg-161718"]
span_ids = ["span-123", "span-456", "span-789", "span-101112", "span-131415", "span-161718"]
commits = ["5e5342f", "a1b2c3d", "e4f5g6h", "i7j8k9l", "m0n1o2p", "q3r4s5t"]
metadata_values = ["{'key': 'value'}", "{'type': 'log'}", "{'source': 'app'}", "{'source': 'server'}", "{'source': 'client'}", "{'source': 'database'}"]

# Generate timestamps for the past 6 months
end_time = datetime.utcnow()
start_time = end_time - timedelta(days=180)

# Function to generate a random timestamp within the specified range
def generate_random_timestamp():
    return start_time + timedelta(seconds=random.randint(0, int((end_time - start_time).total_seconds())))

# API endpoint
api_url = "http://localhost:3000/log"

# Number of requests
num_requests = 500

for _ in range(num_requests):
    data = {
        "level": random.choice(levels),
        "message": random.choice(messages),
        "resourceId": random.choice(resource_ids),
        "timestamp": generate_random_timestamp().strftime("%Y-%m-%dT%H:%M:%SZ"),
        "traceId": random.choice(trace_ids),
        "spanId": random.choice(span_ids),
        "commit": random.choice(commits),
        "metadata": random.choice(metadata_values)
    }

    # Convert data to JSON
    json_data = json.dumps(data)

    # Make a POST request to the API
    response = requests.post(api_url, data=json_data, headers={"Content-Type": "application/json"})

    # Print the response status code and sleep for 10ms
    print(f"Status Code: {response.status_code}")
    time.sleep(0.01)
