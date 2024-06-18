# counter to 1 million 
from datetime import datetime

startTime = datetime.now()

for i in range(1, 1000001):
    print(i)
endTime = datetime.now()
timeTaken = endTime - startTime
print(
    f"Time take in seconds: {timeTaken.total_seconds()}"
)