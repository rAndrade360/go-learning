from kafka import KafkaProducer
from kafka.errors import KafkaError

producer = KafkaProducer(bootstrap_servers=['localhost:9092'])

for _ in range(100):
    future = producer.send('kafka-labs', b'raw_bytes')
    try:
        record_metadata = future.get(timeout=10)
    except KafkaError:
        pass
    print (record_metadata.topic)
    print (record_metadata.partition)
    print (record_metadata.offset)