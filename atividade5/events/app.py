from flask import Flask, request, jsonify
import logging

app = Flask(__name__)


logging.basicConfig(filename='events.log', level=logging.INFO)

events = []

@app.route('/events', methods=['POST'])
def create_event():
    event = request.json
    events.append(event)
    logging.info(f"Event created: {event}")
    return jsonify(event), 201

@app.route('/events', methods=['GET'])
def get_events():
    return jsonify(events), 200

@app.route('/events/<int:event_id>', methods=['PUT'])
def update_event(event_id):
    event = request.json
    events[event_id] = event
    logging.info(f"Event updated: {event}")
    return jsonify(event), 200

@app.route('/events/<int:event_id>', methods=['DELETE'])
def delete_event(event_id):
    event = events.pop(event_id)
    logging.info(f"Event deleted: {event}")
    return '', 204

if __name__ == '__main__':
    app.run(port=5000)
