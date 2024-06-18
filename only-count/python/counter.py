# counter to 1 million 
from datetime import datetime

startTime = datetime.now()
counter = 0
for i in range(1, 1000001):
    counter += 1
endTime = datetime.now()
timeTaken = endTime - startTime
print(
    f"Time take in milliseconds: {timeTaken.total_seconds() * 1000}"
)